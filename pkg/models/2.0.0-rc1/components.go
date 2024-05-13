
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]Schema `json:"schemas"`
  Messages map[string]Message `json:"messages"`
  SecuritySchemes map[string]ComponentsSecuritySchemes `json:"securitySchemes"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters"`
  CorrelationIds map[string]ComponentsCorrelationIds `json:"correlationIds"`
  Traits map[string]TraitsAdditionalProperty `json:"traits"`
}