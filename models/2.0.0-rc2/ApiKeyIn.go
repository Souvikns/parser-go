
package models

// ApiKeyIn represents an enum of ApiKeyIn.
type ApiKeyIn uint

const (
  ApiKeyInUser ApiKeyIn = iota
  ApiKeyInPassword
)

// Value returns the value of the enum.
func (op ApiKeyIn) Value() any {
	if op >= ApiKeyIn(len(ApiKeyInValues)) {
		return nil
	}
	return ApiKeyInValues[op]
}

var ApiKeyInValues = []any{"user","password"}
var ValuesToApiKeyIn = map[any]ApiKeyIn{
  ApiKeyInValues[ApiKeyInUser]: ApiKeyInUser,
  ApiKeyInValues[ApiKeyInPassword]: ApiKeyInPassword,
}
