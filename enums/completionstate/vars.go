package completionstate

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Unknown:               "Unknown",
		Running:               "Running",
		Success:               "Success",
		SuccessWithWarning:    "SuccessWithWarning",
		FailedMiddleWithError: "FailedMiddleWithError",
		CompleteWithError:     "CompleteWithError",
	}

	CompletionMap = map[Variant]bool{
		Success:               true,
		SuccessWithWarning:    true,
		FailedMiddleWithError: true,
		CompleteWithError:     true,
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		coredynamic.TypeName(Success),
		Ranges[:])
)
