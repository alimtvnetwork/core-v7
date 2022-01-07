package logtype

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Silent:  "Silent",
		Trace:   "Trace",
		Debug:   "Debug",
		Info:    "Info",
		Warning: "Warning",
		Error:   "Error",
		Fatal:   "Fatal",
		Panic:   "Panic",
	}

	TraceMap = map[Variant]bool{
		Trace: true,
		Debug: true,
		Info:  true,
	}

	ErrorMap = map[Variant]bool{
		Error: true,
		Fatal: true,
		Panic: true,
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		coredynamic.TypeName(Trace),
		Ranges[:])
)
