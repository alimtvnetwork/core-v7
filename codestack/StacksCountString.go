package codestack

func StacksCountString(
	startSkipIndex, count int,
) string {
	stacks := NewStacksDefault(
		startSkipIndex+defaultInternalSkip,
		count,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}
