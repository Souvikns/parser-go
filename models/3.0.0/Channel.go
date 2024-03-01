
package models

// Channel represents a Channel model.
type Channel struct {
  Address string `json:"address"`
  Messages map[string]interface{} `json:"messages"`
  Parameters map[string]interface{} `json:"parameters"`
  Title string `json:"title"`
  Summary string `json:"summary"`
  Description string `json:"description"`
  Servers []Reference `json:"servers"`
  Tags []interface{} `json:"tags"`
  ExternalDocs interface{} `json:"externalDocs"`
  Bindings interface{} `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}