package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func MsgWrapNumber(name string, number interface{}) string {
	return fmt.Sprintf(constants.StringWithBracketWrapNumberFormat, name, number)
}
