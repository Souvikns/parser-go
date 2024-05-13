
package models

// ServerBindingsObjectSolace represents a ServerBindingsObjectSolace model.
type ServerBindingsObjectSolace struct {
  BindingVersion *ServerBindingsObjectSolaceBindingVersion `json:"bindingVersion"`
  MsgVpn string `json:"msgVpn"`
  ClientName string `json:"clientName"`
  MsvVpn string `json:"msvVpn"`
  AdditionalProperties map[string]interface{} `json:"-"`
}