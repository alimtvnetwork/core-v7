package converters

import (
	"errors"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters/coreconverted"
)

// StringsToIntegersWithDefaults On fail use the default int
func StringsToIntegersWithDefaults(
	defaultInt int,
	strArray ...string,
) *coreconverted.Integers {
	results := make([]int, 0, len(strArray))
	var errMessages []string

	for i, v := range strArray {
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
		Values:        results,
		CombinedError: combinedError,
	}
}
