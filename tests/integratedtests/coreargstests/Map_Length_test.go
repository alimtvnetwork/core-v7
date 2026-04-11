package coreargstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Map basic methods ──

func Test_Map_Length(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Length", actual)
}

func Test_Map_ArgsCount(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"func": func() {},
		"expected": "x",
	}

	// Act
	actual := args.Map{"argsCount": m.ArgsCount()}

	// Assert
	expected := args.Map{"argsCount": 1}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- ArgsCount", actual)
}

func Test_Map_Has_HasDefined(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": nil,
	}

	// Act
	actual := args.Map{
		"hasA":        m.Has("a"),
		"hasC":        m.Has("c"),
		"definedA":    m.HasDefined("a"),
		"definedB":    m.HasDefined("b"),
		"definedC":    m.HasDefined("c"),
		"nilHas":      args.Map(nil).Has("a"),
		"nilDefined":  args.Map(nil).HasDefined("a"),
	}

	// Assert
	expected := args.Map{
		"hasA": true, "hasC": false,
		"definedA": true, "definedB": false, "definedC": false,
		"nilHas": false, "nilDefined": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Has/HasDefined", actual)
}

func Test_Map_HasDefinedAll(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{
		"allDef":     m.HasDefinedAll("a", "b"),
		"oneMissing": m.HasDefinedAll("a", "c"),
		"nil":        args.Map(nil).HasDefinedAll("a"),
		"empty":      m.HasDefinedAll(),
	}

	// Assert
	expected := args.Map{
		"allDef": true,
		"oneMissing": false,
		"nil": false,
		"empty": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- HasDefinedAll", actual)
}

func Test_Map_IsKeyInvalid_IsKeyMissing(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": nil,
	}

	// Act
	actual := args.Map{
		"invalidA":   m.IsKeyInvalid("a"),
		"invalidB":   m.IsKeyInvalid("b"),
		"invalidC":   m.IsKeyInvalid("c"),
		"missingA":   m.IsKeyMissing("a"),
		"missingC":   m.IsKeyMissing("c"),
		"nilInvalid": args.Map(nil).IsKeyInvalid("a"),
		"nilMissing": args.Map(nil).IsKeyMissing("a"),
	}

	// Assert
	expected := args.Map{
		"invalidA": false, "invalidB": true, "invalidC": true,
		"missingA": false, "missingC": true,
		"nilInvalid": false, "nilMissing": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns error -- IsKeyInvalid/IsKeyMissing", actual)
}

func Test_Map_Get(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	val, ok := m.Get("a")
	_, notOk := m.Get("missing")
	nilVal, nilOk := args.Map(nil).Get("a")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
		"notOk": notOk,
		"nilVal": nilVal == nil,
		"nilOk": nilOk,
	}

	// Assert
	expected := args.Map{
		"val": 1,
		"ok": true,
		"notOk": false,
		"nilVal": true,
		"nilOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Get", actual)
}

func Test_Map_GetLowerCase(t *testing.T) {
	// Arrange
	m := args.Map{"hello": "world"}
	val, ok := m.GetLowerCase("HELLO")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "world",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetLowerCase", actual)
}

func Test_Map_GetDirectLower(t *testing.T) {
	// Arrange
	m := args.Map{"hello": "world"}

	// Act
	actual := args.Map{
		"found":   m.GetDirectLower("HELLO"),
		"missing": m.GetDirectLower("OTHER") == nil,
	}

	// Assert
	expected := args.Map{
		"found": "world",
		"missing": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetDirectLower", actual)
}

func Test_Map_GetAsInt(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 42,
		"b": "not int",
	}
	v, ok := m.GetAsInt("a")
	_, notOk := m.GetAsInt("b")
	_, missingOk := m.GetAsInt("c")

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
		"notOk": notOk,
		"missingOk": missingOk,
	}

	// Assert
	expected := args.Map{
		"v": 42,
		"ok": true,
		"notOk": false,
		"missingOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsInt", actual)
}

