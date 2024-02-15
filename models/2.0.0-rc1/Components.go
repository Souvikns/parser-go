
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]Schema
  Messages map[string]Message
  SecuritySchemes map[string]interface{}
  Parameters map[string]interface{}
  CorrelationIds map[string]interface{}
  Traits map[string]interface{}
}