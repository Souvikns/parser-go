
package models
import (  
  "encoding/json"
)
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

 
          
func (op *ApiKeyHttpSecuritySchemeIn) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToApiKeyHttpSecuritySchemeIn[v]
	return nil
}

func (op ApiKeyHttpSecuritySchemeIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          