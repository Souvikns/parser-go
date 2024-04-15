package parser

import (
	"encoding/json"
	"errors"

	parseError "github.com/Souvikns/parser-go/error"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)



const (
	JsonFile = "Json"
	YamlFile = "Yaml"
)

type Document struct {
	Doc      map[string]any
	FileType string
	RawDoc   []byte
}

func loadDocument(document []byte) (Document, error) {
	doc := make(map[string]any)
	err := json.Unmarshal(document, &doc)
	if err != nil {
		err := yaml.Unmarshal(document, &doc)
		if err != nil {
			return Document{Doc: doc, FileType: YamlFile, RawDoc: document}, errors.New("invalid document: Only support json or yaml file")
		} else {
			return Document{Doc: doc, FileType: YamlFile, RawDoc: document}, nil
		}
	}
	return Document{Doc: doc, FileType: JsonFile, RawDoc: document}, nil
}

func extractVersion(v map[string]any) (string, bool) {
	version, err := v["asyncapi"]
	return version.(string), err
}

func validateSchema(definition []byte, document any) error {
	schemaLoader := gojsonschema.NewBytesLoader(definition)
	documentLoader := gojsonschema.NewGoLoader(document)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		var errs []error
		for _, err := range result.Errors() {
			errs = append(errs, errors.New(err.String()))
		}
		return parseError.New(errs...)
	}

	return nil
}

func getModel[K Asyncapi_3_0_0 | Asyncapi_2_6_0 | Asyncapi_2_5_0 | Asyncapi_2_4_0 | Asyncapi_2_3_0 | Asyncapi_2_2_0 | Asyncapi_2_1_0 | Asyncapi_2_0_0](version string, document Document, obj *K) {
	data, _ := json.Marshal(document.Doc)
	json.Unmarshal(data, obj)
}
