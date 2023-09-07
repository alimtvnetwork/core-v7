package enumimpltests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
)

var dynamicMapDiffMessageTestCases = []EnumImplDynamicMapTestWrapper{
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff string compiled must be same",
			ArrangeInput: LeftRightDynamicMap{
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
			ActualInput: nil,
			ExpectedInput: "Dynamic map diff string compiled must be same\n\n" +
				"Difference Between Map:\n\n{{" +
				"\n\n\"" +
				"not-exist-in-left\":2,\n\"" +
				"not-exist-in-right\":3,\n\"" +
				"exist-in-left-right-diff-val\":5\n" +
				"\n}}",
			ArrangeExpectedType:    reflect.TypeOf(LeftRightDynamicMap{}),
			ActualExpectedType:     reflect.TypeOf(""),
			ExpectedTypeOfExpected: reflect.TypeOf(""),
			HasError:               false,
			IsValidateError:        true,
		},
	},
}
