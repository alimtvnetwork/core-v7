package reflectinternaltests

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"

	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	pascalFuncNameTestCases = []coretestcases.CaseV1{
		{
			Title: "Some",
			ArrangeInput: args.One{
				First: "someName",
			},
			ExpectedInput: []string{
				"0 : someName -> SomeName",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf(args.One{}),
		},
	}
)
