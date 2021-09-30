package converters

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters/coreconverted"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/msgtype"
)

// StringsToBytesWithDefaults panic if not a number or more than 255
func StringsToBytesWithDefaults(
	defaultByte byte,
	stringsSlice ...string,
) *coreconverted.Bytes {
	results := make([]byte, len(stringsSlice))
	var sliceErr []string

	for i, v := range stringsSlice {
		vInt, err := strconv.Atoi(v)

		if err != nil {
			msg := err.Error() +
				constants.CommaRawValueColonSpace +
				v +
				constants.CommaIndexColonSpace +
				strconv.Itoa(i)
			sliceErr = append(
				sliceErr,
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
			sliceErr = append(
				sliceErr,
				msg)

			results[i] = defaultByte

			continue
		}

		results[i] = byte(vInt)
	}

	return &coreconverted.Bytes{
		Values:        results,
		CombinedError: msgtype.SliceToError(sliceErr),
	}
}
