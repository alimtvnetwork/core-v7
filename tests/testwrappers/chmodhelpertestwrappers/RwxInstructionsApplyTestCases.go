package chmodhelpertestwrappers

import "gitlab.com/auk-go/core/chmodhelper/chmodins"

// RwxInstructionsApplyTestCases https://ss64.com/bash/chmod.html
var RwxInstructionsApplyTestCases = []RwxInstructionTestWrapper{
	{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				Condition: chmodins.Condition{
					IsSkipOnInvalid:   false,
					IsContinueOnError: false,
					IsRecursive:       false,
				},
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "*-x",
					Group: "**x",
					Other: "-w-",
				},
			},
		},
		DefaultRwx:      &DefaultRwx,
		IsErrorExpected: false,
		CreatePaths:     CreatePathInstruction1,
		funcName:        RwxApplyOnPath,
		expected:        DefaultExpected,
	},
}
