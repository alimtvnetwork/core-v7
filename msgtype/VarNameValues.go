package msgtype

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func VarNameValues(
	nameValues ...NameVal,
) string {
	if len(nameValues) == 0 {
		return ""
	}

	items := make([]string, len(nameValues))

	index := 0
	for _, nameValue := range nameValues {
		items[index] = fmt.Sprintf(
			keyValFormat,
			nameValue.Name,
			nameValue.Value)
		index++
	}

	return strings.Join(
		items,
		constants.CommaSpace)
}
