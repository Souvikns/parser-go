
package models

// SchemaAllOf_1DiscriminatorObject represents a SchemaAllOf_1DiscriminatorObject model.
type SchemaAllOf_1DiscriminatorObject struct {
  PropertyName string `json:"propertyName"`
  Mapping map[string]string `json:"mapping"`
  AdditionalProperties map[string]interface{} `json:"-"`
}