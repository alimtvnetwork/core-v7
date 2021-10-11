package enumimpl

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func PrependJoin(
	joiner string,
	prepend interface{},
	anyItems ...interface{},
) string {
	slice := make([]string, len(anyItems)+constants.Capacity1)
	slice[constants.Zero] = fmt.Sprintf(constants.SprintValueFormat, prepend)

	for i, item := range anyItems {
		slice[i+1] = fmt.Sprintf(
			constants.SprintValueFormat,
			item)
	}

	return strings.Join(slice, joiner)
}
