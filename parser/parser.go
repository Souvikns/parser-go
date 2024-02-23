
import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/Souvikns/parser-go/parser"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
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
	// schemes, err := LoadSchemes()
	// if err != nil {
	// 	return err
	// }

	doc, err := loadDocument(data)
	if err != nil {
		return err
	}
	version, _ := extractVersion(doc)

	schemes, err := LoadSchemes()
	if err != nil {
		fmt.Println(err)
	}

	scheme, _ := schemes.Get(version)

	

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

func extractVersion(v map[string]any) (string, bool) {
	version, err := v["asyncapi"]
	return version.(string), err
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func isLocalFile(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func loadDocument(document []byte) (map[string]any, error) {
	doc := make(map[string]any)
	err := json.Unmarshal(document, &doc)
	if err != nil {
		err := yaml.Unmarshal(document, &doc)
		if err != nil {
			return doc, errors.New("Invalid spec type, it should be json or yaml")
		}
	}
	return doc, nil
}
