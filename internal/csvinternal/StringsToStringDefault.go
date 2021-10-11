package csvinternal

import "gitlab.com/evatix-go/core/constants"

func StringsToStringDefault(
	references ...string,
) string {
	return StringsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
