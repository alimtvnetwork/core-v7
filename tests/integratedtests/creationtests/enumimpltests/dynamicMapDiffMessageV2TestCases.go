package enumimpltests

import (
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
			ExpectedInput: []string{
				"Dynamic map diff string compiled must be same",
				"",
				"Difference Between Map:",
				"",
				"{",
				"- Left Map - Has Diff from Right Map:",
				"",
				"  {",
				"  ",
				"    \"exist-in-left-right-diff-val\":\"{\"Left\":5,\"Right\":6}\",",
				"    \"not-exist-in-left\":\"{\"Left\":null,\"Right\":2}\"",
				"  ",
				"  }",
				"",
				"- Right Map - Has Diff from Left Map:",
				"",
				"  {",
				"  ",
				"    \"exist-in-left-right-diff-val\":\"{\"Left\":5,\"Right\":6}\",",
				"    \"not-exist-in-right\":\"3 (type:int) - left - key is missing!\"",
				"  ",
				"  }}",
			},
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
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
			ExpectedInput: []string{
				"",
			},
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - right hand key missing - cl int 5",
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
			ExpectedInput: []string{
				"Dynamic map diff - right hand key missing - cl int 5",
				"",
				"Difference Between Map:",
				"",
				"{",
				"- Right Map - Has Diff from Left Map:",
				"",
				"  {",
				"  ",
				"    \"cl\":\"5 (type:int) - left - key is missing!\"",
				"  ",
				"  }}",
			},
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - left hand key missing - cl {left, right}",
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
			ExpectedInput: []string{
				"Dynamic map diff - left hand key missing - cl {left, right}",
				"",
				"Difference Between Map:",
				"",
				"{",
				"- Left Map - Has Diff from Right Map:",
				"",
				"  {",
				"  ",
				"    \"cl\":\"{\"Left\":null,\"Right\":5}\"",
				"  ",
				"  }}",
			},
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - left cl - key missing",
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
			ExpectedInput: []string{
				"Dynamic map diff - left cl - key missing",
				"",
				"Difference Between Map:",
				"",
				"{",
				"- Left Map - Has Diff from Right Map:",
				"",
				"  {",
				"  ",
				"    \"cl\":\"{\"Left\":null,\"Right\":5}\"",
				"  ",
				"  }}",
			},
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
}
