package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var collectionBasicTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection returns empty state -- new empty collection",
		ExpectedInput: args.Map{
			"isEmpty":   true,
			"hasItems":  false,
			"len":       0,
			"count":     0,
			"lastIndex": -1,
			"cap":       0,
			"hasIndex0": false,
		},
	},
}

var collectionNilReceiverTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection returns 0 length -- nil receiver",
		ExpectedInput: args.Map{
			"len":     0,
			"isEmpty": true,
		},
	},
}

var collectionAddTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection.Add returns 2 -- two adds",
		ExpectedInput: args.Map{
			"len":    2,
			"capGe2": true,
		},
	},
	{
		Title: "Collection.AddNonEmpty returns 1 -- empty then non-empty",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.AddNonEmptyWhitespace returns 1 -- whitespace then non-empty",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.AddError returns 1 -- nil then error",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.AddIf returns 1 -- false then true",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.AddIfMany returns 2 -- false then true with two items",
		ExpectedInput: args.Map{
			"len": 2,
		},
	},
	{
		Title: "Collection.Adds returns 3 -- three items",
		ExpectedInput: args.Map{
			"len": 3,
		},
	},
	{
		Title: "Collection.AddStrings returns 2 -- two-element slice",
		ExpectedInput: args.Map{
			"len": 2,
		},
	},
	{
		Title: "Collection.AddFunc returns 1 -- one function",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.AddFuncErr returns 1 -- no error",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.AddFuncErr returns 0 -- with error",
		ExpectedInput: args.Map{
			"len":    0,
			"called": true,
		},
	},
	{
		Title: "Collection.AddLock returns 1 -- one item",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.AddsLock returns 2 -- two items",
		ExpectedInput: args.Map{
			"len": 2,
		},
	},
}

var collectionMergeTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection.AddCollection returns 2 -- merge two collections",
		ExpectedInput: args.Map{
			"len":         2,
			"lenAfterAdd": 2,
		},
	},
	{
		Title: "Collection.AddCollections returns 2 -- merge with empty",
		ExpectedInput: args.Map{
			"len": 2,
		},
	},
}

var collectionRemoveTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection.RemoveAt returns 2 -- remove middle",
		ExpectedInput: args.Map{
			"success":  true,
			"len":      2,
			"failNeg":  false,
			"failHigh": false,
		},
	},
}

var collectionQueryTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection.ListStrings returns 1 -- single item",
		ExpectedInput: args.Map{
			"lenList":    1,
			"lenListPtr": 1,
		},
	},
	{
		Title: "Collection.LengthLock returns 2 -- two items",
		ExpectedInput: args.Map{
			"len": 2,
		},
	},
	{
		Title: "Collection.IsEmptyLock returns true -- empty collection",
		ExpectedInput: args.Map{
			"isEmpty": true,
		},
	},
	{
		Title: "Collection.HasIndex returns correct -- various indices",
		ExpectedInput: args.Map{
			"has0":    true,
			"has1":    true,
			"has2":    false,
			"hasNeg1": false,
		},
	},
}

var collectionErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection.AsError returns nil -- empty then non-nil",
		ExpectedInput: args.Map{
			"defaultNilEmpty": true,
			"asErrorNilEmpty": true,
			"defaultNonNil":   true,
		},
	},
	{
		Title: "Collection.ToError returns nil -- empty then non-nil",
		ExpectedInput: args.Map{
			"toErrorNilEmpty":   true,
			"toDefaultNilEmpty": true,
			"toErrorNonNil":     true,
		},
	},
}

var collectionMiscTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection.EachItemSplitBy returns 3 -- comma split",
		ExpectedInput: args.Map{
			"len": 3,
		},
	},
	{
		Title: "Collection.ConcatNew returns 1 -- no extra items",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.ConcatNew returns 3 -- with extra items",
		ExpectedInput: args.Map{
			"len": 3,
		},
	},
	{
		Title: "Collection.IsEquals returns true -- same content",
		ExpectedInput: args.Map{
			"equal": true,
		},
	},
	{
		Title: "Collection.IsEqualsWithSensitive returns correct -- case sensitivity",
		ExpectedInput: args.Map{
			"insensitiveEqual": true,
			"sensitiveEqual":   false,
		},
	},
	{
		Title: "Collection.JsonString returns non-empty -- valid collection",
		ExpectedInput: args.Map{
			"jsonNonEmpty":     true,
			"jsonMustNonEmpty": true,
			"stringJSON":       true,
		},
	},
}

var collectionHashmapTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection.AddHashmapsValues returns 1 -- one entry",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.AddHashmapsKeys returns 1 -- one entry",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
	{
		Title: "Collection.AddPointerCollectionsLock returns 1 -- one item",
		ExpectedInput: args.Map{
			"len": 1,
		},
	},
}
