package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Map basic operations
// ==========================================

func Test_Map_Length(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	actual := args.Map{"result": m.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Map_Has(t *testing.T) {
	m := args.Map{"a": 1}
	actual := args.Map{"result": m.Has("a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have 'a'", actual)
	actual := args.Map{"result": m.Has("b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have 'b'", actual)
}

func Test_Map_Has_Nil(t *testing.T) {
	var m args.Map
	actual := args.Map{"result": m.Has("a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should return false", actual)
}

func Test_Map_HasDefined(t *testing.T) {
	m := args.Map{"a": "val", "b": nil}
	actual := args.Map{"result": m.HasDefined("a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be defined", actual)
}

func Test_Map_HasDefined_Nil(t *testing.T) {
	var m args.Map
	actual := args.Map{"result": m.HasDefined("a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should return false", actual)
}

func Test_Map_IsKeyMissing(t *testing.T) {
	m := args.Map{"a": 1}
	actual := args.Map{"result": m.IsKeyMissing("a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "'a' should not be missing", actual)
	actual := args.Map{"result": m.IsKeyMissing("b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "'b' should be missing", actual)
}

func Test_Map_IsKeyMissing_Nil(t *testing.T) {
	var m args.Map
	actual := args.Map{"result": m.IsKeyMissing("a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should return false (per implementation)", actual)
}

func Test_Map_IsKeyInvalid(t *testing.T) {
	m := args.Map{"a": "val"}
	actual := args.Map{"result": m.IsKeyInvalid("a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "'a' should not be invalid", actual)
}

func Test_Map_IsKeyInvalid_Nil(t *testing.T) {
	var m args.Map
	actual := args.Map{"result": m.IsKeyInvalid("a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should return false (per implementation)", actual)
}

func Test_Map_HasDefinedAll(t *testing.T) {
	m := args.Map{"a": "v1", "b": "v2"}
	actual := args.Map{"result": m.HasDefinedAll("a", "b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have all defined", actual)
}

func Test_Map_HasDefinedAll_Nil(t *testing.T) {
	var m args.Map
	actual := args.Map{"result": m.HasDefinedAll("a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_Map_HasDefinedAll_Empty(t *testing.T) {
	m := args.Map{"a": "v1"}
	actual := args.Map{"result": m.HasDefinedAll()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no names should return false", actual)
}

// ==========================================
// Map Get operations
// ==========================================

func Test_Map_Get_Cov(t *testing.T) {
	m := args.Map{"a": "val"}
	item, isValid := m.Get("a")
	actual := args.Map{"result": isValid || item != "val"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return valid item", actual)
}

func Test_Map_Get_Missing(t *testing.T) {
	m := args.Map{"a": "val"}
	_, isValid := m.Get("b")
	actual := args.Map{"result": isValid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing key should not be valid", actual)
}

func Test_Map_Get_Nil(t *testing.T) {
	var m args.Map
	_, isValid := m.Get("a")
	actual := args.Map{"result": isValid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should not be valid", actual)
}

func Test_Map_GetLowerCase(t *testing.T) {
	m := args.Map{"name": "val"}
	item, isValid := m.GetLowerCase("Name")
	actual := args.Map{"result": isValid || item != "val"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should find lowercase", actual)
}

func Test_Map_GetDirectLower(t *testing.T) {
	m := args.Map{"name": "val"}
	item := m.GetDirectLower("Name")
	actual := args.Map{"result": item != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should find lowercase", actual)
}

func Test_Map_GetDirectLower_Missing(t *testing.T) {
	m := args.Map{"name": "val"}
	item := m.GetDirectLower("Missing")
	actual := args.Map{"result": item != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing should return nil", actual)
}

// ==========================================
// Map semantic accessors
// ==========================================

func Test_Map_When(t *testing.T) {
	m := args.Map{"when": "condition"}
	actual := args.Map{"result": m.When() != "condition"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return when value", actual)
}

func Test_Map_Title(t *testing.T) {
	m := args.Map{"title": "test title"}
	actual := args.Map{"result": m.Title() != "test title"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return title value", actual)
}

func Test_Map_Expect(t *testing.T) {
	m := args.Map{"expect": "value"}
	actual := args.Map{"result": m.Expect() != "value"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return expect value", actual)
}

func Test_Map_Actual(t *testing.T) {
	m := args.Map{"actual": "value"}
	actual := args.Map{"result": m.Actual() != "value"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return actual value", actual)
}

func Test_Map_Arrange(t *testing.T) {
	m := args.Map{"arrange": "value"}
	actual := args.Map{"result": m.Arrange() != "value"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return arrange value", actual)
}

func Test_Map_SetActual_Cov(t *testing.T) {
	m := args.Map{}
	m.SetActual("hello")
	actual := args.Map{"result": m.Actual() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set actual", actual)
}

// ==========================================
// Map numbered items
// ==========================================

func Test_Map_FirstItem(t *testing.T) {
	m := args.Map{"first": "val"}
	actual := args.Map{"result": m.FirstItem() != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return first item", actual)
}

func Test_Map_SecondItem(t *testing.T) {
	m := args.Map{"second": "val"}
	actual := args.Map{"result": m.SecondItem() != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return second item", actual)
}

func Test_Map_ThirdItem(t *testing.T) {
	m := args.Map{"third": "val"}
	actual := args.Map{"result": m.ThirdItem() != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return third item", actual)
}

func Test_Map_FourthItem(t *testing.T) {
	m := args.Map{"fourth": "val"}
	actual := args.Map{"result": m.FourthItem() != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return fourth item", actual)
}

func Test_Map_FifthItem(t *testing.T) {
	m := args.Map{"fifth": "val"}
	actual := args.Map{"result": m.FifthItem() != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return fifth item", actual)
}

func Test_Map_SixthItem(t *testing.T) {
	m := args.Map{"sixth": "val"}
	actual := args.Map{"result": m.SixthItem() != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return sixth item", actual)
}

func Test_Map_Seventh(t *testing.T) {
	m := args.Map{"seventh": "val"}
	actual := args.Map{"result": m.Seventh() != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return seventh item", actual)
}

// ==========================================
// Map Expected
// ==========================================

func Test_Map_Expected(t *testing.T) {
	m := args.Map{"expected": "val"}
	actual := args.Map{"result": m.Expected() != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return expected value", actual)
}

func Test_Map_Expected_Alias(t *testing.T) {
	m := args.Map{"expects": "val"}
	actual := args.Map{"result": m.Expected() != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return expected from alias", actual)
}

func Test_Map_HasExpect(t *testing.T) {
	m := args.Map{"expected": "val"}
	actual := args.Map{"result": m.HasExpect()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have expect", actual)
}

func Test_Map_HasFirst(t *testing.T) {
	m := args.Map{"first": "val"}
	actual := args.Map{"result": m.HasFirst()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have first", actual)
}

// ==========================================
// Map Raw / Args / ValidArgs
// ==========================================

func Test_Map_Raw(t *testing.T) {
	m := args.Map{"a": 1}
	raw := m.Raw()
	actual := args.Map{"result": len(raw) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "raw should have 1 item", actual)
}

func Test_Map_Args_Cov(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	a := m.Args("a", "b")
	actual := args.Map{"result": len(a) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 args", actual)
}

func Test_Map_GetByIndex_Cov(t *testing.T) {
	m := args.Map{"a": 1}
	v := m.GetByIndex(0)
	actual := args.Map{"result": v == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return value at index 0", actual)
}

func Test_Map_GetByIndex_OutOfBounds(t *testing.T) {
	m := args.Map{"a": 1}
	v := m.GetByIndex(10)
	actual := args.Map{"result": v != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "out of bounds should return nil", actual)
}

// ==========================================
// Map SortedKeys
// ==========================================

func Test_Map_SortedKeys_Cov(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys, err := m.SortedKeys()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": len(keys) != 2 || keys[0] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted [a b]", actual)
}

func Test_Map_SortedKeys_Empty(t *testing.T) {
	m := args.Map{}
	keys, err := m.SortedKeys()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": len(keys) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty map should return empty keys", actual)
}

func Test_Map_SortedKeysMust(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys := m.SortedKeysMust()
	actual := args.Map{"result": len(keys) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ==========================================
// Map ArgsCount
// ==========================================

func Test_Map_ArgsCount(t *testing.T) {
	// HasFunc() always returns true (FuncWrap returns non-nil),
	// so ArgsCount = len - 1 (func) = 1
	m := args.Map{"a": 1, "b": 2}
	actual := args.Map{"result": m.ArgsCount() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Map_ArgsCount_WithExpected(t *testing.T) {
	// HasExpect=true, HasFunc=true => ArgsCount = 2 - 2 = 0
	m := args.Map{"a": 1, "expected": "val"}
	c := m.ArgsCount()
	actual := args.Map{"result": c != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 (excluding expected+func)", actual)
}

// ==========================================
// Map GetFirstOfNames
// ==========================================

func Test_Map_GetFirstOfNames(t *testing.T) {
	m := args.Map{"name": "val"}
	r := m.GetFirstOfNames("missing", "name")
	actual := args.Map{"result": r != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return first found", actual)
}

func Test_Map_GetFirstOfNames_Empty(t *testing.T) {
	m := args.Map{"name": "val"}
	r := m.GetFirstOfNames()
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty names should return nil", actual)
}
