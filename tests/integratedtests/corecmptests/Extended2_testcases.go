package corecmptests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var isStringsEqualPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsStringsEqualPtr returns true -- both nil",
		ArrangeInput:  args.Map{
			"leftNil": true,
			"rightNil": true,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsStringsEqualPtr returns false -- left nil right non-nil",
		ArrangeInput:  args.Map{
			"leftNil": true,
			"rightNil": false,
			"right": []string{"a"},
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsStringsEqualPtr returns false -- right nil left non-nil",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": true,
			"left": []string{"a"},
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "IsStringsEqualPtr returns true -- equal slices",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": false,
			"left": []string{"a", "b"},
			"right": []string{"a", "b"},
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "IsStringsEqualPtr returns false -- different length",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": false,
			"left": []string{"a"},
			"right": []string{"a", "b"},
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var timePtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "TimePtr returns Equal -- both nil",
		ArrangeInput:  args.Map{
			"leftNil": true,
			"rightNil": true,
		},
		ExpectedInput: args.Map{"isEqual": true},
	},
	{
		Title:         "TimePtr returns NotEqual -- left nil",
		ArrangeInput:  args.Map{
			"leftNil": true,
			"rightNil": false,
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "TimePtr returns NotEqual -- right nil",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": true,
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "TimePtr returns Equal -- both same time",
		ArrangeInput:  args.Map{
			"leftNil": false,
			"rightNil": false,
			"sameTime": true,
		},
		ExpectedInput: args.Map{"isEqual": true},
	},
}
