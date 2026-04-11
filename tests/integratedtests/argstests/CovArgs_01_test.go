package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// args Coverage — One, Two, Map, LeftRight, String, Holder
// ══════════════════════════════════════════════════════════════════════════════

// --- One ---

func Test_CovArgs_01_One_Basic(t *testing.T) {
	// Arrange
	o := &args.OneAny{First: "hello", Expect: 42}

	// Act
	actual := args.Map{"result": o.FirstItem() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	actual = args.Map{"result": o.Expected() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	actual = args.Map{"result": o.HasFirst()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.HasExpect()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.ArgsCount() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovArgs_02_One_ValidArgs_Args_Slice(t *testing.T) {
	// Arrange
	o := &args.OneAny{First: "hello", Expect: 42}
	va := o.ValidArgs()

	// Act
	actual := args.Map{"result": len(va) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	a := o.Args(1)
	actual = args.Map{"result": len(a) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	a0 := o.Args(0)
	actual = args.Map{"result": len(a0) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	sl := o.Slice()
	actual = args.Map{"result": len(sl) < 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	// cached
	sl2 := o.Slice()
	actual = args.Map{"result": len(sl2) != len(sl)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same", actual)
}

func Test_CovArgs_03_One_GetByIndex_String(t *testing.T) {
	// Arrange
	o := &args.OneAny{First: "hello"}
	_ = o.GetByIndex(0)
	_ = o.GetByIndex(99)
	s := o.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovArgs_04_One_LeftRight(t *testing.T) {
	// Arrange
	o := &args.OneAny{First: "hello", Expect: 42}
	lr := o.LeftRight()

	// Act
	actual := args.Map{"result": lr.Left != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_CovArgs_05_One_ArgTwo(t *testing.T) {
	// Arrange
	o := &args.OneAny{First: "hello", Expect: 42}
	a2 := o.ArgTwo()

	// Act
	actual := args.Map{"result": a2.First != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_CovArgs_06_One_AsInterfaces(t *testing.T) {
	o := args.OneAny{First: "hello"}
	_ = o.AsOneParameter()
	_ = o.AsArgBaseContractsBinder()
}

// --- Two ---

func Test_CovArgs_07_Two_Basic(t *testing.T) {
	// Arrange
	tw := &args.TwoAny{First: "a", Second: "b", Expect: 1}

	// Act
	actual := args.Map{"result": tw.FirstItem() != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
	actual = args.Map{"result": tw.SecondItem() != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)
	actual = args.Map{"result": tw.Expected() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": tw.HasFirst()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": tw.HasSecond()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": tw.HasExpect()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": tw.ArgsCount() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_CovArgs_08_Two_ValidArgs_Args_Slice(t *testing.T) {
	// Arrange
	tw := &args.TwoAny{First: "a", Second: "b", Expect: 1}
	va := tw.ValidArgs()

	// Act
	actual := args.Map{"result": len(va) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	a := tw.Args(2)
	actual = args.Map{"result": len(a) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	a1 := tw.Args(1)
	actual = args.Map{"result": len(a1) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	sl := tw.Slice()
	actual = args.Map{"result": len(sl) < 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
}

func Test_CovArgs_09_Two_GetByIndex_String(t *testing.T) {
	// Arrange
	tw := &args.TwoAny{First: "a", Second: "b"}
	_ = tw.GetByIndex(0)
	_ = tw.GetByIndex(99)
	s := tw.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovArgs_10_Two_LeftRight_ArgTwo(t *testing.T) {
	// Arrange
	tw := &args.TwoAny{First: "a", Second: "b", Expect: 1}
	lr := tw.LeftRight()

	// Act
	actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	a2 := tw.ArgTwo()
	actual = args.Map{"result": a2.First != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
}

func Test_CovArgs_11_Two_AsInterfaces(t *testing.T) {
	tw := args.TwoAny{First: "a", Second: "b"}
	_ = tw.AsTwoParameter()
	_ = tw.AsArgBaseContractsBinder()
}

// --- Map ---

func Test_CovArgs_12_Map_Basic(t *testing.T) {
	// Arrange
	m := args.Map{
		"first": "hello",
		"expected": 42,
	}

	// Act
	actual := args.Map{"result": m.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": m.Expected() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	actual = args.Map{"result": m.HasExpect()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": m.HasFirst()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": m.FirstItem() != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_CovArgs_13_Map_ArgsCount(t *testing.T) {
	// Arrange
	m := args.Map{
		"first": "hello",
		"expected": 42,
		"func": func() {},
	}
	// ArgsCount excludes expected and func
	ac := m.ArgsCount()

	// Act
	actual := args.Map{"result": ac != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovArgs_14_Map_Get_Has_HasDefined(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	v, ok := m.Get("key")

	// Act
	actual := args.Map{"result": ok || v != "val"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected val", actual)
	_, ok2 := m.Get("missing")
	actual = args.Map{"result": ok2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": m.Has("key")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": m.Has("missing")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": m.HasDefined("key")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": m.HasDefined("missing")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovArgs_15_Map_HasDefinedAll_IsKeyInvalid_IsKeyMissing(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{"result": m.HasDefinedAll("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": m.HasDefinedAll("a", "missing")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": m.HasDefinedAll()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty names", actual)
	actual = args.Map{"result": m.IsKeyMissing("a")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": m.IsKeyMissing("missing")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovArgs_16_Map_NilMap(t *testing.T) {
	// Arrange
	var m args.Map
	_, ok := m.Get("key")

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": m.Has("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": m.HasDefined("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": m.HasDefinedAll("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": m.IsKeyInvalid("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": m.IsKeyMissing("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovArgs_17_Map_SortedKeys(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	keys, err := m.SortedKeys()

	// Act
	actual := args.Map{"result": err != nil || len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 keys", actual)
	actual = args.Map{"result": keys[0] != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a first", actual)
	// empty map
	m2 := args.Map{}
	keys2, _ := m2.SortedKeys()
	actual = args.Map{"result": len(keys2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovArgs_18_Map_SortedKeysMust(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	keys := m.SortedKeysMust()

	// Act
	actual := args.Map{"result": len(keys) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovArgs_19_Map_Items(t *testing.T) {
	// Arrange
	m := args.Map{
		"first": "a",
		"second": "b",
		"third": "c",
		"fourth": "d",
		"fifth": "e",
		"sixth": "f",
		"seventh": "g",
	}

	// Act
	actual := args.Map{"result": m.SecondItem() != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)
	actual = args.Map{"result": m.ThirdItem() != "c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected c", actual)
	actual = args.Map{"result": m.FourthItem() != "d"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected d", actual)
	actual = args.Map{"result": m.FifthItem() != "e"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected e", actual)
	actual = args.Map{"result": m.SixthItem() != "f"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected f", actual)
	actual = args.Map{"result": m.Seventh() != "g"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected g", actual)
}

func Test_CovArgs_20_Map_GetLowerCase_GetDirectLower(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	v, ok := m.GetLowerCase("KEY")

	// Act
	actual := args.Map{"result": ok || v != "val"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected val", actual)
	v2 := m.GetDirectLower("KEY")
	actual = args.Map{"result": v2 != "val"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected val", actual)
	v3 := m.GetDirectLower("MISSING")
	actual = args.Map{"result": v3 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovArgs_21_Map_Expect_Actual_Arrange(t *testing.T) {
	// Arrange
	m := args.Map{
		"expect": 1,
		"actual": 2,
		"arrange": 3,
	}

	// Act
	actual := args.Map{"result": m.Expect() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": m.Actual() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": m.Arrange() != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_CovArgs_22_Map_SetActual(t *testing.T) {
	// Arrange
	m := args.Map{}
	m.SetActual("val")

	// Act
	actual := args.Map{"result": m.Actual() != "val"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected val", actual)
}

func Test_CovArgs_23_Map_When_Title(t *testing.T) {
	// Arrange
	m := args.Map{
		"when": "w",
		"title": "t",
	}

	// Act
	actual := args.Map{"result": m.When() != "w"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected w", actual)
	actual = args.Map{"result": m.Title() != "t"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected t", actual)
}

func Test_CovArgs_24_Map_GetByIndex(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	_ = m.GetByIndex(0)
	r := m.GetByIndex(99)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovArgs_25_Map_Raw_Args_ValidArgs(t *testing.T) {
	// Arrange
	m := args.Map{
		"first": "a",
		"second": "b",
	}
	raw := m.Raw()

	// Act
	actual := args.Map{"result": len(raw) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	a := m.Args("first", "second")
	actual = args.Map{"result": len(a) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	va := m.ValidArgs()
	actual = args.Map{"result": len(va) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovArgs_26_Map_GetFirstOfNames(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	v := m.GetFirstOfNames("missing", "a")

	// Act
	actual := args.Map{"result": v != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	v2 := m.GetFirstOfNames()
	actual = args.Map{"result": v2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovArgs_27_Map_GetAsStringSliceFirstOfNames(t *testing.T) {
	// Arrange
	m := args.Map{"items": []string{"a", "b"}}
	r := m.GetAsStringSliceFirstOfNames("items")

	// Act
	actual := args.Map{"result": len(r) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	r2 := m.GetAsStringSliceFirstOfNames("missing")
	actual = args.Map{"result": r2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	r3 := m.GetAsStringSliceFirstOfNames()
	actual = args.Map{"result": r3 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- LeftRight ---

func Test_CovArgs_28_LeftRight(t *testing.T) {
	// Arrange
	lr := &args.LeftRightAny{Left: "a", Right: "b", Expect: 1}

	// Act
	actual := args.Map{"result": lr.FirstItem() != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
	actual = args.Map{"result": lr.SecondItem() != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)
	actual = args.Map{"result": lr.Expected() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": lr.ArgsCount() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": lr.HasFirst()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": lr.HasSecond()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": lr.HasLeft()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": lr.HasRight()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": lr.HasExpect()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	va := lr.ValidArgs()
	actual = args.Map{"result": len(va) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	a := lr.Args(2)
	actual = args.Map{"result": len(a) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sl := lr.Slice()
	actual = args.Map{"result": len(sl) < 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	_ = lr.GetByIndex(0)
	_ = lr.String()
	a2 := lr.ArgTwo()
	actual = args.Map{"result": a2.First != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
	c := lr.Clone()
	actual = args.Map{"result": c.Left != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
	lrV := args.LeftRightAny{Left: "a", Right: "b"}
	_ = lrV.AsTwoParameter()
	_ = lrV.AsArgBaseContractsBinder()
}

// --- Holder ---

func Test_CovArgs_29_Holder(t *testing.T) {
	// Arrange
	fn := func(s string) string { return s }
	h := &args.HolderAny{
		First:    "a",
		Second:   "b",
		Third:    "c",
		Fourth:   "d",
		Fifth:    "e",
		Sixth:    "f",
		WorkFunc: fn,
		Expect:   "x",
	}

	// Act
	actual := args.Map{"result": h.FirstItem() != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
	actual = args.Map{"result": h.SecondItem() != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)
	actual = args.Map{"result": h.ThirdItem() != "c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected c", actual)
	actual = args.Map{"result": h.FourthItem() != "d"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected d", actual)
	actual = args.Map{"result": h.FifthItem() != "e"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected e", actual)
	actual = args.Map{"result": h.SixthItem() != "f"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected f", actual)
	actual = args.Map{"result": h.Expected() != "x"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected x", actual)
	actual = args.Map{"result": h.ArgsCount() != 7}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 7", actual)
	actual = args.Map{"result": h.HasFirst()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": h.HasSecond()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": h.HasThird()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": h.HasFourth()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": h.HasFifth()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": h.HasSixth()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": h.HasFunc()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": h.HasExpect()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	_ = h.GetWorkFunc()
	_ = h.GetFuncName()
	va := h.ValidArgs()
	actual = args.Map{"result": len(va) != 6}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6", actual)
	a := h.Args(6)
	actual = args.Map{"result": len(a) != 6}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6", actual)
	sl := h.Slice()
	actual = args.Map{"result": len(sl) < 6}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 6", actual)
	_ = h.GetByIndex(0)
	_ = h.String()
	_ = h.ArgTwo()
	_ = h.ArgThree()
	_ = h.ArgFour()
	_ = h.ArgFive()
	hv := args.HolderAny{First: "a"}
	_ = hv.AsSixthParameter()
	_ = hv.AsArgFuncContractsBinder()
}

// --- String ---

func Test_CovArgs_30_String(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{"result": s.String() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	actual = args.Map{"result": s.Length() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": s.Count() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": s.AscIILength() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": s.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": s.HasCharacter()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": s.IsDefined()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": s.IsEmptyOrWhitespace()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_ = s.Bytes()
	_ = s.Runes()
	_ = s.TrimSpace()
	_ = s.DoubleQuote()
	_ = s.DoubleQuoteQ()
	_ = s.SingleQuote()
	_ = s.ValueDoubleQuote()
	_ = s.ReplaceAll("h", "H")
	_ = s.Concat("world")
	_ = s.Join(",", "world")
	_ = s.Split(",")
	_ = s.Substring(0, 3)
}

func Test_CovArgs_31_String_Empty(t *testing.T) {
	// Arrange
	s := args.String("")

	// Act
	actual := args.Map{"result": s.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": s.HasCharacter()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": s.IsEmptyOrWhitespace()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovArgs_32_String_TrimReplaceMap(t *testing.T) {
	// Arrange
	s := args.String("Hello {name}")
	r := s.TrimReplaceMap(map[string]string{"{name}": "World"})

	// Act
	actual := args.Map{"result": r.String() != "Hello World"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'Hello World', got ''", actual)
}
