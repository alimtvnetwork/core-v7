package converters

// StringsToBytesConditional only take if isTake returns true, breaks and exits if isBreak to true
func StringsToBytesConditional(
	strArray *[]string,
	processor func(in string) (out byte, isTake, isBreak bool),
) *[]byte {
	results := make([]byte, 0, len(*strArray))

	for _, v := range *strArray {
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
