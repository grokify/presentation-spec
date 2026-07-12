# Data Sources

Data sources connect presentations to external data for dynamic content.

## Structure

```json
{
  "dataSources": [
    {
      "id": "metrics",
      "type": "json",
      "path": "./data/metrics.json"
    }
  ]
}
```

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier |
| `type` | string | Yes | Source type |
| `path` | string | No | File path (for file sources) |
| `endpoint` | string | No | URL (for API sources) |
| `auth` | Auth | No | Authentication config |
| `refresh` | string | No | Refresh interval |

## Source Types

### JSON File

Load data from a local JSON file.

```json
{
  "id": "metrics",
  "type": "json",
  "path": "./data/metrics.json"
}
```

### CSV File

Load tabular data from CSV.

```json
{
  "id": "sales",
  "type": "csv",
  "path": "./data/sales.csv"
}
```

### REST API

Fetch data from an HTTP endpoint.

```json
{
  "id": "live-metrics",
  "type": "api",
  "endpoint": "https://api.example.com/metrics",
  "auth": {
    "type": "bearer",
    "token": "${API_TOKEN}"
  }
}
```

## Authentication

### Bearer Token

```json
{
  "auth": {
    "type": "bearer",
    "token": "${API_TOKEN}"
  }
}
```

### Basic Auth

```json
{
  "auth": {
    "type": "basic",
    "username": "${API_USER}",
    "password": "${API_PASS}"
  }
}
```

### API Key

```json
{
  "auth": {
    "type": "apikey",
    "header": "X-API-Key",
    "key": "${API_KEY}"
  }
}
```

## Environment Variables

Use `${VAR_NAME}` syntax to reference environment variables:

```json
{
  "endpoint": "${API_BASE_URL}/metrics",
  "auth": {
    "type": "bearer",
    "token": "${API_TOKEN}"
  }
}
```

## Widget Binding

Widgets reference data sources via `dataSourceId`:

```json
{
  "widgets": [
    {
      "id": "revenue-chart",
      "type": "chart",
      "dataSourceId": "metrics",
      "dataPath": "revenue.monthly"
    }
  ]
}
```

| Field | Description |
|-------|-------------|
| `dataSourceId` | Reference to data source |
| `dataPath` | JSON path to extract data |

## Refresh Intervals

For live data, specify refresh intervals:

```json
{
  "id": "realtime",
  "type": "api",
  "endpoint": "https://api.example.com/live",
  "refresh": "30s"
}
```

| Value | Description |
|-------|-------------|
| `30s` | Every 30 seconds |
| `5m` | Every 5 minutes |
| `1h` | Every hour |

## Example

Complete data source configuration:

```json
{
  "dataSources": [
    {
      "id": "local-metrics",
      "type": "json",
      "path": "./data/q1-metrics.json"
    },
    {
      "id": "sales-data",
      "type": "csv",
      "path": "./data/sales-2025.csv"
    },
    {
      "id": "live-kpis",
      "type": "api",
      "endpoint": "${API_URL}/kpis",
      "auth": {
        "type": "bearer",
        "token": "${API_TOKEN}"
      },
      "refresh": "5m"
    }
  ]
}
```
