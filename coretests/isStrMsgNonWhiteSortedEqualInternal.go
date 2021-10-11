package coretests

import (
	"strings"

	"gitlab.com/evatix-go/core/errcore"
)

func isStrMsgNonWhiteSortedEqualInternal(
	isPrintOnFail bool,
	actual string,
	expectationMessageDef *errcore.ExpectationMessageDef,
) bool {
	actualSortedDefault := GetMessageToSorted(
		false,
		strings.TrimSpace(actual),
		commonJoiner)

	expectedSortedDefault := GetMessageToSorted(
		false,
		expectationMessageDef.ExpectedStringTrim(),
		commonJoiner)

	isEqual := actualSortedDefault == expectedSortedDefault
	isFailed := !isEqual

	// Exception case for mutation, because test updates it
	expectationMessageDef.ActualProcessed = actualSortedDefault
	expectationMessageDef.ExpectedProcessed = expectedSortedDefault
	expectationMessageDef.PrintIfFailed(
		isPrintOnFail,
		isFailed,
		actual)

	return isEqual
}
