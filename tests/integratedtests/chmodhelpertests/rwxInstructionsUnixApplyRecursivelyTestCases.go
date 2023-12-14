package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

// rwxInstructionsUnixApplyRecursivelyTestCases https://ss64.com/bash/chmod.html
var rwxInstructionsUnixApplyRecursivelyTestCases = []chmodhelpertestwrappers.RwxInstructionTestWrapper{
	{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				Condition: chmodins.Condition{
					IsSkipOnInvalid:   false,
					IsContinueOnError: false,
					IsRecursive:       true,
				},
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "*-x",
					Group: "**x",
					Other: "-w-",
				},
			},
		},
		DefaultRwx:      &chmodhelpertestwrappers.DefaultRwx,
		IsErrorExpected: false,
		CreatePaths:     pathInstructionsV2,
		TestFuncName:    chmodhelpertestwrappers.RwxApplyOnPath,
		WhatIsExpected:  chmodhelpertestwrappers.DefaultExpected,
	},
}
