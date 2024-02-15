
package models

// Operation represents a Operation model.
type Operation struct {
  Traits []interface{}
  Summary string
  Description string
  Tags []Tag
  ExternalDocs *ExternalDocs
  OperationId string
  Bindings *BindingsObject
  Message interface{}
  AdditionalProperties map[string]interface{}
}