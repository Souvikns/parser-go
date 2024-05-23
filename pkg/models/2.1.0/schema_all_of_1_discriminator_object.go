
package models

// SchemaAllOf_1DiscriminatorObject represents a SchemaAllOf_1DiscriminatorObject model.
type SchemaAllOf_1DiscriminatorObject struct {
  PropertyName string `json:"propertyName,omitempty"`
  Mapping map[string]string `json:"mapping,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}