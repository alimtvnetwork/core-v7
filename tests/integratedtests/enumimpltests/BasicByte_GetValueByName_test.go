package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BasicByte: GetValueByName wrapped-quote path ──
// Covers BasicByte.go L81-83

func Test_BasicByte_GetValueByName_WrappedQuotePath(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"Alpha", "Beta"})

	// Act — The map stores keys as `"Alpha"` (double-quoted).
	// Passing raw "Alpha" should fail direct lookup, then succeed via wrapped lookup.
	val, err := bb.GetValueByName("Alpha")

	// Assert
	actual := args.Map{
		"value": int(val),
		"noErr": err == nil,
	}
	expected := args.Map{
		"value": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns value -- wrapped quote path", actual)
}

// ── BasicInt8: GetValueByName wrapped-quote path ──
// Covers BasicInt8.go L78-80

func Test_BasicInt8_GetValueByName_WrappedQuotePath(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestInt8", []string{"X", "Y"})

	// Act
	val, err := bi.GetValueByName("X")

	// Assert
	actual := args.Map{
		"value": int(val),
		"noErr": err == nil,
	}
	expected := args.Map{
		"value": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns value -- wrapped quote path Int8", actual)
}

// ── BasicInt16: GetValueByName wrapped-quote path ──
// Covers BasicInt16.go L77-79

func Test_BasicInt16_GetValueByName_WrappedQuotePath(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestInt16", []string{"M", "N"})

	// Act
	val, err := bi.GetValueByName("M")

	// Assert
	actual := args.Map{
		"value": int(val),
		"noErr": err == nil,
	}
	expected := args.Map{
		"value": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns value -- wrapped quote path Int16", actual)
}

// ── BasicInt32: GetValueByName wrapped-quote path ──
// Covers BasicInt32.go L75-77

func Test_BasicInt32_GetValueByName_WrappedQuotePath(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestInt32", []string{"P", "Q"})

	// Act
	val, err := bi.GetValueByName("P")

	// Assert
	actual := args.Map{
		"value": int(val),
		"noErr": err == nil,
	}
	expected := args.Map{
		"value": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns value -- wrapped quote path Int32", actual)
}

// ── BasicUInt16: GetValueByName wrapped-quote path ──
// Covers BasicUInt16.go L77-79

func Test_BasicUInt16_GetValueByName_WrappedQuotePath(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicUInt16.UsingTypeSlice("TestUInt16", []string{"R", "S"})

	// Act
	val, err := bi.GetValueByName("R")

	// Assert
	actual := args.Map{
		"value": int(val),
		"noErr": err == nil,
	}
	expected := args.Map{
		"value": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns value -- wrapped quote path UInt16", actual)
}

// ── BasicString: GetValueByName wrapped-quote path ──
// Covers BasicString.go L139-141

func Test_BasicString_GetValueByName_WrappedQuotePath(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateUsingNamesSpread("TestString", "Foo", "Bar")

	// Act
	val, err := bs.GetValueByName("Foo")

	// Assert
	actual := args.Map{
		"value": val,
		"noErr": err == nil,
	}
	expected := args.Map{
		"value": "Foo",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns value -- wrapped quote path String", actual)
}

// ── DynamicMap: IsRawEqual with isRegardlessType=true ──
// Covers DynamicMap.go L867-881 (isEqualSingle regardless path)

func Test_DynamicMap_IsRawEqual_RegardlessType(t *testing.T) {
	// Arrange
	left := enumimpl.DynamicMap{"a": 1, "b": "hello"}
	right := map[string]any{"a": 1, "b": "hello"}

	// Act
	result := left.IsRawEqual(true, right)

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsRawEqual returns true -- regardless type equal values", actual)
}

func Test_DynamicMap_IsRawEqual_RegardlessType_Different(t *testing.T) {
	// Arrange
	left := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 2}

	// Act
	result := left.IsRawEqual(true, right)

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsRawEqual returns false -- regardless type diff values", actual)
}

// ── DynamicMap: DiffRaw with right-side-only keys ──
// Covers DynamicMap.go L442-444 (diffRightSide unequal values)

func Test_DynamicMap_DiffRaw_RightSideDifference(t *testing.T) {
	// Arrange — same keys but different values, and right has extra key
	left := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "b": 99, "c": 3}

	// Act
	diff := left.DiffRaw(false, right)

	// Assert — "b" differs, "c" only in right
	actual := args.Map{
		"hasDiff": diff.HasAnyItem(),
		"length": diff.Length(),
	}
	expected := args.Map{
		"hasDiff": true,
		"length": 2,
	}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns diffs -- right-side differences", actual)
}

// ── DynamicMap: ConvMapStringString with non-string value ──
// Covers DynamicMap.go L1363-1364 (dead code: KeyValueString always returns isFound=true for existing keys)
// This test confirms coverage — the continue branch is actually unreachable.

func Test_DynamicMap_ConvMapStringString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"key1": "val1", "key2": 42}

	// Act
	result := dm.ConvMapStringString()

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "ConvMapStringString converts all entries -- mixed types", actual)
}
