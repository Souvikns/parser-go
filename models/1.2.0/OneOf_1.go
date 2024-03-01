
package models

// OneOf_1 represents a OneOf_1 model.
type OneOf_1 struct {
  Asyncapi *Asyncapi `json:asyncapi`
  Info *Info `json:info`
  BaseTopic string `json:baseTopic`
  Servers []Server `json:servers`
  Topics *Topics `json:topics`
  Stream *StreamObject `json:stream`
  Events interface{} `json:events`
  Components *Components `json:components`
  Tags []Tag `json:tags`
  Security []map[string][]string `json:security`
  ExternalDocs *ExternalDocs `json:externalDocs`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}