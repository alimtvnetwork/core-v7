package versionindexes

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	UptoBuildIndexes = []int{
		Major.ValueInt(),
		Minor.ValueInt(),
		Patch.ValueInt(),
		Build.ValueInt(),
	}

	UptoPatchIndexes = [...]int{
		Major.ValueInt(),
		Minor.ValueInt(),
		Patch.ValueInt(),
	}

	UptoMinorIndexes = []int{
		Major.ValueInt(),
		Minor.ValueInt(),
	}

	UptoMajorIndexes = []int{
		Major.ValueInt(),
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

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		reflectinternal.TypeName(Major),
		Ranges[:])
)
