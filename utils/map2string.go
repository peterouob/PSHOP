package utils

import "encoding/json"

func Map2Json(data interface{}) string {
	byteStr, _ := json.Marshal(data)
	return string(byteStr)
}
