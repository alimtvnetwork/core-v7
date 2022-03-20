package strutilinternal

import (
	"strings"
)

func ReplaceTemplateMap(
	isCurly bool,
	templateFormat string,
	replacerMap map[string]string, // key ==> find, value ==> replace
) string {
	if templateFormat == "" || len(replacerMap) == 0 {
		return templateFormat
	}

	for finder, replacer := range replacerMap {
		finderCurly := CurlyWrapIf(isCurly, finder)

		templateFormat = strings.ReplaceAll(
			templateFormat,
			finderCurly,
			replacer)
	}

	return templateFormat
}
