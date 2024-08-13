
# AsyncAPI Parser (GO) V0.0.1

AsyncAPI parser written for `Go` applications, supports v2 to v3 AsyncAPI specification. 

### Installation

```
go get github.com/Souvikns/parser-go
```

### Usage 
The package exposes four main functions, `Parse`, `ParseFromFile`, `ParseFromUrl` and `Validate()`.

- `Parse()` - function that validates the passed AsyncAPI document, It expects document as a string format
- `ParseFromFile()` - function that validates the passed AsyncAPI document and it reads the document from a file path that is provided.
- `ParseFromUrl()` - function that validates the passed AsyncAPI document and it reads the document from a URL.
- `Validate()` - function that validates the passed AsyncAPI document, Return boolean to signify the validation status and errors if any. 

The package also exposes AsyncAPI models as per versions. 

|models|
|----|
|Asyncapi_3_0_0|
|Asyncapi_2_6_0|
|Asyncapi_2_5_0|
|Asyncapi_2_4_0|
|Asyncapi_2_3_0|
|Asyncapi_2_2_0|
|Asyncapi_2_1_0|
|Asyncapi_2_1_0|
|Asyncapi_2_0_0|

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

The `Parse` functions takes the `AsyncAPI` document as string and an object of type `interface{}` as argument to validate and parse the specification file.

```go
func main(){
	var asyncapi parser.Asyncapi_3_0_0
	err := parser.Parse(asyncapiDoc, &asyncapi)
}
```


#### `ParseFromFile(filepath string, obj *k) error`

The `ParseFromFile` function takes the filepath as string and an object of type `interface{}` as argument to validate and parse the specification file.

```go
func main(){
	var asyncapi parser.Asyncapi_3_0_0
	err := parser.ParseFromFile("./asyncapi.yaml", &asyncapi)
}
```

#### `ParseFromUrl(url string, obj *k) error`

The `ParseFromUrl` function takes the url as a string and an object of type `interface{}` as argument to validate and parse the specification file. 

```go
func main(){
	var asyncapi parser.Asyncapi_3_0_0
	err := parser.ParseFromUrl("https://localhost:8080/spec", &asyncapi)
}
```

#### `Validate(document string)`

The `Validate` function takes the AsyncAPI specification as a string and returns `true` if valid and `false` if invalid, along with the appropriate errors. 

```go
func main() {
    isValid := parser.Validate(asyncapi) // True if valid and False if invalid
}
```


### Develop

1. Make sure u are using go v1.19 or above
2. Write code and tests 
3. Make use all tests are passing


### Contributing

Read [Contributing](https://github.com/asyncapi/.github/blob/master/CONTRIBUTING.md) guide. 
