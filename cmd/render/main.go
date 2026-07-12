// Command render converts PresentationSpec JSON files to HTML or PDF.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/grokify/presentation-spec/schema"
	"github.com/grokify/presentation-spec/templ"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var (
	outputDir    string
	outputFile   string
	format       string
	validate     bool
	validateOnly bool
	pageSize     string
	landscape    bool
	watch        bool
	speakerNotes bool
)

var rootCmd = &cobra.Command{
	Use:   "render [spec.json]",
	Short: "Render PresentationSpec JSON to HTML or PDF",
	Long: `render converts a PresentationSpec JSON file into a static HTML presentation or PDF.

Output formats:
  - html (default): Static HTML with CSS/JS assets
  - pdf: PDF document via headless Chrome

Examples:
  # Render to HTML directory
  render presentation.json -o output/

  # Render to PDF
  render presentation.json --format pdf -o presentation.pdf

  # Watch for changes and auto-rebuild
  render presentation.json -o output/ --watch

  # Validate only (no output)
  render presentation.json --validate-only

  # Validate before rendering
  render presentation.json --validate -o output/`,
	Args: cobra.ExactArgs(1),
	RunE: runRender,
}

func init() {
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "",
		"Output directory (HTML) or file (PDF)")
	rootCmd.Flags().StringVarP(&outputFile, "file", "f", "",
		"Output HTML file (single file, no assets)")
	rootCmd.Flags().StringVar(&format, "format", "html",
		"Output format: html, pdf")
	rootCmd.Flags().BoolVar(&validate, "validate", false,
		"Validate spec against schema before rendering")
	rootCmd.Flags().BoolVar(&validateOnly, "validate-only", false,
		"Validate spec only, don't render")
	rootCmd.Flags().StringVar(&pageSize, "page-size", "Letter",
		"PDF page size: Letter, A4, Legal, Tabloid")
	rootCmd.Flags().BoolVar(&landscape, "landscape", true,
		"PDF landscape orientation")
	rootCmd.Flags().BoolVarP(&watch, "watch", "w", false,
		"Watch for changes and auto-rebuild (HTML only)")
	rootCmd.Flags().BoolVar(&speakerNotes, "speaker-notes", false,
		"Generate speaker notes view (HTML only)")
}

func runRender(cmd *cobra.Command, args []string) error {
	specPath := args[0]

	// Validate if requested
	if validate || validateOnly {
		result, err := schema.ValidateFile(specPath)
		if err != nil {
			return fmt.Errorf("validation error: %w", err)
		}

		if !result.Valid {
			fmt.Fprintln(os.Stderr, "Validation failed:")
			for _, e := range result.Errors {
				fmt.Fprintf(os.Stderr, "  - %s\n", e)
			}
			return fmt.Errorf("spec is invalid")
		}

		fmt.Println("✓ Spec is valid")

		if validateOnly {
			return nil
		}
	}

	// Create context with signal handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle Ctrl+C gracefully
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nShutting down...")
		cancel()
	}()

	renderer := templ.NewRenderer()

	// Watch mode (HTML only)
	if watch {
		if format != "html" {
			return fmt.Errorf("watch mode only supports HTML format")
		}

		output := outputDir
		if output == "" {
			base := filepath.Base(specPath)
			name := strings.TrimSuffix(base, filepath.Ext(base))
			output = filepath.Join(filepath.Dir(specPath), name+"-output")
		}

		opts := templ.DefaultWatchOptions()
		return renderer.WatchAndRender(ctx, specPath, output, opts)
	}

	// Load the spec
	pres, err := templ.LoadSpec(specPath)
	if err != nil {
		return fmt.Errorf("failed to load spec: %w", err)
	}

	// Handle format
	switch strings.ToLower(format) {
	case "pdf":
		return renderPDF(ctx, renderer, pres, specPath)
	case "html":
		return renderHTML(ctx, renderer, pres, specPath)
	default:
		return fmt.Errorf("unknown format: %s (use html or pdf)", format)
	}
}

func renderPDF(ctx context.Context, renderer *templ.Renderer, pres templ.SpecType, specPath string) error {
	// Determine output path
	output := outputDir
	if output == "" {
		base := filepath.Base(specPath)
		name := strings.TrimSuffix(base, filepath.Ext(base))
		output = filepath.Join(filepath.Dir(specPath), name+".pdf")
	}

	// Ensure .pdf extension
	if !strings.HasSuffix(strings.ToLower(output), ".pdf") {
		output = output + ".pdf"
	}

	opts := templ.DefaultPDFOptions()
	opts.PageSize = pageSize
	opts.Landscape = landscape

	if err := renderer.RenderToPDF(ctx, pres, output, opts); err != nil {
		return fmt.Errorf("failed to render PDF: %w", err)
	}

	fmt.Printf("Rendered PDF to %s\n", output)
	return nil
}

func renderHTML(ctx context.Context, renderer *templ.Renderer, pres templ.SpecType, specPath string) error {
	// Single file output
	if outputFile != "" {
		if err := renderer.RenderToFile(ctx, pres, outputFile); err != nil {
			return fmt.Errorf("failed to render: %w", err)
		}
		fmt.Printf("Rendered to %s\n", outputFile)
		fmt.Println("Note: Assets (CSS/JS) not included. For full output, use -o flag.")
		return nil
	}

	// Directory output
	output := outputDir
	if output == "" {
		base := filepath.Base(specPath)
		name := strings.TrimSuffix(base, filepath.Ext(base))
		output = filepath.Join(filepath.Dir(specPath), name+"-output")
	}

	opts := templ.RenderOptions{
		IncludeSpeakerNotes: speakerNotes,
	}

	if err := renderer.RenderToDirWithOptions(ctx, pres, output, opts); err != nil {
		return fmt.Errorf("failed to render: %w", err)
	}

	fmt.Printf("Rendered to %s/\n", output)
	fmt.Printf("  - index.html\n")
	fmt.Printf("  - assets/style.css\n")
	fmt.Printf("  - assets/navigation.js\n")
	if speakerNotes {
		fmt.Printf("  - speaker.html\n")
		fmt.Printf("  - assets/speaker.css\n")
		fmt.Printf("  - assets/speaker.js\n")
	}
	fmt.Printf("\nOpen %s/index.html in a browser to view.\n", output)
	if speakerNotes {
		fmt.Printf("Open %s/speaker.html for speaker notes view.\n", output)
	}

	return nil
}
