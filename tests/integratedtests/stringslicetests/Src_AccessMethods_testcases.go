package stringslicetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var srcFirstTestCases = []coretestcases.CaseV1{
	{
		Title: "First returns first element -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"result": "a",
		},
	},
}

var srcFirstOrDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstOrDefault returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "FirstOrDefault returns first element -- non-empty",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": "a",
		},
	},
}

var srcLastTestCases = []coretestcases.CaseV1{
	{
		Title: "Last returns last element -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"result": "b",
		},
	},
}

var srcLastOrDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "LastOrDefault returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "LastOrDefault returns last element -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"result": "b",
		},
	},
}

var srcIndexAtTestCases = []coretestcases.CaseV1{
	{
		Title: "IndexAt returns element at index -- index 1",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
			"index": 1,
		},
		ExpectedInput: args.Map{
			"result": "b",
		},
	},
}

var srcSafeIndexAtTestCases = []coretestcases.CaseV1{
	{
		Title: "SafeIndexAt returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
			"index": 0,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "SafeIndexAt returns empty -- negative index",
		ArrangeInput: args.Map{
			"input": []string{"a"},
			"index": -1,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "SafeIndexAt returns empty -- out of bounds",
		ArrangeInput: args.Map{
			"input": []string{"a"},
			"index": 5,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "SafeIndexAt returns element -- valid index",
		ArrangeInput: args.Map{
			"input": []string{"a"},
			"index": 0,
		},
		ExpectedInput: args.Map{
			"result": "a",
		},
	},
}

var srcFirstPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstPtr returns first element -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"result": "a",
		},
	},
}

var srcFirstOrDefaultPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstOrDefaultPtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "FirstOrDefaultPtr returns first element -- non-empty",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": "a",
		},
	},
}

var srcLastPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "LastPtr returns last element -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"result": "b",
		},
	},
}

var srcLastOrDefaultPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "LastOrDefaultPtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "LastOrDefaultPtr returns last element -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"result": "b",
		},
	},
}

var srcFirstOrDefaultWithTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstOrDefaultWith returns default -- nil input",
		ArrangeInput: args.Map{
			"input":   nil,
			"default": "def",
		},
		ExpectedInput: args.Map{
			"result": "def",
			"ok":     false,
		},
	},
	{
		Title: "FirstOrDefaultWith returns first element -- non-empty",
		ArrangeInput: args.Map{
			"input":   []string{"a"},
			"default": "def",
		},
		ExpectedInput: args.Map{
			"result": "a",
			"ok":     true,
		},
	},
}

var srcSafeIndexAtUsingLastIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "SafeIndexAtUsingLastIndex returns element -- valid lastIndex",
		ArrangeInput: args.Map{
			"input":     []string{"a", "b"},
			"index":     1,
			"lastIndex": 0,
		},
		ExpectedInput: args.Map{
			"result": "a",
		},
	},
	{
		Title: "SafeIndexAtUsingLastIndex returns empty -- nil input",
		ArrangeInput: args.Map{
			"input":     nil,
			"index":     0,
			"lastIndex": 0,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "SafeIndexAtUsingLastIndex returns empty -- negative lastIndex",
		ArrangeInput: args.Map{
			"input":     []string{"a"},
			"index":     0,
			"lastIndex": -1,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
}

var srcSafeIndexAtWithTestCases = []coretestcases.CaseV1{
	{
		Title: "SafeIndexAtWith returns default -- nil input",
		ArrangeInput: args.Map{
			"input":   nil,
			"index":   0,
			"default": "def",
		},
		ExpectedInput: args.Map{
			"result": "def",
		},
	},
	{
		Title: "SafeIndexAtWith returns element -- valid index",
		ArrangeInput: args.Map{
			"input":   []string{"a"},
			"index":   0,
			"default": "def",
		},
		ExpectedInput: args.Map{
			"result": "a",
		},
	},
}

var srcSafeIndexAtWithPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "SafeIndexAtWithPtr returns default -- nil input",
		ArrangeInput: args.Map{
			"input":   nil,
			"index":   0,
			"default": "def",
		},
		ExpectedInput: args.Map{
			"result": "def",
		},
	},
	{
		Title: "SafeIndexAtWithPtr returns element -- valid index",
		ArrangeInput: args.Map{
			"input":   []string{"a"},
			"index":   0,
			"default": "def",
		},
		ExpectedInput: args.Map{
			"result": "a",
		},
	},
}

var srcLastIndexPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "LastIndexPtr returns correct index -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"result": 1,
		},
	},
}

var srcLastSafeIndexPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "LastSafeIndexPtr returns correct index -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"result": 1,
		},
	},
}

var srcIndexesDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "IndexesDefault returns elements -- valid indexes",
		ArrangeInput: args.Map{
			"input":   []string{"a", "b", "c"},
			"indexes": []int{0, 2},
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "IndexesDefault returns empty -- nil input",
		ArrangeInput: args.Map{
			"input":   nil,
			"indexes": []int{0},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "IndexesDefault returns empty -- no indexes",
		ArrangeInput: args.Map{
			"input":   []string{"a"},
			"indexes": []int{},
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcSafeRangeItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "SafeRangeItems returns sub-slice -- valid range",
		ArrangeInput: args.Map{
			"input": []string{"a", "b", "c"},
			"start": 0,
			"end":   2,
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
	{
		Title: "SafeRangeItems returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
			"start": 0,
			"end":   2,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcSafeRangeItemsPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "SafeRangeItemsPtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
			"start": 0,
			"end":   2,
		},
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
	{
		Title: "SafeRangeItemsPtr returns sub-slice -- valid range",
		ArrangeInput: args.Map{
			"input": []string{"a", "b", "c"},
			"start": 0,
			"end":   2,
		},
		ExpectedInput: args.Map{
			"length": 2,
		},
	},
}

var srcFirstLastDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstLastDefault returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"first": "",
			"last":  "",
		},
	},
	{
		Title: "FirstLastDefault returns correct values -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"first": "a",
			"last":  "b",
		},
	},
}

var srcFirstLastDefaultPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstLastDefaultPtr returns empty -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"first": "",
			"last":  "",
		},
	},
	{
		Title: "FirstLastDefaultPtr returns correct values -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"first": "a",
			"last":  "b",
		},
	},
}

var srcFirstLastDefaultStatusTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstLastDefaultStatus returns invalid -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"isValid": false,
		},
	},
	{
		Title: "FirstLastDefaultStatus returns hasFirst -- single item",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"isValid":  false,
			"hasFirst": true,
		},
	},
	{
		Title: "FirstLastDefaultStatus returns hasLast -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"isValid": true,
			"hasLast": true,
		},
	},
}

var srcFirstLastDefaultStatusPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "FirstLastDefaultStatusPtr returns invalid -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"isValid": false,
		},
	},
	{
		Title: "FirstLastDefaultStatusPtr returns valid -- two items",
		ArrangeInput: args.Map{
			"input": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"isValid": true,
		},
	},
}
