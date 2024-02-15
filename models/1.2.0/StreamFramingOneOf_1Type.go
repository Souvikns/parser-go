
package models

// StreamFramingOneOf_1Type represents an enum of StreamFramingOneOf_1Type.
type StreamFramingOneOf_1Type uint

const (
  StreamFramingOneOf_1TypeSse StreamFramingOneOf_1Type = iota
)

// Value returns the value of the enum.
func (op StreamFramingOneOf_1Type) Value() any {
	if op >= StreamFramingOneOf_1Type(len(StreamFramingOneOf_1TypeValues)) {
		return nil
	}
	return StreamFramingOneOf_1TypeValues[op]
}

var StreamFramingOneOf_1TypeValues = []any{"sse"}
var ValuesToStreamFramingOneOf_1Type = map[any]StreamFramingOneOf_1Type{
  StreamFramingOneOf_1TypeValues[StreamFramingOneOf_1TypeSse]: StreamFramingOneOf_1TypeSse,
}
