
package models

// OperationBindingsObjectHttp represents a OperationBindingsObjectHttp model.
type OperationBindingsObjectHttp struct {
  BindingVersion *OperationBindingsObjectHttpBindingVersion `json:"bindingVersion,omitempty"`
  Method *BindingsMinusHttpMinus_0Dot_3Dot_0MinusOperationMethod `json:"method,omitempty"`
  Query *Schema `json:"query,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}