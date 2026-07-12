package templ

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/grokify/presentation-spec/spec"
)

// PDFOptions configures PDF export.
type PDFOptions struct {
	// PageSize is the page size (e.g., "Letter", "A4").
	PageSize string

	// Landscape sets landscape orientation.
	Landscape bool

	// Scale is the scale factor (default 1.0).
	Scale float64

	// PrintBackground includes background graphics.
	PrintBackground bool

	// Timeout is the maximum time to wait for rendering.
	Timeout time.Duration
}

// DefaultPDFOptions returns default PDF export options.
func DefaultPDFOptions() PDFOptions {
	return PDFOptions{
		PageSize:        "Letter",
		Landscape:       true,
		Scale:           1.0,
		PrintBackground: true,
		Timeout:         60 * time.Second,
	}
}

// RenderToPDF renders a presentation to a PDF file.
func (r *Renderer) RenderToPDF(ctx context.Context, pres spec.PresentationSpec, outputPath string, opts PDFOptions) error {
	// Create a temporary directory for the HTML
	tmpDir, err := os.MkdirTemp("", "presentation-pdf-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Render HTML to temp directory
	if err := r.RenderToDir(ctx, pres, tmpDir); err != nil {
		return fmt.Errorf("failed to render HTML: %w", err)
	}

	// Convert HTML to PDF using chromedp
	htmlPath := filepath.Join(tmpDir, "index.html")
	return HTMLToPDF(ctx, htmlPath, outputPath, opts)
}

// HTMLToPDF converts an HTML file to PDF using headless Chrome.
// Requires Chrome or Chromium to be installed on the system.
func HTMLToPDF(ctx context.Context, htmlPath, outputPath string, opts PDFOptions) error {
	// Ensure output directory exists
	outDir := filepath.Dir(outputPath)
	if outDir != "" && outDir != "." {
		if err := os.MkdirAll(outDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	// Set up chromedp context with timeout
	allocCtx, cancel := chromedp.NewExecAllocator(ctx,
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("disable-dev-shm-usage", true),
			chromedp.Flag("disable-software-rasterizer", true),
		)...,
	)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	if opts.Timeout > 0 {
		var cancelTimeout context.CancelFunc
		taskCtx, cancelTimeout = context.WithTimeout(taskCtx, opts.Timeout)
		defer cancelTimeout()
	}

	// Convert file path to URL
	absPath, err := filepath.Abs(htmlPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	fileURL := "file://" + absPath

	// Get page dimensions based on page size and orientation
	width, height := getPageDimensions(opts.PageSize, opts.Landscape)

	// Generate PDF
	var pdfData []byte
	err = chromedp.Run(taskCtx,
		chromedp.EmulateViewport(int64(width*96/72), int64(height*96/72)), // Convert points to pixels at 96 DPI
		chromedp.Navigate(fileURL),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfData, _, err = page.PrintToPDF().
				WithPaperWidth(width / 72).   // Convert points to inches
				WithPaperHeight(height / 72). // Convert points to inches
				WithScale(opts.Scale).
				WithPrintBackground(opts.PrintBackground).
				WithLandscape(opts.Landscape).
				WithMarginTop(0.5).
				WithMarginBottom(0.5).
				WithMarginLeft(0.5).
				WithMarginRight(0.5).
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %w", err)
	}

	// Write PDF to file
	if err := os.WriteFile(outputPath, pdfData, 0644); err != nil { //nolint:gosec // PDF output needs to be readable
		return fmt.Errorf("failed to write PDF file: %w", err)
	}

	return nil
}

// getPageDimensions returns page dimensions in points (1/72 inch).
func getPageDimensions(pageSize string, landscape bool) (width, height float64) {
	// Standard page sizes in points
	sizes := map[string][2]float64{
		"Letter":  {612, 792},  // 8.5 x 11 inches
		"A4":      {595, 842},  // 210 x 297 mm
		"Legal":   {612, 1008}, // 8.5 x 14 inches
		"Tabloid": {792, 1224}, // 11 x 17 inches
	}

	dims, ok := sizes[pageSize]
	if !ok {
		dims = sizes["Letter"]
	}

	if landscape {
		return dims[1], dims[0]
	}
	return dims[0], dims[1]
}
