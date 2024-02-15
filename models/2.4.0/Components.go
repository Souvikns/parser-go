
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]interface{}
  Servers map[string]interface{}
  Channels map[string]ChannelItem
  ServerVariables map[string]ServerVariable
  Messages map[string]interface{}
  SecuritySchemes map[string]interface{}
  Parameters map[string]interface{}
  CorrelationIds map[string]interface{}
  OperationTraits map[string]OperationTrait
  MessageTraits map[string]MessageTrait
  ServerBindings map[string]BindingsObject
  ChannelBindings map[string]BindingsObject
  OperationBindings map[string]BindingsObject
  MessageBindings map[string]BindingsObject
  AdditionalProperties map[string]interface{}
}