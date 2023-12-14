package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

type getAssertSimpleTestCaseWrapper struct{}

// GetAssertMessageUsingSimpleTestCaseWrapper
//
//	Gives generic and consistent test message using msgformats.QuickIndexTitleInputActualExpectedMessageFormat
func (it getAssertSimpleTestCaseWrapper) String(
	testCaseIndex int,
	testCaseWrapper SimpleTestCaseWrapper,
) string {
	return fmt.Sprintf(
		msgformats.QuickIndexTitleInputActualExpectedMessageFormat,
		testCaseIndex,
		testCaseWrapper.CaseTitle(),
		testCaseWrapper.Input(),
		testCaseWrapper.Actual(),
		testCaseWrapper.Expected(),
	)
}

func (it getAssertSimpleTestCaseWrapper) Lines(
	testCaseWrapper SimpleTestCaseWrapper,
) (actualLines, expectedLines []string) {
	toStringsFunc := GetAssert.ToStrings
	actualLines = toStringsFunc(testCaseWrapper.Actual())
	expectedLines = toStringsFunc(testCaseWrapper.Expected())

	return actualLines, expectedLines
}

// CaseLinesUsingDoubleQuoteLinesToString
//
// Actual lines and then expected lines.
func (it getAssertSimpleTestCaseWrapper) CaseLinesUsingDoubleQuoteLinesToString(
	testCaseIndex int,
	testCaseWrapper SimpleTestCaseWrapper,
) string {
	toStringsFunc := GetAssert.ToStrings
	prefixSpaces := 4
	actualLines := toStringsFunc(testCaseWrapper.Actual())
	expectedLines := toStringsFunc(testCaseWrapper.Expected())

	actual := GetAssert.ConvertLinesToDoubleQuoteThenString(prefixSpaces, actualLines)
	expected := GetAssert.ConvertLinesToDoubleQuoteThenString(prefixSpaces, expectedLines)
	title := testCaseWrapper.CaseTitle()

	return fmt.Sprintf(
		msgformats.QuickLinesFormat,
		testCaseIndex,
		title,
		testCaseIndex,
		title,
		actual,
		testCaseIndex,
		title,
		expected,
	)
}

func GetAssertMessage(testCaseMessenger TestCaseMessenger, counter int) string {
	return GetAssert.Quick(
		testCaseMessenger.Value(),
		testCaseMessenger.Actual(),
		testCaseMessenger.Expected(),
		counter,
	)
}

func GetTestHeader(testCaseMessenger TestCaseMessenger) string {
	return fmt.Sprintf(
		"CompareMethod : [%s]",
		testCaseMessenger.GetFuncName(),
	)
}
