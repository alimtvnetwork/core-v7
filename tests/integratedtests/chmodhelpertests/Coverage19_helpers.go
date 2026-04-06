package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/chmodhelper"
)

// mustRwxWrapper creates an RwxWrapper from a 9-char rwx string (without leading hyphen).
// Panics on error — use only in test setup (Arrange phase).
//
// Source API: chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen(string) (RwxWrapper, error)
// Input format: "rwxr-xr-x" (9 chars, no leading hyphen)
func mustRwxWrapper(rwx string) chmodhelper.RwxWrapper {
	wrapper, err := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen(rwx)
	if err != nil {
		panic(err)
	}

	return wrapper
}
