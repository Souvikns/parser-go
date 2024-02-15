
package models

// StreamFramingOneOf_0Delimiter represents an enum of StreamFramingOneOf_0Delimiter.
type StreamFramingOneOf_0Delimiter uint

const (
  StreamFramingOneOf_0DelimiterBackslashRBackslashN StreamFramingOneOf_0Delimiter = iota
  StreamFramingOneOf_0DelimiterBackslashN
)

// Value returns the value of the enum.
func (op StreamFramingOneOf_0Delimiter) Value() any {
	if op >= StreamFramingOneOf_0Delimiter(len(StreamFramingOneOf_0DelimiterValues)) {
		return nil
	}
	return StreamFramingOneOf_0DelimiterValues[op]
}

var StreamFramingOneOf_0DelimiterValues = []any{"\r\n","\n"}
var ValuesToStreamFramingOneOf_0Delimiter = map[any]StreamFramingOneOf_0Delimiter{
  StreamFramingOneOf_0DelimiterValues[StreamFramingOneOf_0DelimiterBackslashRBackslashN]: StreamFramingOneOf_0DelimiterBackslashRBackslashN,
  StreamFramingOneOf_0DelimiterValues[StreamFramingOneOf_0DelimiterBackslashN]: StreamFramingOneOf_0DelimiterBackslashN,
}
