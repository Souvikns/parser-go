
package models

// OperationBindingsObjectHttp represents a OperationBindingsObjectHttp model.
type OperationBindingsObjectHttp struct {
  BindingVersion *OperationBindingsObjectHttpBindingVersion `json:"bindingVersion"`
  Method *BindingsMinusHttpMinus_0Dot_2Dot_0MinusOperationMethod `json:"method"`
  Query *Schema `json:"query"`
  AdditionalProperties map[string]interface{} `json:"-"`
}