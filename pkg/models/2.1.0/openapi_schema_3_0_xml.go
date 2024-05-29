
package models

// OpenapiSchema_3_0Xml represents a OpenapiSchema_3_0Xml model.
type OpenapiSchema_3_0Xml struct {
  Name string `json:"name,omitempty"`
  Namespace string `json:"namespace,omitempty"`
  Prefix string `json:"prefix,omitempty"`
  Attribute bool `json:"attribute,omitempty"`
  Wrapped bool `json:"wrapped,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}