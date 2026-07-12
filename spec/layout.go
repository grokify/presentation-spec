package spec

// Layout defines the layout configuration for a slide.
type Layout struct {
	// Template is the layout template name.
	Template LayoutTemplate `json:"template"`

	// Regions defines the layout regions.
	Regions []Region `json:"regions,omitempty"`

	// Gap is the gap between regions (e.g., "16px", "1rem").
	Gap string `json:"gap,omitempty"`

	// Padding is the padding around the layout (e.g., "24px", "2rem").
	Padding string `json:"padding,omitempty"`
}

// LayoutTemplate enumerates predefined layout templates.
type LayoutTemplate string

const (
	// LayoutTemplateSingle is a single-region layout.
	LayoutTemplateSingle LayoutTemplate = "single"

	// LayoutTemplateTwoColumn is a two-column layout.
	LayoutTemplateTwoColumn LayoutTemplate = "two_column"

	// LayoutTemplateThreeColumn is a three-column layout.
	LayoutTemplateThreeColumn LayoutTemplate = "three_column"

	// LayoutTemplateFourColumn is a four-column layout.
	LayoutTemplateFourColumn LayoutTemplate = "four_column"

	// LayoutTemplateHeaderContent is a header with content below.
	LayoutTemplateHeaderContent LayoutTemplate = "header_content"

	// LayoutTemplateContentFooter is content with footer below.
	LayoutTemplateContentFooter LayoutTemplate = "content_footer"

	// LayoutTemplateKPIRowContentGrid is a row of KPIs with content grid below.
	LayoutTemplateKPIRowContentGrid LayoutTemplate = "kpi_row_content_grid"

	// LayoutTemplateMetricDashboard is a dashboard layout for metrics.
	LayoutTemplateMetricDashboard LayoutTemplate = "metric_dashboard"

	// LayoutTemplateSidebar is a sidebar with main content.
	LayoutTemplateSidebar LayoutTemplate = "sidebar"

	// LayoutTemplateGrid is a flexible grid layout.
	LayoutTemplateGrid LayoutTemplate = "grid"

	// LayoutTemplateCustom allows custom region definitions.
	LayoutTemplateCustom LayoutTemplate = "custom"
)

// Region defines a region within a layout.
type Region struct {
	// ID is the unique identifier for this region.
	ID string `json:"id"`

	// Role describes the semantic role of this region.
	Role RegionRole `json:"role,omitempty"`

	// Column is the starting column (1-based, for grid layouts).
	Column int `json:"column,omitempty"`

	// Row is the starting row (1-based, for grid layouts).
	Row int `json:"row,omitempty"`

	// ColumnSpan is the number of columns this region spans.
	ColumnSpan int `json:"columnSpan,omitempty"`

	// RowSpan is the number of rows this region spans.
	RowSpan int `json:"rowSpan,omitempty"`

	// Width is the region width (e.g., "50%", "300px").
	Width string `json:"width,omitempty"`

	// Height is the region height (e.g., "auto", "200px").
	Height string `json:"height,omitempty"`

	// Align is the content alignment within the region.
	Align Alignment `json:"align,omitempty"`

	// VerticalAlign is the vertical content alignment.
	VerticalAlign VerticalAlignment `json:"verticalAlign,omitempty"`
}

// RegionRole describes the semantic role of a region.
type RegionRole string

const (
	RegionRoleHeader  RegionRole = "header"
	RegionRoleContent RegionRole = "content"
	RegionRoleFooter  RegionRole = "footer"
	RegionRoleSidebar RegionRole = "sidebar"
	RegionRoleKPI     RegionRole = "kpi"
	RegionRoleChart   RegionRole = "chart"
	RegionRoleTable   RegionRole = "table"
)

// VerticalAlignment specifies vertical alignment.
type VerticalAlignment string

const (
	VerticalAlignTop    VerticalAlignment = "top"
	VerticalAlignMiddle VerticalAlignment = "middle"
	VerticalAlignBottom VerticalAlignment = "bottom"
)
