
package models

// ApiKeyHttpSecurityScheme represents a ApiKeyHttpSecurityScheme model.
type ApiKeyHttpSecurityScheme struct {
  ReservedType *ApiKeyHttpSecuritySchemeType
  Name string
  In *ApiKeyHttpSecuritySchemeIn
  Description string
  AdditionalProperties map[string]interface{}
}