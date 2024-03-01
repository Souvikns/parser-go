
package models

// Server represents a Server model.
type Server struct {
  Url string `json:url`
  Description string `json:description`
  Protocol string `json:protocol`
  ProtocolVersion string `json:protocolVersion`
  Variables map[string]interface{} `json:variables`
  Security []map[string][]string `json:security`
  Bindings *BindingsObject `json:bindings`
  Tags []Tag `json:tags`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}