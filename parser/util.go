package parser

import (
	"encoding/json"
	"errors"
	AsyncApi_2Dot_6Dot_0SchemaDot "github.com/Souvikns/parser-go/models/2.6.0"
	AsyncApi_3Dot_0Dot_0SchemaDot "github.com/Souvikns/parser-go/models/3.0.0"
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

func getModel[K AsyncApi_3Dot_0Dot_0SchemaDot.AsyncApi_3Dot_0Dot_0SchemaDot | AsyncApi_2Dot_6Dot_0SchemaDot.AsyncApi_2Dot_6Dot_0SchemaDot](version string, document Document, obj *K) {
	data, _ := json.Marshal(document.Doc)
	json.Unmarshal(data, obj)
}
