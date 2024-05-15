
package models

// Channel represents a Channel model.
type Channel struct {
  Address *string `json:"address"`
  Messages map[string]ChannelMessagesAdditionalProperty `json:"messages"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters"`
  Title string `json:"title"`
  Summary string `json:"summary"`
  Description string `json:"description"`
  Servers []Reference `json:"servers"`
  Tags []ChannelTagsItem `json:"tags"`
  ExternalDocs *ChannelExternalDocs `json:"externalDocs"`
  Bindings *ChannelBindings `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}