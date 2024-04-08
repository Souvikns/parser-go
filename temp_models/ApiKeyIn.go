
package models

import (
	"encoding/json"
)

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

// func (op ApiKeyIn) Key(value string) any {

// }

var ApiKeyInValues = []any{"user","password"}
var ValuesToApiKeyIn = map[any]ApiKeyIn{
  ApiKeyInValues[ApiKeyInUser]: ApiKeyInUser,
  ApiKeyInValues[ApiKeyInPassword]: ApiKeyInPassword,
}

func (op *ApiKeyIn) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToApiKeyIn[v]
	return nil
}

func (op ApiKeyIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
}