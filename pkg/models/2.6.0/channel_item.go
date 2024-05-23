
package models

// ChannelItem represents a ChannelItem model.
type ChannelItem struct {
  Ref string `json:"$ref,omitempty"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters,omitempty"`
  Description string `json:"description,omitempty"`
  Servers []string `json:"servers,omitempty"`
  Publish *Operation `json:"publish,omitempty"`
  Subscribe *Operation `json:"subscribe,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  Bindings *BindingsObject `json:"bindings,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}