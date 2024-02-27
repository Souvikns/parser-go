package parser

import (
	"errors"
	"github.com/Souvikns/parser-go/scheme"
	"io"
	"net/http"
	"os"
)

func Parse(document string) (any, error) {
	doc, err := loadDocument([]byte(document))
	if err != nil {
		return nil, err
	}
	version, found := extractVersion(doc.Doc)
	if !found {
		return nil, errors.New("missing asyncapi property")
	}

	// load correct scheme according to the version of the asyncapi document
	// if the document has not errors then create the appropriate model generated from the document
	// and return it.
	schemes, err := scheme.LoadSchemes()
	if err != nil {
		return nil, err
	}
	schema, err := schemes.Get(version)
	if err != nil {
		return nil, err
	}

	err = validateSchema(schema.Definition, document)
	if err != nil {
		return nil, err
	}

	return getModel(version, doc), nil

}

func ParseFromFile(filepath string) (any, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return Parse(string(data))
}

func ParseFromUrl(url string) (any, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return Parse(string(body))
}
