
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Summary string `json:summary`
  Description string `json:description`
  Tags []Tag `json:tags`
  ExternalDocs *ExternalDocs `json:externalDocs`
  OperationId string `json:operationId`
  ProtocolInfo map[string]map[string]interface{} `json:protocolInfo`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}