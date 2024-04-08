package parser

import (
	"encoding/json"
	"errors"

	// "github.com/Souvikns/parser-go/models/3.0.0"
	// m2 "github.com/Souvikns/parser-go/models/2.6.0"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)



type AsyncAPIObject struct {
	models.AsyncApi_3Dot_0Dot_0SchemaDot
}

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
			return Document{Doc: doc, FileType: YamlFile, RawDoc: document}, errors.New("Invalid document: Only support json or yaml file")
		}
	}
	return Document{Doc: doc, FileType: JsonFile, RawDoc: document}, err
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

func getModel(version string, document Document) any {
	data, _ := json.Marshal(document.Doc)
	switch version {
	case "3.0.0":
		var m models.AsyncApi_3Dot_0Dot_0SchemaDot
		json.Unmarshal(data, &m)
		return m
	case "2.6.0":
		var m models.AsyncApi_2Dot_6Dot_0SchemaDot
		json.Unmarshal(data, &m)
		return m
	}

	return nil
}
