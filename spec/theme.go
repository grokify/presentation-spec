package spec

// Theme defines visual styling and branding for a presentation.
type Theme struct {
	// Name is the theme name (e.g., "corporate", "minimal", "dark").
	Name string `json:"name,omitempty"`

	// Density controls information density ("compact", "normal", "spacious").
	Density Density `json:"density,omitempty"`

	// AspectRatio is the slide aspect ratio (e.g., "16:9", "4:3").
	AspectRatio AspectRatio `json:"aspectRatio,omitempty"`

	// Brand contains branding configuration.
	Brand *Brand `json:"brand,omitempty"`

	// Colors defines the color palette.
	Colors *ColorPalette `json:"colors,omitempty"`

	// Typography defines font settings.
	Typography *Typography `json:"typography,omitempty"`
}

// Density controls information density.
type Density string

const (
	DensityCompact  Density = "compact"
	DensityNormal   Density = "normal"
	DensitySpacious Density = "spacious"
)

// AspectRatio is the slide aspect ratio.
type AspectRatio string

const (
	AspectRatio16x9 AspectRatio = "16:9"
	AspectRatio4x3  AspectRatio = "4:3"
	AspectRatio1x1  AspectRatio = "1:1"
)

// Brand contains branding configuration.
type Brand struct {
	// Logo is the logo image URL or path.
	Logo string `json:"logo,omitempty"`

	// LogoPosition is where to place the logo.
	LogoPosition LogoPosition `json:"logoPosition,omitempty"`

	// CompanyName is the company name for branding.
	CompanyName string `json:"companyName,omitempty"`

	// PrimaryColor is the primary brand color (hex).
	PrimaryColor string `json:"primaryColor,omitempty"`

	// SecondaryColor is the secondary brand color (hex).
	SecondaryColor string `json:"secondaryColor,omitempty"`
}

// LogoPosition specifies logo placement.
type LogoPosition string

const (
	LogoPositionTopLeft     LogoPosition = "top-left"
	LogoPositionTopRight    LogoPosition = "top-right"
	LogoPositionBottomLeft  LogoPosition = "bottom-left"
	LogoPositionBottomRight LogoPosition = "bottom-right"
)

// ColorPalette defines the color scheme.
type ColorPalette struct {
	// Primary is the primary color (hex).
	Primary string `json:"primary,omitempty"`

	// Secondary is the secondary color (hex).
	Secondary string `json:"secondary,omitempty"`

	// Accent is the accent color (hex).
	Accent string `json:"accent,omitempty"`

	// Background is the background color (hex).
	Background string `json:"background,omitempty"`

	// Surface is the surface color for cards, etc. (hex).
	Surface string `json:"surface,omitempty"`

	// Text is the primary text color (hex).
	Text string `json:"text,omitempty"`

	// TextMuted is the muted text color (hex).
	TextMuted string `json:"textMuted,omitempty"`

	// Success is the success indicator color (hex).
	Success string `json:"success,omitempty"`

	// Warning is the warning indicator color (hex).
	Warning string `json:"warning,omitempty"`

	// Error is the error indicator color (hex).
	Error string `json:"error,omitempty"`
}

// Typography defines font settings.
type Typography struct {
	// FontFamily is the primary font family.
	FontFamily string `json:"fontFamily,omitempty"`

	// HeadingFontFamily is the heading font family.
	HeadingFontFamily string `json:"headingFontFamily,omitempty"`

	// CodeFontFamily is the monospace font family.
	CodeFontFamily string `json:"codeFontFamily,omitempty"`

	// BaseFontSize is the base font size (e.g., "16px").
	BaseFontSize string `json:"baseFontSize,omitempty"`

	// LineHeight is the base line height (e.g., "1.5").
	LineHeight string `json:"lineHeight,omitempty"`
}
