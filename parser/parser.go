package parser

import (
	"errors"
	AsyncApi_2Dot_6Dot_0SchemaDot "github.com/Souvikns/parser-go/models/2.6.0"
	AsyncApi_3Dot_0Dot_0SchemaDot "github.com/Souvikns/parser-go/models/3.0.0"
	"github.com/Souvikns/parser-go/scheme"
	// "io"
	// "net/http"
	// "os"
)

func Parse[K AsyncApi_3Dot_0Dot_0SchemaDot.AsyncApi_3Dot_0Dot_0SchemaDot | AsyncApi_2Dot_6Dot_0SchemaDot.AsyncApi_2Dot_6Dot_0SchemaDot](document string, obj *K) error {
	doc, err := loadDocument([]byte(document))
	if err != nil {
		return err
	}
	version, found := extractVersion(doc.Doc)
	if !found {
		return  errors.New("missing asyncapi property")
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

// func ParseFromFile(filepath string) (any, error) {
// 	data, err := os.ReadFile(filepath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return Parse(string(data))
// }

// func ParseFromUrl(url string) (any, error) {
// 	response, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer response.Body.Close()
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return Parse(string(body))
// }
