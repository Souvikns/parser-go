package parser

import "errors"

func Parse(document string) (any, error) {
	doc, err := loadDocument([]byte(document))
	if err !=nil {
		return nil, err
	}
	version, found := extractVersion(doc)
	if found == false {
		return nil, errors.New("Missing asyncapi property")
	}

	return version, nil

}

func ParseFromFile(filepath string) (any, error) {
	return Parse("")
}

func ParseFromUrl(url string) (any, error) { 
	return Parse("")
}

