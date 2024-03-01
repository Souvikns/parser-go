
package models

// Operation represents a Operation model.
type Operation struct {
  Action *OperationAction `json:"action"`
  Channel *Reference `json:"channel"`
  Messages []Reference `json:"messages"`
  Reply interface{} `json:"reply"`
  Traits []interface{} `json:"traits"`
  Title string `json:"title"`
  Summary string `json:"summary"`
  Description string `json:"description"`
  Security []interface{} `json:"security"`
  Tags []interface{} `json:"tags"`
  ExternalDocs interface{} `json:"externalDocs"`
  Bindings interface{} `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}