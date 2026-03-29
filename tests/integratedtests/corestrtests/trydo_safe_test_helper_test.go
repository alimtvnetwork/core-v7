package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/internal/trydo"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

// safeTest wraps a test body with trydo.Block so that any panic is
// caught and reported as a GoConvey assertion failure instead of
// crashing the entire test binary (which would wipe the coverage profile).
//
// Usage:
//
//	func Test_Something(t *testing.T) {
//	    safeTest(t, "Something", func() {
//	        // ... original test body ...
//	    })
//	}
func safeTest(t *testing.T, name string, fn func()) {
	t.Helper()

	trydo.Block{
		Try: fn,
		Catch: func(e trydo.Exception) {
			msg := fmt.Sprintf("[PANIC RECOVERED] %s: %v", name, e)
			t.Log(msg)

			convey.Convey(
				name+" (panic recovered)", t, func() {
					convey.So(
						msg,
						should.BeEmpty,
					)
				},
			)
		},
	}.Do()
}
