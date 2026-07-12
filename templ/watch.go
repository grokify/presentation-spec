package templ

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/grokify/presentation-spec/spec"
)

// WatchOptions configures watch mode behavior.
type WatchOptions struct {
	// Interval is the polling interval for file changes.
	Interval time.Duration

	// OnChange is called when a change is detected and rebuild completes.
	OnChange func(err error)

	// Verbose enables verbose logging.
	Verbose bool
}

// DefaultWatchOptions returns default watch options.
func DefaultWatchOptions() WatchOptions {
	return WatchOptions{
		Interval: 500 * time.Millisecond,
		Verbose:  true,
	}
}

// WatchAndRender watches a spec file and re-renders on changes.
func (r *Renderer) WatchAndRender(ctx context.Context, specPath, outputDir string, opts WatchOptions) error {
	absPath, err := filepath.Abs(specPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	var lastModTime time.Time

	// Initial render
	if err := r.renderOnce(ctx, absPath, outputDir, opts.Verbose); err != nil {
		return err
	}

	// Get initial mod time
	info, err := os.Stat(absPath)
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}
	lastModTime = info.ModTime()

	if opts.Verbose {
		fmt.Printf("Watching %s for changes...\n", specPath)
		fmt.Println("Press Ctrl+C to stop.")
	}

	ticker := time.NewTicker(opts.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			info, err := os.Stat(absPath)
			if err != nil {
				if opts.Verbose {
					fmt.Printf("Error checking file: %v\n", err)
				}
				continue
			}

			if info.ModTime().After(lastModTime) {
				lastModTime = info.ModTime()

				if opts.Verbose {
					fmt.Printf("\n[%s] Change detected, rebuilding...\n",
						time.Now().Format("15:04:05"))
				}

				err := r.renderOnce(ctx, absPath, outputDir, opts.Verbose)
				if opts.OnChange != nil {
					opts.OnChange(err)
				}
			}
		}
	}
}

func (r *Renderer) renderOnce(ctx context.Context, specPath, outputDir string, verbose bool) error {
	pres, err := LoadSpec(specPath)
	if err != nil {
		if verbose {
			fmt.Printf("Error loading spec: %v\n", err)
		}
		return err
	}

	if err := r.RenderToDir(ctx, pres, outputDir); err != nil {
		if verbose {
			fmt.Printf("Error rendering: %v\n", err)
		}
		return err
	}

	if verbose {
		fmt.Printf("Rendered to %s/\n", outputDir)
	}

	return nil
}

// LoadSpecWithValidation loads and validates a spec file.
func LoadSpecWithValidation(path string) (spec.PresentationSpec, error) {
	pres, err := LoadSpec(path)
	if err != nil {
		return pres, err
	}

	// Basic validation
	if pres.Version == "" {
		return pres, fmt.Errorf("spec missing version field")
	}
	if pres.Metadata.Title == "" {
		return pres, fmt.Errorf("spec missing metadata.title field")
	}

	return pres, nil
}
