# CLI Reference

## render

The main command for rendering PresentationSpec files.

### Usage

```bash
render [spec.json] [flags]
render [command]
```

### Commands

| Command | Description |
|---------|-------------|
| `serve` | Start a dev server with live reload |
| `help` | Help about any command |
| `completion` | Generate shell completion scripts |

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--output` | `-o` | `<name>-output/` | Output directory (HTML) or file (PDF) |
| `--file` | `-f` | | Output single HTML file (no assets) |
| `--format` | | `html` | Output format: `html`, `pdf` |
| `--validate` | | `false` | Validate spec before rendering |
| `--validate-only` | | `false` | Validate only, don't render |
| `--watch` | `-w` | `false` | Watch for changes and auto-rebuild |
| `--speaker-notes` | | `false` | Generate speaker notes view |
| `--page-size` | | `Letter` | PDF page size: Letter, A4, Legal, Tabloid |
| `--landscape` | | `true` | PDF landscape orientation |
| `--help` | `-h` | | Show help |

### Examples

#### Render to HTML

```bash
# Default output directory (presentation-output/)
render presentation.json

# Custom output directory
render presentation.json -o dist/

# Single HTML file (no assets)
render presentation.json -f presentation.html
```

#### Render to PDF

```bash
# Default output (presentation.pdf)
render presentation.json --format pdf

# Custom output path
render presentation.json --format pdf -o report.pdf

# A4 paper size
render presentation.json --format pdf --page-size A4

# Portrait orientation
render presentation.json --format pdf --landscape=false
```

#### Validation

```bash
# Validate before rendering
render presentation.json --validate -o output/

# Validate only (no output)
render presentation.json --validate-only
```

#### Watch Mode

```bash
# Watch for changes
render presentation.json -o output/ --watch
```

#### Speaker Notes

```bash
# Generate with speaker notes view
render presentation.json -o output/ --speaker-notes
```

---

## render serve

Start a development server with live reload.

### Usage

```bash
render serve [spec.json] [flags]
```

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--port` | `-p` | `8080` | Port to listen on |
| `--host` | | `localhost` | Host to bind to |
| `--output` | `-o` | `<name>-output/` | Output directory |
| `--no-watch` | | `false` | Disable auto-rebuild |
| `--no-browser` | | `false` | Don't open browser on start |

### Examples

```bash
# Default (localhost:8080)
render serve presentation.json

# Custom port
render serve presentation.json --port 3000

# Bind to all interfaces
render serve presentation.json --host 0.0.0.0

# Disable auto-rebuild
render serve presentation.json --no-watch
```

---

## genschema

Generate JSON Schema from Go types.

### Usage

```bash
genschema [flags]
```

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--output` | `-o` | `schema/presentation.schema.json` | Output path |

### Examples

```bash
# Default output
genschema

# Custom output path
genschema -o docs/schema.json
```

---

## Shell Completion

Generate shell completion scripts for enhanced CLI experience.

### Bash

```bash
render completion bash > /etc/bash_completion.d/render
```

### Zsh

```bash
render completion zsh > "${fpath[1]}/_render"
```

### Fish

```bash
render completion fish > ~/.config/fish/completions/render.fish
```

### PowerShell

```powershell
render completion powershell | Out-String | Invoke-Expression
```
