package corecsv

func StringsToCsvStringsDefault(
	references ...string,
) []string {
	return StringsToCsvStrings(
		true,
		false,
		references...)
}
