
package models

// Server represents a Server model.
type Server struct {
  Host string `json:"host"`
  Pathname string `json:"pathname"`
  Title string `json:"title"`
  Summary string `json:"summary"`
  Description string `json:"description"`
  Protocol string `json:"protocol"`
  ProtocolVersion string `json:"protocolVersion"`
  Variables map[string]interface{} `json:"variables"`
  Security []interface{} `json:"security"`
  Tags []interface{} `json:"tags"`
  ExternalDocs interface{} `json:"externalDocs"`
  Bindings interface{} `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}