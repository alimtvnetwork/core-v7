package coreindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreindexes"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// HasIndex
// ═══════════════════════════════════════════

func Test_HasIndex_Found(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.HasIndex([]int{1, 2, 3}, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasIndex returns true -- found", actual)
}

func Test_HasIndex_NotFound(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.HasIndex([]int{1, 2, 3}, 99)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasIndex returns false -- not found", actual)
}

func Test_HasIndex_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.HasIndex([]int{}, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasIndex returns false -- empty slice", actual)
}

// ═══════════════════════════════════════════
// IsWithinIndexRange
// ═══════════════════════════════════════════

func Test_IsWithinIndexRange_Within(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.IsWithinIndexRange(2, 5)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsWithinIndexRange returns true -- within", actual)
}

func Test_IsWithinIndexRange_Exact(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.IsWithinIndexRange(4, 5)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsWithinIndexRange returns true -- exact last index", actual)
}

func Test_IsWithinIndexRange_Beyond(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.IsWithinIndexRange(5, 5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsWithinIndexRange returns false -- beyond", actual)
}

// ═══════════════════════════════════════════
// LastIndex
// ═══════════════════════════════════════════

func Test_LastIndex_Normal(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.LastIndex(5)}

	// Assert
	expected := args.Map{"result": 4}
	expected.ShouldBeEqual(t, 0, "LastIndex returns 4 -- length 5", actual)
}

func Test_LastIndex_One(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.LastIndex(1)}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "LastIndex returns 0 -- length 1", actual)
}

// ═══════════════════════════════════════════
// NameByIndex — remaining indexes
// ═══════════════════════════════════════════

func Test_NameByIndex_Second(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.NameByIndex(1)}

	// Assert
	expected := args.Map{"result": "Second"}
	expected.ShouldBeEqual(t, 0, "NameByIndex returns Second -- index 1", actual)
}

func Test_NameByIndex_Fifth(t *testing.T) {
	// Act
	actual := args.Map{"result": coreindexes.NameByIndex(4)}

	// Assert
	expected := args.Map{"result": "Fifth"}
	expected.ShouldBeEqual(t, 0, "NameByIndex returns Fifth -- index 4", actual)
}
