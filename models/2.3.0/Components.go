
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]interface{} `json:schemas`
  Servers map[string]interface{} `json:servers`
  Channels map[string]ChannelItem `json:channels`
  Messages map[string]interface{} `json:messages`
  SecuritySchemes map[string]interface{} `json:securitySchemes`
  Parameters map[string]interface{} `json:parameters`
  CorrelationIds map[string]interface{} `json:correlationIds`
  OperationTraits map[string]OperationTrait `json:operationTraits`
  MessageTraits map[string]MessageTrait `json:messageTraits`
  ServerBindings map[string]BindingsObject `json:serverBindings`
  ChannelBindings map[string]BindingsObject `json:channelBindings`
  OperationBindings map[string]BindingsObject `json:operationBindings`
  MessageBindings map[string]BindingsObject `json:messageBindings`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}