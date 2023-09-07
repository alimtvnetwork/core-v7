package chmodhelpertestwrappers

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/chmodhelper/chmodclasstype"
)

var SingleRwxTestCases = []chmodhelper.SingleRwx{
	{
		Rwx:       "r-x",
		ClassType: chmodclasstype.Other,
	},
	{
		Rwx:       "---",
		ClassType: chmodclasstype.Other,
	},
	{
		Rwx:       "--x",
		ClassType: chmodclasstype.Other,
	},
	{
		Rwx:       "r-x",
		ClassType: chmodclasstype.Other,
	},
}
