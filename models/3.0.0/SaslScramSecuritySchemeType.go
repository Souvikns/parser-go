
package models

// SaslScramSecuritySchemeType represents an enum of SaslScramSecuritySchemeType.
type SaslScramSecuritySchemeType uint

const (
  SaslScramSecuritySchemeTypeScramSha256 SaslScramSecuritySchemeType = iota
  SaslScramSecuritySchemeTypeScramSha512
)

// Value returns the value of the enum.
func (op SaslScramSecuritySchemeType) Value() any {
	if op >= SaslScramSecuritySchemeType(len(SaslScramSecuritySchemeTypeValues)) {
		return nil
	}
	return SaslScramSecuritySchemeTypeValues[op]
}

var SaslScramSecuritySchemeTypeValues = []any{"scramSha256","scramSha512"}
var ValuesToSaslScramSecuritySchemeType = map[any]SaslScramSecuritySchemeType{
  SaslScramSecuritySchemeTypeValues[SaslScramSecuritySchemeTypeScramSha256]: SaslScramSecuritySchemeTypeScramSha256,
  SaslScramSecuritySchemeTypeValues[SaslScramSecuritySchemeTypeScramSha512]: SaslScramSecuritySchemeTypeScramSha512,
}
