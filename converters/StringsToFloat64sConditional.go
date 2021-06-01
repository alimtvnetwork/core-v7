package converters

// StringsToFloat64sConditional handle convert from processor function either throw or ignore
func StringsToFloat64sConditional(
	strArray *[]string,
	processor func(in string) (out float64, isTake, isBreak bool),
) *[]float64 {
	results := make([]float64, 0, len(*strArray))

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
