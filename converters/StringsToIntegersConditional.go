package converters

// StringsToIntegersConditional handle converts from processor func
func StringsToIntegersConditional(
	strArray *[]string,
	processor func(in string) (out int, isTake, isBreak bool),
) *[]int {
	results := make([]int, 0, len(*strArray))

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
