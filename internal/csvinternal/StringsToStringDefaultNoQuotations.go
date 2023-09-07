package csvinternal

import "gitlab.com/auk-go/core/constants"

func StringsToStringDefaultNoQuotations(
	references ...string,
) string {
	return StringsToCsvString(
		constants.CommaSpace,
		false,
		false,
		references...)
}
