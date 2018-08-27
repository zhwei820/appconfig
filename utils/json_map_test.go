package utils

import (
	"testing"
)

func TestUnmarshalToMap(t *testing.T) {

	json := `{
		"test": "string",
		"test2": "string",
		"test3": "string",
		"test4": "string",
		"test5": ["x"]
	}`
	res, err := UnmarshalToMap([]byte(json))
	if err != nil {
		t.Log(err)
	}
	t.Log(res)

}

func TestGetJsonKeys(t *testing.T) {
	json := `{
		"test": "string",
		"test2": "string",
		"test3": "string",
		"test4": "string",
		"test5": ["x"]
	}`
	res := GetJsonKeys([]byte(json))

	t.Log(res)
	t.Log(len(res))

}
