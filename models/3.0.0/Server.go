
package models

// Server represents a Server model.
type Server struct {
  Host string
  Pathname string
  Title string
  Summary string
  Description string
  Protocol string
  ProtocolVersion string
  Variables map[string]interface{}
  Security []interface{}
  Tags []interface{}
  ExternalDocs interface{}
  Bindings interface{}
  AdditionalProperties map[string]interface{}
}