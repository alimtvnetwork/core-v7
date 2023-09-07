package stringutil

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// RemoveManyBySplitting Remove as per removes then splits by the given separator
func RemoveManyBySplitting(
	content string,
	splitsBy string,
	removeRequests ...string,
) []string {
	for _, remove := range removeRequests {
		content = strings.ReplaceAll(content, remove, constants.EmptyString)
	}

	return strings.Split(content, splitsBy)
}
