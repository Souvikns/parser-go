
package models
import (  
  "encoding/json"
)
// AvroSchemaV1AvroFieldOrder represents an enum of AvroSchemaV1AvroFieldOrder.
type AvroSchemaV1AvroFieldOrder uint

const (
  AvroSchemaV1AvroFieldOrderAscending AvroSchemaV1AvroFieldOrder = iota
  AvroSchemaV1AvroFieldOrderDescending
  AvroSchemaV1AvroFieldOrderIgnore
)

// Value returns the value of the enum.
func (op AvroSchemaV1AvroFieldOrder) Value() any {
	if op >= AvroSchemaV1AvroFieldOrder(len(AvroSchemaV1AvroFieldOrderValues)) {
		return nil
	}
	return AvroSchemaV1AvroFieldOrderValues[op]
}

var AvroSchemaV1AvroFieldOrderValues = []any{"ascending","descending","ignore"}
var ValuesToAvroSchemaV1AvroFieldOrder = map[any]AvroSchemaV1AvroFieldOrder{
  AvroSchemaV1AvroFieldOrderValues[AvroSchemaV1AvroFieldOrderAscending]: AvroSchemaV1AvroFieldOrderAscending,
  AvroSchemaV1AvroFieldOrderValues[AvroSchemaV1AvroFieldOrderDescending]: AvroSchemaV1AvroFieldOrderDescending,
  AvroSchemaV1AvroFieldOrderValues[AvroSchemaV1AvroFieldOrderIgnore]: AvroSchemaV1AvroFieldOrderIgnore,
}

 
func (op *AvroSchemaV1AvroFieldOrder) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToAvroSchemaV1AvroFieldOrder[v]
  return nil
}

func (op AvroSchemaV1AvroFieldOrder) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}