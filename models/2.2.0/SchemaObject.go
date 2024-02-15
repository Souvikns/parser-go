
package models

// SchemaObject represents a SchemaObject model.
type SchemaObject struct {
  Id string
  Schema string
  Ref string
  Comment string
  Title string
  Description string
  ReservedDefault interface{}
  ReadOnly bool
  WriteOnly bool
  Examples []interface{}
  MultipleOf float64
  Maximum float64
  ExclusiveMaximum interface{}
  Minimum float64
  ExclusiveMinimum interface{}
  MaxLength int
  MinLength int
  Pattern string
  AdditionalItems interface{}
  Items interface{}
  MaxItems int
  MinItems int
  UniqueItems bool
  Contains interface{}
  MaxProperties int
  MinProperties int
  Required []string
  AdditionalProperties interface{}
  Definitions map[string]interface{}
  Properties map[string]interface{}
  PatternProperties map[string]interface{}
  Dependencies map[string]interface{}
  PropertyNames interface{}
  ReservedConst interface{}
  Enum []interface{}
  ReservedType interface{}
  Format string
  ContentMediaType string
  ContentEncoding string
  ReservedIf interface{}
  Then interface{}
  ReservedElse interface{}
  AllOf []interface{}
  AnyOf []interface{}
  OneOf []interface{}
  Not interface{}
  Discriminator interface{}
  ExternalDocs *ExternalDocs
  Deprecated bool
  ReservedAdditionalProperties map[string]interface{}
}