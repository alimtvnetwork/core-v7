package codestack

import (
	"runtime"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

func New(skipIndex int) Trace {
	pc, file, line, isOkay := runtime.Caller(skipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	splitsByDot := strings.Split(fullFuncName, constants.Dot)
	first, last := stringslice.FirstLastDefault(splitsByDot)

	return Trace{
		SkipIndex:         skipIndex,
		PackageName:       first,
		MethodName:        last,
		PackageMethodName: fullFuncName,
		FileName:          file,
		Line:              line,
		IsOkay:            isOkay,
	}
}
