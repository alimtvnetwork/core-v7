package chmodhelpertests

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	dirFilesWithContentCreateReadTestCases = []coretestcases.CaseV1{
		{
			Title: "DirFilesWithContent - testing any file reading - writing",
			ArrangeInput: []args.One{
				{
					First: pathInstructionsV3,
				},
			},
			ExpectedInput: []string{
				"0 : file-1.txt",
				"         0. some lines",
				"         1. alim",
				"0 : file-2.txt",
				"         0. some lines file - 2",
				"         1. alim",
				"0 : file-3.txt",
				"         0. some lines file - 3",
				"         1. alim",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.One{}),
		},
	}
)
