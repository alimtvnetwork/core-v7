package chmodhelpertestwrappers

import "gitlab.com/auk-go/core/chmodhelper/chmodins"

type RwxCompileValueTestWrapper struct {
	Existing, Input, Expected chmodins.RwxOwnerGroupOther
}
