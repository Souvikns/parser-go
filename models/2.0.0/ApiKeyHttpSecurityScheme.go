
package models

// ApiKeyHttpSecurityScheme represents a ApiKeyHttpSecurityScheme model.
type ApiKeyHttpSecurityScheme struct {
  ReservedType *ApiKeyHttpSecuritySchemeType `json:type`
  Name string `json:name`
  In *ApiKeyHttpSecuritySchemeIn `json:in`
  Description string `json:description`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}