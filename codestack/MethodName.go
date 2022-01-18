package codestack

func MethodName() (methodName string) {
	_, _, methodName = MethodNamePackageNameUsingStackSkip(defaultInternalSkip)

	return methodName
}
