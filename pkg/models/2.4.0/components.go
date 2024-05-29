
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]Schema `json:"schemas,omitempty"`
  Servers map[string]ServersAdditionalProperty `json:"servers,omitempty"`
  Channels map[string]ChannelItem `json:"channels,omitempty"`
  ServerVariables map[string]ServerVariable `json:"serverVariables,omitempty"`
  Messages map[string]Message `json:"messages,omitempty"`
  SecuritySchemes map[string]ComponentsSecuritySchemes `json:"securitySchemes,omitempty"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters,omitempty"`
  CorrelationIds map[string]ComponentsCorrelationIds `json:"correlationIds,omitempty"`
  OperationTraits map[string]OperationTrait `json:"operationTraits,omitempty"`
  MessageTraits map[string]MessageTrait `json:"messageTraits,omitempty"`
  ServerBindings map[string]BindingsObject `json:"serverBindings,omitempty"`
  ChannelBindings map[string]BindingsObject `json:"channelBindings,omitempty"`
  OperationBindings map[string]BindingsObject `json:"operationBindings,omitempty"`
  MessageBindings map[string]BindingsObject `json:"messageBindings,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}