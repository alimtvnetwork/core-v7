package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// String type
// ═══════════════════════════════════════════

func Test_String_Basic_FromStringBasic(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{
		"str":       s.String(),
		"len":       s.Length(),
		"count":     s.Count(),
		"asciiLen":  s.AscIILength(),
		"isEmpty":   s.IsEmpty(),
		"hasCh":     s.HasCharacter(),
		"isDef":     s.IsDefined(),
		"isEmptyWS": s.IsEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"str": "hello", "len": 5, "count": 5, "asciiLen": 5,
		"isEmpty": false, "hasCh": true, "isDef": true, "isEmptyWS": false,
	}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- basic", actual)
}

func Test_String_Empty_FromStringBasic(t *testing.T) {
	// Arrange
	s := args.String("")

	// Act
	actual := args.Map{
		"isEmpty": s.IsEmpty(),
		"isEmptyWS": s.IsEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"isEmptyWS": true,
	}
	expected.ShouldBeEqual(t, 0, "String returns empty -- empty", actual)
}

func Test_String_Ops(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{
		"concat":     s.Concat(" world").String(),
		"trimSpace":  args.String("  hi  ").TrimSpace().String(),
		"replaceAll": s.ReplaceAll("l", "r").String(),
		"dq":         s.DoubleQuote().String() != "",
		"dqq":        s.DoubleQuoteQ().String() != "",
		"sq":         s.SingleQuote().String() != "",
		"vdq":        s.ValueDoubleQuote().String() != "",
		"bytesLen":   len(s.Bytes()),
		"runesLen":   len(s.Runes()),
	}

	// Assert
	expected := args.Map{
		"concat": "hello world", "trimSpace": "hi", "replaceAll": "herro",
		"dq": true, "dqq": true, "sq": true, "vdq": true,
		"bytesLen": 5, "runesLen": 5,
	}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- ops", actual)
}

