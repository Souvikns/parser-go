
package models
import (  
  "encoding/json"
)
// JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes represents an enum of JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes.
type JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes uint

const (
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesArray JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes = iota
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesBoolean
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesInteger
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesNull
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesNumber
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesObject
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesString
)

// Value returns the value of the enum.
func (op JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes) Value() any {
	if op >= JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes(len(JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues)) {
		return nil
	}
	return JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues[op]
}

var JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues = []any{"array","boolean","integer","null","number","object","string"}
var ValuesToJsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes = map[any]JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes{
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues[JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesArray]: JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesArray,
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues[JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesBoolean]: JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesBoolean,
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues[JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesInteger]: JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesInteger,
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues[JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesNull]: JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesNull,
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues[JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesNumber]: JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesNumber,
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues[JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesObject]: JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesObject,
  JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesValues[JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesString]: JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypesString,
}

 
          
func (op *JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToJsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes[v]
	return nil
}

func (op JsonMinusSchemaMinusDraftMinus_07MinusSchemaSimpleTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          