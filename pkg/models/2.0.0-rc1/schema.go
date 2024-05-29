
package models

// Schema represents a Schema model.
type Schema struct {
  Ref string `json:"$ref,omitempty"`
  Format string `json:"format,omitempty"`
  Title *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"title,omitempty"`
  Description *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"description,omitempty"`
  ReservedDefault *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"default,omitempty"`
  MultipleOf *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"multipleOf,omitempty"`
  Maximum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"maximum,omitempty"`
  ExclusiveMaximum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"exclusiveMaximum,omitempty"`
  Minimum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"minimum,omitempty"`
  ExclusiveMinimum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"exclusiveMinimum,omitempty"`
  MaxLength *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"maxLength,omitempty"`
  MinLength *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"minLength,omitempty"`
  Pattern *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"pattern,omitempty"`
  MaxItems *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"maxItems,omitempty"`
  MinItems *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"minItems,omitempty"`
  UniqueItems *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"uniqueItems,omitempty"`
  MaxProperties *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"maxProperties,omitempty"`
  MinProperties *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"minProperties,omitempty"`
  Required *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"required,omitempty"`
  Enum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"enum,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  AdditionalProperties *SchemaAdditionalProperties `json:"additionalProperties,omitempty"`
  ReservedType *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"type,omitempty"`
  Items *SchemaItems `json:"items,omitempty"`
  AllOf []Schema `json:"allOf,omitempty"`
  OneOf []Schema `json:"oneOf,omitempty"`
  AnyOf []Schema `json:"anyOf,omitempty"`
  Not *Schema `json:"not,omitempty"`
  Properties map[string]Schema `json:"properties,omitempty"`
  Discriminator string `json:"discriminator,omitempty"`
  ReadOnly bool `json:"readOnly,omitempty"`
  Xml *Xml `json:"xml,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  Example interface{} `json:"example,omitempty"`
  Examples []interface{} `json:"examples,omitempty"`
  ReservedAdditionalProperties map[string]interface{} `json:"-,omitempty"`
}