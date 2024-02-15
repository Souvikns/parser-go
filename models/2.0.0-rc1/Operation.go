
package models

// Operation represents a Operation model.
type Operation struct {
  Traits []interface{}
  Summary string
  Description string
  Tags []Tag
  ExternalDocs *ExternalDocs
  OperationId string
  ProtocolInfo map[string]map[string]interface{}
  Message interface{}
  AdditionalProperties map[string]interface{}
}