package defaultcapacitytests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var ofSplitsTestCases = []coretestcases.CaseV1{
	{
		Title: "No limit (-1) returns OfSearch result",
		ArrangeInput: args.Map{
			"wholeLength": 100,
			"limit":       -1,
		},
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title: "Limit >= wholeLength returns OfSearch result",
		ArrangeInput: args.Map{
			"wholeLength": 50,
			"limit":       100,
		},
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title: "Limit == wholeLength returns OfSearch result",
		ArrangeInput: args.Map{
			"wholeLength": 50,
			"limit":       50,
		},
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title: "Limit < wholeLength returns limit",
		ArrangeInput: args.Map{
			"wholeLength": 100,
			"limit":       25,
		},
		ExpectedInput: args.Map{"result": 25},
	},
}

var predictiveDefaultSmallTestCases = []coretestcases.CaseV1{
	{
		Title:         "Positive length yields positive result",
		ArrangeInput:  100,
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title:         "Zero length yields positive result from additionalCap",
		ArrangeInput:  0,
		ExpectedInput: args.Map{"isPositive": true},
	},
}

var predictiveFiftyPercentTestCases = []coretestcases.CaseV1{
	{
		Title:         "Positive length with additional cap",
		ArrangeInput:  args.Map{
			"possibleLen": 100,
			"additionalCap": 10,
		},
		ExpectedInput: args.Map{"isPositive": true},
	},
	{
		Title:         "Zero length returns additional cap",
		ArrangeInput:  args.Map{
			"possibleLen": 0,
			"additionalCap": 5,
		},
		ExpectedInput: args.Map{"result": 5},
	},
}
