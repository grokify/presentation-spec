package spec

// SlideSpec represents a single slide in the presentation.
type SlideSpec struct {
	// ID is the unique identifier for this slide.
	ID string `json:"id"`

	// SectionID links this slide to a narrative section.
	SectionID string `json:"sectionId,omitempty"`

	// Type is the slide type (e.g., "title", "content", "dashboard").
	Type SlideType `json:"type"`

	// Title is the slide title.
	Title string `json:"title,omitempty"`

	// KeyMessage is the main takeaway for the audience.
	KeyMessage string `json:"keyMessage,omitempty"`

	// Layout defines the slide layout configuration.
	Layout *Layout `json:"layout,omitempty"`

	// Widgets are the content elements on the slide.
	Widgets []Widget `json:"widgets,omitempty"`

	// SpeakerNotes are notes visible only to the presenter.
	SpeakerNotes string `json:"speakerNotes,omitempty"`

	// Quality contains quality assessment information.
	Quality *Quality `json:"quality,omitempty"`
}

// SlideType enumerates the supported slide types.
type SlideType string

const (
	// SlideTypeTitle is a title slide.
	SlideTypeTitle SlideType = "title"

	// SlideTypeContent is a standard content slide.
	SlideTypeContent SlideType = "content"

	// SlideTypeDashboard is a data-heavy dashboard slide.
	SlideTypeDashboard SlideType = "dashboard"

	// SlideTypeSection is a section divider slide.
	SlideTypeSection SlideType = "section"

	// SlideTypeSummary is a summary or conclusion slide.
	SlideTypeSummary SlideType = "summary"

	// SlideTypeBlank is a blank slide.
	SlideTypeBlank SlideType = "blank"
)

// Quality contains quality assessment information for a slide.
type Quality struct {
	// DataFreshness indicates how recent the data is.
	DataFreshness string `json:"dataFreshness,omitempty"`

	// Confidence is the confidence level (0.0 to 1.0).
	Confidence float64 `json:"confidence,omitempty"`

	// Sources lists the data sources used.
	Sources []string `json:"sources,omitempty"`

	// LastUpdated is the timestamp of the last update (ISO 8601).
	LastUpdated string `json:"lastUpdated,omitempty"`
}
