package errcore

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type ExpectationMessageDef struct {
	CaseIndex         int
	FuncName          string
	TestCaseName      string
	When              string
	Expected          interface{}
	ActualProcessed   interface{}
	ExpectedProcessed interface{}
	IsNonWhiteSort    bool
	expectedString    *string
}

func (it ExpectationMessageDef) ExpectedSafeString() string {
	if it.expectedString != nil {
		return *it.expectedString
	}

	var expectedStr string

	if it.Expected != nil {
		expectedStr = fmt.Sprintf(
			constants.SprintValueFormat,
			it.Expected)
	}

	it.expectedString = &expectedStr

	return *it.expectedString
}

func (it ExpectationMessageDef) ExpectedStringTrim() string {
	return strings.TrimSpace(it.ExpectedString())
}

func (it ExpectationMessageDef) ExpectedString() string {
	if it.Expected == nil {
		panic("ExpectationMessageDef! Expected needs to be set before getting it!")
	}

	return it.ExpectedSafeString()
}

func (it ExpectationMessageDef) ToString(actual interface{}) string {
	return GetWhenActualAndExpectProcessedMessage(
		actual,
		&it)
}

func (it ExpectationMessageDef) PrintIf(
	isPrint bool,
	actual interface{},
) {
	if !isPrint {
		return
	}

	it.Print(actual)
}

func (it ExpectationMessageDef) PrintIfFailed(
	isPrintOnFail,
	isFailed bool,
	actual interface{},
) {
	if isPrintOnFail && isFailed {
		it.Print(actual)
	}
}

func (it ExpectationMessageDef) Print(actual interface{}) {
	msg := MsgHeaderPlusEnding(
		it.When,
		it.ToString(actual))

	fmt.Println(msg)
}
