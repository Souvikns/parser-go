
package models
import (  
  "encoding/json"
)
// StreamFramingOneOf_0Type represents an enum of StreamFramingOneOf_0Type.
type StreamFramingOneOf_0Type uint

const (
  StreamFramingOneOf_0TypeChunked StreamFramingOneOf_0Type = iota
)

// Value returns the value of the enum.
func (op StreamFramingOneOf_0Type) Value() any {
	if op >= StreamFramingOneOf_0Type(len(StreamFramingOneOf_0TypeValues)) {
		return nil
	}
	return StreamFramingOneOf_0TypeValues[op]
}

var StreamFramingOneOf_0TypeValues = []any{"chunked"}
var ValuesToStreamFramingOneOf_0Type = map[any]StreamFramingOneOf_0Type{
  StreamFramingOneOf_0TypeValues[StreamFramingOneOf_0TypeChunked]: StreamFramingOneOf_0TypeChunked,
}

 
func (op *StreamFramingOneOf_0Type) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToStreamFramingOneOf_0Type[v]
  return nil
}

func (op StreamFramingOneOf_0Type) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}