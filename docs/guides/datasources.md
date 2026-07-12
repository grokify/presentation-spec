# Working with Data Sources

This guide covers connecting presentations to external data.

## Overview

Data sources allow widgets to pull data from files or APIs instead of hardcoding values. This enables:

- **Dynamic dashboards** - Live metrics from APIs
- **Reusable templates** - Same presentation, different data
- **Data separation** - Keep data in JSON/CSV files

## File-Based Data

### JSON Files

Create a data file:

```json
// data/metrics.json
{
  "revenue": {
    "value": "1.2M",
    "trend": "+15%",
    "status": "good"
  },
  "users": {
    "value": "45K",
    "trend": "+8%",
    "status": "good"
  }
}
```

Reference it in your presentation:

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

### CSV Files

For tabular data, use CSV:

```csv
month,revenue,users
Jan,100000,40000
Feb,120000,42000
Mar,145000,45000
```

```json
{
  "dataSources": [
    {
      "id": "monthly-data",
      "type": "csv",
      "path": "./data/monthly.csv"
    }
  ]
}
```

## API Data Sources

### Basic API

```json
{
  "dataSources": [
    {
      "id": "api-metrics",
      "type": "api",
      "endpoint": "https://api.example.com/metrics"
    }
  ]
}
```

### With Authentication

**Bearer Token:**

```json
{
  "id": "secure-api",
  "type": "api",
  "endpoint": "https://api.example.com/data",
  "auth": {
    "type": "bearer",
    "token": "${API_TOKEN}"
  }
}
```

**Basic Auth:**

```json
{
  "auth": {
    "type": "basic",
    "username": "${API_USER}",
    "password": "${API_PASS}"
  }
}
```

**API Key:**

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

Use `${VAR_NAME}` to reference environment variables:

```json
{
  "endpoint": "${API_BASE_URL}/metrics",
  "auth": {
    "type": "bearer",
    "token": "${API_TOKEN}"
  }
}
```

Set variables before rendering:

```bash
export API_BASE_URL="https://api.example.com"
export API_TOKEN="your-token"
presspec render presentation.json --output ./output
```

## Binding Widgets to Data

### dataSourceId

Reference the data source:

```json
{
  "id": "revenue-card",
  "type": "metric_card",
  "dataSourceId": "metrics"
}
```

### dataPath

Extract specific data using JSON path:

```json
{
  "id": "revenue-card",
  "type": "metric_card",
  "dataSourceId": "metrics",
  "dataPath": "revenue"
}
```

For nested data:

```json
{
  "dataPath": "quarterly.q1.revenue"
}
```

## Chart Data Binding

Bind chart series to data:

```json
{
  "type": "chart",
  "chartType": "line",
  "dataSourceId": "monthly-data",
  "data": {
    "labelsPath": "months",
    "seriesPath": "values"
  }
}
```

## Live Refresh

For real-time dashboards, set refresh intervals:

```json
{
  "id": "live-kpis",
  "type": "api",
  "endpoint": "https://api.example.com/live",
  "refresh": "30s"
}
```

Supported intervals:

| Format | Example | Description |
|--------|---------|-------------|
| `Ns` | `30s` | N seconds |
| `Nm` | `5m` | N minutes |
| `Nh` | `1h` | N hours |

## Error Handling

When data loading fails:

1. **File not found** - Check path is relative to presentation file
2. **API error** - Verify endpoint and credentials
3. **Parse error** - Validate JSON/CSV format

The CLI shows detailed errors:

```bash
presspec render presentation.json -v
# Error: failed to load data source "metrics": file not found: ./data/metrics.json
```

## Best Practices

1. **Use environment variables for secrets** - Never hardcode API tokens
2. **Keep data files alongside presentation** - Use relative paths
3. **Validate data format** - Ensure data matches widget expectations
4. **Set appropriate refresh intervals** - Balance freshness vs. API load

## Complete Example

```json
{
  "version": "1.0",
  "metadata": {
    "title": "Live Dashboard"
  },
  "dataSources": [
    {
      "id": "local-config",
      "type": "json",
      "path": "./data/config.json"
    },
    {
      "id": "live-metrics",
      "type": "api",
      "endpoint": "${API_URL}/metrics",
      "auth": {
        "type": "bearer",
        "token": "${API_TOKEN}"
      },
      "refresh": "1m"
    }
  ],
  "slides": [
    {
      "id": "dashboard",
      "type": "dashboard",
      "widgets": [
        {
          "id": "revenue",
          "type": "metric_card",
          "title": "Revenue",
          "dataSourceId": "live-metrics",
          "dataPath": "kpis.revenue"
        }
      ]
    }
  ]
}
```
