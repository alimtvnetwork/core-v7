package codestack

func MethodNameUsingStackSkip(stackSkipIndex int) (methodName string) {
	_, _, methodName = MethodNamePackageNameUsingStackSkip(
		stackSkipIndex + defaultInternalSkip)

	return methodName
}
