package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/filemode"
	"gitlab.com/auk-go/core/internal/pathinternal"
)

var (
	defaultCaseDirV3 = pathinternal.JoinTemp("core", "case-v3")

	pathInstructionsV3 = []chmodhelper.DirFilesWithContent{
		{
			Dir: defaultCaseDirV3,
			Files: []chmodhelper.FileWithContent{
				{
					RelativePath: "file-1.txt",
					FileMode:     filemode.FileDefault,
					Content: []string{
						"some lines",
						"alim",
					},
				},
				{
					RelativePath: "file-2.txt",
					FileMode:     filemode.FileDefault,
					Content: []string{
						"some lines file - 2",
						"alim",
					},
				},
				{
					RelativePath: "file-3.txt",
					FileMode:     filemode.FileDefault,
					Content: []string{
						"some lines file - 3",
						"alim",
					},
				},
			},
			DirFileMode: filemode.DirDefault,
		},
	}
)
