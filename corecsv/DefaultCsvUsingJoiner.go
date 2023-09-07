package corecsv

func DefaultCsvUsingJoiner(
	joiner string,
	references ...string,
) string {
	return StringsToCsvString(
		joiner,
		true,
		false,
		references...)
}
