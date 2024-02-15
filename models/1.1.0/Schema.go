
package models

// Schema represents a Schema model.
type Schema struct {
  Ref string
  Format string
  Title *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Description *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  ReservedDefault *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  MultipleOf *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Maximum *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  ExclusiveMaximum *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Minimum *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  ExclusiveMinimum *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  MaxLength *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  MinLength *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Pattern *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  MaxItems *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  MinItems *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  UniqueItems *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  MaxProperties *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  MinProperties *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Required *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Enum *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  AdditionalProperties interface{}
  ReservedType *JsonMinusSchemaMinusDraftMinus_07MinusSchema
  Items interface{}
  AllOf []Schema
  OneOf []Schema
  AnyOf []Schema
  Not *Schema
  Properties map[string]Schema
  Discriminator string
  ReadOnly bool
  Xml *Xml
  ExternalDocs *ExternalDocs
  Example interface{}
  ReservedAdditionalProperties map[string]interface{}
}