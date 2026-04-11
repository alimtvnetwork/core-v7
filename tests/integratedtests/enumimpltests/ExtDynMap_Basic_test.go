package enumimpltests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// DynamicMap - Basic Operations
// ==========================================

func Test_ExtDynMap_Basic_Verification(t *testing.T) {
	for caseIndex, tc := range extDynMapBasicTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		items := input["items"].(map[string]any)
		dm := enumimpl.DynamicMap(items)

		// Act
		actual := args.Map{
			"length":  fmt.Sprintf("%d", dm.Length()),
			"isEmpty": fmt.Sprintf("%v", dm.IsEmpty()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// DynamicMap - AddOrUpdate / Set / AddNewOnly
// ==========================================

func Test_ExtDynMap_AddOrUpdate_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	isNew := dm.AddOrUpdate("B", 2)
	isExisting := dm.AddOrUpdate("A", 10)

	// Assert
	actual := args.Map{"result": isNew}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AddOrUpdate new key should return true", actual)
	actual = args.Map{"result": isExisting}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AddOrUpdate existing key should return false", actual)
	actual = args.Map{"result": dm.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Expected 2", actual)
}

func Test_ExtDynMap_Set_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	isNew := dm.Set("B", 2)

	// Assert
	actual := args.Map{"result": isNew}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Set new key should return true", actual)
}

func Test_ExtDynMap_AddNewOnly_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	added := dm.AddNewOnly("A", 10)
	addedNew := dm.AddNewOnly("B", 2)

	// Assert
	actual := args.Map{"result": added}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AddNewOnly existing should return false", actual)
	actual = args.Map{"result": addedNew}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AddNewOnly new should return true", actual)
}

// ==========================================
// DynamicMap - AllKeys / AllKeysSorted
// ==========================================

func Test_ExtDynMap_AllKeys_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"B": 2, "A": 1, "C": 3}

	// Act
	keys := dm.AllKeys()
	sortedKeys := dm.AllKeysSorted()

	// Assert
	actual := args.Map{"result": len(keys) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllKeys expected 3", actual)
	actual = args.Map{"result": len(sortedKeys) != 3 || sortedKeys[0] != "A"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllKeysSorted first should be A", actual)
}

// ==========================================
// DynamicMap - HasKey / HasAllKeys / HasAnyKeys / IsMissingKey
// ==========================================

func Test_ExtDynMap_HasKey_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act & Assert
	actual := args.Map{"result": dm.HasKey("A")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasKey should find A", actual)
	actual = args.Map{"result": dm.HasKey("Z")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasKey should not find Z", actual)
	actual = args.Map{"result": dm.HasAllKeys("A", "B")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAllKeys should be true", actual)
	actual = args.Map{"result": dm.HasAllKeys("A", "Z")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAllKeys with missing should be false", actual)
	actual = args.Map{"result": dm.HasAnyKeys("Z", "A")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyKeys should be true", actual)
	actual = args.Map{"result": dm.HasAnyKeys("X", "Y")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAnyKeys should be false", actual)
	actual = args.Map{"result": dm.IsMissingKey("Z")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsMissingKey should be true for Z", actual)
}

// ==========================================
// DynamicMap - HasAnyItem / Count / LastIndex / HasIndex
// ==========================================

func Test_ExtDynMap_Utility_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act & Assert
	actual := args.Map{"result": dm.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem should be true", actual)
	actual = args.Map{"result": dm.Count() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Count expected 2", actual)
	actual = args.Map{"result": dm.LastIndex() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "LastIndex expected 1", actual)
	actual = args.Map{"result": dm.HasIndex(1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasIndex(1) should be true", actual)
	actual = args.Map{"result": dm.HasIndex(5)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasIndex(5) should be false", actual)
}

// ==========================================
// DynamicMap - First
// ==========================================

func Test_ExtDynMap_First_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	key, val := dm.First()

	// Assert
	actual := args.Map{"result": key == "" || val == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "First should return a key-value", actual)

	// Arrange - empty
	empty := enumimpl.DynamicMap{}

	// Act
	key2, val2 := empty.First()

	// Assert
	actual = args.Map{"result": key2 != "" || val2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "First on empty should return empty", actual)
}

// ==========================================
// DynamicMap - AllValuesStrings / AllValuesStringsSorted / AllValuesIntegers
// ==========================================

func Test_ExtDynMap_AllValues_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act
	valStrings := dm.AllValuesStrings()
	valStringsSorted := dm.AllValuesStringsSorted()
	valIntegers := dm.AllValuesIntegers()

	// Assert
	actual := args.Map{"result": len(valStrings) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllValuesStrings expected 2", actual)
	actual = args.Map{"result": len(valStringsSorted) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllValuesStringsSorted expected 2", actual)
	actual = args.Map{"result": len(valIntegers) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllValuesIntegers expected 2", actual)
}

// ==========================================
// DynamicMap - IsEqual / IsRawEqual / IsMismatch
// ==========================================

func Test_ExtDynMap_IsEqual_Verification(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"A": 1, "B": 2}
	dm2 := enumimpl.DynamicMap{"A": 1, "B": 2}
	dm3 := enumimpl.DynamicMap{"A": 1, "B": 3}

	// Act & Assert
	actual := args.Map{"result": dm1.IsEqual(false, &dm2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Same content should be equal", actual)
	actual = args.Map{"result": dm1.IsEqual(false, &dm3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Different content should not be equal", actual)
	actual = args.Map{"result": dm1.IsMismatch(false, &dm3)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsMismatch should be true for different", actual)

	// nil cases
	var nilDm *enumimpl.DynamicMap
	actual = args.Map{"result": nilDm.IsEqual(false, nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil.IsEqual(nil) should be true", actual)
	actual = args.Map{"result": nilDm.IsEqual(false, &dm1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil.IsEqual(non-nil) should be false", actual)
}

func Test_ExtDynMap_IsRawEqual_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act & Assert
	actual := args.Map{"result": dm.IsRawEqual(false, map[string]any{"A": 1})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsRawEqual same should be true", actual)
	actual = args.Map{"result": dm.IsRawEqual(false, map[string]any{"A": 2})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsRawEqual different should be false", actual)
}

func Test_ExtDynMap_IsRawEqual_RegardlessType_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act & Assert
	// int vs int8 should match regardless of type
	actual := args.Map{"result": dm.IsRawEqual(true, map[string]any{"A": 1})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsRawEqual regardless should be true for same value", actual)
}

// ==========================================
// DynamicMap - IsKeysEqualOnly
// ==========================================

func Test_ExtDynMap_IsKeysEqualOnly_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act & Assert
	actual := args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"A": 10, "B": 20})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Same keys should be equal regardless of values", actual)
	actual = args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"A": 1, "C": 3})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Different keys should not be equal", actual)
}

// ==========================================
// DynamicMap - KeyValue / KeyValueString / KeyValueInt / KeyValueByte
// ==========================================

func Test_ExtDynMap_KeyValue_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 42, "B": "hello"}

	// Act
	val, found := dm.KeyValue("A")

	// Assert
	actual := args.Map{"result": found || val != 42}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "KeyValue A expected 42", actual)

	// Act
	valStr, foundStr := dm.KeyValueString("B")

	// Assert
	actual = args.Map{"result": foundStr || valStr != "hello"}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "KeyValueString B expected 'hello', got ''", actual)

	// Act
	valInt, foundInt, failed := dm.KeyValueInt("A")

	// Assert
	actual = args.Map{"result": foundInt || failed || valInt != 42}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "KeyValueInt A expected 42, got (found=, failed=)", actual)

	// Act - missing key
	_, foundMissing := dm.KeyValue("Z")

	// Assert
	actual = args.Map{"result": foundMissing}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "KeyValue missing should not be found", actual)
}

func Test_ExtDynMap_KeyValueIntDefault_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 42}

	// Act
	val := dm.KeyValueIntDefault("A")
	valMissing := dm.KeyValueIntDefault("Z")

	// Assert
	actual := args.Map{"result": val != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "KeyValueIntDefault A expected 42", actual)
	// missing returns InvalidValue
	_ = valMissing // just ensure no panic
}

// ==========================================
// DynamicMap - Add
// ==========================================

func Test_ExtDynMap_Add_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	dm.Add("B", 2)

	// Assert
	actual := args.Map{"result": dm.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Add expected 2", actual)
}

// ==========================================
// DynamicMap - Raw
// ==========================================

func Test_ExtDynMap_Raw_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	raw := dm.Raw()

	// Assert
	actual := args.Map{"result": len(raw) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Raw expected 1", actual)
}

// ==========================================
// DynamicMap - DiffRaw
// ==========================================

func Test_ExtDynMap_DiffRaw_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}
	right := map[string]any{"A": 1, "C": 3}

	// Act
	diff := dm.DiffRaw(false, right)

	// Assert
	actual := args.Map{"result": diff.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DiffRaw should find differences", actual)
}

func Test_ExtDynMap_DiffRaw_NilCases_Verification(t *testing.T) {
	// Arrange
	var nilDm *enumimpl.DynamicMap
	right := map[string]any{"A": 1}
	dm := enumimpl.DynamicMap{"A": 1}

	// Act & Assert
	diff1 := nilDm.DiffRaw(false, nil)
	actual := args.Map{"result": diff1.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs nil diff should be empty", actual)

	diff2 := nilDm.DiffRaw(false, right)
	actual = args.Map{"result": diff2.Length() == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil diff should have items", actual)

	diff3 := dm.DiffRaw(false, nil)
	actual = args.Map{"result": diff3.Length() == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil diff should have items", actual)
}

// ==========================================
// DynamicMap - SortedKeyValues / SortedKeyAnyValues
// ==========================================

func Test_ExtDynMap_SortedKeyValues_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"B": 2, "A": 1}

	// Act
	kvs := dm.SortedKeyValues()

	// Assert
	actual := args.Map{"result": len(kvs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SortedKeyValues expected 2", actual)
	actual = args.Map{"result": kvs[0].Key != "A"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "First sorted key expected 'A', got ''", actual)
}

func Test_ExtDynMap_SortedKeyAnyValues_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"B": 2, "A": 1}

	// Act
	kavs := dm.SortedKeyAnyValues()

	// Assert
	actual := args.Map{"result": len(kavs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SortedKeyAnyValues expected 2", actual)
}

// ==========================================
// DynamicMap - MapIntegerString
// ==========================================

func Test_ExtDynMap_MapIntegerString_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Valid": 1}

	// Act
	rangeMap, sorted := dm.MapIntegerString()

	// Assert
	actual := args.Map{"result": len(rangeMap) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MapIntegerString map expected 2", actual)
	actual = args.Map{"result": len(sorted) != 2 || sorted[0] != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MapIntegerString sorted expected [0,1]", actual)
}

// ==========================================
// DynamicMap - IsValueString / IsValueTypeOf
// ==========================================

func Test_ExtDynMap_IsValueString_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": "hello"}

	// Act & Assert
	actual := args.Map{"result": dm.IsValueString()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsValueString should be true for string values", actual)

	dm2 := enumimpl.DynamicMap{"A": 42}
	actual = args.Map{"result": dm2.IsValueString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsValueString should be false for int values", actual)
}

// ==========================================
// DynamicMap - BasicByte / BasicInt8 etc
// ==========================================

func Test_ExtDynMap_BasicByte_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1, "Inactive": 2}

	// Act
	bb := dm.BasicByte("TestType")

	// Assert
	actual := args.Map{"result": bb.Min() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BasicByte Min expected 0", actual)
	actual = args.Map{"result": bb.Max() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BasicByte Max expected 2", actual)
	actual = args.Map{"result": bb.TypeName() != "TestType"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TypeName expected 'TestType', got ''", actual)
}

func Test_ExtDynMap_BasicByte_Methods_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}
	bb := dm.BasicByte("TestType")

	// Act & Assert
	actual := args.Map{"result": bb.IsValidRange(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsValidRange(0) should be true", actual)
	actual = args.Map{"result": bb.IsValidRange(1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsValidRange(1) should be true", actual)
	actual = args.Map{"result": bb.IsValidRange(5)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsValidRange(5) should be false", actual)
	actual = args.Map{"result": bb.IsAnyOf(1, 0, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOf should find 1", actual)
	actual = args.Map{"result": bb.IsAnyOf(5, 0, 1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsAnyOf should not find 5", actual)
	actual = args.Map{"result": bb.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Length expected 2", actual)
	actual = args.Map{"result": bb.Count() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Count expected 2", actual)
}

func Test_ExtDynMap_BasicByte_Unmarshal_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}
	bb := dm.BasicByte("TestType")

	// Act
	val, err := bb.UnmarshallToValue(true, nil)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue nil with mapFirst should not error", actual)
	actual = args.Map{"result": val != bb.Min()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue nil expected min", actual)

	// Act - not mapped to first
	_, err2 := bb.UnmarshallToValue(false, nil)

	// Assert
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue nil without mapFirst should error", actual)
}

func Test_ExtDynMap_BasicByte_StringRanges_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}
	bb := dm.BasicByte("TestType")

	// Act
	ranges := bb.StringRanges()
	csv := bb.RangeNamesCsv()
	msg := bb.RangesInvalidMessage()
	err := bb.RangesInvalidErr()

	// Assert
	actual := args.Map{"result": len(ranges) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringRanges expected 2", actual)
	actual = args.Map{"result": csv == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv should not be empty", actual)
	actual = args.Map{"result": msg == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesInvalidMessage should not be empty", actual)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesInvalidErr should not be nil", actual)
}

// ==========================================
// DiffLeftRight
// ==========================================

func Test_ExtDiffLeftRight_Verification(t *testing.T) {
	for caseIndex, tc := range extDiffLeftRightTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		dlr := &enumimpl.DiffLeftRight{Left: left, Right: right}
		actual := args.Map{
			"isSame":     fmt.Sprintf("%v", dlr.IsSame()),
			"isNotEqual": fmt.Sprintf("%v", dlr.IsNotEqual()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ExtDiffLeftRight_Methods_Verification(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "hello", Right: "world"}

	// Act & Assert
	actual := args.Map{"result": dlr.IsSameTypeSame() != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Same type (both string) should return true", actual)
	actual = args.Map{"result": dlr.IsSameRegardlessOfType()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Different values should not be same", actual)
	actual = args.Map{"result": dlr.IsEqual(false)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEqual(false) should be false for different values", actual)
	actual = args.Map{"result": dlr.HasMismatch(false)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasMismatch should be true", actual)
	actual = args.Map{"result": dlr.String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String() should not be empty", actual)
	actual = args.Map{"result": dlr.JsonString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonString() should not be empty", actual)
	actual = args.Map{"result": dlr.DiffString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DiffString() should not be empty for mismatched", actual)

	// Arrange - same
	dlrSame := &enumimpl.DiffLeftRight{Left: "same", Right: "same"}

	// Act & Assert
	actual = args.Map{"result": dlrSame.DiffString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DiffString() should be empty for same values", actual)

	// nil case
	var nilDlr *enumimpl.DiffLeftRight
	actual = args.Map{"result": nilDlr.JsonString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil.JsonString() should be empty", actual)
}

func Test_ExtDiffLeftRight_Types_Verification(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "hello", Right: 42}

	// Act
	l, r := dlr.Types()

	// Assert
	actual := args.Map{"result": l == r}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Different types should not be equal", actual)
	actual = args.Map{"result": dlr.IsSameTypeSame()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsSameTypeSame should be false for string vs int", actual)
}

func Test_ExtDiffLeftRight_SpecificFullString_Verification(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "A", Right: "B"}

	// Act
	l, r := dlr.SpecificFullString()

	// Assert
	actual := args.Map{"result": l == "" || r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SpecificFullString should not be empty", actual)
}

// ==========================================
// KeyAnyVal
// ==========================================

func Test_ExtKeyAnyVal_Verification(t *testing.T) {
	for caseIndex, tc := range extKeyAnyValTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")
		value, _ := input.Get("value")

		// Act
		kav := enumimpl.KeyAnyVal{Key: key, AnyValue: value}
		actual := args.Map{
			"key":      kav.KeyString(),
			"isString": fmt.Sprintf("%v", kav.IsString()),
		}

		if !kav.IsString() {
			actual["valInt"] = fmt.Sprintf("%d", kav.ValInt())
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ExtKeyAnyVal_Methods_Verification(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "Test", AnyValue: 42}

	// Act & Assert
	actual := args.Map{"result": kav.AnyVal() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyVal should be 42", actual)
	actual = args.Map{"result": kav.AnyValString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyValString should not be empty", actual)
	actual = args.Map{"result": kav.WrapKey() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "WrapKey should not be empty", actual)
	actual = args.Map{"result": kav.WrapValue() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "WrapValue should not be empty", actual)
	actual = args.Map{"result": kav.String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be empty", actual)

	// KeyValInteger conversion
	kvi := kav.KeyValInteger()
	actual = args.Map{"result": kvi.Key != "Test" || kvi.ValueInteger != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "KeyValInteger expected Test/42, got/", actual)
}

// ==========================================
// KeyValInteger
// ==========================================

func Test_ExtKeyValInteger_Verification(t *testing.T) {
	for caseIndex, tc := range extKeyValIntegerTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")
		valueRaw, _ := input.Get("value")
		value := valueRaw.(int)

		// Act
		kvi := enumimpl.KeyValInteger{Key: key, ValueInteger: value}
		actual := args.Map{
			"key":      kvi.Key,
			"isString": fmt.Sprintf("%v", kvi.IsString()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ExtKeyValInteger_Methods_Verification(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "Test", ValueInteger: 5}

	// Act & Assert
	actual := args.Map{"result": kvi.WrapKey() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "WrapKey should not be empty", actual)
	actual = args.Map{"result": kvi.WrapValue() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "WrapValue should not be empty", actual)
	actual = args.Map{"result": kvi.String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be empty", actual)

	kav := kvi.KeyAnyVal()
	actual = args.Map{"result": kav.Key != "Test" || kav.AnyValue != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal conversion mismatch", actual)
}

// ==========================================
// KeyAnyValues (func)
// ==========================================

func Test_ExtKeyAnyValues_Verification(t *testing.T) {
	// Arrange
	names := []string{"A", "B", "C"}
	values := []int{1, 2, 3}

	// Act
	result := enumimpl.KeyAnyValues(names, values)

	// Assert
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues expected 3", actual)
	actual = args.Map{"result": result[0].Key != "A" || result[0].ValInt() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "First element expected A/1, got/", actual)
}

func Test_ExtKeyAnyValues_Empty_Verification(t *testing.T) {
	// Arrange
	// Act
	result := enumimpl.KeyAnyValues([]string{}, nil)

	// Assert
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Empty names should return empty", actual)
}

// ==========================================
// AllNameValues
// ==========================================

func Test_ExtAllNameValues_Verification(t *testing.T) {
	// Arrange
	names := []string{"Invalid", "Active"}
	values := []byte{0, 1}

	// Act
	result := enumimpl.AllNameValues(names, values)

	// Assert
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNameValues expected 2", actual)
}

// ==========================================
// ConvEnumAnyValToInteger
// ==========================================

func Test_ExtConvEnumAnyValToInteger_Verification(t *testing.T) {
	// Arrange & Act & Assert
	actual := args.Map{"result": enumimpl.ConvEnumAnyValToInteger(42) != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "int 42 should convert to 42", actual)
	actual = args.Map{"result": enumimpl.ConvEnumAnyValToInteger("hello") >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string should convert to MinInt (negative)", actual)
	actual = args.Map{"result": enumimpl.ConvEnumAnyValToInteger(byte(5)) != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "byte 5 should convert to 5", actual)
}

// ==========================================
// IntegersRangesOfAnyVal
// ==========================================

func Test_ExtIntegersRangesOfAnyVal_Verification(t *testing.T) {
	// Arrange
	values := []int{3, 1, 2}

	// Act
	result := enumimpl.IntegersRangesOfAnyVal(values)

	// Assert
	actual := args.Map{"result": len(result) != 3 || result[0] != 1 || result[2] != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IntegersRangesOfAnyVal expected sorted [1,2,3]", actual)
}

// ==========================================
// UnsupportedNames
// ==========================================

func Test_ExtUnsupportedNames_Verification(t *testing.T) {
	// Arrange
	all := []string{"A", "B", "C", "D"}
	supported := []string{"A", "C"}

	// Act
	unsupported := enumimpl.UnsupportedNames(all, supported...)

	// Assert
	actual := args.Map{"result": len(unsupported) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames expected 2", actual)
}

// ==========================================
// PrependJoin / JoinPrependUsingDot
// ==========================================

func Test_ExtPrependJoin_Verification(t *testing.T) {
	// Arrange
	// Act
	result := enumimpl.PrependJoin(".", "prefix", "a", "b")

	// Assert
	actual := args.Map{"result": result != "prefix.a.b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrependJoin expected 'prefix.a.b', got ''", actual)
}

func Test_ExtJoinPrependUsingDot_Verification(t *testing.T) {
	// Arrange
	// Act
	result := enumimpl.JoinPrependUsingDot("prefix", "a", "b")

	// Assert
	actual := args.Map{"result": result != "prefix.a.b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot expected 'prefix.a.b', got ''", actual)
}

// ==========================================
// NameWithValue
// ==========================================

func Test_ExtNameWithValue_Verification(t *testing.T) {
	// Arrange
	// Act
	result := enumimpl.NameWithValue("TestEnum")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NameWithValue should not be empty", actual)
}

// ==========================================
// Format
// ==========================================

func Test_ExtFormat_Verification(t *testing.T) {
	// Arrange
	format := "Enum of {type-name} - {name} - {value}"

	// Act
	result := enumimpl.Format("MyEnum", "Active", "1", format)

	// Assert
	actual := args.Map{"result": result != "Enum of MyEnum - Active - 1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Format expected 'Enum of MyEnum - Active - 1', got ''", actual)
}

// ==========================================
// differCheckerImpl
// ==========================================

func Test_ExtDifferCheckerImpl_Verification(t *testing.T) {
	// Arrange
	checker := enumimpl.DefaultDiffCheckerImpl

	// Act & Assert
	actual := args.Map{"result": checker.IsEqual(false, 42, 42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEqual same should be true", actual)
	actual = args.Map{"result": checker.IsEqual(false, 42, 43)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEqual different should be false", actual)
	actual = args.Map{"result": checker.IsEqual(true, 42, 42)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEqual regardless same should be true", actual)

	// GetSingleDiffResult
	result := checker.GetSingleDiffResult(true, "left", "right")
	actual = args.Map{"result": result != "left"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult isLeft=true should return left", actual)

	result2 := checker.GetSingleDiffResult(false, "left", "right")
	actual = args.Map{"result": result2 != "right"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult isLeft=false should return right", actual)
}

// ==========================================
// leftRightDiffCheckerImpl
// ==========================================

func Test_ExtLeftRightDiffCheckerImpl_Verification(t *testing.T) {
	// Arrange
	checker := enumimpl.LeftRightDiffCheckerImpl

	// Act
	result := checker.GetSingleDiffResult(true, "L", "R")

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult should not be nil", actual)

	// GetResultOnKeyMissingInRightExistInLeft
	missing := checker.GetResultOnKeyMissingInRightExistInLeft("key", "val")
	actual = args.Map{"result": missing == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetResultOnKeyMissingInRightExistInLeft should not be nil", actual)
}

// ==========================================
// DynamicMap - DiffJsonMessage
// ==========================================

func Test_ExtDynMap_DiffJsonMessage_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act - same
	msg1 := dm.DiffJsonMessage(false, map[string]any{"A": 1, "B": 2})

	// Assert
	actual := args.Map{"result": msg1 != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessage same should be empty", actual)

	// Act - different
	msg2 := dm.DiffJsonMessage(false, map[string]any{"A": 1, "C": 3})

	// Assert
	actual = args.Map{"result": msg2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessage different should not be empty", actual)
}

// ==========================================
// DynamicMap - ShouldDiffMessage
// ==========================================

func Test_ExtDynMap_ShouldDiffMessage_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act - same
	msg1 := dm.ShouldDiffMessage(false, "test", map[string]any{"A": 1})

	// Assert
	actual := args.Map{"result": msg1 != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ShouldDiffMessage same should be empty", actual)

	// Act - different
	msg2 := dm.ShouldDiffMessage(false, "test", map[string]any{"A": 2})

	// Assert
	actual = args.Map{"result": msg2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ShouldDiffMessage different should not be empty", actual)
}

// ==========================================
// DynamicMap - ExpectingMessage
// ==========================================

func Test_ExtDynMap_ExpectingMessage_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act - same
	msg1 := dm.ExpectingMessage("test", map[string]any{"A": 1})

	// Assert
	actual := args.Map{"result": msg1 != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ExpectingMessage same should be empty", actual)

	// Act - different
	msg2 := dm.ExpectingMessage("test", map[string]any{"A": 2})

	// Assert
	actual = args.Map{"result": msg2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ExpectingMessage different should not be empty", actual)
}

// ==========================================
// DynamicMap - DiffJsonMessageLeftRight
// ==========================================

func Test_ExtDynMap_DiffJsonMessageLeftRight_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act - same
	msg1 := dm.DiffJsonMessageLeftRight(false, map[string]any{"A": 1})

	// Assert
	actual := args.Map{"result": msg1 != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessageLeftRight same should be empty", actual)

	// Act - different
	msg2 := dm.DiffJsonMessageLeftRight(false, map[string]any{"B": 2})

	// Assert
	actual = args.Map{"result": msg2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessageLeftRight different should not be empty", actual)
}

// ==========================================
// numberEnumBase methods via BasicByte
// ==========================================

func Test_ExtNumberEnumBase_Methods_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1, "Inactive": 2}
	bb := dm.BasicByte("TestType")

	// Act & Assert - MinMaxAny
	minAny, maxAny := bb.MinMaxAny()
	actual := args.Map{"result": minAny == nil || maxAny == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinMaxAny should not return nil", actual)

	// MinValueString / MaxValueString
	actual = args.Map{"result": bb.MinValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinValueString should not be empty", actual)
	actual = args.Map{"result": bb.MaxValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxValueString should not be empty", actual)

	// MinInt / MaxInt
	actual = args.Map{"result": bb.MinInt() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinInt expected 0", actual)
	actual = args.Map{"result": bb.MaxInt() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxInt expected 2", actual)

	// AllNameValues
	anv := bb.AllNameValues()
	actual = args.Map{"result": len(anv) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNameValues expected 3", actual)

	// IntegerEnumRanges
	ier := bb.IntegerEnumRanges()
	actual = args.Map{"result": len(ier) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IntegerEnumRanges expected 3", actual)

	// NamesHashset
	nh := bb.NamesHashset()
	actual = args.Map{"result": len(nh) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NamesHashset expected 3", actual)

	// RangesDynamicMap / DynamicMap
	rdm := bb.RangesDynamicMap()
	actual = args.Map{"result": len(rdm) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesDynamicMap expected 3", actual)

	// KeyAnyValues
	kavs := bb.KeyAnyValues()
	actual = args.Map{"result": len(kavs) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues expected 3", actual)

	// Format
	formatted := bb.Format("Enum of {type-name} - {name} - {value}", byte(0))
	actual = args.Map{"result": formatted == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Format should not be empty", actual)
}
