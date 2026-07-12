package spec

import (
	"encoding/json"
	"testing"
)

func TestPresentationSpec_JSONRoundTrip(t *testing.T) {
	pres := PresentationSpec{
		Version: "1.0",
		Metadata: Metadata{
			Title:       "Test Presentation",
			Description: "A test presentation for unit testing",
			Author:      "Test Author",
			Date:        "2025-01-15",
			Tags:        []string{"test", "demo"},
		},
		Theme: &Theme{
			Name:        "corporate",
			Density:     DensityNormal,
			AspectRatio: AspectRatio16x9,
		},
		Slides: []SlideSpec{
			{
				ID:    "title",
				Type:  SlideTypeTitle,
				Title: "Welcome",
			},
			{
				ID:         "content",
				Type:       SlideTypeContent,
				Title:      "Main Content",
				KeyMessage: "This is the key message",
				Widgets: []Widget{
					{
						ID:    "metric1",
						Type:  WidgetTypeMetricCard,
						Title: "Revenue",
						Value: "$1.2M",
						Unit:  "USD",
					},
				},
			},
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(pres)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	// Unmarshal back
	var decoded PresentationSpec
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	// Verify fields
	if decoded.Version != pres.Version {
		t.Errorf("version mismatch: got %q, want %q", decoded.Version, pres.Version)
	}

	if decoded.Metadata.Title != pres.Metadata.Title {
		t.Errorf("title mismatch: got %q, want %q", decoded.Metadata.Title, pres.Metadata.Title)
	}

	if len(decoded.Slides) != len(pres.Slides) {
		t.Errorf("slide count mismatch: got %d, want %d", len(decoded.Slides), len(pres.Slides))
	}

	if decoded.Theme.Density != DensityNormal {
		t.Errorf("density mismatch: got %q, want %q", decoded.Theme.Density, DensityNormal)
	}
}

func TestSlideSpec_JSONRoundTrip(t *testing.T) {
	slide := SlideSpec{
		ID:           "dashboard",
		SectionID:    "metrics",
		Type:         SlideTypeDashboard,
		Title:        "KPI Dashboard",
		KeyMessage:   "All metrics trending positive",
		SpeakerNotes: "Emphasize the growth trend",
		Layout: &Layout{
			Template: LayoutTemplateKPIRowContentGrid,
			Regions: []Region{
				{ID: "kpi-row", Role: "kpi"},
				{ID: "main-content", Role: "content"},
			},
		},
		Widgets: []Widget{
			{
				ID:         "revenue",
				Type:       WidgetTypeMetricCard,
				Title:      "Revenue",
				Value:      "1.5M",
				Unit:       "$",
				Status:     StatusGood,
				Trend:      TrendUp,
				TrendValue: "+12%",
			},
		},
		Quality: &Quality{
			DataFreshness: "real-time",
			Confidence:    0.95,
		},
	}

	data, err := json.Marshal(slide)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded SlideSpec
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.ID != slide.ID {
		t.Errorf("ID mismatch: got %q, want %q", decoded.ID, slide.ID)
	}

	if decoded.Type != SlideTypeDashboard {
		t.Errorf("type mismatch: got %q, want %q", decoded.Type, SlideTypeDashboard)
	}

	if decoded.Layout.Template != LayoutTemplateKPIRowContentGrid {
		t.Errorf("layout template mismatch: got %q, want %q",
			decoded.Layout.Template, LayoutTemplateKPIRowContentGrid)
	}

	if len(decoded.Widgets) != 1 {
		t.Fatalf("expected 1 widget, got %d", len(decoded.Widgets))
	}

	if decoded.Widgets[0].Status != StatusGood {
		t.Errorf("widget status mismatch: got %q, want %q",
			decoded.Widgets[0].Status, StatusGood)
	}
}

func TestWidget_JSONRoundTrip(t *testing.T) {
	tests := []struct {
		name   string
		widget Widget
	}{
		{
			name: "metric_card",
			widget: Widget{
				ID:     "metric1",
				Type:   WidgetTypeMetricCard,
				Title:  "Revenue",
				Value:  "1.2M",
				Unit:   "$",
				Target: "1M",
				Status: StatusGood,
			},
		},
		{
			name: "chart",
			widget: Widget{
				ID:        "chart1",
				Type:      WidgetTypeChart,
				Title:     "Sales Trend",
				ChartType: ChartTypeLine,
				Data: &ChartData{
					Labels: []string{"Q1", "Q2", "Q3", "Q4"},
					Series: []ChartSeries{
						{Name: "Sales", Values: []float64{100, 150, 200, 250}},
					},
				},
			},
		},
		{
			name: "table",
			widget: Widget{
				ID:    "table1",
				Type:  WidgetTypeTable,
				Title: "Feature Status",
				Columns: []TableColumn{
					{Key: "feature", Header: "Feature", Width: "50%"},
					{Key: "status", Header: "Status", Width: "50%"},
				},
				Rows: []map[string]any{
					{"feature": "Auth", "status": "Complete"},
				},
			},
		},
		{
			name: "risk_list",
			widget: Widget{
				ID:    "risks",
				Type:  WidgetTypeRiskList,
				Title: "Open Risks",
				Items: []ListItem{
					{
						Text:       "API rate limits",
						Severity:   SeverityHigh,
						Mitigation: "Implement caching",
						Owner:      "Platform Team",
					},
				},
			},
		},
		{
			name: "checklist",
			widget: Widget{
				ID:    "checklist1",
				Type:  WidgetTypeChecklist,
				Title: "Launch Checklist",
				Items: []ListItem{
					{Text: "Code review", Checked: true},
					{Text: "Security audit", Checked: false},
				},
			},
		},
		{
			name: "callout",
			widget: Widget{
				ID:          "callout1",
				Type:        WidgetTypeCallout,
				Title:       "Important",
				CalloutType: CalloutTypeWarning,
				Content:     "Please review carefully",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.widget)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}

			var decoded Widget
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}

			if decoded.Type != tt.widget.Type {
				t.Errorf("type mismatch: got %q, want %q", decoded.Type, tt.widget.Type)
			}

			if decoded.ID != tt.widget.ID {
				t.Errorf("ID mismatch: got %q, want %q", decoded.ID, tt.widget.ID)
			}
		})
	}
}

