package schema

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/grokify/presentation-spec/spec"
)

func TestValidateJSON_ValidSpec(t *testing.T) {
	validJSON := []byte(`{
		"version": "1.0",
		"metadata": {
			"title": "Test Presentation",
			"description": "A valid test presentation"
		},
		"slides": [
			{
				"id": "title-slide",
				"type": "title",
				"title": "Welcome"
			},
			{
				"id": "content-slide",
				"type": "content",
				"title": "Content",
				"keyMessage": "Important information"
			}
		]
	}`)

	result, err := ValidateJSON(validJSON)
	if err != nil {
		t.Fatalf("ValidateJSON failed: %v", err)
	}

	if !result.Valid {
		t.Errorf("expected valid spec, got errors: %v", result.Errors)
	}
}

func TestValidateJSON_InvalidJSON(t *testing.T) {
	invalidJSON := []byte(`{invalid json}`)

	_, err := ValidateJSON(invalidJSON)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestValidateFile(t *testing.T) {
	tmpDir := t.TempDir()
	specPath := filepath.Join(tmpDir, "valid.json")

	specJSON := `{
		"version": "1.0",
		"metadata": {
			"title": "File Test"
		},
		"slides": []
	}`

	if err := os.WriteFile(specPath, []byte(specJSON), 0600); err != nil {
		t.Fatal(err)
	}

	result, err := ValidateFile(specPath)
	if err != nil {
		t.Fatalf("ValidateFile failed: %v", err)
	}

	if !result.Valid {
		t.Errorf("expected valid spec, got errors: %v", result.Errors)
	}
}

func TestValidateFile_NotFound(t *testing.T) {
	_, err := ValidateFile("/nonexistent/spec.json")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestValidateSpec_ValidStruct(t *testing.T) {
	pres := spec.PresentationSpec{
		Version: "1.0",
		Metadata: spec.Metadata{
			Title:       "Struct Test",
			Description: "Test from struct",
		},
		Slides: []spec.SlideSpec{
			{
				ID:    "slide-1",
				Type:  spec.SlideTypeTitle,
				Title: "Title",
			},
		},
	}

	result, err := ValidateSpec(pres)
	if err != nil {
		t.Fatalf("ValidateSpec failed: %v", err)
	}

	if !result.Valid {
		t.Errorf("expected valid spec, got errors: %v", result.Errors)
	}
}

func TestValidateJSON_WithWidgets(t *testing.T) {
	specJSON := []byte(`{
		"version": "1.0",
		"metadata": {"title": "Widget Test"},
		"slides": [
			{
				"id": "dashboard",
				"type": "dashboard",
				"title": "Dashboard",
				"widgets": [
					{
						"id": "metric1",
						"type": "metric_card",
						"title": "Revenue",
						"value": "1.2M",
						"unit": "$",
						"status": "good"
					},
					{
						"id": "chart1",
						"type": "chart",
						"chartType": "bar",
						"data": {
							"labels": ["Q1", "Q2", "Q3"],
							"series": [
								{"name": "Sales", "values": [100, 150, 200]}
							]
						}
					}
				]
			}
		]
	}`)

	result, err := ValidateJSON(specJSON)
	if err != nil {
		t.Fatalf("ValidateJSON failed: %v", err)
	}

	if !result.Valid {
		t.Errorf("expected valid spec with widgets, got errors: %v", result.Errors)
	}
}

func TestValidateJSON_WithDataSources(t *testing.T) {
	specJSON := []byte(`{
		"version": "1.0",
		"metadata": {"title": "DataSource Test"},
		"slides": [],
		"dataSources": [
			{
				"id": "api-data",
				"type": "api",
				"endpoint": "https://api.example.com/data",
				"headers": {"Accept": "application/json"},
				"auth": {
					"type": "bearer",
					"token": "secret-token"
				}
			},
			{
				"id": "file-data",
				"type": "json",
				"path": "./data/metrics.json"
			}
		]
	}`)

	result, err := ValidateJSON(specJSON)
	if err != nil {
		t.Fatalf("ValidateJSON failed: %v", err)
	}

	if !result.Valid {
		t.Errorf("expected valid spec with data sources, got errors: %v", result.Errors)
	}
}

func TestValidateJSON_WithTheme(t *testing.T) {
	specJSON := []byte(`{
		"version": "1.0",
		"metadata": {"title": "Theme Test"},
		"slides": [],
		"theme": {
			"name": "corporate",
			"density": "compact",
			"aspectRatio": "16:9",
			"brand": {
				"logo": "/assets/logo.png",
				"primaryColor": "#1a73e8",
				"secondaryColor": "#34a853"
			},
			"colors": {
				"primary": "#1a73e8",
				"secondary": "#34a853",
				"background": "#ffffff"
			},
			"typography": {
				"headingFontFamily": "Inter",
				"fontFamily": "Open Sans"
			}
		}
	}`)

	result, err := ValidateJSON(specJSON)
	if err != nil {
		t.Fatalf("ValidateJSON failed: %v", err)
	}

	if !result.Valid {
		t.Errorf("expected valid spec with theme, got errors: %v", result.Errors)
	}
}

func TestPresentationSchemaJSON(t *testing.T) {
	schemaBytes := PresentationSchemaJSON()

	if len(schemaBytes) == 0 {
		t.Error("expected non-empty schema")
	}

	// Verify it's valid JSON by trying to use it
	result, err := ValidateJSON([]byte(`{"version": "1.0", "metadata": {"title": "test"}, "slides": []}`))
	if err != nil {
		t.Errorf("schema should be usable: %v", err)
	}

	if !result.Valid {
		t.Errorf("minimal spec should be valid: %v", result.Errors)
	}
}
