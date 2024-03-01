
package models

// OperationBindingsObjectSolace represents a OperationBindingsObjectSolace model.
type OperationBindingsObjectSolace struct {
  BindingVersion *OperationBindingsObjectSolaceBindingVersion `json:"bindingVersion"`
  Destinations []interface{} `json:"destinations"`
  TimeToLive int `json:"timeToLive"`
  Priority int `json:"priority"`
  DmqEligible bool `json:"dmqEligible"`
}