package executionleveltype

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Level1:  "Level1",
		Level2:  "Level2",
		Level3:  "Level3",
		Level4:  "Level4",
		Level5:  "Level5",
		Level6:  "Level6",
		Level7:  "Level7",
		Level8:  "Level8",
		Level9:  "Level9",
		Level10: "Level10",
		Level11: "Level11",
		Level12: "Level12",
		Level13: "Level13",
		Level14: "Level14",
		Level15: "Level15",
		Level16: "Level16",
		Level17: "Level17",
		Level18: "Level18",
		Level19: "Level19",
		Level20: "Level20",
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		coredynamic.TypeName(Level1),
		Ranges[:])
)
