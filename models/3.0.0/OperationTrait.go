
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Title string
  Summary string
  Description string
  Security []interface{}
  Tags []interface{}
  ExternalDocs interface{}
  Bindings interface{}
  AdditionalProperties map[string]interface{}
}