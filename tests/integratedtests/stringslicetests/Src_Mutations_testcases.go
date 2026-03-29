package stringslicetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var srcCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "Clone returns copy -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "ClonePtr returns copy -- one item",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
}

var srcCloneUsingCapTestCases = []coretestcases.CaseV1{
	{
		Title: "CloneUsingCap returns copy -- one item cap 5",
		ArrangeInput: args.Map{
			"input": []string{"a"},
			"cap":   5,
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
	{
		Title: "CloneUsingCap returns empty -- nil input cap 0",
		ArrangeInput: args.Map{
			"input": nil,
			"cap":   0,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcMergeNewTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeNew returns merged -- one slice plus two items",
		ArrangeInput: args.Map{
			"input":  []string{"a"},
			"extras": []string{"b", "c"},
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
	{
		Title: "MergeNew returns empty -- nil no extras",
		ArrangeInput: args.Map{
			"input":  nil,
			"extras": []string{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcMergeNewSimpleTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeNewSimple returns merged -- two slices",
		ArrangeInput: args.Map{
			"left":  []string{"a"},
			"right": []string{"b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcInPlaceReverseTestCases = []coretestcases.CaseV1{
	{
		Title: "InPlaceReverse returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "InPlaceReverse reverses in place -- three items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b", "c"},
		},
		ExpectedInput: args.Map{
			"first": "c",
		},
	},
	{
		Title: "InPlaceReverse reverses in place -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"first": "b",
		},
	},
	{
		Title: "InPlaceReverse keeps single item -- one item",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"first": "a",
		},
	},
}

var srcPrependNewTestCases = []coretestcases.CaseV1{
	{
		Title: "PrependNew returns prepended -- one item prepended",
		ArrangeInput: args.Map{
			"input":   []string{"b"},
			"prepend": []string{"a"},
		},
		ExpectedInput: args.Map{
			"length": 2,
			"first":  "a",
		},
	},
	{
		Title: "PrependNew returns empty -- nil no prepend",
		ArrangeInput: args.Map{
			"input":   nil,
			"prepend": []string{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcAppendLineNewTestCases = []coretestcases.CaseV1{
	{
		Title: "AppendLineNew returns appended -- one item appended",
		ArrangeInput: args.Map{
			"input":  []string{"a"},
			"append": "b",
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcPrependLineNewTestCases = []coretestcases.CaseV1{
	{
		Title: "PrependLineNew returns prepended -- one line prepended",
		ArrangeInput: args.Map{
			"line":  "a",
			"input": []string{"b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcSortIfTestCases = []coretestcases.CaseV1{
	{
		Title: "SortIf sorts -- isSort true",
		ArrangeInput: args.Map{
			"input":  []string{"b", "a"},
			"isSort": true,
		},
		ExpectedInput: args.Map{
			"first": "a",
		},
	},
}

var srcExpandBySplitTestCases = []coretestcases.CaseV1{
	{
		Title: "ExpandBySplit returns expanded -- comma separator",
		ArrangeInput: args.Map{
			"input":     []string{"a,b", "c"},
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"minLength": 3,
		},
	},
	{
		Title: "ExpandBySplit returns empty -- nil input",
		ArrangeInput: args.Map{
			"input":     nil,
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}
