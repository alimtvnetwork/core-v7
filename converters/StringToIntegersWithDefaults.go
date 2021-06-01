package converters

import (
	"errors"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters/coreconverted"
)

func StringToIntegersWithDefaults(
	stringInput,
	separator string,
	defaultInt int,
) *coreconverted.Integers {
	if stringInput == "" {
		return &coreconverted.Integers{
			Values:        &[]int{},
			CombinedError: nil,
		}
	}

	splits := strings.Split(stringInput, separator)
	results := make([]int, len(splits))
	var errMessages []string

	for i, v := range splits {
		vInt, err := strconv.Atoi(v)

		if err != nil {
			results[i] = defaultInt
			errMessage := constants.IndexColonSpace +
				strconv.Itoa(i) +
				err.Error()
			errMessages = append(
				errMessages,
				errMessage)

			continue
		}

		results[i] = vInt
	}

	var combinedError error
	if len(errMessages) > 0 {
		errCompiledMessage := strings.Join(errMessages, constants.NewLineUnix)
		combinedError = errors.New(errCompiledMessage)
	}

	return &coreconverted.Integers{
		Values:        &results,
		CombinedError: combinedError,
	}
}
