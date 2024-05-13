package ref

import (
	"os"

	"github.com/bserdar/gojsonpath"
	"github.com/lestrrat-go/jsref"
	"github.com/lestrrat-go/jsref/provider"
	"github.com/xeipuuv/gojsonpointer"
)



func GetRef(refString string, doc any) ([]any, error) {
	path, err := gojsonpath.Parse(refString)
	if err != nil {
		return nil, err
	}

	result, err := gojsonpath.Find(gojsonpath.MapModel{Doc: doc}, path)
	return result, err
}

func ResolveRef(pointer string, data any, options []jsref.Option) (any, error) {
	mp := provider.NewMap()
	rootPath, _ := os.Getwd()
	fp := provider.NewFS(rootPath)

	res := jsref.New()
	res.AddProvider(mp)
	res.AddProvider(fp)

	return res.Resolve(data, pointer, options...)
}

func Update(jsonObject *map[string]any, pointerString string, data any) {
	pointer, _ := gojsonpointer.NewJsonPointer(pointerString)
	pointer.Set(jsonObject, data)
}

