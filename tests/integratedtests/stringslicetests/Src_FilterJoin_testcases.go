package stringslicetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var srcNonEmptySliceTestCases = []coretestcases.CaseV1{
	{
		Title: "NonEmptySlice returns filtered -- mixed items",
		ArrangeInput: args.Map{
			"input": []string{"a", "", "b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "NonEmptySlice returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcNonEmptySlicePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "NonEmptySlicePtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "NonEmptySlicePtr returns filtered -- mixed items",
		ArrangeInput: args.Map{
			"input": []string{"a", "", "b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcNonWhitespaceTestCases = []coretestcases.CaseV1{
	{
		Title: "NonWhitespace returns filtered -- mixed items",
		ArrangeInput: args.Map{
			"input": []string{"a", "  ", "b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "NonWhitespace returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "NonWhitespace returns empty -- empty slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcNonWhitespacePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "NonWhitespacePtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "NonWhitespacePtr returns filtered -- mixed items",
		ArrangeInput: args.Map{
			"input": []string{"a", "  ", "b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcNonNullStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "NonNullStrings returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "NonNullStrings returns filtered -- mixed items",
		ArrangeInput: args.Map{
			"input": []string{"a", "", "b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcNonEmptyStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "NonEmptyStrings returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "NonEmptyStrings returns empty -- empty slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "NonEmptyStrings returns filtered -- mixed items",
		ArrangeInput: args.Map{
			"input": []string{"a", "", "b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcNonEmptyIfTestCases = []coretestcases.CaseV1{
	{
		Title: "NonEmptyIf returns filtered -- isFilter true",
		ArrangeInput: args.Map{
			"input":    []string{"a", "", "b"},
			"isFilter": true,
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcNonEmptyJoinTestCases = []coretestcases.CaseV1{
	{
		Title: "NonEmptyJoin returns joined -- mixed items",
		ArrangeInput: args.Map{
			"input":     []string{"a", "", "b"},
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"nonEmpty": true,
		},
	},
	{
		Title: "NonEmptyJoin returns empty -- nil input",
		ArrangeInput: args.Map{
			"input":     nil,
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "NonEmptyJoin returns empty -- empty slice",
		ArrangeInput: args.Map{
			"input":     []string{},
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
}

var srcNonEmptyJoinPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "NonEmptyJoinPtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input":     nil,
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "NonEmptyJoinPtr returns joined -- mixed items",
		ArrangeInput: args.Map{
			"input":     []string{"a", "", "b"},
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"nonEmpty": true,
		},
	},
}

var srcNonWhitespaceJoinTestCases = []coretestcases.CaseV1{
	{
		Title: "NonWhitespaceJoin returns empty -- nil input",
		ArrangeInput: args.Map{
			"input":     nil,
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "NonWhitespaceJoin returns joined -- mixed items",
		ArrangeInput: args.Map{
			"input":     []string{"a", " ", "b"},
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"nonEmpty": true,
		},
	},
}

var srcNonWhitespaceJoinPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "NonWhitespaceJoinPtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input":     nil,
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "NonWhitespaceJoinPtr returns joined -- mixed items",
		ArrangeInput: args.Map{
			"input":     []string{"a", " ", "b"},
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"nonEmpty": true,
		},
	},
}

var srcJoinWithTestCases = []coretestcases.CaseV1{
	{
		Title: "JoinWith returns joined -- two items",
		ArrangeInput: args.Map{
			"separator": ",",
			"items":     []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"nonEmpty": true,
		},
	},
	{
		Title: "JoinWith returns empty -- no items",
		ArrangeInput: args.Map{
			"separator": ",",
			"items":     []string{},
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
}

var srcJoinsTestCases = []coretestcases.CaseV1{
	{
		Title: "Joins returns joined -- two items comma",
		ArrangeInput: args.Map{
			"separator": ",",
			"items":     []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"result": "a,b",
		},
	},
	{
		Title: "Joins returns empty -- no items",
		ArrangeInput: args.Map{
			"separator": ",",
			"items":     []string{},
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
}

var srcTrimmedEachWordsTestCases = []coretestcases.CaseV1{
	{
		Title: "TrimmedEachWords returns trimmed non-empty -- mixed items",
		ArrangeInput: args.Map{
			"input": []string{" a ", " ", "b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "TrimmedEachWords returns nil -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
	{
		Title: "TrimmedEachWords returns empty -- empty slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcTrimmedEachWordsPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "TrimmedEachWordsPtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "TrimmedEachWordsPtr returns trimmed -- mixed items",
		ArrangeInput: args.Map{
			"input": []string{" a ", "b"},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcTrimmedEachWordsIfTestCases = []coretestcases.CaseV1{
	{
		Title: "TrimmedEachWordsIf returns trimmed -- isTrim true",
		ArrangeInput: args.Map{
			"input":  []string{" a ", ""},
			"isTrim": true,
		},
		ExpectedInput: args.Map{
			"length": 1,
		},
	},
}

var srcSplitContentsByWhitespaceTestCases = []coretestcases.CaseV1{
	{
		Title: "SplitContentsByWhitespace returns split -- space separated",
		ArrangeInput: args.Map{
			"input": "a b c",
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}

var srcSplitTrimmedNonEmptyAllTestCases = []coretestcases.CaseV1{
	{
		Title: "SplitTrimmedNonEmptyAll returns split trimmed -- comma separated",
		ArrangeInput: args.Map{
			"input":     "a, b, c",
			"separator": ",",
		},
		ExpectedInput: args.Map{
			"length": 3,
		},
	},
}
