
package models
import (  
  "encoding/json"
)
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

 
          
func (op *ApiKeyHttpSecuritySchemeType) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToApiKeyHttpSecuritySchemeType[v]
	return nil
}

func (op ApiKeyHttpSecuritySchemeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          