
package models

// JsonMinusSchemaMinusDraftMinus_07MinusSchema represents a JsonMinusSchemaMinusDraftMinus_07MinusSchema model.
type JsonMinusSchemaMinusDraftMinus_07MinusSchema struct {
  Id string `json:"id"`
  Schema string `json:"$schema"`
  Title string `json:"title"`
  Description string `json:"description"`
  ReservedDefault interface{} `json:"default"`
  MultipleOf float64 `json:"multipleOf"`
  Maximum float64 `json:"maximum"`
  ExclusiveMaximum bool `json:"exclusiveMaximum"`
  Minimum float64 `json:"minimum"`
  ExclusiveMinimum bool `json:"exclusiveMinimum"`
  MaxLength int `json:"maxLength"`
  MinLength int `json:"minLength"`
  Pattern string `json:"pattern"`
  AdditionalItems *JsonMinusSchemaMinusDraftMinus_07MinusSchemaAdditionalItems `json:"additionalItems"`
  Items *JsonMinusSchemaMinusDraftMinus_07MinusSchemaItems `json:"items"`
  MaxItems int `json:"maxItems"`
  MinItems int `json:"minItems"`
  UniqueItems bool `json:"uniqueItems"`
  MaxProperties int `json:"maxProperties"`
  MinProperties int `json:"minProperties"`
  Required []string `json:"required"`
  AdditionalProperties *JsonMinusSchemaMinusDraftMinus_07MinusSchemaAdditionalProperties `json:"additionalProperties"`
  Definitions map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"definitions"`
  Properties map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"properties"`
  PatternProperties map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"patternProperties"`
  Dependencies map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchemaDependenciesAdditionalProperty `json:"dependencies"`
  Enum []interface{} `json:"enum"`
  ReservedType *JsonMinusSchemaMinusDraftMinus_07MinusSchemaType `json:"type"`
  Format string `json:"format"`
  AllOf []JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"allOf"`
  AnyOf []JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"anyOf"`
  OneOf []JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"oneOf"`
  Not *JsonMinusSchemaMinusDraftMinus_07MinusSchema `json:"not"`
  ReservedAdditionalProperties map[string]interface{} `json:"-"`
}