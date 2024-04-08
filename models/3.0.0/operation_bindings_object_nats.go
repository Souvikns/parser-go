
package models

// OperationBindingsObjectNats represents a OperationBindingsObjectNats model.
type OperationBindingsObjectNats struct {
  BindingVersion *OperationBindingsObjectNatsBindingVersion `json:"bindingVersion"`
  Queue string `json:"queue"`
  AdditionalProperties map[string]interface{} `json:"-"`
}