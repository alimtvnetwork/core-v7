package coretests

import (
	"testing"

	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/osconsts"
)

// SkipOnWindows Skip on Windows
func SkipOnWindows(t *testing.T) {
	if osconsts.IsWindows {
		t.Skip(errcore.WindowsIgnoreType)
	}
}
