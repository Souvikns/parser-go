
package models

// OperationReplyAddress represents a OperationReplyAddress model.
type OperationReplyAddress struct {
  Location string `json:"location,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}