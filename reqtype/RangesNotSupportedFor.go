package reqtype

import (
	"gitlab.com/evatix-go/core/errcore"
)

func RangesNotSupportedFor(
	message string,
	requests ...Request,
) error {
	if len(requests) == 0 {
		return nil
	}

	referencesMessage := RangesStringDefaultJoiner(
		requests...)

	return errcore.NotSupported.Error(
		message,
		referencesMessage)
}
