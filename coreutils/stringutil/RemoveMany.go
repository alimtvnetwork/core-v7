package stringutil

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func RemoveMany(
	content string,
	removeRequests ...string,
) string {
	if content == "" {
		return content
	}

	for _, remove := range removeRequests {
		content = strings.ReplaceAll(
			content,
			remove,
			constants.EmptyString)
	}

	return content
}
