package coretests

import (
	"testing"

	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/osconsts"
)

// SkipOnWindows Skip on Windows
func SkipOnWindows(t *testing.T) {
	if osconsts.IsWindows {
		t.Skip(errcore.WindowsIgnoreMessage)
	}
}
