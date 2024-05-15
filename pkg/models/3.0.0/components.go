
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]ComponentsSchemas `json:"schemas"`
  Servers map[string]ComponentsServers `json:"servers"`
  Channels map[string]ComponentsChannels `json:"channels"`
  ServerVariables map[string]ComponentsServerVariables `json:"serverVariables"`
  Operations map[string]ComponentsOperations `json:"operations"`
  Messages map[string]ComponentsMessages `json:"messages"`
  SecuritySchemes map[string]ComponentsSecuritySchemes `json:"securitySchemes"`
  Parameters map[string]ComponentsParameters `json:"parameters"`
  CorrelationIds map[string]ComponentsCorrelationIds `json:"correlationIds"`
  OperationTraits map[string]ComponentsOperationTraits `json:"operationTraits"`
  MessageTraits map[string]ComponentsMessageTraits `json:"messageTraits"`
  Replies map[string]ComponentsReplies `json:"replies"`
  ReplyAddresses map[string]ComponentsReplyAddresses `json:"replyAddresses"`
  ServerBindings map[string]ComponentsServerBindings `json:"serverBindings"`
  ChannelBindings map[string]ComponentsChannelBindings `json:"channelBindings"`
  OperationBindings map[string]ComponentsOperationBindings `json:"operationBindings"`
  MessageBindings map[string]ComponentsMessageBindings `json:"messageBindings"`
  Tags map[string]ComponentsTags `json:"tags"`
  ExternalDocs map[string]ComponentsExternalDocs `json:"externalDocs"`
  AdditionalProperties map[string]interface{} `json:"-"`
}