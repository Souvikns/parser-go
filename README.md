
# AsyncAPI Parser (GO) V0.0.1

AsyncAPI parser written for `Go` applications, supports v2 to v3 AsyncAPI specification. 

### Installation

```
go get github.com/Souvikns/parser-go
```

### Usage 
The package exposes three main functions, `Parse`, `ParseFromFile`, `ParseFromUrl`

- `Parse()` - function that validates the passed AsyncAPI document, It expects document as a string format
- `ParseFromFile()` - function that validates the passed AsyncAPI document and it reads the document from a file path that is provided.
- `ParseFromUrl()` - function that validates the passed AsyncAPI document and it reads the document from a URL.

### Example

```go
package main

import(
    "fmt"
    "log"
    "github.com/Souvikns/parser-go/parser"
)

func main() {
    var asyncapi parser.Asyncapi_3_0_0
    err := parser.ParseFromFile("./asyncapi.yaml", &asyncapi)
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Println(asyncapi.AsyncAPI)
}
```

### API

#### `Parse(document string, obj *k) error`

```go
func main(){
	var asyncapi parser.Asyncapi_3_0_0
	err := parser.Parse(asyncapiDoc, &asyncapi)
}
```


#### `ParseFromFile(filepath string, obj *k) error`

```go
func main(){
	var asyncapi parser.Asyncapi_3_0_0
	err := parser.ParseFromFile("./asyncapi.yaml", &asyncapi)
}
```

#### `ParseFromUrl(url string, obj *k) error`


```go
func main(){
	var asyncapi parser.Asyncapi_3_0_0
	err := parser.ParseFromUrl("https://localhost:8080/spec", &asyncapi)
}
```


### Develop

1. Make sure u are using go v1.19 or above
2. Write code and tests 
3. Make use all tests are passing


### Contributing

Read [Contributing](https://github.com/asyncapi/.github/blob/master/CONTRIBUTING.md) guide. 
