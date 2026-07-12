# Installation

## Prerequisites

- Go 1.21 or later
- Chrome or Chromium (for PDF export)

## Install CLI Tools

Install the render command for generating presentations:

```bash
go install github.com/grokify/presentation-spec/cmd/render@latest
```

Install the schema generator (optional, for development):

```bash
go install github.com/grokify/presentation-spec/cmd/genschema@latest
```

## Verify Installation

```bash
render --help
```

You should see:

```
render converts a PresentationSpec JSON file into a static HTML presentation or PDF.

Output formats:
  - html (default): Static HTML with CSS/JS assets
  - pdf: PDF document via headless Chrome

Usage:
  render [spec.json] [flags]
  render [command]

Available Commands:
  serve       Start a dev server with live reload
  ...
```

## Use as a Library

Add PresentationSpec to your Go project:

```bash
go get github.com/grokify/presentation-spec
```

Import the packages you need:

```go
import (
    "github.com/grokify/presentation-spec/spec"
    "github.com/grokify/presentation-spec/templ"
    "github.com/grokify/presentation-spec/schema"
)
```

## Chrome for PDF Export

PDF export requires Chrome or Chromium. Install via:

=== "macOS"

    ```bash
    brew install --cask google-chrome
    ```

=== "Ubuntu/Debian"

    ```bash
    wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
    sudo dpkg -i google-chrome-stable_current_amd64.deb
    ```

=== "Windows"

    Download from [google.com/chrome](https://www.google.com/chrome/)

## Next Steps

- [Quick Start Guide](quickstart.md) - Create your first presentation
- [CLI Reference](cli.md) - Learn all CLI commands and options
