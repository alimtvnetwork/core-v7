package corerangestests

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/issetter"
)

var (
	validIntRangeTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "2-5, 7-10, 15-20 --- ranges generate for int",
				ArrangeInput: []corerange.MinMaxInt{
					{
						Min: 2,
						Max: 5,
					},
					{
						Min: 7,
						Max: 10,
					},
					{
						Min: 15,
						Max: 20,
					},
				},
				ExpectedInput: []int{
					2, 3, 4,
					5, 7, 8,
					9, 10, 15,
					16, 17, 18,
					19, 20,
				},
				VerifyTypeOf: &coretests.VerifyTypeOf{
					ArrangeInput:  reflect.TypeOf([]corerange.MinMaxInt{}),
					ActualInput:   reflect.TypeOf([]int{}),
					ExpectedInput: reflect.TypeOf([]int{}),
				},
				IsEnable: issetter.True,
			},
		},
	}

	validInt8RangeTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "2-5, 7-10, 15-20 --- ranges generate for int8",
				ArrangeInput: []corerange.MinMaxInt8{
					{
						Min: 2,
						Max: 5,
					},
					{
						Min: 7,
						Max: 10,
					},
					{
						Min: 15,
						Max: 20,
					},
				},
				ExpectedInput: []int8{
					2, 3, 4,
					5, 7, 8,
					9, 10, 15,
					16, 17, 18,
					19, 20,
				},
				VerifyTypeOf: &coretests.VerifyTypeOf{
					ArrangeInput:  reflect.TypeOf([]corerange.MinMaxInt8{}),
					ActualInput:   reflect.TypeOf([]int8{}),
					ExpectedInput: reflect.TypeOf([]int8{}),
				},
				IsEnable: issetter.True,
			},
		},
	}

	validStartEndRangesTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "2-5, 7-10, 15-20 --- ranges generate for int8",
				ArrangeInput: []corerange.StartEndInt{
					{
						Start: 2,
						End:   5,
					},
					{
						Start: 7,
						End:   10,
					},
					{
						Start: 15,
						End:   20,
					},
				},
				ExpectedInput: []int{
					2, 3, 4,
					5, 7, 8,
					9, 10, 15,
					16, 17, 18,
					19, 20,
				},
				VerifyTypeOf: &coretests.VerifyTypeOf{
					ArrangeInput:  reflect.TypeOf([]corerange.StartEndInt{}),
					ActualInput:   reflect.TypeOf([]int{}),
					ExpectedInput: reflect.TypeOf([]int{}),
				},
				IsEnable: issetter.True,
			},
		},
	}

	startEndRangesStringFunctionsVerificationTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Verifying String, StringColon, StringHyphen, StringSpace functions values",
				ArrangeInput: []corerange.StartEndInt{
					{
						Start: 2,
						End:   5,
					},
					{
						Start: 7,
						End:   10,
					},
					{
						Start: 15,
						End:   20,
					},
				},
				VerifyTypeOf: &coretests.VerifyTypeOf{
					ArrangeInput:  reflect.TypeOf([]corerange.StartEndInt{}),
					ActualInput:   reflect.TypeOf([]string{}),
					ExpectedInput: reflect.TypeOf([]string{}),
				},
				IsEnable:        issetter.True,
				HasError:        false,
				IsValidateError: false,
			},
			Validator: corevalidator.SliceValidator{
				Condition: corevalidator.DefaultTrimCoreCondition,
				CompareAs: stringcompareas.Equal,
				ExpectedLines: []string{
					"StartEnd : 2-5",
					"    [0] func : String        | result : 2-5",
					"    [0] func : StringColon   | result : 2:5",
					"    [0] func : StringHyphen  | result : 2-5",
					"    [0] func : StringSpace   | result : 2 5",
					"StartEnd : 7-10",
					"    [1] func : String        | result : 7-10",
					"    [1] func : StringColon   | result : 7:10",
					"    [1] func : StringHyphen  | result : 7-10",
					"    [1] func : StringSpace   | result : 7 10",
					"StartEnd : 15-20",
					"    [2] func : String        | result : 15-20",
					"    [2] func : StringColon   | result : 15:20",
					"    [2] func : StringHyphen  | result : 15-20",
					"    [2] func : StringSpace   | result : 15 20",
				},
			},
		},
	}

	someRange     = corerange.NewRangeIntUsingValues(5, 25, true)
	minMaxDefault = corerange.MinMaxInt{
		Min: 5,
		Max: 25,
	}
	rangeInt16        = corerange.NewRangeInt16("5:25", ":", nil)
	range16WithinFunc = func(x int) bool {
		return rangeInt16.IsWithinRange(int16(x))
	}

	isWithInFuncsMap = map[string]isWithInDefinitionFunc{
		reflect.TypeOf(someRange).String():     someRange.IsValidPlusWithinRange,
		reflect.TypeOf(minMaxDefault).String(): minMaxDefault.IsWithinRange,
		reflect.TypeOf(rangeInt16).String():    range16WithinFunc,
	}
)
