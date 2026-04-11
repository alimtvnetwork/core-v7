package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// DiffLeftRight
// ==========================================

func Test_DiffLeftRight_Types_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: 1}
	l, r := d.Types()

	// Act
	actual := args.Map{"result": l == nil || r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "types should not be nil", actual)
}

func Test_DiffLeftRight_IsSameTypeSame_True_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"result": d.IsSameTypeSame()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same type should return true", actual)
}

func Test_DiffLeftRight_IsSameTypeSame_False_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: 1}

	// Act
	actual := args.Map{"result": d.IsSameTypeSame()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different types should return false", actual)
}

func Test_DiffLeftRight_IsSame_True(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}

	// Act
	actual := args.Map{"result": d.IsSame()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same values should return true", actual)
}

func Test_DiffLeftRight_IsSame_False(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"result": d.IsSame()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different values should return false", actual)
}

func Test_DiffLeftRight_IsSameRegardlessOfType(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}

	// Act
	actual := args.Map{"result": d.IsSameRegardlessOfType()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same value regardless of type should be true", actual)
}

func Test_DiffLeftRight_IsEqual_Regardless_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}

	// Act
	actual := args.Map{"result": d.IsEqual(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal regardless", actual)
}

func Test_DiffLeftRight_IsEqual_Strict_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}

	// Act
	actual := args.Map{"result": d.IsEqual(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal strict", actual)
}

func Test_DiffLeftRight_HasMismatch_Regardless_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"result": d.HasMismatch(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have mismatch", actual)
}

func Test_DiffLeftRight_HasMismatch_Strict_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"result": d.HasMismatch(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have mismatch", actual)
}

func Test_DiffLeftRight_IsNotEqual(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"result": d.IsNotEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should not be equal", actual)
}

func Test_DiffLeftRight_String(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"result": d.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_DiffLeftRight_JsonString_Nil_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	var d *enumimpl.DiffLeftRight

	// Act
	actual := args.Map{"result": d.JsonString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_DiffLeftRight_SpecificFullString_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	l, r := d.SpecificFullString()

	// Act
	actual := args.Map{"result": l == "" || r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_DiffLeftRight_DiffString_Same_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}

	// Act
	actual := args.Map{"result": d.DiffString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same values should return empty diff", actual)
}

func Test_DiffLeftRight_DiffString_Different_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"result": d.DiffString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different values should return non-empty diff", actual)
}

// ==========================================
// DynamicMap
// ==========================================

func Test_DynamicMap_AddOrUpdate_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"key1": "val1"}
	isNew := dm.AddOrUpdate("key2", "val2")

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be newly added", actual)
	isNew2 := dm.AddOrUpdate("key1", "updated")
	actual = args.Map{"result": isNew2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be updated, not new", actual)
}

func Test_DynamicMap_AllKeys(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	keys := dm.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_DynamicMap_AllKeys_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	keys := dm.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_DynamicMap_AllKeysSorted(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	keys := dm.AllKeysSorted()

	// Act
	actual := args.Map{"result": len(keys) != 2 || keys[0] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted [a b]", actual)
}

func Test_DynamicMap_AllValuesStrings(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	vals := dm.AllValuesStrings()

	// Act
	actual := args.Map{"result": len(vals) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_DynamicMap_AllValuesStringsSorted_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "b", "c": "a"}
	vals := dm.AllValuesStringsSorted()

	// Act
	actual := args.Map{"result": len(vals) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_DynamicMap_Length(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_DynamicMap_Length_Nil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap

	// Act
	actual := args.Map{"result": dm.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_DynamicMap_Count_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.Count() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_DynamicMap_IsEmpty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}

	// Act
	actual := args.Map{"result": dm.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_DynamicMap_HasAnyItem(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_DynamicMap_HasKey(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.HasKey("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have key 'a'", actual)
	actual = args.Map{"result": dm.HasKey("b")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have key 'b'", actual)
}

func Test_DynamicMap_HasAllKeys_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.HasAllKeys("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have all keys", actual)
	actual = args.Map{"result": dm.HasAllKeys("a", "c")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have all keys", actual)
}

func Test_DynamicMap_HasAnyKeys_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.HasAnyKeys("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have any key", actual)
	actual = args.Map{"result": dm.HasAnyKeys("b", "c")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have any key", actual)
}

func Test_DynamicMap_IsMissingKey(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsMissingKey("b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be missing 'b'", actual)
}

func Test_DynamicMap_Raw(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	raw := dm.Raw()

	// Act
	actual := args.Map{"result": len(raw) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "raw should have 1 item", actual)
}

func Test_DynamicMap_First(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	k, v := dm.First()

	// Act
	actual := args.Map{"result": k == "" || v == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return first item", actual)
}

func Test_DynamicMap_First_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	k, v := dm.First()

	// Act
	actual := args.Map{"result": k != "" || v != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty map should return empty first", actual)
}

func Test_DynamicMap_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var a, b *enumimpl.DynamicMap

	// Act
	actual := args.Map{"result": a.IsEqual(true, b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_DynamicMap_IsEqual_OneNil(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsEqual(true, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should not be equal", actual)
}

func Test_DynamicMap_IsEqual_Same(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	dm2 := &enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsEqual(true, dm2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_DynamicMap_IsMismatch_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	dm2 := &enumimpl.DynamicMap{"a": 2}

	// Act
	actual := args.Map{"result": dm.IsMismatch(false, dm2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be mismatch", actual)
}

// ==========================================
// Format
// ==========================================

func Test_Format_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	result := enumimpl.Format(
		"MyEnum",
		"Invalid",
		"0",
		"Enum of {type-name} - {name} - {value}",
	)

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ==========================================
// NameWithValue
// ==========================================

func Test_NameWithValue_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	result := enumimpl.NameWithValue("test")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ==========================================
// UnsupportedNames
// ==========================================

func Test_UnsupportedNames_AllSupported_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	result := enumimpl.UnsupportedNames([]string{"a", "b"}, "a", "b")

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 unsupported", actual)
}

func Test_UnsupportedNames_SomeUnsupported(t *testing.T) {
	// Arrange
	result := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a")

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 unsupported", actual)
}

// ==========================================
// OnlySupportedErr
// ==========================================

func Test_OnlySupportedErr_AllSupported_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b"}, "a", "b")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all supported should return nil", actual)
}

func Test_OnlySupportedErr_SomeUnsupported(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b", "c"}, "a")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unsupported should return error", actual)
}

func Test_OnlySupportedErr_EmptyAllNames_FromDiffLeftRightTypes(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty allNames should return nil", actual)
}
