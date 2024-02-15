
package models

// MultiFormatSchemaElseSchemaFormatAnyOf_2 represents an enum of MultiFormatSchemaElseSchemaFormatAnyOf_2.
type MultiFormatSchemaElseSchemaFormatAnyOf_2 uint

const (
  MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotOaiDotOpenapiSemicolonVersionEqual_3Dot_0Dot_0 MultiFormatSchemaElseSchemaFormatAnyOf_2 = iota
  MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotOaiDotOpenapiPlusJsonSemicolonVersionEqual_3Dot_0Dot_0
  MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotOaiDotOpenapiPlusYamlSemicolonVersionEqual_3Dot_0Dot_0
  MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotApacheDotAvroSemicolonVersionEqual_1Dot_9Dot_0
  MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotApacheDotAvroPlusJsonSemicolonVersionEqual_1Dot_9Dot_0
  MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotApacheDotAvroPlusYamlSemicolonVersionEqual_1Dot_9Dot_0
  MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashRamlPlusYamlSemicolonVersionEqual_1Dot_0
)

// Value returns the value of the enum.
func (op MultiFormatSchemaElseSchemaFormatAnyOf_2) Value() any {
	if op >= MultiFormatSchemaElseSchemaFormatAnyOf_2(len(MultiFormatSchemaElseSchemaFormatAnyOf_2Values)) {
		return nil
	}
	return MultiFormatSchemaElseSchemaFormatAnyOf_2Values[op]
}

var MultiFormatSchemaElseSchemaFormatAnyOf_2Values = []any{"application/vnd.oai.openapi;version=3.0.0","application/vnd.oai.openapi+json;version=3.0.0","application/vnd.oai.openapi+yaml;version=3.0.0","application/vnd.apache.avro;version=1.9.0","application/vnd.apache.avro+json;version=1.9.0","application/vnd.apache.avro+yaml;version=1.9.0","application/raml+yaml;version=1.0"}
var ValuesToMultiFormatSchemaElseSchemaFormatAnyOf_2 = map[any]MultiFormatSchemaElseSchemaFormatAnyOf_2{
  MultiFormatSchemaElseSchemaFormatAnyOf_2Values[MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotOaiDotOpenapiSemicolonVersionEqual_3Dot_0Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotOaiDotOpenapiSemicolonVersionEqual_3Dot_0Dot_0,
  MultiFormatSchemaElseSchemaFormatAnyOf_2Values[MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotOaiDotOpenapiPlusJsonSemicolonVersionEqual_3Dot_0Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotOaiDotOpenapiPlusJsonSemicolonVersionEqual_3Dot_0Dot_0,
  MultiFormatSchemaElseSchemaFormatAnyOf_2Values[MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotOaiDotOpenapiPlusYamlSemicolonVersionEqual_3Dot_0Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotOaiDotOpenapiPlusYamlSemicolonVersionEqual_3Dot_0Dot_0,
  MultiFormatSchemaElseSchemaFormatAnyOf_2Values[MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotApacheDotAvroSemicolonVersionEqual_1Dot_9Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotApacheDotAvroSemicolonVersionEqual_1Dot_9Dot_0,
  MultiFormatSchemaElseSchemaFormatAnyOf_2Values[MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotApacheDotAvroPlusJsonSemicolonVersionEqual_1Dot_9Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotApacheDotAvroPlusJsonSemicolonVersionEqual_1Dot_9Dot_0,
  MultiFormatSchemaElseSchemaFormatAnyOf_2Values[MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotApacheDotAvroPlusYamlSemicolonVersionEqual_1Dot_9Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashVndDotApacheDotAvroPlusYamlSemicolonVersionEqual_1Dot_9Dot_0,
  MultiFormatSchemaElseSchemaFormatAnyOf_2Values[MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashRamlPlusYamlSemicolonVersionEqual_1Dot_0]: MultiFormatSchemaElseSchemaFormatAnyOf_2ApplicationSlashRamlPlusYamlSemicolonVersionEqual_1Dot_0,
}
