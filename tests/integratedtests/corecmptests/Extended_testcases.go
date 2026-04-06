package corecmptests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var byteCompareTestCases = []coretestcases.CaseV1{
	{
		Title:         "ByteCompare returns Equal -- equal bytes",
		ArrangeInput:  args.Map{
			"when": "equal bytes",
			"left": 5,
			"right": 5,
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title:         "ByteCompare returns LeftLess -- left less",
		ArrangeInput:  args.Map{
			"when": "left less",
			"left": 3,
			"right": 7,
		},
		ExpectedInput: args.Map{"name": "LeftLess"},
	},
	{
		Title:         "ByteCompare returns LeftGreater -- left greater",
		ArrangeInput:  args.Map{
			"when": "left greater",
			"left": 9,
			"right": 2,
		},
		ExpectedInput: args.Map{"name": "LeftGreater"},
	},
}

var isStringsEqualWithoutOrderTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsStringsEqualWithoutOrder returns true -- same strings different order",
		ArrangeInput:  args.Map{
			"when": "same unordered",
			"left": []string{"b", "a", "c"},
			"right": []string{"c", "a", "b"},
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsStringsEqualWithoutOrder returns false -- different strings",
		ArrangeInput:  args.Map{
			"when": "different",
			"left": []string{"a", "b"},
			"right": []string{"c", "d"},
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsStringsEqualWithoutOrder returns false -- different length",
		ArrangeInput:  args.Map{
			"when": "diff length",
			"left": []string{"a"},
			"right": []string{"a", "b"},
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var versionSliceByteTestCases = []coretestcases.CaseV1{
	{
		Title:         "VersionSliceByte returns Equal -- equal versions",
		ArrangeInput:  args.Map{
			"when": "equal versions",
			"left": []int{1, 2, 3},
			"right": []int{1, 2, 3},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title:         "VersionSliceByte returns LeftLess -- left less",
		ArrangeInput:  args.Map{
			"when": "left less",
			"left": []int{1, 2, 0},
			"right": []int{1, 2, 3},
		},
		ExpectedInput: args.Map{"name": "LeftLess"},
	},
	{
		Title:         "VersionSliceByte returns LeftGreater -- left greater",
		ArrangeInput:  args.Map{
			"when": "left greater",
			"left": []int{1, 3, 0},
			"right": []int{1, 2, 0},
		},
		ExpectedInput: args.Map{"name": "LeftGreater"},
	},
	{
		Title:         "VersionSliceByte returns Equal -- both nil",
		ArrangeInput:  args.Map{
			"when": "both nil",
			"left": nil,
			"right": nil,
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title:         "VersionSliceByte returns LeftLess -- left shorter",
		ArrangeInput:  args.Map{
			"when": "left shorter",
			"left": []int{1, 2},
			"right": []int{1, 2, 3},
		},
		ExpectedInput: args.Map{"name": "LeftLess"},
	},
}

var versionSliceIntegerTestCases = []coretestcases.CaseV1{
	{
		Title:         "VersionSliceInteger returns Equal -- equal versions",
		ArrangeInput:  args.Map{
			"when": "equal",
			"left": []int{1, 0, 5},
			"right": []int{1, 0, 5},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title:         "VersionSliceInteger returns LeftLess -- left less",
		ArrangeInput:  args.Map{
			"when": "left less",
			"left": []int{1, 0, 3},
			"right": []int{1, 0, 5},
		},
		ExpectedInput: args.Map{"name": "LeftLess"},
	},
	{
		Title:         "VersionSliceInteger returns LeftGreater -- left greater",
		ArrangeInput:  args.Map{
			"when": "left greater",
			"left": []int{2, 0, 0},
			"right": []int{1, 9, 9},
		},
		ExpectedInput: args.Map{"name": "LeftGreater"},
	},
	{
		Title:         "VersionSliceInteger returns Equal -- both nil",
		ArrangeInput:  args.Map{
			"when": "both nil",
			"left": nil,
			"right": nil,
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title:         "VersionSliceInteger returns NotEqual -- one nil",
		ArrangeInput:  args.Map{
			"when": "one nil",
			"left": []int{1},
			"right": nil,
		},
		ExpectedInput: args.Map{"name": "NotEqual"},
	},
}

var isStringsEqualBothNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsStringsEqual returns true -- both nil",
		ArrangeInput:  args.Map{"when": "both nil"},
		ExpectedInput: args.Map{"result": true},
	},
}

var isStringsEqualOneNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsStringsEqual returns false -- one nil",
		ArrangeInput:  args.Map{"when": "one nil"},
		ExpectedInput: args.Map{"result": false},
	},
}

var compareExtendedMethodsTestCases = []coretestcases.CaseV1{
	{
		Title:        "Compare returns correct booleans -- Equal value",
		ArrangeInput: args.Map{
			"when": "Equal value",
			"value": 0,
		},
		ExpectedInput: args.Map{
			"isLess":                false,
			"isLessEqual":          true,
			"isGreater":            false,
			"isGreaterEqual":       true,
			"isDefined":            true,
			"isInconclusive":       false,
			"isNotEqual":           false,
			"isNotEqualLogically":  false,
			"isDefinedProperly":    true,
		},
	},
	{
		Title:        "Compare returns correct booleans -- LeftGreater value",
		ArrangeInput: args.Map{
			"when": "LeftGreater value",
			"value": 1,
		},
		ExpectedInput: args.Map{
			"isLess":                false,
			"isLessEqual":          false,
			"isGreater":            true,
			"isGreaterEqual":       true,
			"isDefined":            true,
			"isInconclusive":       false,
			"isNotEqual":           false,
			"isNotEqualLogically":  true,
			"isDefinedProperly":    true,
		},
	},
	{
		Title:        "Compare returns correct booleans -- LeftLess value",
		ArrangeInput: args.Map{
			"when": "LeftLess value",
			"value": 3,
		},
		ExpectedInput: args.Map{
			"isLess":                true,
			"isLessEqual":          true,
			"isGreater":            false,
			"isGreaterEqual":       false,
			"isDefined":            true,
			"isInconclusive":       false,
			"isNotEqual":           false,
			"isNotEqualLogically":  true,
			"isDefinedProperly":    true,
		},
	},
	{
		Title:        "Compare returns correct booleans -- Inconclusive value",
		ArrangeInput: args.Map{
			"when": "Inconclusive value",
			"value": 6,
		},
		ExpectedInput: args.Map{
			"isLess":                false,
			"isLessEqual":          false,
			"isGreater":            false,
			"isGreaterEqual":       false,
			"isDefined":            false,
			"isInconclusive":       true,
			"isNotEqual":           false,
			"isNotEqualLogically":  true,
			"isDefinedProperly":    false,
		},
	},
}
