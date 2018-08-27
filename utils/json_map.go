package utils

import (
	"github.com/francoispqt/gojay"
	"github.com/astaxie/beego/logs"
)

// define our custom map type implementing MarshalerJSONObject and UnmarshalerJSONObject
type myMap map[string]string

// Implementing Unmarshaler
func (m myMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	str := ""
	err := dec.AddString(&str)
	if err != nil {
		return err
	}
	m[k] = str
	return nil
}

// Her we put the number of keys
// If number of keys is unknown return 0, it will parse all keys
func (m myMap) NKeys() int {
	return 0
}

// Implementing Marshaler
func (m myMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range m {
		enc.AddStringKey(k, v)
	}
}

func (m myMap) IsNil() bool {
	return m == nil
}

// json str to map[string][string]
func UnmarshalToMap(b []byte) (map[string]string, error) {
	nM := myMap(make(map[string]string))
	err := gojay.Unmarshal(b, nM)

	return map[string]string(nM), err
}

// get 1st level json keys by UnmarshalToMap
func GetJsonKeys(b []byte) ([]string) {
	nM, err := UnmarshalToMap(b)
	if err != nil {
		logs.Error("UnmarshalToMap error:src: %v, dst: %v, error: %v", string(b), nM, err)
	}
	var keys []string
	for key := range nM {
		keys = append(keys, key)
	}
	return keys

}
