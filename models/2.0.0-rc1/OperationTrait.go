
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Summary string
  Description string
  Tags []Tag
  ExternalDocs *ExternalDocs
  OperationId string
  ProtocolInfo map[string]map[string]interface{}
  AdditionalProperties map[string]interface{}
}