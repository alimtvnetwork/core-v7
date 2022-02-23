package csvinternal

func StringsToCsvStringsDefault(
	references ...string,
) []string {
	return StringsToCsvStrings(
		true,
		false,
		references...)
}