func Test_String_Split_FromStringBasic(t *testing.T) {
	// Arrange
	s := args.String("a,b,c")
	parts := s.Split(",")

	// Act
	actual := args.Map{"len": len(parts)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Split", actual)
}

func Test_String_Join_FromStringBasic(t *testing.T) {
	// Arrange
	s := args.String("hello")
	joined := s.Join("-", "world", "go")

	// Act
	actual := args.Map{"hasContent": len(joined) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Join", actual)
}

func Test_String_Substring_FromStringBasic(t *testing.T) {
	// Arrange
	s := args.String("hello")
	sub := s.Substring(1, 4)

	// Act
	actual := args.Map{"sub": sub.String()}

	// Assert
	expected := args.Map{"sub": "ell"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Substring", actual)
}

// ═══════════════════════════════════════════
// LeftRight
// ═══════════════════════════════════════════

func Test_LeftRight_Basic_FromStringBasic(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{Left: "hello", Right: 42, Expect: true}

	// Act
	actual := args.Map{
		"first":     lr.FirstItem(),
		"second":    lr.SecondItem(),
		"expected":  lr.Expected(),
		"count":     lr.ArgsCount(),
		"hasFirst":  lr.HasFirst(),
		"hasSecond": lr.HasSecond(),
		"hasLeft":   lr.HasLeft(),
		"hasRight":  lr.HasRight(),
		"hasExpect": lr.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"first": "hello", "second": 42, "expected": true, "count": 2,
		"hasFirst": true, "hasSecond": true, "hasLeft": true, "hasRight": true,
		"hasExpect": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- basic", actual)
}

func Test_LeftRight_Slice(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{Left: "a", Right: 1}
	s := lr.Slice()
	s2 := lr.Slice() // cached

	// Act
	actual := args.Map{
		"len": len(s),
		"cached": len(s2) == len(s),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"cached": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Slice", actual)
}

func Test_LeftRight_GetByIndex_FromStringBasic(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{Left: "a", Right: 1}

	// Act
	actual := args.Map{
		"idx0": lr.GetByIndex(0),
		"idx1": lr.GetByIndex(1),
	}

	// Assert
	expected := args.Map{
		"idx0": "a",
		"idx1": 1,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- GetByIndex", actual)
}

func Test_LeftRight_String_FromStringBasic(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{Left: "a", Right: 1}

	// Act
	actual := args.Map{"hasContent": len(lr.String()) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- String", actual)
}

func Test_LeftRight_Clone_FromStringBasic(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{Left: "a", Right: 1, Expect: "exp"}
	cloned := lr.Clone()

	// Act
	actual := args.Map{
		"left": cloned.Left,
		"right": cloned.Right,
	}

	// Assert
	expected := args.Map{
		"left": "a",
		"right": 1,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Clone", actual)
}

func Test_LeftRight_Args(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{Left: "a", Right: 1}
	a1 := lr.Args(1)
	a2 := lr.Args(2)
	va := lr.ValidArgs()

	// Act
	actual := args.Map{
		"a1Len": len(a1),
		"a2Len": len(a2),
		"vaLen": len(va),
	}

	// Assert
	expected := args.Map{
		"a1Len": 1,
		"a2Len": 2,
		"vaLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Args", actual)
}

func Test_LeftRight_ArgTwo_FromStringBasic(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{Left: "a", Right: 1}
	two := lr.ArgTwo()

	// Act
	actual := args.Map{
		"first": two.First,
		"second": two.Second,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": 1,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- ArgTwo", actual)
}

// ═══════════════════════════════════════════
// Holder — comprehensive
// ═══════════════════════════════════════════

func Test_Holder_AllPositional(t *testing.T) {
	// Arrange
	h := &args.Holder[func() string]{
		First: "a", Second: "b", Third: "c",
		Fourth: "d", Fifth: "e", Sixth: "f",
		Expect: "exp",
	}

	// Act
	actual := args.Map{
		"first": h.FirstItem(), "second": h.SecondItem(), "third": h.ThirdItem(),
		"fourth": h.FourthItem(), "fifth": h.FifthItem(), "sixth": h.SixthItem(),
		"expected": h.Expected(), "count": h.ArgsCount(),
		"hasFirst": h.HasFirst(), "hasSecond": h.HasSecond(), "hasThird": h.HasThird(),
		"hasFourth": h.HasFourth(), "hasFifth": h.HasFifth(), "hasSixth": h.HasSixth(),
		"hasExpect": h.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"fourth": "d", "fifth": "e", "sixth": "f",
		"expected": "exp", "count": 7,
		"hasFirst": true, "hasSecond": true, "hasThird": true,
		"hasFourth": true, "hasFifth": true, "hasSixth": true,
		"hasExpect": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- all positional", actual)
}

func Test_Holder_Args_FromStringBasic(t *testing.T) {
	// Arrange
	h := &args.Holder[any]{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f"}
	a1 := h.Args(1)
	a3 := h.Args(3)
	a6 := h.Args(6)
	va := h.ValidArgs()

	// Act
	actual := args.Map{
		"a1Len": len(a1), "a3Len": len(a3), "a6Len": len(a6), "vaLen": len(va),
	}

	// Assert
	expected := args.Map{
		"a1Len": 1,
		"a3Len": 3,
		"a6Len": 6,
		"vaLen": 6,
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- Args", actual)
}

func Test_Holder_ArgTwo(t *testing.T) {
	// Arrange
	h := &args.Holder[any]{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}
	two := h.ArgTwo()
	three := h.ArgThree()
	four := h.ArgFour()
	five := h.ArgFive()

	// Act
	actual := args.Map{
		"twoFirst": two.First, "threeThird": three.Third,
		"fourFourth": four.Fourth, "fiveFifth": five.Fifth,
	}

	// Assert
	expected := args.Map{
		"twoFirst": "a", "threeThird": "c", "fourFourth": "d", "fiveFifth": "e",
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- ArgTwo/Three/Four/Five", actual)
}

func Test_Holder_WithFunc(t *testing.T) {
	// Arrange
	fn := func(a, b string) string { return a + b }
	h := &args.Holder[func(string, string) string]{First: "a", Second: "b", WorkFunc: fn}

	// Act
	actual := args.Map{
		"hasFunc":  h.HasFunc(),
		"funcName": h.GetFuncName() != "",
		"getFunc":  h.GetWorkFunc() != nil,
	}

	// Assert
	expected := args.Map{
		"hasFunc": true,
		"funcName": true,
		"getFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder returns non-empty -- with func", actual)
}

func Test_Holder_Invoke(t *testing.T) {
	// Arrange
	fn := func(a, b string) string { return a + b }
	h := &args.Holder[func(string, string) string]{First: "hello", Second: " world", WorkFunc: fn}
	results, err := h.Invoke("hello", " world")

	// Act
	actual := args.Map{
		"noErr":  err == nil,
		"result": results[0],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"result": "hello world",
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- Invoke", actual)
}

func Test_Holder_Slice(t *testing.T) {
	// Arrange
	h := &args.Holder[any]{First: "a", Second: "b"}
	s := h.Slice()

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- Slice", actual)
}

func Test_Holder_String(t *testing.T) {
	// Arrange
	h := &args.Holder[any]{First: "a"}

	// Act
	actual := args.Map{"hasContent": len(h.String()) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- String", actual)
}

// ═══════════════════════════════════════════
// FuncWrap — typed helpers
// ═══════════════════════════════════════════

func Test_FuncWrap_TypedHelpers(t *testing.T) {
	// Arrange
	boolFn := func() bool { return true }
	errFn := func() error { return nil }
	strFn := func() string { return "hi" }
	voidFn := func() {}
	valErrFn := func() (int, error) { return 1, nil }

	boolFW := args.NewFuncWrap.Default(boolFn)
	errFW := args.NewFuncWrap.Default(errFn)
	strFW := args.NewFuncWrap.Default(strFn)
	voidFW := args.NewFuncWrap.Default(voidFn)
	valErrFW := args.NewFuncWrap.Default(valErrFn)

	// Act
	actual := args.Map{
		"isBool":     boolFW.IsBoolFunc(),
		"isErr":      errFW.IsErrorFunc(),
		"isStr":      strFW.IsStringFunc(),
		"isVoid":     voidFW.IsVoidFunc(),
		"isAny":      strFW.IsAnyFunc(),
		"isValErr":   valErrFW.IsValueErrorFunc(),
		"isAnyErr":   valErrFW.IsAnyErrorFunc(),
	}

	// Assert
	expected := args.Map{
		"isBool": true, "isErr": true, "isStr": true,
		"isVoid": true, "isAny": true, "isValErr": true, "isAnyErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- typed helpers", actual)
}

func Test_FuncWrap_InvokeAsBool(t *testing.T) {
	// Arrange
	fn := func() bool { return true }
	fw := args.NewFuncWrap.Default(fn)
	b, err := fw.InvokeAsBool()

	// Act
	actual := args.Map{
		"b": b,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"b": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InvokeAsBool", actual)
}

func Test_FuncWrap_InvokeAsString(t *testing.T) {
	// Arrange
	fn := func() string { return "hello" }
	fw := args.NewFuncWrap.Default(fn)
	s, err := fw.InvokeAsString()

	// Act
	actual := args.Map{
		"s": s,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"s": "hello",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InvokeAsString", actual)
}

func Test_FuncWrap_InvokeAsAny(t *testing.T) {
	// Arrange
	fn := func() int { return 42 }
	fw := args.NewFuncWrap.Default(fn)
	v, err := fw.InvokeAsAny()

	// Act
	actual := args.Map{
		"v": v,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"v": 42,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InvokeAsAny", actual)
}

func Test_FuncWrap_InvokeAsError(t *testing.T) {
	// Arrange
	fn := func() error { return nil }
	fw := args.NewFuncWrap.Default(fn)
	funcErr, procErr := fw.InvokeAsError()

	// Act
	actual := args.Map{
		"funcErr": funcErr == nil,
		"procErr": procErr == nil,
	}

	// Assert
	expected := args.Map{
		"funcErr": true,
		"procErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns error -- InvokeAsError", actual)
}

func Test_FuncWrap_InvokeAsAnyError(t *testing.T) {
	// Arrange
	fn := func() (int, error) { return 42, nil }
	fw := args.NewFuncWrap.Default(fn)
	v, funcErr, procErr := fw.InvokeAsAnyError()

	// Act
	actual := args.Map{
		"v": v,
		"funcErr": funcErr == nil,
		"procErr": procErr == nil,
	}

	// Assert
	expected := args.Map{
		"v": 42,
		"funcErr": true,
		"procErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns error -- InvokeAsAnyError", actual)
}

// ═══════════════════════════════════════════
// FuncWrap — args info
// ═══════════════════════════════════════════

func Test_FuncWrap_ArgsInfo(t *testing.T) {
	// Arrange
	fn := func(a string, b int) (string, error) { return a, nil }
	fw := args.NewFuncWrap.Default(fn)
	inTypes := fw.GetInArgsTypes()
	outTypes := fw.GetOutArgsTypes()
	inNames := fw.GetInArgsTypesNames()
	outNames := fw.GetOutArgsTypesNames()

	// Act
	actual := args.Map{
		"inLen":     len(inTypes),
		"outLen":    len(outTypes),
		"inNames":   len(inNames),
		"outNames":  len(outNames),
		"argsLen":   fw.ArgsLength(),
		"retLen":    fw.ReturnLength(),
	}

	// Assert
	expected := args.Map{
		"inLen": 2, "outLen": 2, "inNames": 2, "outNames": 2,
		"argsLen": 2, "retLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- args info", actual)
}

func Test_FuncWrap_InArgNames(t *testing.T) {
	// Arrange
	fn := func(a string) string { return a }
	fw := args.NewFuncWrap.Default(fn)
	names := fw.InArgNames()
	outNames := fw.OutArgNames()

	// Act
	actual := args.Map{
		"inLen": len(names),
		"outLen": len(outNames),
	}

	// Assert
	expected := args.Map{
		"inLen": 1,
		"outLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InArgNames", actual)
}

func Test_FuncWrap_InArgNamesEachLine(t *testing.T) {
	// Arrange
	fn := func(a, b string) string { return a + b }
	fw := args.NewFuncWrap.Default(fn)
	lines := fw.InArgNamesEachLine()
	outLines := fw.OutArgNamesEachLine()

	// Act
	actual := args.Map{
		"inLen": len(lines) > 0,
		"outLen": len(outLines) > 0,
	}

	// Assert
	expected := args.Map{
		"inLen": true,
		"outLen": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InArgNamesEachLine", actual)
}

func Test_FuncWrap_IsTypeMatches(t *testing.T) {
	// Arrange
	fn := func(a string) int { return 0 }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{
		"inMatch":   fw.IsInTypeMatches("hello"),
		"outMatch":  fw.IsOutTypeMatches(0),
		"inNoMatch": fw.IsInTypeMatches(42),
	}

	// Assert
	expected := args.Map{
		"inMatch": true,
		"outMatch": true,
		"inNoMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- IsTypeMatches", actual)
}

// ═══════════════════════════════════════════
// FuncWrap — validation
// ═══════════════════════════════════════════

func Test_FuncWrap_Validation(t *testing.T) {
	// Arrange
	fn := func(a string) string { return a }
	fw := args.NewFuncWrap.Default(fn)
	nilFW := args.NewFuncWrap.Default(nil)

	// Act
	actual := args.Map{
		"validErr":   fw.ValidationError() == nil,
		"invalidErr": fw.InvalidError() == nil,
		"nilValid":   nilFW.ValidationError() != nil,
		"nilInvalid": nilFW.InvalidError() != nil,
	}

	// Assert
	expected := args.Map{
		"validErr": true, "invalidErr": true,
		"nilValid": true, "nilInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns non-empty -- validation", actual)
}

func Test_FuncWrap_ValidateMethodArgs(t *testing.T) {
	// Arrange
	fn := func(a string) string { return a }
	fw := args.NewFuncWrap.Default(fn)
	noErr := fw.ValidateMethodArgs([]any{"hello"})
	hasErr := fw.ValidateMethodArgs([]any{"hello", "extra"})

	// Act
	actual := args.Map{
		"noErr": noErr == nil,
		"hasErr": hasErr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns non-empty -- ValidateMethodArgs", actual)
}

func Test_FuncWrap_MustBeValid_Panic(t *testing.T) {
	// Arrange
	nilFW := args.NewFuncWrap.Default(nil)
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "FuncWrap panics -- MustBeValid panic", actual)
	}()
	nilFW.MustBeValid()
}

// ═══════════════════════════════════════════
// FuncWrap — invoke methods
// ═══════════════════════════════════════════

func Test_FuncWrap_VoidCall_FromStringBasic(t *testing.T) {
	// Arrange
	called := false
	fn := func() { called = true }
	fw := args.NewFuncWrap.Default(fn)
	_, err := fw.VoidCall()

	// Act
	actual := args.Map{
		"called": called,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"called": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- VoidCall", actual)
}

func Test_FuncWrap_GetFirstResponse(t *testing.T) {
	// Arrange
	fn := func() string { return "first" }
	fw := args.NewFuncWrap.Default(fn)
	first, err := fw.GetFirstResponseOfInvoke()

	// Act
	actual := args.Map{
		"first": first,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"first": "first",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- GetFirstResponseOfInvoke", actual)
}

func Test_FuncWrap_InvokeResultOfIndex_FromStringBasic(t *testing.T) {
	// Arrange
	fn := func() (string, int) { return "a", 1 }
	fw := args.NewFuncWrap.Default(fn)
	r, err := fw.InvokeResultOfIndex(1)

	// Act
	actual := args.Map{
		"r": r,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"r": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InvokeResultOfIndex", actual)
}

func Test_FuncWrap_InvokeMust(t *testing.T) {
	// Arrange
	fn := func() string { return "ok" }
	fw := args.NewFuncWrap.Default(fn)
	results := fw.InvokeMust()

	// Act
	actual := args.Map{"r": results[0]}

	// Assert
	expected := args.Map{"r": "ok"}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InvokeMust", actual)
}

// ═══════════════════════════════════════════
// FuncWrap — nil/invalid return 
// ═══════════════════════════════════════════

func Test_FuncWrap_Invalid_Args(t *testing.T) {
	// Arrange
	nilFW := args.NewFuncWrap.Default(nil)

	// Act
	actual := args.Map{
		"argsCount":   nilFW.ArgsCount(),
		"outCount":    nilFW.OutArgsCount(),
		"inTypes":     len(nilFW.GetInArgsTypes()),
		"outTypes":    len(nilFW.GetOutArgsTypes()),
		"inNames":     len(nilFW.GetInArgsTypesNames()),
		"outNames":    len(nilFW.GetOutArgsTypesNames()),
		"inArgNames":  len(nilFW.InArgNames()),
		"outArgNames": len(nilFW.OutArgNames()),
		"isBool":      nilFW.IsBoolFunc(),
		"isErr":       nilFW.IsErrorFunc(),
		"isStr":       nilFW.IsStringFunc(),
		"isAny":       nilFW.IsAnyFunc(),
		"isValErr":    nilFW.IsValueErrorFunc(),
		"isVoid":      nilFW.IsVoidFunc(),
	}

	// Assert
	expected := args.Map{
		"argsCount": -1, "outCount": -1,
		"inTypes": 0, "outTypes": 0, "inNames": 0, "outNames": 0,
		"inArgNames": 0, "outArgNames": 0,
		"isBool": false, "isErr": false, "isStr": false,
		"isAny": false, "isValErr": false, "isVoid": false,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns error -- invalid args", actual)
}

// ═══════════════════════════════════════════
// FourFunc / FiveFunc / SixFunc
// ═══════════════════════════════════════════

func Test_FourFunc(t *testing.T) {
	// Arrange
	ff := &args.FourFunc[string, int, bool, float64]{
		First: "a", Second: 1, Third: true, Fourth: 3.14,
		WorkFunc: func() string { return "hi" },
	}

	// Act
	actual := args.Map{
		"first":   ff.FirstItem(),
		"hasFunc": ff.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "FourFunc returns correct value -- with args", actual)
}

func Test_FiveFunc(t *testing.T) {
	// Arrange
	ff := &args.FiveFunc[string, int, bool, float64, byte]{
		First: "a", Second: 1, Third: true, Fourth: 3.14, Fifth: byte(5),
		WorkFunc: func() string { return "hi" },
	}

	// Act
	actual := args.Map{
		"fifth":   ff.FifthItem(),
		"hasFunc": ff.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"fifth": byte(5),
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "FiveFunc returns correct value -- with args", actual)
}

func Test_SixFunc(t *testing.T) {
	// Arrange
	sf := &args.SixFunc[string, int, bool, float64, byte, uint]{
		First: "a", Sixth: uint(6),
		WorkFunc: func() string { return "hi" },
	}

	// Act
	actual := args.Map{
		"sixth":   sf.SixthItem(),
		"hasFunc": sf.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"sixth": uint(6),
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "SixFunc returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// DynamicFunc
// ═══════════════════════════════════════════

func Test_DynamicFunc(t *testing.T) {
	// Arrange
	df := &args.DynamicFuncAny{
		Params:   args.Map{"first": "hello"},
		WorkFunc: func() string { return "hi" },
		Expect:   42,
	}

	// Act
	actual := args.Map{
		"first":   df.FirstItem(),
		"expect":  df.Expected(),
		"hasFunc": df.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"first": "hello",
		"expect": 42,
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// Map — additional methods
// ═══════════════════════════════════════════

func Test_Map_WorkFunc_FromStringBasic(t *testing.T) {
	// Arrange
	fn := func() string { return "hello" }
	m := args.Map{
		"func": args.NewFuncWrap.Default(fn),
		"a": 1,
	}
	fw := m.FuncWrap()

	// Act
	actual := args.Map{
		"hasFunc":  m.HasFunc(),
		"fwValid":  fw.IsValid(),
		"count":    m.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"hasFunc": true,
		"fwValid": true,
		"count": 1,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- WorkFunc", actual)
}

func Test_Map_GetTyped(t *testing.T) {
	// Arrange
	m := args.Map{
		"str": "hello",
		"int": 42,
		"bool": true,
		"strs": []string{"a", "b"},
	}
	str, _ := m.GetAsString("str")
	intVal, _ := m.GetAsInt("int")
	boolVal, _ := m.GetAsBool("bool")
	strs, _ := m.GetAsStrings("strs")

	// Act
	actual := args.Map{
		"str":     str,
		"int":     intVal,
		"bool":    boolVal,
		"strs":    len(strs),
		"defBool": m.GetAsBoolDefault("missing", true),
	}

	// Assert
	expected := args.Map{
		"str": "hello", "int": 42, "bool": true, "strs": 2, "defBool": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetTyped", actual)
}

func Test_Map_SortedKeysMust_FromStringBasic(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
		"c": 3,
	}
	keys := m.SortedKeysMust()

	// Act
	actual := args.Map{
		"first": keys[0],
		"last": keys[2],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- SortedKeysMust", actual)
}
