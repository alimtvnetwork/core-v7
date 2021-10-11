package csvinternal

import "gitlab.com/evatix-go/core/constants"

func AnyItemsToStringDefault(
	references ...interface{},
) string {
	return AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
