package reflectinternal

import (
	"fmt"
	"path/filepath"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/jsoninternal"
)

type StackTrace struct {
	SkipIndex int
	PackageName,
	MethodName,
	PackageMethodName string
	FilePath string
	Line     int
	IsOkay   bool
}

func (it StackTrace) Message() string {
	return it.getCompiledMessage()
}

// ShortString
//
// Returns lazy or at once "Method (LineNumber) -> FileFullPath:LineNumber"
// using format shortStringFormat "%s (%d) -> %s:%d"
//
// Format :
//   - https://prnt.sc/25ypcyc : "%s (%d) -> %s:%d"
//
// Example :
//   - "Method (LineNumber) -> FileFullPath:LineNumber"
func (it StackTrace) ShortString() string {
	shortString := fmt.Sprintf(
		shortStringFormat,
		it.PackageMethodName,
		it.Line,
		it.FilePath,
		it.Line,
	)

	return shortString
}

func (it *StackTrace) IsNil() bool {
	return it == nil
}

func (it *StackTrace) HasIssues() bool {
	return it == nil || !it.IsOkay || it.PackageMethodName == "" || it.PackageName == ""
}

func (it *StackTrace) IsNotNil() bool {
	return it != nil
}

func (it *StackTrace) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.Message()
}

func (it StackTrace) StringUsingFmt(formatterFunc func(trace StackTrace) string) string {
	return formatterFunc(it)
}

func (it StackTrace) FileWithLine() FileWithLine {
	return FileWithLine{
		FilePath: it.FilePath,
		Line:     it.Line,
	}
}

// FullFilePath
//
// Returns the full file reflectPath
func (it StackTrace) FullFilePath() string {
	return it.FilePath
}

// FileName
//
// Returns the file name only
func (it StackTrace) FileName() string {
	return filepath.Base(it.FilePath)
}

func (it StackTrace) LineNumber() int {
	return it.Line
}

// FileWithLineString
//
// Format :
//   - https://prnt.sc/25yorfh : "%s:%d"
//
// Example :
//   - "FilePath:LineNumber"
func (it StackTrace) FileWithLineString() string {
	return fmt.Sprintf(
		fileWithLineFormat,
		it.FilePath,
		it.Line,
	)
}

func (it StackTrace) getCompiledMessage() string {
	message := fmt.Sprintf(
		funcPrintFormat,
		it.PackageMethodName,
		it.Line,
		it.FilePath,
		it.Line,
	)

	return message
}

func (it StackTrace) JsonModel() StackTrace {
	return it
}

func (it StackTrace) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *StackTrace) Dispose() {
	if it == nil {
		return
	}

	it.SkipIndex = 0
	it.PackageName = ""
	it.MethodName = ""
	it.PackageMethodName = ""
	it.FilePath = ""
	it.Line = 0
	it.IsOkay = false
}

func (it StackTrace) JsonString() string {
	s, _ := jsoninternal.Pretty.AnyTo.String(it)

	return s
}

func (it StackTrace) Clone() StackTrace {
	return StackTrace{
		SkipIndex:         it.SkipIndex,
		PackageName:       it.PackageName,
		MethodName:        it.MethodName,
		PackageMethodName: it.PackageMethodName,
		FilePath:          it.FilePath,
		Line:              it.Line,
		IsOkay:            it.IsOkay,
	}
}

func (it *StackTrace) ClonePtr() *StackTrace {
	if it == nil {
		return nil
	}

	trace := it.Clone()

	return &trace
}
