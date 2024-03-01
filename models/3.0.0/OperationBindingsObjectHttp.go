
package models

// OperationBindingsObjectHttp represents a OperationBindingsObjectHttp model.
type OperationBindingsObjectHttp struct {
  BindingVersion *OperationBindingsObjectHttpBindingVersion `json:"bindingVersion"`
  Method *BindingsMinusHttpMinus_0Dot_2Dot_0MinusOperationMethod `json:"method"`
  Query interface{} `json:"query"`
  AdditionalProperties map[string]interface{} `json:"-"`
}