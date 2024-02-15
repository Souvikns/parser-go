
package models

// JsonMinusSchemaMinusDraftMinus_07MinusSchema represents a JsonMinusSchemaMinusDraftMinus_07MinusSchema model.
type JsonMinusSchemaMinusDraftMinus_07MinusSchema struct {
  Id string
  Schema string
  Title string
  Description string
  ReservedDefault interface{}
  MultipleOf float64
  Maximum float64
  ExclusiveMaximum bool
  Minimum float64
  ExclusiveMinimum bool
  MaxLength int
  MinLength int
  Pattern string
  AdditionalItems interface{}
  Items interface{}
  MaxItems int
  MinItems int
  UniqueItems bool
  MaxProperties int
  MinProperties int
  Required []string
  AdditionalProperties interface{}
  Definitions map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Properties map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchema
  PatternProperties map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Dependencies map[string]interface{}
  Enum []interface{}
  ReservedType interface{}
  Format string
  AllOf []JsonMinusSchemaMinusDraftMinus_07MinusSchema
  AnyOf []JsonMinusSchemaMinusDraftMinus_07MinusSchema
  OneOf []JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Not *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  ReservedAdditionalProperties map[string]interface{}
}