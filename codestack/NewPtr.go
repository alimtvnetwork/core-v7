package codestack

import (
	"runtime"
)

func NewPtr(skipIndex int) *Trace {
	pc, file, line, isOkay := runtime.Caller(skipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	packageName, methodName := MethodNamePackageName(fullFuncName)

	return &Trace{
		SkipIndex:         skipIndex,
		PackageName:       packageName,
		MethodName:        methodName,
		PackageMethodName: fullFuncName,
		FileName:          file,
		Line:              line,
		IsOkay:            isOkay,
	}
}
