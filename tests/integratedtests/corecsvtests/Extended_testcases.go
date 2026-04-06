package corecsvtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var anyItemsToStringDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "Single item returns its string",
		ArrangeInput:  args.Map{"items": []any{"hello"}},
		ExpectedInput: args.Map{"notEmpty": true},
	},
	{
		Title:         "Multiple items returns csv",
		ArrangeInput:  args.Map{"items": []any{"a", "b", "c"}},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var compileStringersToCsvStringsTestCases = []coretestcases.CaseV1{
	{
		Title:         "Empty funcs returns empty",
		ArrangeInput:  args.Map{"count": 0},
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title:         "Single quote mode",
		ArrangeInput:  args.Map{
			"count": 1,
			"includeQuote": true,
			"singleQuote": true,
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title:         "Double quote mode",
		ArrangeInput:  args.Map{
			"count": 1,
			"includeQuote": true,
			"singleQuote": false,
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title:         "No quote mode",
		ArrangeInput:  args.Map{
			"count": 1,
			"includeQuote": false,
			"singleQuote": false,
		},
		ExpectedInput: args.Map{"length": 1},
	},
}
