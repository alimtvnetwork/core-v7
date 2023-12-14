package coretests

import (
	"strings"

	"gitlab.com/auk-go/core/errcore"
)

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
		expectationMessageDef,
	)
}

func IsStrMsgNonWhiteSortedEqual(
	isPrintOnFail bool,
	actual string,
	expectationMessageDef *errcore.ExpectationMessageDef,
) bool {
	if expectationMessageDef.IsNonWhiteSort {
		return isStrMsgNonWhiteSortedEqualInternal(
			isPrintOnFail,
			actual,
			expectationMessageDef,
		)
	}

	trimActual := strings.TrimSpace(actual)
	trimExpected := expectationMessageDef.ExpectedStringTrim()
	isEqual := trimActual == trimExpected
	isFailed := !isEqual

	expectationMessageDef.ActualProcessed = trimActual
	expectationMessageDef.ExpectedProcessed = trimExpected
	expectationMessageDef.PrintIfFailed(
		isPrintOnFail,
		isFailed,
		actual,
	)

	return isEqual
}

func isStrMsgNonWhiteSortedEqualInternal(
	isPrintOnFail bool,
	actual string,
	expectationMessageDef *errcore.ExpectationMessageDef,
) bool {
	actualSortedDefault := GetAssert.SortedMessage(
		false,
		strings.TrimSpace(actual),
		commonJoiner,
	)

	expectedSortedDefault := GetAssert.SortedMessage(
		false,
		expectationMessageDef.ExpectedStringTrim(),
		commonJoiner,
	)

	isEqual := actualSortedDefault == expectedSortedDefault
	isFailed := !isEqual

	// Exception case for mutation, because test updates it
	expectationMessageDef.ActualProcessed = actualSortedDefault
	expectationMessageDef.ExpectedProcessed = expectedSortedDefault
	expectationMessageDef.PrintIfFailed(
		isPrintOnFail,
		isFailed,
		actual,
	)

	return isEqual
}
