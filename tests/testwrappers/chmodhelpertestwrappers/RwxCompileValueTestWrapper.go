package chmodhelpertestwrappers

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

type RwxCompileValueTestWrapper struct {
	Existing, Input, Expected chmodins.RwxOwnerGroupOther
}
