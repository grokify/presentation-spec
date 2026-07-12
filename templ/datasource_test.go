package templ

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/grokify/presentation-spec/spec"
)

func TestDataLoader_LoadJSONFile(t *testing.T) {
	// Create temp directory and JSON file
	tmpDir := t.TempDir()
	jsonPath := filepath.Join(tmpDir, "data.json")

	testData := map[string]any{
		"name":  "test",
		"value": float64(42),
	}
	jsonBytes, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("failed to marshal test data: %v", err)
	}

	if err := os.WriteFile(jsonPath, jsonBytes, 0600); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	loader := NewDataLoader(tmpDir)
	ctx := context.Background()

	src := spec.DataSource{
		ID:   "test-json",
		Type: spec.DataSourceTypeJSON,
		Path: "data.json",
	}

	loaded, err := loader.Load(ctx, src)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if loaded.ID != "test-json" {
		t.Errorf("expected ID 'test-json', got %q", loaded.ID)
	}

	data, ok := loaded.Data.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any, got %T", loaded.Data)
	}

	if data["name"] != "test" {
		t.Errorf("expected name 'test', got %v", data["name"])
	}

	if data["value"] != float64(42) {
		t.Errorf("expected value 42, got %v", data["value"])
	}
}

func TestDataLoader_LoadCSVFile(t *testing.T) {
	tmpDir := t.TempDir()
	csvPath := filepath.Join(tmpDir, "data.csv")

	csvContent := `name,value,status
Alice,100,active
Bob,200,inactive`

	if err := os.WriteFile(csvPath, []byte(csvContent), 0600); err != nil {
		t.Fatalf("failed to write CSV file: %v", err)
	}

	loader := NewDataLoader(tmpDir)
	ctx := context.Background()

	src := spec.DataSource{
		ID:   "test-csv",
		Type: spec.DataSourceTypeCSV,
		Path: "data.csv",
	}

	loaded, err := loader.Load(ctx, src)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	data, ok := loaded.Data.([]map[string]string)
	if !ok {
		t.Fatalf("expected []map[string]string, got %T", loaded.Data)
	}

	if len(data) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(data))
	}

	if data[0]["name"] != "Alice" {
		t.Errorf("expected first name 'Alice', got %q", data[0]["name"])
	}

	if data[1]["value"] != "200" {
		t.Errorf("expected second value '200', got %q", data[1]["value"])
	}
}

func TestDataLoader_LoadAPI(t *testing.T) {
	testData := map[string]any{
		"status": "ok",
		"count":  float64(10),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check auth header
		if auth := r.Header.Get("Authorization"); auth != "Bearer test-token" {
			t.Errorf("expected Bearer auth, got %q", auth)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(testData); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	loader := NewDataLoader("")
	ctx := context.Background()

	src := spec.DataSource{
		ID:       "test-api",
		Type:     spec.DataSourceTypeAPI,
		Endpoint: server.URL,
		Auth: &spec.DataSourceAuth{
			Type:  spec.AuthTypeBearer,
			Token: "test-token",
		},
	}

	loaded, err := loader.Load(ctx, src)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	data, ok := loaded.Data.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any, got %T", loaded.Data)
	}

	if data["status"] != "ok" {
		t.Errorf("expected status 'ok', got %v", data["status"])
	}
}

func TestDataLoader_LoadAPIWithBasicAuth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != "testuser" || pass != "testpass" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"authenticated": true}`))
	}))
	defer server.Close()

	loader := NewDataLoader("")
	ctx := context.Background()

	src := spec.DataSource{
		ID:       "test-basic-auth",
		Type:     spec.DataSourceTypeAPI,
		Endpoint: server.URL,
		Auth: &spec.DataSourceAuth{
			Type:     spec.AuthTypeBasic,
			Username: "testuser",
			Password: "testpass",
		},
	}

	loaded, err := loader.Load(ctx, src)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	data, ok := loaded.Data.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any, got %T", loaded.Data)
	}

	if data["authenticated"] != true {
		t.Errorf("expected authenticated true, got %v", data["authenticated"])
	}
}

func TestDataLoader_LoadAPIWithAPIKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Custom-Key")
		if apiKey != "my-api-key" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"valid": true}`))
	}))
	defer server.Close()

	loader := NewDataLoader("")
	ctx := context.Background()

	src := spec.DataSource{
		ID:       "test-api-key",
		Type:     spec.DataSourceTypeAPI,
		Endpoint: server.URL,
		Auth: &spec.DataSourceAuth{ //nolint:gosec // Test credentials
			Type:         spec.AuthTypeAPIKey,
			APIKey:       "my-api-key",
			APIKeyHeader: "X-Custom-Key",
		},
	}

	loaded, err := loader.Load(ctx, src)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	data, ok := loaded.Data.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any, got %T", loaded.Data)
	}

	if data["valid"] != true {
		t.Errorf("expected valid true, got %v", data["valid"])
	}
}

