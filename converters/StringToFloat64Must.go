package converters

func StringToFloat64Must(input string) float64 {
	value, err2 := StringToFloat64(input)

	if err2 != nil {
		panic(err2)
	}

	return value
}
