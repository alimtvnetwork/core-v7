package coretests

import (
	"testing"

	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/osconsts"
)

// Skip on Unix
func SkipOnUnix(t *testing.T) {
	if !osconsts.IsWindows {
		t.Skip(msgtype.UnixIgnoreMessage)
	}
}
