
package models

// MessageObjectTraitsItem represents a MessageObjectTraitsItem model.
type MessageObjectTraitsItem struct {
  Reference `json:"-,omitempty`
  MessageTrait `json:"-,omitempty`
  ModelinaArrType []Union `json:"-,omitempty`
}