func TestDataLoader_LoadAll(t *testing.T) {
	tmpDir := t.TempDir()

	// Create two JSON files
	if err := os.WriteFile(filepath.Join(tmpDir, "data1.json"), []byte(`{"id": 1}`), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "data2.json"), []byte(`{"id": 2}`), 0600); err != nil {
		t.Fatal(err)
	}

	loader := NewDataLoader(tmpDir)
	ctx := context.Background()

	sources := []spec.DataSource{
		{ID: "src1", Type: spec.DataSourceTypeJSON, Path: "data1.json"},
		{ID: "src2", Type: spec.DataSourceTypeJSON, Path: "data2.json"},
	}

	store, err := loader.LoadAll(ctx, sources)
	if err != nil {
		t.Fatalf("LoadAll failed: %v", err)
	}

	if _, ok := store.Get("src1"); !ok {
		t.Error("expected src1 in store")
	}

	if _, ok := store.Get("src2"); !ok {
		t.Error("expected src2 in store")
	}

	if len(store.All()) != 2 {
		t.Errorf("expected 2 items in store, got %d", len(store.All()))
	}
}

func TestDataLoader_InlineType(t *testing.T) {
	loader := NewDataLoader("")
	ctx := context.Background()

	src := spec.DataSource{
		ID:   "inline-data",
		Type: spec.DataSourceTypeInline,
	}

	loaded, err := loader.Load(ctx, src)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if loaded.ID != "inline-data" {
		t.Errorf("expected ID 'inline-data', got %q", loaded.ID)
	}

	// Inline type should have nil Data (data is in spec itself)
	if loaded.Data != nil {
		t.Errorf("expected nil Data for inline type, got %v", loaded.Data)
	}
}

func TestDataLoader_FileNotFound(t *testing.T) {
	loader := NewDataLoader("/nonexistent")
	ctx := context.Background()

	src := spec.DataSource{
		ID:   "missing",
		Type: spec.DataSourceTypeJSON,
		Path: "missing.json",
	}

	_, err := loader.Load(ctx, src)
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestDataLoader_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmpDir, "invalid.json"), []byte(`{invalid`), 0600); err != nil {
		t.Fatal(err)
	}

	loader := NewDataLoader(tmpDir)
	ctx := context.Background()

	src := spec.DataSource{
		ID:   "invalid",
		Type: spec.DataSourceTypeJSON,
		Path: "invalid.json",
	}

	_, err := loader.Load(ctx, src)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestDataLoader_UnsupportedType(t *testing.T) {
	loader := NewDataLoader("")
	ctx := context.Background()

	src := spec.DataSource{
		ID:   "unsupported",
		Type: spec.DataSourceTypeDatabase,
	}

	_, err := loader.Load(ctx, src)
	if err == nil {
		t.Error("expected error for unsupported database type")
	}
}
