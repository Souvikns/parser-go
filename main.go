package main

import (
	"fmt"
	"log"
	"github.com/Souvikns/parser-go/parser"
)

func main() {
	var asyncapi parser.Asyncapi_3_0_0
	err := parser.ParseFromFile("./spec.yaml", &asyncapi)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(asyncapi.Channels["userSignedUp"].Address)
}
