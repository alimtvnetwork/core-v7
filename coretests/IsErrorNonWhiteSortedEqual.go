package coretests

import "gitlab.com/auk-go/core/errcore"

func IsErrorNonWhiteSortedEqual(
	isPrintOnFail bool,
	actual error,
	expectationMessageDef *errcore.ExpectationMessageDef,
) bool {
	var actualErrorMessage string

	if actual != nil {
		actualErrorMessage = actual.Error()
	}

	expectedString := expectationMessageDef.ExpectedString()

	if expectedString == "" && actualErrorMessage == "" {
		return true
	}

	return IsStrMsgNonWhiteSortedEqual(
		isPrintOnFail,
		actualErrorMessage,
		expectationMessageDef)
}
