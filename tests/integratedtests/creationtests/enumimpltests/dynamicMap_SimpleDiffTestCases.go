package enumimpltests

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/coretests"
)

var dynamicMapSimpleDiffTestCases = []EnumImplDynamicMapTestWrapper{
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map simple diff [someKey2] mismatch verify",
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
			ExpectedInput: []string{
				"Dynamic map simple diff [someKey2] mismatch verify",
				"",
				"Difference Between Map:",
				"",
				"{",
				"- Left Map - Has Diff from Right Map:",
				"",
				"  {",
				"  ",
				"    \"someKey2\":4",
				"  ",
				"  }",
				"",
				"- Right Map - Has Diff from Left Map:",
				"",
				"  {",
				"  ",
				"    \"someKey2\":2",
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
			Title: "Dynamic map simple diff [someKey2], [someKey4] mismatch verify",
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
			ExpectedInput: []string{
				"Dynamic map simple diff [someKey2], [someKey4] mismatch verify",
				"",
				"Difference Between Map:",
				"",
				"{",
				"- Left Map - Has Diff from Right Map:",
				"",
				"  {",
				"  ",
				"    \"someKey4\":4",
				"  ",
				"  }",
				"",
				"- Right Map - Has Diff from Left Map:",
				"",
				"  {",
				"  ",
				"    \"someKey2\":2",
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
			Title: "Dynamic map simple diff [someKey2], [someKey4] mismatch verify",
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
			ActualInput: nil,
			ExpectedInput: []string{
				"",
			},
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
}
