package parser

import (
	"encoding/json"
	"errors"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

func loadDocument(document []byte) (map[string]any, error) {
	doc := make(map[string]any)
	err := json.Unmarshal(document, &doc)
	if err != nil {
		err := yaml.Unmarshal(document, &doc)
		if err != nil {
			return doc, errors.New("Invalid document: Only support json or yaml file")
		}

	}

	return doc, err
}

func extractVersion(v map[string]any) (string, bool) {
	version, err := v["asyncapi"]
	return version.(string), err
}

func validateSchema(definition []byte, document string) error {
	schemaLoader := gojsonschema.NewBytesLoader(definition)
	documentLoader := gojsonschema.NewGoLoader(document)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {

	}

	return nil
}
