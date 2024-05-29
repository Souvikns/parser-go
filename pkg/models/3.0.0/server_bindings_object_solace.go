
package models

// ServerBindingsObjectSolace represents a ServerBindingsObjectSolace model.
type ServerBindingsObjectSolace struct {
  BindingVersion *ServerBindingsObjectSolaceBindingVersion `json:"bindingVersion,omitempty"`
  MsgVpn string `json:"msgVpn,omitempty"`
  ClientName string `json:"clientName,omitempty"`
  MsvVpn string `json:"msvVpn,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}