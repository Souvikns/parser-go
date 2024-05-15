
package models

// ServerBindingsObjectIbmmq represents a ServerBindingsObjectIbmmq model.
type ServerBindingsObjectIbmmq struct {
  BindingVersion *ServerBindingsObjectIbmmqBindingVersion `json:"bindingVersion"`
  GroupId string `json:"groupId"`
  CcdtQueueManagerName string `json:"ccdtQueueManagerName"`
  CipherSpec string `json:"cipherSpec"`
  MultiEndpointServer bool `json:"multiEndpointServer"`
  HeartBeatInterval int `json:"heartBeatInterval"`
  AdditionalProperties map[string]interface{} `json:"-"`
}