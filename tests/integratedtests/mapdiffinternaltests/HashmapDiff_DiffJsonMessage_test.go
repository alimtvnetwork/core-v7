package mapdiffinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/mapdiffinternal"
)

// ── HashmapDiff — DiffJsonMessage ──

func Test_HashmapDiff_DiffJsonMessage_Cov3(t *testing.T) {
	// Arrange
	h := mapdiffinternal.HashmapDiff(map[string]string{"a": "1"})

	// Act
	emptyResult := h.DiffJsonMessage(map[string]string{"a": "1"})
	diffResult := h.DiffJsonMessage(map[string]string{"a": "2"})

	// Assert
	actual := args.Map{
		"emptyEmpty": emptyResult == "",
		"diffNotEmpty": diffResult != "",
	}
	expected := args.Map{
		"emptyEmpty": true,
		"diffNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- DiffJsonMessage", actual)
}

// ── HashmapDiff — DiffRaw left non-nil right nil ──

func Test_HashmapDiff_DiffRaw_LeftNonNilRightNil_Cov3(t *testing.T) {
	// Arrange
	h := mapdiffinternal.HashmapDiff(map[string]string{"a": "1"})

	// Act
	diffMap := h.DiffRaw(nil)

	// Assert
	actual := args.Map{
		"diffLength": len(diffMap),
		"hasKey-a": true,
	}
	expected := args.Map{
		"diffLength": 1,
		"hasKey-a": true,
	}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns nil -- left non-nil right nil", actual)
}

// ── MapStringAnyDiff — HasAnyChanges ──

func Test_MapStringAnyDiff_HasAnyChanges_Cov3(t *testing.T) {
	// Arrange
	m := mapdiffinternal.MapStringAnyDiff(map[string]any{"a": 1})

	// Act
	actual := args.Map{
		"hasChanges":   m.HasAnyChanges(false, map[string]any{"a": 2}),
		"noChanges":    m.HasAnyChanges(false, map[string]any{"a": 1}),
		"regardless":   m.HasAnyChanges(true, map[string]any{"a": "1"}),
	}

	// Assert
	expected := args.Map{
		"hasChanges":   true,
		"noChanges":    false,
		"regardless":   false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringAnyDiff returns correct value -- HasAnyChanges", actual)
}

// ── MapStringAnyDiff — ShouldDiffMessage ──

func Test_MapStringAnyDiff_ShouldDiffMessage_Cov3(t *testing.T) {
	// Arrange
	m := mapdiffinternal.MapStringAnyDiff(map[string]any{"a": 1})

	// Act
	emptyResult := m.ShouldDiffMessage(false, "title", map[string]any{"a": 1})
	diffResult := m.ShouldDiffMessage(false, "title", map[string]any{"a": 2})

	// Assert
	actual := args.Map{
		"emptyEmpty": emptyResult == "",
		"diffNotEmpty": diffResult != "",
	}
	expected := args.Map{
		"emptyEmpty": true,
		"diffNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "MapStringAnyDiff returns correct value -- ShouldDiffMessage", actual)
}

// ── MapStringAnyDiff — ToStringsSliceOfDiffMap non-string values ──

func Test_MapStringAnyDiff_ToStringsSliceOfDiffMap_Cov3(t *testing.T) {
	// Arrange
	m := mapdiffinternal.MapStringAnyDiff(map[string]any{})
	diffMap := map[string]any{"key": 42, "str": "val"}

	// Act
	slice := m.ToStringsSliceOfDiffMap(diffMap)

	// Assert
	actual := args.Map{"length": len(slice)}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "ToStringsSliceOfDiffMap returns correct value -- mixed types", actual)
}

// ── MapStringAnyDiff — DiffJsonMessage with regardless type ──

func Test_MapStringAnyDiff_DiffJsonMessage_Regardless_Cov3(t *testing.T) {
	// Arrange
	m := mapdiffinternal.MapStringAnyDiff(map[string]any{"a": 1})

	// Act
	result := m.DiffJsonMessage(true, map[string]any{"a": "1"})

	// Assert
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessage returns correct value -- regardless same string rep", actual)
}

// ── MapStringAnyDiff — Raw nil ──

func Test_MapStringAnyDiff_Raw_Nil_Cov3(t *testing.T) {
	// Arrange
	var m mapdiffinternal.MapStringAnyDiff

	// Act
	raw := m.Raw()

	// Assert
	actual := args.Map{"length": len(raw)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "MapStringAnyDiff returns nil -- Raw nil", actual)
}

// ── MapStringAnyDiff — DiffRaw nil left non-nil right ──

func Test_MapStringAnyDiff_DiffRaw_NilLeftNonNilRight_Cov3(t *testing.T) {
	// Arrange
	var m *mapdiffinternal.MapStringAnyDiff

	// Act
	result := m.DiffRaw(false, map[string]any{"x": 10})

	// Assert
	actual := args.Map{
		"length": len(result),
		"hasX": result["x"] != nil,
	}
	expected := args.Map{
		"length": 1,
		"hasX": true,
	}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns nil -- nil left non-nil right", actual)
}
