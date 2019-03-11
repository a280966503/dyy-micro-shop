package utils

import "encoding/json"

func JsonToString(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
