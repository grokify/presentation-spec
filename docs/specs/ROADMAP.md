# PresentationSpec Roadmap

This document tracks planned features and enhancements for PresentationSpec.

## Completed Features

| Feature | Description | Status |
|---------|-------------|--------|
| Core spec types | Go-first type definitions for presentations | ✓ Done |
| JSON Schema generation | Auto-generate schema from Go types | ✓ Done |
| HTML rendering | Static HTML output via templ templates | ✓ Done |
| PDF export | PDF generation via headless Chrome (chromedp) | ✓ Done |
| Schema validation | Validate specs against JSON Schema | ✓ Done |
| Watch mode | Auto-rebuild on file changes | ✓ Done |
| Dev server | Live reload development server | ✓ Done |
| Data source loading | Load data from JSON, CSV, REST APIs | ✓ Done |
| Syntax highlighting | Code blocks via Prism.js | ✓ Done |
| Mermaid diagrams | Diagram rendering via Mermaid.js | ✓ Done |
| Chart.js integration | Rich charts via Chart.js | ✓ Done |
| Speaker notes view | Presenter view with timer and notes | ✓ Done |
| Unit tests | Test coverage for core packages | ✓ Done |

## Planned Features

### High Priority

#### PPTX Export
- **Description**: Export presentations to PowerPoint format
- **Implementation**: Use `github.com/unidoc/unioffice` library
- **Complexity**: High
- **Notes**: Map widgets to PowerPoint shapes, handle layouts

#### Image Export
- **Description**: Export individual slides as PNG or SVG images
- **Implementation**: Use chromedp to capture screenshots
- **Complexity**: Medium
- **Notes**: Useful for sharing slides on social media, embedding in docs

#### Custom Theme System
- **Description**: Allow user-defined themes via JSON/YAML configuration
- **Implementation**: Theme loading, CSS variable generation
- **Complexity**: Medium
- **Notes**: Support color palettes, typography, spacing presets

#### Reveal.js Export
- **Description**: Export to Reveal.js HTML presentation format
- **Implementation**: Alternative templ templates for Reveal.js structure
- **Complexity**: Medium
- **Notes**: Popular alternative presentation framework

### Developer Experience

#### CI/CD Setup
- **Description**: GitHub Actions workflows for build, test, release
- **Implementation**: `.github/workflows/` configuration
- **Complexity**: Low
- **Tasks**:
  - Build and test on push
  - Lint with golangci-lint
  - Generate and publish schema on release
  - Cross-platform binary releases

#### Documentation Site
- **Description**: MkDocs documentation with examples and API reference
- **Implementation**: `mkdocs.yml` configuration, markdown docs
- **Complexity**: Medium
- **Tasks**:
  - Getting started guide
  - Spec reference documentation
  - Widget gallery with examples
  - Theme customization guide
  - CLI reference

#### Config File Support
- **Description**: Project-level configuration via `.presentationrc` or `presentation.yaml`
- **Implementation**: Viper configuration loading
- **Complexity**: Low
- **Notes**: Default output directory, theme, export options

#### CLI Autocomplete
- **Description**: Shell completions for bash, zsh, fish
- **Implementation**: Cobra completion generation
- **Complexity**: Low
- **Notes**: Improve CLI discoverability

### Polish

#### Accessibility Improvements
- **Description**: Enhanced ARIA labels, focus management, screen reader support
- **Implementation**: Update templ templates with accessibility attributes
- **Complexity**: Medium
- **Tasks**:
  - ARIA landmarks for slide regions
  - Focus management during navigation
  - Skip links
  - High contrast mode support
  - Reduced motion support

#### Print Stylesheet
- **Description**: Optimized CSS for print and PDF output
- **Implementation**: `@media print` CSS rules
- **Complexity**: Low
- **Notes**: Page breaks, hide navigation, optimize colors

#### Slide Transitions
- **Description**: CSS animations between slides
- **Implementation**: CSS transitions, configurable per-slide
- **Complexity**: Medium
- **Notes**: Fade, slide, zoom effects

#### Footer/Header Customization
- **Description**: Configurable branding elements on all slides
- **Implementation**: Add footer/header to spec, render in templates
- **Complexity**: Low
- **Tasks**:
  - Logo placement
  - Page numbers
  - Date/author
  - Custom text

### Advanced Features

#### Presenter Remote Control
- **Description**: WebSocket-based remote control for presentations
- **Implementation**: WebSocket server, mobile-friendly control page
- **Complexity**: High
- **Tasks**:
  - WebSocket server in dev server
  - Remote control web page
  - QR code for connection
  - Sync multiple viewers

#### Multi-File Specs
- **Description**: Split large presentations across multiple JSON files
- **Implementation**: File references, merge logic
- **Complexity**: Medium
- **Notes**: `$ref` support, include directives

#### Template Overrides
- **Description**: Allow custom templ templates per project
- **Implementation**: Template loading from project directory
- **Complexity**: Medium
- **Notes**: Override individual widget or slide templates

#### Marp Export
- **Description**: Export to Marp Markdown format
- **Implementation**: Markdown generation with Marp directives
- **Complexity**: Medium
- **Notes**: Alternative to Reveal.js

#### Plugin System
- **Description**: Extensible architecture for custom widgets and exports
- **Implementation**: Go plugin interface or embedded scripting
- **Complexity**: High
- **Notes**: Long-term architectural consideration

### Widget Enhancements

#### QR Code Widget
- **Description**: Generate QR codes for URLs or text
- **Implementation**: Add `qr_code` widget type, use go-qrcode library
- **Complexity**: Low

#### Video Widget
- **Description**: Embed video content in slides
- **Implementation**: Add `video` widget type with src, autoplay options
- **Complexity**: Low

#### Embed Widget
- **Description**: Embed external content (iframes)
- **Implementation**: Add `embed` widget type with URL and sandbox options
- **Complexity**: Low

#### Timeline Widget
- **Description**: Visual timeline for roadmaps, histories
- **Implementation**: Add `timeline` widget type with events
- **Complexity**: Medium

#### Comparison Widget
- **Description**: Side-by-side comparison layout
- **Implementation**: Add `comparison` widget type with before/after
- **Complexity**: Low

### Data Features

#### Database Data Source
- **Description**: Load data directly from SQL databases
- **Implementation**: Database driver support in data loader
- **Complexity**: High
- **Notes**: Connection string handling, query execution

#### GraphQL Data Source
- **Description**: Load data from GraphQL APIs
- **Implementation**: GraphQL client in data loader
- **Complexity**: Medium

#### Data Transformation
- **Description**: Transform loaded data with jq-style expressions
- **Implementation**: Integrate gojq or similar
- **Complexity**: Medium
- **Notes**: Already have `transform` field in DataSource spec

#### Data Caching
- **Description**: Cache loaded data with TTL
- **Implementation**: File-based or in-memory cache
- **Complexity**: Medium
- **Notes**: Already have `cache` field in DataSource spec

## Contributing

To contribute to any of these features:

1. Open an issue to discuss the approach
2. Reference this roadmap item in the issue
3. Submit a PR with tests and documentation

## Version Planning

### v0.2.0
- Image export (PNG/SVG)
- CI/CD setup
- Print stylesheet improvements

### v0.3.0
- Custom theme system
- Config file support
- CLI autocomplete

### v0.4.0
- PPTX export
- Documentation site

### v1.0.0
- Accessibility audit complete
- All high-priority features
- Comprehensive documentation
- Stable API
