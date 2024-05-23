
package models

// Server represents a Server model.
type Server struct {
  Url string `json:"url,omitempty"`
  Description string `json:"description,omitempty"`
  Scheme *ServerScheme `json:"scheme,omitempty"`
  SchemeVersion string `json:"schemeVersion,omitempty"`
  Variables map[string]ServerVariable `json:"variables,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}