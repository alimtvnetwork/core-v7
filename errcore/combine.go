package errcore

import (
	"gitlab.com/evatix-go/core/constants"
)

func Combine(genericMsg, otherMsg, reference string) string {
	return genericMsg +
		constants.Space +
		otherMsg +
		constants.Space +
		ReferenceStart +
		reference +
		ReferenceEnd
}
