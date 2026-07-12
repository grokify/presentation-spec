// Package spec defines the Go types for PresentationSpec, a format-agnostic
// intermediate representation for presentations.
package spec

// PresentationSpec is the root type for a presentation specification.
// It contains all metadata, content, and configuration needed to render
// a presentation to multiple output formats deterministically.
type PresentationSpec struct {
	// Version is the specification version (e.g., "1.0.0").
	Version string `json:"version"`

	// Metadata contains presentation-level metadata.
	Metadata Metadata `json:"metadata"`

	// Theme defines visual styling and branding.
	Theme *Theme `json:"theme,omitempty"`

	// Narrative defines the storyline structure.
	Narrative *Narrative `json:"narrative,omitempty"`

	// Slides is the ordered list of slides in the presentation.
	Slides []SlideSpec `json:"slides"`

	// DataSources defines external data sources referenced by widgets.
	DataSources []DataSource `json:"dataSources,omitempty"`

	// Exports configures output format targets.
	Exports *ExportConfig `json:"exports,omitempty"`
}

// Metadata contains presentation-level metadata.
type Metadata struct {
	// Title is the presentation title.
	Title string `json:"title"`

	// Subtitle is an optional subtitle.
	Subtitle string `json:"subtitle,omitempty"`

	// Author is the presentation author.
	Author string `json:"author,omitempty"`

	// Date is the presentation date (ISO 8601 format).
	Date string `json:"date,omitempty"`

	// Description is a brief description of the presentation.
	Description string `json:"description,omitempty"`

	// Tags are optional categorization tags.
	Tags []string `json:"tags,omitempty"`
}
