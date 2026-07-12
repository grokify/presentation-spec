# Layouts

Layouts define how widgets are arranged on a slide.

## Structure

```json
{
  "layout": {
    "template": "kpi_row_content_grid",
    "regions": [
      {"id": "kpi-row", "role": "kpi"},
      {"id": "main-content", "role": "content", "columns": 3},
      {"id": "sidebar", "role": "sidebar", "columns": 1}
    ]
  }
}
```

## Fields

| Field | Type | Description |
|-------|------|-------------|
| `template` | string | Predefined layout template |
| `regions` | Region[] | Custom regions |

## Templates

### single

Single full-width region.

```
┌─────────────────────┐
│                     │
│       Content       │
│                     │
└─────────────────────┘
```

### two_column

Two equal columns.

```
┌──────────┬──────────┐
│          │          │
│   Left   │   Right  │
│          │          │
└──────────┴──────────┘
```

### three_column

Three equal columns.

```
┌──────┬──────┬──────┐
│      │      │      │
│  1   │  2   │  3   │
│      │      │      │
└──────┴──────┴──────┘
```

### kpi_row_content_grid

KPI metrics row with content grid below.

```
┌─────┬─────┬─────┬─────┐
│ KPI │ KPI │ KPI │ KPI │
├─────┴─────┴─────┼─────┤
│                 │     │
│     Content     │Side │
│                 │     │
└─────────────────┴─────┘
```

### metric_dashboard

Auto-fit grid for metrics.

```
┌─────┬─────┬─────┬─────┐
│     │     │     │     │
├─────┼─────┼─────┼─────┤
│     │     │     │     │
└─────┴─────┴─────┴─────┘
```

## Regions

### Fields

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Region identifier |
| `role` | string | Semantic role |
| `columns` | number | Column span |
| `rows` | number | Row span |

### Roles

| Role | Description |
|------|-------------|
| `kpi` | KPI metrics region |
| `content` | Main content area |
| `sidebar` | Side panel |
| `header` | Header region |
| `footer` | Footer region |

## Widget Placement

Widgets reference regions via `regionId`:

```json
{
  "layout": {
    "template": "two_column",
    "regions": [
      {"id": "left", "role": "content"},
      {"id": "right", "role": "content"}
    ]
  },
  "widgets": [
    {"id": "chart1", "regionId": "left", "type": "chart", ...},
    {"id": "table1", "regionId": "right", "type": "table", ...}
  ]
}
```
