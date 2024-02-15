
package models

// BearerHttpSecuritySchemeScheme represents an enum of BearerHttpSecuritySchemeScheme.
type BearerHttpSecuritySchemeScheme uint

const (
  BearerHttpSecuritySchemeSchemeBearer BearerHttpSecuritySchemeScheme = iota
)

// Value returns the value of the enum.
func (op BearerHttpSecuritySchemeScheme) Value() any {
	if op >= BearerHttpSecuritySchemeScheme(len(BearerHttpSecuritySchemeSchemeValues)) {
		return nil
	}
	return BearerHttpSecuritySchemeSchemeValues[op]
}

var BearerHttpSecuritySchemeSchemeValues = []any{"bearer"}
var ValuesToBearerHttpSecuritySchemeScheme = map[any]BearerHttpSecuritySchemeScheme{
  BearerHttpSecuritySchemeSchemeValues[BearerHttpSecuritySchemeSchemeBearer]: BearerHttpSecuritySchemeSchemeBearer,
}
