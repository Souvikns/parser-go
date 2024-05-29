
package models

// ApiKey represents a ApiKey model.
type ApiKey struct {
  ReservedType *ApiKeyType `json:"type,omitempty"`
  In *ApiKeyIn `json:"in,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}