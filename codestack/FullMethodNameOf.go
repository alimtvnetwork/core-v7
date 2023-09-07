package codestack

func FullMethodNameOf(fullName string) (packageName string) {
	fullMethodNameOf, _, _ := MethodNamePackageName(
		fullName)

	return fullMethodNameOf
}
