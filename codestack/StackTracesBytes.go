package codestack

func StackTracesBytes(stackSkipIndex int) []byte {
	return NewStacksDefaultCount(stackSkipIndex + defaultInternalSkip).
		StackTracesBytes()
}

func StackTracesBytesDefault() []byte {
	return NewStacksDefaultCount(defaultInternalSkip).
		StackTracesBytes()
}
