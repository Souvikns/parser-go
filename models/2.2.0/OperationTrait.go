
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Summary string
  Description string
  Tags []Tag
  ExternalDocs *ExternalDocs
  OperationId string
  Bindings *BindingsObject
  AdditionalProperties map[string]interface{}
}