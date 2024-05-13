
package models
import (  
  "encoding/json"
)
// Asyncapi represents an enum of Asyncapi.
type Asyncapi uint

const (
  AsyncapiNumber_2Dot_0Dot_0MinusRc1 Asyncapi = iota
)

// Value returns the value of the enum.
func (op Asyncapi) Value() any {
	if op >= Asyncapi(len(AsyncapiValues)) {
		return nil
	}
	return AsyncapiValues[op]
}

var AsyncapiValues = []any{"2.0.0-rc1"}
var ValuesToAsyncapi = map[any]Asyncapi{
  AsyncapiValues[AsyncapiNumber_2Dot_0Dot_0MinusRc1]: AsyncapiNumber_2Dot_0Dot_0MinusRc1,
}

 
func (op *Asyncapi) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToAsyncapi[v]
  return nil
}

func (op Asyncapi) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}