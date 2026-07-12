package templ

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// ServerOptions configures the dev server.
type ServerOptions struct {
	// Port is the HTTP port to listen on.
	Port int

	// Host is the host to bind to.
	Host string

	// Watch enables watch mode for auto-rebuild.
	Watch bool

	// WatchInterval is the polling interval for file changes.
	WatchInterval time.Duration

	// OpenBrowser opens the browser on start.
	OpenBrowser bool

	// Verbose enables verbose logging.
	Verbose bool
}

// DefaultServerOptions returns default server options.
func DefaultServerOptions() ServerOptions {
	return ServerOptions{
		Port:          8080,
		Host:          "localhost",
		Watch:         true,
		WatchInterval: 500 * time.Millisecond,
		OpenBrowser:   true,
		Verbose:       true,
	}
}

// Server is a development server with live reload.
type Server struct {
	renderer  *Renderer
	opts      ServerOptions
	specPath  string
	outputDir string

	mu          sync.RWMutex
	lastModTime time.Time
	buildError  error
}

// NewServer creates a new dev server.
func NewServer(specPath, outputDir string, opts ServerOptions) *Server {
	return &Server{
		renderer:  NewRenderer(),
		opts:      opts,
		specPath:  specPath,
		outputDir: outputDir,
	}
}

// Start starts the dev server.
func (s *Server) Start(ctx context.Context) error {
	// Get absolute paths
	absSpec, err := filepath.Abs(s.specPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute spec path: %w", err)
	}
	s.specPath = absSpec

	absOutput, err := filepath.Abs(s.outputDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute output path: %w", err)
	}
	s.outputDir = absOutput

	// Initial build
	if err := s.rebuild(ctx); err != nil {
		return fmt.Errorf("initial build failed: %w", err)
	}

	// Start file watcher
	if s.opts.Watch {
		go s.watchFiles(ctx)
	}

	// Set up HTTP handler
	mux := http.NewServeMux()

	// Serve static files from output directory
	fileServer := http.FileServer(http.Dir(s.outputDir))
	mux.Handle("/", s.injectLiveReload(fileServer))

	// Live reload endpoint
	mux.HandleFunc("/livereload", s.handleLiveReload)

	addr := fmt.Sprintf("%s:%d", s.opts.Host, s.opts.Port)
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// Start server in goroutine
	errCh := make(chan error, 1)
	go func() {
		if s.opts.Verbose {
			fmt.Printf("\nDev server running at http://%s\n", addr)
			fmt.Println("Press Ctrl+C to stop.")
		}
		errCh <- server.ListenAndServe()
	}()

	// Wait for context cancellation or server error
	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	case err := <-errCh:
		return err
	}
}

func (s *Server) rebuild(ctx context.Context) error {
	pres, err := LoadSpec(s.specPath)
	if err != nil {
		s.mu.Lock()
		s.buildError = err
		s.mu.Unlock()
		return err
	}

	if err := s.renderer.RenderToDir(ctx, pres, s.outputDir); err != nil {
		s.mu.Lock()
		s.buildError = err
		s.mu.Unlock()
		return err
	}

	s.mu.Lock()
	s.buildError = nil
	s.mu.Unlock()

	return nil
}

func (s *Server) watchFiles(ctx context.Context) {
	// Get initial mod time
	info, err := os.Stat(s.specPath)
	if err == nil {
		s.mu.Lock()
		s.lastModTime = info.ModTime()
		s.mu.Unlock()
	}

	ticker := time.NewTicker(s.opts.WatchInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			info, err := os.Stat(s.specPath)
			if err != nil {
				continue
			}

			s.mu.RLock()
			lastMod := s.lastModTime
			s.mu.RUnlock()

			if info.ModTime().After(lastMod) {
				s.mu.Lock()
				s.lastModTime = info.ModTime()
				s.mu.Unlock()

				if s.opts.Verbose {
					fmt.Printf("[%s] Rebuilding...\n", time.Now().Format("15:04:05"))
				}

				if err := s.rebuild(ctx); err != nil {
					if s.opts.Verbose {
						fmt.Printf("Build error: %v\n", err)
					}
				} else if s.opts.Verbose {
					fmt.Println("Rebuild complete")
				}
			}
		}
	}
}

func (s *Server) injectLiveReload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only inject for HTML files
		if r.URL.Path == "/" || filepath.Ext(r.URL.Path) == ".html" {
			// Read the file - path is cleaned via filepath.Join which removes ..
			path := filepath.Join(s.outputDir, filepath.Clean(r.URL.Path))
			if r.URL.Path == "/" {
				path = filepath.Join(s.outputDir, "index.html")
			}

			content, err := os.ReadFile(path) //nolint:gosec // Path is cleaned above
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// Inject live reload script before </body>
			liveReloadScript := `<script>
(function() {
  let lastCheck = Date.now();
  setInterval(function() {
    fetch('/livereload?since=' + lastCheck)
      .then(r => r.json())
      .then(data => {
        if (data.reload) {
          location.reload();
        }
        lastCheck = Date.now();
      })
      .catch(() => {});
  }, 1000);
})();
</script>`

			// Insert before </body>
			html := string(content)
			html = html[:len(html)-14] + liveReloadScript + html[len(html)-14:]

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = fmt.Fprint(w, html) //nolint:gosec // HTML is from our generated templates
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) handleLiveReload(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	modTime := s.lastModTime
	buildErr := s.buildError
	s.mu.RUnlock()

	// Check if client's timestamp is older than our last mod time
	sinceStr := r.URL.Query().Get("since")
	var since int64
	if sinceStr != "" {
		_, _ = fmt.Sscanf(sinceStr, "%d", &since)
	}

	reload := modTime.UnixMilli() > since && buildErr == nil

	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, `{"reload":%t,"error":%q}`, reload, func() string {
		if buildErr != nil {
			return buildErr.Error()
		}
		return ""
	}())
}
