package chmodhelpertestwrappers

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodclasstype"
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
