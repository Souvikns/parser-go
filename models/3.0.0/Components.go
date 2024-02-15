
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]interface{}
  Servers map[string]interface{}
  Channels map[string]interface{}
  ServerVariables map[string]interface{}
  Operations map[string]interface{}
  Messages map[string]interface{}
  SecuritySchemes map[string]interface{}
  Parameters map[string]interface{}
  CorrelationIds map[string]interface{}
  OperationTraits map[string]interface{}
  MessageTraits map[string]interface{}
  Replies map[string]interface{}
  ReplyAddresses map[string]interface{}
  ServerBindings map[string]interface{}
  ChannelBindings map[string]interface{}
  OperationBindings map[string]interface{}
  MessageBindings map[string]interface{}
  Tags map[string]interface{}
  ExternalDocs map[string]interface{}
  AdditionalProperties map[string]interface{}
}