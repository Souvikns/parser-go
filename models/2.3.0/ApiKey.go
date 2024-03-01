
package models

// ApiKey represents a ApiKey model.
type ApiKey struct {
  ReservedType *ApiKeyType `json:type`
  In *ApiKeyIn `json:in`
  Description string `json:description`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}