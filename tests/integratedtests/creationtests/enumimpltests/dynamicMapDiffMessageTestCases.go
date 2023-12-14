package enumimpltests

import (
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
			ExpectedInput: []string{
				"Dynamic map diff string compiled must be same",
				"",
				"Difference Between Map:",
				"",
				"{{",
				"",
				"  \"not-exist-in-left\":2,",
				"  \"not-exist-in-right\":3,",
				"  \"exist-in-left-right-diff-val\":5",
				"",
				"}}",
			},
			VerifyTypeOf:    coretests.NewVerifyTypeOf(LeftRightDynamicMap{}),
			HasError:        false,
			IsValidateError: true,
		},
	},
}
