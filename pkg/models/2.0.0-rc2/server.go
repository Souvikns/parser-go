
package models

// Server represents a Server model.
type Server struct {
  Url string `json:"url"`
  Description string `json:"description"`
  Protocol string `json:"protocol"`
  ProtocolVersion string `json:"protocolVersion"`
  Variables map[string]ServerVariable `json:"variables"`
  Security []map[string][]string `json:"security"`
  Bindings *BindingsObject `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}