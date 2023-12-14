package chmodhelpertestwrappers

import "gitlab.com/auk-go/core/chmodhelper/chmodins"

var VerifyRwxChmodUsingRwxInstructionsTestCases = []VerifyRwxChmodUsingRwxInstructionsWrapper{
	{
		Header: "rwx - missing paths",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "rwx",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   false,
				IsContinueOnError: false,
				IsRecursive:       false,
			},
		},
		Locations:            SimpleLocations,
		ExpectedErrorMessage: "Path missing or having other access issues! Ref(s) { \"[/temp/core/test-cases-3s /temp/core/test-cases-3x]\" }",
	},
	{
		Header: "rwx - expectation failed",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   true,
				IsContinueOnError: true,
				IsRecursive:       false,
			},
		},
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Path:/temp/core/test-cases-2 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual\n" +
			"Path:/temp/core/test-cases-3 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
	},
	{
		Header: "Recursive not supported",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   true,
				IsContinueOnError: true,
				IsRecursive:       true,
			},
		},
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Not Supported: Feature or method is not supported yet. " +
			"IsRecursive is not supported for Verify chmod. Ref(s) { " +
			"\"[" +
			"/temp/core/test-cases-2 " +
			"/temp/core/test-cases-3s " +
			"/temp/core/test-cases-3x " +
			"/temp/core/test-cases-3]\" }",
	},
	{
		Header: "Missing paths + Expectation failed",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   false,
				IsContinueOnError: true,
				IsRecursive:       false,
			},
		},
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Path missing or having other access issues! Ref(s) { " +
			"\"[/temp/core/test-cases-3s /temp/core/test-cases-3x]\" " +
			"}\n" +
			"Path:/temp/core/test-cases-2 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual\n" +
			"Path:/temp/core/test-cases-3 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
	},
	{
		Header: "Expectation and missing paths, isContinue false so will fail for missing paths only",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   false,
				IsContinueOnError: false,
				IsRecursive:       false,
			},
		},
		Locations:            SimpleLocations,
		ExpectedErrorMessage: "\"[/temp/core/test-cases-3s /temp/core/test-cases-3x]\" Path Ref(s) access having issues! missing or other { }",
	},
}
