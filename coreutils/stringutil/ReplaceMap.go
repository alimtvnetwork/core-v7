package stringutil

import "strings"

func ReplaceMap(
	replaceMapper map[string]string,
	content string,
) string {
	if content == "" {
		return content
	}

	if len(replaceMapper) == 0 {
		return content
	}

	for key, val := range replaceMapper {
		content = strings.ReplaceAll(
			content,
			key,
			val)
	}

	return content
}
