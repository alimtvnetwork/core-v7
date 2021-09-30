package codestack

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

type Trace struct {
	SkipIndex int
	PackageName,
	MethodName,
	PackageMethodName string
	FileName string
	Line     int
	IsOkay   bool
	message  corestr.SimpleStringOnce
}

func (it *Trace) Message() string {
	return it.
		message.
		GetPlusSetOnUninitializedFunc(
			it.getCompiledMessage)
}

func (it *Trace) IsNil() bool {
	return it == nil
}

func (it *Trace) HasIssues() bool {
	return it == nil || !it.IsOkay || it.PackageMethodName == ""
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
		FileName: it.FileName,
		Line:     it.Line,
	}
}

func (it *Trace) FileWithLineString() string {
	return fmt.Sprintf(fileWithLineFormat,
		it.FileName,
		it.Line)
}

func (it *Trace) getCompiledMessage() string {
	message := fmt.Sprintf(funcPrintFormat,
		it.PackageMethodName,
		it.Line,
		it.FileName,
		it.Line)

	return message
}

func (it Trace) JsonModel() Trace {
	return it
}

func (it *Trace) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Trace) JsonString() string {
	jsonResult := it.Json()

	return jsonResult.JsonString()
}

func (it Trace) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it Trace) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

func (it *Trace) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Trace, error) {
	err := jsonResult.Unmarshal(&it)

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
		FileName:          it.FileName,
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
