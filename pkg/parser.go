package pkg

import (
	"encoding/json"
	"errors"

	parserError "github.com/Souvikns/parser-go/pkg/error"
	schemes "github.com/asyncapi/spec-json-schemas/v6"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

const (
	JsonFile = "json"
	YamlFile = "yaml"
)

type Document struct {
	Doc      map[string]any
	FileType string
	Rawdoc   []byte
}

func Validate(document string) (bool, error) {
	doc, err := loadDocument([]byte(document))
	if err != nil {
		return false, err
	}

	version, found := extractVersion(doc.Doc)
	if !found {
		return false, errors.New("missing asyncapi property")
	}
	schema, err := schemes.Get(version)
	if err != nil {
		return false, err
	}
	validationErrors := validateSchema(schema, doc.Doc)
	if validationErrors != nil {
		return false, err
	}
	return true, nil
}

func loadDocument(document []byte) (Document, error) {
	doc := make(map[string]any)
	err := json.Unmarshal(document, &doc)
	if err != nil {
		err := yaml.Unmarshal(document, &doc)
		if err != nil {
			return Document{Doc: doc, FileType: YamlFile, Rawdoc: document}, errors.New("Invalid document type: Only support json or yaml")
		} else {
			return Document{Doc: doc, FileType: YamlFile, Rawdoc: document}, nil
		}
	}
	return Document{Doc: doc, FileType: JsonFile, Rawdoc: document}, nil
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
		return parserError.New(errs...)
	}
	return nil
}
