package schema

import (
	_ "embed"
)

//go:embed presentation.schema.json
var presentationSchemaJSON []byte

// PresentationSchemaJSON returns the embedded JSON Schema for PresentationSpec.
func PresentationSchemaJSON() []byte {
	return presentationSchemaJSON
}
