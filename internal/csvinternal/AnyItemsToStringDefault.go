package csvinternal

import "gitlab.com/auk-go/core/constants"

func AnyItemsToStringDefault(
	references ...interface{},
) string {
	return AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
