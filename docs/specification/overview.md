# Specification Overview

PresentationSpec is a format-agnostic intermediate representation (IR) for presentations. It defines the structure and content of presentations in a way that can be rendered to multiple output formats.

## Design Principles

### Go-First

Go types are the **source of truth**. The JSON Schema is generated from Go structs, not hand-written. This ensures:

- Type safety in Go code
- Automatic schema updates when types change
- Consistent validation across languages

### Format-Agnostic

The spec describes *what* to present, not *how* to render it. This enables:

- Multiple output formats (HTML, PDF, PPTX)
- Different themes and layouts
- Future format support without spec changes

### Data-Driven

Presentations can reference external data sources, enabling:

- Real-time data in presentations
- Automated dashboard generation
- Data refresh without spec changes

## Structure

A PresentationSpec consists of:

```
PresentationSpec
├── version          # Spec version
├── metadata         # Title, author, date, tags
├── theme            # Visual styling
├── narrative        # Story structure
├── slides[]         # Content slides
│   ├── layout       # Slide layout
│   └── widgets[]    # Content elements
├── dataSources[]    # External data
└── exports          # Export configuration
```

## Core Types

| Type | Description |
|------|-------------|
| `PresentationSpec` | Root container for a presentation |
| `Metadata` | Title, description, author, date |
| `Theme` | Colors, typography, density |
| `Narrative` | Story structure with sections |
| `SlideSpec` | Individual slide content |
| `Widget` | Content element (metric, chart, etc.) |
| `Layout` | Slide layout template |
| `DataSource` | External data reference |
| `ExportConfig` | Export format settings |

## Slide Types

| Type | Use Case |
|------|----------|
| `title` | Opening slide with title and subtitle |
| `content` | Standard content slide |
| `dashboard` | Data-heavy KPI dashboard |
| `section` | Section divider |
| `summary` | Conclusion or summary |
| `blank` | Empty slide for custom content |

## Widget Types

| Type | Use Case |
|------|----------|
| `metric_card` | Single KPI with trend |
| `chart` | Data visualization |
| `table` | Tabular data |
| `risk_list` | Risk items with severity |
| `decision_list` | Decisions needing action |
| `checklist` | Task completion list |
| `content_block` | Markdown/HTML text |
| `diagram` | Mermaid/PlantUML diagrams |
| `callout` | Highlighted information |
| `quote` | Quotations |
| `image` | Images with captions |
| `code` | Code with syntax highlighting |

## JSON Schema

The JSON Schema is available at:

```
schema/presentation.schema.json
```

Use it for:

- Editor autocompletion (VS Code, JetBrains)
- Validation in CI/CD pipelines
- Documentation generation

### VS Code Setup

Add to your `.vscode/settings.json`:

```json
{
  "json.schemas": [
    {
      "fileMatch": ["*.presentation.json"],
      "url": "./schema/presentation.schema.json"
    }
  ]
}
```

## Versioning

The spec uses semantic versioning:

- **Major**: Breaking changes to spec structure
- **Minor**: New fields or widget types (backwards compatible)
- **Patch**: Documentation or schema fixes

Current version: `1.0`
