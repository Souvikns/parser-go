
package models
import (  
  "encoding/json"
)
// BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope represents an enum of BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope uint

const (
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeMessageAttributes BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope = iota
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeMessageBody
)

// Value returns the value of the enum.
func (op BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope) Value() any {
	if op >= BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope(len(BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeValues)) {
		return nil
	}
	return BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeValues[op]
}

var BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeValues = []any{"MessageAttributes","MessageBody"}
var ValuesToBindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope = map[any]BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope{
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeMessageAttributes]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeMessageAttributes,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeMessageBody]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScopeMessageBody,
}

 
func (op *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope[v]
  return nil
}

func (op BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}