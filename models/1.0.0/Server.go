
package models

// Server represents a Server model.
type Server struct {
  Url string
  Description string
  Scheme *ServerScheme
  SchemeVersion string
  Variables map[string]ServerVariable
  AdditionalProperties map[string]interface{}
}