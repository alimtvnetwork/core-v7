package codestack

func MethodNameOf(fullName string) (packageName string) {
	_, _, methodName := MethodNamePackageName(
		fullName)

	return methodName
}
