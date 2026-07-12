package spec

// Widget represents a content element on a slide.
// The Type field acts as a discriminator for the widget variant.
type Widget struct {
	// Type is the widget type discriminator.
	Type WidgetType `json:"type"`

	// ID is the unique identifier for this widget.
	ID string `json:"id,omitempty"`

	// RegionID links this widget to a layout region.
	RegionID string `json:"regionId,omitempty"`

	// DataSourceID references an external data source.
	DataSourceID string `json:"dataSourceId,omitempty"`

	// Title is the widget title.
	Title string `json:"title,omitempty"`

	// --- MetricCard fields ---

	// Value is the primary metric value (for metric_card).
	Value string `json:"value,omitempty"`

	// Unit is the unit of measurement (for metric_card).
	Unit string `json:"unit,omitempty"`

	// Trend is the trend direction (for metric_card).
	Trend Trend `json:"trend,omitempty"`

	// TrendValue is the trend percentage or delta (for metric_card).
	TrendValue string `json:"trendValue,omitempty"`

	// Status is the status indicator (for metric_card, risk_list items).
	Status Status `json:"status,omitempty"`

	// Target is the target value (for metric_card).
	Target string `json:"target,omitempty"`

	// Sparkline contains sparkline data points (for metric_card).
	Sparkline []float64 `json:"sparkline,omitempty"`

	// --- ContentBlock fields ---

	// Content is the markdown content (for content_block, callout).
	Content string `json:"content,omitempty"`

	// Format specifies the content format (for content_block).
	Format ContentFormat `json:"format,omitempty"`

	// --- List fields (risk_list, decision_list, checklist) ---

	// Items contains list items (for risk_list, decision_list, checklist).
	Items []ListItem `json:"items,omitempty"`

	// --- Chart fields ---

	// ChartType specifies the chart type (for chart).
	ChartType ChartType `json:"chartType,omitempty"`

	// Data contains chart data (for chart, table).
	Data *ChartData `json:"data,omitempty"`

	// Options contains chart rendering options (for chart).
	Options *ChartOptions `json:"options,omitempty"`

	// --- Table fields ---

	// Columns defines table columns (for table).
	Columns []TableColumn `json:"columns,omitempty"`

	// Rows contains table row data (for table).
	Rows []map[string]any `json:"rows,omitempty"`

	// --- Diagram fields ---

	// DiagramType specifies the diagram type (for diagram).
	DiagramType DiagramType `json:"diagramType,omitempty"`

	// Source is the diagram source code (for diagram).
	Source string `json:"source,omitempty"`

	// --- Callout fields ---

	// CalloutType specifies the callout style (for callout).
	CalloutType CalloutType `json:"calloutType,omitempty"`

	// --- Quote fields ---

	// Quote is the quoted text (for quote).
	Quote string `json:"quote,omitempty"`

	// Attribution is the quote attribution (for quote).
	Attribution string `json:"attribution,omitempty"`

	// --- Image fields ---

	// Src is the image source URL or path (for image).
	Src string `json:"src,omitempty"`

	// Alt is the image alt text (for image).
	Alt string `json:"alt,omitempty"`

	// Caption is the image caption (for image).
	Caption string `json:"caption,omitempty"`

	// --- Code fields ---

	// Language is the programming language (for code).
	Language string `json:"language,omitempty"`

	// Code is the code content (for code).
	Code string `json:"code,omitempty"`

	// Highlight specifies lines to highlight (for code).
	Highlight []int `json:"highlight,omitempty"`
}

// WidgetType enumerates the supported widget types.
type WidgetType string

const (
	// WidgetTypeMetricCard displays a single metric with optional trend.
	WidgetTypeMetricCard WidgetType = "metric_card"

	// WidgetTypeContentBlock displays formatted text content.
	WidgetTypeContentBlock WidgetType = "content_block"

	// WidgetTypeRiskList displays a list of risks with severity.
	WidgetTypeRiskList WidgetType = "risk_list"

	// WidgetTypeDecisionList displays decisions requiring action.
	WidgetTypeDecisionList WidgetType = "decision_list"

	// WidgetTypeChecklist displays a checklist with completion status.
	WidgetTypeChecklist WidgetType = "checklist"

	// WidgetTypeChart displays data visualizations.
	WidgetTypeChart WidgetType = "chart"

	// WidgetTypeTable displays tabular data.
	WidgetTypeTable WidgetType = "table"

	// WidgetTypeDiagram displays diagrams (Mermaid, PlantUML, etc.).
	WidgetTypeDiagram WidgetType = "diagram"

	// WidgetTypeCallout displays highlighted information.
	WidgetTypeCallout WidgetType = "callout"

	// WidgetTypeQuote displays a quotation.
	WidgetTypeQuote WidgetType = "quote"

	// WidgetTypeImage displays an image.
	WidgetTypeImage WidgetType = "image"

	// WidgetTypeCode displays code with syntax highlighting.
	WidgetTypeCode WidgetType = "code"
)

