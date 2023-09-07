package coretests

import (
	"testing"

	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/osconsts"
)

// SkipOnUnix Skip on Unix
func SkipOnUnix(t *testing.T) {
	if osconsts.IsUnixGroup {
		t.Skip(errcore.UnixIgnoreType)
	}
}
