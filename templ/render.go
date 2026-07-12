// Package templ provides HTML rendering for PresentationSpec using templ templates.
package templ

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/grokify/presentation-spec/spec"
)

// SpecType is an alias for spec.PresentationSpec for external use.
type SpecType = spec.PresentationSpec

// Renderer renders PresentationSpec to HTML.
type Renderer struct {
	// StaticDir is the directory containing static assets (CSS, JS).
	// If empty, embedded defaults are used.
	StaticDir string

	// IncludeMermaid enables Mermaid diagram support.
	IncludeMermaid bool
}

// NewRenderer creates a new Renderer with default settings.
func NewRenderer() *Renderer {
	return &Renderer{
		IncludeMermaid: true,
	}
}

// RenderToWriter renders a presentation to the provided writer.
func (r *Renderer) RenderToWriter(ctx context.Context, pres spec.PresentationSpec, w io.Writer) error {
	component := Presentation(pres)
	return component.Render(ctx, w)
}

// RenderToFile renders a presentation to an HTML file.
func (r *Renderer) RenderToFile(ctx context.Context, pres spec.PresentationSpec, outputPath string) error {
	// Ensure output directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create output file
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	if err := r.RenderToWriter(ctx, pres, f); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

// RenderToDir renders a presentation to a directory with assets.
func (r *Renderer) RenderToDir(ctx context.Context, pres spec.PresentationSpec, outputDir string) error {
	return r.RenderToDirWithOptions(ctx, pres, outputDir, RenderOptions{})
}

// RenderOptions configures rendering behavior.
type RenderOptions struct {
	// IncludeSpeakerNotes generates a speaker notes view.
	IncludeSpeakerNotes bool
}

// RenderToDirWithOptions renders a presentation to a directory with options.
func (r *Renderer) RenderToDirWithOptions(ctx context.Context, pres spec.PresentationSpec, outputDir string, opts RenderOptions) error {
	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create assets directory
	assetsDir := filepath.Join(outputDir, "assets")
	if err := os.MkdirAll(assetsDir, 0755); err != nil {
		return fmt.Errorf("failed to create assets directory: %w", err)
	}

	// Write CSS
	cssPath := filepath.Join(assetsDir, "style.css")
	if err := os.WriteFile(cssPath, []byte(DefaultCSS), 0644); err != nil { //nolint:gosec // Assets need to be readable
		return fmt.Errorf("failed to write CSS: %w", err)
	}

	// Write JS
	jsPath := filepath.Join(assetsDir, "navigation.js")
	if err := os.WriteFile(jsPath, []byte(DefaultJS), 0644); err != nil { //nolint:gosec // Assets need to be readable
		return fmt.Errorf("failed to write JS: %w", err)
	}

	// Write speaker assets if needed
	if opts.IncludeSpeakerNotes {
		speakerCSSPath := filepath.Join(assetsDir, "speaker.css")
		if err := os.WriteFile(speakerCSSPath, []byte(SpeakerCSS), 0644); err != nil { //nolint:gosec // Assets need to be readable
			return fmt.Errorf("failed to write speaker CSS: %w", err)
		}

		speakerJSPath := filepath.Join(assetsDir, "speaker.js")
		if err := os.WriteFile(speakerJSPath, []byte(SpeakerJS), 0644); err != nil { //nolint:gosec // Assets need to be readable
			return fmt.Errorf("failed to write speaker JS: %w", err)
		}
	}

	// Render main HTML
	htmlPath := filepath.Join(outputDir, "index.html")
	if err := r.RenderToFile(ctx, pres, htmlPath); err != nil {
		return err
	}

	// Render speaker notes view if requested
	if opts.IncludeSpeakerNotes {
		speakerPath := filepath.Join(outputDir, "speaker.html")
		if err := r.RenderSpeakerView(ctx, pres, speakerPath); err != nil {
			return fmt.Errorf("failed to render speaker view: %w", err)
		}
	}

	return nil
}

// RenderSpeakerView renders the speaker notes view to a file.
func (r *Renderer) RenderSpeakerView(ctx context.Context, pres spec.PresentationSpec, outputPath string) error {
	// Ensure output directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create output file
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	component := SpeakerView(pres)
	if err := component.Render(ctx, f); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

// LoadSpec loads a PresentationSpec from a JSON file.
func LoadSpec(path string) (spec.PresentationSpec, error) {
	var pres spec.PresentationSpec

	data, err := os.ReadFile(path)
	if err != nil {
		return pres, fmt.Errorf("failed to read spec file: %w", err)
	}

	if err := json.Unmarshal(data, &pres); err != nil {
		return pres, fmt.Errorf("failed to parse spec JSON: %w", err)
	}

	return pres, nil
}

// LoadSpecWithData loads a spec and all its data sources.
func LoadSpecWithData(ctx context.Context, path string) (spec.PresentationSpec, *DataStore, error) {
	pres, err := LoadSpec(path)
	if err != nil {
		return pres, nil, err
	}

	if len(pres.DataSources) == 0 {
		return pres, NewDataStore(), nil
	}

	baseDir := filepath.Dir(path)
	loader := NewDataLoader(baseDir)

	store, err := loader.LoadAll(ctx, pres.DataSources)
	if err != nil {
		return pres, store, fmt.Errorf("failed to load data sources: %w", err)
	}

	return pres, store, nil
}

// DefaultCSS is the default stylesheet for rendered presentations.
const DefaultCSS = `/* PresentationSpec Default Styles */
:root {
  --color-primary: #3b82f6;
  --color-secondary: #10b981;
  --color-background: #ffffff;
  --color-surface: #f8fafc;
  --color-text: #1e293b;
  --color-text-muted: #64748b;
  --color-success: #22c55e;
  --color-warning: #f59e0b;
  --color-error: #ef4444;
  --font-family: system-ui, -apple-system, sans-serif;
  --font-family-heading: system-ui, -apple-system, sans-serif;
  --font-family-code: ui-monospace, monospace;
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: var(--font-family);
  color: var(--color-text);
  background: var(--color-background);
  line-height: 1.6;
}

.presentation-container {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}

.presentation {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}

/* Slides */
.slide {
  width: 100%;
  min-height: 100vh;
  padding: 2rem 4rem;
  display: none;
  flex-direction: column;
}

.slide.active {
  display: flex;
}

.slide:first-child {
  display: flex;
}

.slide-type-title {
  justify-content: center;
  align-items: center;
  text-align: center;
}

.slide-type-section {
  justify-content: center;
  align-items: center;
  text-align: center;
  background: var(--color-surface);
}

.slide-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.title-slide-content h1 {
  font-size: 3rem;
  font-weight: 700;
  margin-bottom: 1rem;
}

.slide-header {
  margin-bottom: 2rem;
}

.slide-title {
  font-size: 2rem;
  font-weight: 600;
  font-family: var(--font-family-heading);
}

.slide-key-message {
  color: var(--color-text-muted);
  font-size: 1.1rem;
  margin-top: 0.5rem;
}

.slide-body {
  flex: 1;
  display: grid;
  gap: 1.5rem;
  align-content: start;
}

/* Layouts */
.layout-single { grid-template-columns: 1fr; }
.layout-two_column { grid-template-columns: 1fr 1fr; }
.layout-three_column { grid-template-columns: 1fr 1fr 1fr; }
.layout-four_column { grid-template-columns: 1fr 1fr 1fr 1fr; }
.layout-kpi_row_content_grid {
  grid-template-columns: repeat(4, 1fr);
  grid-template-rows: auto 1fr;
}
.layout-metric_dashboard {
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
}

/* Regions */
.region { padding: 1rem; }
.region-kpi { display: flex; gap: 1rem; }
.region-content { grid-column: span 3; }
.region-sidebar { grid-column: span 1; }

/* Widgets */
.widget {
  background: var(--color-surface);
  border-radius: 0.5rem;
  padding: 1rem;
}

/* Metric Card */
.metric-card {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.metric-title {
  font-size: 0.875rem;
  color: var(--color-text-muted);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.metric-value-container {
  display: flex;
  align-items: baseline;
  gap: 0.25rem;
}

.metric-value {
  font-size: 2rem;
  font-weight: 700;
}

.metric-unit {
  font-size: 1rem;
  color: var(--color-text-muted);
}

.metric-trend {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  font-size: 0.875rem;
}

.trend-up { color: var(--color-success); }
.trend-down { color: var(--color-error); }
.trend-flat { color: var(--color-text-muted); }

.metric-target {
  font-size: 0.75rem;
  color: var(--color-text-muted);
}

.sparkline-svg {
  width: 100%;
  height: 30px;
  margin-top: 0.5rem;
}

/* Status colors */
.status-good { border-left: 3px solid var(--color-success); }
.status-warning { border-left: 3px solid var(--color-warning); }
.status-critical { border-left: 3px solid var(--color-error); }
.status-neutral { border-left: 3px solid var(--color-text-muted); }

/* Content Block */
.content-block { line-height: 1.7; }
.content-title { font-size: 1.25rem; margin-bottom: 0.75rem; }
.content-body h2 { font-size: 1.5rem; margin: 1rem 0 0.5rem; }
.content-body h3 { font-size: 1.25rem; margin: 0.75rem 0 0.5rem; }
.content-body h4 { font-size: 1.1rem; margin: 0.5rem 0 0.25rem; }
.content-body ul, .content-body ol { margin: 0.5rem 0; padding-left: 1.5rem; }
.content-body li { margin: 0.25rem 0; }
.content-body p { margin: 0.5rem 0; }
.content-body strong { font-weight: 600; }
.content-body code {
  background: var(--color-surface);
  padding: 0.125rem 0.375rem;
  border-radius: 0.25rem;
  font-family: var(--font-family-code);
  font-size: 0.875em;
}

/* Lists */
.list-title { font-size: 1.125rem; margin-bottom: 0.75rem; }

.risk-items, .decision-items, .checklist-items {
  list-style: none;
  padding: 0;
}

.risk-item, .decision-item, .checklist-item {
  padding: 0.75rem;
  border-radius: 0.375rem;
  margin-bottom: 0.5rem;
  background: var(--color-background);
}

.risk-header { display: flex; align-items: flex-start; gap: 0.5rem; }

.severity-badge {
  font-size: 0.7rem;
  padding: 0.125rem 0.5rem;
  border-radius: 1rem;
  text-transform: uppercase;
  font-weight: 600;
}

.severity-low { background: #dcfce7; color: #166534; }
.severity-medium { background: #fef3c7; color: #92400e; }
.severity-high { background: #fed7aa; color: #c2410c; }
.severity-critical { background: #fee2e2; color: #b91c1c; }

.risk-mitigation, .risk-owner {
  font-size: 0.875rem;
  margin-top: 0.5rem;
  color: var(--color-text-muted);
}

.decision-meta {
  display: flex;
  gap: 1rem;
  margin-top: 0.5rem;
  font-size: 0.875rem;
  color: var(--color-text-muted);
}

.checklist-item {
  display: flex;
  align-items: flex-start;
  gap: 0.5rem;
}

.check-indicator {
  font-size: 1.25rem;
  line-height: 1;
}

.checklist-item.checked { color: var(--color-success); }
.checklist-item.unchecked { color: var(--color-text-muted); }
.checklist-item.checked .check-text { text-decoration: line-through; }

/* Charts */
.chart-container { min-height: 200px; }
.chart-title { font-size: 1.125rem; margin-bottom: 0.75rem; }
.chart-svg { width: 100%; height: 200px; }
.chart-fallback {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 150px;
  color: var(--color-text-muted);
}

/* Tables */
.data-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.875rem;
}

.data-table th, .data-table td {
  padding: 0.75rem;
  text-align: left;
  border-bottom: 1px solid var(--color-surface);
}

.data-table th {
  font-weight: 600;
  background: var(--color-surface);
}

.align-center { text-align: center; }
.align-right { text-align: right; }

/* Callout */
.callout {
  display: flex;
  gap: 1rem;
  padding: 1rem;
  border-radius: 0.5rem;
}

.callout-info { background: #eff6ff; border-left: 3px solid #3b82f6; }
.callout-warning { background: #fffbeb; border-left: 3px solid #f59e0b; }
.callout-error { background: #fef2f2; border-left: 3px solid #ef4444; }
.callout-success { background: #f0fdf4; border-left: 3px solid #22c55e; }
.callout-tip { background: #faf5ff; border-left: 3px solid #a855f7; }

.callout-icon { font-size: 1.25rem; }
.callout-title { font-weight: 600; margin-bottom: 0.25rem; }

/* Quote */
.quote-container {
  padding: 1.5rem;
  border-left: 4px solid var(--color-primary);
  background: var(--color-surface);
}

.quote-text {
  font-size: 1.25rem;
  font-style: italic;
  margin-bottom: 0.75rem;
}

.quote-attribution {
  color: var(--color-text-muted);
}

/* Image */
.image-container { text-align: center; }
.image { max-width: 100%; height: auto; border-radius: 0.5rem; }
.image-caption { font-size: 0.875rem; color: var(--color-text-muted); margin-top: 0.5rem; }

/* Code */
.code-container { font-family: var(--font-family-code); }
.code-header {
  display: flex;
  justify-content: space-between;
  padding: 0.5rem 1rem;
  background: #1e293b;
  color: #e2e8f0;
  border-radius: 0.5rem 0.5rem 0 0;
  font-size: 0.75rem;
}
.code-block {
  background: #0f172a;
  color: #e2e8f0;
  padding: 1rem;
  border-radius: 0 0 0.5rem 0.5rem;
  overflow-x: auto;
  font-size: 0.875rem;
  line-height: 1.5;
}
.code-block code { background: none; padding: 0; }

/* Diagrams */
.diagram-container { text-align: center; }
.diagram-source {
  text-align: left;
  background: var(--color-surface);
  padding: 1rem;
  border-radius: 0.5rem;
  font-family: var(--font-family-code);
  font-size: 0.875rem;
  white-space: pre-wrap;
}

/* Navigation */
.slide-nav {
  position: fixed;
  bottom: 1rem;
  right: 1rem;
  display: flex;
  align-items: center;
  gap: 0.75rem;
  background: rgba(255, 255, 255, 0.9);
  padding: 0.5rem 1rem;
  border-radius: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  z-index: 100;
}

.nav-prev, .nav-next {
  background: none;
  border: none;
  font-size: 1.25rem;
  cursor: pointer;
  padding: 0.25rem 0.5rem;
  color: var(--color-text);
}

.nav-prev:hover, .nav-next:hover {
  color: var(--color-primary);
}

.slide-counter {
  font-size: 0.875rem;
  color: var(--color-text-muted);
}

/* Aspect ratios */
.aspect-16x9 .slide { aspect-ratio: 16/9; min-height: auto; }
.aspect-4x3 .slide { aspect-ratio: 4/3; min-height: auto; }
.aspect-1x1 .slide { aspect-ratio: 1/1; min-height: auto; }

/* Density */
.density-compact .slide { padding: 1rem 2rem; }
.density-spacious .slide { padding: 3rem 6rem; }

/* Print */
@media print {
  .slide-nav { display: none; }
  .slide {
    display: flex !important;
    page-break-after: always;
    min-height: auto;
  }
}
`

// SpeakerCSS is the stylesheet for the speaker notes view.
const SpeakerCSS = `/* Speaker Notes View */
.speaker-view {
  font-family: var(--font-family, system-ui, -apple-system, sans-serif);
  background: #1a1a2e;
  color: #eee;
  min-height: 100vh;
  margin: 0;
  padding: 0;
}

.speaker-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  padding: 1rem;
  box-sizing: border-box;
}

.speaker-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem 1rem;
  background: #16213e;
  border-radius: 0.5rem;
  margin-bottom: 1rem;
}

.speaker-header h1 {
  font-size: 1.25rem;
  margin: 0;
}

.speaker-timer {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.timer-elapsed {
  font-size: 2rem;
  font-family: monospace;
  font-weight: bold;
}

.timer-btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 0.25rem;
  background: #0f3460;
  color: #eee;
  cursor: pointer;
  font-size: 0.875rem;
}

.timer-btn:hover {
  background: #1a4d7e;
}

.speaker-main {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 1rem;
  flex: 1;
  min-height: 0;
}

.speaker-slides {
  display: grid;
  grid-template-columns: 3fr 2fr;
  gap: 1rem;
}

.current-slide-container,
.next-slide-container,
.speaker-notes-pane {
  background: #16213e;
  border-radius: 0.5rem;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.pane-title {
  font-size: 0.875rem;
  color: #888;
  margin: 0 0 0.5rem 0;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.current-slide,
.next-slide {
  flex: 1;
  position: relative;
  min-height: 0;
}

.slide-preview {
  display: none;
  flex-direction: column;
  height: 100%;
}

.slide-preview.active {
  display: flex;
}

.slide-preview-header {
  padding: 0.5rem;
  background: #0f3460;
  border-radius: 0.25rem;
  margin-bottom: 0.5rem;
}

.slide-num {
  font-size: 0.75rem;
  color: #888;
}

.slide-title {
  font-weight: 600;
  margin-left: 1rem;
}

.slide-frame {
  flex: 1;
  border: none;
  border-radius: 0.25rem;
  background: #fff;
  width: 100%;
}

.end-slide {
  display: flex;
  align-items: center;
  justify-content: center;
}

.end-slide-content {
  font-size: 1.5rem;
  color: #888;
}

.notes-content {
  flex: 1;
  overflow-y: auto;
  line-height: 1.6;
}

.notes-item {
  display: none;
}

.notes-item.active {
  display: block;
}

.notes-item ul, .notes-item ol {
  margin: 0.5rem 0;
  padding-left: 1.5rem;
}

.no-notes {
  color: #666;
  font-style: italic;
}

.speaker-controls {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 2rem;
  padding: 1rem;
  background: #16213e;
  border-radius: 0.5rem;
  margin-top: 1rem;
}

.control-btn {
  padding: 0.75rem 2rem;
  border: none;
  border-radius: 0.25rem;
  background: #e94560;
  color: #fff;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
}

.control-btn:hover {
  background: #ff6b6b;
}

.slide-counter {
  font-size: 1.25rem;
}
`

// SpeakerJS is the JavaScript for the speaker notes view.
const SpeakerJS = `// Speaker Notes View
(function() {
  let currentSlide = 0;
  let timerRunning = false;
  let timerStart = null;
  let timerInterval = null;
  let totalSlides = 0;

  function init() {
    const previews = document.querySelectorAll('.current-slide .slide-preview');
    totalSlides = previews.length;
    showSlide(0);
  }

  function showSlide(index) {
    if (index < 0) index = 0;
    if (index >= totalSlides) index = totalSlides - 1;

    currentSlide = index;

    // Update current slide
    document.querySelectorAll('.current-slide .slide-preview').forEach((el, i) => {
      el.classList.toggle('active', i === index);
    });

    // Update next slide
    document.querySelectorAll('.next-slide .slide-preview').forEach((el) => {
      const slideNum = parseInt(el.getAttribute('data-slide'));
      el.classList.toggle('active', slideNum === index + 1 || (slideNum === 'end' && index === totalSlides - 1));
    });

    // Update notes
    document.querySelectorAll('.notes-item').forEach((el, i) => {
      el.classList.toggle('active', i === index);
    });

    // Update counter
    document.getElementById('current-num').textContent = index + 1;

    // Sync main window if open
    if (window.mainWindow && !window.mainWindow.closed) {
      window.mainWindow.location.hash = document.querySelectorAll('.current-slide .slide-preview')[index]
        ?.querySelector('.slide-frame')?.src.split('#')[1] || '';
    }
  }

  window.navigateSpeaker = function(delta) {
    showSlide(currentSlide + delta);
  };

  window.toggleTimer = function() {
    const btn = document.getElementById('timer-toggle');
    if (timerRunning) {
      clearInterval(timerInterval);
      timerRunning = false;
      btn.textContent = 'Resume';
    } else {
      if (!timerStart) {
        timerStart = Date.now();
      }
      timerInterval = setInterval(updateTimer, 1000);
      timerRunning = true;
      btn.textContent = 'Pause';
    }
  };

  window.resetTimer = function() {
    clearInterval(timerInterval);
    timerRunning = false;
    timerStart = null;
    document.getElementById('elapsed').textContent = '00:00:00';
    document.getElementById('timer-toggle').textContent = 'Start';
  };

  function updateTimer() {
    const elapsed = Date.now() - timerStart;
    const hours = Math.floor(elapsed / 3600000);
    const minutes = Math.floor((elapsed % 3600000) / 60000);
    const seconds = Math.floor((elapsed % 60000) / 1000);
    document.getElementById('elapsed').textContent =
      String(hours).padStart(2, '0') + ':' +
      String(minutes).padStart(2, '0') + ':' +
      String(seconds).padStart(2, '0');
  }

  // Keyboard navigation
  document.addEventListener('keydown', function(e) {
    switch(e.key) {
      case 'ArrowRight':
      case 'ArrowDown':
      case ' ':
      case 'PageDown':
        e.preventDefault();
        navigateSpeaker(1);
        break;
      case 'ArrowLeft':
      case 'ArrowUp':
      case 'PageUp':
        e.preventDefault();
        navigateSpeaker(-1);
        break;
      case 'Home':
        e.preventDefault();
        showSlide(0);
        break;
      case 'End':
        e.preventDefault();
        showSlide(totalSlides - 1);
        break;
    }
  });

  // Initialize on load
  document.addEventListener('DOMContentLoaded', init);
})();
`

// DefaultJS is the default JavaScript for slide navigation.
const DefaultJS = `// PresentationSpec Navigation
(function() {
  let currentSlide = 0;
  const slides = document.querySelectorAll('.slide');
  const counter = document.getElementById('current-slide');

  function showSlide(index) {
    if (index < 0) index = 0;
    if (index >= slides.length) index = slides.length - 1;

    slides.forEach((slide, i) => {
      slide.classList.toggle('active', i === index);
    });

    currentSlide = index;
    if (counter) counter.textContent = currentSlide + 1;

    // Update URL hash
    window.location.hash = slides[index].id || ('slide-' + (index + 1));
  }

  window.navigateSlide = function(delta) {
    showSlide(currentSlide + delta);
  };

  // Keyboard navigation
  document.addEventListener('keydown', function(e) {
    switch(e.key) {
      case 'ArrowRight':
      case 'ArrowDown':
      case ' ':
      case 'PageDown':
        e.preventDefault();
        navigateSlide(1);
        break;
      case 'ArrowLeft':
      case 'ArrowUp':
      case 'PageUp':
        e.preventDefault();
        navigateSlide(-1);
        break;
      case 'Home':
        e.preventDefault();
        showSlide(0);
        break;
      case 'End':
        e.preventDefault();
        showSlide(slides.length - 1);
        break;
    }
  });

  // Handle hash navigation
  function handleHash() {
    const hash = window.location.hash.slice(1);
    if (hash) {
      const target = document.getElementById(hash);
      if (target && target.classList.contains('slide')) {
        const index = Array.from(slides).indexOf(target);
        if (index >= 0) showSlide(index);
      }
    }
  }

  window.addEventListener('hashchange', handleHash);

  // Initialize
  if (window.location.hash) {
    handleHash();
  } else {
    showSlide(0);
  }
})();

// Initialize Mermaid diagrams
if (typeof mermaid !== 'undefined') {
  mermaid.initialize({
    startOnLoad: true,
    theme: 'default',
    securityLevel: 'loose',
  });
}

// Initialize Prism.js syntax highlighting
if (typeof Prism !== 'undefined') {
  Prism.highlightAll();
}

// Initialize Chart.js charts
if (typeof Chart !== 'undefined') {
  document.querySelectorAll('[data-chart-type]').forEach(function(el) {
    try {
      const chartType = el.getAttribute('data-chart-type');
      const chartData = JSON.parse(el.getAttribute('data-chart-data') || '{}');
      const chartOptions = JSON.parse(el.getAttribute('data-chart-options') || '{}');

      if (!chartData.labels || !chartData.series) return;

      // Find or create canvas
      let canvas = el.querySelector('canvas');
      if (!canvas) {
        // Clear existing content (SVG fallback)
        el.innerHTML = '';
        canvas = document.createElement('canvas');
        el.appendChild(canvas);
      }

      // Convert our data format to Chart.js format
      const datasets = chartData.series.map(function(s, i) {
        const colors = ['#3b82f6', '#10b981', '#f59e0b', '#ef4444', '#8b5cf6'];
        return {
          label: s.name || ('Series ' + (i + 1)),
          data: s.values,
          backgroundColor: s.color || colors[i % colors.length],
          borderColor: s.color || colors[i % colors.length],
          fill: chartType === 'area',
        };
      });

      new Chart(canvas, {
        type: chartType === 'area' ? 'line' : chartType,
        data: {
          labels: chartData.labels,
          datasets: datasets,
        },
        options: {
          responsive: true,
          maintainAspectRatio: false,
          plugins: {
            legend: { display: chartOptions.showLegend || false },
          },
          scales: chartType !== 'pie' && chartType !== 'donut' ? {
            x: {
              display: true,
              title: { display: !!chartOptions.xAxisLabel, text: chartOptions.xAxisLabel }
            },
            y: {
              display: true,
              title: { display: !!chartOptions.yAxisLabel, text: chartOptions.yAxisLabel }
            },
          } : undefined,
        },
      });
    } catch (e) {
      console.error('Failed to initialize chart:', e);
    }
  });
}
`
