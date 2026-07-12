package spec

// ExportConfig configures output format targets.
type ExportConfig struct {
	// Targets specifies the export targets.
	Targets []ExportTarget `json:"targets,omitempty"`

	// OutputDir is the base output directory.
	OutputDir string `json:"outputDir,omitempty"`

	// Filename is the base filename (without extension).
	Filename string `json:"filename,omitempty"`
}

// ExportTarget represents a single export target configuration.
type ExportTarget struct {
	// Format is the output format.
	Format ExportFormat `json:"format"`

	// Enabled controls whether this target is active.
	Enabled bool `json:"enabled,omitempty"`

	// OutputPath is the output path for this target.
	OutputPath string `json:"outputPath,omitempty"`

	// Options contains format-specific options.
	Options *ExportOptions `json:"options,omitempty"`
}

// ExportFormat enumerates supported export formats.
type ExportFormat string

const (
	// ExportFormatWeb is HTML/web output.
	ExportFormatWeb ExportFormat = "web"

	// ExportFormatPDF is PDF output.
	ExportFormatPDF ExportFormat = "pdf"

	// ExportFormatPPTX is PowerPoint output.
	ExportFormatPPTX ExportFormat = "pptx"

	// ExportFormatPNG is PNG image output.
	ExportFormatPNG ExportFormat = "png"

	// ExportFormatSVG is SVG image output.
	ExportFormatSVG ExportFormat = "svg"

	// ExportFormatMarkdown is Markdown output.
	ExportFormatMarkdown ExportFormat = "markdown"

	// ExportFormatRevealJS is Reveal.js HTML output.
	ExportFormatRevealJS ExportFormat = "revealjs"

	// ExportFormatMarp is Marp Markdown output.
	ExportFormatMarp ExportFormat = "marp"
)

// ExportOptions contains format-specific export options.
type ExportOptions struct {
	// Quality is the image quality (0-100, for PNG/JPEG).
	Quality int `json:"quality,omitempty"`

	// Scale is the image scale factor.
	Scale float64 `json:"scale,omitempty"`

	// PageSize is the page size for PDF (e.g., "A4", "Letter").
	PageSize string `json:"pageSize,omitempty"`

	// Landscape controls page orientation for PDF.
	Landscape bool `json:"landscape,omitempty"`

	// IncludeSpeakerNotes controls whether to include speaker notes.
	IncludeSpeakerNotes bool `json:"includeSpeakerNotes,omitempty"`

	// IncludeQuality controls whether to include quality metadata.
	IncludeQuality bool `json:"includeQuality,omitempty"`

	// Theme is the theme override for this export.
	Theme string `json:"theme,omitempty"`

	// Template is the template file for web/revealjs exports.
	Template string `json:"template,omitempty"`
}
