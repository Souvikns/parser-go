
package models

// UserPassword represents a UserPassword model.
type UserPassword struct {
  ReservedType *UserPasswordType `json:"type"`
  Description string `json:"description"`
  AdditionalProperties map[string]interface{} `json:"-"`
}