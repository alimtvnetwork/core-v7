package corecsv

func DefaultCsvStrings(
	references ...string,
) []string {
	return StringsToCsvStrings(
		true,
		false,
		references...)
}
