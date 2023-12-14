package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/internal/pathinternal"
)

var (
	dirCreateBasePath = pathinternal.JoinTemp("core", "case-dir-create")

	createDirTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - if",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"/if/some-dir/first.txt",
						"/if/some-dir-2/first.txt",
						"/if/some-dir-3/first.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core/case-dir-create/if/some-dir - isCreated : true",
				"0 - 1 : core/case-dir-create/if/some-dir-2 - isCreated : true",
				"0 - 2 : core/case-dir-create/if/some-dir-3 - isCreated : true",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}

	createDirIfMissingTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - if-missing",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"/if-missing/first.txt",
						"/if-missing/second.txt",
						"/if-missing/third.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core/case-dir-create/if-missing - isCreated : true",
				"0 - 1 : core/case-dir-create/if-missing - isCreated : true",
				"0 - 2 : core/case-dir-create/if-missing - isCreated : true",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}

	createDirDirectTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - direct create - if exist fails",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"/first.txt",
						"/f/first.txt",
						"/s/first.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core/case-dir-create/first.txt - already exist as file, err: dir : , applyChmod :-rwxr-xr-x, path exist but it is not a dir.",
				"0 - 1 : core/case-dir-create/f/first.txt - already exist as file, err: dir : , applyChmod :-rwxr-xr-x, path exist but it is not a dir.",
				"0 - 2 : core/case-dir-create/s/first.txt - already exist as file, err: dir : , applyChmod :-rwxr-xr-x, path exist but it is not a dir.",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}

	createDirByCheckingTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - direct create - by checking",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"by-checking\\a.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core/case-dir-create/by-checking/a.txt - no error during 2nd invoke of createDir.Direct",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}
)
