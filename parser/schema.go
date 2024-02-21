package parser

import (
	"errors"
	"os"
	"strings"
)

type Schema struct {
	Version    string
	Definition []byte
}

func (s Schema) ToString() string {
	return string(s.Definition)
}

type Schemas struct {
	Scheme []Schema
}

func LoadSchemes() (schemas Schemas, err error) {
	files, err := os.ReadDir("generator/definitions/")
	if err != nil {
		return schemas, err
	}

	for _, f := range files {
		data, err := os.ReadFile("generator/definitions/" + f.Name())
		if err != nil {
			return schemas, err
		}
		schemas.Scheme = append(schemas.Scheme, Schema{Version: getVersion(f.Name()), Definition: data})
	}
	return schemas, nil
}

func getVersion(filename string) string {
	filename = strings.Replace(filename, "-without-$id.json", "", 1)
	return filename
}


func (s Schemas) Get(version string) (Schema, error) {
	for _, schema := range s.Scheme {
		if schema.Version == version {
			return schema, nil
		}
	}
	return Schema{}, errors.New("Schema not found")
}