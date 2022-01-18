package codestack

func PackageNameUsingStackSkip(stackSkipIndex int) (packageName string) {
	_, packageName, _ = MethodNamePackageNameUsingStackSkip(
		stackSkipIndex + defaultInternalSkip)

	return packageName
}
