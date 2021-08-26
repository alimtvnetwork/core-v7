package coretests

import (
	"fmt"
)

func GetTestHeader(testCaseMessenger TestCaseMessenger) string {
	return fmt.Sprintf("CompareMethod : [%s]",
		testCaseMessenger.FuncName(),
	)
}
