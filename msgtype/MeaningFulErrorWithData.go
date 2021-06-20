package msgtype

import (
	"gitlab.com/evatix-go/core/internal/strutilinternal"
)

func MeaningFulErrorWithData(
	msgType Variation,
	funcName string,
	err error,
	data interface{},
) error {
	if err == nil {
		return nil
	}

	return msgType.Error(funcName, err.Error()+strutilinternal.AnyToString(data))
}
