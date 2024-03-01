
package models

// Reference represents a Reference model.
type Reference struct {
  Ref string `json:$ref`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}