package coreappendtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var appendAnyItemsTestCases = []coretestcases.CaseV1{
	{
		Title:         "AppendAnyItemsToString comma -- not empty",
		ArrangeInput:  args.Map{
			"when": "comma joiner",
			"joiner": ", ",
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var prependAnyItemsTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAnyItemsToString comma -- not empty",
		ArrangeInput:  args.Map{
			"when": "comma joiner",
			"joiner": ", ",
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var prependAppendStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAppendAnyItemsToString -- not empty",
		ArrangeInput:  args.Map{
			"when": "comma joiner",
			"joiner": ", ",
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var prependAppendSkipNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAppendSkipOnNil -- skips nil items (3 total: prefix + item1 + item3 + suffix)",
		ArrangeInput:  args.Map{"when": "nil in middle"},
		ExpectedInput: args.Map{"length": 4},
	},
}

var prependAppendNilPrependTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAppend nil prepend -- skips prepend (item1 + suffix)",
		ArrangeInput:  args.Map{"when": "nil prepend"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var prependAppendNilAppendTestCases = []coretestcases.CaseV1{
	{
		Title:         "PrependAppend nil append -- skips append (prefix + item1)",
		ArrangeInput:  args.Map{"when": "nil append"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var mapAppendTestCases = []coretestcases.CaseV1{
	{
		Title:         "MapStringStringAppend -- appends 2 items to existing 1",
		ArrangeInput:  args.Map{"when": "append 2 items"},
		ExpectedInput: args.Map{"length": 3},
	},
}

var mapAppendSkipEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "MapStringStringAppend skip empty -- skips empty value",
		ArrangeInput:  args.Map{"when": "skip empty"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var mapAppendEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "MapStringStringAppend empty append -- unchanged",
		ArrangeInput:  args.Map{"when": "empty append map"},
		ExpectedInput: args.Map{"length": 1},
	},
}
