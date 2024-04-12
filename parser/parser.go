package parser

import (
	"errors"
	"github.com/Souvikns/parser-go/scheme"
	"io"
	"net/http"
	"os"
)

func Parse[K Asyncapi_3_0_0 | Asyncapi_2_6_0 | Asyncapi_2_5_0 | Asyncapi_2_4_0 | Asyncapi_2_3_0 | Asyncapi_2_2_0 | Asyncapi_2_1_0 | Asyncapi_2_0_0](document string, obj *K) error {
	doc, err := loadDocument([]byte(document))
	if err != nil {
		return err
	}
	version, found := extractVersion(doc.Doc)
	if !found {
		return errors.New("missing asyncapi property")
	}

	// load correct scheme according to the version of the asyncapi document
	// if the document has not errors then create the appropriate model generated from the document
	// and return it.
	schemes, err := scheme.LoadSchemes()
	if err != nil {
		return err
	}
	schema, err := schemes.Get(version)
	if err != nil {
		return err
	}

	err = validateSchema(schema.Definition, document)
	if err != nil {
		return err
	}

	getModel(version, doc, obj)
	return nil
}

func ParseFromFile[K Asyncapi_3_0_0 | Asyncapi_2_6_0 | Asyncapi_2_5_0 | Asyncapi_2_4_0 | Asyncapi_2_3_0 | Asyncapi_2_2_0 | Asyncapi_2_1_0 | Asyncapi_2_0_0](filepath string, obj *K) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = Parse(string(data), obj)
	if err !=nil {
		return err
	}
	return nil
}

func ParseFromUrlan[K Asyncapi_3_0_0 | Asyncapi_2_6_0 | Asyncapi_2_5_0 | Asyncapi_2_4_0 | Asyncapi_2_3_0 | Asyncapi_2_2_0 | Asyncapi_2_1_0 | Asyncapi_2_0_0](url string, obj *K) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = Parse(string(body), obj)
	if err != nil {
		return err
	}
	return nil
}
