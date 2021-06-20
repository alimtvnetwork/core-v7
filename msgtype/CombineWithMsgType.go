package msgtype

import "gitlab.com/evatix-go/core/constants"

func CombineWithMsgType(
	genericMsg Variation,
	otherMsg string,
	reference interface{},
) string {
	if otherMsg == "" {
		return genericMsg.String() +
			constants.Space +
			ReferenceStart +
			ToValueString(reference) +
			ReferenceEnd
	}

	return genericMsg.String() +
		constants.Space +
		otherMsg +
		constants.Space +
		ReferenceStart +
		ToValueString(reference) +
		ReferenceEnd
}
