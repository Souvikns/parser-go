
package models

// Components represents a Components model.
type Components struct {
  Schemas map[string]Schema `json:"schemas,omitempty"`
  Messages map[string]Message `json:"messages,omitempty"`
  SecuritySchemes map[string]ComponentsSecuritySchemes `json:"securitySchemes,omitempty"`
  Parameters map[string]ParametersAdditionalProperty `json:"parameters,omitempty"`
  CorrelationIds map[string]ComponentsCorrelationIds `json:"correlationIds,omitempty"`
  Traits map[string]TraitsAdditionalProperty `json:"traits,omitempty"`
}