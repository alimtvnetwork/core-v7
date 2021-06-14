package reqtype

import "errors"

func RangesNotMeetError(
	message string,
	reqs ...Request,
) error {
	if len(reqs) == 0 {
		return nil
	}

	msg := RangesNotMeet(message, reqs...)

	return errors.New(msg)
}
