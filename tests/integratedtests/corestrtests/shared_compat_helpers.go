package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// caseV1Compat is a lightweight assertion helper used across many coverage files.
// Moved here from Coverage41_Iteration37_test.go so split-recovery subfolders can see it.
type caseV1Compat struct {
	Name     string
	Expected any
	Actual   any
	Args     args.Map
}

func (it caseV1Compat) ShouldBeEqual(t *testing.T) {
	expected := args.Map{"value": it.Expected}
	actual := args.Map{"value": it.Actual}
	expected.ShouldBeEqual(t, 0, it.Name, actual)
}

// testErr is a simple error type used in LinkedList/SimpleSlice tests.
// Moved here from Coverage40_Seg5_Hashset_test.go.
type testErr struct{}

func (e *testErr) Error() string { return "test" }
