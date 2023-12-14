package codestack

import "gitlab.com/auk-go/core/constants"

type stacksTo struct{}

func (it stacksTo) Bytes(stackSkipIndex int) []byte {
	return New.
		StackTrace.
		DefaultCount(stackSkipIndex + defaultInternalSkip).
		StackTracesBytes()
}

func (it stacksTo) BytesDefault() []byte {
	return New.
		StackTrace.
		DefaultCount(defaultInternalSkip).
		StackTracesBytes()
}

func (it stacksTo) String(
	startSkipIndex, count int,
) string {
	stacks := it.newStacks(
		startSkipIndex+defaultInternalSkip,
		count,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}

func (it stacksTo) newStacks(startSkipIndex int, count int) TraceCollection {
	stacks := New.
		StackTrace.
		Default(startSkipIndex+defaultInternalSkip, count)

	return stacks
}

func (it stacksTo) newStacksDefault(startSkipIndex int) TraceCollection {
	stacks := New.
		StackTrace.
		DefaultCount(startSkipIndex + defaultInternalSkip)

	return stacks
}

func (it stacksTo) StringUsingFmt(
	formatter Formatter,
	startSkipIndex, count int,
) string {
	stacks := it.newStacks(
		startSkipIndex+defaultInternalSkip,
		count,
	)

	toString := stacks.JoinUsingFmt(
		formatter,
		constants.NewLineSpaceHyphenSpace,
	)
	stacks.Dispose()

	return toString
}

func (it stacksTo) JsonString(
	startSkipIndex int,
) string {
	stacks := it.newStacksDefault(
		startSkipIndex + defaultInternalSkip,
	)

	json := stacks.JsonPtr()
	stacks.Dispose()
	json.HandleError()

	return json.JsonString()
}

func (it stacksTo) JsonStringDefault() string {
	return it.JsonString(defaultInternalSkip)
}

func (it stacksTo) StringNoCount(
	startSkipIndex int,
) string {
	stacks := it.newStacksDefault(
		startSkipIndex + defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}

func (it stacksTo) StringDefault() string {
	stacks := it.newStacksDefault(
		defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}
