package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func ErrPrintWithTestIndex(
	caseIndex int,
	err error,
) {
	if err != nil {
		fmt.Print(
			"Case Index: ",
			caseIndex,
			constants.CommaSpace,
			constants.NewLineUnix,
			err)
	}
}
