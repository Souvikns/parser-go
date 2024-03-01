
package models
import (  
  "encoding/json"
)
// MultiFormatSchemaElseSchemaFormatAnyOf_1 represents an enum of MultiFormatSchemaElseSchemaFormatAnyOf_1.
type MultiFormatSchemaElseSchemaFormatAnyOf_1 uint

const (
  MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashSchemaPlusJsonSemicolonVersionEqualDraftMinus_07 MultiFormatSchemaElseSchemaFormatAnyOf_1 = iota
  MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashSchemaPlusYamlSemicolonVersionEqualDraftMinus_07
  MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashVndDotAaiDotAsyncapiSemicolonVersionEqual_3Dot_0Dot_0
  MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashVndDotAaiDotAsyncapiPlusJsonSemicolonVersionEqual_3Dot_0Dot_0
  MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashVndDotAaiDotAsyncapiPlusYamlSemicolonVersionEqual_3Dot_0Dot_0
)

// Value returns the value of the enum.
func (op MultiFormatSchemaElseSchemaFormatAnyOf_1) Value() any {
	if op >= MultiFormatSchemaElseSchemaFormatAnyOf_1(len(MultiFormatSchemaElseSchemaFormatAnyOf_1Values)) {
		return nil
	}
	return MultiFormatSchemaElseSchemaFormatAnyOf_1Values[op]
}

var MultiFormatSchemaElseSchemaFormatAnyOf_1Values = []any{"application/schema+json;version=draft-07","application/schema+yaml;version=draft-07","application/vnd.aai.asyncapi;version=3.0.0","application/vnd.aai.asyncapi+json;version=3.0.0","application/vnd.aai.asyncapi+yaml;version=3.0.0"}
var ValuesToMultiFormatSchemaElseSchemaFormatAnyOf_1 = map[any]MultiFormatSchemaElseSchemaFormatAnyOf_1{
  MultiFormatSchemaElseSchemaFormatAnyOf_1Values[MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashSchemaPlusJsonSemicolonVersionEqualDraftMinus_07]: MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashSchemaPlusJsonSemicolonVersionEqualDraftMinus_07,
  MultiFormatSchemaElseSchemaFormatAnyOf_1Values[MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashSchemaPlusYamlSemicolonVersionEqualDraftMinus_07]: MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashSchemaPlusYamlSemicolonVersionEqualDraftMinus_07,
  MultiFormatSchemaElseSchemaFormatAnyOf_1Values[MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashVndDotAaiDotAsyncapiSemicolonVersionEqual_3Dot_0Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashVndDotAaiDotAsyncapiSemicolonVersionEqual_3Dot_0Dot_0,
  MultiFormatSchemaElseSchemaFormatAnyOf_1Values[MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashVndDotAaiDotAsyncapiPlusJsonSemicolonVersionEqual_3Dot_0Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashVndDotAaiDotAsyncapiPlusJsonSemicolonVersionEqual_3Dot_0Dot_0,
  MultiFormatSchemaElseSchemaFormatAnyOf_1Values[MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashVndDotAaiDotAsyncapiPlusYamlSemicolonVersionEqual_3Dot_0Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_1ApplicationSlashVndDotAaiDotAsyncapiPlusYamlSemicolonVersionEqual_3Dot_0Dot_0,
}

 
          
func (op *MultiFormatSchemaElseSchemaFormatAnyOf_1) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToMultiFormatSchemaElseSchemaFormatAnyOf_1[v]
	return nil
}

func (op MultiFormatSchemaElseSchemaFormatAnyOf_1) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          