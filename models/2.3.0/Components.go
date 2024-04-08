
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]Schema `json:"schemas"`
  Servers map[string]ServersAdditionalProperty `json:"servers"`
  Channels map[string]ChannelItem `json:"channels"`
  Messages map[string]Message `json:"messages"`
  SecuritySchemes map[string]ComponentsSecuritySchemes `json:"securitySchemes"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters"`
  CorrelationIds map[string]ComponentsCorrelationIds `json:"correlationIds"`
  OperationTraits map[string]OperationTrait `json:"operationTraits"`
  MessageTraits map[string]MessageTrait `json:"messageTraits"`
  ServerBindings map[string]BindingsObject `json:"serverBindings"`
  ChannelBindings map[string]BindingsObject `json:"channelBindings"`
  OperationBindings map[string]BindingsObject `json:"operationBindings"`
  MessageBindings map[string]BindingsObject `json:"messageBindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}