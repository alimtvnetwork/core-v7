package errcore

import (
	"gitlab.com/evatix-go/core/constants"
)

func CombineWithMsgType(
	genericMsg Variation,
	otherMsg string,
	reference interface{},
) string {
	if otherMsg == "" {
		return genericMsg.String() +
			getReferenceMessage(reference)
	}

	return genericMsg.String() +
		constants.Space +
		otherMsg +
		getReferenceMessage(reference)
}
