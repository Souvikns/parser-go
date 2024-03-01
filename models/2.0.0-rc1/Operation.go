
package models

// Operation represents a Operation model.
type Operation struct {
  Traits []interface{} `json:traits`
  Summary string `json:summary`
  Description string `json:description`
  Tags []Tag `json:tags`
  ExternalDocs *ExternalDocs `json:externalDocs`
  OperationId string `json:operationId`
  ProtocolInfo map[string]map[string]interface{} `json:protocolInfo`
  Message interface{} `json:message`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}