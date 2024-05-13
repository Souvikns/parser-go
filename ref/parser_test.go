package ref

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetRef(t *testing.T) {
	jsonText := `{
		"channels": {
			"index":{
				"$ref": "#/components/channels/index"
			}
		},
		"components": {
			"channels": {
				"index": {
					"address" : "/"
				}
			}
		}
	}`
	var jsonDocument map[string]any
	json.Unmarshal([]byte(jsonText), &jsonDocument)
	refs, err := GetRef("*..$ref", jsonDocument)
	t.Log(refs)
	assert.Contains(t, refs[0], "#/components/channels/index")
	assert.Nil(t, err)
}

func TestResolveRef(t *testing.T) {
	jsonText := `{
		"components": {
			"channels": {
				"index": {
					"address" : "/"
				}
			}
		}
	}`
	var jsonDocument map[string]any
	json.Unmarshal([]byte(jsonText), &jsonDocument)
	result, err := ResolveRef("#/components/channels/index", jsonDocument, nil)
	t.Log(result)
	assert.Contains(t, result.(map[string]any)["address"], "/")
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	jsonText := `{
		"channels": {
			"index":{
				"$ref": "#/components/channels/index"
			}
		},
		"components": {
			"channels": {
				"index": {
					"address" : "/"
				}
			}
		}
	}`

	var jsonDocument map[string]any
	json.Unmarshal([]byte(jsonText), &jsonDocument)
	result, _ := ResolveRef("#/components/channels/index", jsonDocument, nil)
	t.Log(result)
	Update(&jsonDocument, "/channels/index", result)
	t.Log(jsonDocument)
	assert.Nil(t, jsonDocument)
}


// func TestUpdateObject(t *testing.T) {
// 	jsonText := `{
// 		"components": {
// 			"channels": {
// 				"index": {
// 					"address" : "/"
// 				}
// 			}
// 		}
// 	}`
// 	var jsonDocument map[string]interface{}
// 	json.Unmarshal([]byte(jsonText), &jsonDocument)
// 	pointerString := "/components/channels/index/address"
// 	pointer, _ := gojsonpointer.NewJsonPointer(pointerString)
// 	pointer.Set(jsonDocument, "/index")
// 	add, _, _ := pointer.Get(jsonDocument)
// 	t.Log(add)
// 	b, _ := json.Marshal(jsonDocument)
// 	t.Log(string(b))
// 	assert.Nil(t, b)
// }