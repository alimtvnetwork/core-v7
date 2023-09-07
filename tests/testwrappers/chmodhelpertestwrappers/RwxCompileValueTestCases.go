package chmodhelpertestwrappers

import "gitlab.com/auk-go/core/chmodhelper/chmodins"

var RwxCompileValueTestCases = []RwxCompileValueTestWrapper{
	{
		Existing: DefaultRwx,
		Input: chmodins.RwxOwnerGroupOther{
			Owner: "*-x",
			Group: "**x",
			Other: "-w-",
		},
		Expected: DefaultExpected,
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
