
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]interface{} `json:"schemas"`
  Servers map[string]interface{} `json:"servers"`
  Channels map[string]interface{} `json:"channels"`
  ServerVariables map[string]interface{} `json:"serverVariables"`
  Operations map[string]interface{} `json:"operations"`
  Messages map[string]interface{} `json:"messages"`
  SecuritySchemes map[string]interface{} `json:"securitySchemes"`
  Parameters map[string]interface{} `json:"parameters"`
  CorrelationIds map[string]interface{} `json:"correlationIds"`
  OperationTraits map[string]interface{} `json:"operationTraits"`
  MessageTraits map[string]interface{} `json:"messageTraits"`
  Replies map[string]interface{} `json:"replies"`
  ReplyAddresses map[string]interface{} `json:"replyAddresses"`
  ServerBindings map[string]interface{} `json:"serverBindings"`
  ChannelBindings map[string]interface{} `json:"channelBindings"`
  OperationBindings map[string]interface{} `json:"operationBindings"`
  MessageBindings map[string]interface{} `json:"messageBindings"`
  Tags map[string]interface{} `json:"tags"`
  ExternalDocs map[string]interface{} `json:"externalDocs"`
  AdditionalProperties map[string]interface{} `json:"-"`
}