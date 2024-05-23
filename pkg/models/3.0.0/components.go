
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]ComponentsSchemas `json:"schemas,omitempty"`
  Servers map[string]ComponentsServers `json:"servers,omitempty"`
  Channels map[string]ComponentsChannels `json:"channels,omitempty"`
  ServerVariables map[string]ComponentsServerVariables `json:"serverVariables,omitempty"`
  Operations map[string]ComponentsOperations `json:"operations,omitempty"`
  Messages map[string]ComponentsMessages `json:"messages,omitempty"`
  SecuritySchemes map[string]ComponentsSecuritySchemes `json:"securitySchemes,omitempty"`
  Parameters map[string]ComponentsParameters `json:"parameters,omitempty"`
  CorrelationIds map[string]ComponentsCorrelationIds `json:"correlationIds,omitempty"`
  OperationTraits map[string]ComponentsOperationTraits `json:"operationTraits,omitempty"`
  MessageTraits map[string]ComponentsMessageTraits `json:"messageTraits,omitempty"`
  Replies map[string]ComponentsReplies `json:"replies,omitempty"`
  ReplyAddresses map[string]ComponentsReplyAddresses `json:"replyAddresses,omitempty"`
  ServerBindings map[string]ComponentsServerBindings `json:"serverBindings,omitempty"`
  ChannelBindings map[string]ComponentsChannelBindings `json:"channelBindings,omitempty"`
  OperationBindings map[string]ComponentsOperationBindings `json:"operationBindings,omitempty"`
  MessageBindings map[string]ComponentsMessageBindings `json:"messageBindings,omitempty"`
  Tags map[string]ComponentsTags `json:"tags,omitempty"`
  ExternalDocs map[string]ComponentsExternalDocs `json:"externalDocs,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}