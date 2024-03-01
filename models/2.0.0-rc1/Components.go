
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]Schema `json:schemas`
  Messages map[string]Message `json:messages`
  SecuritySchemes map[string]interface{} `json:securitySchemes`
  Parameters map[string]interface{} `json:parameters`
  CorrelationIds map[string]interface{} `json:correlationIds`
  Traits map[string]interface{} `json:traits`
}