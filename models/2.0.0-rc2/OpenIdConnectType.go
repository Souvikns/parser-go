
package models

// OpenIdConnectType represents an enum of OpenIdConnectType.
type OpenIdConnectType uint

const (
  OpenIdConnectTypeOpenIdConnect OpenIdConnectType = iota
)

// Value returns the value of the enum.
func (op OpenIdConnectType) Value() any {
	if op >= OpenIdConnectType(len(OpenIdConnectTypeValues)) {
		return nil
	}
	return OpenIdConnectTypeValues[op]
}

var OpenIdConnectTypeValues = []any{"openIdConnect"}
var ValuesToOpenIdConnectType = map[any]OpenIdConnectType{
  OpenIdConnectTypeValues[OpenIdConnectTypeOpenIdConnect]: OpenIdConnectTypeOpenIdConnect,
}
