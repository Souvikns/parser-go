
package models

// OperationBindingsObjectNats represents a OperationBindingsObjectNats model.
type OperationBindingsObjectNats struct {
  BindingVersion *OperationBindingsObjectNatsBindingVersion `json:"bindingVersion,omitempty"`
  Queue string `json:"queue,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}