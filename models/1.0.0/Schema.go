
package models

// Schema represents a Schema model.
type Schema struct {
  Ref string `json:"$ref"`
  Format string `json:"format"`
  Title *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"title"`
  Description *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"description"`
  ReservedDefault *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"default"`
  MultipleOf *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"multipleOf"`
  Maximum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"maximum"`
  ExclusiveMaximum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"exclusiveMaximum"`
  Minimum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"minimum"`
  ExclusiveMinimum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"exclusiveMinimum"`
  MaxLength *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"maxLength"`
  MinLength *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"minLength"`
  Pattern *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"pattern"`
  MaxItems *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"maxItems"`
  MinItems *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"minItems"`
  UniqueItems *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"uniqueItems"`
  MaxProperties *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"maxProperties"`
  MinProperties *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"minProperties"`
  Required *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"required"`
  Enum *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"enum"`
  AdditionalProperties *SchemaAdditionalProperties `json:"additionalProperties"`
  ReservedType *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"type"`
  Items *SchemaItems `json:"items"`
  AllOf []Schema `json:"allOf"`
  Properties map[string]Schema `json:"properties"`
  Discriminator string `json:"discriminator"`
  ReadOnly bool `json:"readOnly"`
  Xml *Xml `json:"xml"`
  ExternalDocs *ExternalDocs `json:"externalDocs"`
  Example interface{} `json:"example"`
  ReservedAdditionalProperties map[string]interface{} `json:"-"`
}