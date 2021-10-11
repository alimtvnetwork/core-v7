package codestack

func StacksString(
	startSkipIndex int,
) string {
	stacks := NewStacksDefaultCount(
		startSkipIndex + defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}
