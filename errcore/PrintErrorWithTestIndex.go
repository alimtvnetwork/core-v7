package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func PrintErrorWithTestIndex(
	caseIndex int,
	header string,
	err error,
) {
	if err != nil {
		fmt.Print(
			"Case Index: ",
			caseIndex,
			constants.NewLineUnix,
			" \t    Title: ",
			header,
			constants.NewLineUnix,
			" \t  Summary: ",
			err)
	}
}
