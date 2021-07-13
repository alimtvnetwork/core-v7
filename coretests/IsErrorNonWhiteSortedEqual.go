package coretests

import "gitlab.com/evatix-go/core/msgtype"

func IsErrorNonWhiteSortedEqual(
	isPrintOnFail bool,
	actual error,
	expectationMessageDef *msgtype.ExpectationMessageDef,
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
