
package models

// ChannelItem represents a ChannelItem model.
type ChannelItem struct {
  Ref string `json:"$ref"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters"`
  Description string `json:"description"`
  Servers []string `json:"servers"`
  Publish *Operation `json:"publish"`
  Subscribe *Operation `json:"subscribe"`
  Deprecated bool `json:"deprecated"`
  Bindings *BindingsObject `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}