package lib

import "encoding/json"

func ParseJSON(b []byte) map[string]interface{} {
	var resp map[string]interface{}
	json.Unmarshal(b, &resp)

	return resp
}
