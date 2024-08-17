package helper

import "encoding/json"

func ToJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func FromJSON[T any](data []byte, v *T) error {
	return json.Unmarshal(data, v)
}
