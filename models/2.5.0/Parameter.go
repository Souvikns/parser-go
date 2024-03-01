
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string `json:description`
  Schema interface{} `json:schema`
  Location string `json:location`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}