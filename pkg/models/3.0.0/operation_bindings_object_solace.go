
package models

// OperationBindingsObjectSolace represents a OperationBindingsObjectSolace model.
type OperationBindingsObjectSolace struct {
  BindingVersion *OperationBindingsObjectSolaceBindingVersion `json:"bindingVersion,omitempty"`
  Destinations []BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItem `json:"destinations,omitempty"`
  TimeToLive int `json:"timeToLive,omitempty"`
  Priority int `json:"priority,omitempty"`
  DmqEligible bool `json:"dmqEligible,omitempty"`
}