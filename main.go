package main

import (
	"fmt"
	"github.com/Souvikns/parser-go/parser"
)

func main() {
	var asyncapi parser.Asyncapi_3_0_0
	err := parser.ParseFromFile("./spec.yaml", &asyncapi)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(asyncapi.Asyncapi)
}

