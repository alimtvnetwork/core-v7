package argstests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// Dynamic — nil receiver, getters, caching
// ═══════════════════════════════════════════

func Test_Dynamic_NilReceiver(t *testing.T) {
	// Arrange
	var d *args.DynamicAny

	// Act
	actual := args.Map{
		"argsCount":  d.ArgsCount(),
		"getWorkFn":  d.GetWorkFunc() == nil,
		"hasFirst":   d.HasFirst(),
		"hasDefined": d.HasDefined("x"),
		"has":        d.Has("x"),
		"isInvalid":  d.IsKeyInvalid("x"),
		"isMissing":  d.IsKeyMissing("x"),
		"hasExpect":  d.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"argsCount":  0,
		"getWorkFn":  true,
		"hasFirst":   false,
		"hasDefined": false,
		"has":        false,
		"isInvalid":  false,
		"isMissing":  false,
		"hasExpect":  false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- nil receiver", actual)
}

func Test_Dynamic_GetMethods(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"str":   "hello",
			"num":   42,
			"strs":  []string{"a", "b"},
			"items": []any{1, 2},
		},
		Expect: "expected",
	}

	strVal, strOK := d.GetAsString("str")
	strDef := d.GetAsStringDefault("str")
	missDef := d.GetAsStringDefault("missing")
	numVal, numOK := d.GetAsInt("num")
	numDef := d.GetAsIntDefault("num", 0)
	misNumDef := d.GetAsIntDefault("missing", 99)
	strsVal, strsOK := d.GetAsStrings("strs")
	itemsVal, itemsOK := d.GetAsAnyItems("items")
	_, misStrOK := d.GetAsStrings("missing")
	_, misItemOK := d.GetAsAnyItems("missing")

	// Act
	actual := args.Map{
		"str": strVal, "strOK": strOK,
		"strDef": strDef, "missDef": missDef,
		"num": numVal, "numOK": numOK,
		"numDef": numDef, "misNumDef": misNumDef,
		"strsLen": len(strsVal), "strsOK": strsOK,
		"itemsLen": len(itemsVal), "itemsOK": itemsOK,
		"misStrOK": misStrOK, "misItemOK": misItemOK,
	}

	// Assert
	expected := args.Map{
		"str": "hello", "strOK": true,
		"strDef": "hello", "missDef": "",
		"num": 42, "numOK": true,
		"numDef": 42, "misNumDef": 99,
		"strsLen": 2, "strsOK": true,
		"itemsLen": 2, "itemsOK": true,
		"misStrOK": false, "misItemOK": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- get methods", actual)
}

func Test_Dynamic_GetLowerCase(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"actual": "val",
			"arrange": "arr",
		},
	}
	lcVal, lcOK := d.GetLowerCase("ACTUAL")
	directLower := d.GetDirectLower("ACTUAL")
	missingLower := d.GetDirectLower("NONEXIST")

	// Act
	actualVal := d.Actual()
	arrangeVal := d.Arrange()

	actual := args.Map{
		"lcVal": lcVal, "lcOK": lcOK,
		"directLower":  directLower,
		"missingLower": missingLower == nil,
		"actual":       actualVal,
		"arrange":      arrangeVal,
	}

	// Assert
	expected := args.Map{
		"lcVal": "val", "lcOK": true,
		"directLower":  "val",
		"missingLower": true,
		"actual":       "val",
		"arrange":      "arr",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- getLowerCase", actual)
}

func Test_Dynamic_HasDefinedAll(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"a": 1,
			"b": 2,
		},
	}

	// Act
	actual := args.Map{
		"allDefined":   d.HasDefinedAll("a", "b"),
		"missingOne":   d.HasDefinedAll("a", "c"),
		"emptyNames":   d.HasDefinedAll(),
	}

	// Assert
	expected := args.Map{
		"allDefined":   true,
		"missingOne":   false,
		"emptyNames":   false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- hasDefinedAll", actual)
}

func Test_Dynamic_SliceCaching(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{"a": 1},
		Expect: "ex",
	}
	s1 := d.Slice()
	s2 := d.Slice()
	str1 := d.String()
	str2 := d.String()

	// Act
	actual := args.Map{
		"sameSlice":  len(s1) == len(s2),
		"sameString": str1 == str2,
		"strNonEmpty": str1 != "",
	}

	// Assert
	expected := args.Map{
		"sameSlice": true, "sameString": true, "strNonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- slice caching", actual)
}

func Test_Dynamic_ItemAccessors(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"first": 1, "second": 2, "third": 3,
			"fourth": 4, "fifth": 5, "sixth": 6,
		},
		Expect: "exp",
	}

	// Act
	actual := args.Map{
		"first":  d.FirstItem(),
		"second": d.SecondItem(),
		"third":  d.ThirdItem(),
		"fourth": d.FourthItem(),
		"fifth":  d.FifthItem(),
		"sixth":  d.SixthItem(),
		"expect": d.Expected(),
	}

	// Assert
	expected := args.Map{
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "sixth": 6,
		"expect": "exp",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- item accessors", actual)
}

func Test_Dynamic_ValidArgs(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"a": 1,
			"b": nil,
			"c": 3,
		},
	}
	va := d.ValidArgs()
	namedArgs := d.Args("a", "c")

	// Act
	actual := args.Map{
		"validLen": len(va),
		"namedLen": len(namedArgs),
	}

	// Assert
	expected := args.Map{
		"validLen": 2,
		"namedLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns non-empty -- validArgs", actual)
}

func Test_Dynamic_GetByIndex(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"a": 1}}

	// Act
	actual := args.Map{
		"idx0NotNil": d.GetByIndex(0) != nil,
		"idx99Nil":   d.GetByIndex(99) == nil,
	}

	// Assert
	expected := args.Map{
		"idx0NotNil": true,
		"idx99Nil": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- getByIndex", actual)
}

func Test_Dynamic_AsInterfaces(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"a": 1}}

	// Act
	actual := args.Map{
		"mapper":   d.AsArgsMapper() != nil,
		"funcBind": d.AsArgFuncNameContractsBinder() != nil,
		"baseBind": d.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"mapper": true,
		"funcBind": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- asInterfaces", actual)
}

