package main

import (
	"encoding/json"
	"fmt"
	"github.com/Souvikns/parser-go/models/3.0.0"
)

func main() {
	data := `"password"`
	var v models.ApiKeyIn
	json.Unmarshal([]byte(data), &v)
	fmt.Println(v.Value())
	d , err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))

	var apikeyin models.ApiKeyIn = 0
	fmt.Println(apikeyin.Value())
}