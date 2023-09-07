package corecsv

func DefaultAnyCsvStrings(
	references ...interface{},
) []string {
	return AnyItemsToCsvStrings(
		true,
		false,
		references...)
}
