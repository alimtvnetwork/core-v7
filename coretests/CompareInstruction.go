package coretests

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
)

type ComparingInstruction struct {
	actualHashset                    *corestr.Hashset
	actual                           string
	Header                           string
	MatchingAsEqual                  string
	ComparingItems                   []Compare
	HasWhitespace, IsMatchingAsEqual bool
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

	it.actualHashset = corestr.NewHashsetUsingStrings(&whitespaceRemovedSplits)

	return it.actualHashset
}

func (it *ComparingInstruction) IsMatches(
	index int,
	isPrint bool,
) bool {
	isMatchesEqual := !it.IsMatchingAsEqual || it.IsMatchingAsEqual &&
		IsStringMessageWithoutWhitespaceSortedEqual(
			isPrint,
			it.HasWhitespace,
			it.Header,
			it.actual,
			it.MatchingAsEqual,
			index)

	for i, item := range it.ComparingItems {
		isMatchesEqual = item.IsMatch(
			isPrint,
			i,
			it) &&
			isMatchesEqual
	}

	return isMatchesEqual
}
