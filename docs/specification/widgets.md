# Widgets

Widgets are the content elements within slides. Each widget has a `type` field that determines its structure and rendering.

## Common Fields

All widgets share these fields:

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique identifier |
| `type` | string | Widget type discriminator |
| `title` | string | Widget title |
| `regionId` | string | Target layout region |
| `dataSourceId` | string | Reference to data source |

## Metric Card

Display a single KPI metric with optional trend and sparkline.

```json
{
  "id": "revenue",
  "type": "metric_card",
  "title": "Revenue",
  "value": "1.2M",
  "unit": "$",
  "status": "good",
  "trend": "up",
  "trendValue": "+15%",
  "target": "1M",
  "sparkline": [100, 120, 115, 140, 160, 180]
}
```

| Field | Type | Description |
|-------|------|-------------|
| `value` | string | Primary metric value |
| `unit` | string | Unit of measurement |
| `status` | string | `good`, `warning`, `critical`, `neutral` |
| `trend` | string | `up`, `down`, `flat` |
| `trendValue` | string | Trend amount (e.g., "+15%") |
| `target` | string | Target value |
| `sparkline` | number[] | Data points for sparkline |

## Chart

Data visualization using Chart.js.

```json
{
  "id": "sales-chart",
  "type": "chart",
  "title": "Sales Trend",
  "chartType": "line",
  "data": {
    "labels": ["Q1", "Q2", "Q3", "Q4"],
    "series": [
      {"name": "2024", "values": [100, 150, 200, 250]},
      {"name": "2025", "values": [120, 180, 220, 300]}
    ]
  },
  "options": {
    "showLegend": true,
    "showGrid": true,
    "yAxisLabel": "Revenue ($K)"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `chartType` | string | `bar`, `line`, `pie`, `donut`, `area`, `scatter`, `radar` |
| `data.labels` | string[] | X-axis labels |
| `data.series` | array | Data series with name, values, color |
| `options.showLegend` | boolean | Show legend |
| `options.showGrid` | boolean | Show grid lines |
| `options.stacked` | boolean | Stack series |
| `options.xAxisLabel` | string | X-axis label |
| `options.yAxisLabel` | string | Y-axis label |

## Table

Tabular data display.

```json
{
  "id": "features-table",
  "type": "table",
  "title": "Feature Status",
  "columns": [
    {"key": "feature", "header": "Feature", "width": "50%"},
    {"key": "owner", "header": "Owner", "width": "25%"},
    {"key": "status", "header": "Status", "width": "25%", "align": "center"}
  ],
  "rows": [
    {"feature": "User Auth", "owner": "Alice", "status": "Complete"},
    {"feature": "Dashboard", "owner": "Bob", "status": "In Progress"}
  ]
}
```

| Field | Type | Description |
|-------|------|-------------|
| `columns` | array | Column definitions |
| `columns[].key` | string | Row data key |
| `columns[].header` | string | Column header text |
| `columns[].width` | string | Column width |
| `columns[].align` | string | `left`, `center`, `right` |
| `rows` | array | Row data objects |

## Risk List

Display risks with severity and mitigation.

```json
{
  "id": "risks",
  "type": "risk_list",
  "title": "Open Risks",
  "items": [
    {
      "text": "API rate limits may impact launch",
      "severity": "high",
      "mitigation": "Implement caching layer",
      "owner": "Platform Team"
    },
    {
      "text": "Documentation incomplete",
      "severity": "medium",
      "mitigation": "Doc sprint this week",
      "owner": "Tech Writing"
    }
  ]
}
```

| Field | Type | Description |
|-------|------|-------------|
| `items` | array | Risk items |
| `items[].text` | string | Risk description |
| `items[].severity` | string | `low`, `medium`, `high`, `critical` |
| `items[].mitigation` | string | Mitigation plan |
| `items[].owner` | string | Responsible party |
| `items[].status` | string | Item status |

## Decision List

Display decisions requiring action.

```json
{
  "id": "decisions",
  "type": "decision_list",
  "title": "Pending Decisions",
  "items": [
    {
      "text": "Choose deployment strategy",
      "description": "Phased rollout vs big-bang launch",
      "owner": "VP Product",
      "deadline": "2025-01-20"
    }
  ]
}
```

| Field | Type | Description |
|-------|------|-------------|
| `items` | array | Decision items |
| `items[].text` | string | Decision title |
| `items[].description` | string | Additional context |
| `items[].owner` | string | Decision maker |
| `items[].deadline` | string | Due date |

## Checklist

Task completion list.

```json
{
  "id": "launch-checklist",
  "type": "checklist",
  "title": "Launch Checklist",
  "items": [
    {"text": "Feature freeze", "checked": true},
    {"text": "Security review", "checked": true},
    {"text": "Load testing", "checked": false},
    {"text": "Documentation", "checked": false}
  ]
}
```

| Field | Type | Description |
|-------|------|-------------|
| `items` | array | Checklist items |
| `items[].text` | string | Item text |
| `items[].checked` | boolean | Completion status |

## Content Block

Formatted text content.

```json
{
  "id": "summary",
  "type": "content_block",
  "title": "Summary",
  "content": "## Key Points\n\n- Revenue up 15%\n- User growth on track\n- Launch scheduled for Q2",
  "format": "markdown"
}
```

| Field | Type | Description |
|-------|------|-------------|
| `content` | string | Text content |
| `format` | string | `markdown`, `html`, `plain` |

## Diagram

Diagrams via Mermaid or other tools.

```json
{
  "id": "architecture",
  "type": "diagram",
  "title": "System Architecture",
  "diagramType": "mermaid",
  "source": "graph LR\n  A[Client] --> B[API]\n  B --> C[Database]"
}
```

| Field | Type | Description |
|-------|------|-------------|
| `diagramType` | string | `mermaid`, `plantuml`, `graphviz`, `d2` |
| `source` | string | Diagram source code |

## Callout

Highlighted information box.

```json
{
  "id": "warning",
  "type": "callout",
  "title": "Important",
  "calloutType": "warning",
  "content": "API changes require client updates before launch."
}
```

| Field | Type | Description |
|-------|------|-------------|
| `calloutType` | string | `info`, `warning`, `error`, `success`, `tip` |
| `content` | string | Callout text |

## Quote

Quotation with attribution.

```json
{
  "id": "quote",
  "type": "quote",
  "quote": "The best way to predict the future is to create it.",
  "attribution": "Peter Drucker"
}
```

| Field | Type | Description |
|-------|------|-------------|
| `quote` | string | Quoted text |
| `attribution` | string | Quote source |

## Image

Image with optional caption.

```json
{
  "id": "hero",
  "type": "image",
  "src": "/images/hero.png",
  "alt": "Product screenshot",
  "caption": "New dashboard design"
}
```

| Field | Type | Description |
|-------|------|-------------|
| `src` | string | Image URL or path |
| `alt` | string | Alt text for accessibility |
| `caption` | string | Image caption |

## Code

Code block with syntax highlighting.

```json
{
  "id": "example",
  "type": "code",
  "title": "API Example",
  "language": "go",
  "code": "func main() {\n  fmt.Println(\"Hello\")\n}",
  "highlight": [2]
}
```

| Field | Type | Description |
|-------|------|-------------|
| `language` | string | Programming language |
| `code` | string | Code content |
| `highlight` | number[] | Lines to highlight |
