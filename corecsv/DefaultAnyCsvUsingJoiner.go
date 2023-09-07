package corecsv

import "strings"

func DefaultAnyCsvUsingJoiner(
	joiner string,
	references ...interface{},
) string {
	csvItems := AnyItemsToCsvStrings(
		true,
		false,
		references...)

	return strings.Join(
		csvItems,
		joiner)
}
