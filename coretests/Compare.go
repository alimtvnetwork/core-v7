package coretests

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type Compare struct {
	sortedString   *string
	sortedStrings  []string
	MatchingLength int // 0/-1 means all, number means, that number must match
	StringContains string
}

func (it *Compare) SortedStrings() []string {
	if it.sortedStrings != nil {
		return it.sortedStrings
	}

	it.sortedStrings = GetAssert.SortedArray(
		false,
		true,
		strings.TrimSpace(it.StringContains),
	)

	return it.sortedStrings
}

func (it *Compare) SortedString() string {
	if it.sortedString != nil {
		return *it.sortedString
	}

	sortedStrings := it.SortedStrings()
	sortedString := strings.Join(
		sortedStrings,
		commonJoiner,
	)

	it.sortedString = &sortedString

	return *it.sortedString
}

func (it *Compare) GetPrintMessage(index int) string {
	return fmt.Sprintf(
		"\n\tIndex:%d\n\tString Contains:%s\n\tString Processed:%s",
		index,
		it.StringContains,
		it.SortedString(),
	)
}

func (it *Compare) IsMatch(
	isPrint bool,
	index int,
	instruction *ComparingInstruction,
) bool {
	actualHashset := instruction.ActualHashset()
	sortedStrings := it.SortedStrings()

	// all
	if it.MatchingLength <= constants.Zero {
		isMatch := actualHashset.HasAll(sortedStrings...)

		if !isMatch && isPrint {
			compiledMessage := it.GetPrintMessage(index)

			fmt.Println(compiledMessage)
		}

		return isMatch
	}

	foundMatches := 0

	for _, item := range sortedStrings {
		if actualHashset.Has(item) {
			foundMatches++
		}
	}

	isMatch := foundMatches >= it.MatchingLength
	if !isMatch && isPrint {
		compiledMessage := it.GetPrintMessage(index)

		fmt.Println(compiledMessage)
	}

	return isMatch
}
