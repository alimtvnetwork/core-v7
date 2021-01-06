package msgtype

import "gitlab.com/evatix-go/core/constants"

func CombineWithMsgType(
	genericMsg Variation,
	otherMsg string,
	reference interface{},
) string {
	return genericMsg.String() +
		constants.Space +
		otherMsg +
		constants.Space +
		ReferenceStart +
		ToValueString(reference) +
		ReferenceEnd
}
