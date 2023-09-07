package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

// GetAssertMessageUsingSimpleTestCaseWrapper
//
//  Gives generic and consistent test message using msgformats.QuickIndexTitleInputActualExpectedMessageFormat
func GetAssertMessageUsingSimpleTestCaseWrapper(testCaseIndex int, testCaseWrapper SimpleTestCaseWrapper) string {
	return fmt.Sprintf(
		msgformats.QuickIndexTitleInputActualExpectedMessageFormat,
		testCaseIndex,
		testCaseWrapper.CaseTitle(),
		testCaseWrapper.Input(),
		testCaseWrapper.Actual(),
		testCaseWrapper.Expected())
}
