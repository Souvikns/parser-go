
package models

// ApiKeyHttpSecuritySchemeType represents an enum of ApiKeyHttpSecuritySchemeType.
type ApiKeyHttpSecuritySchemeType uint

const (
  ApiKeyHttpSecuritySchemeTypeHttpApiKey ApiKeyHttpSecuritySchemeType = iota
)

// Value returns the value of the enum.
func (op ApiKeyHttpSecuritySchemeType) Value() any {
	if op >= ApiKeyHttpSecuritySchemeType(len(ApiKeyHttpSecuritySchemeTypeValues)) {
		return nil
	}
	return ApiKeyHttpSecuritySchemeTypeValues[op]
}

var ApiKeyHttpSecuritySchemeTypeValues = []any{"httpApiKey"}
var ValuesToApiKeyHttpSecuritySchemeType = map[any]ApiKeyHttpSecuritySchemeType{
  ApiKeyHttpSecuritySchemeTypeValues[ApiKeyHttpSecuritySchemeTypeHttpApiKey]: ApiKeyHttpSecuritySchemeTypeHttpApiKey,
}
