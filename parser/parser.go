package parser

import (
	"errors"
	"github.com/xeipuuv/gojsonschema"
)

type Parser struct {
	schemaLoader gojsonschema.JSONLoader
}

func NewParser(schema []byte) *Parser {
	return &Parser{
		schemaLoader: gojsonschema.NewBytesLoader(schema),
	}
}

func (p *Parser) Parse(data []byte) error {
	documentLoader := gojsonschema.NewBytesLoader(data)
	result, err := gojsonschema.Validate(p.schemaLoader, documentLoader)
	if err != nil {
		return err
	}
	if !result.Valid() {
		var errs []error
		for _, err := range result.Errors() {
			errs = append(errs, errors.New(err.String()))
		}
		return errs[0]
	}
	return nil
}

func extractVersion(v interface{}) (string, error) {
	switch doc := v.(type) {
	case map[string]interface{}:
		if doc["asyncapi"] == nil {
			return "", errors.New("the `asyncapi` field is missing")
		}

		return doc["asyncapi"].(string), nil
	default:
		return "", errors.New("only map[string]interface{} type is supported")
	}
}
