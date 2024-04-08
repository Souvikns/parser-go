package main

import (
	"fmt"
	"log"

	"github.com/Souvikns/parser-go/parser"
	model "github.com/Souvikns/parser-go/models/3.0.0"
	"os"
)

func main() {
	spec, err := os.ReadFile(("./spec.yaml"))
	if err !=nil {
		log.Fatalf(err.Error())
	}
	var asyncapi model.AsyncApi_3Dot_0Dot_0SchemaDot
	err = parser.Parse(string(spec), &asyncapi)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(asyncapi.Asyncapi)
}