package corecsv

import "gitlab.com/evatix-go/core/constants"

func DefaultAnyCsv(
	references ...interface{},
) string {
	return AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
