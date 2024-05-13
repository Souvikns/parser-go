
package models

// ChannelItem represents a ChannelItem model.
type ChannelItem struct {
  Ref string `json:"$ref"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters"`
  Publish *Operation `json:"publish"`
  Subscribe *Operation `json:"subscribe"`
  Deprecated bool `json:"deprecated"`
  ProtocolInfo map[string]map[string]interface{} `json:"protocolInfo"`
  AdditionalProperties map[string]interface{} `json:"-"`
}