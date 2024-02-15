
package models

// SaslGssapiSecuritySchemeType represents an enum of SaslGssapiSecuritySchemeType.
type SaslGssapiSecuritySchemeType uint

const (
  SaslGssapiSecuritySchemeTypeGssapi SaslGssapiSecuritySchemeType = iota
)

// Value returns the value of the enum.
func (op SaslGssapiSecuritySchemeType) Value() any {
	if op >= SaslGssapiSecuritySchemeType(len(SaslGssapiSecuritySchemeTypeValues)) {
		return nil
	}
	return SaslGssapiSecuritySchemeTypeValues[op]
}

var SaslGssapiSecuritySchemeTypeValues = []any{"gssapi"}
var ValuesToSaslGssapiSecuritySchemeType = map[any]SaslGssapiSecuritySchemeType{
  SaslGssapiSecuritySchemeTypeValues[SaslGssapiSecuritySchemeTypeGssapi]: SaslGssapiSecuritySchemeTypeGssapi,
}
