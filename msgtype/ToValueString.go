package msgtype

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func ToValueString(reference interface{}) string {
	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		reference)
}
