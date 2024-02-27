package parser

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func Parse(document string) (any, error) {
	doc, err := loadDocument([]byte(document))
	if err != nil {
		return nil, err
	}
	version, found := extractVersion(doc)
	if !found {
		return nil, errors.New("missing asyncapi property")
	}

	return version, nil

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