// Status represents a status indicator value.
type Status string

const (
	StatusGood     Status = "good"
	StatusWarning  Status = "warning"
	StatusCritical Status = "critical"
	StatusNeutral  Status = "neutral"
)

// Trend represents a trend direction.
type Trend string

const (
	TrendUp   Trend = "up"
	TrendDown Trend = "down"
	TrendFlat Trend = "flat"
)

// ContentFormat specifies the format of content.
type ContentFormat string

const (
	ContentFormatMarkdown ContentFormat = "markdown"
	ContentFormatHTML     ContentFormat = "html"
	ContentFormatPlain    ContentFormat = "plain"
)

// ListItem represents an item in a list widget.
type ListItem struct {
	// Text is the item text.
	Text string `json:"text"`

	// Severity is the risk severity (for risk_list).
	Severity Severity `json:"severity,omitempty"`

	// Status is the item status.
	Status Status `json:"status,omitempty"`

	// Mitigation is the risk mitigation (for risk_list).
	Mitigation string `json:"mitigation,omitempty"`

	// Owner is the responsible party (for risk_list, decision_list).
	Owner string `json:"owner,omitempty"`

	// Deadline is the due date (for decision_list).
	Deadline string `json:"deadline,omitempty"`

	// Checked indicates completion status (for checklist).
	Checked bool `json:"checked,omitempty"`

	// Description provides additional detail.
	Description string `json:"description,omitempty"`
}

// Severity represents risk severity levels.
type Severity string

const (
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
	SeverityHigh     Severity = "high"
	SeverityCritical Severity = "critical"
)

// ChartType specifies the type of chart.
type ChartType string

const (
	ChartTypeBar         ChartType = "bar"
	ChartTypeLine        ChartType = "line"
	ChartTypePie         ChartType = "pie"
	ChartTypeDonut       ChartType = "donut"
	ChartTypeArea        ChartType = "area"
	ChartTypeScatter     ChartType = "scatter"
	ChartTypeRadar       ChartType = "radar"
	ChartTypeGauge       ChartType = "gauge"
	ChartTypeHeatmap     ChartType = "heatmap"
	ChartTypeTreemap     ChartType = "treemap"
	ChartTypeSankey      ChartType = "sankey"
	ChartTypeWaterfall   ChartType = "waterfall"
	ChartTypeFunnel      ChartType = "funnel"
	ChartTypeCandlestick ChartType = "candlestick"
)

// ChartData contains data for chart rendering.
type ChartData struct {
	// Labels are the data labels (x-axis for most charts).
	Labels []string `json:"labels,omitempty"`

	// Series contains the data series.
	Series []ChartSeries `json:"series,omitempty"`
}

// ChartSeries represents a single data series in a chart.
type ChartSeries struct {
	// Name is the series name.
	Name string `json:"name,omitempty"`

	// Values are the data values.
	Values []float64 `json:"values"`

	// Color is the series color (optional).
	Color string `json:"color,omitempty"`
}

// ChartOptions contains chart rendering options.
type ChartOptions struct {
	// ShowLegend controls legend visibility.
	ShowLegend bool `json:"showLegend,omitempty"`

	// ShowGrid controls grid visibility.
	ShowGrid bool `json:"showGrid,omitempty"`

	// ShowLabels controls data label visibility.
	ShowLabels bool `json:"showLabels,omitempty"`

	// Stacked enables stacked chart mode.
	Stacked bool `json:"stacked,omitempty"`

	// XAxisLabel is the x-axis label.
	XAxisLabel string `json:"xAxisLabel,omitempty"`

	// YAxisLabel is the y-axis label.
	YAxisLabel string `json:"yAxisLabel,omitempty"`
}

// TableColumn defines a table column.
type TableColumn struct {
	// Key is the column key (matches row data keys).
	Key string `json:"key"`

	// Header is the column header text.
	Header string `json:"header"`

	// Width is the column width (e.g., "100px", "20%").
	Width string `json:"width,omitempty"`

	// Align is the text alignment.
	Align Alignment `json:"align,omitempty"`
}

// Alignment specifies text alignment.
type Alignment string

const (
	AlignLeft   Alignment = "left"
	AlignCenter Alignment = "center"
	AlignRight  Alignment = "right"
)

// DiagramType specifies the diagram format.
type DiagramType string

const (
	DiagramTypeMermaid  DiagramType = "mermaid"
	DiagramTypePlantUML DiagramType = "plantuml"
	DiagramTypeGraphviz DiagramType = "graphviz"
	DiagramTypeD2       DiagramType = "d2"
)

// CalloutType specifies the callout style.
type CalloutType string

const (
	CalloutTypeInfo    CalloutType = "info"
	CalloutTypeWarning CalloutType = "warning"
	CalloutTypeError   CalloutType = "error"
	CalloutTypeSuccess CalloutType = "success"
	CalloutTypeTip     CalloutType = "tip"
)
