package coretests

import (
	"strings"

	"gitlab.com/evatix-go/core/errcore"
)

func IsStrMsgNonWhiteSortedEqual(
	isPrintOnFail bool,
	actual string,
	expectationMessageDef *errcore.ExpectationMessageDef,
) bool {
	if expectationMessageDef.IsNonWhiteSort {
		return isStrMsgNonWhiteSortedEqualInternal(
			isPrintOnFail,
			actual,
			expectationMessageDef)
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
		actual)

	return isEqual
}
