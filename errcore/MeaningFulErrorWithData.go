package errcore

import (
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
)

func MeaningfulErrorWithData(
	msgType Variation,
	funcName string,
	err error,
	data interface{},
) error {
	if err == nil {
		return nil
	}

	return msgType.Error(
		funcName,
		err.Error()+utilstringinternal.AnyToString(data))
}
