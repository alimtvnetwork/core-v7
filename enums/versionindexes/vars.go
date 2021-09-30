package versionindexes

import (
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

var (
	UptoBuildIndexes = []int{
		Major,
		Minor,
		Patch,
		Build,
	}
	UptoPatchIndexes = []int{
		Major,
		Minor,
		Patch,
	}
	UptoMinorIndexes = []int{
		Major,
		Minor,
	}
	UptoMajorIndexes = []int{
		Major,
	}

	Ranges = [...]string{
		Major: "Major",
		Minor: "Minor",
		Patch: "Patch",
		Build: "Build",
	}

	AllVersionIndexes = []Index{
		Major,
		Minor,
		Patch,
		Build,
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		reflectinternal.TypeName(Major),
		Ranges[:])
)
