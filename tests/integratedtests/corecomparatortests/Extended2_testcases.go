package corecomparatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var compareLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.Is returns true -- Equal vs Equal",
		ArrangeInput:  args.Map{
			"when": "both equal",
			"left": 0,
			"right": 0,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "Compare.Is returns false -- LeftGreater vs Equal",
		ArrangeInput:  args.Map{
			"when": "greater vs equal",
			"left": 1,
			"right": 0,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "Compare.Is returns true -- LeftLess vs LeftLess",
		ArrangeInput:  args.Map{
			"when": "both left less",
			"left": 3,
			"right": 3,
		},
		ExpectedInput: args.Map{"result": true},
	},
}

var compareIsAnyOfTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.IsAnyOf returns true -- Equal in Equal,LeftLess",
		ArrangeInput:  args.Map{
			"when": "Equal",
			"value": 0,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "Compare.IsAnyOf returns false -- LeftGreater not in Equal,LeftLess",
		ArrangeInput:  args.Map{
			"when": "LeftGreater",
			"value": 1,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var compareIsAnyOfEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.IsAnyOf returns true -- empty values list",
		ArrangeInput:  args.Map{"when": "empty values"},
		ExpectedInput: args.Map{"result": true},
	},
}

var compareNameValueTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.NameValue returns not empty -- Equal value",
		ArrangeInput:  args.Map{
			"when": "Equal",
			"value": 0,
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var compareCsvStringsTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.CsvStrings returns length 2 -- 2 values",
		ArrangeInput:  args.Map{"when": "2 values"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var compareCsvStringsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.CsvStrings returns length 0 -- no values",
		ArrangeInput:  args.Map{"when": "no values"},
		ExpectedInput: args.Map{"length": 0},
	},
}

var compareValueConversionsTestCases = []coretestcases.CaseV1{
	{
		Title:        "Compare.ValueConversions returns correct values -- Equal (0)",
		ArrangeInput: args.Map{
			"when": "Equal (0)",
			"value": 0,
		},
		ExpectedInput: args.Map{
			"valueByte":        0,
			"valueInt":         0,
			"toNumberString":   "0",
			"numberString":     "0",
			"numberJsonString": "\"0\"",
		},
	},
	{
		Title:        "Compare.ValueConversions returns correct values -- LeftGreater (1)",
		ArrangeInput: args.Map{
			"when": "LeftGreater (1)",
			"value": 1,
		},
		ExpectedInput: args.Map{
			"valueByte":        1,
			"valueInt":         1,
			"toNumberString":   "1",
			"numberString":     "1",
			"numberJsonString": "\"1\"",
		},
	},
}

var compareMarshalJsonTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.MarshalJSON returns no error and not empty -- Equal",
		ArrangeInput:  args.Map{"when": "Equal"},
		ExpectedInput: args.Map{
			"hasError": false,
			"notEmpty": true,
		},
	},
}

var compareOnlySupportedErrTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.OnlySupportedErr returns no error -- Equal in list",
		ArrangeInput:  args.Map{
			"when": "Equal supported",
			"value": 0,
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title:         "Compare.OnlySupportedErr returns error -- LeftGreater not in list",
		ArrangeInput:  args.Map{
			"when": "LeftGreater unsupported",
			"value": 1,
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}

var compareOnlySupportedDirectErrTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.OnlySupportedDirectErr returns no error -- Equal in list",
		ArrangeInput:  args.Map{
			"when": "Equal supported",
			"value": 0,
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title:         "Compare.OnlySupportedDirectErr returns error -- LeftGreater not in list",
		ArrangeInput:  args.Map{
			"when": "LeftGreater unsupported",
			"value": 1,
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}

var compareOnlySupportedEmptyMsgTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.OnlySupportedErr returns error -- empty message unsupported",
		ArrangeInput:  args.Map{"when": "empty message unsupported"},
		ExpectedInput: args.Map{"hasError": true},
	},
}

var minLengthTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.MinLength returns left -- left 3 smaller than right 5",
		ArrangeInput:  args.Map{
			"when": "3 vs 5",
			"left": 3,
			"right": 5,
		},
		ExpectedInput: args.Map{"result": 3},
	},
	{
		Title:         "Compare.MinLength returns right -- right 4 smaller than left 7",
		ArrangeInput:  args.Map{
			"when": "7 vs 4",
			"left": 7,
			"right": 4,
		},
		ExpectedInput: args.Map{"result": 4},
	},
	{
		Title:         "Compare.MinLength returns either -- both equal 5",
		ArrangeInput:  args.Map{
			"when": "5 vs 5",
			"left": 5,
			"right": 5,
		},
		ExpectedInput: args.Map{"result": 5},
	},
}

var compareIsAnyNamesOfTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.IsAnyNamesOf returns true -- Equal name in list",
		ArrangeInput:  args.Map{"when": "Equal name in list"},
		ExpectedInput: args.Map{"result": true},
	},
}

var compareIsInconclusiveOrNotEqualTestCases = []coretestcases.CaseV1{
	{
		Title:         "Compare.IsInconclusiveOrNotEqual returns true -- Inconclusive value 6",
		ArrangeInput:  args.Map{
			"when": "Inconclusive",
			"value": 6,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "Compare.IsInconclusiveOrNotEqual returns true -- NotEqual value 5",
		ArrangeInput:  args.Map{
			"when": "NotEqual",
			"value": 5,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "Compare.IsInconclusiveOrNotEqual returns false -- Equal value 0",
		ArrangeInput:  args.Map{
			"when": "Equal",
			"value": 0,
		},
		ExpectedInput: args.Map{"result": false},
	},
}
