
package models

// CoreSchemaMetaMinusSchemaObject represents a CoreSchemaMetaMinusSchemaObject model.
type CoreSchemaMetaMinusSchemaObject struct {
  Id string `json:"$id,omitempty"`
  Schema string `json:"$schema,omitempty"`
  Ref string `json:"$ref,omitempty"`
  Comment string `json:"$comment,omitempty"`
  Title string `json:"title,omitempty"`
  Description string `json:"description,omitempty"`
  ReservedDefault interface{} `json:"default,omitempty"`
  ReadOnly bool `json:"readOnly,omitempty"`
  WriteOnly bool `json:"writeOnly,omitempty"`
  Examples []interface{} `json:"examples,omitempty"`
  MultipleOf float64 `json:"multipleOf,omitempty"`
  Maximum float64 `json:"maximum,omitempty"`
  ExclusiveMaximum float64 `json:"exclusiveMaximum,omitempty"`
  Minimum float64 `json:"minimum,omitempty"`
  ExclusiveMinimum float64 `json:"exclusiveMinimum,omitempty"`
  MaxLength int `json:"maxLength,omitempty"`
  MinLength int `json:"minLength,omitempty"`
  Pattern string `json:"pattern,omitempty"`
  AdditionalItems *CoreSchemaMetaMinusSchema `json:"additionalItems,omitempty"`
  Items *JsonMinusSchemaMinusDraftMinus_07MinusSchemaItems `json:"items,omitempty"`
  MaxItems int `json:"maxItems,omitempty"`
  MinItems int `json:"minItems,omitempty"`
  UniqueItems bool `json:"uniqueItems,omitempty"`
  Contains *CoreSchemaMetaMinusSchema `json:"contains,omitempty"`
  MaxProperties int `json:"maxProperties,omitempty"`
  MinProperties int `json:"minProperties,omitempty"`
  Required []string `json:"required,omitempty"`
  AdditionalProperties *CoreSchemaMetaMinusSchema `json:"additionalProperties,omitempty"`
  Definitions map[string]AnySchema `json:"definitions,omitempty"`
  Properties map[string]AnySchema `json:"properties,omitempty"`
  PatternProperties map[string]AnySchema `json:"patternProperties,omitempty"`
  Dependencies map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchemaDependenciesAdditionalProperty `json:"dependencies,omitempty"`
  PropertyNames *CoreSchemaMetaMinusSchema `json:"propertyNames,omitempty"`
  ReservedConst interface{} `json:"const,omitempty"`
  Enum []interface{} `json:"enum,omitempty"`
  ReservedType *JsonMinusSchemaMinusDraftMinus_07MinusSchemaType `json:"type,omitempty"`
  Format string `json:"format,omitempty"`
  ContentMediaType string `json:"contentMediaType,omitempty"`
  ContentEncoding string `json:"contentEncoding,omitempty"`
  ReservedIf *CoreSchemaMetaMinusSchema `json:"if,omitempty"`
  Then *CoreSchemaMetaMinusSchema `json:"then,omitempty"`
  ReservedElse *CoreSchemaMetaMinusSchema `json:"else,omitempty"`
  AllOf []AnySchema `json:"allOf,omitempty"`
  AnyOf []AnySchema `json:"anyOf,omitempty"`
  OneOf []AnySchema `json:"oneOf,omitempty"`
  Not *CoreSchemaMetaMinusSchema `json:"not,omitempty"`
  ReservedAdditionalProperties map[string]ReservedAdditionalProperties `json:"-",omitempty`
}