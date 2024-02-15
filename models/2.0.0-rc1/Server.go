
package models

// Server represents a Server model.
type Server struct {
  Url string
  Description string
  Protocol string
  ProtocolVersion string
  Variables map[string]ServerVariable
  BaseChannel string
  Security []map[string][]string
  AdditionalProperties map[string]interface{}
}