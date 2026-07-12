# PresentationSpec

A Go-first specification library for creating data-driven presentations.

## Overview

PresentationSpec provides a format-agnostic intermediate representation (IR) for presentations, enabling:

- **Go-first design**: Go types are the source of truth, JSON Schema is generated automatically
- **Multiple output formats**: HTML, PDF, and more
- **Rich widget library**: 12 widget types optimized for dashboards and executive presentations
- **Data integration**: Connect to JSON files, CSV, and REST APIs
- **Live development**: Dev server with hot reload for rapid iteration

## Quick Example

```json
{
  "version": "1.0",
  "metadata": {
    "title": "Q1 Results",
    "author": "Product Team"
  },
  "slides": [
    {
      "id": "title",
      "type": "title",
      "title": "Q1 2025 Results"
    },
    {
      "id": "metrics",
      "type": "dashboard",
      "title": "Key Metrics",
      "widgets": [
        {
          "id": "revenue",
          "type": "metric_card",
          "title": "Revenue",
          "value": "1.2M",
          "unit": "$",
          "status": "good",
          "trend": "up",
          "trendValue": "+15%"
        }
      ]
    }
  ]
}
```

## Features

<div class="grid cards" markdown>

-   :material-code-braces:{ .lg .middle } __Go-First Design__

    ---

    Define presentations in Go types. JSON Schema is generated automatically for validation and editor support.

-   :material-file-pdf-box:{ .lg .middle } __Multiple Formats__

    ---

    Export to HTML for web viewing, PDF for printing, with more formats planned.

-   :material-chart-box:{ .lg .middle } __Rich Widgets__

    ---

    12 widget types including metric cards, charts, tables, risk lists, and more.

-   :material-database:{ .lg .middle } __Data Sources__

    ---

    Load data from JSON files, CSV, REST APIs with authentication support.

-   :material-refresh:{ .lg .middle } __Live Reload__

    ---

    Development server with automatic rebuild on file changes.

-   :material-presentation:{ .lg .middle } __Speaker Notes__

    ---

    Presenter view with timer, notes, and slide previews.

</div>

## Installation

```bash
go install github.com/grokify/presentation-spec/cmd/render@latest
go install github.com/grokify/presentation-spec/cmd/genschema@latest
```

## Getting Started

1. Create a presentation spec in JSON
2. Render to HTML: `render presentation.json -o output/`
3. Open `output/index.html` in a browser

See the [Quick Start Guide](getting-started/quickstart.md) for a complete tutorial.

## Widget Types

| Widget | Description |
|--------|-------------|
| `metric_card` | KPI metrics with trends and sparklines |
| `chart` | Bar, line, pie, area charts via Chart.js |
| `table` | Data tables with configurable columns |
| `risk_list` | Risk items with severity and mitigation |
| `decision_list` | Decisions requiring action |
| `checklist` | Task lists with completion status |
| `content_block` | Markdown/HTML content |
| `diagram` | Mermaid, PlantUML, Graphviz diagrams |
| `callout` | Info, warning, error, success callouts |
| `quote` | Quotations with attribution |
| `image` | Images with captions |
| `code` | Code blocks with syntax highlighting |

## Links

- [GitHub Repository](https://github.com/grokify/presentation-spec)
- [Roadmap](specs/ROADMAP.md)
- [Release Notes](releases/v0.1.0.md)
