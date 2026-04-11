package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── TraceCollection via StackTrace ──

func Test_TraceCollection_Default(t *testing.T) {
	// Arrange & Act — use exported StackTrace to get a TraceCollection
	tc := codestack.New.StackTrace.Default(0, 5)

	// Assert
	actual := args.Map{"hasItems": tc.Length() > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "StackTrace Default returns non-empty collection", actual)
}

func Test_TraceCollection_Empty(t *testing.T) {
	// Arrange & Act — use zero-value TraceCollection
	tc := codestack.TraceCollection{}

	// Assert
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection zero-value returns empty -- default", actual)
}

func Test_TraceCollection_Add_FromTraceCollectionDefau(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	trace := codestack.New.Default()

	// Act
	tc.Add(trace)

	// Assert
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection Add increases length -- single trace", actual)
}

func Test_GetSinglePageCollection_ZeroPagePanic(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())

	// Act
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		tc.GetSinglePageCollection(2, 0)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection panics -- zero pageIndex", actual)
}
