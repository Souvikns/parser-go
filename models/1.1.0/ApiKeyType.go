
package models

// ApiKeyType represents an enum of ApiKeyType.
type ApiKeyType uint

const (
  ApiKeyTypeApiKey ApiKeyType = iota
)

// Value returns the value of the enum.
func (op ApiKeyType) Value() any {
	if op >= ApiKeyType(len(ApiKeyTypeValues)) {
		return nil
	}
	return ApiKeyTypeValues[op]
}

var ApiKeyTypeValues = []any{"apiKey"}
var ValuesToApiKeyType = map[any]ApiKeyType{
  ApiKeyTypeValues[ApiKeyTypeApiKey]: ApiKeyTypeApiKey,
}
