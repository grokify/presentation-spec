# Slides

Slides are the main content containers in a presentation.

## Structure

```json
{
  "id": "dashboard",
  "sectionId": "metrics",
  "type": "dashboard",
  "title": "Key Metrics",
  "keyMessage": "All metrics trending positive",
  "layout": { ... },
  "widgets": [ ... ],
  "speakerNotes": "Emphasize Q1 growth",
  "quality": { ... }
}
```

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique slide identifier |
| `sectionId` | string | No | Link to narrative section |
| `type` | SlideType | Yes | Slide type |
| `title` | string | No | Slide title |
| `keyMessage` | string | No | Main takeaway |
| `layout` | Layout | No | Layout configuration |
| `widgets` | Widget[] | No | Content widgets |
| `speakerNotes` | string | No | Presenter notes |
| `quality` | Quality | No | Quality metadata |

## Slide Types

### Title

Opening slide with prominent title.

```json
{
  "id": "title",
  "type": "title",
  "title": "Q1 2025 Results",
  "widgets": [
    {
      "id": "subtitle",
      "type": "content_block",
      "content": "Executive Review | January 2025"
    }
  ]
}
```

### Content

Standard content slide.

```json
{
  "id": "overview",
  "type": "content",
  "title": "Overview",
  "keyMessage": "Strong performance across all metrics",
  "widgets": [ ... ]
}
```

### Dashboard

Data-heavy slide with multiple metrics.

```json
{
  "id": "kpis",
  "type": "dashboard",
  "title": "KPI Dashboard",
  "layout": {
    "template": "kpi_row_content_grid"
  },
  "widgets": [
    {"type": "metric_card", ...},
    {"type": "chart", ...}
  ]
}
```

### Section

Section divider slide.

```json
{
  "id": "section-1",
  "type": "section",
  "title": "Financial Results"
}
```

### Summary

Conclusion slide.

```json
{
  "id": "summary",
  "type": "summary",
  "title": "Key Takeaways",
  "keyMessage": "On track for Q2 launch"
}
```

### Blank

Empty slide for custom content.

```json
{
  "id": "custom",
  "type": "blank",
  "widgets": [ ... ]
}
```

## Quality Metadata

Track data freshness and confidence.

```json
{
  "quality": {
    "dataFreshness": "real-time",
    "confidence": 0.95,
    "sources": ["analytics", "crm"],
    "lastUpdated": "2025-01-15T10:30:00Z"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `dataFreshness` | string | Data recency description |
| `confidence` | number | Confidence level (0.0-1.0) |
| `sources` | string[] | Data source names |
| `lastUpdated` | string | Last update timestamp |
