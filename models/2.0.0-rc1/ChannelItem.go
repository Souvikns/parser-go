
package models

// ChannelItem represents a ChannelItem model.
type ChannelItem struct {
  Ref string
  Parameters map[string]interface{}
  Publish *Operation
  Subscribe *Operation
  Deprecated bool
  ProtocolInfo map[string]map[string]interface{}
  AdditionalProperties map[string]interface{}
}