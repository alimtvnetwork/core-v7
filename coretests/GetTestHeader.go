package coretests

import (
	"fmt"
)

func GetTestHeader(testCaseMessenger TestCaseMessenger) string {
	return fmt.Sprintf("Method : [%s]",
		testCaseMessenger.FuncName(),
	)
}
