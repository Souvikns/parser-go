package parser

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	parserError "github.com/Souvikns/parser-go/pkg/error"
	schemes "github.com/asyncapi/spec-json-schemas/v6"
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

func Validate(document string) (bool, error) {
	doc, err := loadDocument([]byte(document))
	if err != nil {
		return false, err
	}
	version, found := extractVersion((doc.Doc))
	if !found {
		return false, errors.New("missing asyncapi property")
	}
	schema, err := schemes.Get(version)
	if err != nil {
		return false, err
	}
	err = validateSchema(schema, doc.Doc)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Parse[K Asyncapi_3_0_0 | Asyncapi_2_6_0 | Asyncapi_2_5_0 | Asyncapi_2_4_0 | Asyncapi_2_3_0 | Asyncapi_2_2_0 | Asyncapi_2_1_0 | Asyncapi_2_0_0](document string, obj *K) error {
	doc, err := loadDocument([]byte(document))
	if err != nil {
		return err
	}
	version, found := extractVersion(doc.Doc)
	if !found {
		return errors.New("missing asyncapi property")
	}
	schema, err := schemes.Get(version)

	if err != nil {
		return err
	}

	err = validateSchema(schema, doc.Doc)
	if err != nil {
		return err
	}

	parseModel(doc, obj)

	return nil

}

func ParseFromFile[K Asyncapi_3_0_0 | Asyncapi_2_6_0 | Asyncapi_2_5_0 | Asyncapi_2_4_0 | Asyncapi_2_3_0 | Asyncapi_2_2_0 | Asyncapi_2_1_0 | Asyncapi_2_0_0](filepath string, obj *K) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	return Parse(string(data), obj)
}

func ParseFromUrl[K Asyncapi_3_0_0 | Asyncapi_2_6_0 | Asyncapi_2_5_0 | Asyncapi_2_4_0 | Asyncapi_2_3_0 | Asyncapi_2_2_0 | Asyncapi_2_1_0 | Asyncapi_2_0_0](url string, obj *K) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return Parse(string(body), obj)
}

func loadDocument(document []byte) (Document, error) {
	doc := make(map[string]any)
	err := json.Unmarshal(document, &doc)
	if err != nil {
		err := yaml.Unmarshal(document, &doc)
		if err != nil {
			return Document{Doc: doc, FileType: YamlFile, RawDoc: document}, errors.New("invalid document type: Only support json or yaml")
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
		return parserError.New(errs...)
	}
	return nil
}

func parseModel[K Asyncapi_3_0_0 | Asyncapi_2_6_0 | Asyncapi_2_5_0 | Asyncapi_2_4_0 | Asyncapi_2_3_0 | Asyncapi_2_2_0 | Asyncapi_2_1_0 | Asyncapi_2_0_0](document Document, obj *K) {
	data, _ := json.Marshal(document.Doc)
	json.Unmarshal(data, obj)
}
