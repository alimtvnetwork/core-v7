package codestack

func PackageName() (packageName string) {
	_, packageName, _ = MethodNamePackageNameUsingStackSkip(
		defaultInternalSkip)

	return packageName
}
