package corecsv

import "gitlab.com/evatix-go/core/constants"

func DefaultCsv(
	references ...string,
) string {
	return StringsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
