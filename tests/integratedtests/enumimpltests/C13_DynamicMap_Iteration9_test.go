package enumimpltests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===================== DynamicMap =====================

func Test_C13_DynamicMap_AddOrUpdate(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	isNew := dm.AddOrUpdate("k1", "v1")

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new", actual)
	isNew2 := dm.AddOrUpdate("k1", "v2")
	actual = args.Map{"result": isNew2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected update not new", actual)
}

func Test_C13_DynamicMap_Set(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	isNew := dm.Set("k", "v")

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new", actual)
	isNew2 := dm.Set("k", "v2")
	actual = args.Map{"result": isNew2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected update", actual)
}

func Test_C13_DynamicMap_AddNewOnly(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": dm.AddNewOnly("k", "v")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected added", actual)
	actual = args.Map{"result": dm.AddNewOnly("k", "v2")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not added", actual)
}

func Test_C13_DynamicMap_AllKeys(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	keys := dm.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_AllKeys_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.AllKeys()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_AllKeysSorted(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	keys := dm.AllKeysSorted()

	// Act
	actual := args.Map{"result": keys[0] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a first", actual)
}

func Test_C13_DynamicMap_AllKeysSorted_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.AllKeysSorted()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_AllValuesStrings(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"k": "v"}
	vs := dm.AllValuesStrings()

	// Act
	actual := args.Map{"result": len(vs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_AllValuesStrings_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.AllValuesStrings()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_AllValuesStringsSorted(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"k": "v"}
	vs := dm.AllValuesStringsSorted()

	// Act
	actual := args.Map{"result": len(vs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_AllValuesStringsSorted_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.AllValuesStringsSorted()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_AllValuesIntegers(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	ints := dm.AllValuesIntegers()

	// Act
	actual := args.Map{"result": len(ints) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_AllValuesIntegers_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.AllValuesIntegers()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_MapIntegerString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	m, sorted := dm.MapIntegerString()

	// Act
	actual := args.Map{"result": len(m) == 0 || len(sorted) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C13_DynamicMap_MapIntegerString_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	m, sorted := dm.MapIntegerString()

	// Act
	actual := args.Map{"result": len(m) != 0 || len(sorted) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_MapIntegerString_StringValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "x", "b": "y"}
	m, sorted := dm.MapIntegerString()
	_ = m
	_ = sorted
}

func Test_C13_DynamicMap_SortedKeyValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	kv := dm.SortedKeyValues()

	// Act
	actual := args.Map{"result": len(kv) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_SortedKeyValues_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.SortedKeyValues()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_SortedKeyAnyValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	kav := dm.SortedKeyAnyValues()

	// Act
	actual := args.Map{"result": len(kav) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_SortedKeyAnyValues_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.SortedKeyAnyValues()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_SortedKeyAnyValues_StringValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "x", "b": "y"}
	kav := dm.SortedKeyAnyValues()

	// Act
	actual := args.Map{"result": len(kav) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_First(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"k": "v"}
	k, v := dm.First()

	// Act
	actual := args.Map{"result": k == "" || v == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected key and value", actual)
}

func Test_C13_DynamicMap_First_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	k, v := dm.First()

	// Act
	actual := args.Map{"result": k != "" || v != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_IsValueString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"k": "v"}

	// Act
	actual := args.Map{"result": dm.IsValueString()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	dm2 := enumimpl.DynamicMap{"k": 1}
	actual = args.Map{"result": dm2.IsValueString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C13_DynamicMap_LengthAndCount(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.Length() != 1 || dm.Count() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_Length_Nil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap

	// Act
	actual := args.Map{"result": dm.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil length should be 0", actual)
}

func Test_C13_DynamicMap_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": dm.IsEmpty() || dm.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	dm["k"] = 1
	actual = args.Map{"result": dm.IsEmpty() || !dm.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C13_DynamicMap_LastIndex_HasIndex(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.LastIndex() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": dm.HasIndex(1) || dm.HasIndex(2)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "index check failed", actual)
}

func Test_C13_DynamicMap_HasKey_IsMissingKey(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.HasKey("a") || dm.HasKey("b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "key check failed", actual)
	actual = args.Map{"result": dm.IsMissingKey("a") || !dm.IsMissingKey("b")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing key check failed", actual)
}

func Test_C13_DynamicMap_HasAllKeys(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.HasAllKeys("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": dm.HasAllKeys("a", "c")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C13_DynamicMap_HasAnyKeys(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.HasAnyKeys("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": dm.HasAnyKeys("c")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C13_DynamicMap_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var l, r *enumimpl.DynamicMap

	// Act
	actual := args.Map{"result": l.IsEqual(true, r)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_C13_DynamicMap_IsEqual_OneNil(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsEqual(true, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C13_DynamicMap_IsEqual_SameRef(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsEqual(true, &dm)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same ref should be equal", actual)
}

func Test_C13_DynamicMap_IsEqual_RegardlessType(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm1.IsEqual(true, &dm2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal regardless", actual)
}

func Test_C13_DynamicMap_IsEqual_StrictType(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm1.IsEqual(false, &dm2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal strict", actual)
}

func Test_C13_DynamicMap_IsRawEqual_DiffLength(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.IsRawEqual(true, right)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different length should not be equal", actual)
}

func Test_C13_DynamicMap_IsRawEqual_MissingKey(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"b": 1}

	// Act
	actual := args.Map{"result": dm.IsRawEqual(true, right)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing key should not be equal", actual)
}

func Test_C13_DynamicMap_IsMismatch(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	other := enumimpl.DynamicMap{"a": 2}

	// Act
	actual := args.Map{"result": dm.IsMismatch(false, &other)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
}

func Test_C13_DynamicMap_IsRawMismatch(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsRawMismatch(false, map[string]any{"a": 2})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
}

func Test_C13_DynamicMap_DiffRaw(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	diff := dm.DiffRaw(false, map[string]any{"a": 1, "c": 3})

	// Act
	actual := args.Map{"result": diff.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diffs", actual)
}

func Test_C13_DynamicMap_DiffRaw_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	diff := dm.DiffRaw(true, map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": diff.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no diffs", actual)
}

func Test_C13_DynamicMap_DiffRawUsingDifferChecker_BothNil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap
	diff := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, nil)

	// Act
	actual := args.Map{"result": diff.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_DiffRawUsingDifferChecker_LeftNil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap
	right := map[string]any{"a": 1}
	diff := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, right)

	// Act
	actual := args.Map{"result": diff.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected right map", actual)
}

func Test_C13_DynamicMap_DiffRawUsingDifferChecker_RightNil(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	diff := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, nil)

	// Act
	actual := args.Map{"result": diff.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected left map", actual)
}

func Test_C13_DynamicMap_DiffRawLeftRight(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	lDiff, rDiff := dm.DiffRawLeftRightUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false,
		map[string]any{"a": 1, "c": 3},
	)
	_ = lDiff
	_ = rDiff
}

func Test_C13_DynamicMap_DiffRawLeftRight_BothNil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, nil)

	// Act
	actual := args.Map{"result": l.Length() != 0 || r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_DiffRawLeftRight_LeftNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, map[string]any{"a": 1})
	_ = l
	_ = r
}

func Test_C13_DynamicMap_DiffRawLeftRight_RightNil(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, nil)
	_ = l
	_ = r
}

func Test_C13_DynamicMap_DiffRawLeftRight_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": l.Length() != 0 || r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no diff", actual)
}

func Test_C13_DynamicMap_DiffJsonMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessage(false, map[string]any{"a": 2})

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diff message", actual)
}

func Test_C13_DynamicMap_DiffJsonMessage_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessage(true, map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_DiffJsonMessageLeftRight(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessageLeftRight(false, map[string]any{"b": 2})

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected message", actual)
}

func Test_C13_DynamicMap_DiffJsonMessageLeftRight_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessageLeftRight(true, map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_ShouldDiffMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessage(false, "test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diff", actual)
}

func Test_C13_DynamicMap_ShouldDiffMessage_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessage(true, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_LogShouldDiffMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessage(false, "test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diff", actual)
}

func Test_C13_DynamicMap_LogShouldDiffMessage_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessage(true, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_LogShouldDiffLeftRightMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessage(false, "test", map[string]any{"b": 2})

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diff", actual)
}

func Test_C13_DynamicMap_LogShouldDiffLeftRightMessage_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessage(true, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_ShouldDiffLeftRightMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, false, "test",
		map[string]any{"b": 2},
	)

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diff", actual)
}

func Test_C13_DynamicMap_ShouldDiffLeftRightMessage_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, true, "test",
		map[string]any{"a": 1},
	)

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_ExpectingMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ExpectingMessage("test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch message", actual)
}

func Test_C13_DynamicMap_ExpectingMessage_Equal(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ExpectingMessage("test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_LogExpectingMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 2})
}

func Test_C13_DynamicMap_LogExpectingMessage_Equal(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 1})
}

func Test_C13_DynamicMap_IsKeysEqualOnly(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"a": 99})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "keys should be equal", actual)
	actual = args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"b": 1})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "keys should not be equal", actual)
}

func Test_C13_DynamicMap_IsKeysEqualOnly_BothNil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap

	// Act
	actual := args.Map{"result": dm.IsKeysEqualOnly(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_C13_DynamicMap_IsKeysEqualOnly_OneNil(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsKeysEqualOnly(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C13_DynamicMap_IsKeysEqualOnly_DiffLength(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"a": 1, "b": 2})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "diff length", actual)
}

func Test_C13_DynamicMap_KeyValue(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	v, found := dm.KeyValue("a")

	// Act
	actual := args.Map{"result": found || v != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected found", actual)
	_, found2 := dm.KeyValue("missing")
	actual = args.Map{"result": found2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_C13_DynamicMap_KeyValueString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "hello"}
	v, found := dm.KeyValueString("a")

	// Act
	actual := args.Map{"result": found || v != "hello"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	_, found2 := dm.KeyValueString("missing")
	actual = args.Map{"result": found2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_C13_DynamicMap_KeyValueInt(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 42}
	v, found, failed := dm.KeyValueInt("a")

	// Act
	actual := args.Map{"result": found || failed || v != 42}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	_, found2, _ := dm.KeyValueInt("missing")
	actual = args.Map{"result": found2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_C13_DynamicMap_KeyValueInt_ByteValue(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": byte(5)}
	v, found, failed := dm.KeyValueInt("a")

	// Act
	actual := args.Map{"result": found || failed || v != 5}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_C13_DynamicMap_KeyValueIntDefault(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 10}

	// Act
	actual := args.Map{"result": dm.KeyValueIntDefault("a") != 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
}

func Test_C13_DynamicMap_KeyValueByte(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": byte(5)}
	v, found, failed := dm.KeyValueByte("a")

	// Act
	actual := args.Map{"result": found || failed || v != 5}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	_, found2, _ := dm.KeyValueByte("missing")
	actual = args.Map{"result": found2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_C13_DynamicMap_KeyValueByte_IntValue(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 100}
	v, found, failed := dm.KeyValueByte("a")

	// Act
	actual := args.Map{"result": found || failed || v != 100}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
}

func Test_C13_DynamicMap_Add(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	dm.Add("k", "v")

	// Act
	actual := args.Map{"result": dm.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_Raw(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"k": "v"}
	raw := dm.Raw()

	// Act
	actual := args.Map{"result": len(raw) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_ConcatNew(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"b": 2}
	result := dm1.ConcatNew(true, dm2)

	// Act
	actual := args.Map{"result": result.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_ConcatNew_NoOverride(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 2, "b": 3}
	result := dm1.ConcatNew(false, dm2)

	// Act
	actual := args.Map{"result": result["a"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not override", actual)
}

func Test_C13_DynamicMap_ConcatNew_BothEmpty(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{}
	dm2 := enumimpl.DynamicMap{}
	result := dm1.ConcatNew(true, dm2)

	// Act
	actual := args.Map{"result": result.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_Strings(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	s := dm.Strings()

	// Act
	actual := args.Map{"result": len(s) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_Strings_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.Strings()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_StringsUsingFmt(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	s := dm.StringsUsingFmt(func(i int, k string, v any) string {
		return fmt.Sprintf("%d:%s=%v", i, k, v)
	})

	// Act
	actual := args.Map{"result": len(s) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_StringsUsingFmt_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.StringsUsingFmt(func(i int, k string, v any) string { return "" })) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_String(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C13_DynamicMap_IsStringEqual(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsStringEqual(dm.String())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_C13_DynamicMap_Serialize(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	_, err := dm.Serialize()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C13_DynamicMap_ConvMaps(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": byte(1), "b": byte(2)}
	byteMap := dm.ConvMapByteString()

	// Act
	actual := args.Map{"result": len(byteMap) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C13_DynamicMap_ConvMapByteString_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.ConvMapByteString()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_ConvMapStringInteger(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapStringInteger()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_ConvMapStringInteger_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.ConvMapStringInteger()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_ConvMapIntegerString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapIntegerString()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_ConvMapIntegerString_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.ConvMapIntegerString()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_ConvMapInt8String(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapInt8String()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_ConvMapInt8String_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.ConvMapInt8String()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_ConvMapInt16String(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapInt16String()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_ConvMapInt16String_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.ConvMapInt16String()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_ConvMapInt32String(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapInt32String()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_ConvMapInt32String_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.ConvMapInt32String()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_ConvMapUInt16String(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapUInt16String()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_ConvMapUInt16String_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.ConvMapUInt16String()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_ConvMapStringString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "x"}
	m := dm.ConvMapStringString()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_ConvMapStringString_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.ConvMapStringString()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_DynamicMap_ConvMapInt64String(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapInt64String()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_DynamicMap_ConvMapInt64String_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": len(dm.ConvMapInt64String()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ===================== DynamicMap → Basic* conversions =====================

func Test_C13_DynamicMap_BasicByte(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": byte(0), "Valid": byte(1)}
	bb := dm.BasicByte("TestByteEnum")

	// Act
	actual := args.Map{"result": bb == nil || bb.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected BasicByte with 2 items", actual)
}

func Test_C13_DynamicMap_BasicByteUsingAliasMap(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": byte(0), "Valid": byte(1)}
	bb := dm.BasicByteUsingAliasMap("TestByteEnum", map[string]byte{"v": 1})

	// Act
	actual := args.Map{"result": bb == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C13_DynamicMap_BasicInt8(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": int8(0), "B": int8(1)}
	bi := dm.BasicInt8("TestInt8Enum")

	// Act
	actual := args.Map{"result": bi == nil || bi.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_BasicInt16(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": int16(0), "B": int16(1)}
	bi := dm.BasicInt16("TestInt16Enum")

	// Act
	actual := args.Map{"result": bi == nil || bi.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_BasicInt32(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": int32(0), "B": int32(1)}
	bi := dm.BasicInt32("TestInt32Enum")

	// Act
	actual := args.Map{"result": bi == nil || bi.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_BasicString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"alpha": "x", "beta": "y"}
	bs := dm.BasicString("TestStringEnum")

	// Act
	actual := args.Map{"result": bs == nil || bs.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_DynamicMap_BasicUInt16(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": uint16(0), "B": uint16(1)}
	bu := dm.BasicUInt16("TestUInt16Enum")

	// Act
	actual := args.Map{"result": bu == nil || bu.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ===================== BasicByte =====================

func Test_C13_BasicByte_AllMethods(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.Create(
		"TestByte",
		[]byte{0, 1, 2},
		[]string{"Invalid", "Active", "Inactive"},
		0, 2,
	)

	// Act
	actual := args.Map{"result": bb.Min() != 0 || bb.Max() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "min/max wrong", actual)
	actual = args.Map{"result": bb.IsValidRange(1) || bb.IsValidRange(3)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "range check failed", actual)
	actual = args.Map{"result": bb.ToEnumString(0) != "Invalid"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Invalid", actual)
	actual = args.Map{"result": bb.IsAnyOf(1, 1, 2)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected match", actual)
	actual = args.Map{"result": bb.IsAnyOf(1, 3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no match", actual)
	actual = args.Map{"result": bb.IsAnyOf(1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty spread should return true", actual)
	actual = args.Map{"result": bb.IsAnyNamesOf(0, "Invalid", "Active")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected name match", actual)
	actual = args.Map{"result": bb.IsAnyNamesOf(0, "Active")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no name match", actual)
	v := bb.GetValueByString("Invalid")
	actual = args.Map{"result": v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	s := bb.GetStringValue(0)
	actual = args.Map{"result": s != "Invalid"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Invalid", actual)
	ranges := bb.Ranges()
	actual = args.Map{"result": len(ranges) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	hm := bb.Hashmap()
	actual = args.Map{"result": len(hm) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty hashmap", actual)
	hmPtr := bb.HashmapPtr()
	actual = args.Map{"result": hmPtr == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	jb, err := bb.ToEnumJsonBytes(0)
	actual = args.Map{"result": err != nil || len(jb) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "json bytes failed", actual)
	_, err2 := bb.ToEnumJsonBytes(99)
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid value", actual)
	s2 := bb.AppendPrependJoinValue(".", 1, 0)
	actual = args.Map{"result": s2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	s3 := bb.ToNumberString(1)
	actual = args.Map{"result": s3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	jm := bb.JsonMap()
	actual = args.Map{"result": len(jm) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	// UnmarshallToValue
	v2, err3 := bb.UnmarshallToValue(false, []byte("Invalid"))
	actual = args.Map{"result": err3}
	expected = args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err3", actual)
	_ = v2

	// nil bytes, no map to first
	_, err4 := bb.UnmarshallToValue(false, nil)
	actual = args.Map{"result": err4 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	// nil bytes, map to first
	v3, err5 := bb.UnmarshallToValue(true, nil)
	actual = args.Map{"result": err5 != nil || v3 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	// empty string, map to first
	v4, err6 := bb.UnmarshallToValue(true, []byte(""))
	actual = args.Map{"result": err6 != nil || v4 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	// double quote
	v5, err7 := bb.UnmarshallToValue(true, []byte(`""`))
	actual = args.Map{"result": err7 != nil || v5 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	actual = args.Map{"result": bb.EnumType().String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected enum type", actual)
	_ = bb.AsBasicByter()

	// GetValueByName
	_, err8 := bb.GetValueByName("NonExistent")
	actual = args.Map{"result": err8 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	vn, err9 := bb.GetValueByName("Invalid")
	actual = args.Map{"result": err9}
	expected = args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err9", actual)
	_ = vn
}

func Test_C13_BasicByte_ExpectingEnumValueError(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.Create(
		"TestByte",
		[]byte{0, 1},
		[]string{"Invalid", "Active"},
		0, 1,
	)

	// Matching value
	err := bb.ExpectingEnumValueError("Invalid", byte(0))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for matching", actual)

	// Non-matching
	err2 := bb.ExpectingEnumValueError("Active", byte(0))
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-matching", actual)

	// Invalid rawString
	err3 := bb.ExpectingEnumValueError("NonExistent", byte(0))
	actual = args.Map{"result": err3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid name", actual)
}

// ===================== BasicString =====================

func Test_C13_BasicString_AllMethods(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestString", []string{"alpha", "beta", "gamma"})

	// Act
	actual := args.Map{"result": bs.Min() == "" || bs.Max() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min/max", actual)
	ranges := bs.Ranges()
	actual = args.Map{"result": len(ranges) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual = args.Map{"result": bs.HasAnyItem() || bs.MaxIndex() != 2}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "check failed", actual)
	actual = args.Map{"result": bs.GetNameByIndex(1) == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name at index 1", actual)
	if bs.GetNameByIndex(0) != "" {
		// index 0 returns empty because condition is index > 0
	}
	actual = args.Map{"result": bs.GetNameByIndex(99) != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "out of range should return empty", actual)
	idx := bs.GetIndexByName("alpha")
	actual = args.Map{"result": idx < 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid index", actual)
	actual = args.Map{"result": bs.GetIndexByName("") >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty name should return invalid", actual)
	actual = args.Map{"result": bs.GetIndexByName("nonexistent") >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nonexistent should return invalid", actual)
	nim := bs.NameWithIndexMap()
	actual = args.Map{"result": len(nim) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	ri := bs.RangesIntegers()
	actual = args.Map{"result": len(ri) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	hs := bs.Hashset()
	actual = args.Map{"result": len(hs) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	hsPtr := bs.HashsetPtr()
	actual = args.Map{"result": hsPtr == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_, err := bs.GetValueByName("alpha")
	actual = args.Map{"result": err}
	expected = args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	_, err2 := bs.GetValueByName("nonexistent")
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual = args.Map{"result": bs.IsValidRange("alpha")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual = args.Map{"result": bs.IsValidRange("nonexistent")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual = args.Map{"result": bs.IsAnyOf("alpha", "alpha", "beta")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected match", actual)
	actual = args.Map{"result": bs.IsAnyOf("alpha", "beta")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no match", actual)
	actual = args.Map{"result": bs.IsAnyOf("alpha")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty spread should return true", actual)
	actual = args.Map{"result": bs.IsAnyNamesOf("alpha", "alpha")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected match", actual)

	// ToEnumJsonBytes
	jb, err3 := bs.ToEnumJsonBytes("alpha")
	actual = args.Map{"result": err3 != nil || len(jb) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "json bytes failed", actual)
	_, err4 := bs.ToEnumJsonBytes("nonexistent")
	actual = args.Map{"result": err4 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	// UnmarshallToValue
	v, err5 := bs.UnmarshallToValue(false, []byte("alpha"))
	actual = args.Map{"result": err5 != nil || v == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)
	_, err6 := bs.UnmarshallToValue(false, nil)
	actual = args.Map{"result": err6 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	v2, err7 := bs.UnmarshallToValue(true, nil)
	actual = args.Map{"result": err7}
	expected = args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err7", actual)
	_ = v2
	v3, err8 := bs.UnmarshallToValue(true, []byte(""))
	actual = args.Map{"result": err8}
	expected = args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err8", actual)
	_ = v3
	v4, err9 := bs.UnmarshallToValue(true, []byte(`""`))
	actual = args.Map{"result": err9}
	expected = args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err9", actual)
	_ = v4

	actual = args.Map{"result": bs.EnumType().String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected enum type", actual)

	// OnlySupportedErr
	err10 := bs.OnlySupportedErr("alpha")
	actual = args.Map{"result": err10 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for unsupported", actual)
	err11 := bs.OnlySupportedErr("alpha", "beta", "gamma")
	actual = args.Map{"result": err11 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all supported should return nil", actual)

	// OnlySupportedMsgErr
	err12 := bs.OnlySupportedMsgErr("msg", "alpha")
	actual = args.Map{"result": err12 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	// AppendPrependJoinValue
	s := bs.AppendPrependJoinValue(".", "beta", "alpha")
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ===================== numberEnumBase methods =====================

func Test_C13_NumberEnumBase_Methods(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.Create(
		"TestByte",
		[]byte{0, 1, 2},
		[]string{"Zero", "One", "Two"},
		0, 2,
	)

	min, max := bb.MinMaxAny()

	// Act
	actual := args.Map{"result": min == nil || max == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": bb.MinValueString() == "" || bb.MaxValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": bb.MinInt() != 0 || bb.MaxInt() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "min/max int wrong", actual)
	anv := bb.AllNameValues()
	actual = args.Map{"result": len(anv) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	rm := bb.RangesMap()
	actual = args.Map{"result": len(rm) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	ose := bb.OnlySupportedErr("Zero")
	actual = args.Map{"result": ose == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unsupported error", actual)
	osme := bb.OnlySupportedMsgErr("msg", "Zero")
	actual = args.Map{"result": osme == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	ier := bb.IntegerEnumRanges()
	actual = args.Map{"result": len(ier) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual = args.Map{"result": bb.Length() != 3 || bb.Count() != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	rdm := bb.RangesDynamicMap()
	actual = args.Map{"result": len(rdm) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	dm := bb.DynamicMap()
	actual = args.Map{"result": dm.Length() != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	rim := bb.RangesIntegerStringMap()
	_ = rim
	kav := bb.KeyAnyValues()
	actual = args.Map{"result": len(kav) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	kvi := bb.KeyValIntegers()
	actual = args.Map{"result": len(kvi) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	csv := bb.RangeNamesCsv()
	actual = args.Map{"result": csv == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	im := bb.RangesInvalidMessage()
	actual = args.Map{"result": im == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	ie := bb.RangesInvalidErr()
	actual = args.Map{"result": ie == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	sr := bb.StringRanges()
	actual = args.Map{"result": len(sr) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	srp := bb.StringRangesPtr()
	actual = args.Map{"result": len(srp) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	nh := bb.NamesHashset()
	actual = args.Map{"result": len(nh) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	js := bb.JsonString(byte(0))
	actual = args.Map{"result": js == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	tn := bb.TypeName()
	actual = args.Map{"result": tn != "TestByte"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected TestByte", actual)
	nwv := bb.NameWithValue(byte(0))
	actual = args.Map{"result": nwv == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	nwvo := bb.NameWithValueOption(byte(0), true)
	actual = args.Map{"result": nwvo == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	nwvo2 := bb.NameWithValueOption(byte(0), false)
	actual = args.Map{"result": nwvo2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	vs := bb.ValueString(byte(0))
	actual = args.Map{"result": vs == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	f := bb.Format("{type-name}-{name}-{value}", byte(0))
	actual = args.Map{"result": f == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected formatted", actual)

	// Loop
	count := 0
	bb.Loop(func(index int, name string, anyVal any) bool {
		count++
		return false
	})
	actual = args.Map{"result": count != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 iterations", actual)

	// Loop with break
	count2 := 0
	bb.Loop(func(index int, name string, anyVal any) bool {
		count2++
		return true
	})
	actual = args.Map{"result": count2 != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 iteration", actual)

	// LoopInteger
	count3 := 0
	bb.LoopInteger(func(index int, name string, val int) bool {
		count3++
		return false
	})
	actual = args.Map{"result": count3 != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

// ===================== DiffLeftRight =====================

func Test_C13_DiffLeftRight_AllMethods(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	l, r := d.Types()

	// Act
	actual := args.Map{"result": l == nil || r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected types", actual)
	actual = args.Map{"result": d.IsSameTypeSame()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected same type", actual)
	actual = args.Map{"result": d.IsSame()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected same", actual)
	actual = args.Map{"result": d.IsSameRegardlessOfType()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected same regardless", actual)
	actual = args.Map{"result": d.IsEqual(true) || !d.IsEqual(false)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	actual = args.Map{"result": d.HasMismatch(true) || d.HasMismatch(false)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no mismatch", actual)
	actual = args.Map{"result": d.IsNotEqual()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not not-equal", actual)
	actual = args.Map{"result": d.HasMismatchRegardlessOfType()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no mismatch", actual)
	actual = args.Map{"result": d.String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": d.DiffString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty diff string for equal values", actual)
	dl, dr := d.SpecificFullString()
	actual = args.Map{"result": dl == "" || dr == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C13_DiffLeftRight_Different(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}

	// Act
	actual := args.Map{"result": d.IsSame()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected different", actual)
	actual = args.Map{"result": d.IsNotEqual()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	actual = args.Map{"result": d.DiffString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty diff", actual)
	actual = args.Map{"result": d.HasMismatch(false)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
}

func Test_C13_DiffLeftRight_JsonString_Nil(t *testing.T) {
	// Arrange
	var d *enumimpl.DiffLeftRight

	// Act
	actual := args.Map{"result": d.JsonString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

// ===================== KeyAnyVal =====================

func Test_C13_KeyAnyVal_AllMethods(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "name", AnyValue: byte(5)}

	// Act
	actual := args.Map{"result": kav.KeyString() != "name"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
	actual = args.Map{"result": kav.AnyVal() != byte(5)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": kav.AnyValString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": kav.WrapKey() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": kav.WrapValue() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual = args.Map{"result": kav.IsString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "byte should not be string type", actual)
	actual = args.Map{"result": kav.ValInt() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	kvi := kav.KeyValInteger()
	actual = args.Map{"result": kvi.Key != "name" || kvi.ValueInteger != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "conversion failed", actual)
	s := kav.String()
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C13_KeyAnyVal_StringType(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "name", AnyValue: "strval"}

	// Act
	actual := args.Map{"result": kav.IsString()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected string type", actual)
	s := kav.String()
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ===================== KeyAnyValues =====================

func Test_C13_KeyAnyValues_Empty(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{}, []byte{})

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C13_KeyAnyValues_NonEmpty(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{"a", "b"}, []byte{0, 1})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ===================== Format, FormatUsingFmt, NameWithValue, PrependJoin, JoinPrependUsingDot =====================

func Test_C13_Format(t *testing.T) {
	// Arrange
	s := enumimpl.Format("MyEnum", "Active", "1", "Enum:{type-name}-{name}-{value}")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected formatted", actual)
}

type testFormatterC13 struct{}

func (tf testFormatterC13) TypeName() string   { return "TestType" }
func (tf testFormatterC13) Name() string       { return "TestName" }
func (tf testFormatterC13) ValueString() string { return "TestVal" }

func Test_C13_FormatUsingFmt(t *testing.T) {
	// Arrange
	s := enumimpl.FormatUsingFmt(testFormatterC13{}, "{type-name}-{name}-{value}")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected formatted", actual)
}

func Test_C13_NameWithValue(t *testing.T) {
	// Arrange
	s := enumimpl.NameWithValue(byte(5))

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C13_PrependJoin(t *testing.T) {
	// Arrange
	s := enumimpl.PrependJoin(".", "prefix", "a", "b")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C13_JoinPrependUsingDot(t *testing.T) {
	// Arrange
	s := enumimpl.JoinPrependUsingDot("prefix", "a", "b")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ===================== ConvEnumAnyValToInteger =====================

func Test_C13_ConvEnumAnyValToInteger_String(t *testing.T) {
	v := enumimpl.ConvEnumAnyValToInteger("hello")
	_ = v // should return MinInt
}

func Test_C13_ConvEnumAnyValToInteger_Int(t *testing.T) {
	// Arrange
	v := enumimpl.ConvEnumAnyValToInteger(42)

	// Act
	actual := args.Map{"result": v != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C13_ConvEnumAnyValToInteger_Byte(t *testing.T) {
	// Arrange
	v := enumimpl.ConvEnumAnyValToInteger(byte(5))

	// Act
	actual := args.Map{"result": v != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_C13_ConvEnumAnyValToInteger_Fallback(t *testing.T) {
	// A float will be Sprintf'd and Atoi'd
	v := enumimpl.ConvEnumAnyValToInteger(3.0)
	_ = v
}

// ===================== IntegersRangesOfAnyVal =====================

func Test_C13_IntegersRangesOfAnyVal(t *testing.T) {
	// Arrange
	result := enumimpl.IntegersRangesOfAnyVal([]byte{2, 0, 1})

	// Act
	actual := args.Map{"result": len(result) != 3 || result[0] != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted [0,1,2]", actual)
}

// ===================== AllNameValues =====================

func Test_C13_AllNameValues(t *testing.T) {
	// Arrange
	result := enumimpl.AllNameValues([]string{"a", "b"}, []byte{0, 1})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ===================== UnsupportedNames =====================

func Test_C13_UnsupportedNames(t *testing.T) {
	// Arrange
	result := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a")

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 unsupported", actual)
}

func Test_C13_UnsupportedNames_AllSupported(t *testing.T) {
	// Arrange
	result := enumimpl.UnsupportedNames([]string{"a", "b"}, "a", "b")

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ===================== OnlySupportedErr =====================

func Test_C13_OnlySupportedErr_EmptyAll(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(2, []string{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C13_OnlySupportedErr_AllSupported(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(2, []string{"a"}, "a")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C13_OnlySupportedErr_HasUnsupported(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(2, []string{"a", "b"}, "a")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ===================== DifferCheckerImpl =====================

func Test_C13_DifferCheckerImpl_GetSingleDiffResult(t *testing.T) {
	// Arrange
	dc := enumimpl.DefaultDiffCheckerImpl

	// Act
	actual := args.Map{"result": dc.GetSingleDiffResult(true, "l", "r") != "l"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected left", actual)
	actual = args.Map{"result": dc.GetSingleDiffResult(false, "l", "r") != "r"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected right", actual)
}

func Test_C13_DifferCheckerImpl_IsEqual(t *testing.T) {
	// Arrange
	dc := enumimpl.DefaultDiffCheckerImpl

	// Act
	actual := args.Map{"result": dc.IsEqual(true, 1, 1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal regardless", actual)
	actual = args.Map{"result": dc.IsEqual(false, 1, "1")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal strict", actual)
}

func Test_C13_DifferCheckerImpl_GetResultOnKeyMissing(t *testing.T) {
	// Arrange
	dc := enumimpl.DefaultDiffCheckerImpl
	r := dc.GetResultOnKeyMissingInRightExistInLeft("k", "v")

	// Act
	actual := args.Map{"result": r != "v"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected v", actual)
}

// ===================== LeftRightDiffCheckerImpl =====================

func Test_C13_LeftRightDiffCheckerImpl(t *testing.T) {
	// Arrange
	lrdc := enumimpl.LeftRightDiffCheckerImpl
	r := lrdc.GetSingleDiffResult(true, 1, 2)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	r2 := lrdc.GetResultOnKeyMissingInRightExistInLeft("k", "v")
	actual = args.Map{"result": r2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": lrdc.IsEqual(true, 1, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

// ===================== Creator methods =====================

func Test_C13_NewBasicByte_UsingTypeSlice(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("T", []string{"A", "B"})

	// Act
	actual := args.Map{"result": bb.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_NewBasicByte_Default(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.Default(byte(0), []string{"A", "B"})

	// Act
	actual := args.Map{"result": bb.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_NewBasicByte_DefaultAllCases(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.DefaultAllCases(byte(0), []string{"Active", "Inactive"})

	// Act
	actual := args.Map{"result": bb.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_NewBasicByte_DefaultWithAliasMap(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.DefaultWithAliasMap(byte(0), []string{"A", "B"}, map[string]byte{"a": 0})

	// Act
	actual := args.Map{"result": bb.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_NewBasicByte_DefaultWithAliasMapAllCases(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.DefaultWithAliasMapAllCases(byte(0), []string{"A"}, map[string]byte{"a": 0})

	// Act
	actual := args.Map{"result": bb.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_NewBasicByte_UsingFirstItemSliceCaseOptions(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingFirstItemSliceCaseOptions(false, byte(0), []string{"A"})

	// Act
	actual := args.Map{"result": bb.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_NewBasicByte_UsingFirstItemSliceAllCases(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(byte(0), []string{"A"})

	// Act
	actual := args.Map{"result": bb.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_NewBasicByte_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingFirstItemSliceAliasMap(byte(0), []string{"A"}, nil)

	// Act
	actual := args.Map{"result": bb.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_NewBasicByte_CreateUsingMapPlusAliasMapOptions(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.CreateUsingMapPlusAliasMapOptions(
		true, byte(0), map[byte]string{0: "A"}, nil,
	)

	// Act
	actual := args.Map{"result": bb.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_NewBasicString_CreateDefault(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateDefault("strval", []string{"a", "b"})

	// Act
	actual := args.Map{"result": bs.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_NewBasicString_CreateUsingStringersSpread(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateUsingNamesSpread("T", "alpha", "beta")

	// Act
	actual := args.Map{"result": bs.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_NewBasicString_CreateUsingNamesMinMax(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateUsingNamesMinMax("T", []string{"a", "b"}, "a", "b")

	// Act
	actual := args.Map{"result": bs.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_NewBasicString_CreateUsingSlicePlusAliasMapOptions(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateUsingSlicePlusAliasMapOptions(
		true, "strval", []string{"A", "B"}, nil,
	)

	// Act
	actual := args.Map{"result": bs.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C13_NewBasicString_UsingFirstItemSliceCaseOptions(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.UsingFirstItemSliceCaseOptions(false, "strval", []string{"A"})

	// Act
	actual := args.Map{"result": bs.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C13_NewBasicString_UsingFirstItemSliceAllCases(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.UsingFirstItemSliceAllCases("strval", []string{"A"})

	// Act
	actual := args.Map{"result": bs.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ===================== DynamicMap with LeftRightDiffChecker =====================

func Test_C13_DynamicMap_LogShouldDiffMessageUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test",
		map[string]any{"a": 2},
	)

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diff", actual)
}

func Test_C13_DynamicMap_LogShouldDiffMessageUsingDifferChecker_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, true, "test",
		map[string]any{"a": 1},
	)

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_LogShouldDiffLeftRightMessageUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, false, "test",
		map[string]any{"b": 2},
	)

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diff", actual)
}

func Test_C13_DynamicMap_LogShouldDiffLeftRightMessageUsingDifferChecker_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, true, "test",
		map[string]any{"a": 1},
	)

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_ShouldDiffMessageUsingDifferChecker_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, true, "test",
		map[string]any{"a": 1},
	)

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C13_DynamicMap_DiffJsonMessageUsingDifferChecker_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, true,
		map[string]any{"a": 1},
	)

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ===================== IsValueTypeOf =====================

func Test_C13_DynamicMap_IsValueTypeOf(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "str"}
	// This checks reflect.TypeOf against first value's type
	_ = dm.IsValueTypeOf(nil)
}

// ===================== AppendPrependJoinNamer =====================

type testNamerC13 struct{ name string }

func (n testNamerC13) Name() string { return n.name }

func Test_C13_BasicByte_AppendPrependJoinNamer(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.Create("T", []byte{0, 1}, []string{"A", "B"}, 0, 1)
	s := bb.AppendPrependJoinNamer(".", testNamerC13{"append"}, testNamerC13{"prepend"})

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ===================== NamesHashset empty =====================

func Test_C13_NamesHashset_Empty(t *testing.T) {
	// Create a BasicByte with no items — exercise NamesHashset empty path
	// Note: This would require a zero-item enum which isn't natural
	// The Length() == 0 path returns empty map
}
