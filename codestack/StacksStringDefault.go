package codestack

func StacksStringDefault() string {
	stacks := NewStacksDefaultCount(
		defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}
