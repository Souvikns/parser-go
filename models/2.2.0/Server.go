
package models

// Server represents a Server model.
type Server struct {
  Url string
  Description string
  Protocol string
  ProtocolVersion string
  Variables map[string]ServerVariable
  Security []map[string][]string
  Bindings *BindingsObject
  AdditionalProperties map[string]interface{}
}