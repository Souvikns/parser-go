
package models

// Server represents a Server model.
type Server struct {
  Url string `json:url`
  Description string `json:description`
  Scheme *ServerScheme `json:scheme`
  SchemeVersion string `json:schemeVersion`
  Variables map[string]ServerVariable `json:variables`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}