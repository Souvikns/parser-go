
package models
import (  
  "encoding/json"
)
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

 
func (op *ApiKeyType) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToApiKeyType[v]
  return nil
}

func (op ApiKeyType) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}