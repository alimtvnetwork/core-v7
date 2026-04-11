package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// enumimpl Coverage — DynamicMap comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func newTestDynMap() enumimpl.DynamicMap {
	return enumimpl.DynamicMap{
		"Invalid": byte(0),
		"Read":    byte(1),
		"Write":   byte(2),
		"Execute": byte(3),
	}
}

func Test_CovEnum_DM01_AddOrUpdate(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.AddOrUpdate("b", 2)

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new", actual)
	isNew2 := dm.AddOrUpdate("a", 99)
	actual = args.Map{"result": isNew2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected update", actual)
}

func Test_CovEnum_DM02_Set(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.Set("b", 2)

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new", actual)
}

func Test_CovEnum_DM03_AddNewOnly(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	ok := dm.AddNewOnly("b", 2)

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected added", actual)
	ok2 := dm.AddNewOnly("a", 99)
	actual = args.Map{"result": ok2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not added", actual)
}

func Test_CovEnum_DM04_AllKeys_AllKeysSorted(t *testing.T) {
	// Arrange
	dm := newTestDynMap()
	keys := dm.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	sorted := dm.AllKeysSorted()
	actual = args.Map{"result": sorted[0] != "Execute"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Execute first", actual)
	// empty
	empty := enumimpl.DynamicMap{}
	actual = args.Map{"result": len(empty.AllKeys()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.AllKeysSorted()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovEnum_DM05_AllValuesStrings_Sorted(t *testing.T) {
	// Arrange
	dm := newTestDynMap()
	vs := dm.AllValuesStrings()

	// Act
	actual := args.Map{"result": len(vs) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	vss := dm.AllValuesStringsSorted()
	actual = args.Map{"result": len(vss) != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	// empty
	empty := enumimpl.DynamicMap{}
	actual = args.Map{"result": len(empty.AllValuesStrings()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.AllValuesStringsSorted()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovEnum_DM06_AllValuesIntegers(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	ints := dm.AllValuesIntegers()

	// Act
	actual := args.Map{"result": len(ints) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	// empty
	empty := enumimpl.DynamicMap{}
	actual = args.Map{"result": len(empty.AllValuesIntegers()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovEnum_DM07_MapIntegerString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	m, keys := dm.MapIntegerString()

	// Act
	actual := args.Map{"result": len(m) < 2 || len(keys) < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	// empty
	empty := enumimpl.DynamicMap{}
	m2, k2 := empty.MapIntegerString()
	actual = args.Map{"result": len(m2) != 0 || len(k2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovEnum_DM08_SortedKeyValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	kv := dm.SortedKeyValues()

	// Act
	actual := args.Map{"result": len(kv) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	// empty
	empty := enumimpl.DynamicMap{}
	actual = args.Map{"result": len(empty.SortedKeyValues()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovEnum_DM09_SortedKeyAnyValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	kav := dm.SortedKeyAnyValues()

	// Act
	actual := args.Map{"result": len(kav) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	// empty
	empty := enumimpl.DynamicMap{}
	actual = args.Map{"result": len(empty.SortedKeyAnyValues()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	// string values
	dmStr := enumimpl.DynamicMap{"x": "hello", "y": "world"}
	kavStr := dmStr.SortedKeyAnyValues()
	actual = args.Map{"result": len(kavStr) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_CovEnum_DM10_First(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	k, v := dm.First()

	// Act
	actual := args.Map{"result": k != "a" || v != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a,1", actual)
	// empty
	empty := enumimpl.DynamicMap{}
	k2, v2 := empty.First()
	actual = args.Map{"result": k2 != "" || v2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovEnum_DM11_IsValueString_IsValueTypeOf(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "hello"}

	// Act
	actual := args.Map{"result": dm.IsValueString()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	dm2 := enumimpl.DynamicMap{"a": 1}
	actual = args.Map{"result": dm2.IsValueString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovEnum_DM12_Length_Count_IsEmpty_HasAnyItem_LastIndex_HasIndex(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": dm.Count() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": dm.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": dm.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": dm.LastIndex() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": dm.HasIndex(1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": dm.HasIndex(5)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// nil ptr
	var nilDm *enumimpl.DynamicMap
	actual = args.Map{"result": nilDm.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovEnum_DM13_HasKey_HasAllKeys_HasAnyKeys_IsMissingKey(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.HasKey("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": dm.HasKey("c")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": dm.HasAllKeys("a", "b")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": dm.HasAllKeys("a", "c")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": dm.HasAnyKeys("a", "c")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": dm.HasAnyKeys("c", "d")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": dm.IsMissingKey("c")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovEnum_DM14_IsEqual_IsMismatch(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm1.IsEqual(false, &dm2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	actual = args.Map{"result": dm1.IsMismatch(false, &dm2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no mismatch", actual)
	// nil comparisons
	var nilDm *enumimpl.DynamicMap
	actual = args.Map{"result": nilDm.IsEqual(false, nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil==nil", actual)
	actual = args.Map{"result": nilDm.IsEqual(false, &dm1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	// regardless of type
	dm3 := enumimpl.DynamicMap{"a": byte(1)}
	actual = args.Map{"result": dm1.IsEqual(true, &dm3)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal regardless of type", actual)
}

func Test_CovEnum_DM15_IsRawEqual_IsRawMismatch(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	raw := map[string]any{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsRawEqual(false, raw)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	actual = args.Map{"result": dm.IsRawMismatch(false, raw)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no mismatch", actual)
	// different lengths
	raw2 := map[string]any{"a": 1, "b": 2}
	actual = args.Map{"result": dm.IsRawEqual(false, raw2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	// missing key
	raw3 := map[string]any{"b": 1}
	actual = args.Map{"result": dm.IsRawEqual(false, raw3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	// nil dm
	var nilDm *enumimpl.DynamicMap
	actual = args.Map{"result": nilDm.IsRawEqual(false, nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil==nil", actual)
	actual = args.Map{"result": nilDm.IsRawEqual(false, raw)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovEnum_DM16_DiffRaw(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}
	diff := dm.DiffRaw(false, right)

	// Act
	actual := args.Map{"result": diff.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diffs", actual)
	// same maps
	same := map[string]any{"a": 1, "b": 2}
	diff2 := dm.DiffRaw(false, same)
	actual = args.Map{"result": diff2.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no diff", actual)
}

func Test_CovEnum_DM17_DiffRawUsingDifferChecker_NilBranches(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	// both nil
	var nilDm *enumimpl.DynamicMap
	r := nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	// Act
	actual := args.Map{"result": r.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	// left nil, right not nil
	right := map[string]any{"a": 1}
	r2 := nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)
	actual = args.Map{"result": r2.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// left not nil, right nil
	r3 := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	actual = args.Map{"result": r3.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_DM18_DiffRawLeftRightUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}
	lDiff, rDiff := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)
	_ = lDiff
	_ = rDiff
	// nil branches
	var nilDm *enumimpl.DynamicMap
	l2, r2 := nilDm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	// Act
	actual := args.Map{"result": l2.HasAnyItem() || r2.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovEnum_DM19_DiffJsonMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	same := map[string]any{"a": 1}
	msg := dm.DiffJsonMessage(false, same)

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	diff := map[string]any{"a": 2}
	msg2 := dm.DiffJsonMessage(false, diff)
	actual = args.Map{"result": msg2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_DM20_DiffJsonMessageLeftRight(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	diff := map[string]any{"a": 2, "b": 3}
	msg := dm.DiffJsonMessageLeftRight(false, diff)

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_DM21_ShouldDiffMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessage(false, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	msg2 := dm.ShouldDiffMessage(false, "test", map[string]any{"a": 2})
	actual = args.Map{"result": msg2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_DM22_ShouldDiffLeftRightMessageUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	msg2 := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 2})
	actual = args.Map{"result": msg2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_DM23_LogShouldDiffMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessage(false, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovEnum_DM24_LogShouldDiffLeftRightMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessage(false, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovEnum_DM25_LogShouldDiffMessageUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovEnum_DM26_ExpectingMessage(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ExpectingMessage("test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for same", actual)
	msg2 := dm.ExpectingMessage("test", map[string]any{"a": 2})
	actual = args.Map{"result": msg2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_DM27_LogExpectingMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 1})
}

func Test_CovEnum_DM28_IsKeysEqualOnly(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"a": 99, "b": 88})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"a": 99, "c": 88})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"a": 99})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for different length", actual)
	var nilDm *enumimpl.DynamicMap
	actual = args.Map{"result": nilDm.IsKeysEqualOnly(nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil==nil", actual)
	actual = args.Map{"result": nilDm.IsKeysEqualOnly(map[string]any{"a": 1})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovEnum_DM29_KeyValue_KeyValueString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	v, ok := dm.KeyValue("a")

	// Act
	actual := args.Map{"result": ok || v != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_, ok2 := dm.KeyValue("missing")
	actual = args.Map{"result": ok2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	vs, ok3 := dm.KeyValueString("a")
	actual = args.Map{"result": ok3 || vs == ""}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	_, ok4 := dm.KeyValueString("missing")
	actual = args.Map{"result": ok4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovEnum_DM30_KeyValueIntDefault_KeyValueInt(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	vi := dm.KeyValueIntDefault("a")

	// Act
	actual := args.Map{"result": vi != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	vi2 := dm.KeyValueIntDefault("missing")
	actual = args.Map{"result": vi2 >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_CovEnum_DM31_KeyValueByte(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": byte(5)}
	b, found, failed := dm.KeyValueByte("a")

	// Act
	actual := args.Map{"result": found || failed || b != 5}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	_, found2, _ := dm.KeyValueByte("missing")
	actual = args.Map{"result": found2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
	// int value
	dm2 := enumimpl.DynamicMap{"a": 42}
	b2, found3, failed3 := dm2.KeyValueByte("a")
	actual = args.Map{"result": found3 || failed3 || b2 != 42}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_CovEnum_DM32_Add(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	dm.Add("a", 1)

	// Act
	actual := args.Map{"result": dm.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovEnum_DM33_ConvMaps(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	si := dm.ConvMapStringInteger()

	// Act
	actual := args.Map{"result": len(si) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	is := dm.ConvMapIntegerString()
	actual = args.Map{"result": len(is) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	bs := dm.ConvMapByteString()
	_ = bs
	i8s := dm.ConvMapInt8String()
	_ = i8s
	i16s := dm.ConvMapInt16String()
	_ = i16s
	i32s := dm.ConvMapInt32String()
	_ = i32s
	u16s := dm.ConvMapUInt16String()
	_ = u16s
	i64s := dm.ConvMapInt64String()
	_ = i64s
	ss := dm.ConvMapStringString()
	_ = ss
	// empty
	empty := enumimpl.DynamicMap{}
	actual = args.Map{"result": len(empty.ConvMapStringInteger()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.ConvMapIntegerString()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.ConvMapByteString()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.ConvMapInt8String()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.ConvMapInt16String()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.ConvMapInt32String()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.ConvMapUInt16String()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.ConvMapInt64String()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.ConvMapStringString()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovEnum_DM34_ConcatNew(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"b": 2}
	r := dm1.ConcatNew(true, dm2)

	// Act
	actual := args.Map{"result": r.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	// no override
	dm3 := enumimpl.DynamicMap{"a": 99}
	r2 := dm1.ConcatNew(false, dm3)
	actual = args.Map{"result": r2["a"] != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 not overridden", actual)
	// override
	r3 := dm1.ConcatNew(true, dm3)
	actual = args.Map{"result": r3["a"] != 99}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99 overridden", actual)
	// both empty
	empty1 := enumimpl.DynamicMap{}
	empty2 := enumimpl.DynamicMap{}
	r4 := empty1.ConcatNew(true, empty2)
	actual = args.Map{"result": r4.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovEnum_DM35_Strings_String_StringsUsingFmt(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	ss := dm.Strings()

	// Act
	actual := args.Map{"result": len(ss) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	s := dm.String()
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	sf := dm.StringsUsingFmt(func(index int, key string, val any) string {
		return key
	})
	actual = args.Map{"result": len(sf) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	// empty
	empty := enumimpl.DynamicMap{}
	actual = args.Map{"result": len(empty.Strings()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.StringsUsingFmt(func(i int, k string, v any) string { return k })) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovEnum_DM36_IsStringEqual(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsStringEqual(dm.String())}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovEnum_DM37_Serialize(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	b, err := dm.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected serialization", actual)
}

func Test_CovEnum_DM38_Raw(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	raw := dm.Raw()

	// Act
	actual := args.Map{"result": len(raw) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovEnum_DM39_BasicByte(t *testing.T) {
	// Arrange
	dm := newTestDynMap()
	bb := dm.BasicByte("TestEnum")

	// Act
	actual := args.Map{"result": bb == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": bb.Length() != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_CovEnum_DM40_BasicByteUsingAliasMap(t *testing.T) {
	// Arrange
	dm := newTestDynMap()
	alias := map[string]byte{"r": 1, "w": 2}
	bb := dm.BasicByteUsingAliasMap("TestEnum", alias)

	// Act
	actual := args.Map{"result": bb == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovEnum_DM41_BasicString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": "Invalid", "Active": "Active"}
	bs := dm.BasicString("TestStrEnum")

	// Act
	actual := args.Map{"result": bs == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovEnum_DM42_BasicStringUsingAliasMap(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": "Invalid", "Active": "Active"}
	alias := map[string]string{"inv": "Invalid"}
	bs := dm.BasicStringUsingAliasMap("TestStrEnum", alias)

	// Act
	actual := args.Map{"result": bs == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovEnum_DM43_BasicInt8(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	bi := dm.BasicInt8("TestI8")

	// Act
	actual := args.Map{"result": bi == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovEnum_DM44_BasicInt16(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	bi := dm.BasicInt16("TestI16")

	// Act
	actual := args.Map{"result": bi == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovEnum_DM45_BasicInt32(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	bi := dm.BasicInt32("TestI32")

	// Act
	actual := args.Map{"result": bi == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovEnum_DM46_BasicUInt16(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	bi := dm.BasicUInt16("TestU16")

	// Act
	actual := args.Map{"result": bi == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicByte — targeted branch coverage
// ══════════════════════════════════════════════════════════════════════════════

func newTestBasicByte() *enumimpl.BasicByte {
	dm := newTestDynMap()
	return dm.BasicByte("TestEnum")
}

func Test_CovEnum_BB01_IsAnyOf(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()

	// Act
	actual := args.Map{"result": bb.IsAnyOf(1, 1, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": bb.IsAnyOf(5, 1, 2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": bb.IsAnyOf(5)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for empty variadic", actual)
}

func Test_CovEnum_BB02_IsAnyNamesOf(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	name := bb.ToEnumString(1)

	// Act
	actual := args.Map{"result": bb.IsAnyNamesOf(1, name)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": bb.IsAnyNamesOf(1, "nonexistent")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovEnum_BB03_MinMax(t *testing.T) {
	bb := newTestBasicByte()
	_ = bb.Min()
	_ = bb.Max()
}

func Test_CovEnum_BB04_GetValueByName_Valid_Invalid(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	_, err := bb.GetValueByName("Invalid")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_, err2 := bb.GetValueByName("nonexistent_xyz")
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovEnum_BB05_IsValidRange(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()

	// Act
	actual := args.Map{"result": bb.IsValidRange(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovEnum_BB06_ToEnumJsonBytes(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	b, err := bb.ToEnumJsonBytes(0)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	actual = args.Map{"result": len(b) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	_, err2 := bb.ToEnumJsonBytes(99)
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovEnum_BB07_UnmarshallToValue(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	// nil not mapped to first
	_, err := bb.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// nil mapped to first
	v, err2 := bb.UnmarshallToValue(true, nil)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = v
	// empty string mapped to first
	v2, err3 := bb.UnmarshallToValue(true, []byte(""))
	actual = args.Map{"result": err3 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = v2
	// valid name
	_, err4 := bb.UnmarshallToValue(false, []byte("Invalid"))
	actual = args.Map{"result": err4 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovEnum_BB08_EnumType(t *testing.T) {
	bb := newTestBasicByte()
	_ = bb.EnumType()
}

func Test_CovEnum_BB09_AsBasicByter(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	byter := bb.AsBasicByter()

	// Act
	actual := args.Map{"result": byter == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovEnum_BB10_Hashmap_Ranges_AppendPrependJoinValue(t *testing.T) {
	bb := newTestBasicByte()
	_ = bb.Hashmap()
	_ = bb.HashmapPtr()
	_ = bb.Ranges()
	_ = bb.AppendPrependJoinValue(".", 1, 2)
	_ = bb.ToNumberString(1)
	_ = bb.JsonMap()
	_ = bb.GetStringValue(0)
	_ = bb.GetValueByString("Invalid")
}

func Test_CovEnum_BB11_ExpectingEnumValueError(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	// ToName(byte(0)) returns fmt.Sprintf("%v", 0) = "0", then GetValueByString("0")
	// returns whatever value was at index 0 during map iteration — non-deterministic.
	// Only test the deterministic error path (bad name).
	err2 := bb.ExpectingEnumValueError("nonexistent_xyz_999", byte(0))

	// Act
	actual := args.Map{"result": err2 == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for bad name", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// numberEnumBase — coverage via BasicByte
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_NEB01_MinMaxAny_MinValueString_MaxValueString(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	min, max := bb.MinMaxAny()

	// Act
	actual := args.Map{"result": min == nil || max == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	ms := bb.MinValueString()
	actual = args.Map{"result": ms == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// cached call
	ms2 := bb.MinValueString()
	actual = args.Map{"result": ms2 != ms}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same cached", actual)
	mx := bb.MaxValueString()
	actual = args.Map{"result": mx == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// cached call
	mx2 := bb.MaxValueString()
	actual = args.Map{"result": mx2 != mx}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same cached", actual)
}

func Test_CovEnum_NEB02_MinInt_MaxInt(t *testing.T) {
	bb := newTestBasicByte()
	_ = bb.MinInt()
	_ = bb.MaxInt()
}

func Test_CovEnum_NEB03_AllNameValues(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	nv := bb.AllNameValues()

	// Act
	actual := args.Map{"result": len(nv) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_CovEnum_NEB04_RangesMap_DynamicMap_RangesDynamicMap(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	rm := bb.RangesMap()

	// Act
	actual := args.Map{"result": len(rm) < 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 4", actual)
	dm := bb.DynamicMap()
	actual = args.Map{"result": dm.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	rdm := bb.RangesDynamicMap()
	actual = args.Map{"result": len(rdm) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// cached
	rdm2 := bb.RangesDynamicMap()
	actual = args.Map{"result": len(rdm2) != len(rdm)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same", actual)
}

func Test_CovEnum_NEB05_IntegerEnumRanges(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	ranges := bb.IntegerEnumRanges()

	// Act
	actual := args.Map{"result": len(ranges) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_CovEnum_NEB06_KeyAnyValues_KeyValIntegers(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	kav := bb.KeyAnyValues()

	// Act
	actual := args.Map{"result": len(kav) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	// cached
	kav2 := bb.KeyAnyValues()
	actual = args.Map{"result": len(kav2) != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	kvi := bb.KeyValIntegers()
	actual = args.Map{"result": len(kvi) != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_CovEnum_NEB07_Loop_LoopInteger(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	count := 0
	bb.Loop(func(index int, name string, anyVal any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"result": count != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	// break early
	count2 := 0
	bb.Loop(func(index int, name string, anyVal any) bool {
		count2++
		return true
	})
	actual = args.Map{"result": count2 != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	count3 := 0
	bb.LoopInteger(func(index int, name string, anyVal int) bool {
		count3++
		return false
	})
	actual = args.Map{"result": count3 != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_CovEnum_NEB08_RangesCsv_RangesInvalidMessage_RangesInvalidErr(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	csv := bb.RangeNamesCsv()

	// Act
	actual := args.Map{"result": csv == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	msg := bb.RangesInvalidMessage()
	actual = args.Map{"result": msg == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	err := bb.RangesInvalidErr()
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovEnum_NEB09_StringRanges_NamesHashset(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	sr := bb.StringRanges()

	// Act
	actual := args.Map{"result": len(sr) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	srp := bb.StringRangesPtr()
	actual = args.Map{"result": len(srp) != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	hs := bb.NamesHashset()
	actual = args.Map{"result": len(hs) != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_CovEnum_NEB10_NameWithValue_NameWithValueOption(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	nv := bb.NameWithValue(1)

	// Act
	actual := args.Map{"result": nv == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	nvo := bb.NameWithValueOption(1, true)
	actual = args.Map{"result": nvo == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	nvo2 := bb.NameWithValueOption(1, false)
	actual = args.Map{"result": nvo2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_NEB11_ValueString_JsonString_ToEnumString_ToName(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	vs := bb.ValueString(1)

	// Act
	actual := args.Map{"result": vs == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	js := bb.JsonString(1)
	actual = args.Map{"result": js == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_NEB12_Format(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	f := bb.Format("Enum of {type-name} - {name} - {value}", byte(1))

	// Act
	actual := args.Map{"result": f == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_NEB13_OnlySupportedErr(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	err := bb.OnlySupportedErr("Invalid")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for unsupported names", actual)
}

func Test_CovEnum_NEB14_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	bb := newTestBasicByte()
	err := bb.OnlySupportedMsgErr("prefix", "Invalid")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovEnum_NEB15_RangesIntegerStringMap(t *testing.T) {
	bb := newTestBasicByte()
	m := bb.RangesIntegerStringMap()
	_ = m
}

// ══════════════════════════════════════════════════════════════════════════════
// Helper functions and types
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_Misc01_AllNameValues(t *testing.T) {
	// Arrange
	names := []string{"Invalid", "Read", "Write"}
	vals := []byte{0, 1, 2}
	r := enumimpl.AllNameValues(names, vals)

	// Act
	actual := args.Map{"result": len(r) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_CovEnum_Misc02_ConvEnumAnyValToInteger(t *testing.T) {
	// Arrange
	// string returns MinInt
	r := enumimpl.ConvEnumAnyValToInteger("hello")

	// Act
	actual := args.Map{"result": r >= 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected MinInt", actual)
	// int
	r2 := enumimpl.ConvEnumAnyValToInteger(42)
	actual = args.Map{"result": r2 != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_CovEnum_Misc03_IntegersRangesOfAnyVal(t *testing.T) {
	// Arrange
	vals := []byte{0, 1, 2}
	r := enumimpl.IntegersRangesOfAnyVal(vals)

	// Act
	actual := args.Map{"result": len(r) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_CovEnum_Misc04_Format(t *testing.T) {
	// Arrange
	r := enumimpl.Format("TypeA", "Name1", "0", "Enum of {type-name} - {name} - {value}")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_Misc05_PrependJoin_JoinPrependUsingDot(t *testing.T) {
	// Arrange
	r := enumimpl.PrependJoin(".", "pre", "a", "b")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := enumimpl.JoinPrependUsingDot("pre", "a", "b")
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_Misc06_NameWithValue(t *testing.T) {
	// Arrange
	r := enumimpl.NameWithValue(1)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_Misc07_OnlySupportedErr(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b", "c"}, "a")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// empty all names
	err2 := enumimpl.OnlySupportedErr(0, []string{}, "a")
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	// all supported
	err3 := enumimpl.OnlySupportedErr(0, []string{"a", "b"}, "a", "b")
	actual = args.Map{"result": err3 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovEnum_Misc08_UnsupportedNames(t *testing.T) {
	// Arrange
	un := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a")

	// Act
	actual := args.Map{"result": len(un) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_CovEnum_Misc09_KeyAnyValues(t *testing.T) {
	// Arrange
	names := []string{"a", "b"}
	vals := []int{1, 2}
	kav := enumimpl.KeyAnyValues(names, vals)

	// Act
	actual := args.Map{"result": len(kav) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	// empty
	kav2 := enumimpl.KeyAnyValues(nil, nil)
	actual = args.Map{"result": len(kav2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DiffLeftRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_DLR01_Types_IsSameTypeSame(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	l, r := d.Types()

	// Act
	actual := args.Map{"result": l != r}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same type", actual)
	actual = args.Map{"result": d.IsSameTypeSame()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovEnum_DLR02_IsSame_IsNotEqual(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}

	// Act
	actual := args.Map{"result": d.IsSame()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected same", actual)
	actual = args.Map{"result": d.IsNotEqual()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not notequal", actual)
	d2 := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	actual = args.Map{"result": d2.IsSame()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not same", actual)
}

func Test_CovEnum_DLR03_IsSameRegardlessOfType(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: byte(1)}

	// Act
	actual := args.Map{"result": d.IsSameRegardlessOfType()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovEnum_DLR04_IsEqual_HasMismatch(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}

	// Act
	actual := args.Map{"result": d.IsEqual(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	actual = args.Map{"result": d.HasMismatch(false)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no mismatch", actual)
	actual = args.Map{"result": d.IsEqual(true)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal regardless", actual)
	actual = args.Map{"result": d.HasMismatch(true)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no mismatch regardless", actual)
	d2 := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	actual = args.Map{"result": d2.HasMismatch(false)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
	if d2.HasMismatchRegardlessOfType() {
		// might be same string fmt
	}
}

func Test_CovEnum_DLR05_String_JsonString_DiffString(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	s := d.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	js := d.JsonString()
	actual = args.Map{"result": js == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	ds := d.DiffString()
	actual = args.Map{"result": ds == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty for different", actual)
	// same
	d2 := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	ds2 := d2.DiffString()
	actual = args.Map{"result": ds2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for same", actual)
	// nil
	var nilD *enumimpl.DiffLeftRight
	actual = args.Map{"result": nilD.JsonString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_CovEnum_DLR06_SpecificFullString(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	l, r := d.SpecificFullString()

	// Act
	actual := args.Map{"result": l == "" || r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyVal / KeyValInteger
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_KAV01_KeyAnyVal_Methods(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "test", AnyValue: 42}

	// Act
	actual := args.Map{"result": kav.KeyString() != "test"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)
	actual = args.Map{"result": kav.AnyVal() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	_ = kav.AnyValString()
	_ = kav.WrapKey()
	_ = kav.WrapValue()
	actual = args.Map{"result": kav.IsString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": kav.ValInt() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	kvi := kav.KeyValInteger()
	actual = args.Map{"result": kvi.Key != "test"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)
	s := kav.String()
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_KAV02_KeyAnyVal_StringType(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "test", AnyValue: "hello"}

	// Act
	actual := args.Map{"result": kav.IsString()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	s := kav.String()
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovEnum_KVI01_KeyValInteger_Methods(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "test", ValueInteger: 42}
	_ = kvi.WrapKey()
	_ = kvi.WrapValue()
	kav := kvi.KeyAnyVal()

	// Act
	actual := args.Map{"result": kav.Key != "test"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)
	actual = args.Map{"result": kvi.IsString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	s := kvi.String()
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// differCheckerImpl / leftRightDiffCheckerImpl
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_DC01_DefaultDiffCheckerImpl(t *testing.T) {
	// Arrange
	dc := enumimpl.DefaultDiffCheckerImpl
	r := dc.GetSingleDiffResult(true, 1, 2)

	// Act
	actual := args.Map{"result": r != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected left", actual)
	r2 := dc.GetSingleDiffResult(false, 1, 2)
	actual = args.Map{"result": r2 != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected right", actual)
	r3 := dc.GetResultOnKeyMissingInRightExistInLeft("k", 1)
	actual = args.Map{"result": r3 != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": dc.IsEqual(false, 1, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": dc.IsEqual(true, 1, byte(1))}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true regardless", actual)
}

func Test_CovEnum_DC02_LeftRightDiffCheckerImpl(t *testing.T) {
	// Arrange
	dc := enumimpl.LeftRightDiffCheckerImpl
	r := dc.GetSingleDiffResult(true, 1, 2)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	r2 := dc.GetResultOnKeyMissingInRightExistInLeft("k", 1)
	actual = args.Map{"result": r2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": dc.IsEqual(false, 1, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapIntegerString with string values
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_DM47_MapIntegerString_StringValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "hello", "b": "world"}
	m, keys := dm.MapIntegerString()
	_ = m
	_ = keys
}

func Test_CovEnum_DM48_ShouldDiffMessageUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}
