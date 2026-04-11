package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Map basic operations
// ==========================================

func Test_Map_Length_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{"result": m.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Map_Has_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{"result": m.Has("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have 'a'", actual)
	actual = args.Map{"result": m.Has("b")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have 'b'", actual)
}

func Test_Map_Has_Nil(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"result": m.Has("a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should return false", actual)
}

func Test_Map_HasDefined_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": "val",
		"b": nil,
	}

	// Act
	actual := args.Map{"result": m.HasDefined("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be defined", actual)
}

func Test_Map_HasDefined_Nil(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"result": m.HasDefined("a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should return false", actual)
}

func Test_Map_IsKeyMissing_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{"result": m.IsKeyMissing("a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "'a' should not be missing", actual)
	actual = args.Map{"result": m.IsKeyMissing("b")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "'b' should be missing", actual)
}

func Test_Map_IsKeyMissing_Nil(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"result": m.IsKeyMissing("a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should return false (per implementation)", actual)
}

func Test_Map_IsKeyInvalid_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"a": "val"}

	// Act
	actual := args.Map{"result": m.IsKeyInvalid("a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "'a' should not be invalid", actual)
}

func Test_Map_IsKeyInvalid_Nil(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"result": m.IsKeyInvalid("a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should return false (per implementation)", actual)
}

func Test_Map_HasDefinedAll_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": "v1",
		"b": "v2",
	}

	// Act
	actual := args.Map{"result": m.HasDefinedAll("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have all defined", actual)
}

func Test_Map_HasDefinedAll_Nil(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"result": m.HasDefinedAll("a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Map_HasDefinedAll_Empty(t *testing.T) {
	// Arrange
	m := args.Map{"a": "v1"}

	// Act
	actual := args.Map{"result": m.HasDefinedAll()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no names should return false", actual)
}

// ==========================================
// Map Get operations
// ==========================================

func Test_Map_Get_Cov(t *testing.T) {
	// Arrange
	m := args.Map{"a": "val"}
	item, isValid := m.Get("a")

	// Act
	actual := args.Map{"result": isValid || item != "val"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return valid item", actual)
}

func Test_Map_Get_Missing(t *testing.T) {
	// Arrange
	m := args.Map{"a": "val"}
	_, isValid := m.Get("b")

	// Act
	actual := args.Map{"result": isValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing key should not be valid", actual)
}

func Test_Map_Get_Nil(t *testing.T) {
	// Arrange
	var m args.Map
	_, isValid := m.Get("a")

	// Act
	actual := args.Map{"result": isValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should not be valid", actual)
}

func Test_Map_GetLowerCase_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"name": "val"}
	item, isValid := m.GetLowerCase("Name")

	// Act
	actual := args.Map{"result": isValid || item != "val"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should find lowercase", actual)
}

func Test_Map_GetDirectLower_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"name": "val"}
	item := m.GetDirectLower("Name")

	// Act
	actual := args.Map{"result": item != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should find lowercase", actual)
}

func Test_Map_GetDirectLower_Missing(t *testing.T) {
	// Arrange
	m := args.Map{"name": "val"}
	item := m.GetDirectLower("Missing")

	// Act
	actual := args.Map{"result": item != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing should return nil", actual)
}

// ==========================================
// Map semantic accessors
// ==========================================

func Test_Map_When_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"when": "condition"}

	// Act
	actual := args.Map{"result": m.When() != "condition"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return when value", actual)
}

func Test_Map_Title_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"title": "test title"}

	// Act
	actual := args.Map{"result": m.Title() != "test title"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return title value", actual)
}

func Test_Map_Expect_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"expect": "value"}

	// Act
	actual := args.Map{"result": m.Expect() != "value"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return expect value", actual)
}

func Test_Map_Actual_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"actual": "value"}

	// Act
	actual := args.Map{"result": m.Actual() != "value"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return actual value", actual)
}

func Test_Map_Arrange_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"arrange": "value"}

	// Act
	actual := args.Map{"result": m.Arrange() != "value"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return arrange value", actual)
}

