
package models

// Channel represents a Channel model.
type Channel struct {
  Address *string `json:"address,omitempty"`
  Messages map[string]ChannelMessagesAdditionalProperty `json:"messages,omitempty"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters,omitempty"`
  Title string `json:"title,omitempty"`
  Summary string `json:"summary,omitempty"`
  Description string `json:"description,omitempty"`
  Servers []Reference `json:"servers,omitempty"`
  Tags []ChannelTagsItem `json:"tags,omitempty"`
  ExternalDocs *ChannelExternalDocs `json:"externalDocs,omitempty"`
  Bindings *ChannelBindings `json:"bindings,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}