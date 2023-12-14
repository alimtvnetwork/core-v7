package converters

import (
	"errors"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/constants/bitsize"
	"gitlab.com/auk-go/core/converters/coreconverted"
	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/strutilinternal"
	"gitlab.com/auk-go/core/simplewrap"
)

type stringsTo struct{}

func (it stringsTo) Hashset(
	lines []string,
) map[string]bool {
	length := len(lines)
	hashset := make(map[string]bool, length)

	for _, s := range lines {
		hashset[s] = true
	}

	return hashset
}

func (it stringsTo) HashmapTrimColon(
	lines ...string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(lines).
		LineSplitMapOptions(
			true,
			constants.Colon,
		)
}

func (it stringsTo) HashmapTrimHyphen(
	lines ...string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(lines).
		LineSplitMapOptions(
			true,
			constants.Hyphen,
		)
}

func (it stringsTo) HashmapOptions(
	isTrim bool,
	splitter string,
	lines ...string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(lines).
		LineSplitMapOptions(
			isTrim,
			splitter,
		)
}

func (it stringsTo) HashmapTrim(
	splitter string,
	lines []string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(lines).
		LineSplitMapTrim(splitter)
}

// HashmapUsingFuncOptions
//
//	Skips if empty after trim
func (it stringsTo) HashmapUsingFuncOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key, val string),
	lines ...string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(lines).
		LineProcessorMapOptions(
			isTrimBefore,
			processorFunc,
		)
}

// HashmapUsingFuncTrim
//
//	Skips if empty after trim
func (it stringsTo) HashmapUsingFuncTrim(
	processorFunc func(line string) (key, val string),
	lines ...string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(lines).
		LineProcessorMapOptions(
			true,
			processorFunc,
		)
}

// MapStringIntegerUsingFunc
//
//	Skips if empty after trim
func (it stringsTo) MapStringIntegerUsingFunc(
	isTrimBefore bool,
	processorFunc func(line string) (key string, val int),
	lines ...string,
) map[string]int {
	return strutilinternal.
		SliceToMapConverter(lines).
		LineProcessorMapStringIntegerOptions(
			isTrimBefore,
			processorFunc,
		)
}

// MapStringAnyUsingFunc
//
//	Skips if empty after trim
func (it stringsTo) MapStringAnyUsingFunc(
	isTrimBefore bool,
	processorFunc func(line string) (key string, val interface{}),
	lines ...string,
) map[string]interface{} {
	return strutilinternal.
		SliceToMapConverter(lines).
		LineProcessorMapStringAnyOptions(
			isTrimBefore,
			processorFunc,
		)
}

func (it stringsTo) MapConverter(
	lines ...string,
) StringsToMapConverter {
	return lines
}

// PointerStrings
//
//	Will give empty or converted results array (not nil)
//	It doesn't copy but points to same string address in the array
//
//	Example code : https://play.golang.org/p/_OkY82E2kO9
func (it stringsTo) PointerStrings(pointerToStrings *[]string) *[]*string {
	if pointerToStrings == nil || *pointerToStrings == nil {
		var emptyResult []*string

		return &emptyResult
	}

	newArray := make([]*string, len(*pointerToStrings))

	for i := range *pointerToStrings {
		// direct access important here.
		newArray[i] = &(*pointerToStrings)[i]
	}

	return &newArray
}

// PointerStringsCopy
//
//	will give empty or converted results array (not nil)
//	Copy each item to the new array
func (it stringsTo) PointerStringsCopy(pointerToStrings *[]string) *[]*string {
	if pointerToStrings == nil || *pointerToStrings == nil {
		var emptyResult []*string

		return &emptyResult
	}

	newArray := make([]*string, len(*pointerToStrings))

	for i, value := range *pointerToStrings {
		// here copy is important
		valueCopy := value
		newArray[i] = &valueCopy
	}

	return &newArray
}

// IntegersConditional handle converts from processor func
func (it stringsTo) IntegersConditional(
	processor func(in string) (out int, isTake, isBreak bool),
	lines ...string,
) []int {
	results := make([]int, 0, len(lines))

	for _, v := range lines {
		out, isTake, isBreak := processor(v)

		if isTake {
			results = append(results, out)
		}

		if isBreak {
			break
		}
	}

	return results
}

