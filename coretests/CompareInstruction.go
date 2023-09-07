package coretests

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
)

type ComparingInstruction struct {
	FunName,
	Header,
	TestCaseName,
	MatchingAsEqualExpectation string
	ComparingItems                   []Compare
	HasWhitespace, IsMatchingAsEqual bool
	actualHashset                    *corestr.Hashset
	actual                           string
}

func (it *ComparingInstruction) Actual() string {
	return it.actual
}

func (it *ComparingInstruction) SetActual(actual string) {
	it.actual = actual
	it.actualHashset = nil
}

func (it *ComparingInstruction) ActualHashset() *corestr.Hashset {
	if it.actualHashset != nil {
		return it.actualHashset
	}

	whitespaceRemovedSplits := GetMessageToSortedArray(
		false,
		true,
		it.Actual())

	it.actualHashset = corestr.New.Hashset.StringsPtr(&whitespaceRemovedSplits)

	return it.actualHashset
}

func (it *ComparingInstruction) IsMatch(
	caseIndexPlusIsPrint *CaseIndexPlusIsPrint,
) bool {
	isMatchesEqual := it.isMatchingEqual(caseIndexPlusIsPrint)

	for i, item := range it.ComparingItems {
		isMatchesEqual = item.IsMatch(
			caseIndexPlusIsPrint.IsPrint,
			i,
			it) &&
			isMatchesEqual
	}

	return isMatchesEqual
}

func (it *ComparingInstruction) isMatchingEqual(caseIndexPlusIsPrint *CaseIndexPlusIsPrint) bool {
	if !it.IsMatchingAsEqual {
		return true
	}

	expectation := &errcore.ExpectationMessageDef{
		CaseIndex:      caseIndexPlusIsPrint.CaseIndex,
		FuncName:       it.FunName,
		TestCaseName:   it.TestCaseName,
		When:           it.Header,
		Expected:       it.MatchingAsEqualExpectation,
		IsNonWhiteSort: it.HasWhitespace,
	}

	return IsStrMsgNonWhiteSortedEqual(
		caseIndexPlusIsPrint.IsPrint,
		it.actual,
		expectation)
}
