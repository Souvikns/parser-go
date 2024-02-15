
package models

// Asyncapi represents an enum of Asyncapi.
type Asyncapi uint

const (
  AsyncapiNumber_1Dot_0Dot_0 Asyncapi = iota
)

// Value returns the value of the enum.
func (op Asyncapi) Value() any {
	if op >= Asyncapi(len(AsyncapiValues)) {
		return nil
	}
	return AsyncapiValues[op]
}

var AsyncapiValues = []any{"1.0.0"}
var ValuesToAsyncapi = map[any]Asyncapi{
  AsyncapiValues[AsyncapiNumber_1Dot_0Dot_0]: AsyncapiNumber_1Dot_0Dot_0,
}
