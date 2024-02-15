
package models

// OneOf_0 represents a OneOf_0 model.
type OneOf_0 struct {
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