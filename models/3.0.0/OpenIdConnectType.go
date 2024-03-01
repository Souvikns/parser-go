
package models
import (  
  "encoding/json"
)
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

 
          
func (op *OpenIdConnectType) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToOpenIdConnectType[v]
	return nil
}

func (op OpenIdConnectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          