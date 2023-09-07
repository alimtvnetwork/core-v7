package errcore

import (
	"gitlab.com/auk-go/core/constants"
)

func Combine(
	genericMsg,
	otherMsg,
	reference string,
) string {
	return genericMsg +
		constants.Space +
		otherMsg +
		constants.Space +
		ReferenceStart +
		reference +
		ReferenceEnd
}
