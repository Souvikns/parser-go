
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Summary string `json:"summary,omitempty"`
  Description string `json:"description,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  OperationId string `json:"operationId,omitempty"`
  ProtocolInfo map[string]map[string]interface{} `json:"protocolInfo,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}