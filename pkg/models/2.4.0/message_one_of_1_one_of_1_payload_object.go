
package models

// MessageOneOf_1OneOf_1PayloadObject represents a MessageOneOf_1OneOf_1PayloadObject model.
type MessageOneOf_1OneOf_1PayloadObject struct {
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
  ExclusiveMaximum *JsonMinusSchemaMinusDraftMinus_07MinusSchemaExclusiveMaximum `json:"exclusiveMaximum,omitempty"`
  Minimum float64 `json:"minimum,omitempty"`
  ExclusiveMinimum *JsonMinusSchemaMinusDraftMinus_07MinusSchemaExclusiveMinimum `json:"exclusiveMinimum,omitempty"`
  MaxLength int `json:"maxLength,omitempty"`
  MinLength int `json:"minLength,omitempty"`
  Pattern string `json:"pattern,omitempty"`
  AdditionalItems *MessageOneOf_1OneOf_1Payload `json:"additionalItems,omitempty"`
  Items *JsonMinusSchemaMinusDraftMinus_07MinusSchemaItems `json:"items,omitempty"`
  MaxItems int `json:"maxItems,omitempty"`
  MinItems int `json:"minItems,omitempty"`
  UniqueItems bool `json:"uniqueItems,omitempty"`
  Contains *MessageOneOf_1OneOf_1Payload `json:"contains,omitempty"`
  MaxProperties int `json:"maxProperties,omitempty"`
  MinProperties int `json:"minProperties,omitempty"`
  Required []string `json:"required,omitempty"`
  AdditionalProperties *MessageOneOf_1OneOf_1Payload `json:"additionalProperties,omitempty"`
  Definitions map[string]MessageOneOf_1OneOf_1Payload `json:"definitions,omitempty"`
  Properties map[string]MessageOneOf_1OneOf_1Payload `json:"properties,omitempty"`
  PatternProperties map[string]MessageOneOf_1OneOf_1Payload `json:"patternProperties,omitempty"`
  Dependencies map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchemaDependenciesAdditionalProperty `json:"dependencies,omitempty"`
  PropertyNames *MessageOneOf_1OneOf_1Payload `json:"propertyNames,omitempty"`
  ReservedConst interface{} `json:"const,omitempty"`
  Enum []interface{} `json:"enum,omitempty"`
  ReservedType *JsonMinusSchemaMinusDraftMinus_07MinusSchemaType `json:"type,omitempty"`
  Format string `json:"format,omitempty"`
  ContentMediaType string `json:"contentMediaType,omitempty"`
  ContentEncoding string `json:"contentEncoding,omitempty"`
  ReservedIf *MessageOneOf_1OneOf_1Payload `json:"if,omitempty"`
  Then *MessageOneOf_1OneOf_1Payload `json:"then,omitempty"`
  ReservedElse *MessageOneOf_1OneOf_1Payload `json:"else,omitempty"`
  AllOf []MessageOneOf_1OneOf_1Payload `json:"allOf,omitempty"`
  AnyOf []MessageOneOf_1OneOf_1Payload `json:"anyOf,omitempty"`
  OneOf []MessageOneOf_1OneOf_1Payload `json:"oneOf,omitempty"`
  Not *MessageOneOf_1OneOf_1Payload `json:"not,omitempty"`
  Discriminator *SchemaAllOf_1Discriminator `json:"discriminator,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  Nullable bool `json:"nullable,omitempty"`
  Example interface{} `json:"example,omitempty"`
  Xml *OpenapiSchema_3_0Xml `json:"xml,omitempty"`
  ReservedAdditionalProperties map[string]ReservedAdditionalProperties `json:"-",omitempty`
}