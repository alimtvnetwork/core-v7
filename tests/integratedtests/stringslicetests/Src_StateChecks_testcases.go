package stringslicetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var srcIsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmpty returns true -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEmpty returns true -- empty slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEmpty returns false -- non-empty slice",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

var srcHasAnyItemTestCases = []coretestcases.CaseV1{
	{
		Title: "HasAnyItem returns false -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "HasAnyItem returns true -- non-empty slice",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var srcEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty returns empty slice -- no input",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}

var srcIsEmptyPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmptyPtr returns true -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEmptyPtr returns true -- empty slice",
		ArrangeInput: args.Map{
			"input": []string{},
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsEmptyPtr returns false -- non-empty slice",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

var srcHasAnyItemPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "HasAnyItemPtr returns false -- nil input",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "HasAnyItemPtr returns true -- non-empty slice",
		ArrangeInput: args.Map{
			"input": []string{"a"},
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var srcEmptyPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "EmptyPtr returns empty slice -- no input",
		ExpectedInput: args.Map{
			"length": 0,
		},
	},
}
