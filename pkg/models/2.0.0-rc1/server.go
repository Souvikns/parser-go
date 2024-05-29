
package models

// Server represents a Server model.
type Server struct {
  Url string `json:"url,omitempty"`
  Description string `json:"description,omitempty"`
  Protocol string `json:"protocol,omitempty"`
  ProtocolVersion string `json:"protocolVersion,omitempty"`
  Variables map[string]ServerVariable `json:"variables,omitempty"`
  BaseChannel string `json:"baseChannel,omitempty"`
  Security []map[string][]string `json:"security,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}