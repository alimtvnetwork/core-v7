package converters

import (
	"errors"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters/coreconverted"
	"gitlab.com/evatix-go/core/defaulterr"
)

// StringsToBytesWithDefaults panic if not a number or more than 255
func StringsToBytesWithDefaults(
	strArray *[]string,
	defaultByte byte,
) *coreconverted.Bytes {
	results := make([]byte, len(*strArray))
	var errMessages []string

	for i, v := range *strArray {
		vInt, err := strconv.Atoi(v)

		if err != nil {
			msg := err.Error() +
				constants.CommaRawValueColonSpace +
				v +
				constants.CommaIndexColonSpace +
				strconv.Itoa(i)
			errMessages = append(
				errMessages,
				msg)

			results[i] = defaultByte

			continue
		}

		if vInt > constants.MaxUnit8AsInt {
			msg := defaulterr.CannotConvertStringToByte.Error() +
				constants.CommaRawValueColonSpace +
				v +
				constants.CommaIndexColonSpace +
				strconv.Itoa(i)
			errMessages = append(
				errMessages,
				msg)

			results[i] = defaultByte

			continue
		}

		results[i] = byte(vInt)
	}

	var combinedError error
	if len(errMessages) > 0 {
		errCompiledMessage := strings.Join(errMessages, constants.NewLineUnix)
		combinedError = errors.New(errCompiledMessage)
	}

	return &coreconverted.Bytes{
		Values:        &results,
		CombinedError: combinedError,
	}
}
