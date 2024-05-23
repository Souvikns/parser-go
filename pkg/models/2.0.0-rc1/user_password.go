
package models

// UserPassword represents a UserPassword model.
type UserPassword struct {
  ReservedType *UserPasswordType `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}