func TestDataSource_JSONRoundTrip(t *testing.T) {
	ds := DataSource{
		ID:       "api-metrics",
		Type:     DataSourceTypeAPI,
		Endpoint: "https://api.example.com/metrics",
		Headers: map[string]string{
			"Accept": "application/json",
		},
		Auth: &DataSourceAuth{
			Type:  AuthTypeBearer,
			Token: "secret-token",
		},
		RefreshInterval: "5m",
		Cache:           true,
	}

	data, err := json.Marshal(ds)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded DataSource
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.Type != DataSourceTypeAPI {
		t.Errorf("type mismatch: got %q, want %q", decoded.Type, DataSourceTypeAPI)
	}

	if decoded.Auth.Type != AuthTypeBearer {
		t.Errorf("auth type mismatch: got %q, want %q", decoded.Auth.Type, AuthTypeBearer)
	}

	if decoded.Endpoint != ds.Endpoint {
		t.Errorf("endpoint mismatch: got %q, want %q", decoded.Endpoint, ds.Endpoint)
	}
}

func TestNarrative_JSONRoundTrip(t *testing.T) {
	narrative := Narrative{
		Storyline: "This presentation covers our Q1 progress",
		Sections: []Section{
			{ID: "intro", Title: "Introduction"},
			{ID: "metrics", Title: "Key Metrics"},
			{ID: "next-steps", Title: "Next Steps"},
		},
	}

	data, err := json.Marshal(narrative)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded Narrative
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.Storyline != narrative.Storyline {
		t.Errorf("storyline mismatch")
	}

	if len(decoded.Sections) != 3 {
		t.Errorf("expected 3 sections, got %d", len(decoded.Sections))
	}
}

func TestExportConfig_JSONRoundTrip(t *testing.T) {
	export := ExportConfig{
		Targets: []ExportTarget{
			{
				Format:     ExportFormatPDF,
				OutputPath: "./output/presentation.pdf",
				Options: &ExportOptions{
					PageSize:  "Letter",
					Landscape: true,
				},
			},
			{
				Format:     ExportFormatWeb,
				OutputPath: "./output/web/",
			},
		},
	}

	data, err := json.Marshal(export)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded ExportConfig
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(decoded.Targets) != 2 {
		t.Errorf("expected 2 targets, got %d", len(decoded.Targets))
	}

	if decoded.Targets[0].Format != ExportFormatPDF {
		t.Errorf("first target format mismatch: got %q, want %q",
			decoded.Targets[0].Format, ExportFormatPDF)
	}
}
