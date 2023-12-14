package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

var pathInstructionsV2 = []chmodhelper.DirFilesWithRwxPermission{
	{
		DirWithFiles: chmodhelper.DirWithFiles{
			Dir: "/temp/core/test-cases",
			Files: []string{
				"file-1.txt",
				"file-2.txt",
				"file-3.txt",
			},
		},
		ApplyRwx: chmodhelpertestwrappers.DefaultRwx,
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
		ApplyRwx: chmodhelpertestwrappers.DefaultRwx,
	},
	{
		DirWithFiles: chmodhelper.DirWithFiles{
			Dir: "/temp/core/test-cases-3",
			Files: []string{
				"file-1.txt",
				"file-2.txt",
				"file-3.txt",
			},
		},
		ApplyRwx: chmodhelpertestwrappers.DefaultRwx,
	},
}
