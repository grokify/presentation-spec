package spec

// DataSource defines an external data source referenced by widgets.
type DataSource struct {
	// ID is the unique identifier for this data source.
	ID string `json:"id"`

	// Type is the data source type.
	Type DataSourceType `json:"type"`

	// Path is the file path (for file-based sources).
	Path string `json:"path,omitempty"`

	// Endpoint is the API endpoint URL (for API sources).
	Endpoint string `json:"endpoint,omitempty"`

	// Query is the query string (for database sources).
	Query string `json:"query,omitempty"`

	// Connection is the connection string or name (for database sources).
	Connection string `json:"connection,omitempty"`

	// RefreshInterval is how often to refresh the data (e.g., "5m", "1h").
	RefreshInterval string `json:"refreshInterval,omitempty"`

	// Cache controls whether to cache the data.
	Cache bool `json:"cache,omitempty"`

	// Transform is a jq-style transformation to apply to the data.
	Transform string `json:"transform,omitempty"`

	// Headers are HTTP headers to include (for API sources).
	Headers map[string]string `json:"headers,omitempty"`

	// Auth specifies authentication configuration.
	Auth *DataSourceAuth `json:"auth,omitempty"`
}

// DataSourceType enumerates data source types.
type DataSourceType string

const (
	// DataSourceTypeInline indicates data is embedded in the spec.
	DataSourceTypeInline DataSourceType = "inline"

	// DataSourceTypeFile indicates data is in a local file.
	DataSourceTypeFile DataSourceType = "file"

	// DataSourceTypeAPI indicates data is from an API endpoint.
	DataSourceTypeAPI DataSourceType = "api"

	// DataSourceTypeDatabase indicates data is from a database.
	DataSourceTypeDatabase DataSourceType = "database"

	// DataSourceTypeCSV indicates data is from a CSV file.
	DataSourceTypeCSV DataSourceType = "csv"

	// DataSourceTypeJSON indicates data is from a JSON file.
	DataSourceTypeJSON DataSourceType = "json"
)

// DataSourceAuth specifies authentication for a data source.
type DataSourceAuth struct {
	// Type is the authentication type.
	Type AuthType `json:"type"`

	// Token is the bearer token (for bearer auth).
	Token string `json:"token,omitempty"`

	// Username is the username (for basic auth).
	Username string `json:"username,omitempty"`

	// Password is the password (for basic auth).
	Password string `json:"password,omitempty"`

	// APIKey is the API key (for API key auth).
	APIKey string `json:"apiKey,omitempty"`

	// APIKeyHeader is the header name for the API key.
	APIKeyHeader string `json:"apiKeyHeader,omitempty"`
}

// AuthType enumerates authentication types.
type AuthType string

const (
	AuthTypeNone   AuthType = "none"
	AuthTypeBearer AuthType = "bearer"
	AuthTypeBasic  AuthType = "basic"
	AuthTypeAPIKey AuthType = "apiKey"
)
