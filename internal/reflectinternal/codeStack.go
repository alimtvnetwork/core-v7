package reflectinternal

import (
	"fmt"
	"runtime"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type codeStack struct{}

func (it codeStack) New(skipStack int) StackTrace {
	pc, file, line, isOkay := runtime.Caller(skipStack + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	fullMethodSignature, packageName, methodName := GetFunc.All(fullFuncName)

	return StackTrace{
		SkipIndex:         skipStack,
		PackageName:       packageName,
		MethodName:        methodName,
		PackageMethodName: fullMethodSignature,
		FilePath:          file,
		Line:              line,
		IsOkay:            isOkay,
	}
}

func (it codeStack) NewDefault() StackTrace {
	return it.New(defaultInternalSkip)
}

func (it codeStack) MethodName(skipStack int) string {
	pc, _, _, _ := runtime.Caller(skipStack + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	_, _, methodName := GetFunc.All(fullFuncName)

	return methodName
}

func (it codeStack) MethodNameWithLine(skipStack int) string {
	pc, _, line, _ := runtime.Caller(skipStack + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	_, _, methodName := GetFunc.All(fullFuncName)

	return fmt.Sprintf(
		"%s:%d",
		methodName,
		line,
	)
}

func (it codeStack) FileWithLine(skipStack int) string {
	pc, file, line, _ := runtime.Caller(skipStack + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	_, _, methodName := GetFunc.All(fullFuncName)

	return fmt.Sprintf(
		shortStringFormat,
		methodName,
		line,
		file,
		line,
	)
}

func (it codeStack) LastFileWithLines(skipStack, count int) []string {
	lines := make([]string, 0, count)

	for i := 0; i < count; i++ {
		lines = append(lines, it.FileWithLine(skipStack+i))
	}

	return lines
}

func (it codeStack) LastFileWithLine(skipStack, count int) string {
	lines := it.LastFileWithLines(defaultInternalSkip+skipStack, count)

	return strings.Join(lines, constants.NewLineUnix)
}

func (it codeStack) FilePath(skipStack int) string {
	_, file, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return file
	}

	return constants.EmptyString
}

func (it codeStack) NewFileWithLines(skipStack, count int) []FileWithLine {
	lines := make([]FileWithLine, 0, count)

	for i := 0; i < count; i++ {
		_, file, line, isOkay := runtime.Caller(skipStack + defaultInternalSkip + i)

		if !isOkay {
			return lines
		}

		f := FileWithLine{
			FilePath: file,
			Line:     line,
		}

		lines = append(
			lines,
			f,
		)
	}

	return lines
}

func (it codeStack) NewFileWithLine(skipStack int) FileWithLine {
	_, file, line, _ := runtime.Caller(skipStack + defaultInternalSkip)

	return FileWithLine{
		FilePath: file,
		Line:     line,
	}
}

func (it codeStack) NewStacks(skipStack, count int) []StackTrace {
	slice := make([]StackTrace, 0, count)

	for i := 0; i < count; i++ {
		f := it.New(defaultInternalSkip + skipStack)

		slice = append(
			slice,
			f,
		)
	}

	return slice
}

func (it codeStack) StacksStrings(skipStack int) []string {
	return it.StacksStringsCount(
		skipStack+defaultInternalSkip,
		defaultStackCount,
	)
}

func (it codeStack) StacksStringsCount(skipStack, count int) []string {
	fileWithLines := it.NewFileWithLines(
		skipStack,
		count,
	)

	lines := make([]string, 0, len(fileWithLines))

	for _, fileWithLine := range fileWithLines {
		newLine := fmt.Sprintf(
			fileWithLineFormat,
			fileWithLine.FilePath,
			fileWithLine.Line,
		)

		lines = append(lines, newLine)
	}

	return lines
}

func (it codeStack) StacksString(skipStack int) string {
	lines := it.StacksStrings(skipStack + defaultInternalSkip)

	return strings.Join(lines, constants.NewLineUnix)
}

func (it codeStack) StacksStringDefault(skipStack int) string {
	return it.StacksStringCount(
		skipStack+defaultInternalSkip,
		defaultStackCount,
	)
}

func (it codeStack) StacksStringCount(skipStack, count int) string {
	lines := it.StacksStringsCount(skipStack+defaultInternalSkip, count)

	joinedLines := strings.Join(lines, "\n  - ")

	return fmt.Sprintf("Stack-Trace:\n  - %s", joinedLines)
}

func (it codeStack) SingleStack(skipStack int) string {
	lines := it.StacksStringsCount(skipStack+defaultInternalSkip, 1)

	if len(lines) > 0 {
		return lines[0]
	}

	return ""
}
