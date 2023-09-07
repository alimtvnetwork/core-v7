package codestack

func PackageNameOf(fullName string) (packageName string) {
	_, packageName, _ = MethodNamePackageName(
		fullName)

	return packageName
}
