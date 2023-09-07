package codestack

func StacksJsonString(
	startSkipIndex int,
) string {
	stacks := NewStacksDefaultCount(
		startSkipIndex + defaultInternalSkip,
	)

	json := stacks.JsonPtr()
	stacks.Dispose()
	json.HandleError()

	return json.JsonString()
}
