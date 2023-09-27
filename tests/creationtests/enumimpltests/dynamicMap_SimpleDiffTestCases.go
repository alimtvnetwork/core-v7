package enumimpltests

import (
	"reflect"
	
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/coretests"
)

var dynamicMapSimpleDiffTestCases = []EnumImplDynamicMapTestWrapper{
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map simple diff `someKey2` mismatch verify",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"someKey":  1,
						"someKey2": 2,
						"someKey3": 3,
					},
					Right: map[string]interface{}{
						"someKey":  1,
						"someKey2": 4,
						"someKey3": 3,
					},
				},
				DifferChecker: enumimpl.DefaultDiffCheckerImpl,
			},
			ActualInput: nil,
			ExpectedInput: "Dynamic map simple diff `someKey2` mismatch verify\n\n" +
				"Difference Between Map:\n\n{\n" +
				"- Left Map - Has Diff from Right Map:\n" +
				"{\n\n\"someKey2\":4\n\n}\n\n\n" +
				"- Right Map - Has Diff from Left Map:\n" +
				"{\n\n\"someKey2\":2\n\n}}",
			ArrangeExpectedType:    reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
			ActualExpectedType:     reflect.TypeOf(""),
			ExpectedTypeOfExpected: reflect.TypeOf(""),
			HasError:               false,
			IsValidateError:        true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map simple diff `someKey2`, `someKey4` mismatch verify",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"someKey":  1,
						"someKey2": 2,
						"someKey3": 3,
					},
					Right: map[string]interface{}{
						"someKey":  1,
						"someKey4": 4,
						"someKey3": 3,
					},
				},
				DifferChecker: enumimpl.DefaultDiffCheckerImpl,
			},
			ActualInput: nil,
			ExpectedInput: "Dynamic map simple diff `someKey2`, `someKey4` mismatch verify\n\n" +
				"Difference Between Map:\n\n{\n" +
				"- Left Map - Has Diff from Right Map:\n" +
				"{\n\n\"someKey4\":4\n\n}\n\n\n- Right Map - Has Diff from Left Map:\n{\n\n\"someKey2\":2\n\n}}",
			ArrangeExpectedType:    reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
			ActualExpectedType:     reflect.TypeOf(""),
			ExpectedTypeOfExpected: reflect.TypeOf(""),
			HasError:               false,
			IsValidateError:        true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map simple diff `someKey2`, `someKey4` mismatch verify",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"someKey":  1,
						"someKey2": 2,
						"someKey4": 4,
						"someKey3": 3,
					},
					Right: map[string]interface{}{
						"someKey":  1,
						"someKey2": 2,
						"someKey4": 4,
						"someKey3": 3,
					},
				},
				DifferChecker: enumimpl.DefaultDiffCheckerImpl,
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
}
