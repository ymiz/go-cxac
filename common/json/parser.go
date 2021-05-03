package json

import jsoniter "github.com/json-iterator/go"

func Unmarshal(jsonBytes []byte, v interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(jsonBytes, &v)
}

func Marshal(v interface{}) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(&v)
}
