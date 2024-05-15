
package models

// MessageOneOf_1OneOf_1PayloadObject represents a MessageOneOf_1OneOf_1PayloadObject model.
type MessageOneOf_1OneOf_1PayloadObject struct {
  Id string `json:"$id"`
  Schema string `json:"$schema"`
  Ref string `json:"$ref"`
  Comment string `json:"$comment"`
  Title string `json:"title"`
  Description string `json:"description"`
  ReservedDefault interface{} `json:"default"`
  ReadOnly bool `json:"readOnly"`
  WriteOnly bool `json:"writeOnly"`
  Examples []interface{} `json:"examples"`
  MultipleOf float64 `json:"multipleOf"`
  Maximum float64 `json:"maximum"`
  ExclusiveMaximum *JsonMinusSchemaMinusDraftMinus_07MinusSchemaExclusiveMaximum `json:"exclusiveMaximum"`
  Minimum float64 `json:"minimum"`
  ExclusiveMinimum *JsonMinusSchemaMinusDraftMinus_07MinusSchemaExclusiveMinimum `json:"exclusiveMinimum"`
  MaxLength int `json:"maxLength"`
  MinLength int `json:"minLength"`
  Pattern string `json:"pattern"`
  AdditionalItems *MessageOneOf_1OneOf_1Payload `json:"additionalItems"`
  Items *JsonMinusSchemaMinusDraftMinus_07MinusSchemaItems `json:"items"`
  MaxItems int `json:"maxItems"`
  MinItems int `json:"minItems"`
  UniqueItems bool `json:"uniqueItems"`
  Contains *MessageOneOf_1OneOf_1Payload `json:"contains"`
  MaxProperties int `json:"maxProperties"`
  MinProperties int `json:"minProperties"`
  Required []string `json:"required"`
  AdditionalProperties *MessageOneOf_1OneOf_1Payload `json:"additionalProperties"`
  Definitions map[string]MessageOneOf_1OneOf_1Payload `json:"definitions"`
  Properties map[string]MessageOneOf_1OneOf_1Payload `json:"properties"`
  PatternProperties map[string]MessageOneOf_1OneOf_1Payload `json:"patternProperties"`
  Dependencies map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchemaDependenciesAdditionalProperty `json:"dependencies"`
  PropertyNames *MessageOneOf_1OneOf_1Payload `json:"propertyNames"`
  ReservedConst interface{} `json:"const"`
  Enum []interface{} `json:"enum"`
  ReservedType *JsonMinusSchemaMinusDraftMinus_07MinusSchemaType `json:"type"`
  Format string `json:"format"`
  ContentMediaType string `json:"contentMediaType"`
  ContentEncoding string `json:"contentEncoding"`
  ReservedIf *MessageOneOf_1OneOf_1Payload `json:"if"`
  Then *MessageOneOf_1OneOf_1Payload `json:"then"`
  ReservedElse *MessageOneOf_1OneOf_1Payload `json:"else"`
  AllOf []MessageOneOf_1OneOf_1Payload `json:"allOf"`
  AnyOf []MessageOneOf_1OneOf_1Payload `json:"anyOf"`
  OneOf []MessageOneOf_1OneOf_1Payload `json:"oneOf"`
  Not *MessageOneOf_1OneOf_1Payload `json:"not"`
  Discriminator *SchemaAllOf_1Discriminator `json:"discriminator"`
  ExternalDocs *ExternalDocs `json:"externalDocs"`
  Deprecated bool `json:"deprecated"`
  Nullable bool `json:"nullable"`
  Example interface{} `json:"example"`
  Xml *OpenapiSchema_3_0Xml `json:"xml"`
  ReservedAdditionalProperties map[string]ReservedAdditionalProperties `json:"-"`
}