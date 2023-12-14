package chmodhelpertestwrappers

import (
	"gitlab.com/auk-go/core/chmodhelper"
)

var PathInstructionsV1 = []chmodhelper.DirFilesWithRwxPermission{
	{
		DirWithFiles: chmodhelper.DirWithFiles{
			Dir: "/temp/core/test-cases",
			Files: []string{
				"file-1.txt",
				"file-2.txt",
				"file-3.txt",
			},
		},
		ApplyRwx: DefaultRwx,
	},
	{
		DirWithFiles: chmodhelper.DirWithFiles{
			Dir: "/temp/core/test-cases-2",
			Files: []string{
				"file-1.txt",
				"file-2.txt",
				"file-3.txt",
			},
		},
		ApplyRwx: DefaultRwx,
	},
}
