package converters

import "strings"

func StringToIntegersConditional(
	stringInput,
	separator string,
	processor func(in string) (out int, isTake, isBreak bool),
) *[]int {
	if stringInput == "" {
		return &[]int{}
	}

	splits := strings.Split(stringInput, separator)
	results := make([]int, 0, len(splits))

	for _, v := range splits {
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
