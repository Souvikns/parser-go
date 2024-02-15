
package models

// Operation represents a Operation model.
type Operation struct {
  Action *OperationAction
  Channel *Reference
  Messages []Reference
  Reply interface{}
  Traits []interface{}
  Title string
  Summary string
  Description string
  Security []interface{}
  Tags []interface{}
  ExternalDocs interface{}
  Bindings interface{}
  AdditionalProperties map[string]interface{}
}