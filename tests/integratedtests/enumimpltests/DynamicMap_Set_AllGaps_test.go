package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage20 — enumimpl remaining 20 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── DynamicMap.Set on nil receiver (line 26-28) ──

func Test_DynamicMap_Set_NilPtr(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap

	// Act
	result := dm.Set("key", "val")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	actual.ShouldBeEqual(t, 1, "DynamicMap Set nil pointer", expected)
}

// ── DynamicMap.AddNewOnly on nil receiver (line 44-46) ──

func Test_DynamicMap_AddNewOnly_NilPtr(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap

	// Act
	result := dm.AddNewOnly("key", "val")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	actual.ShouldBeEqual(t, 1, "DynamicMap AddNewOnly nil pointer", expected)
}

// ── DynamicMap.Set on nil map (creates new map) (line 26-28 else) ──

func Test_DynamicMap_Set_NilMap_CreatesNew(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap(nil)

	// Act
	result := dm.Set("key", "val")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	actual.ShouldBeEqual(t, 1, "DynamicMap Set nil map creates new", expected)
}

// ── DynamicMap.isEqualSingle regardless type (line 873-887) ──

func Test_DynamicMap_IsEqual_RegardlessType(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	other := enumimpl.DynamicMap{"a": "1"}

	// Act — compare int 1 vs string "1" regardless of type
	result := dm.IsEqual(true, &other)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	actual.ShouldBeEqual(t, 1, "DynamicMap IsEqual regardless type", expected)
}

// ── DynamicMap.DiffRawUsingDifferChecker with custom checker (line 448-450) ──

func Test_DynamicMap_DiffRawUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "b": 3}

	// Act
	diffMap := dm.DiffRawUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl,
		false,
		right,
	)

	// Assert
	actual := args.Map{"hasDiff": len(diffMap) > 0}
	expected := args.Map{"hasDiff": true}
	actual.ShouldBeEqual(t, 1, "DynamicMap DiffRawUsingDifferChecker", expected)
}

// ── toHashset empty input (line 4-6) ──
// This is an unexported function — tested indirectly via enum creation

// ── toStringPrintableDynamicMap empty (line 11-13) ──
// Unexported — tested indirectly

// ── BasicByte/Int8/Int16/Int32/UInt16/String: isFoundByWrapped unmarshal path ──
// These branches require a JSON value wrapped in quotes that exists in the
// double-quote hash map but NOT in the standard map. This is structurally
// unreachable in normal usage — documented as accepted dead code.

// ── newBasicStringCreator: sliceNamesToMap loop (line 302-311) ──
// Tested indirectly via BasicString creation

func Test_BasicString_UnmarshalToValue_NotFound(t *testing.T) {
	// Arrange
	// Create a BasicString enum with known names
	bs := enumimpl.New.BasicString.Create(
		"TestEnum",
		[]string{"Alpha", "Beta", "Gamma"},
	)

	// Act — unmarshal with a name that doesn't exist
	val, err := bs.UnmarshallToValue(false, []byte(`"Nonexistent"`))

	// Assert
	actual := args.Map{
		"val":      val,
		"hasError": err != nil,
	}
	expected := args.Map{
		"val":      "",
		"hasError": true,
	}
	actual.ShouldBeEqual(t, 1, "BasicString UnmarshalToValue not found", expected)
}
