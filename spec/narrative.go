package spec

// Narrative defines the storyline structure of a presentation.
type Narrative struct {
	// Storyline is a brief description of the presentation's narrative arc.
	Storyline string `json:"storyline,omitempty"`

	// Sections define the logical sections of the presentation.
	Sections []Section `json:"sections,omitempty"`

	// Audience describes the target audience.
	Audience string `json:"audience,omitempty"`

	// Objective states the presentation objective.
	Objective string `json:"objective,omitempty"`

	// Duration is the expected presentation duration (e.g., "30m", "1h").
	Duration string `json:"duration,omitempty"`
}

// Section represents a logical section in the narrative.
type Section struct {
	// ID is the unique identifier for this section.
	ID string `json:"id"`

	// Title is the section title.
	Title string `json:"title"`

	// Description provides additional context for this section.
	Description string `json:"description,omitempty"`

	// Order is the section order (1-based).
	Order int `json:"order,omitempty"`

	// KeyPoints are the main points to cover in this section.
	KeyPoints []string `json:"keyPoints,omitempty"`
}
