
package models

// Server represents a Server model.
type Server struct {
  Url string
  Description string
  Protocol string
  ProtocolVersion string
  Variables map[string]interface{}
  Security []map[string][]string
  Bindings *BindingsObject
  Tags []Tag
  AdditionalProperties map[string]interface{}
}