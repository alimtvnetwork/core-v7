package enumimpltests

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/coretests"
)

var dynamicMapDiffTestCases = []EnumImplDynamicMapTestWrapper{
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "dynamic map must yield diff properly.",
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
			ExpectedInput: enumimpl.DynamicMap{
				"exist-in-left-right-diff-val": 5,
				"not-exist-in-left":            2,
				"not-exist-in-right":           3,
			},
			VerifyTypeOf:    typeVerifyOfForDynamicMapDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
}