// IntegersWithDefaults On fail use the default int
func (it stringsTo) IntegersWithDefaults(
	defaultInt int,
	lines ...string,
) *coreconverted.Integers {
	results := make([]int, 0, len(lines))
	var errMessages []string

	for i, v := range lines {
		vInt, err := strconv.Atoi(v)

		if err != nil {
			results[i] = defaultInt
			errMessage := constants.IndexColonSpace +
				strconv.Itoa(i) +
				err.Error()
			errMessages = append(
				errMessages,
				errMessage,
			)

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

// IntegersOptionPanic
//
//	panic if not a number
func (it stringsTo) IntegersOptionPanic(
	isPanic bool,
	lines ...string,
) []int {
	results := make([]int, len(lines))

	for i, v := range lines {
		vInt, err := strconv.Atoi(v)

		if isPanic && err != nil {
			panic(err)
		} else if err != nil {
			continue
		}

		results[i] = vInt
	}

	return results
}

// IntegersSkipErrors
//
//	no errors captured.
func (it stringsTo) IntegersSkipErrors(
	lines ...string,
) []int {
	return it.IntegersOptionPanic(
		false,
		lines...,
	)
}

func (it stringsTo) IntegersSkipMapAndDefaultValue(
	defaultVal int,
	skipValues map[string]bool,
	lines ...string,
) []int {
	results := make([]int, len(lines))

	for i, v := range lines {
		if skipValues[v] {
			continue
		}

		vInt, err := strconv.Atoi(strings.TrimSpace(v))

		if err != nil {
			results[i] = defaultVal
			continue
		}

		results[i] = vInt
	}

	return results
}

func (it stringsTo) IntegersSkipAndDefaultValue(
	defaultVal int,
	skipValue string,
	lines ...string,
) []int {
	results := make([]int, len(lines))

	for i, v := range lines {
		if skipValue == v {
			continue
		}

		vInt, err := strconv.Atoi(v)

		if err != nil {
			results[i] = defaultVal
			continue
		}

		results[i] = vInt
	}

	return results
}

// BytesConditional only take if isTake returns true, breaks and exits if isBreak to true
func (it stringsTo) BytesConditional(
	processor func(in string) (out byte, isTake, isBreak bool),
	stringsSlice []string,
) []byte {
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

	return results
}

// BytesWithDefaults
//
//	panic if not a number or more than 255
func (it stringsTo) BytesWithDefaults(
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
				msg,
			)

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
				msg,
			)

			results[i] = defaultByte

			continue
		}

		results[i] = byte(vInt)
	}

	return &coreconverted.Bytes{
		Values:        results,
		CombinedError: errcore.SliceToError(sliceErr),
	}
}

func (it stringsTo) Csv(isSkipQuoteOnlyOnExistence bool, stringsSlice ...string) string {
	csvLines := simplewrap.DoubleQuoteWrapElements(
		isSkipQuoteOnlyOnExistence,
		stringsSlice...,
	)

	return strings.Join(csvLines, constants.Comma)
}

func (it stringsTo) CsvUsingPtrStrings(isSkipQuoteOnlyOnExistence bool, stringsSlice *[]string) string {
	if stringsSlice == nil {
		return ""
	}

	csvLines := simplewrap.DoubleQuoteWrapElements(
		isSkipQuoteOnlyOnExistence,
		*stringsSlice...,
	)

	return strings.Join(csvLines, constants.Comma)
}

func (it stringsTo) CsvWithIndexes(lines []string) string {
	csvLines := simplewrap.DoubleQuoteWrapElementsWithIndexes(
		lines...,
	)

	return strings.Join(csvLines, constants.Comma)
}

// BytesMust
//
//	panic if not a number or more than 255 or less than 0
func (it stringsTo) BytesMust(lines ...string) []byte {
	results := make([]byte, len(lines))

	for i, v := range lines {
		vInt, err := StringTo.Byte(v)

		if err != nil {
			panic(err)
		}

		results[i] = vInt
	}

	return results
}

// Float64sMust
//
//	panic if not a number
func (it stringsTo) Float64sMust(lines ...string) []float64 {
	results := make([]float64, len(lines))

	for i, v := range lines {
		vFloat, err := strconv.ParseFloat(v, bitsize.Of64)

		if err != nil {
			panic(err)
		}

		results[i] = vFloat
	}

	return results
}

// Float64sConditional
//
//	handle convert from processor function either throw or ignore
func (it stringsTo) Float64sConditional(
	processor func(in string) (out float64, isTake, isBreak bool),
	lines []string,
) []float64 {
	results := make([]float64, 0, len(lines))

	for _, v := range lines {
		out, isTake, isBreak := processor(v)

		if isTake {
			results = append(results, out)
		}

		if isBreak {
			break
		}
	}

	return results
}

// PtrOfPtrToPtrStrings will give empty or converted results array (not nil)
func (it stringsTo) PtrOfPtrToPtrStrings(pointerStringOfArray *[]*string) *[]string {
	if pointerStringOfArray == nil || *pointerStringOfArray == nil {
		var emptyResult []string

		return &emptyResult
	}

	newArray := make([]string, len(*pointerStringOfArray))

	for i, value := range *pointerStringOfArray {
		newArray[i] = *value
	}

	return &newArray
}

func (it stringsTo) PtrOfPtrToMapStringBool(inputArray *[]*string) map[string]bool {
	if inputArray == nil || len(*inputArray) == 0 {
		return map[string]bool{}
	}

	length := len(*inputArray)
	hashset := make(map[string]bool, length)

	for _, s := range *inputArray {
		sC := *s
		hashset[sC] = true
	}

	return hashset
}

func (it stringsTo) CloneIf(
	isClone bool,
	items ...string,
) []string {
	if len(items) == 0 || isClone {
		return items
	}

	newArray := make([]string, len(items))

	for i, value := range items {
		newArray[i] = value
	}

	return newArray
}
