
package models

// ChannelItem represents a ChannelItem model.
type ChannelItem struct {
  Ref string `json:"$ref,omitempty"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters,omitempty"`
  Publish *Operation `json:"publish,omitempty"`
  Subscribe *Operation `json:"subscribe,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  ProtocolInfo map[string]map[string]interface{} `json:"protocolInfo,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}