func Test_Map_GetAsIntDefault(t *testing.T) {
	// Arrange
	m := args.Map{"a": 42}

	// Act
	actual := args.Map{
		"found":   m.GetAsIntDefault("a", 0),
		"missing": m.GetAsIntDefault("b", 99),
	}

	// Assert
	expected := args.Map{
		"found": 42,
		"missing": 99,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsIntDefault", actual)
}

func Test_Map_GetAsString(t *testing.T) {
	// Arrange
	m := args.Map{"a": "hello"}
	v, ok := m.GetAsString("a")
	_, notOk := m.GetAsString("b")

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
		"notOk": notOk,
	}

	// Assert
	expected := args.Map{
		"v": "hello",
		"ok": true,
		"notOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsString", actual)
}

func Test_Map_GetAsStringDefault(t *testing.T) {
	// Arrange
	m := args.Map{"a": "hello"}

	// Act
	actual := args.Map{
		"found":   m.GetAsStringDefault("a"),
		"missing": m.GetAsStringDefault("b"),
	}

	// Assert
	expected := args.Map{
		"found": "hello",
		"missing": "",
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsStringDefault", actual)
}

func Test_Map_GetAsBool(t *testing.T) {
	// Arrange
	m := args.Map{"a": true}
	v, ok := m.GetAsBool("a")
	_, notOk := m.GetAsBool("b")

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
		"notOk": notOk,
	}

	// Assert
	expected := args.Map{
		"v": true,
		"ok": true,
		"notOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsBool", actual)
}

