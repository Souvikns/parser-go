
package models

// SchemaObject represents a SchemaObject model.
type SchemaObject struct {
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
  ExclusiveMaximum float64 `json:"exclusiveMaximum"`
  Minimum float64 `json:"minimum"`
  ExclusiveMinimum float64 `json:"exclusiveMinimum"`
  MaxLength int `json:"maxLength"`
  MinLength int `json:"minLength"`
  Pattern string `json:"pattern"`
  AdditionalItems *CoreSchemaMetaMinusSchema `json:"additionalItems"`
  Items *JsonMinusSchemaMinusDraftMinus_07MinusSchemaItems `json:"items"`
  MaxItems int `json:"maxItems"`
  MinItems int `json:"minItems"`
  UniqueItems bool `json:"uniqueItems"`
  Contains *AnySchema `json:"contains"`
  MaxProperties int `json:"maxProperties"`
  MinProperties int `json:"minProperties"`
  Required []string `json:"required"`
  AdditionalProperties *CoreSchemaMetaMinusSchema `json:"additionalProperties"`
  Definitions map[string]AnySchema `json:"definitions"`
  Properties map[string]AnySchema `json:"properties"`
  PatternProperties map[string]AnySchema `json:"patternProperties"`
  Dependencies map[string]JsonMinusSchemaMinusDraftMinus_07MinusSchemaDependenciesAdditionalProperty `json:"dependencies"`
  PropertyNames *AnySchema `json:"propertyNames"`
  ReservedConst interface{} `json:"const"`
  Enum []interface{} `json:"enum"`
  ReservedType *JsonMinusSchemaMinusDraftMinus_07MinusSchemaType `json:"type"`
  Format string `json:"format"`
  ContentMediaType string `json:"contentMediaType"`
  ContentEncoding string `json:"contentEncoding"`
  ReservedIf *CoreSchemaMetaMinusSchema `json:"if"`
  Then *CoreSchemaMetaMinusSchema `json:"then"`
  ReservedElse *CoreSchemaMetaMinusSchema `json:"else"`
  AllOf []AnySchema `json:"allOf"`
  AnyOf []AnySchema `json:"anyOf"`
  OneOf []AnySchema `json:"oneOf"`
  Not *AnySchema `json:"not"`
  Discriminator string `json:"discriminator"`
  ExternalDocs *SchemaAllOf_1ExternalDocs `json:"externalDocs"`
  Deprecated bool `json:"deprecated"`
  ReservedAdditionalProperties map[string]ReservedAdditionalProperties `json:"-"`
}