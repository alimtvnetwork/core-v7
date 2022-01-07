package codestack

import "gitlab.com/evatix-go/core/constants"

func StacksCountStringUsingFmt(
	formatter Formatter,
	startSkipIndex, count int,
) string {
	stacks := NewStacksDefault(
		startSkipIndex+defaultInternalSkip,
		count,
	)

	toString := stacks.JoinUsingFmt(
		formatter,
		constants.NewLineSpaceHyphenSpace)
	stacks.Dispose()

	return toString
}
