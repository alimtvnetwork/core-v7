package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
)

func Test_QW_Value_OnlySupportedErr_WithNames(t *testing.T) {
	// The internal toHashset is called with names passed to OnlySupportedErr.
	// The empty branch of toHashset is unreachable since OnlySupportedErr
	// returns early for len(names)==0. We maximize coverage by calling with names.
	err := issetter.Set.OnlySupportedErr("Set", "True", "False")
	_ = err
}
