
package models

// ApiKeyHttpSecuritySchemeIn represents an enum of ApiKeyHttpSecuritySchemeIn.
type ApiKeyHttpSecuritySchemeIn uint

const (
  ApiKeyHttpSecuritySchemeInHeader ApiKeyHttpSecuritySchemeIn = iota
  ApiKeyHttpSecuritySchemeInQuery
  ApiKeyHttpSecuritySchemeInCookie
)

// Value returns the value of the enum.
func (op ApiKeyHttpSecuritySchemeIn) Value() any {
	if op >= ApiKeyHttpSecuritySchemeIn(len(ApiKeyHttpSecuritySchemeInValues)) {
		return nil
	}
	return ApiKeyHttpSecuritySchemeInValues[op]
}

var ApiKeyHttpSecuritySchemeInValues = []any{"header","query","cookie"}
var ValuesToApiKeyHttpSecuritySchemeIn = map[any]ApiKeyHttpSecuritySchemeIn{
  ApiKeyHttpSecuritySchemeInValues[ApiKeyHttpSecuritySchemeInHeader]: ApiKeyHttpSecuritySchemeInHeader,
  ApiKeyHttpSecuritySchemeInValues[ApiKeyHttpSecuritySchemeInQuery]: ApiKeyHttpSecuritySchemeInQuery,
  ApiKeyHttpSecuritySchemeInValues[ApiKeyHttpSecuritySchemeInCookie]: ApiKeyHttpSecuritySchemeInCookie,
}
