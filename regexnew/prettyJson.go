package regexnew

import (
	"bytes"
	"encoding/json"

	"gitlab.com/auk-go/core/constants"
)

// prettyJson
//
// Warning:
//
//	swallows error
func prettyJson(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	allBytes, err := json.Marshal(anyItem)

	if err != nil || len(allBytes) == 0 {
		return ""
	}

	var prettyJSON bytes.Buffer

	json.Indent(
		&prettyJSON,
		allBytes,
		constants.EmptyString,
		constants.Tab)

	return prettyJSON.String()
}