func Test_Dynamic_Get_NilReceiver(t *testing.T) {
	// Arrange
	var d *args.DynamicAny
	item, valid := d.Get("x")

	// Act
	actual := args.Map{
		"item": item == nil,
		"valid": valid,
	}

	// Assert
	expected := args.Map{
		"item": true,
		"valid": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- Get nil receiver", actual)
}

// ═══════════════════════════════════════════
// DynamicFunc — nil receiver, func methods
// ═══════════════════════════════════════════

func Test_DynamicFunc_NilReceiver(t *testing.T) {
	// Arrange
	var df *args.DynamicFuncAny

	// Act
	actual := args.Map{
		"argsCount":  df.ArgsCount(),
		"length":     df.Length(),
		"hasDefined": df.HasDefined("x"),
		"has":        df.Has("x"),
		"isInvalid":  df.IsKeyInvalid("x"),
		"isMissing":  df.IsKeyMissing("x"),
		"hasFunc":    df.HasFunc(),
		"hasExpect":  df.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"argsCount": 0, "length": 0,
		"hasDefined": false, "has": false,
		"isInvalid": false, "isMissing": false,
		"hasFunc": false, "hasExpect": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns nil -- nil receiver", actual)
}

func Test_DynamicFunc_FuncMethods(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	df := args.DynamicFuncAny{
		Params:   args.Map{"a": 1},
		WorkFunc: fn,
		Expect:   "ex",
	}

	// Act
	actual := args.Map{
		"getWorkFn": df.GetWorkFunc() != nil,
		"hasFunc":   df.HasFunc(),
		"hasExpect": df.HasExpect(),
		"funcName":  df.GetFuncName() != "",
		"expected":  df.Expected(),
	}

	// Assert
	expected := args.Map{
		"getWorkFn": true, "hasFunc": true,
		"hasExpect": true, "funcName": true,
		"expected": "ex",
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- func methods", actual)
}

func Test_DynamicFunc_GetMethods(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{
			"str":   "hello",
			"num":   42,
			"strs":  []string{"a"},
			"items": []any{1},
		},
	}
	strVal, strOK := df.GetAsString("str")
	numVal, numOK := df.GetAsInt("num")
	strsVal, strsOK := df.GetAsStrings("strs")
	itemsVal, itemsOK := df.GetAsAnyItems("items")

	// Act
	actual := args.Map{
		"str": strVal, "strOK": strOK,
		"num": numVal, "numOK": numOK,
		"strsLen": len(strsVal), "strsOK": strsOK,
		"itemsLen": len(itemsVal), "itemsOK": itemsOK,
	}

	// Assert
	expected := args.Map{
		"str": "hello", "strOK": true,
		"num": 42, "numOK": true,
		"strsLen": 1, "strsOK": true,
		"itemsLen": 1, "itemsOK": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- get methods", actual)
}

func Test_DynamicFunc_WhenTitle(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{
			"when": "w",
			"title": "t",
		},
	}

	// Act
	actual := args.Map{
		"when": df.When(),
		"title": df.Title(),
	}

	// Assert
	expected := args.Map{
		"when": "w",
		"title": "t",
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- when title", actual)
}

func Test_DynamicFunc_GetLowerCase(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{
			"actual": "val",
			"arrange": "arr",
		},
	}
	lcVal, lcOK := df.GetLowerCase("ACTUAL")
	directLower := df.GetDirectLower("ACTUAL")
	missingLower := df.GetDirectLower("NONEXIST")

	// Act
	actual := args.Map{
		"lcVal": lcVal, "lcOK": lcOK,
		"directLower":  directLower,
		"missingLower": missingLower == nil,
		"actual":       df.Actual(),
		"arrange":      df.Arrange(),
	}

	// Assert
	expected := args.Map{
		"lcVal": "val", "lcOK": true,
		"directLower":  "val",
		"missingLower": true,
		"actual":       "val",
		"arrange":      "arr",
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- getLowerCase", actual)
}

func Test_DynamicFunc_HasDefinedAll(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{
			"a": 1,
			"b": 2,
		},
	}

	// Act
	actual := args.Map{
		"allDefined": df.HasDefinedAll("a", "b"),
		"missingOne": df.HasDefinedAll("a", "c"),
		"emptyNames": df.HasDefinedAll(),
	}

	// Assert
	expected := args.Map{
		"allDefined": true, "missingOne": false, "emptyNames": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- hasDefinedAll", actual)
}

func Test_DynamicFunc_SliceCaching(t *testing.T) {
	// Arrange
	fn := func() {}
	df := args.DynamicFuncAny{
		Params:   args.Map{"a": 1},
		WorkFunc: fn,
		Expect:   "ex",
	}
	s1 := df.Slice()
	s2 := df.Slice()
	str1 := df.String()
	str2 := df.String()

	// Act
	actual := args.Map{
		"sameSlice":  len(s1) == len(s2),
		"sameString": str1 == str2,
		"strNonEmpty": str1 != "",
	}

	// Assert
	expected := args.Map{
		"sameSlice": true, "sameString": true, "strNonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- slice caching", actual)
}

func Test_DynamicFunc_ItemAccessors(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{
			"first": 1, "second": 2, "third": 3,
			"fourth": 4, "fifth": 5, "sixth": 6,
		},
	}

	// Act
	actual := args.Map{
		"first": df.FirstItem(), "second": df.SecondItem(),
		"third": df.ThirdItem(), "fourth": df.FourthItem(),
		"fifth": df.FifthItem(), "sixth": df.SixthItem(),
	}

	// Assert
	expected := args.Map{
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "sixth": 6,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- item accessors", actual)
}

func Test_DynamicFunc_GetByIndex(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"a": 1}}

	// Act
	actual := args.Map{
		"idx0NotNil": df.GetByIndex(0) != nil,
		"idx99Nil":   df.GetByIndex(99) == nil,
	}

	// Assert
	expected := args.Map{
		"idx0NotNil": true,
		"idx99Nil": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- getByIndex", actual)
}

func Test_DynamicFunc_ValidArgs(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{
			"a": 1,
			"b": nil,
			"c": 3,
		},
	}
	va := df.ValidArgs()
	namedArgs := df.Args("a", "c")

	// Act
	actual := args.Map{
		"validLen": len(va),
		"namedLen": len(namedArgs),
	}

	// Assert
	expected := args.Map{
		"validLen": 2,
		"namedLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns non-empty -- validArgs", actual)
}

func Test_DynamicFunc_AsInterfaces(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"a": 1}}

	// Act
	actual := args.Map{
		"mapper":   df.AsArgsMapper() != nil,
		"funcBind": df.AsArgFuncNameContractsBinder() != nil,
		"baseBind": df.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"mapper": true,
		"funcBind": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- asInterfaces", actual)
}

