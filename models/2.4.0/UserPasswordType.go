
package models

// UserPasswordType represents an enum of UserPasswordType.
type UserPasswordType uint

const (
  UserPasswordTypeUserPassword UserPasswordType = iota
)

// Value returns the value of the enum.
func (op UserPasswordType) Value() any {
	if op >= UserPasswordType(len(UserPasswordTypeValues)) {
		return nil
	}
	return UserPasswordTypeValues[op]
}

var UserPasswordTypeValues = []any{"userPassword"}
var ValuesToUserPasswordType = map[any]UserPasswordType{
  UserPasswordTypeValues[UserPasswordTypeUserPassword]: UserPasswordTypeUserPassword,
}
