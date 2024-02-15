
package models

// OneOf_1 represents a OneOf_1 model.
type OneOf_1 struct {
  Asyncapi *Asyncapi
  Info *Info
  BaseTopic string
  Servers []Server
  Topics *Topics
  Stream *StreamObject
  Events interface{}
  Components *Components
  Tags []Tag
  Security []map[string][]string
  ExternalDocs *ExternalDocs
  AdditionalProperties map[string]interface{}
}