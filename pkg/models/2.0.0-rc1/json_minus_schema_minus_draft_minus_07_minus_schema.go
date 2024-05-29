
package models

// JsonMinusSchemaMinusDraftMinus_07MinusSchema represents a JsonMinusSchemaMinusDraftMinus_07MinusSchema model.
type JsonMinusSchemaMinusDraftMinus_07MinusSchema struct {
  Id string `json:"id,omitempty"`
  Schema string `json:"$schema,omitempty"`
  Title string `json:"title,omitempty"`
  Description string `json:"description,omitempty"`
  ReservedDefault interface{} `json:"default,omitempty"`
  MultipleOf float64 `json:"multipleOf,omitempty"`
  Maximum float64 `json:"maximum,omitempty"`
  ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
  Minimum float64 `json:"minimum,omitempty"`
  ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`
  MaxLength int `json:"maxLength,omitempty"`
  MinLength int `json:"minLength,omitempty"`
  Pattern string `json:"pattern,omitempty"`
  AdditionalItems *JsonMinusSchemaMinusDraftMinus_07MinusSchemaAdditionalItems `json:"additionalItems,omitempty"`
  Items *JsonMinusSchemaMinusDraftMinus_07MinusSchemaItems `json:"items,omitempty"`
  MaxItems int `json:"maxItems,omitempty"`
  MinItems int `json:"minItems,omitempty"`
  UniqueItems bool `json:"uniqueItems,omitempty"`
  MaxProperties int `json:"maxProperties,omitempty"`
  MinProperties int `json:"minProperties,omitempty"`
  Required []string `json:"required,omitempty"`
  AdditionalProperties *JsonMinusSchemaMinusDraftMinus_07MinusSchemaAdditionalProperties `json:"additionalProperties,omitempty"`
  Definitions map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"definitions,omitempty"`
  Properties map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"properties,omitempty"`
  PatternProperties map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"patternProperties,omitempty"`
  Dependencies map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchemaDependenciesAdditionalProperty `json:"dependencies,omitempty"`
  Enum []interface{} `json:"enum,omitempty"`
  ReservedType *JsonMinusSchemaMinusDraftMinus_07MinusSchemaType `json:"type,omitempty"`
  Format string `json:"format,omitempty"`
  AllOf []JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"allOf,omitempty"`
  AnyOf []JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"anyOf,omitempty"`
  OneOf []JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"oneOf,omitempty"`
  Not *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"not,omitempty"`
  ReservedAdditionalProperties map[string]interface{} `json:"-,omitempty"`
}