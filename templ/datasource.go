package templ

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/grokify/presentation-spec/spec"
)

// DataLoader loads data from configured data sources.
type DataLoader struct {
	// BaseDir is the base directory for resolving relative file paths.
	BaseDir string

	// HTTPClient is the HTTP client to use for API requests.
	HTTPClient *http.Client

	// Timeout is the default timeout for requests.
	Timeout time.Duration
}

// NewDataLoader creates a new DataLoader with default settings.
func NewDataLoader(baseDir string) *DataLoader {
	return &DataLoader{
		BaseDir: baseDir,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Timeout: 30 * time.Second,
	}
}

// LoadedData holds the data loaded from a data source.
type LoadedData struct {
	// ID is the data source ID.
	ID string

	// Data is the loaded data (typically map[string]any or []any).
	Data any

	// Raw is the raw bytes of the loaded data.
	Raw []byte

	// LoadedAt is when the data was loaded.
	LoadedAt time.Time

	// Source is the original data source configuration.
	Source spec.DataSource
}

// DataStore holds loaded data indexed by source ID.
type DataStore struct {
	data map[string]*LoadedData
}

// NewDataStore creates an empty DataStore.
func NewDataStore() *DataStore {
	return &DataStore{
		data: make(map[string]*LoadedData),
	}
}

// Get retrieves loaded data by source ID.
func (ds *DataStore) Get(id string) (*LoadedData, bool) {
	d, ok := ds.data[id]
	return d, ok
}

// Set stores loaded data by source ID.
func (ds *DataStore) Set(id string, data *LoadedData) {
	ds.data[id] = data
}

// All returns all loaded data.
func (ds *DataStore) All() map[string]*LoadedData {
	return ds.data
}

// LoadAll loads data from all data sources in a presentation spec.
func (dl *DataLoader) LoadAll(ctx context.Context, sources []spec.DataSource) (*DataStore, error) {
	store := NewDataStore()

	for _, src := range sources {
		data, err := dl.Load(ctx, src)
		if err != nil {
			return store, fmt.Errorf("failed to load data source %q: %w", src.ID, err)
		}
		store.Set(src.ID, data)
	}

	return store, nil
}

// Load loads data from a single data source.
func (dl *DataLoader) Load(ctx context.Context, src spec.DataSource) (*LoadedData, error) {
	var data any
	var raw []byte
	var err error

	switch src.Type {
	case spec.DataSourceTypeInline:
		// Inline data is already in the spec, nothing to load
		return &LoadedData{
			ID:       src.ID,
			LoadedAt: time.Now(),
			Source:   src,
		}, nil

	case spec.DataSourceTypeFile, spec.DataSourceTypeJSON:
		data, raw, err = dl.loadJSONFile(src.Path)

	case spec.DataSourceTypeCSV:
		data, raw, err = dl.loadCSVFile(src.Path)

	case spec.DataSourceTypeAPI:
		data, raw, err = dl.loadAPI(ctx, src)

	case spec.DataSourceTypeDatabase:
		return nil, fmt.Errorf("database data source type not yet supported")

	default:
		return nil, fmt.Errorf("unknown data source type: %s", src.Type)
	}

	if err != nil {
		return nil, err
	}

	return &LoadedData{
		ID:       src.ID,
		Data:     data,
		Raw:      raw,
		LoadedAt: time.Now(),
		Source:   src,
	}, nil
}

func (dl *DataLoader) loadJSONFile(path string) (any, []byte, error) {
	fullPath := dl.resolvePath(path)

	raw, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read JSON file %q: %w", fullPath, err)
	}

	var data any
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, raw, fmt.Errorf("failed to parse JSON file %q: %w", fullPath, err)
	}

	return data, raw, nil
}

func (dl *DataLoader) loadCSVFile(path string) (any, []byte, error) {
	fullPath := dl.resolvePath(path)

	f, err := os.Open(fullPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open CSV file %q: %w", fullPath, err)
	}
	defer func() { _ = f.Close() }()

	// Read all content for Raw
	raw, err := io.ReadAll(f)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CSV file %q: %w", fullPath, err)
	}

	// Re-open for CSV parsing
	f2, err := os.Open(fullPath)
	if err != nil {
		return nil, raw, fmt.Errorf("failed to reopen CSV file %q: %w", fullPath, err)
	}
	defer func() { _ = f2.Close() }()

	reader := csv.NewReader(f2)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, raw, fmt.Errorf("failed to parse CSV file %q: %w", fullPath, err)
	}

	// Convert to slice of maps using first row as headers
	if len(records) == 0 {
		return []map[string]string{}, raw, nil
	}

	headers := records[0]
	var data []map[string]string

	for i := 1; i < len(records); i++ {
		row := make(map[string]string)
		for j, header := range headers {
			if j < len(records[i]) {
				row[header] = records[i][j]
			}
		}
		data = append(data, row)
	}

	return data, raw, nil
}

func (dl *DataLoader) loadAPI(ctx context.Context, src spec.DataSource) (any, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, src.Endpoint, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request for %q: %w", src.Endpoint, err)
	}

	// Add custom headers
	for k, v := range src.Headers {
		req.Header.Set(k, v)
	}

	// Add authentication
	if src.Auth != nil {
		switch src.Auth.Type {
		case spec.AuthTypeBearer:
			req.Header.Set("Authorization", "Bearer "+src.Auth.Token)
		case spec.AuthTypeBasic:
			req.SetBasicAuth(src.Auth.Username, src.Auth.Password)
		case spec.AuthTypeAPIKey:
			header := src.Auth.APIKeyHeader
			if header == "" {
				header = "X-API-Key"
			}
			req.Header.Set(header, src.Auth.APIKey)
		}
	}

	resp, err := dl.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch %q: %w", src.Endpoint, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, nil, fmt.Errorf("API returned status %d for %q", resp.StatusCode, src.Endpoint)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response from %q: %w", src.Endpoint, err)
	}

	var data any
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, raw, fmt.Errorf("failed to parse JSON response from %q: %w", src.Endpoint, err)
	}

	return data, raw, nil
}

func (dl *DataLoader) resolvePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(dl.BaseDir, path)
}
