
package models

// Xml represents a Xml model.
type Xml struct {
  Name string `json:"name,omitempty"`
  Namespace string `json:"namespace,omitempty"`
  Prefix string `json:"prefix,omitempty"`
  Attribute bool `json:"attribute,omitempty"`
  Wrapped bool `json:"wrapped,omitempty"`
}