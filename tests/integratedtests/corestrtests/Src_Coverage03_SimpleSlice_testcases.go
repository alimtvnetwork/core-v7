package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// Branch: SimpleSlice basic state checks (empty, length, lastIndex)
var srcC03BasicStateTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice Empty returns empty state -- new empty slice",
		ExpectedInput: args.Map{
			"isEmpty":   true,
			"hasAny":    false,
			"length":    0,
			"count":     0,
			"lastIndex": -1,
		},
	},
}

// Branch: SimpleSlice nil receiver
var srcC03NilReceiverTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice nil receiver returns safe defaults -- nil pointer",
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

// Branch: Add methods (Add, AddSplit, AddIf, Adds, Append, AppendFmt, AppendFmtIf)
var srcC03AddMethodsTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice Add returns correct length -- two items added",
		ArrangeInput: args.Map{
			"method": "Add",
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "SimpleSlice AddSplit returns correct length -- comma separated",
		ArrangeInput: args.Map{
			"method": "AddSplit",
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
	{
		Title: "SimpleSlice AddIf returns correct length -- one true one false",
		ArrangeInput: args.Map{
			"method": "AddIf",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "SimpleSlice Adds returns correct length -- empty then two items",
		ArrangeInput: args.Map{
			"method": "Adds",
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "SimpleSlice Append returns correct length -- empty then one item",
		ArrangeInput: args.Map{
			"method": "Append",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "SimpleSlice AppendFmt returns correct length -- empty fmt skipped",
		ArrangeInput: args.Map{
			"method": "AppendFmt",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "SimpleSlice AppendFmtIf returns correct length -- false skipped true added",
		ArrangeInput: args.Map{
			"method": "AppendFmtIf",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
}

// Branch: Title/Value and Curly wrap add methods
var srcC03TitleValueTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice AddAsTitleValue returns 1 -- key val pair",
		ArrangeInput: args.Map{
			"method": "AddAsTitleValue",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "SimpleSlice AddAsCurlyTitleWrap returns 1 -- key val pair",
		ArrangeInput: args.Map{
			"method": "AddAsCurlyTitleWrap",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "SimpleSlice AddAsCurlyTitleWrapIf returns 1 -- false skipped true added",
		ArrangeInput: args.Map{
			"method": "AddAsCurlyTitleWrapIf",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "SimpleSlice AddAsTitleValueIf returns 1 -- false skipped true added",
		ArrangeInput: args.Map{
			"method": "AddAsTitleValueIf",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
}

// Branch: InsertAt including out-of-range
var srcC03InsertAtTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice InsertAt returns correct item at index -- middle insert",
	ExpectedInput: args.Map{
		"length":  3,
		"atIndex": "b",
	},
}

// Branch: AddsIf, AddError, AddStruct, AddPointer
var srcC03ConditionalAddsTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice AddsIf returns 2 -- false skipped true added",
		ArrangeInput: args.Map{
			"method": "AddsIf",
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "SimpleSlice AddError returns 1 -- nil skipped error added",
		ArrangeInput: args.Map{
			"method": "AddError",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "SimpleSlice AddStruct returns 1 -- nil skipped struct added",
		ArrangeInput: args.Map{
			"method": "AddStruct",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "SimpleSlice AddPointer returns 1 -- nil skipped pointer added",
		ArrangeInput: args.Map{
			"method": "AddPointer",
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
}

// Branch: AsError / AsDefaultError
var srcC03AsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice AsDefaultError returns nil -- empty slice",
		ArrangeInput: args.Map{
			"hasItems": false,
		},
		ExpectedInput: args.Map{
			"defaultErrorNil": true,
			"asErrorNil":      true,
		},
	},
	{
		Title: "SimpleSlice AsDefaultError returns non-nil -- has items",
		ArrangeInput: args.Map{
			"hasItems": true,
		},
		ExpectedInput: args.Map{
			"defaultErrorNil": false,
		},
	},
}

// Branch: First/Last methods
var srcC03FirstLastTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice First Last returns correct values -- three items",
	ExpectedInput: args.Map{
		"first":                 "a",
		"last":                  "c",
		"firstDynamic":          "a",
		"lastDynamic":           "c",
		"firstOrDefault":        "a",
		"lastOrDefault":         "c",
		"firstOrDefaultDynamic": "a",
		"lastOrDefaultDynamic":  "c",
	},
}

// Branch: FirstOrDefault empty
var srcC03FirstLastEmptyTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice FirstOrDefault returns empty -- empty slice",
	ExpectedInput: args.Map{
		"firstOrDefault": "",
		"lastOrDefault":  "",
	},
}

// Branch: Skip/Take
var srcC03SkipTakeTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice Skip Take returns correct lengths -- three items",
	ExpectedInput: args.Map{
		"skip1":   2,
		"skip100": 0,
		"take2":   2,
		"take100": 3,
	},
}

// Branch: CountFunc
var srcC03CountFuncTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice CountFunc returns 2 -- items longer than 1",
		ArrangeInput: args.Map{
			"items": "a,bb,ccc",
		},
		ExpectedInput: args.Map{
			"count": 2,
		},
	},
	{
		Title: "SimpleSlice CountFunc returns 0 -- empty slice",
		ArrangeInput: args.Map{
			"items": "",
		},
		ExpectedInput: args.Map{
			"count": 0,
		},
	},
}

// Branch: IsContains / IsContainsFunc
var srcC03ContainsTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice IsContains returns true -- item exists",
		ArrangeInput: args.Map{
			"search": "a",
			"items":  "a,b",
		},
		ExpectedInput: args.Map{
			"found": true,
		},
	},
	{
		Title: "SimpleSlice IsContains returns false -- item missing",
		ArrangeInput: args.Map{
			"search": "c",
			"items":  "a,b",
		},
		ExpectedInput: args.Map{
			"found": false,
		},
	},
	{
		Title: "SimpleSlice IsContains returns false -- empty slice",
		ArrangeInput: args.Map{
			"search": "a",
			"items":  "",
		},
		ExpectedInput: args.Map{
			"found": false,
		},
	},
}

// Branch: IndexOf / IndexOfFunc
var srcC03IndexOfTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice IndexOf returns 1 -- item at index 1",
		ArrangeInput: args.Map{
			"search": "b",
			"items":  "a,b,c",
		},
		ExpectedInput: args.Map{
			"index": 1,
		},
	},
	{
		Title: "SimpleSlice IndexOf returns -1 -- item not found",
		ArrangeInput: args.Map{
			"search": "z",
			"items":  "a,b,c",
		},
		ExpectedInput: args.Map{
			"index": -1,
		},
	},
	{
		Title: "SimpleSlice IndexOf returns -1 -- empty slice",
		ArrangeInput: args.Map{
			"search": "a",
			"items":  "",
		},
		ExpectedInput: args.Map{
			"index": -1,
		},
	},
}

// Branch: HasIndex
var srcC03HasIndexTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice HasIndex returns correct flags -- two items",
	ExpectedInput: args.Map{
		"has0":    true,
		"has1":    true,
		"has2":    false,
		"hasNeg1": false,
	},
}

// Branch: Strings/List
var srcC03StringsListTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice Strings List returns correct length -- one item",
	ExpectedInput: args.Map{
		"stringsLen": 1,
		"listLen":    1,
	},
}

// Branch: WrapQuotes
var srcC03WrapQuotesTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice WrapQuotes methods execute without panic -- one item",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}
