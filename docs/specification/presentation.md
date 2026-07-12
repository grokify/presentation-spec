# Presentation

The root `PresentationSpec` type defines a complete presentation.

## Structure

```json
{
  "version": "1.0",
  "metadata": { ... },
  "theme": { ... },
  "narrative": { ... },
  "slides": [ ... ],
  "dataSources": [ ... ],
  "exports": { ... }
}
```

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `version` | string | Yes | Spec version (e.g., "1.0") |
| `metadata` | Metadata | Yes | Presentation metadata |
| `theme` | Theme | No | Visual styling |
| `narrative` | Narrative | No | Story structure |
| `slides` | SlideSpec[] | Yes | Presentation slides |
| `dataSources` | DataSource[] | No | External data sources |
| `exports` | ExportConfig | No | Export configuration |

## Metadata

```json
{
  "metadata": {
    "title": "Q1 2025 Review",
    "description": "Quarterly business review",
    "author": "Product Team",
    "date": "2025-01-15",
    "tags": ["quarterly", "review", "2025"]
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `title` | string | Presentation title |
| `description` | string | Brief description |
| `author` | string | Author name |
| `date` | string | Creation date (ISO 8601) |
| `tags` | string[] | Categorization tags |

## Example

```json
{
  "version": "1.0",
  "metadata": {
    "title": "Product Launch",
    "author": "Launch Team",
    "date": "2025-01-15"
  },
  "theme": {
    "name": "corporate",
    "aspectRatio": "16:9"
  },
  "slides": [
    {
      "id": "title",
      "type": "title",
      "title": "Product Launch 2025"
    },
    {
      "id": "agenda",
      "type": "content",
      "title": "Agenda",
      "widgets": [
        {
          "id": "agenda-list",
          "type": "content_block",
          "content": "1. Overview\n2. Timeline\n3. Resources\n4. Q&A"
        }
      ]
    }
  ]
}
```
