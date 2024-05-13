
package models
import (  
  "encoding/json"
)
// BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation represents an enum of BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation.
type BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation uint

const (
  BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationHeader BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation = iota
  BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationPayload
)

// Value returns the value of the enum.
func (op BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation) Value() any {
	if op >= BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation(len(BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationValues)) {
		return nil
	}
	return BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationValues[op]
}

var BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationValues = []any{"header","payload"}
var ValuesToBindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation = map[any]BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation{
  BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationValues[BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationHeader]: BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationHeader,
  BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationValues[BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationPayload]: BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocationPayload,
}

 
func (op *BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation[v]
  return nil
}

func (op BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}