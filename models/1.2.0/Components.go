
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]Schema `json:schemas`
  Messages map[string]Message `json:messages`
  SecuritySchemes map[string]interface{} `json:securitySchemes`
  Parameters map[string]Parameter `json:parameters`
}