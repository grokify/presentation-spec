package templ

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/grokify/presentation-spec/spec"
)

func TestLoadSpec(t *testing.T) {
	tmpDir := t.TempDir()
	specPath := filepath.Join(tmpDir, "test.json")

	specJSON := `{
		"version": "1.0",
		"metadata": {
			"title": "Test Presentation",
			"description": "A test presentation"
		},
		"slides": [
			{
				"id": "slide-1",
				"type": "title",
				"title": "Hello World"
			}
		]
	}`

	if err := os.WriteFile(specPath, []byte(specJSON), 0600); err != nil {
		t.Fatalf("failed to write spec file: %v", err)
	}

	pres, err := LoadSpec(specPath)
	if err != nil {
		t.Fatalf("LoadSpec failed: %v", err)
	}

	if pres.Version != "1.0" {
		t.Errorf("expected version '1.0', got %q", pres.Version)
	}

	if pres.Metadata.Title != "Test Presentation" {
		t.Errorf("expected title 'Test Presentation', got %q", pres.Metadata.Title)
	}

	if len(pres.Slides) != 1 {
		t.Errorf("expected 1 slide, got %d", len(pres.Slides))
	}

	if pres.Slides[0].ID != "slide-1" {
		t.Errorf("expected slide ID 'slide-1', got %q", pres.Slides[0].ID)
	}
}

func TestLoadSpec_FileNotFound(t *testing.T) {
	_, err := LoadSpec("/nonexistent/spec.json")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestLoadSpec_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	specPath := filepath.Join(tmpDir, "invalid.json")

	if err := os.WriteFile(specPath, []byte(`{invalid json}`), 0600); err != nil {
		t.Fatal(err)
	}

	_, err := LoadSpec(specPath)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestRenderer_RenderToWriter(t *testing.T) {
	pres := spec.PresentationSpec{
		Version: "1.0",
		Metadata: spec.Metadata{
			Title:       "Test",
			Description: "Test presentation",
		},
		Slides: []spec.SlideSpec{
			{
				ID:    "title",
				Type:  spec.SlideTypeTitle,
				Title: "Hello World",
			},
		},
	}

	renderer := NewRenderer()
	ctx := context.Background()

	var buf bytes.Buffer
	if err := renderer.RenderToWriter(ctx, pres, &buf); err != nil {
		t.Fatalf("RenderToWriter failed: %v", err)
	}

	html := buf.String()

	if !strings.Contains(html, "<!doctype html>") {
		t.Error("expected HTML doctype")
	}

	if !strings.Contains(html, "<title>Test</title>") {
		t.Error("expected title element")
	}

	if !strings.Contains(html, "Hello World") {
		t.Error("expected slide title in output")
	}

	if !strings.Contains(html, "mermaid") {
		t.Error("expected Mermaid script reference")
	}
}

func TestRenderer_RenderToFile(t *testing.T) {
	pres := spec.PresentationSpec{
		Version: "1.0",
		Metadata: spec.Metadata{
			Title: "Test File Output",
		},
		Slides: []spec.SlideSpec{
			{
				ID:    "slide1",
				Type:  spec.SlideTypeContent,
				Title: "Content Slide",
			},
		},
	}

	tmpDir := t.TempDir()
	outputPath := filepath.Join(tmpDir, "output.html")

	renderer := NewRenderer()
	ctx := context.Background()

	if err := renderer.RenderToFile(ctx, pres, outputPath); err != nil {
		t.Fatalf("RenderToFile failed: %v", err)
	}

	// Verify file exists and has content
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	if len(content) == 0 {
		t.Error("expected non-empty output file")
	}

	if !strings.Contains(string(content), "Test File Output") {
		t.Error("expected title in output")
	}
}

func TestRenderer_RenderToDir(t *testing.T) {
	pres := spec.PresentationSpec{
		Version: "1.0",
		Metadata: spec.Metadata{
			Title: "Test Dir Output",
		},
		Slides: []spec.SlideSpec{
			{
				ID:    "slide1",
				Type:  spec.SlideTypeTitle,
				Title: "Title Slide",
			},
		},
	}

	tmpDir := t.TempDir()
	outputDir := filepath.Join(tmpDir, "presentation")

	renderer := NewRenderer()
	ctx := context.Background()

	if err := renderer.RenderToDir(ctx, pres, outputDir); err != nil {
		t.Fatalf("RenderToDir failed: %v", err)
	}

	// Verify index.html exists
	indexPath := filepath.Join(outputDir, "index.html")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Error("expected index.html to exist")
	}

	// Verify assets directory exists
	assetsDir := filepath.Join(outputDir, "assets")
	if _, err := os.Stat(assetsDir); os.IsNotExist(err) {
		t.Error("expected assets directory to exist")
	}

	// Verify CSS file
	cssPath := filepath.Join(assetsDir, "style.css")
	cssContent, err := os.ReadFile(cssPath)
	if err != nil {
		t.Errorf("failed to read CSS file: %v", err)
	}
	if !strings.Contains(string(cssContent), ".presentation") {
		t.Error("expected CSS to contain .presentation selector")
	}

	// Verify JS file
	jsPath := filepath.Join(assetsDir, "navigation.js")
	jsContent, err := os.ReadFile(jsPath)
	if err != nil {
		t.Errorf("failed to read JS file: %v", err)
	}
	if !strings.Contains(string(jsContent), "navigateSlide") {
		t.Error("expected JS to contain navigateSlide function")
	}
}

func TestLoadSpecWithData(t *testing.T) {
	tmpDir := t.TempDir()

	// Create data file
	dataPath := filepath.Join(tmpDir, "metrics.json")
	if err := os.WriteFile(dataPath, []byte(`{"revenue": 1000000}`), 0600); err != nil {
		t.Fatal(err)
	}

	// Create spec file with data source
	specJSON := `{
		"version": "1.0",
		"metadata": {"title": "Test"},
		"slides": [],
		"dataSources": [
			{
				"id": "metrics",
				"type": "json",
				"path": "metrics.json"
			}
		]
	}`
	specPath := filepath.Join(tmpDir, "spec.json")
	if err := os.WriteFile(specPath, []byte(specJSON), 0600); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	pres, store, err := LoadSpecWithData(ctx, specPath)
	if err != nil {
		t.Fatalf("LoadSpecWithData failed: %v", err)
	}

	if pres.Version != "1.0" {
		t.Errorf("expected version '1.0', got %q", pres.Version)
	}

	loaded, ok := store.Get("metrics")
	if !ok {
		t.Fatal("expected metrics data in store")
	}

	data, ok := loaded.Data.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any, got %T", loaded.Data)
	}

	if data["revenue"] != float64(1000000) {
		t.Errorf("expected revenue 1000000, got %v", data["revenue"])
	}
}

func TestLoadSpecWithData_NoDataSources(t *testing.T) {
	tmpDir := t.TempDir()

	specJSON := `{
		"version": "1.0",
		"metadata": {"title": "Test"},
		"slides": []
	}`
	specPath := filepath.Join(tmpDir, "spec.json")
	if err := os.WriteFile(specPath, []byte(specJSON), 0600); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	pres, store, err := LoadSpecWithData(ctx, specPath)
	if err != nil {
		t.Fatalf("LoadSpecWithData failed: %v", err)
	}

	if pres.Metadata.Title != "Test" {
		t.Errorf("expected title 'Test', got %q", pres.Metadata.Title)
	}

	if len(store.All()) != 0 {
		t.Errorf("expected empty store, got %d items", len(store.All()))
	}
}
