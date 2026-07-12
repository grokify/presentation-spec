# PresentationSpec

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Docs][docs-mkdoc-svg]][docs-mkdoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/presentation-spec/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/presentation-spec/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/presentation-spec/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/presentation-spec/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/presentation-spec/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/presentation-spec/actions/workflows/go-sast-codeql.yaml
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/presentation-spec
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/presentation-spec
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://grokify.github.io/presentation-spec
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fpresentation-spec
 [loc-svg]: https://tokei.rs/b1/github/grokify/presentation-spec
 [repo-url]: https://github.com/grokify/presentation-spec
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/presentation-spec/blob/main/LICENSE

A Go-first specification library for presentations. Define your presentation once in a format-agnostic intermediate representation (IR), then render to multiple output formats deterministically.

## Features

- 🎯 **Go-first**: Go types are the source of truth
- 📋 **JSON Schema**: Auto-generated schema from Go types for validation and editor support
- 🔄 **Format-agnostic**: Render to web, PDF, PPTX, and more
- 📊 **Data-driven**: Connect widgets to external data sources
- 🛡️ **Type-safe**: Strong typing prevents common errors

## Installation

```bash
go get github.com/grokify/presentation-spec
```

## Quick Start

### Define a Presentation

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/grokify/presentation-spec/spec"
)

func main() {
    pres := spec.PresentationSpec{
        Version: "1.0.0",
        Metadata: spec.Metadata{
            Title:  "Hello World",
            Author: "Your Name",
        },
        Slides: []spec.SlideSpec{
            {
                ID:    "title-slide",
                Type:  spec.SlideTypeTitle,
                Title: "Hello World",
            },
        },
    }

    data, _ := json.MarshalIndent(pres, "", "  ")
    fmt.Println(string(data))
}
```

### Validate Against Schema

The embedded JSON Schema can be used for validation:

```go
import "github.com/grokify/presentation-spec/schema"

schemaJSON := schema.PresentationSchemaJSON()
// Use with your preferred JSON Schema validator
```

## Rendering to HTML

Render presentations to static HTML using the CLI:

```bash
# Render to a directory with assets
go run ./cmd/render presentation.json -o output/

# Output structure
output/
├── index.html           # Full presentation
└── assets/
    ├── style.css        # Stylesheet
    └── navigation.js    # Keyboard navigation
```

Or use the renderer programmatically:

```go
import (
    "context"
    "github.com/grokify/presentation-spec/templ"
)

pres, _ := templ.LoadSpec("presentation.json")
renderer := templ.NewRenderer()
renderer.RenderToDir(context.Background(), pres, "output/")
```

The rendered HTML includes:

- Keyboard navigation (arrow keys, space, page up/down)
- URL hash navigation (#slide-id)
- Responsive CSS with theme support
- SVG charts for basic visualizations

## Project Structure

```
presentation-spec/
├── spec/                    # Core Go types (source of truth)
│   ├── presentation.go      # PresentationSpec type
│   ├── slide.go             # SlideSpec type
│   ├── widget.go            # Widget types
│   ├── layout.go            # Layout types
│   ├── theme.go             # Theme types
│   ├── narrative.go         # Narrative types
│   ├── datasource.go        # DataSource types
│   └── export.go            # Export configuration
├── templ/                   # HTML rendering (a-h/templ)
│   ├── presentation.templ   # Main template
│   ├── slide.templ          # Slide template
│   ├── render.go            # Renderer API
│   └── widgets/             # Widget templates
├── schema/
│   ├── generate.go          # Schema generator
│   ├── embed.go             # Embedded schema
│   └── presentation.schema.json
├── cmd/
│   ├── genschema/           # Schema generation CLI
│   └── render/              # HTML rendering CLI
└── examples/                # Example presentations
```

## Core Concepts

### PresentationSpec

The root type containing all presentation content and configuration:

- **Metadata**: Title, author, date, description
- **Theme**: Visual styling and branding
- **Narrative**: Storyline structure and sections
- **Slides**: Ordered list of slides
- **DataSources**: External data connections
- **Exports**: Output format configuration

### Widgets

Content elements that populate slides. Widget types include:

| Type | Description |
|------|-------------|
| `metric_card` | Single metric with trend indicator |
| `content_block` | Formatted text (markdown/html) |
| `risk_list` | List of risks with severity |
| `decision_list` | Decisions requiring action |
| `checklist` | Checkbox list with status |
| `chart` | Data visualizations |
| `table` | Tabular data |
| `diagram` | Mermaid/PlantUML diagrams |
| `callout` | Highlighted information |
| `quote` | Quotations with attribution |
| `image` | Images with captions |
| `code` | Syntax-highlighted code |

### Layouts

Predefined layout templates:

- `single` - Single content region
- `two_column` - Two equal columns
- `kpi_row_content_grid` - Row of KPIs with content below
- `metric_dashboard` - Dashboard-style metric grid
- And more...

## Regenerating Generated Files

When Go types change, regenerate the JSON Schema:

```bash
go run ./cmd/genschema
```

When templ files change, regenerate the Go code:

```bash
templ generate
```

Install templ CLI with:

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

## Examples

See the `examples/` directory:

- `basic/hello-world.json` - Minimal presentation
- `dashboard/launch-readiness.json` - Full dashboard example
- `executive/board-update.json` - Executive summary example

## License

MIT
