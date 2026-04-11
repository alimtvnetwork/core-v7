package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newStacksCreator.All with isBreakOnceInvalid=false ──
// Covers TraceCollection.go L75 (continue path on invalid trace)

func Test_AddsUsingSkip_SkipContinue(t *testing.T) {
	// Arrange & Act — large skip index means invalid traces
	// isSkipInvalid=true, isBreakOnceInvalid=false → continue past invalid
	tc := codestack.New.StackTrace.All(true, false, 900, 5)

	// Assert — all traces should be skipped
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "All continues past invalid -- no break", actual)
}

// ── newStacksCreator.DefaultCount — captures current call stack ──
// Covers newTraceCollection.go indirectly

func Test_StackTrace_Default(t *testing.T) {
	// Arrange & Act — use DefaultCount to avoid double-skip
	tc := codestack.New.StackTrace.DefaultCount(0)

	// Assert
	actual := args.Map{"hasItems": tc.HasAnyItem()}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "StackTrace Default returns items", actual)
}

// ── FilterWithLimit: natural loop exhaustion ──
// Covers TraceCollection.go L529 (end-of-loop return)

func Test_FilterWithLimit_NaturalExhaustion(t *testing.T) {
	// Arrange — use DefaultCount which reliably captures frames
	tcVal := codestack.New.StackTrace.DefaultCount(0)
	tc := &tcVal
	takeAll := func(trace *codestack.Trace) (bool, bool) {
		return true, false
	}

	// Act — limit larger than collection
	result := tc.FilterWithLimit(100, takeAll)

	// Assert
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "FilterWithLimit returns all -- natural loop end", actual)
}

// ── GetSinglePageCollection: page=0 causes panic ──
// Covers TraceCollection.go L419-426

func Test_GetSinglePageCollection_NegativePagePanic(t *testing.T) {
	// Arrange — need enough items so length >= eachPageSize, otherwise early return
	tcVal := codestack.New.StackTrace.DefaultCount(0)
	tc := &tcVal

	// Ensure we have at least 5 items (eachPageSize) so the panic path is reachable
	if tc.Length() < 5 {
		// DefaultCount captures runtime stack which typically has enough frames.
		// If not, skip — the panic path requires length >= eachPageSize.
		t.Skip("not enough stack frames to test paging panic path")
	}

	// Act — pageIndex=0 → skipItems = eachPageSize*(0-1) = negative → panic
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		tc.GetSinglePageCollection(5, 0)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection panics on zero pageIndex", actual)
}

// ── AddsUsingSkipUsingFilter: skip-continue and end-of-loop return ──
// Covers TraceCollection.go L119-120 (continue) and L141 (return)

func Test_AddsUsingSkipUsingFilter_SkipContinue(t *testing.T) {
	// Arrange — start with items using DefaultCount
	tcVal := codestack.New.StackTrace.DefaultCount(0)
	tc := &tcVal
	initialLen := tc.Length()
	takeAll := func(trace *codestack.Trace) (bool, bool) {
		return true, false
	}

	// Act — isSkipInvalid=true, isBreakOnceInvalid=false, high skip index
	tc.AddsUsingSkipUsingFilter(true, false, 900, 5, takeAll)

	// Assert — original items preserved
	actual := args.Map{"hasItems": tc.Length() >= initialLen}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "AddsUsingSkipUsingFilter skips invalid -- continues", actual)
}

// ── AddsUsingSkipUsingFilter: valid traces with filter that takes all, no break ──
// Covers TraceCollection.go L141 (end-of-loop return after processing all)

func Test_AddsUsingSkipUsingFilter_AllValid_NoBreak(t *testing.T) {
	// Arrange
	tcVal := codestack.TraceCollection{}
	tc := &tcVal
	takeAll := func(trace *codestack.Trace) (bool, bool) {
		return true, false
	}

	// Act — start from current stack, small count
	tc.AddsUsingSkipUsingFilter(false, false, 1, 3, takeAll)

	// Assert
	actual := args.Map{"hasItems": tc.HasAnyItem()}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "AddsUsingSkipUsingFilter returns at loop end -- no break", actual)
}
