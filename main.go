package main

import (
	"fmt"
	"log"

	"github.com/Souvikns/parser-go/parser"
	model "github.com/Souvikns/parser-go/models/3.0.0"
)

func main() {
	var asyncapi model.AsyncApi_3Dot_0Dot_0SchemaDot
	err := parser.ParseFromFile("./spec.yaml", &asyncapi)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(asyncapi.Asyncapi)
}