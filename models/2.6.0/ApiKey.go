
package models

// ApiKey represents a ApiKey model.
type ApiKey struct {
  ReservedType *ApiKeyType
  In *ApiKeyIn
  Description string
  AdditionalProperties map[string]interface{}
}