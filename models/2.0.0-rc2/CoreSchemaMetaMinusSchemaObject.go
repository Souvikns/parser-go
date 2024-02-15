
package models

// CoreSchemaMetaMinusSchemaObject represents a CoreSchemaMetaMinusSchemaObject model.
type CoreSchemaMetaMinusSchemaObject struct {
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
  ExclusiveMaximum float64
  Minimum float64
  ExclusiveMinimum float64
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
  ReservedAdditionalProperties map[string]interface{}
}