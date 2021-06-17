package chmodhelpertestwrappers

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

// RwxInstructionsUnixApplyRecursivelyTestCases https://ss64.com/bash/chmod.html
var RwxInstructionsUnixApplyRecursivelyTestCases = []RwxInstructionTestWrapper{
	{
		RwxInstructions: []*chmodins.RwxInstruction{
			{
				Condition: chmodins.Condition{
					IsSkipOnNonExist:  false,
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
		DefaultRwx:      &DefaultRwx,
		IsErrorExpected: false,
		CreatePaths:     CreatePathInstruction2,
		funcName:        RwxApplyOnPath,
		expected:        DefaultExpected,
	},
}
