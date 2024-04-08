
package models

// Xml represents a Xml model.
type Xml struct {
  Name string `json:"name"`
  Namespace string `json:"namespace"`
  Prefix string `json:"prefix"`
  Attribute bool `json:"attribute"`
  Wrapped bool `json:"wrapped"`
}