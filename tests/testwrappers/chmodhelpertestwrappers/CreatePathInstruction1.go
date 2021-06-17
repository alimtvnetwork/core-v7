package chmodhelpertestwrappers

var CreatePathInstruction1 = []*CreatePathsInstruction{
	{
		Dir: "/temp/core/test-cases",
		Files: []string{
			"file-1.txt",
			"file-2.txt",
			"file-3.txt",
		},
		ApplyRwx: DefaultRwx,
	},
	{
		Dir: "/temp/core/test-cases-2",
		Files: []string{
			"file-1.txt",
			"file-2.txt",
			"file-3.txt",
		},
		ApplyRwx: DefaultRwx,
	},
}
