
package models

// Asyncapi represents an enum of Asyncapi.
type Asyncapi uint

const (
  AsyncapiNumber_2Dot_0Dot_0MinusRc2 Asyncapi = iota
)

// Value returns the value of the enum.
func (op Asyncapi) Value() any {
	if op >= Asyncapi(len(AsyncapiValues)) {
		return nil
	}
	return AsyncapiValues[op]
}

var AsyncapiValues = []any{"2.0.0-rc2"}
var ValuesToAsyncapi = map[any]Asyncapi{
  AsyncapiValues[AsyncapiNumber_2Dot_0Dot_0MinusRc2]: AsyncapiNumber_2Dot_0Dot_0MinusRc2,
}
