
package models

// OpenapiSchema_3_0Xml represents a OpenapiSchema_3_0Xml model.
type OpenapiSchema_3_0Xml struct {
  Name string `json:"name"`
  Namespace string `json:"namespace"`
  Prefix string `json:"prefix"`
  Attribute bool `json:"attribute"`
  Wrapped bool `json:"wrapped"`
  AdditionalProperties map[string]interface{} `json:"-"`
}