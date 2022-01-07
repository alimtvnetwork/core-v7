package eventcategory

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Log:     "Log",
		Success: "Success",
		Error:   "Error",
		Failure: "Failure",
	}

	ErrorMap = map[Variant]bool{
		Failure: true,
		Error:   true,
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		coredynamic.TypeName(Success),
		Ranges[:])
)
