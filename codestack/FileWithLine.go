package codestack

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type FileWithLine struct {
	FilePath string // absolute file path
	Line     int    // line number
}

func (it *FileWithLine) FullFilePath() string {
	return it.FilePath
}

func (it *FileWithLine) LineNumber() int {
	return it.Line
}

func (it *FileWithLine) IsNil() bool {
	return it == nil
}

func (it *FileWithLine) IsNotNil() bool {
	return it != nil
}

func (it *FileWithLine) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.FileWithLine()
}

func (it FileWithLine) StringUsingFmt(formatterFunc func(fileWithLine FileWithLine) string) string {
	return formatterFunc(it)
}

func (it *FileWithLine) FileWithLine() string {
	return fmt.Sprintf(fileWithLineFormat,
		it.FilePath,
		it.Line)
}

func (it FileWithLine) JsonModel() FileWithLine {
	return it
}

func (it *FileWithLine) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *FileWithLine) JsonString() string {
	jsonResult := it.Json()

	return jsonResult.JsonString()
}

func (it FileWithLine) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it FileWithLine) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

func (it *FileWithLine) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*FileWithLine, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *FileWithLine) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *FileWithLine {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *FileWithLine) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *FileWithLine) AsFileLiner() FileWithLiner {
	return it
}
