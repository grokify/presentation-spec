package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/grokify/presentation-spec/templ"
	"github.com/spf13/cobra"
)

var (
	servePort      int
	serveHost      string
	serveNoWatch   bool
	serveNoBrowser bool
)

var serveCmd = &cobra.Command{
	Use:   "serve [spec.json]",
	Short: "Start a dev server with live reload",
	Long: `serve starts a development server that serves the rendered presentation
and automatically rebuilds when the spec file changes.

Examples:
  # Start dev server
  render serve presentation.json

  # Specify port
  render serve presentation.json --port 3000

  # Disable auto-rebuild
  render serve presentation.json --no-watch`,
	Args: cobra.ExactArgs(1),
	RunE: runServe,
}

func init() {
	serveCmd.Flags().IntVarP(&servePort, "port", "p", 8080,
		"Port to listen on")
	serveCmd.Flags().StringVar(&serveHost, "host", "localhost",
		"Host to bind to")
	serveCmd.Flags().BoolVar(&serveNoWatch, "no-watch", false,
		"Disable auto-rebuild on file changes")
	serveCmd.Flags().BoolVar(&serveNoBrowser, "no-browser", false,
		"Don't open browser on start")
	serveCmd.Flags().StringVarP(&outputDir, "output", "o", "",
		"Output directory for rendered files")

	rootCmd.AddCommand(serveCmd)
}

func runServe(cmd *cobra.Command, args []string) error {
	specPath := args[0]

	// Determine output directory
	output := outputDir
	if output == "" {
		base := filepath.Base(specPath)
		name := strings.TrimSuffix(base, filepath.Ext(base))
		output = filepath.Join(filepath.Dir(specPath), name+"-output")
	}

	// Create context with signal handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nShutting down server...")
		cancel()
	}()

	opts := templ.DefaultServerOptions()
	opts.Port = servePort
	opts.Host = serveHost
	opts.Watch = !serveNoWatch
	opts.OpenBrowser = !serveNoBrowser

	server := templ.NewServer(specPath, output, opts)
	return server.Start(ctx)
}
