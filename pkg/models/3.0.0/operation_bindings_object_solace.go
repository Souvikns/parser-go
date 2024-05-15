
package models

// OperationBindingsObjectSolace represents a OperationBindingsObjectSolace model.
type OperationBindingsObjectSolace struct {
  BindingVersion *OperationBindingsObjectSolaceBindingVersion `json:"bindingVersion"`
  Destinations []BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItem `json:"destinations"`
  TimeToLive int `json:"timeToLive"`
  Priority int `json:"priority"`
  DmqEligible bool `json:"dmqEligible"`
}