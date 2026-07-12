// Package schema provides JSON Schema generation from Go types.
package schema

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/presentation-spec/spec"
	"github.com/invopop/jsonschema"
)

// Generator generates JSON Schema from Go types.
type Generator struct {
	reflector *jsonschema.Reflector
}

// NewGenerator creates a new schema generator.
func NewGenerator() *Generator {
	r := &jsonschema.Reflector{
		DoNotReference:             false,
		ExpandedStruct:             false,
		AllowAdditionalProperties:  false,
		RequiredFromJSONSchemaTags: true,
	}
	return &Generator{reflector: r}
}

// GeneratePresentationSchema generates the JSON Schema for PresentationSpec.
func (g *Generator) GeneratePresentationSchema() *jsonschema.Schema {
	schema := g.reflector.Reflect(&spec.PresentationSpec{})
	schema.ID = jsonschema.ID("https://github.com/grokify/presentation-spec/schema/presentation.schema.json")
	schema.Title = "PresentationSpec"
	schema.Description = "A format-agnostic intermediate representation for presentations."
	return schema
}

// SchemaJSON returns the JSON Schema as a formatted JSON byte slice.
func (g *Generator) SchemaJSON() ([]byte, error) {
	schema := g.GeneratePresentationSchema()
	return json.MarshalIndent(schema, "", "  ")
}

// WriteSchema writes the JSON Schema to a file.
func (g *Generator) WriteSchema(path string) error {
	data, err := g.SchemaJSON()
	if err != nil {
		return fmt.Errorf("failed to generate schema JSON: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil { //nolint:gosec // Generated schema needs to be world-readable
		return fmt.Errorf("failed to write schema file: %w", err)
	}
	return nil
}
