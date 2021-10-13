package errcore

import (
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
)

func MeaningfulErrorWithData(
	rawErrType RawErrorType,
	funcName string,
	err error,
	data interface{},
) error {
	if err == nil {
		return nil
	}

	return rawErrType.Error(
		funcName,
		err.Error()+utilstringinternal.AnyToString(data))
}
