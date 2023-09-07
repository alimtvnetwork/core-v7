package codestack

import (
	"fmt"
	"path/filepath"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
)

type Trace struct {
	SkipIndex int
	PackageName,
	MethodName,
	PackageMethodName string
	FilePath    string
	Line        int
	IsOkay      bool
	message     corestr.SimpleStringOnce
	shortString corestr.SimpleStringOnce
}

func (it *Trace) Message() string {
	if it.message.IsInitialized() {
		return it.message.String()
	}

	return it.
		message.
		GetPlusSetOnUninitializedFunc(
			it.getCompiledMessage)
}

// ShortString
//
// Returns lazy or at once "Method (LineNumber) -> FileFullPath:LineNumber"
// using format shortStringFormat "%s (%d) -> %s:%d"
//
// Format :
//  - https://prnt.sc/25ypcyc : "%s (%d) -> %s:%d"
//
// Example :
//  - "Method (LineNumber) -> FileFullPath:LineNumber"
func (it *Trace) ShortString() string {
	if it.shortString.IsInitialized() {
		return it.shortString.String()
	}

	shortString := fmt.Sprintf(
		shortStringFormat,
		it.PackageMethodName,
		it.Line,
		it.FilePath,
		it.Line)

	return it.
		shortString.
		GetPlusSetOnUninitialized(
			shortString)
}

func (it *Trace) IsNil() bool {
	return it == nil
}

func (it *Trace) HasIssues() bool {
	return it == nil || !it.IsOkay || it.PackageMethodName == "" || it.PackageName == ""
}

func (it *Trace) IsNotNil() bool {
	return it != nil
}

func (it *Trace) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.Message()
}

func (it Trace) StringUsingFmt(formatterFunc func(trace Trace) string) string {
	return formatterFunc(it)
}

func (it *Trace) FileWithLine() FileWithLine {
	return FileWithLine{
		FilePath: it.FilePath,
		Line:     it.Line,
	}
}

// FullFilePath
//
// Returns the full file path
func (it *Trace) FullFilePath() string {
	return it.FilePath
}

// FileName
//
// Returns the file name only
func (it *Trace) FileName() string {
	return filepath.Base(it.FilePath)
}

func (it *Trace) LineNumber() int {
	return it.Line
}

// FileWithLineString
//
// Format :
//  - https://prnt.sc/25yorfh : "%s:%d"
//
// Example :
//  - "FilePath:LineNumber"
func (it *Trace) FileWithLineString() string {
	return fmt.Sprintf(
		fileWithLineFormat,
		it.FilePath,
		it.Line)
}

func (it *Trace) getCompiledMessage() string {
	message := fmt.Sprintf(funcPrintFormat,
		it.PackageMethodName,
		it.Line,
		it.FilePath,
		it.Line)

	return message
}

func (it Trace) JsonModel() Trace {
	return it
}

func (it *Trace) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Trace) Dispose() {
	if it == nil {
		return
	}

	it.SkipIndex = constants.Zero
	it.PackageName = constants.EmptyString
	it.MethodName = constants.EmptyString
	it.PackageMethodName = constants.EmptyString
	it.FilePath = constants.EmptyString
	it.Line = constants.Zero
	it.IsOkay = false
	it.message.Dispose()
	it.shortString.Dispose()
}

func (it *Trace) JsonString() string {
	jsonResult := it.Json()

	return jsonResult.JsonString()
}

func (it Trace) Json() corejson.Result {
	return corejson.New(it)
}

func (it Trace) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Trace) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Trace, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *Trace) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Trace {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Trace) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it Trace) Clone() Trace {
	return Trace{
		SkipIndex:         it.SkipIndex,
		PackageName:       it.PackageName,
		MethodName:        it.MethodName,
		PackageMethodName: it.PackageMethodName,
		FilePath:          it.FilePath,
		Line:              it.Line,
		IsOkay:            it.IsOkay,
	}
}

func (it *Trace) ClonePtr() *Trace {
	if it == nil {
		return nil
	}

	trace := it.Clone()

	return &trace
}

func (it *Trace) AsFileLiner() FileWithLiner {
	return it
}
