package corecsvtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// DefaultCsv
// ==========================================================================

var ext2DefaultCsvTestCases = []coretestcases.CaseV1{
	{
		Title:         "DefaultCsv single item -- no joiner",
		ArrangeInput:  args.Map{"items": []string{"hello"}},
		ExpectedInput: args.Map{
			"notEmpty": true,
			"hasComma": false,
		},
	},
	{
		Title:         "DefaultCsv multiple items -- comma joined",
		ArrangeInput:  args.Map{"items": []string{"a", "b", "c"}},
		ExpectedInput: args.Map{
			"notEmpty": true,
			"hasComma": true,
		},
	},
	{
		Title:         "DefaultCsv empty items -- returns empty",
		ArrangeInput:  args.Map{"items": []string{}},
		ExpectedInput: args.Map{
			"notEmpty": false,
			"hasComma": false,
		},
	},
}

// ==========================================================================
// DefaultCsvStrings
// ==========================================================================

var ext2DefaultCsvStringsTestCases = []coretestcases.CaseV1{
	{
		Title:         "DefaultCsvStrings multiple items -- quoted strings",
		ArrangeInput:  args.Map{"items": []string{"a", "b"}},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "DefaultCsvStrings empty -- empty slice",
		ArrangeInput:  args.Map{"items": []string{}},
		ExpectedInput: args.Map{"length": 0},
	},
}

// ==========================================================================
// DefaultCsvUsingJoiner
// ==========================================================================

var ext2DefaultCsvUsingJoinerTestCases = []coretestcases.CaseV1{
	{
		Title:         "DefaultCsvUsingJoiner pipe joiner -- joined",
		ArrangeInput:  args.Map{
			"joiner": " | ",
			"items": []string{"x", "y"},
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// DefaultAnyCsv
// ==========================================================================

var ext2DefaultAnyCsvTestCases = []coretestcases.CaseV1{
	{
		Title:         "DefaultAnyCsv single item -- returns string",
		ArrangeInput:  args.Map{"items": []any{42}},
		ExpectedInput: args.Map{"notEmpty": true},
	},
	{
		Title:         "DefaultAnyCsv multiple -- returns csv",
		ArrangeInput:  args.Map{"items": []any{"a", 1, true}},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// DefaultAnyCsvStrings
// ==========================================================================

var ext2DefaultAnyCsvStringsTestCases = []coretestcases.CaseV1{
	{
		Title:         "DefaultAnyCsvStrings multiple items -- returns slice",
		ArrangeInput:  args.Map{"items": []any{"a", "b"}},
		ExpectedInput: args.Map{"length": 2},
	},
}

// ==========================================================================
// DefaultAnyCsvUsingJoiner
// ==========================================================================

var ext2DefaultAnyCsvUsingJoinerTestCases = []coretestcases.CaseV1{
	{
		Title:         "DefaultAnyCsvUsingJoiner pipe joiner -- joined",
		ArrangeInput:  args.Map{
			"joiner": " | ",
			"items": []any{"a", "b"},
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// StringsToCsvStringsDefault
// ==========================================================================

var ext2StringsToCsvStringsDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringsToCsvStringsDefault multiple items -- quoted",
		ArrangeInput:  args.Map{"items": []string{"a", "b"}},
		ExpectedInput: args.Map{"length": 2},
	},
}

// ==========================================================================
// StringsToStringDefault
// ==========================================================================

var ext2StringsToStringDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringsToStringDefault multiple items -- comma joined",
		ArrangeInput:  args.Map{"items": []string{"a", "b"}},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// StringsToCsvString -- all three quote branches
// ==========================================================================

var ext2StringsToCsvStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringsToCsvString single quote -- has single quotes",
		ArrangeInput:  args.Map{
			"joiner": ", ",
			"quote": true,
			"singleQuote": true,
			"items": []string{"a"},
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
	{
		Title:         "StringsToCsvString double quote -- has double quotes",
		ArrangeInput:  args.Map{
			"joiner": ", ",
			"quote": true,
			"singleQuote": false,
			"items": []string{"a"},
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
	{
		Title:         "StringsToCsvString no quote -- plain",
		ArrangeInput:  args.Map{
			"joiner": ", ",
			"quote": false,
			"singleQuote": false,
			"items": []string{"a"},
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
	{
		Title:         "StringsToCsvString empty -- empty string",
		ArrangeInput:  args.Map{
			"joiner": ", ",
			"quote": false,
			"singleQuote": false,
			"items": []string{},
		},
		ExpectedInput: args.Map{"notEmpty": false},
	},
}

// ==========================================================================
// StringsToCsvStrings -- all three quote branches
// ==========================================================================

var ext2StringsToCsvStringsTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringsToCsvStrings single quote",
		ArrangeInput:  args.Map{
			"quote": true,
			"singleQuote": true,
			"items": []string{"a", "b"},
		},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "StringsToCsvStrings double quote",
		ArrangeInput:  args.Map{
			"quote": true,
			"singleQuote": false,
			"items": []string{"a", "b"},
		},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "StringsToCsvStrings no quote",
		ArrangeInput:  args.Map{
			"quote": false,
			"singleQuote": false,
			"items": []string{"a", "b"},
		},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "StringsToCsvStrings empty -- empty slice",
		ArrangeInput:  args.Map{
			"quote": false,
			"singleQuote": false,
			"items": []string{},
		},
		ExpectedInput: args.Map{"length": 0},
	},
}

// ==========================================================================
// CompileStringersToString
// ==========================================================================

var ext2CompileStringersToStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "CompileStringersToString with funcs -- joined string",
		ArrangeInput:  args.Map{
			"count": 2,
			"joiner": ", ",
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// CompileStringersToStringDefault
// ==========================================================================

var ext2CompileStringersToStringDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "CompileStringersToStringDefault with funcs -- csv string",
		ArrangeInput:  args.Map{"count": 2},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// StringFunctionsToString -- all three branches
// ==========================================================================

var ext2StringFunctionsToStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringFunctionsToString single quote",
		ArrangeInput:  args.Map{
			"quote": true,
			"singleQuote": true,
			"count": 2,
		},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "StringFunctionsToString double quote",
		ArrangeInput:  args.Map{
			"quote": true,
			"singleQuote": false,
			"count": 2,
		},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "StringFunctionsToString no quote",
		ArrangeInput:  args.Map{
			"quote": false,
			"singleQuote": false,
			"count": 2,
		},
		ExpectedInput: args.Map{"length": 2},
	},
	{
		Title:         "StringFunctionsToString empty -- empty slice",
		ArrangeInput:  args.Map{
			"quote": false,
			"singleQuote": false,
			"count": 0,
		},
		ExpectedInput: args.Map{"length": 0},
	},
}
