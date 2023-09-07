package codestack

func JoinPackageNameWithRelative(
	fullNameExtractPackageName, relativeNamesJoin string,
) (packageName string) {
	_, packageName, _ = MethodNamePackageName(
		fullNameExtractPackageName)

	return packageName + "." + relativeNamesJoin
}
