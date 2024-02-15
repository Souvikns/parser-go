
package models

// BearerHttpSecuritySchemeType represents an enum of BearerHttpSecuritySchemeType.
type BearerHttpSecuritySchemeType uint

const (
  BearerHttpSecuritySchemeTypeHttp BearerHttpSecuritySchemeType = iota
)

// Value returns the value of the enum.
func (op BearerHttpSecuritySchemeType) Value() any {
	if op >= BearerHttpSecuritySchemeType(len(BearerHttpSecuritySchemeTypeValues)) {
		return nil
	}
	return BearerHttpSecuritySchemeTypeValues[op]
}

var BearerHttpSecuritySchemeTypeValues = []any{"http"}
var ValuesToBearerHttpSecuritySchemeType = map[any]BearerHttpSecuritySchemeType{
  BearerHttpSecuritySchemeTypeValues[BearerHttpSecuritySchemeTypeHttp]: BearerHttpSecuritySchemeTypeHttp,
}
