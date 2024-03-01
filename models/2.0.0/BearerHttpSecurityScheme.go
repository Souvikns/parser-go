
package models

// BearerHttpSecurityScheme represents a BearerHttpSecurityScheme model.
type BearerHttpSecurityScheme struct {
  Scheme *BearerHttpSecuritySchemeScheme `json:scheme`
  BearerFormat string `json:bearerFormat`
  ReservedType *BearerHttpSecuritySchemeType `json:type`
  Description string `json:description`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}