package reflectinternal

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/jsoninternal"
)

type FileWithLine struct {
	FilePath string // absolute file reflectPath
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
	return fmt.Sprintf(
		fileWithLineFormat,
		it.FilePath,
		it.Line,
	)
}

func (it FileWithLine) JsonModel() FileWithLine {
	return it
}

func (it *FileWithLine) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *FileWithLine) JsonString() string {
	s, _ := jsoninternal.Pretty.AnyTo.String(it)

	return s
}
