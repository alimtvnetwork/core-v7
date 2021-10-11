package reqtype

import "gitlab.com/evatix-go/core/errcore"

func RangesOnlySupportedFor(
	message string,
	requests ...Request,
) error {
	if len(requests) == 0 {
		return nil
	}

	referencesMessage := RangesStringDefaultJoiner(
		requests...)

	return errcore.RangesOnlySupported.Error(
		message,
		referencesMessage)
}