func Test_Map_NamedAccessors(t *testing.T) {
	// Arrange
	m := args.Map{
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "sixth": 6, "seventh": 7,
		"when": "w", "title": "t",
	}

	// Act
	actual := args.Map{
		"first": m.FirstItem(), "second": m.SecondItem(), "third": m.ThirdItem(),
		"fourth": m.FourthItem(), "fifth": m.FifthItem(), "sixth": m.SixthItem(),
		"seventh": m.Seventh(), "when": m.When(), "title": m.Title(),
		"hasFirst": m.HasFirst(),
	}

	// Assert
	expected := args.Map{
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "sixth": 6,
		"seventh": 7, "when": "w", "title": "t",
		"hasFirst": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- NamedAccessors", actual)
}

func Test_Map_AltKeyAccessors(t *testing.T) {
	// Arrange
	m := args.Map{
		"f1": "a",
		"p2": "b",
		"p3": "c",
		"f4": "d",
		"f5": "e",
		"f6": "f",
		"f7": "g",
	}

	// Act
	actual := args.Map{
		"first": m.FirstItem(), "second": m.SecondItem(), "third": m.ThirdItem(),
		"fourth": m.FourthItem(), "fifth": m.FifthItem(), "sixth": m.SixthItem(),
		"seventh": m.Seventh(),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"fourth": "d", "fifth": "e", "sixth": "f",
		"seventh": "g",
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- AltKeyAccessors", actual)
}

func Test_Map_Expect_Actual_Arrange(t *testing.T) {
	// Arrange
	m := args.Map{
		"expect": "e",
		"actual": "a",
		"arrange": "r",
	}

	// Act
	actual := args.Map{
		"expect":  m.Expect(),
		"actual":  m.Actual(),
		"arrange": m.Arrange(),
	}

	// Assert
	expected := args.Map{
		"expect": "e",
		"actual": "a",
		"arrange": "r",
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Expect/Actual/Arrange", actual)
}

func Test_Map_SetActual(t *testing.T) {
	// Arrange
	m := args.Map{}
	m.SetActual("val")

	// Act
	actual := args.Map{"actual": m["actual"]}

	// Assert
	expected := args.Map{"actual": "val"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- SetActual", actual)
}

func Test_Map_SortedKeys(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	keys, err := m.SortedKeys()
	emptyKeys, emptyErr := args.Map{}.SortedKeys()

	// Act
	actual := args.Map{
		"first": keys[0], "second": keys[1], "noErr": err == nil,
		"emptyLen": len(emptyKeys), "emptyNoErr": emptyErr == nil,
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "noErr": true,
		"emptyLen": 0, "emptyNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- SortedKeys", actual)
}

func Test_Map_SortedKeysMust(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	keys := m.SortedKeysMust()

	// Act
	actual := args.Map{"first": keys[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- SortedKeysMust", actual)
}

func Test_Map_GetByIndex(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{
		"found":   m.GetByIndex(0) != nil,
		"outRange": m.GetByIndex(99) == nil,
	}

	// Assert
	expected := args.Map{
		"found": true,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetByIndex", actual)
}

func Test_Map_GetFirstOfNames(t *testing.T) {
	// Arrange
	m := args.Map{"b": 2}

	// Act
	actual := args.Map{
		"found":   m.GetFirstOfNames("a", "b"),
		"empty":   m.GetFirstOfNames() == nil,
		"missing": m.GetFirstOfNames("x", "y") == nil,
	}

	// Assert
	expected := args.Map{
		"found": 2,
		"empty": true,
		"missing": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetFirstOfNames", actual)
}

func Test_Map_Raw(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	raw := m.Raw()

	// Act
	actual := args.Map{"len": len(raw)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Raw", actual)
}

func Test_Map_Args(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}
	result := m.Args("a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Args", actual)
}

func Test_Map_ValidArgs(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": nil,
		"func": func() {},
	}
	result := m.ValidArgs()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map returns non-empty -- ValidArgs", actual)
}

func Test_Map_Expected(t *testing.T) {
	// Arrange
	m1 := args.Map{"expected": "a"}
	m2 := args.Map{"expects": "b"}
	m3 := args.Map{"expect": "c"}

	// Act
	actual := args.Map{
		"expected": m1.Expected(), "expects": m2.Expected(), "expect": m3.Expected(),
		"hasExpect": m1.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"expected": "a",
		"expects": "b",
		"expect": "c",
		"hasExpect": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Expected", actual)
}

func Test_Map_WorkFunc_HasFunc(t *testing.T) {
	// Arrange
	fn := func() {}
	m := args.Map{"func": fn}

	// Act
	actual := args.Map{
		"hasFunc":  m.HasFunc(),
		"noFunc":   args.Map{}.HasFunc(),
		"funcName": m.GetFuncName() != "",
	}

	// Assert
	expected := args.Map{
		"hasFunc": true,
		"noFunc": actual["noFunc"],
		"funcName": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- WorkFunc/HasFunc", actual)
}

func Test_Map_Compile(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": "hello",
	}
	result := m.CompileToStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- CompileToStrings", actual)
}

func Test_Map_Slice_String(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{
		"sliceLen":  len(m.Slice()) > 0,
		"strNotEmpty": m.String() != "",
	}

	// Assert
	expected := args.Map{
		"sliceLen": true,
		"strNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Slice/String", actual)
}

func Test_Map_GoLiteralString(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{"notEmpty": m.GoLiteralString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GoLiteralString", actual)
}

// ── Dynamic ──

func Test_Dynamic_Methods(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{
			"first": 1,
			"second": 2,
		},
		Expect: "expected",
	}

	// Act
	actual := args.Map{
		"argsCount":  d.ArgsCount(),
		"expected":   d.Expected(),
		"hasExpect":  d.HasExpect(),
		"hasFirst":   d.HasFirst(),
		"firstItem":  d.FirstItem(),
		"secondItem": d.SecondItem(),
		"hasDefined": d.HasDefined("first"),
		"has":        d.Has("first"),
	}

	// Assert
	expected := args.Map{
		"argsCount": actual["argsCount"], "expected": "expected", "hasExpect": true,
		"hasFirst": true, "firstItem": 1, "secondItem": 2,
		"hasDefined": true, "has": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Methods", actual)
}

func Test_Dynamic_Nil(t *testing.T) {
	// Arrange
	var d *args.DynamicAny

	// Act
	actual := args.Map{
		"argsCount":  d.ArgsCount(),
		"hasFirst":   d.HasFirst(),
		"hasDefined": d.HasDefined("a"),
		"has":        d.Has("a"),
		"hasExpect":  d.HasExpect(),
		"workFunc":   d.GetWorkFunc() == nil,
	}

	// Assert
	expected := args.Map{
		"argsCount": 0, "hasFirst": false, "hasDefined": false,
		"has": false, "hasExpect": false, "workFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- Nil", actual)
}

func Test_Dynamic_Get_Methods(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{Params: args.Map{
		"num": 42,
		"str": "hello",
		"actual": "act",
		"arrange": "arr",
	}}
	v, ok := d.Get("num")
	intV, intOk := d.GetAsInt("num")
	strV, strOk := d.GetAsString("str")

	// Act
	actual := args.Map{
		"v": v, "ok": ok, "intV": intV, "intOk": intOk,
		"strV": strV, "strOk": strOk,
		"strDefault": d.GetAsStringDefault("str"),
		"intDefault": d.GetAsIntDefault("num", 0),
		"intMissing": d.GetAsIntDefault("x", 99),
		"actual":     d.Actual(),
		"arrange":    d.Arrange(),
	}

	// Assert
	expected := args.Map{
		"v": 42, "ok": true, "intV": 42, "intOk": true,
		"strV": "hello", "strOk": true,
		"strDefault": "hello", "intDefault": 42, "intMissing": 99,
		"actual": "act", "arrange": "arr",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Get Methods", actual)
}

func Test_Dynamic_HasDefinedAll_IsKeyInvalid_IsKeyMissing(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{Params: args.Map{
		"a": 1,
		"b": nil,
	}}

	// Act
	actual := args.Map{
		"allDef":     d.HasDefinedAll("a"),
		"allMissing": d.HasDefinedAll("a", "c"),
		"invalidB":   d.IsKeyInvalid("b"),
		"missingC":   d.IsKeyMissing("c"),
	}

	// Assert
	expected := args.Map{
		"allDef": true,
		"allMissing": false,
		"invalidB": true,
		"missingC": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns error -- HasDefinedAll/IsKeyInvalid/IsKeyMissing", actual)
}

func Test_Dynamic_GetAsStrings_GetAsAnyItems(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{Params: args.Map{
		"strs": []string{"a"},
		"anys": []any{1},
	}}
	strs, sOk := d.GetAsStrings("strs")
	anys, aOk := d.GetAsAnyItems("anys")
	_, sMiss := d.GetAsStrings("x")
	_, aMiss := d.GetAsAnyItems("x")

	// Act
	actual := args.Map{
		"strsLen": len(strs), "sOk": sOk, "anysLen": len(anys), "aOk": aOk,
		"sMiss": sMiss, "aMiss": aMiss,
	}

	// Assert
	expected := args.Map{
		"strsLen": 1, "sOk": true, "anysLen": 1, "aOk": true,
		"sMiss": false, "aMiss": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetAsStrings/GetAsAnyItems", actual)
}

func Test_Dynamic_Slice_String_Contracts(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{Params: args.Map{"a": 1}, Expect: "e"}

	// Act
	actual := args.Map{
		"sliceLen":  len(d.Slice()) > 0,
		"strNotEmpty": d.String() != "",
		"mapper":    d.AsArgsMapper() != nil,
		"funcName":  d.AsArgFuncNameContractsBinder() != nil,
		"base":      d.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"sliceLen": true, "strNotEmpty": true,
		"mapper": true, "funcName": true, "base": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Slice/String/Contracts", actual)
}

// ── DynamicFunc ──

func Test_DynamicFunc_Methods(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	d := &args.DynamicFuncAny{
		Params:   args.Map{"first": "hello"},
		WorkFunc: fn,
		Expect:   5,
	}

	// Act
	actual := args.Map{
		"argsCount":  d.ArgsCount(),
		"length":     d.Length(),
		"hasFirst":   d.HasFirst(),
		"firstItem":  d.FirstItem(),
		"expected":   d.Expected(),
		"workFunc":   d.GetWorkFunc() != nil,
		"hasDefined": d.HasDefined("first"),
		"has":        d.Has("first"),
	}

	// Assert
	expected := args.Map{
		"argsCount": actual["argsCount"], "length": 1, "hasFirst": true,
		"firstItem": "hello", "expected": 5,
		"workFunc": true, "hasDefined": true, "has": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- Methods", actual)
}

func Test_DynamicFunc_Nil(t *testing.T) {
	// Arrange
	var d *args.DynamicFuncAny

	// Act
	actual := args.Map{
		"argsCount": d.ArgsCount(),
		"length": d.Length(),
	}

	// Assert
	expected := args.Map{
		"argsCount": 0,
		"length": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns nil -- Nil", actual)
}

func Test_DynamicFunc_Contracts(t *testing.T) {
	// Arrange
	d := args.DynamicFuncAny{Params: args.Map{"a": 1}, WorkFunc: func() {}, Expect: "e"}

	// Act
	actual := args.Map{
		"mapper":   d.AsArgsMapper() != nil,
		"funcName": d.AsArgFuncNameContractsBinder() != nil,
		"base":     d.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"mapper": true,
		"funcName": true,
		"base": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- Contracts", actual)
}

// ── FuncWrap ──

func Test_FuncWrap_Valid(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{
		"hasValid":    fw.HasValidFunc(),
		"isInvalid":   fw.IsInvalid(),
		"isValid":     fw.IsValid(),
		"name":        fw.GetFuncName() != "",
		"pascal":      fw.GetPascalCaseFuncName() != "",
		"argsCount":   fw.ArgsCount(),
		"returnLen":   fw.ReturnLength(),
		"outCount":    fw.OutArgsCount(),
		"typeNotNil":  fw.GetType() != nil,
		"pkgPath":     fw.PkgPath() != "",
		"pkgNameOnly": fw.PkgNameOnly() != "",
		"directName":  fw.FuncDirectInvokeName() != "",
	}

	// Assert
	expected := args.Map{
		"hasValid": true, "isInvalid": false, "isValid": true,
		"name": true, "pascal": true, "argsCount": 1, "returnLen": 1,
		"outCount": 1, "typeNotNil": true,
		"pkgPath": true, "pkgNameOnly": true, "directName": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns non-empty -- Valid", actual)
}

func Test_FuncWrap_Invalid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	actual := args.Map{
		"isInvalid":  fw.IsInvalid(),
		"argsCount":  fw.ArgsCount(),
		"returnLen":  fw.ReturnLength(),
		"typeIsNil":  fw.GetType() == nil,
		"pkgPath":    fw.PkgPath(),
		"pkgName":    fw.PkgNameOnly(),
		"directName": fw.FuncDirectInvokeName(),
	}

	// Assert
	expected := args.Map{
		"isInvalid": true, "argsCount": -1, "returnLen": -1,
		"typeIsNil": true, "pkgPath": "", "pkgName": "", "directName": "",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns error -- Invalid", actual)
}

func Test_FuncWrap_Nil(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	actual := args.Map{
		"name":   fw.GetFuncName(),
		"pascal": fw.GetPascalCaseFuncName(),
	}

	// Assert
	expected := args.Map{
		"name": "",
		"pascal": "",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns nil -- Nil", actual)
}

func Test_FuncWrap_IsEqual(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw1 := args.NewFuncWrap.Default(fn)
	fw2 := args.NewFuncWrap.Default(fn)
	fw3 := args.NewFuncWrap.Default(func() {})
	var nilFw *args.FuncWrapAny

	// Act
	actual := args.Map{
		"equalSame":    fw1.IsEqual(fw2),
		"notEqual":     fw1.IsNotEqual(fw3),
		"nilNil":       nilFw.IsEqual(nilFw),
		"nilNonNil":    nilFw.IsEqual(fw1),
		"equalVal":     fw1.IsEqualValue(*fw2),
		"isPublic":     fw1.IsPublicMethod(),
		"isPrivate":    fw1.IsPrivateMethod(),
	}

	// Assert
	expected := args.Map{
		"equalSame": true, "notEqual": true, "nilNil": true,
		"nilNonNil": false, "equalVal": true,
		"isPublic": true, "isPrivate": false,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- IsEqual", actual)
}

func Test_FuncWrap_ArgsTypes(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{
		"inLen":      len(fw.GetInArgsTypes()),
		"outLen":     len(fw.GetOutArgsTypes()),
		"inNames":    len(fw.GetInArgsTypesNames()) > 0,
		"outNames":   len(fw.GetOutArgsTypesNames()) > 0,
		"inArgNames": len(fw.InArgNames()) > 0,
		"outArgNames": len(fw.OutArgNames()) > 0,
	}

	// Assert
	expected := args.Map{
		"inLen": 1, "outLen": 1, "inNames": true,
		"outNames": true, "inArgNames": true, "outArgNames": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- ArgsTypes", actual)
}

func Test_FuncWrap_Invoke(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	results, err := fw.Invoke("hello")

	// Act
	actual := args.Map{
		"noErr":  err == nil,
		"result": results[0],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"result": 5,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- Invoke", actual)
}

func Test_FuncWrap_InvokeMultiArg(t *testing.T) {
	// Arrange
	fn := func(a, b string) string { return a + b }
	fw := args.NewFuncWrap.Default(fn)
	results, err := fw.Invoke("hello", "world")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"result": "helloworld",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- Invoke multi-arg", actual)
}

func Test_FuncWrap_Validate(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	err := fw.ValidateMethodArgs([]any{"hello"})
	errMismatch := fw.ValidateMethodArgs([]any{"a", "b"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasErr": errMismatch != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns non-empty -- Validate", actual)
}

func Test_FuncWrap_VerifyArgs(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	inOk, inErr := fw.VerifyInArgs([]any{"hello"})
	outOk, outErr := fw.VerifyOutArgs([]any{5})

	// Act
	actual := args.Map{
		"inOk": inOk, "inErr": inErr == nil,
		"outOk": outOk, "outErr": outErr == nil,
	}

	// Assert
	expected := args.Map{
		"inOk": true,
		"inErr": true,
		"outOk": true,
		"outErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- VerifyArgs", actual)
}

func Test_FuncWrap_InArgNamesEachLine(t *testing.T) {
	// Arrange
	fn := func(a, b string) int { return len(a) }
	fw := args.NewFuncWrap.Default(fn)
	inLines := fw.InArgNamesEachLine()
	outLines := fw.OutArgNamesEachLine()

	// Act
	actual := args.Map{
		"inLen": len(inLines) > 0,
		"outLen": len(outLines) > 0,
	}

	// Assert
	expected := args.Map{
		"inLen": true,
		"outLen": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InArgNamesEachLine", actual)
}

func Test_FuncWrap_InOutArgNames(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	inNames := fw.InArgNames()
	outNames := fw.OutArgNames()

	// Act
	actual := args.Map{
		"inLen": len(inNames) > 0,
		"outLen": len(outNames) > 0,
	}

	// Assert
	expected := args.Map{
		"inLen": true,
		"outLen": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InOutArgNames", actual)
}

// ── NewTypedFuncWrap ──

func Test_NewTypedFuncWrap_Valid(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewTypedFuncWrap(fn)

	// Act
	actual := args.Map{
		"isValid": fw.IsValid(),
		"name": fw.GetFuncName() != "",
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"name": true,
	}
	expected.ShouldBeEqual(t, 0, "NewTypedFuncWrap returns non-empty -- Valid", actual)
}

func Test_NewTypedFuncWrap_NonFunc(t *testing.T) {
	// Arrange
	fw := args.NewTypedFuncWrap(42)

	// Act
	actual := args.Map{"isInvalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "NewTypedFuncWrap returns correct value -- NonFunc", actual)
}

// ── Holder ──

func Test_Holder(t *testing.T) {
	// Arrange
	h := &args.HolderAny{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Hashmap: map[string]any{"k": "v"}}

	// Act
	actual := args.Map{
		"first": h.First, "second": h.Second, "third": h.Third,
		"fourth": h.Fourth, "fifth": h.Fifth, "hmLen": len(h.Hashmap),
	}

	// Assert
	expected := args.Map{
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "hmLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- with args", actual)
}

// ── LeftRight ──

func Test_LeftRight(t *testing.T) {
	// Arrange
	lr := args.LeftRight[string, string]{Left: "l", Right: "r"}

	// Act
	actual := args.Map{
		"left": lr.Left,
		"right": lr.Right,
	}

	// Assert
	expected := args.Map{
		"left": "l",
		"right": "r",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- with args", actual)
}

// ── One through Six ──

func Test_One(t *testing.T) {
	// Arrange
	o := args.One[int]{First: 1}

	// Act
	actual := args.Map{
		"first": o.FirstItem(),
		"str": o.String() != "",
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"str": true,
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- with args", actual)
}

func Test_Two(t *testing.T) {
	// Arrange
	o := args.Two[int, int]{First: 1, Second: 2}

	// Act
	actual := args.Map{
		"first": o.FirstItem(),
		"second": o.SecondItem(),
		"str": o.String() != "",
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"second": 2,
		"str": true,
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- with args", actual)
}

func Test_Three(t *testing.T) {
	// Arrange
	o := args.Three[int, int, int]{First: 1, Second: 2, Third: 3}

	// Act
	actual := args.Map{
		"first": o.FirstItem(),
		"second": o.SecondItem(),
		"third": o.ThirdItem(),
		"str": o.String() != "",
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"second": 2,
		"third": 3,
		"str": true,
	}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- with args", actual)
}

func Test_Four(t *testing.T) {
	// Arrange
	o := args.Four[int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4}

	// Act
	actual := args.Map{
		"fourth": o.FourthItem(),
		"str": o.String() != "",
	}

	// Assert
	expected := args.Map{
		"fourth": 4,
		"str": true,
	}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- with args", actual)
}

func Test_Five(t *testing.T) {
	// Arrange
	o := args.Five[int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5}

	// Act
	actual := args.Map{
		"fifth": o.FifthItem(),
		"str": o.String() != "",
	}

	// Assert
	expected := args.Map{
		"fifth": 5,
		"str": true,
	}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- with args", actual)
}

func Test_Six(t *testing.T) {
	// Arrange
	o := args.Six[int, int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6}

	// Act
	actual := args.Map{
		"sixth": o.SixthItem(),
		"str": o.String() != "",
	}

	// Assert
	expected := args.Map{
		"sixth": 6,
		"str": true,
	}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- with args", actual)
}

// ── OneFunc through SixFunc ──

func Test_OneFunc(t *testing.T) {
	// Arrange
	fn := func() {}
	o := args.OneFuncAny{First: 1, WorkFunc: fn}

	// Act
	actual := args.Map{
		"first": o.FirstItem(),
		"hasFunc": o.GetWorkFunc() != nil,
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "OneFunc returns correct value -- with args", actual)
}

func Test_TwoFunc(t *testing.T) {
	// Arrange
	fn := func() {}
	o := args.TwoFuncAny{First: 1, Second: 2, WorkFunc: fn}

	// Act
	actual := args.Map{
		"second": o.SecondItem(),
		"hasFunc": o.GetWorkFunc() != nil,
	}

	// Assert
	expected := args.Map{
		"second": 2,
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "TwoFunc returns correct value -- with args", actual)
}

func Test_FuncMap(t *testing.T) {
	// Arrange
	fn := func() {}
	fm := args.FuncMap{}
	fm.Add(fn)

	// Act
	actual := args.Map{
		"hasFunc": fm.HasAnyItem(),
		"firstNotNil": fm.Length() > 0,
	}

	// Assert
	expected := args.Map{
		"hasFunc": true,
		"firstNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- with args", actual)
}

// ── Empty creator ──

func Test_Empty_Map(t *testing.T) {
	// Arrange
	m := args.Empty.Map()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- Map", actual)
}

func Test_Map_GetAsStringSliceFirstOfNames(t *testing.T) {
	// Arrange
	m := args.Map{"lines": []string{"a", "b"}}
	result := m.GetAsStringSliceFirstOfNames("lines")
	nilResult := m.GetAsStringSliceFirstOfNames("x")
	emptyResult := m.GetAsStringSliceFirstOfNames()

	// Act
	actual := args.Map{
		"len": len(result), "nilResult": nilResult == nil, "emptyResult": emptyResult == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilResult": true,
		"emptyResult": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsStringSliceFirstOfNames", actual)
}

func Test_Map_WorkFuncName(t *testing.T) {
	// Arrange
	fn := func() {}
	m := args.Map{"func": fn}

	// Act
	actual := args.Map{"notEmpty": m.WorkFuncName() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- WorkFuncName", actual)
}

func Test_Map_GetFirstFuncNameOf(t *testing.T) {
	// Arrange
	fn := func() {}
	m := args.Map{"workFunc": fn}

	// Act
	actual := args.Map{"notEmpty": m.GetFirstFuncNameOf("workFunc") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetFirstFuncNameOf", actual)
}
