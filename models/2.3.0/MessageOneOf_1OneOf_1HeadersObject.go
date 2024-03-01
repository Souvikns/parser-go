
package models

// MessageOneOf_1OneOf_1HeadersObject represents a MessageOneOf_1OneOf_1HeadersObject model.
type MessageOneOf_1OneOf_1HeadersObject struct {
  Id string `json:$id`
  Schema string `json:$schema`
  Ref string `json:$ref`
  Comment string `json:$comment`
  Title string `json:title`
  Description string `json:description`
  ReservedDefault interface{} `json:default`
  ReadOnly bool `json:readOnly`
  WriteOnly bool `json:writeOnly`
  Examples []interface{} `json:examples`
  MultipleOf float64 `json:multipleOf`
  Maximum float64 `json:maximum`
  ExclusiveMaximum interface{} `json:exclusiveMaximum`
  Minimum float64 `json:minimum`
  ExclusiveMinimum interface{} `json:exclusiveMinimum`
  MaxLength int `json:maxLength`
  MinLength int `json:minLength`
  Pattern string `json:pattern`
  AdditionalItems interface{} `json:additionalItems`
  Items interface{} `json:items`
  MaxItems int `json:maxItems`
  MinItems int `json:minItems`
  UniqueItems bool `json:uniqueItems`
  Contains interface{} `json:contains`
  MaxProperties int `json:maxProperties`
  MinProperties int `json:minProperties`
  Required []string `json:required`
  AdditionalProperties interface{} `json:additionalProperties`
  Definitions map[string]interface{} `json:definitions`
  Properties map[string]interface{} `json:properties`
  PatternProperties map[string]interface{} `json:patternProperties`
  Dependencies map[string]interface{} `json:dependencies`
  PropertyNames interface{} `json:propertyNames`
  ReservedConst interface{} `json:const`
  Enum []interface{} `json:enum`
  ReservedType interface{} `json:type`
  Format string `json:format`
  ContentMediaType string `json:contentMediaType`
  ContentEncoding string `json:contentEncoding`
  ReservedIf interface{} `json:if`
  Then interface{} `json:then`
  ReservedElse interface{} `json:else`
  AllOf []interface{} `json:allOf`
  AnyOf []interface{} `json:anyOf`
  OneOf []interface{} `json:oneOf`
  Not interface{} `json:not`
  Discriminator interface{} `json:discriminator`
  ExternalDocs *ExternalDocs `json:externalDocs`
  Deprecated bool `json:deprecated`
  ReservedAdditionalProperties map[string]interface{} `json:reserved_additionalProperties`
}