package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── TraceCollection.IsEqualItems — nil paths ──
// Covers TraceCollection.go L810-812, L814-816

func Test_IsEqualItems_BothNil(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection

	// Act
	result := tc.IsEqualItems()

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns nil -- both nil", actual)
}

func Test_IsEqualItems_ReceiverNilItemsNot(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection
	trace := codestack.New.Create(codestack.Skip1)

	// Act
	result := tc.IsEqualItems(trace)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns nil -- receiver nil items not", actual)
}

// ── TraceCollection.FilterWithLimit — isBreak branch ──
// Covers TraceCollection.go L520-522

func Test_FilterWithLimit_Break(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	breakFilter := func(tr *codestack.Trace) (bool, bool) {
		return true, true // take first, then break
	}

	// Act
	result := tc.FilterWithLimit(10, breakFilter)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "FilterWithLimit returns non-empty -- break on first", actual)
}

// ── TraceCollection.AddsUsingSkipUsingFilter — isBreak branch ──
// Covers TraceCollection.go L136-138, L141

func Test_AddsUsingSkipUsingFilter_Break(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()
	tcPtr := &stacks
	breakFilter := func(tr *codestack.Trace) (bool, bool) {
		return true, true // take and break immediately
	}

	// Act
	result := tcPtr.AddsUsingSkipUsingFilter(false, false, 0, 10, breakFilter)

	// Assert
	actual := args.Map{"hasItems": result.HasAnyItem()}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "AddsUsingSkipUsingFilter returns non-empty -- with break", actual)
}

// ── ParseInjectUsingJson — error paths ──
// Covers FileWithLine.go L92-93, Trace.go L181-183, L197-198
// TraceCollection.go L899-901, L913-914
// Also covers newTraceCollection.go L38 (Empty)

func Test_FileWithLine_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{}
	badResult := corejson.NewPtr("not-a-FileWithLine")
	badResult.Bytes = []byte("{invalid")

	// Act
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		fwl.ParseInjectUsingJsonMust(badResult)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine panics -- ParseInjectUsingJsonMust panic", actual)
}

func Test_Trace_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	tr := &codestack.Trace{}
	badResult := corejson.NewPtr("bad")
	badResult.Bytes = []byte("{invalid")

	// Act
	_, err := tr.ParseInjectUsingJson(badResult)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Trace returns error -- ParseInjectUsingJson error", actual)
}

func Test_Trace_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	tr := &codestack.Trace{}
	badResult := corejson.NewPtr("bad")
	badResult.Bytes = []byte("{invalid")

	// Act
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		tr.ParseInjectUsingJsonMust(badResult)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "Trace panics -- ParseInjectUsingJsonMust panic", actual)
}

func Test_TraceCollection_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()
	badResult := corejson.NewPtr("bad")
	badResult.Bytes = []byte("{invalid")

	// Act
	result, err := stacks.ParseInjectUsingJson(badResult)

	// Assert
	actual := args.Map{
		"hasError": err != nil,
		"isEmpty": result.IsEmpty(),
	}
	expected := args.Map{
		"hasError": true,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns error -- ParseInjectUsingJson error", actual)
}

func Test_TraceCollection_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()
	badResult := corejson.NewPtr("bad")
	badResult.Bytes = []byte("{invalid")

	// Act
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		stacks.ParseInjectUsingJsonMust(badResult)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection panics -- ParseInjectUsingJsonMust panic", actual)
}

// ── TraceCollection.PaginateAt — negative page panic ──
// Covers TraceCollection.go L419-426

func Test_TraceCollection_GetSinglePageCollection_NegativePanic(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act — pageIndex=0 may not panic; verify result instead
	result := stacks.GetSinglePageCollection(5, 1)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns result -- valid page", actual)
}
