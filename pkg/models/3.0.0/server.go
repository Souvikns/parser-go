
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
  Variables map[string]ServerVariablesAdditionalProperty `json:"variables"`
  Security []SecurityRequirementsItem `json:"security"`
  Tags []ServerTagsItem `json:"tags"`
  ExternalDocs *ServerExternalDocs `json:"externalDocs"`
  Bindings *ServerBindings `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}