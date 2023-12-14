package codestack

import "runtime"

type newCreator struct {
	traces     newTraceCollection
	StackTrace newStacksCreator
}

func (it newCreator) Default() Trace {
	return it.Create(defaultInternalSkip)
}

func (it newCreator) SkipOne() Trace {
	return it.Create(Skip2)
}

func (it newCreator) Ptr(skipIndex int) *Trace {
	trace := it.Create(skipIndex + defaultInternalSkip)

	return &trace
}

func (it newCreator) Create(skipIndex int) Trace {
	pc, file, line, isOkay := runtime.Caller(skipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	fullMethodSignature, packageName, methodName := NameOf.All(fullFuncName)

	return Trace{
		SkipIndex:         skipIndex,
		PackageName:       packageName,
		MethodName:        methodName,
		PackageMethodName: fullMethodSignature,
		FilePath:          file,
		Line:              line,
		IsOkay:            isOkay,
	}
}
