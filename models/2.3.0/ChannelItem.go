
package models

// ChannelItem represents a ChannelItem model.
type ChannelItem struct {
  Ref string
  Parameters map[string]interface{}
  Description string
  Servers []string
  Publish *Operation
  Subscribe *Operation
  Deprecated bool
  Bindings *BindingsObject
  AdditionalProperties map[string]interface{}
}