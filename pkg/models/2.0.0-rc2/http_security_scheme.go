
package models

// HttpSecurityScheme represents a HttpSecurityScheme model.
type HttpSecurityScheme struct {
  ModelinaAnyType interface{} `json:"-,omitempty`
  BearerHttpSecurityScheme `json:"-,omitempty`
  ApiKeyHttpSecurityScheme `json:"-,omitempty`
}