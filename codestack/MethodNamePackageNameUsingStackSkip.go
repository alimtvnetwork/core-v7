package codestack

import "runtime"

func MethodNamePackageNameUsingStackSkip(stackSkipIndex int) (fullMethodName, packageName, methodName string) {
	pc, _, _, _ := runtime.Caller(stackSkipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	return MethodNamePackageName(fullFuncName)
}
