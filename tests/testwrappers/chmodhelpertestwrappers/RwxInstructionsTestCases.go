package chmodhelpertestwrappers

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

var defaultRwx = chmodins.RwxOwnerGroupOther{
	Owner: "rwx",
	Group: "r-x",
	Other: "r--",
}

var defaultExpected = chmodins.RwxOwnerGroupOther{
	Owner: "r-x",
	Group: "r-x",
	Other: "-w-",
}

type RwxCompileValueTestWrapper struct {
	Existing, Input, Expected chmodins.RwxOwnerGroupOther
}

var RwxCompileValueTestCases = []RwxCompileValueTestWrapper{
	{
		Existing: defaultRwx,
		Input: chmodins.RwxOwnerGroupOther{
			Owner: "*-x",
			Group: "**x",
			Other: "-w-",
		},
		Expected: defaultExpected,
	},
	{
		Existing: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r--",
			Other: "--x",
		},
		Input: chmodins.RwxOwnerGroupOther{
			Owner: "***",
			Group: "**x",
			Other: "-w*",
		},
		Expected: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "-wx",
		},
	},
}

// RwxInstructionsApplyTestCases https://ss64.com/bash/chmod.html
var RwxInstructionsApplyTestCases = []RwxInstructionTestWrapper{
	{
		RwxInstructions: []*chmodins.RwxInstruction{
			{
				IsSkipOnNonExist:  false,
				IsContinueOnError: false,
				IsRecursive:       false,
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "*-x",
					Group: "**x",
					Other: "-w-",
				},
			},
		},
		DefaultRwx:      &defaultRwx,
		IsErrorExpected: false,
		CreatePaths: []*CreatePathsInstruction{
			{
				Dir: "/temp/core/test-cases",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
			{
				Dir: "/temp/core/test-cases-2",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
		},
		funcName: RwxApplyOnPath,
		expected: defaultExpected,
	},
}

// RwxInstructionsUnixApplyRecursivelyTestCases https://ss64.com/bash/chmod.html
var RwxInstructionsUnixApplyRecursivelyTestCases = []RwxInstructionTestWrapper{
	{
		RwxInstructions: []*chmodins.RwxInstruction{
			{
				IsSkipOnNonExist:  false,
				IsContinueOnError: false,
				IsRecursive:       true,
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "*-x",
					Group: "**x",
					Other: "-w-",
				},
			},
		},
		DefaultRwx:      &defaultRwx,
		IsErrorExpected: false,
		CreatePaths: []*CreatePathsInstruction{
			{
				Dir: "/temp/core/test-cases",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
			{
				Dir: "/temp/core/test-cases-2",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
			{
				Dir: "/temp/core/test-cases-3",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
		},
		funcName: RwxApplyOnPath,
		expected: defaultExpected,
	},
}
