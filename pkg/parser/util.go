package parser

import (
	"encoding/json"
	"errors"

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