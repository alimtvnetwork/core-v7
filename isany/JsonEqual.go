package isany

import (
	"bytes"
	"encoding/json"
)

func JsonEqual(
	left, right interface{},
) bool {
	leftBytes, leftErr := json.Marshal(left)
	rightBytes, rightErr := json.Marshal(right)

	if leftErr != nil && rightErr != nil && rightErr.Error() != leftErr.Error() {
		return false
	}

	if leftErr != nil || rightErr != nil {
		return false
	}

	return bytes.Equal(leftBytes, rightBytes)
}