func Test_DynamicFunc_Get_NilReceiver(t *testing.T) {
	// Arrange
	var df *args.DynamicFuncAny
	item, valid := df.Get("x")

	// Act
	actual := args.Map{
		"item": item == nil,
		"valid": valid,
	}

	// Assert
	expected := args.Map{
		"item": true,
		"valid": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns nil -- Get nil", actual)
}

func Test_DynamicFunc_HasFirst(t *testing.T) {
	// Arrange
	df1 := args.DynamicFuncAny{Params: args.Map{"first": "hello"}}
	df2 := args.DynamicFuncAny{Params: args.Map{}}

	// Act
	actual := args.Map{
		"hasFirst1": df1.HasFirst(),
		"hasFirst2": df2.HasFirst(),
	}

	// Assert
	expected := args.Map{
		"hasFirst1": true,
		"hasFirst2": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- hasFirst", actual)
}

// ═══════════════════════════════════════════
// One — downcasts, args, interfaces
// ═══════════════════════════════════════════

func Test_One_AllMethods(t *testing.T) {
	// Arrange
	o := args.OneAny{First: "hello", Expect: 42}

	// Act
	actual := args.Map{
		"firstItem":  o.FirstItem(),
		"expected":   o.Expected(),
		"hasFirst":   o.HasFirst(),
		"hasExpect":  o.HasExpect(),
		"argsCount":  o.ArgsCount(),
		"args0":      len(o.Args(0)),
		"args1":      len(o.Args(1)),
		"validLen":   len(o.ValidArgs()),
		"getByIdx0":  o.GetByIndex(0),
		"getByIdx99": o.GetByIndex(99) == nil,
		"strNE":      o.String() != "",
	}

	// Assert
	expected := args.Map{
		"firstItem": "hello", "expected": 42,
		"hasFirst": true, "hasExpect": true,
		"argsCount": 1, "args0": 0, "args1": 1,
		"validLen": 1, "getByIdx0": "hello",
		"getByIdx99": true, "strNE": true,
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- all methods", actual)
}

func Test_One_ArgTwoAndLeftRight(t *testing.T) {
	// Arrange
	o := args.OneAny{First: "hello", Expect: 42}
	a2 := o.ArgTwo()
	lr := o.LeftRight()

	// Act
	actual := args.Map{
		"a2First": a2.First,
		"lrLeft":  lr.Left,
	}

	// Assert
	expected := args.Map{
		"a2First": "hello",
		"lrLeft": "hello",
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- argTwo leftRight", actual)
}

func Test_One_AsInterfaces(t *testing.T) {
	// Arrange
	o := args.OneAny{First: "hello"}

	// Act
	actual := args.Map{
		"oneParam": o.AsOneParameter() != nil,
		"baseBind": o.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"oneParam": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// OneFunc — func methods
// ═══════════════════════════════════════════

func Test_OneFunc_AllMethods(t *testing.T) {
	// Arrange
	fn := func(s string) string { return s + "!" }
	of := args.OneFuncAny{First: "hello", WorkFunc: fn, Expect: "hello!"}

	// Act
	actual := args.Map{
		"firstItem":  of.FirstItem(),
		"expected":   of.Expected(),
		"getWorkFn":  of.GetWorkFunc() != nil,
		"hasFirst":   of.HasFirst(),
		"hasFunc":    of.HasFunc(),
		"hasExpect":  of.HasExpect(),
		"funcNameNE": of.GetFuncName() != "",
		"argsCount":  of.ArgsCount(),
		"args0":      len(of.Args(0)),
		"args1":      len(of.Args(1)),
		"validLen":   len(of.ValidArgs()),
		"getByIdx0":  of.GetByIndex(0),
		"strNE":      of.String() != "",
	}

	// Assert
	expected := args.Map{
		"firstItem": "hello", "expected": "hello!",
		"getWorkFn": true, "hasFirst": true,
		"hasFunc": true, "hasExpect": true,
		"funcNameNE": true, "argsCount": 1,
		"args0": 0, "args1": 1, "validLen": 1,
		"getByIdx0": "hello", "strNE": true,
	}
	expected.ShouldBeEqual(t, 0, "OneFunc returns correct value -- all methods", actual)
}

func Test_OneFunc_ArgTwoAndLeftRight(t *testing.T) {
	// Arrange
	fn := func() {}
	of := args.OneFuncAny{First: "a", WorkFunc: fn, Expect: "b"}
	a2 := of.ArgTwo()
	lr := of.LeftRight()

	// Act
	actual := args.Map{
		"a2First": a2.First,
		"lrLeft":  lr.Left,
		"lrRight": lr.Right != nil,
	}

	// Assert
	expected := args.Map{
		"a2First": "a",
		"lrLeft": "a",
		"lrRight": true,
	}
	expected.ShouldBeEqual(t, 0, "OneFunc returns correct value -- argTwo leftRight", actual)
}

func Test_OneFunc_AsInterfaces(t *testing.T) {
	// Arrange
	of := args.OneFuncAny{First: "a"}

	// Act
	actual := args.Map{
		"oneFuncP": of.AsOneFuncParameter() != nil,
		"funcBind": of.AsArgFuncContractsBinder() != nil,
		"baseBind": of.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"oneFuncP": true,
		"funcBind": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "OneFunc returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// Two — downcasts, args, interfaces
// ═══════════════════════════════════════════

func Test_Two_AllMethods(t *testing.T) {
	// Arrange
	tw := args.TwoAny{First: "a", Second: "b", Expect: "c"}

	// Act
	actual := args.Map{
		"first":    tw.FirstItem(),
		"second":   tw.SecondItem(),
		"expected": tw.Expected(),
		"hasFirst": tw.HasFirst(),
		"hasSec":   tw.HasSecond(),
		"hasExp":   tw.HasExpect(),
		"count":    tw.ArgsCount(),
		"args0":    len(tw.Args(0)),
		"args1":    len(tw.Args(1)),
		"args2":    len(tw.Args(2)),
		"validLen": len(tw.ValidArgs()),
		"strNE":    tw.String() != "",
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "expected": "c",
		"hasFirst": true, "hasSec": true, "hasExp": true,
		"count": 2, "args0": 0, "args1": 1, "args2": 2,
		"validLen": 2, "strNE": true,
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- all methods", actual)
}

func Test_Two_DowncastAndConvert(t *testing.T) {
	// Arrange
	tw := args.TwoAny{First: "a", Second: "b", Expect: "c"}
	a2 := tw.ArgTwo()
	lr := tw.LeftRight()

	// Act
	actual := args.Map{
		"a2First": a2.First,
		"lrLeft":  lr.Left,
		"lrRight": lr.Right,
	}

	// Assert
	expected := args.Map{
		"a2First": "a",
		"lrLeft": "a",
		"lrRight": "b",
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- downcast", actual)
}

func Test_Two_AsInterfaces(t *testing.T) {
	// Arrange
	tw := args.TwoAny{First: "a", Second: "b"}

	// Act
	actual := args.Map{
		"twoP":     tw.AsTwoParameter() != nil,
		"baseBind": tw.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"twoP": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// TwoFunc — all methods
// ═══════════════════════════════════════════

func Test_TwoFunc_AllMethods(t *testing.T) {
	// Arrange
	fn := func(a, b string) string { return a + b }
	tf := args.TwoFuncAny{First: "a", Second: "b", WorkFunc: fn, Expect: "ab"}

	// Act
	actual := args.Map{
		"first":    tf.FirstItem(),
		"second":   tf.SecondItem(),
		"expected": tf.Expected(),
		"getWF":    tf.GetWorkFunc() != nil,
		"hasFirst": tf.HasFirst(),
		"hasSec":   tf.HasSecond(),
		"hasFunc":  tf.HasFunc(),
		"hasExp":   tf.HasExpect(),
		"fnName":   tf.GetFuncName() != "",
		"count":    tf.ArgsCount(),
		"args0":    len(tf.Args(0)),
		"args1":    len(tf.Args(1)),
		"args2":    len(tf.Args(2)),
		"validLen": len(tf.ValidArgs()),
		"strNE":    tf.String() != "",
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "expected": "ab",
		"getWF": true, "hasFirst": true, "hasSec": true,
		"hasFunc": true, "hasExp": true, "fnName": true,
		"count": 2, "args0": 0, "args1": 1, "args2": 2,
		"validLen": 2, "strNE": true,
	}
	expected.ShouldBeEqual(t, 0, "TwoFunc returns correct value -- all methods", actual)
}

func Test_TwoFunc_DowncastAndConvert(t *testing.T) {
	// Arrange
	tf := args.TwoFuncAny{First: "a", Second: "b", Expect: "c"}
	a2 := tf.ArgTwo()
	lr := tf.LeftRight()

	// Act
	actual := args.Map{
		"a2First": a2.First,
		"lrLeft":  lr.Left,
	}

	// Assert
	expected := args.Map{
		"a2First": "a",
		"lrLeft": "a",
	}
	expected.ShouldBeEqual(t, 0, "TwoFunc returns correct value -- downcast", actual)
}

func Test_TwoFunc_AsInterfaces(t *testing.T) {
	// Arrange
	tf := args.TwoFuncAny{First: "a", Second: "b"}

	// Act
	actual := args.Map{
		"twoFuncP": tf.AsTwoFuncParameter() != nil,
		"funcBind": tf.AsArgFuncContractsBinder() != nil,
		"baseBind": tf.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"twoFuncP": true,
		"funcBind": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "TwoFunc returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// Three — all methods
// ═══════════════════════════════════════════

func Test_Three_AllMethods(t *testing.T) {
	// Arrange
	th := args.ThreeAny{First: "a", Second: "b", Third: "c", Expect: "d"}
	a2 := th.ArgTwo()
	a3 := th.ArgThree()
	lr := th.LeftRight()

	// Act
	actual := args.Map{
		"first": th.FirstItem(), "second": th.SecondItem(), "third": th.ThirdItem(),
		"expected": th.Expected(), "hasFirst": th.HasFirst(),
		"hasSec": th.HasSecond(), "hasThird": th.HasThird(),
		"hasExp": th.HasExpect(), "count": th.ArgsCount(),
		"args0": len(th.Args(0)), "args1": len(th.Args(1)),
		"args2": len(th.Args(2)), "args3": len(th.Args(3)),
		"validLen": len(th.ValidArgs()), "strNE": th.String() != "",
		"a2First": a2.First, "a3First": a3.First,
		"lrLeft": lr.Left,
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"expected": "d", "hasFirst": true,
		"hasSec": true, "hasThird": true,
		"hasExp": true, "count": 3,
		"args0": 0, "args1": 1, "args2": 2, "args3": 3,
		"validLen": 3, "strNE": true,
		"a2First": "a", "a3First": "a", "lrLeft": "a",
	}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- all methods", actual)
}

func Test_Three_AsInterfaces(t *testing.T) {
	// Arrange
	th := args.ThreeAny{First: "a", Second: "b", Third: "c"}

	// Act
	actual := args.Map{
		"threeP":   th.AsThreeParameter() != nil,
		"baseBind": th.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"threeP": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// ThreeFunc — all methods
// ═══════════════════════════════════════════

func Test_ThreeFunc_AllMethods(t *testing.T) {
	// Arrange
	fn := func(a, b, c string) string { return a + b + c }
	tf := args.ThreeFuncAny{First: "a", Second: "b", Third: "c", WorkFunc: fn, Expect: "abc"}
	a2 := tf.ArgTwo()
	a3 := tf.ArgThree()
	lr := tf.LeftRight()

	// Act
	actual := args.Map{
		"first": tf.FirstItem(), "second": tf.SecondItem(), "third": tf.ThirdItem(),
		"expected": tf.Expected(), "getWF": tf.GetWorkFunc() != nil,
		"hasFirst": tf.HasFirst(), "hasSec": tf.HasSecond(),
		"hasThird": tf.HasThird(), "hasFunc": tf.HasFunc(),
		"hasExp": tf.HasExpect(), "fnName": tf.GetFuncName() != "",
		"count": tf.ArgsCount(), "args0": len(tf.Args(0)),
		"args3": len(tf.Args(3)), "validLen": len(tf.ValidArgs()),
		"strNE": tf.String() != "",
		"a2First": a2.First, "a3First": a3.First, "lrLeft": lr.Left,
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"expected": "abc", "getWF": true,
		"hasFirst": true, "hasSec": true,
		"hasThird": true, "hasFunc": true,
		"hasExp": true, "fnName": true,
		"count": 3, "args0": 0, "args3": 3, "validLen": 3,
		"strNE": true, "a2First": "a", "a3First": "a", "lrLeft": "a",
	}
	expected.ShouldBeEqual(t, 0, "ThreeFunc returns correct value -- all methods", actual)
}

func Test_ThreeFunc_AsInterfaces(t *testing.T) {
	// Arrange
	tf := args.ThreeFuncAny{First: "a"}

	// Act
	actual := args.Map{
		"threeFuncP": tf.AsThreeFuncParameter() != nil,
		"funcBind":   tf.AsArgFuncContractsBinder() != nil,
		"baseBind":   tf.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"threeFuncP": true,
		"funcBind": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "ThreeFunc returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// Four — all methods
// ═══════════════════════════════════════════

func Test_Four_AllMethods(t *testing.T) {
	// Arrange
	f := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d", Expect: "e"}
	a2 := f.ArgTwo()
	a3 := f.ArgThree()

	// Act
	actual := args.Map{
		"first": f.FirstItem(), "second": f.SecondItem(),
		"third": f.ThirdItem(), "fourth": f.FourthItem(),
		"expected": f.Expected(),
		"hasFirst": f.HasFirst(), "hasSec": f.HasSecond(),
		"hasThird": f.HasThird(), "hasFourth": f.HasFourth(),
		"hasExp": f.HasExpect(), "count": f.ArgsCount(),
		"args0": len(f.Args(0)), "args4": len(f.Args(4)),
		"validLen": len(f.ValidArgs()), "strNE": f.String() != "",
		"a2First": a2.First, "a3First": a3.First,
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c", "fourth": "d",
		"expected": "e",
		"hasFirst": true, "hasSec": true, "hasThird": true, "hasFourth": true,
		"hasExp": true, "count": 4,
		"args0": 0, "args4": 4, "validLen": 4, "strNE": true,
		"a2First": "a", "a3First": "a",
	}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- all methods", actual)
}

func Test_Four_AsInterfaces(t *testing.T) {
	// Arrange
	f := args.FourAny{First: "a"}

	// Act
	actual := args.Map{
		"fourP":    f.AsFourParameter() != nil,
		"baseBind": f.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"fourP": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// FourFunc — all methods
// ═══════════════════════════════════════════

func Test_FourFunc_AllMethods(t *testing.T) {
	// Arrange
	fn := func() {}
	ff := args.FourFuncAny{First: "a", Second: "b", Third: "c", Fourth: "d", WorkFunc: fn, Expect: "e"}
	a2 := ff.ArgTwo()
	a3 := ff.ArgThree()
	a4 := ff.ArgFour()

	// Act
	actual := args.Map{
		"first": ff.FirstItem(), "second": ff.SecondItem(),
		"third": ff.ThirdItem(), "fourth": ff.FourthItem(),
		"expected": ff.Expected(), "getWF": ff.GetWorkFunc() != nil,
		"hasFirst": ff.HasFirst(), "hasSec": ff.HasSecond(),
		"hasThird": ff.HasThird(), "hasFourth": ff.HasFourth(),
		"hasFunc": ff.HasFunc(), "hasExp": ff.HasExpect(),
		"fnName": ff.GetFuncName() != "",
		"count": ff.ArgsCount(), "args0": len(ff.Args(0)), "args4": len(ff.Args(4)),
		"validLen": len(ff.ValidArgs()), "strNE": ff.String() != "",
		"a2First": a2.First, "a3First": a3.First, "a4First": a4.First,
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c", "fourth": "d",
		"expected": "e", "getWF": true,
		"hasFirst": true, "hasSec": true, "hasThird": true, "hasFourth": true,
		"hasFunc": true, "hasExp": true, "fnName": true,
		"count": 4, "args0": 0, "args4": 4, "validLen": 4, "strNE": true,
		"a2First": "a", "a3First": "a", "a4First": "a",
	}
	expected.ShouldBeEqual(t, 0, "FourFunc returns correct value -- all methods", actual)
}

func Test_FourFunc_AsInterfaces(t *testing.T) {
	// Arrange
	ff := args.FourFuncAny{First: "a"}

	// Act
	actual := args.Map{
		"fourFuncP": ff.AsFourFuncParameter() != nil,
		"funcBind":  ff.AsArgFuncContractsBinder() != nil,
		"baseBind":  ff.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"fourFuncP": true,
		"funcBind": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "FourFunc returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// Five — all methods
// ═══════════════════════════════════════════

func Test_Five_AllMethods(t *testing.T) {
	// Arrange
	f := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Expect: "f"}
	a2 := f.ArgTwo()
	a3 := f.ArgThree()
	a4 := f.ArgFour()

	// Act
	actual := args.Map{
		"first": f.FirstItem(), "second": f.SecondItem(),
		"third": f.ThirdItem(), "fourth": f.FourthItem(),
		"fifth": f.FifthItem(), "expected": f.Expected(),
		"hasFirst": f.HasFirst(), "hasSec": f.HasSecond(),
		"hasThird": f.HasThird(), "hasFourth": f.HasFourth(),
		"hasFifth": f.HasFifth(), "hasExp": f.HasExpect(),
		"count": f.ArgsCount(), "args0": len(f.Args(0)), "args5": len(f.Args(5)),
		"validLen": len(f.ValidArgs()), "strNE": f.String() != "",
		"a2First": a2.First, "a3First": a3.First, "a4First": a4.First,
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c", "fourth": "d",
		"fifth": "e", "expected": "f",
		"hasFirst": true, "hasSec": true, "hasThird": true,
		"hasFourth": true, "hasFifth": true, "hasExp": true,
		"count": 5, "args0": 0, "args5": 5, "validLen": 5, "strNE": true,
		"a2First": "a", "a3First": "a", "a4First": "a",
	}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- all methods", actual)
}

func Test_Five_AsInterfaces(t *testing.T) {
	// Arrange
	f := args.FiveAny{First: "a"}

	// Act
	actual := args.Map{
		"fifthP":   f.AsFifthParameter() != nil,
		"baseBind": f.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"fifthP": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// FiveFunc — all methods
// ═══════════════════════════════════════════

func Test_FiveFunc_AllMethods(t *testing.T) {
	// Arrange
	fn := func() {}
	ff := args.FiveFuncAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", WorkFunc: fn, Expect: "f"}
	a2 := ff.ArgTwo()
	a3 := ff.ArgThree()
	a4 := ff.ArgFour()

	// Act
	actual := args.Map{
		"first": ff.FirstItem(), "second": ff.SecondItem(),
		"third": ff.ThirdItem(), "fourth": ff.FourthItem(),
		"fifth": ff.FifthItem(), "expected": ff.Expected(),
		"getWF": ff.GetWorkFunc() != nil,
		"hasFirst": ff.HasFirst(), "hasSec": ff.HasSecond(),
		"hasThird": ff.HasThird(), "hasFourth": ff.HasFourth(),
		"hasFifth": ff.HasFifth(), "hasFunc": ff.HasFunc(),
		"hasExp": ff.HasExpect(), "fnName": ff.GetFuncName() != "",
		"count": ff.ArgsCount(), "args5": len(ff.Args(5)),
		"validLen": len(ff.ValidArgs()), "strNE": ff.String() != "",
		"a2First": a2.First, "a3First": a3.First, "a4First": a4.First,
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c", "fourth": "d",
		"fifth": "e", "expected": "f", "getWF": true,
		"hasFirst": true, "hasSec": true, "hasThird": true,
		"hasFourth": true, "hasFifth": true, "hasFunc": true,
		"hasExp": true, "fnName": true,
		"count": 5, "args5": 5, "validLen": 5, "strNE": true,
		"a2First": "a", "a3First": "a", "a4First": "a",
	}
	expected.ShouldBeEqual(t, 0, "FiveFunc returns correct value -- all methods", actual)
}

func Test_FiveFunc_AsInterfaces(t *testing.T) {
	// Arrange
	ff := args.FiveFuncAny{First: "a"}

	// Act
	actual := args.Map{
		"fifthFuncP": ff.AsFifthFuncParameter() != nil,
		"funcBind":   ff.AsArgFuncContractsBinder() != nil,
		"baseBind":   ff.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"fifthFuncP": true,
		"funcBind": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "FiveFunc returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// Six — all methods
// ═══════════════════════════════════════════

func Test_Six_AllMethods(t *testing.T) {
	// Arrange
	s := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f", Expect: "g"}
	a2 := s.ArgTwo()
	a3 := s.ArgThree()
	a4 := s.ArgFour()
	a5 := s.ArgFive()

	// Act
	actual := args.Map{
		"first": s.FirstItem(), "second": s.SecondItem(),
		"third": s.ThirdItem(), "fourth": s.FourthItem(),
		"fifth": s.FifthItem(), "sixth": s.SixthItem(),
		"expected": s.Expected(),
		"hasFirst": s.HasFirst(), "hasSec": s.HasSecond(),
		"hasThird": s.HasThird(), "hasFourth": s.HasFourth(),
		"hasFifth": s.HasFifth(), "hasSixth": s.HasSixth(),
		"hasExp": s.HasExpect(), "count": s.ArgsCount(),
		"args0": len(s.Args(0)), "args6": len(s.Args(6)),
		"validLen": len(s.ValidArgs()), "strNE": s.String() != "",
		"a2First": a2.First, "a3First": a3.First,
		"a4First": a4.First, "a5First": a5.First,
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c", "fourth": "d",
		"fifth": "e", "sixth": "f", "expected": "g",
		"hasFirst": true, "hasSec": true, "hasThird": true,
		"hasFourth": true, "hasFifth": true, "hasSixth": true,
		"hasExp": true, "count": 6, "args0": 0, "args6": 6,
		"validLen": 6, "strNE": true,
		"a2First": "a", "a3First": "a", "a4First": "a", "a5First": "a",
	}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- all methods", actual)
}

func Test_Six_AsInterfaces(t *testing.T) {
	// Arrange
	s := args.SixAny{First: "a"}

	// Act
	actual := args.Map{
		"sixthP":   s.AsSixthParameter() != nil,
		"baseBind": s.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"sixthP": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// SixFunc — all methods
// ═══════════════════════════════════════════

func Test_SixFunc_AllMethods(t *testing.T) {
	// Arrange
	fn := func() {}
	sf := args.SixFuncAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f", WorkFunc: fn, Expect: "g"}
	a2 := sf.ArgTwo()
	a3 := sf.ArgThree()
	a4 := sf.ArgFour()
	a5 := sf.ArgFive()

	// Act
	actual := args.Map{
		"first": sf.FirstItem(), "second": sf.SecondItem(),
		"third": sf.ThirdItem(), "fourth": sf.FourthItem(),
		"fifth": sf.FifthItem(), "sixth": sf.SixthItem(),
		"expected": sf.Expected(), "getWF": sf.GetWorkFunc() != nil,
		"hasFirst": sf.HasFirst(), "hasSec": sf.HasSecond(),
		"hasThird": sf.HasThird(), "hasFourth": sf.HasFourth(),
		"hasFifth": sf.HasFifth(), "hasSixth": sf.HasSixth(),
		"hasFunc": sf.HasFunc(), "hasExp": sf.HasExpect(),
		"fnName": sf.GetFuncName() != "",
		"count": sf.ArgsCount(), "args6": len(sf.Args(6)),
		"validLen": len(sf.ValidArgs()), "strNE": sf.String() != "",
		"a2First": a2.First, "a3First": a3.First,
		"a4First": a4.First, "a5First": a5.First,
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c", "fourth": "d",
		"fifth": "e", "sixth": "f", "expected": "g", "getWF": true,
		"hasFirst": true, "hasSec": true, "hasThird": true,
		"hasFourth": true, "hasFifth": true, "hasSixth": true,
		"hasFunc": true, "hasExp": true, "fnName": true,
		"count": 6, "args6": 6, "validLen": 6, "strNE": true,
		"a2First": "a", "a3First": "a", "a4First": "a", "a5First": "a",
	}
	expected.ShouldBeEqual(t, 0, "SixFunc returns correct value -- all methods", actual)
}

func Test_SixFunc_AsInterfaces(t *testing.T) {
	// Arrange
	sf := args.SixFuncAny{First: "a"}

	// Act
	actual := args.Map{
		"sixthFuncP": sf.AsSixthFuncParameter() != nil,
		"funcBind":   sf.AsArgFuncContractsBinder() != nil,
		"baseBind":   sf.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"sixthFuncP": true,
		"funcBind": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "SixFunc returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// Holder — all methods
// ═══════════════════════════════════════════

func Test_Holder_AllMethods(t *testing.T) {
	// Arrange
	fn := func() {}
	h := args.HolderAny{
		First: "a", Second: "b", Third: "c",
		Fourth: "d", Fifth: "e", Sixth: "f",
		WorkFunc: fn, Expect: "g",
		Hashmap: args.Map{"extra": "val"},
	}
	a2 := h.ArgTwo()
	a3 := h.ArgThree()
	a4 := h.ArgFour()
	a5 := h.ArgFive()

	// Act
	actual := args.Map{
		"first": h.FirstItem(), "second": h.SecondItem(),
		"third": h.ThirdItem(), "fourth": h.FourthItem(),
		"fifth": h.FifthItem(), "sixth": h.SixthItem(),
		"expected": h.Expected(), "getWF": h.GetWorkFunc() != nil,
		"hasFirst": h.HasFirst(), "hasSec": h.HasSecond(),
		"hasThird": h.HasThird(), "hasFourth": h.HasFourth(),
		"hasFifth": h.HasFifth(), "hasSixth": h.HasSixth(),
		"hasFunc": h.HasFunc(), "hasExp": h.HasExpect(),
		"fnName": h.GetFuncName() != "",
		"count": h.ArgsCount(), "args0": len(h.Args(0)), "args6": len(h.Args(6)),
		"validLen": len(h.ValidArgs()), "strNE": h.String() != "",
		"a2First": a2.First, "a3First": a3.First,
		"a4First": a4.First, "a5First": a5.First,
		"getByIdx0": h.GetByIndex(0),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"fourth": "d", "fifth": "e", "sixth": "f",
		"expected": "g", "getWF": true,
		"hasFirst": true, "hasSec": true, "hasThird": true,
		"hasFourth": true, "hasFifth": true, "hasSixth": true,
		"hasFunc": true, "hasExp": true, "fnName": true,
		"count": 7, "args0": 0, "args6": 6, "validLen": 6, "strNE": true,
		"a2First": "a", "a3First": "a",
		"a4First": "a", "a5First": "a",
		"getByIdx0": "a",
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- all methods", actual)
}

func Test_Holder_AsInterfaces(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a"}

	// Act
	actual := args.Map{
		"sixthP":   h.AsSixthParameter() != nil,
		"funcBind": h.AsArgFuncContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"sixthP": true,
		"funcBind": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// LeftRight — clone, has methods
// ═══════════════════════════════════════════

func Test_LeftRight_AllMethods(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "a", Right: "b", Expect: "c"}
	clone := lr.Clone()
	a2 := lr.ArgTwo()

	// Act
	actual := args.Map{
		"first": lr.FirstItem(), "second": lr.SecondItem(),
		"expected": lr.Expected(),
		"hasFirst": lr.HasFirst(), "hasSec": lr.HasSecond(),
		"hasLeft": lr.HasLeft(), "hasRight": lr.HasRight(),
		"hasExp": lr.HasExpect(), "count": lr.ArgsCount(),
		"args0": len(lr.Args(0)), "args2": len(lr.Args(2)),
		"validLen": len(lr.ValidArgs()), "strNE": lr.String() != "",
		"cloneLeft": clone.Left, "a2First": a2.First,
		"getByIdx0": lr.GetByIndex(0),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "expected": "c",
		"hasFirst": true, "hasSec": true,
		"hasLeft": true, "hasRight": true,
		"hasExp": true, "count": 2,
		"args0": 0, "args2": 2, "validLen": 2, "strNE": true,
		"cloneLeft": "a", "a2First": "a", "getByIdx0": "a",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- all methods", actual)
}

func Test_LeftRight_AsInterfaces(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "a", Right: "b"}

	// Act
	actual := args.Map{
		"twoP":     lr.AsTwoParameter() != nil,
		"baseBind": lr.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"twoP": true,
		"baseBind": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- asInterfaces", actual)
}

// ═══════════════════════════════════════════
// Map — remaining methods
// ═══════════════════════════════════════════

func Test_Map_NilMap(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{
		"hasDefined":    m.HasDefined("x"),
		"has":           m.Has("x"),
		"hasDefinedAll": m.HasDefinedAll("x"),
		"isInvalid":     m.IsKeyInvalid("x"),
		"isMissing":     m.IsKeyMissing("x"),
	}

	// Assert
	expected := args.Map{
		"hasDefined": false, "has": false,
		"hasDefinedAll": false, "isInvalid": false, "isMissing": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- nil", actual)
}

func Test_Map_GetNilMap(t *testing.T) {
	// Arrange
	var m args.Map
	item, valid := m.Get("x")

	// Act
	actual := args.Map{
		"item": item == nil,
		"valid": valid,
	}

	// Assert
	expected := args.Map{
		"item": true,
		"valid": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- Get nil", actual)
}

func Test_Map_GetAsStringDefault(t *testing.T) {
	// Arrange
	m := args.Map{"name": "hello"}

	// Act
	actual := args.Map{
		"found":   m.GetAsStringDefault("name"),
		"missing": m.GetAsStringDefault("nope"),
	}

	// Assert
	expected := args.Map{
		"found": "hello",
		"missing": "",
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsStringDefault", actual)
}

func Test_Map_GetAsIntDefault(t *testing.T) {
	// Arrange
	m := args.Map{"val": 42}

	// Act
	actual := args.Map{
		"found":   m.GetAsIntDefault("val", 0),
		"missing": m.GetAsIntDefault("nope", 99),
	}

	// Assert
	expected := args.Map{
		"found": 42,
		"missing": 99,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsIntDefault", actual)
}

func Test_Map_GetAsAnyItems(t *testing.T) {
	// Arrange
	m := args.Map{"items": []any{1, 2, 3}}
	items, ok := m.GetAsAnyItems("items")
	_, misOK := m.GetAsAnyItems("missing")

	// Act
	actual := args.Map{
		"len": len(items),
		"ok": ok,
		"misOK": misOK,
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"ok": true,
		"misOK": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsAnyItems", actual)
}

func Test_Map_Accessors(t *testing.T) {
	// Arrange
	m := args.Map{
		"when": "w", "title": "t",
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "sixth": 6,
		"seventh": 7, "expect": "e",
	}

	// Act
	actual := args.Map{
		"when": m.When(), "title": m.Title(),
		"first": m.FirstItem(), "second": m.SecondItem(),
		"third": m.ThirdItem(), "fourth": m.FourthItem(),
		"fifth": m.FifthItem(), "sixth": m.SixthItem(),
		"seventh": m.Seventh(), "expect": m.Expect(),
	}

	// Assert
	expected := args.Map{
		"when": "w", "title": "t",
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "sixth": 6,
		"seventh": 7, "expect": "e",
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- accessors", actual)
}

func Test_Map_SetActual(t *testing.T) {
	// Arrange
	m := args.Map{}
	m.SetActual(42)

	// Act
	actual := args.Map{"actual": m.Actual()}

	// Assert
	expected := args.Map{"actual": 42}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- SetActual", actual)
}

func Test_Map_GetAsStringSliceFirstOfNames(t *testing.T) {
	// Arrange
	m := args.Map{"strs": []string{"a", "b"}}
	result := m.GetAsStringSliceFirstOfNames("strs")
	nilResult := m.GetAsStringSliceFirstOfNames("missing")
	emptyResult := m.GetAsStringSliceFirstOfNames()

	// Act
	actual := args.Map{
		"len":    len(result),
		"nilRes": nilResult == nil,
		"empRes": emptyResult == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilRes": true,
		"empRes": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsStringSliceFirstOfNames", actual)
}

func Test_Map_SortedKeys(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	keys, err := m.SortedKeys()
	empty := args.Map{}
	emptyKeys, _ := empty.SortedKeys()

	// Act
	actual := args.Map{
		"errNil":  err == nil,
		"first":   keys[0],
		"second":  keys[1],
		"empLen":  len(emptyKeys),
	}

	// Assert
	expected := args.Map{
		"errNil": true, "first": "a", "second": "b", "empLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- SortedKeys", actual)
}

func Test_Map_CompileToString(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	s := m.CompileToString()

	// Act
	actual := args.Map{"nonEmpty": s != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- CompileToString", actual)
}

func Test_Map_GoLiteralString(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	s := m.GoLiteralString()

	// Act
	actual := args.Map{"nonEmpty": s != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GoLiteralString", actual)
}

func Test_Map_EmptyGoLiteralLines(t *testing.T) {
	// Arrange
	m := args.Map{}
	lines := m.GoLiteralLines()

	// Act
	actual := args.Map{"len": len(lines)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- empty GoLiteralLines", actual)
}

func Test_Map_WorkFuncName(t *testing.T) {
	// Arrange
	fn := func() {}
	m := args.Map{"func": fn}
	name := m.WorkFuncName()

	// Act
	actual := args.Map{"nameNE": name != ""}

	// Assert
	expected := args.Map{"nameNE": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- WorkFuncName", actual)
}

func Test_Map_GetFirstFuncNameOf(t *testing.T) {
	// Arrange
	fn := func() {}
	m := args.Map{"func": fn}
	name := m.GetFirstFuncNameOf("func")

	// Act
	actual := args.Map{"nameNE": name != ""}

	// Assert
	expected := args.Map{"nameNE": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetFirstFuncNameOf", actual)
}

func Test_Map_GetFuncName(t *testing.T) {
	// Arrange
	m1 := args.Map{}
	m2 := args.Map{"func": func() {}}

	// Act
	actual := args.Map{
		"noFunc":  m1.GetFuncName(),
		"hasFunc": m2.GetFuncName() != "",
	}

	// Assert
	expected := args.Map{
		"noFunc": "",
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetFuncName", actual)
}

func Test_Map_Slice(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}
	s := m.Slice()

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Slice", actual)
}

func Test_Map_String(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	s := m.String()

	// Act
	actual := args.Map{"contains": strings.Contains(s, "Map")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- String", actual)
}

// ═══════════════════════════════════════════
// String type — all methods
// ═══════════════════════════════════════════

func Test_String_AllMethods(t *testing.T) {
	// Arrange
	s := args.String("hello world")

	// Act
	actual := args.Map{
		"string":    s.String(),
		"length":    s.Length(),
		"count":     s.Count(),
		"asciiLen":  s.AscIILength(),
		"isEmpty":   s.IsEmpty(),
		"isDefined": s.IsDefined(),
		"hasChar":   s.HasCharacter(),
		"isEmptyWS": s.IsEmptyOrWhitespace(),
		"trimSpace": s.TrimSpace().String(),
		"replace":   s.ReplaceAll("hello", "bye").String(),
		"sub":       s.Substring(0, 5).String(),
		"bytesLen":  len(s.Bytes()),
		"runesLen":  len(s.Runes()),
	}

	// Assert
	expected := args.Map{
		"string": "hello world", "length": 11,
		"count": 11, "asciiLen": 11,
		"isEmpty": false, "isDefined": true, "hasChar": true,
		"isEmptyWS": false, "trimSpace": "hello world",
		"replace": "bye world", "sub": "hello",
		"bytesLen": 11, "runesLen": 11,
	}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- all methods", actual)
}

func Test_String_ConcatJoinSplit(t *testing.T) {
	// Arrange
	s := args.String("hello")
	concat := s.Concat(" world")
	joined := args.String("a").Join("-", "b", "c")
	split := args.String("a,b,c").Split(",")

	// Act
	actual := args.Map{
		"concat":   concat.String(),
		"joined":   joined.String(),
		"splitLen": len(split),
	}

	// Assert
	expected := args.Map{
		"concat": "hello world",
		"joined": "a-b-c",
		"splitLen": 3,
	}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- concat join split", actual)
}

func Test_String_Quotes(t *testing.T) {
	// Arrange
	s := args.String("test")
	dq := s.DoubleQuote()
	dqq := s.DoubleQuoteQ()
	sq := s.SingleQuote()
	vdq := s.ValueDoubleQuote()

	// Act
	actual := args.Map{
		"dqNE":  dq.String() != "",
		"dqqNE": dqq.String() != "",
		"sqNE":  sq.String() != "",
		"vdqNE": vdq.String() != "",
	}

	// Assert
	expected := args.Map{
		"dqNE": true, "dqqNE": true, "sqNE": true, "vdqNE": true,
	}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- quotes", actual)
}

func Test_String_Empty(t *testing.T) {
	// Arrange
	s := args.String("")
	ws := args.String("  ")

	// Act
	actual := args.Map{
		"isEmpty":   s.IsEmpty(),
		"isDefined": s.IsDefined(),
		"hasChar":   s.HasCharacter(),
		"isEmptyWS": ws.IsEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true, "isDefined": false, "hasChar": false, "isEmptyWS": true,
	}
	expected.ShouldBeEqual(t, 0, "String returns empty -- empty", actual)
}

// ═══════════════════════════════════════════
// emptyCreator — all methods
// ═══════════════════════════════════════════

func Test_EmptyCreator(t *testing.T) {
	// Arrange
	m := args.Empty.Map()
	fw := args.Empty.FuncWrap()
	fm := args.Empty.FuncMap()
	h := args.Empty.Holder()

	// Act
	actual := args.Map{
		"mapLen":  len(m),
		"fwInval": fw.IsInvalid(),
		"fmLen":   len(fm),
		"hCount":  h.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"mapLen": 0, "fwInval": true, "fmLen": 0, "hCount": 7,
	}
	expected.ShouldBeEqual(t, 0, "emptyCreator returns empty -- with args", actual)
}

// ═══════════════════════════════════════════
// funcDetector — GetFuncWrap
// ═══════════════════════════════════════════

func Test_FuncDetector_GetFuncWrap(t *testing.T) {
	// Arrange
	fn := func() {}
	fw := args.NewFuncWrap.Default(fn)

	// From Map
	m := args.Map{"func": fn}
	fwFromMap := args.FuncDetector.GetFuncWrap(m)

	// From *FuncWrapAny
	fwFromPtr := args.FuncDetector.GetFuncWrap(fw)

	// From raw function
	fwFromFn := args.FuncDetector.GetFuncWrap(fn)

	// Act
	actual := args.Map{
		"fromMapNE": fwFromMap != nil,
		"fromPtrNE": fwFromPtr != nil,
		"fromFnNE":  fwFromFn != nil,
	}

	// Assert
	expected := args.Map{
		"fromMapNE": true,
		"fromPtrNE": true,
		"fromFnNE": true,
	}
	expected.ShouldBeEqual(t, 0, "funcDetector returns correct value -- GetFuncWrap", actual)
}

// ═══════════════════════════════════════════
// FuncMap — all methods
// ═══════════════════════════════════════════

func Test_FuncMap_BasicMethods(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}
	fn := func(a, b int) int { return a + b }
	fm.Add(fn)

	// Act
	actual := args.Map{
		"isEmpty":    fm.IsEmpty(),
		"length":     fm.Length(),
		"count":      fm.Count(),
		"hasAnyItem": fm.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": false, "length": 1, "count": 1, "hasAnyItem": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- basic", actual)
}

func Test_FuncMap_HasAndGet(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	fm := args.FuncMap{}
	fm.Add(fn)
	name := fm.GetPascalCaseFuncName("test")

	// get a function name we know exists
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	// Act
	actual := args.Map{
		"has":        fm.Has(knownName),
		"contains":   fm.IsContains(knownName),
		"getNotNil":  fm.Get(knownName) != nil,
		"getMissing": fm.Get("nonexistent") == nil,
		"hasMissing": fm.Has("nonexistent"),
		"pascalNE":   name != "",
	}

	// Assert
	expected := args.Map{
		"has": true, "contains": true,
		"getNotNil": true, "getMissing": true,
		"hasMissing": false, "pascalNE": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- has and get", actual)
}

func Test_FuncMap_Adds(t *testing.T) {
	// Arrange
	fn1 := func() {}
	fn2 := func(x int) int { return x }
	fm := args.FuncMap{}
	fm.Adds(fn1, fn2)

	// Act
	actual := args.Map{"len": fm.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- adds", actual)
}

func Test_FuncMap_AddsEmpty(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}
	fm.Adds()

	// Act
	actual := args.Map{"len": fm.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "FuncMap returns empty -- adds empty", actual)
}

func Test_FuncMap_IsValidInvalid(t *testing.T) {
	// Arrange
	fn := func() {}
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	// Act
	actual := args.Map{
		"isValid":     fm.IsValidFuncOf(knownName),
		"isInvalid":   fm.IsInvalidFunc(knownName),
		"misValid":    fm.IsValidFuncOf("missing"),
		"misInvalid":  fm.IsInvalidFunc("missing"),
	}

	// Assert
	expected := args.Map{
		"isValid": true, "isInvalid": false,
		"misValid": false, "misInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns error -- isValid/isInvalid", actual)
}

func Test_FuncMap_FuncInfo(t *testing.T) {
	// Arrange
	fn := func(a int) string { return "" }
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	// Act
	actual := args.Map{
		"pkgPath":    fm.PkgPath(knownName) != "",
		"pkgName":    fm.PkgNameOnly(knownName) != "",
		"invokeName": fm.FuncDirectInvokeName(knownName) != "",
		"argsCount":  fm.ArgsCount(knownName),
		"argsLen":    fm.ArgsLength(knownName),
		"retLen":     fm.ReturnLength(knownName),
		"misPkg":     fm.PkgPath("missing"),
		"misPkgName": fm.PkgNameOnly("missing"),
		"misInvoke":  fm.FuncDirectInvokeName("missing"),
		"misArgs":    fm.ArgsCount("missing"),
		"misRet":     fm.ReturnLength("missing"),
	}

	// Assert
	expected := args.Map{
		"pkgPath": true, "pkgName": true, "invokeName": true,
		"argsCount": 1, "argsLen": 1, "retLen": 1,
		"misPkg": "", "misPkgName": "", "misInvoke": "",
		"misArgs": 0, "misRet": 0,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- func info", actual)
}

func Test_FuncMap_IsPublicPrivate(t *testing.T) {
	// Arrange
	fn := func() {}
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	// Act
	actual := args.Map{
		"misPub":  fm.IsPublicMethod("missing"),
		"misPriv": fm.IsPrivateMethod("missing"),
	}
	// Only test missing — function literal naming is runtime-dependent

	// Assert
	expected := args.Map{
		"misPub": false, "misPriv": false,
	}
	_ = knownName
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- public/private", actual)
}

func Test_FuncMap_GetType(t *testing.T) {
	// Arrange
	fn := func() {}
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	// Act
	actual := args.Map{
		"typeNotNil": fm.GetType(knownName) != nil,
		"misType":    fm.GetType("missing") == nil,
		"outArgs":    len(fm.GetOutArgsTypes(knownName)),
		"inArgs":     len(fm.GetInArgsTypes(knownName)),
		"inArgNames": len(fm.GetInArgsTypesNames(knownName)),
		"misOut":     len(fm.GetOutArgsTypes("missing")),
		"misIn":      len(fm.GetInArgsTypes("missing")),
		"misInNames": len(fm.GetInArgsTypesNames("missing")),
	}

	// Assert
	expected := args.Map{
		"typeNotNil": true, "misType": true,
		"outArgs": 0, "inArgs": 0, "inArgNames": 0,
		"misOut": 0, "misIn": 0, "misInNames": 0,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- getType", actual)
}

func Test_FuncMap_Invoke(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	results, err := fm.Invoke(knownName, 3, 4)
	_, misErr := fm.Invoke("missing")
	// VoidCall passes no args to a 2-arg func — expect error
	voidCallRes, voidErr := fm.VoidCall(knownName)

	// Act
	actual := args.Map{
		"errNil":   err == nil,
		"resLen":   len(results),
		"misErr":   misErr != nil,
		"voidErr":  voidErr != nil,
		"voidLen":  len(voidCallRes),
	}

	// Assert
	expected := args.Map{
		"errNil": true, "resLen": 1,
		"misErr": true, "voidErr": true, "voidLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- invoke", actual)
}

func Test_FuncMap_VoidCallNoReturn(t *testing.T) {
	// Arrange
	fn := func() {}
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	// VoidCallNoReturn on a zero-arg func should succeed
	err := fm.VoidCallNoReturn(knownName)
	misErr := fm.VoidCallNoReturn("missing")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"misErr": misErr != nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"misErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- VoidCallNoReturn", actual)
}

func Test_FuncMap_ValidationError(t *testing.T) {
	// Arrange
	fn := func() {}
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	validErr := fm.ValidationError(knownName)
	misErr := fm.ValidationError("missing")

	// Act
	actual := args.Map{
		"validErr": validErr == nil,
		"misErr":   misErr != nil,
	}

	// Assert
	expected := args.Map{
		"validErr": true,
		"misErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns error -- validationError", actual)
}

func Test_FuncMap_InvalidError(t *testing.T) {
	// Arrange
	emptyFm := args.FuncMap{}
	fn := func() {}
	fm := args.FuncMap{}
	fm.Add(fn)

	// Act
	actual := args.Map{
		"emptyErr": emptyFm.InvalidError() != nil,
		"nonEmpty": fm.InvalidError() == nil,
	}

	// Assert
	expected := args.Map{
		"emptyErr": true,
		"nonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns error -- invalidError", actual)
}

func Test_FuncMap_InvalidErrorByName(t *testing.T) {
	// Arrange
	emptyFm := args.FuncMap{}
	fn := func() {}
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	// Act
	actual := args.Map{
		"emptyErr":  emptyFm.InvalidErrorByName("any") != nil,
		"foundNil":  fm.InvalidErrorByName(knownName) == nil,
		"misErr":    fm.InvalidErrorByName("missing") != nil,
	}

	// Assert
	expected := args.Map{
		"emptyErr": true, "foundNil": true, "misErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns error -- invalidErrorByName", actual)
}

func Test_FuncMap_GetEmptyPascalCase(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}
	name := fm.GetPascalCaseFuncName("test")

	// Act
	actual := args.Map{"name": name}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "FuncMap returns empty -- empty getPascalCase", actual)
}

func Test_FuncMap_VerifyArgs(t *testing.T) {
	// Arrange
	fn := func(a int) int { return a }
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	_, misInErr := fm.VerifyInArgs("missing", []any{1})
	_, misOutErr := fm.VerifyOutArgs("missing", []any{1})

	// Act
	actual := args.Map{
		"misInErr":  misInErr != nil,
		"misOutErr": misOutErr != nil,
	}

	// Assert
	expected := args.Map{
		"misInErr": true, "misOutErr": true,
	}
	_ = knownName
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- verifyArgs missing", actual)
}

func Test_FuncMap_ValidateMethodArgs(t *testing.T) {
	// Arrange
	fn := func(a int) int { return a }
	fm := args.FuncMap{}
	fm.Add(fn)

	misErr := fm.ValidateMethodArgs("missing", []any{1})

	// Act
	actual := args.Map{"misErr": misErr != nil}

	// Assert
	expected := args.Map{"misErr": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns non-empty -- validateMethodArgs missing", actual)
}

func Test_FuncMap_GetFirstResponseOfInvoke(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	resp, err := fm.GetFirstResponseOfInvoke(knownName, 3, 4)
	_, misErr := fm.GetFirstResponseOfInvoke("missing")

	// Act
	actual := args.Map{
		"resp":   resp,
		"errNil": err == nil,
		"misErr": misErr != nil,
	}

	// Assert
	expected := args.Map{
		"resp": 7,
		"errNil": true,
		"misErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- getFirstResponse", actual)
}

func Test_FuncMap_InvokeResultOfIndex(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	fm := args.FuncMap{}
	fm.Add(fn)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}

	resp, err := fm.InvokeResultOfIndex(knownName, 0, 3, 4)
	_, misErr := fm.InvokeResultOfIndex("missing", 0)

	// Act
	actual := args.Map{
		"resp":   resp,
		"errNil": err == nil,
		"misErr": misErr != nil,
	}

	// Assert
	expected := args.Map{
		"resp": 7,
		"errNil": true,
		"misErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- invokeResultOfIndex", actual)
}

func Test_FuncMap_InvokeFirstAndError(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	fm := args.FuncMap{}
	fm.Add(fn)

	_, _, misErr := fm.InvokeFirstAndError("missing")

	// Act
	actual := args.Map{"misErr": misErr != nil}

	// Assert
	expected := args.Map{"misErr": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns error -- invokeFirstAndError missing", actual)
}

func Test_FuncMap_EmptyHas(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	actual := args.Map{
		"has": fm.Has("x"),
		"get": fm.Get("x") == nil,
	}

	// Assert
	expected := args.Map{
		"has": false,
		"get": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns empty -- empty has/get", actual)
}

func Test_FuncMap_InArgsVerifyRv_Missing(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}
	fm.Add(func() {})
	_, inErr := fm.InArgsVerifyRv("missing", nil)
	_, outErr := fm.OutArgsVerifyRv("missing", nil)

	// Act
	actual := args.Map{
		"inErr": inErr != nil,
		"outErr": outErr != nil,
	}

	// Assert
	expected := args.Map{
		"inErr": true,
		"outErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- InArgsVerifyRv missing", actual)
}
