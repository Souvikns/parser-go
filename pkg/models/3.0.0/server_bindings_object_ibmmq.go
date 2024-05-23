
package models

// ServerBindingsObjectIbmmq represents a ServerBindingsObjectIbmmq model.
type ServerBindingsObjectIbmmq struct {
  BindingVersion *ServerBindingsObjectIbmmqBindingVersion `json:"bindingVersion,omitempty"`
  GroupId string `json:"groupId,omitempty"`
  CcdtQueueManagerName string `json:"ccdtQueueManagerName,omitempty"`
  CipherSpec string `json:"cipherSpec,omitempty"`
  MultiEndpointServer bool `json:"multiEndpointServer,omitempty"`
  HeartBeatInterval int `json:"heartBeatInterval,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}