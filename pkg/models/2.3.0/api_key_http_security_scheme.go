
package models

// ApiKeyHttpSecurityScheme represents a ApiKeyHttpSecurityScheme model.
type ApiKeyHttpSecurityScheme struct {
  ReservedType *ApiKeyHttpSecuritySchemeType `json:"type,omitempty"`
  Name string `json:"name,omitempty"`
  In *ApiKeyHttpSecuritySchemeIn `json:"in,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}