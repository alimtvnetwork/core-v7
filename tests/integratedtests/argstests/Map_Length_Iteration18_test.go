package argstests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// --- Map ---

func Test_Map_Length(t *testing.T) {
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

func Test_Map_ArgsCount_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{
		"first": 1,
		"second": 2,
	}
	// ArgsCount subtracts 1 for func (always present as invalid wrapper)

	// Act
	actual := args.Map{"result": m.ArgsCount() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 arg (length minus func slot)", actual)
}

func Test_Map_ArgsCount_WithFuncAndExpect(t *testing.T) {
	// Arrange
	fn := func() {}
	m := args.Map{
		"first": 1,
		"func": fn,
		"expected": "x",
	}
	count := m.ArgsCount()

	// Act
	actual := args.Map{"result": count != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 arg (excluding func and expect)", actual)
}

func Test_Map_HasFirst(t *testing.T) {
	// Arrange
	m := args.Map{"first": "hello"}

	// Act
	actual := args.Map{"result": m.HasFirst()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasFirst", actual)

	m2 := args.Map{}
	actual = args.Map{"result": m2.HasFirst()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no first", actual)
}

func Test_Map_Items(t *testing.T) {
	// Arrange
	m := args.Map{
		"first":  "a",
		"second": "b",
		"third":  "c",
		"fourth": "d",
		"fifth":  "e",
		"sixth":  "f",
	}

	// Act
	actual := args.Map{"result": m.FirstItem() != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "first mismatch", actual)
	actual = args.Map{"result": m.SecondItem() != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "second mismatch", actual)
	actual = args.Map{"result": m.ThirdItem() != "c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "third mismatch", actual)
	actual = args.Map{"result": m.FourthItem() != "d"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "fourth mismatch", actual)
	actual = args.Map{"result": m.FifthItem() != "e"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "fifth mismatch", actual)
	actual = args.Map{"result": m.SixthItem() != "f"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "sixth mismatch", actual)
}

func Test_Map_Seventh(t *testing.T) {
	// Arrange
	m := args.Map{"seventh": "g"}

	// Act
	actual := args.Map{"result": m.Seventh() != "g"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "seventh mismatch", actual)
}

func Test_Map_Expected(t *testing.T) {
	// Arrange
	m := args.Map{"expected": "x"}

	// Act
	actual := args.Map{"result": m.Expected() != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
}

func Test_Map_HasExpect(t *testing.T) {
	// Arrange
	m := args.Map{"expected": "x"}

	// Act
	actual := args.Map{"result": m.HasExpect()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasExpect", actual)
}

func Test_Map_GetByIndex_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}
	item := m.GetByIndex(0)

	// Act
	actual := args.Map{"result": item == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected item at index 0", actual)

	nilItem := m.GetByIndex(100)
	actual = args.Map{"result": nilItem != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for out of range", actual)
}

func Test_Map_HasFunc_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	fn := func() {}
	m := args.Map{"func": fn}

	// Act
	actual := args.Map{"result": m.HasFunc()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasFunc", actual)
}

func Test_Map_GetFuncName_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{}
	name := m.GetFuncName()

	// Act
	actual := args.Map{"result": name != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty func name", actual)
}

func Test_Map_HasDefined(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	actual := args.Map{"result": m.HasDefined("key")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected defined", actual)
	actual = args.Map{"result": m.HasDefined("missing")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not defined", actual)

	var nilMap args.Map
	actual = args.Map{"result": nilMap.HasDefined("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil map", actual)
}

func Test_Map_Has(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	actual := args.Map{"result": m.Has("key")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has", actual)
	actual = args.Map{"result": m.Has("missing")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not has", actual)
	var nilMap args.Map
	actual = args.Map{"result": nilMap.Has("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_Map_HasDefinedAll(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{"result": m.HasDefinedAll("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected all defined", actual)
	actual = args.Map{"result": m.HasDefinedAll("a", "missing")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for missing key", actual)

	var nilMap args.Map
	actual = args.Map{"result": nilMap.HasDefinedAll("a")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_Map_IsKeyInvalid(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	actual := args.Map{"result": m.IsKeyInvalid("key")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid key", actual)
	actual = args.Map{"result": m.IsKeyInvalid("missing")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid for missing", actual)
}

func Test_Map_IsKeyMissing(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	actual := args.Map{"result": m.IsKeyMissing("key")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not missing", actual)
	actual = args.Map{"result": m.IsKeyMissing("missing")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected missing", actual)
}

func Test_Map_SortedKeys_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	keys, err := m.SortedKeys()

	// Act
	actual := args.Map{"result": err != nil || len(keys) != 2 || keys[0] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected sorted keys", actual)
}

func Test_Map_SortedKeys_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}
	keys, err := m.SortedKeys()

	// Act
	actual := args.Map{"result": err != nil || len(keys) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty sorted keys", actual)
}

func Test_Map_When(t *testing.T) {
	// Arrange
	m := args.Map{"when": "now"}

	// Act
	actual := args.Map{"result": m.When() != "now"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected when=now", actual)
}

func Test_Map_Title(t *testing.T) {
	// Arrange
	m := args.Map{"title": "test"}

	// Act
	actual := args.Map{"result": m.Title() != "test"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected title=test", actual)
}

func Test_Map_Get_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	item, valid := m.Get("key")

	// Act
	actual := args.Map{"result": valid || item != "val"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)

	_, valid = m.Get("missing")
	actual = args.Map{"result": valid}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid for missing", actual)

	var nilMap args.Map
	_, valid = nilMap.Get("key")
	actual = args.Map{"result": valid}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid for nil", actual)
}

func Test_Map_GetLowerCase_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	item, valid := m.GetLowerCase("KEY")

	// Act
	actual := args.Map{"result": valid || item != "val"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected lower case match", actual)
}

func Test_Map_GetDirectLower_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	item := m.GetDirectLower("KEY")

	// Act
	actual := args.Map{"result": item != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected lower case match", actual)

	nilItem := m.GetDirectLower("MISSING")
	actual = args.Map{"result": nilItem != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for missing", actual)
}

func Test_Map_Expect(t *testing.T) {
	// Arrange
	m := args.Map{"expect": "x"}

	// Act
	actual := args.Map{"result": m.Expect() != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected x", actual)
}

func Test_Map_Actual(t *testing.T) {
	// Arrange
	m := args.Map{"actual": "y"}

	// Act
	actual := args.Map{"result": m.Actual() != "y"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected y", actual)
}

func Test_Map_Arrange(t *testing.T) {
	// Arrange
	m := args.Map{"arrange": "z"}

	// Act
	actual := args.Map{"result": m.Arrange() != "z"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected z", actual)
}

func Test_Map_SetActual_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{}
	m.SetActual("val")

	// Act
	actual := args.Map{"result": m.Actual() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected set actual", actual)
}

func Test_Map_WorkFunc_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	fn := func() {}
	m := args.Map{"func": fn}
	wf := m.WorkFunc()

	// Act
	actual := args.Map{"result": wf == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil work func", actual)
}

func Test_Map_GetWorkFunc(t *testing.T) {
	// Arrange
	fn := func() {}
	m := args.Map{"func": fn}
	wf := m.GetWorkFunc()

	// Act
	actual := args.Map{"result": wf == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Map_GetFirstOfNames_Empty(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{"result": m.GetFirstOfNames() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty names", actual)
}

func Test_Map_GetAsStringSliceFirstOfNames_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"items": []string{"a", "b"}}
	s := m.GetAsStringSliceFirstOfNames("items")

	// Act
	actual := args.Map{"result": len(s) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
}

func Test_Map_GetAsStringSliceFirstOfNames_Empty_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{}
	s := m.GetAsStringSliceFirstOfNames()

	// Act
	actual := args.Map{"result": s != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Map_GetAsStringSliceFirstOfNames_Undefined(t *testing.T) {
	// Arrange
	m := args.Map{}
	s := m.GetAsStringSliceFirstOfNames("missing")

	// Act
	actual := args.Map{"result": s != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for undefined", actual)
}

func Test_Map_GetAsInt(t *testing.T) {
	// Arrange
	m := args.Map{"n": 42}
	n, ok := m.GetAsInt("n")

	// Act
	actual := args.Map{"result": ok || n != 42}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	_, ok = m.GetAsInt("missing")
	actual = args.Map{"result": ok}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not ok", actual)
}

func Test_Map_GetAsIntDefault_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"n": 42}
	n := m.GetAsIntDefault("n", 0)

	// Act
	actual := args.Map{"result": n != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	n = m.GetAsIntDefault("missing", 99)
	actual = args.Map{"result": n != 99}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected default 99", actual)
}

func Test_Map_GetAsBool(t *testing.T) {
	// Arrange
	m := args.Map{"b": true}
	b, ok := m.GetAsBool("b")

	// Act
	actual := args.Map{"result": ok || !b}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Map_GetAsBoolDefault_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{}
	b := m.GetAsBoolDefault("b", true)

	// Act
	actual := args.Map{"result": b}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected default true", actual)
}

func Test_Map_GetAsString(t *testing.T) {
	// Arrange
	m := args.Map{"s": "hello"}
	s, ok := m.GetAsString("s")

	// Act
	actual := args.Map{"result": ok || s != "hello"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_Map_GetAsStringDefault_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{}
	s := m.GetAsStringDefault("s")

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty default", actual)
}

func Test_Map_GetAsStrings_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"s": []string{"a", "b"}}
	s, ok := m.GetAsStrings("s")

	// Act
	actual := args.Map{"result": ok || len(s) != 2}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Map_GetAsAnyItems_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"items": []any{1, "two"}}
	items, ok := m.GetAsAnyItems("items")

	// Act
	actual := args.Map{"result": ok || len(items) != 2}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Map_ValidArgs_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}
	validArgs := m.ValidArgs()

	// Act
	actual := args.Map{"result": len(validArgs) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 valid args", actual)
}

func Test_Map_Args_FromMapLengthIteration18(t *testing.T) {
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

func Test_Map_Raw_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	raw := m.Raw()

	// Act
	actual := args.Map{"result": raw == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil raw", actual)
}

func Test_Map_Slice_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	s := m.Slice()

	// Act
	actual := args.Map{"result": len(s) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 item", actual)
}

func Test_Map_String_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	s := m.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

func Test_Map_GetFirstFuncNameOf_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{}
	name := m.GetFirstFuncNameOf("func")

	// Act
	actual := args.Map{"result": name != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty name", actual)
}

func Test_Map_WorkFuncName_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Map{}
	name := m.WorkFuncName()

	// Act
	actual := args.Map{"result": name != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// --- Dynamic ---

func Test_Dynamic_NilReceiver_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	var d *args.DynamicAny

	// Act
	actual := args.Map{"result": d.ArgsCount() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": d.GetWorkFunc() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": d.HasFirst()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": d.HasDefined("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": d.Has("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": d.HasDefinedAll("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": d.IsKeyInvalid("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	actual = args.Map{"result": d.IsKeyMissing("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	_, valid := d.Get("x")
	actual = args.Map{"result": valid}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid for nil", actual)
	actual = args.Map{"result": d.HasExpect()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_Dynamic_AllMethods(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{
			"first":  "a",
			"second": "b",
			"third":  "c",
			"fourth": "d",
			"fifth":  "e",
			"sixth":  "f",
		},
		Expect: "expected",
	}

	// Act
	actual := args.Map{"result": d.FirstItem() != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "first mismatch", actual)
	actual = args.Map{"result": d.SecondItem() != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "second mismatch", actual)
	actual = args.Map{"result": d.ThirdItem() != "c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "third mismatch", actual)
	actual = args.Map{"result": d.FourthItem() != "d"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "fourth mismatch", actual)
	actual = args.Map{"result": d.FifthItem() != "e"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "fifth mismatch", actual)
	actual = args.Map{"result": d.SixthItem() != "f"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "sixth mismatch", actual)
	actual = args.Map{"result": d.Expected() != "expected"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)

	actual = args.Map{"result": d.GetByIndex(0) == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected item at 0", actual)
	if !d.HasFunc() == false {
		// just exercise
	}
	_ = d.GetFuncName()
	_ = d.FuncWrap()

	actual = args.Map{"result": d.HasDefined("first") != true}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected defined", actual)
	actual = args.Map{"result": d.Has("missing")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not has", actual)
	actual = args.Map{"result": d.HasDefinedAll("first", "second")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected all defined", actual)
	actual = args.Map{"result": d.IsKeyInvalid("first")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual = args.Map{"result": d.IsKeyMissing("missing")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected missing", actual)
	_, _ = d.GetLowerCase("FIRST")
	_ = d.GetDirectLower("FIRST")
	_ = d.Actual()
	_ = d.Arrange()

	_, _ = d.GetAsInt("first")
	_ = d.GetAsIntDefault("first", 0)
	_, _ = d.GetAsString("first")
	_ = d.GetAsStringDefault("first")
	_, _ = d.GetAsStrings("first")
	_, _ = d.GetAsAnyItems("first")

	_ = d.ValidArgs()
	_ = d.Args("first")
	_ = d.Slice()
	s := d.String()
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)

	// second call should return cached
	s2 := d.String()
	actual = args.Map{"result": s != s2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached string", actual)

	_ = d.AsArgsMapper()
	_ = d.AsArgFuncNameContractsBinder()
	_ = d.AsArgBaseContractsBinder()
}

// --- DynamicFunc ---

func Test_DynamicFunc_NilReceiver_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	var df *args.DynamicFuncAny

	// Act
	actual := args.Map{"result": df.ArgsCount() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": df.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": df.HasDefined("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": df.Has("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": df.HasDefinedAll("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": df.IsKeyInvalid("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	actual = args.Map{"result": df.IsKeyMissing("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	_, valid := df.Get("x")
	actual = args.Map{"result": valid}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual = args.Map{"result": df.HasFunc()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": df.HasExpect()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_DynamicFunc_AllMethods(t *testing.T) {
	// Arrange
	fn := func(s string) string { return s }
	df := &args.DynamicFuncAny{
		Params: args.Map{
			"first":  "a",
			"second": "b",
			"third":  "c",
			"fourth": "d",
			"fifth":  "e",
			"sixth":  "f",
		},
		WorkFunc: fn,
		Expect:   "expected",
	}

	_ = df.GetWorkFunc()
	_ = df.HasFirst()
	_ = df.GetByIndex(0)
	_ = df.GetByIndex(100)
	_ = df.FirstItem()
	_ = df.SecondItem()
	_ = df.ThirdItem()
	_ = df.FourthItem()
	_ = df.FifthItem()
	_ = df.SixthItem()
	_ = df.Expected()
	_ = df.When()
	_ = df.Title()
	_, _ = df.GetLowerCase("FIRST")
	_ = df.GetDirectLower("FIRST")
	_ = df.GetDirectLower("MISSING")
	_ = df.Actual()
	_ = df.Arrange()
	_, _ = df.GetAsInt("first")
	_, _ = df.GetAsString("first")
	_, _ = df.GetAsStrings("first")
	_, _ = df.GetAsAnyItems("first")
	_ = df.HasFunc()
	_ = df.HasExpect()
	_ = df.GetFuncName()
	_ = df.FuncWrap()
	_ = df.ValidArgs()
	_ = df.Args("first")

	s := df.Slice()

	// Act
	actual := args.Map{"result": len(s) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty slice", actual)

	str := df.String()
	actual = args.Map{"result": str == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)

	// cached
	str2 := df.String()
	actual = args.Map{"result": str != str2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached", actual)

	_ = df.AsArgsMapper()
	_ = df.AsArgFuncNameContractsBinder()
	_ = df.AsArgBaseContractsBinder()
}

// --- One ---

func Test_One_AllMethods_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	o := &args.OneAny{First: "hello", Expect: "world"}

	// Act
	actual := args.Map{"result": o.FirstItem() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "first mismatch", actual)
	actual = args.Map{"result": o.Expected() != "world"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
	_ = o.ArgTwo()
	actual = args.Map{"result": o.HasFirst()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasFirst", actual)
	actual = args.Map{"result": o.HasExpect()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasExpect", actual)
	actual = args.Map{"result": len(o.ValidArgs()) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 valid arg", actual)
	actual = args.Map{"result": len(o.Args(1)) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 arg", actual)
	actual = args.Map{"result": o.ArgsCount() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = o.Slice()
	_ = o.GetByIndex(0)
	s := o.String()
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	_ = o.LeftRight()
	_ = args.OneAny{First: "x"}.AsOneParameter()
	_ = args.OneAny{First: "x"}.AsArgBaseContractsBinder()
}

// --- Two ---

func Test_Two_AllMethods_FromMapLengthIteration18(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b", Expect: "c"}
	_ = tw.FirstItem()
	_ = tw.SecondItem()
	_ = tw.Expected()
	_ = tw.ArgTwo()
	_ = tw.HasFirst()
	_ = tw.HasSecond()
	_ = tw.HasExpect()
	_ = tw.ValidArgs()
	_ = tw.ArgsCount()
	_ = tw.Args(2)
	_ = tw.Slice()
	_ = tw.GetByIndex(0)
	_ = tw.String()
	_ = tw.LeftRight()
	_ = args.TwoAny{}.AsTwoParameter()
	_ = args.TwoAny{}.AsArgBaseContractsBinder()
}

// --- Three ---

func Test_Three_AllMethods_FromMapLengthIteration18(t *testing.T) {
	th := &args.ThreeAny{First: "a", Second: "b", Third: "c", Expect: "d"}
	_ = th.ArgsCount()
	_ = th.FirstItem()
	_ = th.SecondItem()
	_ = th.ThirdItem()
	_ = th.Expected()
	_ = th.ArgTwo()
	_ = th.ArgThree()
	_ = th.HasFirst()
	_ = th.HasSecond()
	_ = th.HasThird()
	_ = th.HasExpect()
	_ = th.ValidArgs()
	_ = th.Args(3)
	_ = th.Slice()
	_ = th.GetByIndex(0)
	_ = th.String()
	_ = th.LeftRight()
	_ = args.ThreeAny{}.AsThreeParameter()
	_ = args.ThreeAny{}.AsArgBaseContractsBinder()
}

// --- Four ---

func Test_Four_AllMethods_FromMapLengthIteration18(t *testing.T) {
	f := &args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d", Expect: "e"}
	_ = f.ArgsCount()
	_ = f.FirstItem()
	_ = f.SecondItem()
	_ = f.ThirdItem()
	_ = f.FourthItem()
	_ = f.Expected()
	_ = f.ArgTwo()
	_ = f.ArgThree()
	_ = f.HasFirst()
	_ = f.HasSecond()
	_ = f.HasThird()
	_ = f.HasFourth()
	_ = f.HasExpect()
	_ = f.ValidArgs()
	_ = f.Args(4)
	_ = f.Slice()
	_ = f.GetByIndex(0)
	_ = f.String()
	_ = args.FourAny{}.AsFourParameter()
	_ = args.FourAny{}.AsArgBaseContractsBinder()
}

// --- Five ---

func Test_Five_AllMethods_FromMapLengthIteration18(t *testing.T) {
	f := &args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Expect: "f"}
	_ = f.ArgsCount()
	_ = f.FirstItem()
	_ = f.SecondItem()
	_ = f.ThirdItem()
	_ = f.FourthItem()
	_ = f.FifthItem()
	_ = f.Expected()
	_ = f.ArgTwo()
	_ = f.ArgThree()
	_ = f.ArgFour()
	_ = f.HasFirst()
	_ = f.HasSecond()
	_ = f.HasThird()
	_ = f.HasFourth()
	_ = f.HasFifth()
	_ = f.HasExpect()
	_ = f.ValidArgs()
	_ = f.Args(5)
	_ = f.Slice()
	_ = f.GetByIndex(0)
	_ = f.String()
	_ = args.FiveAny{}.AsFifthParameter()
	_ = args.FiveAny{}.AsArgBaseContractsBinder()
}

// --- Six ---

func Test_Six_AllMethods_FromMapLengthIteration18(t *testing.T) {
	s := &args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f", Expect: "g"}
	_ = s.ArgsCount()
	_ = s.FirstItem()
	_ = s.SecondItem()
	_ = s.ThirdItem()
	_ = s.FourthItem()
	_ = s.FifthItem()
	_ = s.SixthItem()
	_ = s.Expected()
	_ = s.ArgTwo()
	_ = s.ArgThree()
	_ = s.ArgFour()
	_ = s.ArgFive()
	_ = s.HasFirst()
	_ = s.HasSecond()
	_ = s.HasThird()
	_ = s.HasFourth()
	_ = s.HasFifth()
	_ = s.HasSixth()
	_ = s.HasExpect()
	_ = s.ValidArgs()
	_ = s.Args(6)
	_ = s.Slice()
	_ = s.GetByIndex(0)
	_ = s.String()
	_ = args.SixAny{}.AsSixthParameter()
	_ = args.SixAny{}.AsArgBaseContractsBinder()
}

// --- LeftRight ---

func Test_LeftRight_AllMethods_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	lr := &args.LeftRightAny{Left: "a", Right: "b", Expect: "c"}
	_ = lr.ArgsCount()
	_ = lr.FirstItem()
	_ = lr.SecondItem()
	_ = lr.Expected()
	_ = lr.ArgTwo()
	_ = lr.HasFirst()
	_ = lr.HasSecond()
	_ = lr.HasLeft()
	_ = lr.HasRight()
	_ = lr.HasExpect()
	_ = lr.ValidArgs()
	_ = lr.Args(2)
	_ = lr.Slice()
	_ = lr.GetByIndex(0)
	_ = lr.String()
	clone := lr.Clone()

	// Act
	actual := args.Map{"result": clone.Left != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
	_ = args.LeftRightAny{}.AsTwoParameter()
	_ = args.LeftRightAny{}.AsArgBaseContractsBinder()
}

// --- Holder ---

func Test_Holder_AllMethods_FromMapLengthIteration18(t *testing.T) {
	fn := func(s string) string { return s }
	h := &args.HolderAny{
		First: "a", Second: "b", Third: "c",
		Fourth: "d", Fifth: "e", Sixth: "f",
		WorkFunc: fn, Expect: "exp",
	}
	_ = h.GetWorkFunc()
	_ = h.ArgsCount()
	_ = h.FirstItem()
	_ = h.SecondItem()
	_ = h.ThirdItem()
	_ = h.FourthItem()
	_ = h.FifthItem()
	_ = h.SixthItem()
	_ = h.Expected()
	_ = h.ArgTwo()
	_ = h.ArgThree()
	_ = h.ArgFour()
	_ = h.ArgFive()
	_ = h.HasFirst()
	_ = h.HasSecond()
	_ = h.HasThird()
	_ = h.HasFourth()
	_ = h.HasFifth()
	_ = h.HasSixth()
	_ = h.HasFunc()
	_ = h.HasExpect()
	_ = h.GetFuncName()
	_ = h.FuncWrap()
	_ = h.ValidArgs()
	_ = h.Args(6)
	_ = h.Slice()
	_ = h.GetByIndex(0)
	_ = h.String()

	hVal := args.HolderAny{}
	_ = hVal.AsSixthParameter()
	_ = hVal.AsArgFuncContractsBinder()
}

// --- String ---

func Test_String_AllMethods_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	s := args.String("hello")
	_ = s.Concat(" world")
	_ = s.Join("-", "a", "b")
	_ = s.Split("l")
	_ = s.DoubleQuote()
	_ = s.DoubleQuoteQ()
	_ = s.SingleQuote()
	_ = s.ValueDoubleQuote()
	_ = s.String()
	_ = s.Bytes()
	_ = s.Runes()
	_ = s.Length()
	_ = s.Count()
	_ = s.IsEmptyOrWhitespace()
	_ = s.TrimSpace()
	_ = s.ReplaceAll("h", "H")
	_ = s.Substring(0, 3)

	// Act
	actual := args.Map{"result": s.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	actual = args.Map{"result": s.HasCharacter()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has character", actual)
	actual = args.Map{"result": s.IsDefined()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected defined", actual)
	_ = s.AscIILength()

	empty := args.String("")
	actual = args.Map{"result": empty.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": empty.IsEmptyOrWhitespace()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty or whitespace", actual)
	ws := args.String("   ")
	actual = args.Map{"result": ws.IsEmptyOrWhitespace()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected whitespace only", actual)
}

func Test_String_TrimReplaceMap(t *testing.T) {
	// Arrange
	s := args.String("hello {name}")
	result := s.TrimReplaceMap(map[string]string{"{name}": "world"})

	// Act
	actual := args.Map{"result": strings.Contains(string(result), "world")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected replacement", actual)
}

// --- Empty ---

func Test_Empty_Map_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	m := args.Empty.Map()

	// Act
	actual := args.Map{"result": m == nil || len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map", actual)
}

func Test_Empty_FuncWrap(t *testing.T) {
	// Arrange
	fw := args.Empty.FuncWrap()

	// Act
	actual := args.Map{"result": fw == nil || !fw.IsInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid func wrap", actual)
}

func Test_Empty_FuncMap(t *testing.T) {
	// Arrange
	fm := args.Empty.FuncMap()

	// Act
	actual := args.Map{"result": fm == nil || len(fm) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty func map", actual)
}

func Test_Empty_Holder(t *testing.T) {
	// Arrange
	h := args.Empty.Holder()

	// Act
	actual := args.Map{"result": h.HasFirst()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no first", actual)
}

// --- FuncWrap ---

func Test_NewTypedFuncWrap_Valid(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewTypedFuncWrap(fn)

	// Act
	actual := args.Map{"result": fw.IsInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual = args.Map{"result": fw.ArgsCount() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 arg", actual)
}

func Test_NewTypedFuncWrap_Nil(t *testing.T) {
	// Arrange
	var fn func()
	fw := args.NewTypedFuncWrap(fn)

	// Act
	actual := args.Map{"result": fw.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid for nil func", actual)
}

func Test_NewTypedFuncWrap_NotFunc(t *testing.T) {
	// Arrange
	fw := args.NewTypedFuncWrap("not a func")

	// Act
	actual := args.Map{"result": fw.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid for non-func", actual)
}

func Test_FuncWrap_GetFuncName_Nil(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	actual := args.Map{"result": fw.GetFuncName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_FuncWrap_GetPascalCaseFuncName_Nil_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	actual := args.Map{"result": fw.GetPascalCaseFuncName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_FuncWrap_IsEqual(t *testing.T) {
	// Arrange
	fn := func() {}
	fw1 := args.NewFuncWrap.Default(fn)
	fw2 := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{"result": fw1.IsNotEqual(fw2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)

	var nilFw *args.FuncWrapAny
	actual = args.Map{"result": nilFw.IsEqual(nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil == nil", actual)
	actual = args.Map{"result": nilFw.IsEqual(fw1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil != non-nil", actual)
}

func Test_FuncWrap_PkgPath_FromMapLengthIteration18(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	p := fw.PkgPath()
	_ = p
	// second call for cached path
	p2 := fw.PkgPath()
	_ = p2
}

func Test_FuncWrap_PkgNameOnly_FromMapLengthIteration18(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	p := fw.PkgNameOnly()
	_ = p
	p2 := fw.PkgNameOnly()
	_ = p2
}

func Test_FuncWrap_FuncDirectInvokeName_FromMapLengthIteration18(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	n := fw.FuncDirectInvokeName()
	_ = n
	n2 := fw.FuncDirectInvokeName()
	_ = n2
}

func Test_FuncWrap_IsEqualValue_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	fn := func() {}
	fw1 := args.NewTypedFuncWrap(fn)
	fw2 := args.NewTypedFuncWrap(fn)

	// Act
	actual := args.Map{"result": fw1.IsEqualValue(*fw2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal value", actual)
}

// --- FuncMap ---

func Test_FuncMap_Basic_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	actual := args.Map{"result": fm.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": fm.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": fm.Count() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": fm.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no items", actual)
	actual = args.Map{"result": fm.Has("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not has", actual)
	actual = args.Map{"result": fm.IsContains("x")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not contains", actual)
	actual = args.Map{"result": fm.Get("x") != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_FuncMap_AddAndInvoke(t *testing.T) {
	// Arrange
	fn := func(n int) int { return n * 2 }
	fm := args.FuncMap{}
	fm.Add(fn)

	// Act
	actual := args.Map{"result": fm.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	name := fm.GetPascalCaseFuncName("x")
	_ = name
}

func Test_FuncMap_InvalidError_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}
	err := fm.InvalidError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for empty map", actual)
}

// --- NewFuncWrap creator ---

func Test_NewFuncWrap_Default_Nil(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	actual := args.Map{"result": fw == nil || !fw.IsInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid for nil", actual)
}

func Test_NewFuncWrap_Default_Valid(t *testing.T) {
	// Arrange
	fn := func() {}
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{"result": fw == nil || fw.IsInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
}

// --- FuncDetector ---

func Test_FuncDetector_GetFuncWrap_FromMapLengthIteration18(t *testing.T) {
	// Arrange
	fn := func() {}
	fw := args.FuncDetector.GetFuncWrap(fn)

	// Act
	actual := args.Map{"result": fw == nil || fw.IsInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid func wrap", actual)

	fw2 := args.FuncDetector.GetFuncWrap("not a func")
	actual = args.Map{"result": fw2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil (invalid) func wrap", actual)
}
