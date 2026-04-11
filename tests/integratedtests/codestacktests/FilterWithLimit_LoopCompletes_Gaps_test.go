package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage12 — remaining codestack gaps
//
// Target 1: TraceCollection.FilterWithLimit line 529
//   Final return after loop completes without break or limit hit.
//
// Target 2: dirGetter.CurDir line 22 — runtime.Caller fail fallback (dead code)
// Target 3: fileGetter.CurrentFilePath line 74 — runtime.Caller fail fallback (dead code)
//
// Targets 2 & 3 are unreachable in normal execution (runtime.Caller always
// succeeds from test context). Documented as accepted dead-code gaps.
// ══════════════════════════════════════════════════════════════════════════════

func Test_FilterWithLimit_LoopCompletes(t *testing.T) {
	// Arrange — create a small collection, filter takes all, limit > count
	stacks := codestack.New.StackTrace.SkipNone()
	collection := &stacks

	// Act — limit much larger than items; loop will complete naturally
	result := collection.FilterWithLimit(
		9999,
		func(trace *codestack.Trace) (isTake, isBreak bool) {
			return true, false
		},
	)

	// Assert
	convey.Convey("FilterWithLimit returns all items when loop completes without break", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// Coverage note: dirGetter.CurDir line 22 and fileGetter.CurrentFilePath line 74
// are runtime.Caller failure fallbacks — unreachable in normal test execution.
// Documented as accepted dead-code gaps.
