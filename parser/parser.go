package parser

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Parser func(io.Reader, io.Writer) error

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func isLocalFile(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func NewReader(doc string) (io.Reader, error) {
	r := bytes.NewBuffer(nil)
	if isURL(doc) {
		resp, err := http.Get(doc)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		_, err = io.Copy(r, resp.Body)
		return r, err
	}

	if isLocalFile(doc) {
		f, err := os.Open(doc)
		if err != nil {
			return nil, err
		}

		defer f.Close()
	}

	return r, nil
}
