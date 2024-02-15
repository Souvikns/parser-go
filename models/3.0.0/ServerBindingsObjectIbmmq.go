
package models

// ServerBindingsObjectIbmmq represents a ServerBindingsObjectIbmmq model.
type ServerBindingsObjectIbmmq struct {
  BindingVersion *ServerBindingsObjectIbmmqBindingVersion
  GroupId string
  CcdtQueueManagerName string
  CipherSpec string
  MultiEndpointServer bool
  HeartBeatInterval int
  AdditionalProperties map[string]interface{}
}