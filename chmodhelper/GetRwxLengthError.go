package chmodhelper

import "gitlab.com/evatix-go/core/msgtype"

func GetRwxLengthError(rwx string) error {
	length := len(rwx)

	if length != SupportedLength {
		return msgtype.LengthShouldBeEqualToMessage.
			Error(
				"rwx length should be "+SupportedLengthString,
				length)
	}

	return nil
}
