
package models

// StreamFramingOneOf_1Delimiter represents an enum of StreamFramingOneOf_1Delimiter.
type StreamFramingOneOf_1Delimiter uint

const (
  StreamFramingOneOf_1DelimiterBackslashNBackslashN StreamFramingOneOf_1Delimiter = iota
)

// Value returns the value of the enum.
func (op StreamFramingOneOf_1Delimiter) Value() any {
	if op >= StreamFramingOneOf_1Delimiter(len(StreamFramingOneOf_1DelimiterValues)) {
		return nil
	}
	return StreamFramingOneOf_1DelimiterValues[op]
}

var StreamFramingOneOf_1DelimiterValues = []any{"\n\n"}
var ValuesToStreamFramingOneOf_1Delimiter = map[any]StreamFramingOneOf_1Delimiter{
  StreamFramingOneOf_1DelimiterValues[StreamFramingOneOf_1DelimiterBackslashNBackslashN]: StreamFramingOneOf_1DelimiterBackslashNBackslashN,
}
