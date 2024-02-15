
package models

// SaslPlainSecuritySchemeType represents an enum of SaslPlainSecuritySchemeType.
type SaslPlainSecuritySchemeType uint

const (
  SaslPlainSecuritySchemeTypePlain SaslPlainSecuritySchemeType = iota
)

// Value returns the value of the enum.
func (op SaslPlainSecuritySchemeType) Value() any {
	if op >= SaslPlainSecuritySchemeType(len(SaslPlainSecuritySchemeTypeValues)) {
		return nil
	}
	return SaslPlainSecuritySchemeTypeValues[op]
}

var SaslPlainSecuritySchemeTypeValues = []any{"plain"}
var ValuesToSaslPlainSecuritySchemeType = map[any]SaslPlainSecuritySchemeType{
  SaslPlainSecuritySchemeTypeValues[SaslPlainSecuritySchemeTypePlain]: SaslPlainSecuritySchemeTypePlain,
}
