
package models
import (  
  "encoding/json"
)
// BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion represents an enum of BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion.
type BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion uint

const (
  BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersionNumber_0Dot_0Dot_1 BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion = iota
)

// Value returns the value of the enum.
func (op BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion) Value() any {
	if op >= BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion(len(BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersionValues)) {
		return nil
	}
	return BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersionValues[op]
}

var BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersionValues = []any{"0.0.1"}
var ValuesToBindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion = map[any]BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion{
  BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersionValues[BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersionNumber_0Dot_0Dot_1]: BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersionNumber_0Dot_0Dot_1,
}

 
func (op *BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion[v]
  return nil
}

func (op BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}