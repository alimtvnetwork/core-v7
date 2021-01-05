package msgtype

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func Combine(genericMsg, otherMsg, reference string) string {
	return genericMsg +
		" " +
		otherMsg +
		" " +
		ReferenceStart +
		reference +
		ReferenceEnd
}

func ToValueString(reference interface{}) string {
	return fmt.Sprintf(constants.SprintPropertyNameValueFormat, reference)
}

func CombineWithMsgType(genericMsg Variation, otherMsg string, reference interface{}) string {
	return genericMsg.String() +
		" " +
		otherMsg +
		" " +
		ReferenceStart +
		ToValueString(reference) +
		ReferenceEnd
}
