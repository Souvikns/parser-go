
package models

// Server represents a Server model.
type Server struct {
  Host string `json:"host,omitempty"`
  Pathname string `json:"pathname,omitempty"`
  Title string `json:"title,omitempty"`
  Summary string `json:"summary,omitempty"`
  Description string `json:"description,omitempty"`
  Protocol string `json:"protocol,omitempty"`
  ProtocolVersion string `json:"protocolVersion,omitempty"`
  Variables map[string]ServerVariablesAdditionalProperty `json:"variables,omitempty"`
  Security []SecurityRequirementsItem `json:"security,omitempty"`
  Tags []ServerTagsItem `json:"tags,omitempty"`
  ExternalDocs *ServerExternalDocs `json:"externalDocs,omitempty"`
  Bindings *ServerBindings `json:"bindings,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}