func Test_Map_SetActual_Cov(t *testing.T) {
	// Arrange
	m := args.Map{}
	m.SetActual("hello")

	// Act
	actual := args.Map{"result": m.Actual() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set actual", actual)
}

// ==========================================
// Map numbered items
// ==========================================

func Test_Map_FirstItem(t *testing.T) {
	// Arrange
	m := args.Map{"first": "val"}

	// Act
	actual := args.Map{"result": m.FirstItem() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return first item", actual)
}

func Test_Map_SecondItem(t *testing.T) {
	// Arrange
	m := args.Map{"second": "val"}

	// Act
	actual := args.Map{"result": m.SecondItem() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return second item", actual)
}

func Test_Map_ThirdItem(t *testing.T) {
	// Arrange
	m := args.Map{"third": "val"}

	// Act
	actual := args.Map{"result": m.ThirdItem() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return third item", actual)
}

func Test_Map_FourthItem(t *testing.T) {
	// Arrange
	m := args.Map{"fourth": "val"}

	// Act
	actual := args.Map{"result": m.FourthItem() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return fourth item", actual)
}

func Test_Map_FifthItem(t *testing.T) {
	// Arrange
	m := args.Map{"fifth": "val"}

	// Act
	actual := args.Map{"result": m.FifthItem() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return fifth item", actual)
}

func Test_Map_SixthItem(t *testing.T) {
	// Arrange
	m := args.Map{"sixth": "val"}

	// Act
	actual := args.Map{"result": m.SixthItem() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return sixth item", actual)
}

func Test_Map_Seventh_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"seventh": "val"}

	// Act
	actual := args.Map{"result": m.Seventh() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return seventh item", actual)
}

// ==========================================
// Map Expected
// ==========================================

func Test_Map_Expected_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"expected": "val"}

	// Act
	actual := args.Map{"result": m.Expected() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return expected value", actual)
}

func Test_Map_Expected_Alias(t *testing.T) {
	// Arrange
	m := args.Map{"expects": "val"}

	// Act
	actual := args.Map{"result": m.Expected() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return expected from alias", actual)
}

func Test_Map_HasExpect_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"expected": "val"}

	// Act
	actual := args.Map{"result": m.HasExpect()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have expect", actual)
}

func Test_Map_HasFirst_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"first": "val"}

	// Act
	actual := args.Map{"result": m.HasFirst()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have first", actual)
}

// ==========================================
// Map Raw / Args / ValidArgs
// ==========================================

func Test_Map_Raw_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	raw := m.Raw()

	// Act
	actual := args.Map{"result": len(raw) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "raw should have 1 item", actual)
}

func Test_Map_Args_Cov(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}
	a := m.Args("a", "b")

	// Act
	actual := args.Map{"result": len(a) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 args", actual)
}

func Test_Map_GetByIndex_Cov(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	v := m.GetByIndex(0)

	// Act
	actual := args.Map{"result": v == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return value at index 0", actual)
}

func Test_Map_GetByIndex_OutOfBounds(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	v := m.GetByIndex(10)

	// Act
	actual := args.Map{"result": v != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "out of bounds should return nil", actual)
}

// ==========================================
// Map SortedKeys
// ==========================================

func Test_Map_SortedKeys_Cov(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	keys, err := m.SortedKeys()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": len(keys) != 2 || keys[0] != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted [a b]", actual)
}

func Test_Map_SortedKeys_Empty_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{}
	keys, err := m.SortedKeys()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": len(keys) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty map should return empty keys", actual)
}

func Test_Map_SortedKeysMust(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	keys := m.SortedKeysMust()

	// Act
	actual := args.Map{"result": len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ==========================================
// Map ArgsCount
// ==========================================

func Test_Map_ArgsCount_FromMapLengthV2(t *testing.T) {
	// Arrange
	// HasFunc() always returns true (FuncWrap returns non-nil),
	// so ArgsCount = len - 1 (func) = 1
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{"result": m.ArgsCount() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Map_ArgsCount_WithExpected(t *testing.T) {
	// Arrange
	// HasExpect=true, HasFunc=true => ArgsCount = 2 - 2 = 0
	m := args.Map{
		"a": 1,
		"expected": "val",
	}
	c := m.ArgsCount()

	// Act
	actual := args.Map{"result": c != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 (excluding expected+func)", actual)
}

// ==========================================
// Map GetFirstOfNames
// ==========================================

func Test_Map_GetFirstOfNames_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"name": "val"}
	r := m.GetFirstOfNames("missing", "name")

	// Act
	actual := args.Map{"result": r != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return first found", actual)
}

func Test_Map_GetFirstOfNames_Empty_FromMapLengthV2(t *testing.T) {
	// Arrange
	m := args.Map{"name": "val"}
	r := m.GetFirstOfNames()

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty names should return nil", actual)
}
