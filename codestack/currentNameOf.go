package codestack

import (
	"runtime"
)

type currentNameOf struct{}

func (it currentNameOf) Method() (methodName string) {
	_, _, methodName = it.AllStackSkip(defaultInternalSkip)

	return methodName
}

func (it currentNameOf) MethodByFullName(fullName string) (packageName string) {
	_, _, methodName := it.All(
		fullName,
	)

	return methodName
}

func (it currentNameOf) All(fullFuncName string) (fullMethodName, packageName, methodName string) {
	if fullFuncName == "" {
		return "", "", ""
	}

	return getFuncEverything(fullFuncName)
}

func (it currentNameOf) AllStackSkip(stackSkipIndex int) (fullMethodName, packageName, methodName string) {
	pc, _, _, _ := runtime.Caller(stackSkipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	return it.All(fullFuncName)
}
func (it currentNameOf) MethodStackSkip(stackSkipIndex int) (methodName string) {
	_, _, methodName = it.AllStackSkip(
		stackSkipIndex + defaultInternalSkip,
	)

	return methodName
}

func (it currentNameOf) JoinPackageNameWithRelative(
	fullNameExtractPackageName, relativeNamesJoin string,
) (packageName string) {
	_, packageName, _ = it.All(
		fullNameExtractPackageName,
	)

	return packageName + "." + relativeNamesJoin
}

func (it currentNameOf) Package() (packageName string) {
	_, packageName, _ = it.AllStackSkip(
		defaultInternalSkip,
	)

	return packageName
}

func (it currentNameOf) PackageByFullName(fullName string) (packageName string) {
	_, packageName, _ = it.All(
		fullName,
	)

	return packageName
}

func (it currentNameOf) PackageStackSkip(stackSkipIndex int) (packageName string) {
	_, packageName, _ = it.AllStackSkip(
		stackSkipIndex + defaultInternalSkip,
	)

	return packageName
}

func (it currentNameOf) CurrentFuncFullPath(fullName string) (packageName string) {
	fullMethodNameOf, _, _ := NameOf.All(
		fullName,
	)

	return fullMethodNameOf
}
