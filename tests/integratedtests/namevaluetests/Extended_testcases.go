package namevaluetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var instanceStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "Instance String returns non-empty",
		ArrangeInput:  args.Map{
			"name": "key",
			"value": "val",
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var instanceJsonStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "Instance JsonString returns non-empty",
		ArrangeInput:  args.Map{
			"name": "key",
			"value": "val",
		},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var instanceDisposeTestCases = []coretestcases.CaseV1{
	{
		Title:         "Instance Dispose clears values",
		ArrangeInput:  args.Map{
			"name": "key",
			"value": "val",
		},
		ExpectedInput: args.Map{
			"nameEmpty": true,
			"valueEmpty": true,
		},
	},
}

var instanceNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "Nil instance IsNull returns true",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"isNull": true},
	},
}

var extCollectionTestCases = []coretestcases.CaseV1{
	{
		Title:        "Collection operations",
		ArrangeInput: args.Map{},
		ExpectedInput: args.Map{
			"length":          3,
			"hasAnyItem":      true,
			"isEmpty":         false,
			"stringsLen":      3,
			"jsonStringsLen":  3,
			"joinNotEmpty":    true,
			"joinLinesOk":     true,
			"joinCsvOk":       true,
			"joinCsvLineOk":   true,
			"stringOk":        true,
			"jsonStringOk":    true,
			"csvStringsLen":   3,
			"cloneLen":        3,
			"errorNotNil":     true,
			"errorMsgNotNil":  true,
		},
	},
}

var collectionPrependAppendTestCases = []coretestcases.CaseV1{
	{
		Title:        "Prepend and Append operations",
		ArrangeInput: args.Map{},
		ExpectedInput: args.Map{
			"afterPrepend": 2,
			"afterAppend":  3,
			"appendIfTrue": 4,
			"appendIfFalse": 4,
			"prependIfTrue": 5,
			"prependIfFalse": 5,
		},
	},
}

var collectionIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title:         "Same collections are equal",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"isEqual": true},
	},
}

var collectionClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "ClonePtr nil returns nil",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"isNil": true},
	},
	{
		Title:         "ClonePtr non-nil returns clone",
		ArrangeInput:  args.Map{"isNil": false},
		ExpectedInput: args.Map{"isNil": false},
	},
}

var collectionAddsIfTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddsIf true adds items",
		ArrangeInput:  args.Map{"isAdd": true},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title:         "AddsIf false skips items",
		ArrangeInput:  args.Map{"isAdd": false},
		ExpectedInput: args.Map{"length": 0},
	},
}

var collectionConcatTestCases = []coretestcases.CaseV1{
	{
		Title:        "ConcatNew creates new collection",
		ArrangeInput: args.Map{},
		ExpectedInput: args.Map{
			"originalLen": 1,
			"concatLen":   2,
		},
	},
}

var collectionFuncIfTestCases = []coretestcases.CaseV1{
	{
		Title:        "PrependUsingFuncIf and AppendUsingFuncIf",
		ArrangeInput: args.Map{},
		ExpectedInput: args.Map{
			"afterPrependFunc": 2,
			"afterAppendFunc":  3,
		},
	},
}

var collectionAppendPrependIfTestCases = []coretestcases.CaseV1{
	{
		Title:        "AppendPrependIf with both arrays",
		ArrangeInput: args.Map{},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}
