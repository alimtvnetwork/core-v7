package chmodhelpertestwrappers

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
)

var (
	DefaultRwx = chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r--",
	}

	DefaultExpected = chmodins.RwxOwnerGroupOther{
		Owner: "r-x",
		Group: "r-x",
		Other: "-w-",
	}
)
