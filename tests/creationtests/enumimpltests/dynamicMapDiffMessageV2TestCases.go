package enumimpltests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
)

var dynamicMapDiffMessageTestCasesV2 = []EnumImplDynamicMapTestWrapper{
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff string compiled must be same",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"exist":                        1,
						"not-exist-in-right":           3,
						"exist-in-left-right-diff-val": 5,
					},
					Right: map[string]interface{}{
						"exist":                        1,
						"not-exist-in-left":            2,
						"exist-in-left-right-diff-val": 6,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput: nil,
			ExpectedInput: "Dynamic map diff string compiled must be same\n\n" +
				"Difference Between Map:\n\n" +
				"{\n" +
				"- Left Map - Has Diff from Right Map:\n" +
				"{\n\n\"" +
				"exist-in-left-right-diff-val\":\"{\"Left\":5,\"Right\":6}\",\n\"" +
				"not-exist-in-left\":\"{\"Left\":null,\"Right\":2}\"\n\n}" +
				"\n\n\n" +
				"- Right Map - Has Diff from Left Map:\n{\n\n\"" +
				"exist-in-left-right-diff-val\":\"{\"Left\":5,\"Right\":6}\",\n\"" +
				"not-exist-in-right\":\"3 (type:int) - left - key is missing!\"\n\n}}",
			ArrangeExpectedType:    reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
			ActualExpectedType:     reflect.TypeOf(""),
			ExpectedTypeOfExpected: reflect.TypeOf(""),
			HasError:               false,
			IsValidateError:        true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - no changes",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
					Right: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput:            nil,
			ExpectedInput:          "",
			ArrangeExpectedType:    reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
			ActualExpectedType:     reflect.TypeOf(""),
			ExpectedTypeOfExpected: reflect.TypeOf(""),
			HasError:               false,
			IsValidateError:        true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - right hand key missing",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
					Right: map[string]interface{}{
						"a": 1,
						"b": 3,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput: nil,
			ExpectedInput: "Dynamic map diff - right hand key missing\n\n" +
				"Difference Between Map:\n\n{\n" +
				"- Right Map - Has Diff from Left Map:\n" +
				"{\n\n\"" +
				"cl\":\"5 (type:int) - left - key is missing!\"\n\n}}",
			ArrangeExpectedType:    reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
			ActualExpectedType:     reflect.TypeOf(""),
			ExpectedTypeOfExpected: reflect.TypeOf(""),
			HasError:               false,
			IsValidateError:        true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - left hand key missing",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"a": 1,
						"b": 3,
					},
					Right: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput: nil,
			ExpectedInput: "Dynamic map diff - left hand key missing\n\n" +
				"Difference Between Map:\n\n{\n" +
				"- Left Map - Has Diff from Right Map:" +
				"\n{\n\n\"" +
				"cl\":\"{\"Left\":null,\"Right\":5}\"\n\n}\n}",
			ArrangeExpectedType:    reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
			ActualExpectedType:     reflect.TypeOf(""),
			ExpectedTypeOfExpected: reflect.TypeOf(""),
			HasError:               false,
			IsValidateError:        true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - left hand key missing",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"a": 1,
						"b": 3,
					},
					Right: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput: nil,
			ExpectedInput: "Dynamic map diff - left hand key missing\n\n" +
				"Difference Between Map:\n\n{\n" +
				"- Left Map - Has Diff from Right Map:" +
				"\n{\n\n\"" +
				"cl\":\"{\"Left\":null,\"Right\":5}\"\n\n}\n}",
			ArrangeExpectedType:    reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
			ActualExpectedType:     reflect.TypeOf(""),
			ExpectedTypeOfExpected: reflect.TypeOf(""),
			HasError:               false,
			IsValidateError:        true,
		},
	},
}
