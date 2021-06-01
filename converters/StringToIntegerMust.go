package converters

func StringToIntegerMust(
	input string,
) (value int) {
	value, err := StringToInteger(input)

	if err != nil {
		panic(err)
	}

	return value
}
