package schema

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

// ValidationResult contains the result of schema validation.
type ValidationResult struct {
	// Valid indicates whether the document is valid.
	Valid bool

	// Errors contains validation error messages.
	Errors []string
}

// ValidateJSON validates a JSON byte slice against the PresentationSpec schema.
func ValidateJSON(data []byte) (*ValidationResult, error) {
	schemaLoader := gojsonschema.NewBytesLoader(presentationSchemaJSON)
	documentLoader := gojsonschema.NewBytesLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	vr := &ValidationResult{
		Valid:  result.Valid(),
		Errors: make([]string, 0, len(result.Errors())),
	}

	for _, err := range result.Errors() {
		vr.Errors = append(vr.Errors, err.String())
	}

	return vr, nil
}

// ValidateFile validates a JSON file against the PresentationSpec schema.
func ValidateFile(path string) (*ValidationResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return ValidateJSON(data)
}

// ValidateSpec validates a PresentationSpec struct against the schema.
// This marshals the struct to JSON first, then validates.
func ValidateSpec(spec any) (*ValidationResult, error) {
	data, err := json.Marshal(spec)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal spec: %w", err)
	}

	return ValidateJSON(data)
}

// MustBeValid panics if validation fails.
func MustBeValid(result *ValidationResult, err error) {
	if err != nil {
		panic(fmt.Sprintf("validation error: %v", err))
	}
	if !result.Valid {
		panic(fmt.Sprintf("invalid spec: %v", result.Errors))
	}
}
