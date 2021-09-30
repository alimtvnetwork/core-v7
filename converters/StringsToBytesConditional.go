package converters

// StringsToBytesConditional only take if isTake returns true, breaks and exits if isBreak to true
func StringsToBytesConditional(
	processor func(in string) (out byte, isTake, isBreak bool),
	stringsSlice []string,
) *[]byte {
	results := make([]byte, 0, len(stringsSlice))

	for _, v := range stringsSlice {
		out, isTake, isBreak := processor(v)

		if isTake {
			results = append(results, out)
		}

		if isBreak {
			break
		}
	}

	return &results
}
