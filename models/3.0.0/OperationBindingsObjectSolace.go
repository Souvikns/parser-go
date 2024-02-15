
package models

// OperationBindingsObjectSolace represents a OperationBindingsObjectSolace model.
type OperationBindingsObjectSolace struct {
  BindingVersion *OperationBindingsObjectSolaceBindingVersion
  Destinations []interface{}
  TimeToLive int
  Priority int
  DmqEligible bool
}