# Building Dashboards

This guide covers creating data-rich dashboard presentations.

## Overview

Dashboards combine multiple widgets to display KPIs, charts, and status information. PresentationSpec provides specialized layouts and widgets for dashboard-style slides.

## Dashboard Layout

Use the `kpi_row_content_grid` template for typical dashboards:

```json
{
  "id": "metrics-dashboard",
  "type": "dashboard",
  "title": "Q1 Metrics",
  "layout": {
    "template": "kpi_row_content_grid",
    "regions": [
      {"id": "kpi-row", "role": "kpi"},
      {"id": "main", "role": "content", "columns": 3},
      {"id": "sidebar", "role": "sidebar", "columns": 1}
    ]
  }
}
```

This creates:

```
┌─────┬─────┬─────┬─────┐
│ KPI │ KPI │ KPI │ KPI │
├─────┴─────┴─────┼─────┤
│                 │     │
│     Content     │Side │
│                 │     │
└─────────────────┴─────┘
```

## KPI Metrics Row

Add metric cards to the KPI region:

```json
{
  "widgets": [
    {
      "id": "revenue",
      "type": "metric_card",
      "regionId": "kpi-row",
      "title": "Revenue",
      "value": "1.2M",
      "unit": "$",
      "status": "good",
      "trend": "up",
      "trendValue": "+15%",
      "sparkline": [100, 110, 105, 120, 140, 160, 180]
    },
    {
      "id": "users",
      "type": "metric_card",
      "regionId": "kpi-row",
      "title": "Active Users",
      "value": "45K",
      "status": "good",
      "trend": "up",
      "trendValue": "+8%"
    },
    {
      "id": "churn",
      "type": "metric_card",
      "regionId": "kpi-row",
      "title": "Churn Rate",
      "value": "2.1%",
      "status": "warning",
      "trend": "up",
      "trendValue": "+0.3%",
      "target": "< 2%"
    }
  ]
}
```

## Status Indicators

Metric cards support status colors:

| Status | Use Case |
|--------|----------|
| `good` | On target, positive trend |
| `warning` | Needs attention |
| `critical` | Requires action |
| `neutral` | Informational |

## Charts in Dashboards

Add charts to the main content area:

```json
{
  "id": "revenue-trend",
  "type": "chart",
  "regionId": "main",
  "title": "Revenue Trend",
  "chartType": "line",
  "data": {
    "labels": ["Jan", "Feb", "Mar", "Apr", "May", "Jun"],
    "series": [
      {"name": "2024", "values": [100, 120, 115, 140, 155, 170]},
      {"name": "2025", "values": [110, 135, 145, 165, 180, 200]}
    ]
  },
  "options": {
    "showLegend": true,
    "showGrid": true,
    "yAxisLabel": "Revenue ($K)"
  }
}
```

## Metric Dashboard Layout

For all-metrics slides, use `metric_dashboard`:

```json
{
  "layout": {
    "template": "metric_dashboard"
  },
  "widgets": [
    {"id": "m1", "type": "metric_card", "title": "Revenue", ...},
    {"id": "m2", "type": "metric_card", "title": "Users", ...},
    {"id": "m3", "type": "metric_card", "title": "NPS", ...},
    {"id": "m4", "type": "metric_card", "title": "Churn", ...},
    {"id": "m5", "type": "metric_card", "title": "CAC", ...},
    {"id": "m6", "type": "metric_card", "title": "LTV", ...}
  ]
}
```

This auto-fits metrics into a grid.

## Live Data

Connect dashboards to live data sources:

```json
{
  "dataSources": [
    {
      "id": "live-metrics",
      "type": "api",
      "endpoint": "https://api.example.com/metrics",
      "auth": {"type": "bearer", "token": "${API_TOKEN}"},
      "refresh": "30s"
    }
  ],
  "slides": [
    {
      "widgets": [
        {
          "id": "revenue",
          "type": "metric_card",
          "dataSourceId": "live-metrics",
          "dataPath": "revenue"
        }
      ]
    }
  ]
}
```

## Complete Example

```json
{
  "version": "1.0",
  "metadata": {
    "title": "Q1 Dashboard",
    "author": "Analytics Team"
  },
  "theme": {
    "name": "corporate",
    "aspectRatio": "16:9"
  },
  "slides": [
    {
      "id": "dashboard",
      "type": "dashboard",
      "title": "Q1 2025 Performance",
      "keyMessage": "All metrics trending positive",
      "layout": {
        "template": "kpi_row_content_grid"
      },
      "widgets": [
        {
          "id": "revenue",
          "type": "metric_card",
          "regionId": "kpi-row",
          "title": "Revenue",
          "value": "1.2M",
          "unit": "$",
          "status": "good",
          "trend": "up",
          "trendValue": "+15%"
        },
        {
          "id": "users",
          "type": "metric_card",
          "regionId": "kpi-row",
          "title": "Users",
          "value": "45K",
          "status": "good",
          "trend": "up",
          "trendValue": "+8%"
        },
        {
          "id": "chart",
          "type": "chart",
          "regionId": "main",
          "title": "Monthly Trend",
          "chartType": "line",
          "data": {
            "labels": ["Jan", "Feb", "Mar"],
            "series": [
              {"name": "Revenue", "values": [100, 120, 145]}
            ]
          }
        }
      ]
    }
  ]
}
```
