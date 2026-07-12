// Command genschema generates the JSON Schema from Go types.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grokify/presentation-spec/schema"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "genschema",
	Short: "Generate JSON Schema from PresentationSpec Go types",
	Long: `genschema generates a JSON Schema file from the PresentationSpec Go types.
The generated schema can be used for validation and editor support.`,
	RunE: runGenerate,
}

var outputPath string

func init() {
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "schema/presentation.schema.json",
		"Output path for the generated schema")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	// Ensure output directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate schema
	gen := schema.NewGenerator()
	if err := gen.WriteSchema(outputPath); err != nil {
		return err
	}

	fmt.Printf("Schema written to %s\n", outputPath)
	return nil
}
