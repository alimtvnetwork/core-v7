package args

import (
	"strings"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Dynamic — nil and uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func sampleStringFunc(s string) string { return strings.ToUpper(s) }

func Test_I13_Dynamic_NilBranches(t *testing.T) {
	var d *Dynamic[string]
	if d.ArgsCount() != 0 {
		t.Fatal("nil ArgsCount should be 0")
	}
	if d.GetWorkFunc() != nil {
		t.Fatal("nil GetWorkFunc should be nil")
	}
	if d.HasFirst() {
		t.Fatal("nil HasFirst should be false")
	}
}

func Test_I13_Dynamic_GetWorkFunc(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{"func": sampleStringFunc},
	}
	if d.GetWorkFunc() == nil {
		t.Fatal("expected work func")
	}
}

func Test_I13_Dynamic_HasFunc(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{"func": sampleStringFunc},
	}
	if !d.HasFunc() {
		t.Fatal("expected HasFunc true")
	}
	_ = d.GetFuncName()
}

func Test_I13_Dynamic_InvokeMust(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{"func": sampleStringFunc, "first": "hello"},
	}
	results := d.InvokeMust("hello")
	if len(results) == 0 {
		t.Fatal("expected results")
	}
}

func Test_I13_Dynamic_FuncWrap(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{"func": sampleStringFunc},
	}
	fw := d.FuncWrap()
	if fw == nil {
		t.Fatal("expected func wrap")
	}
}

func Test_I13_Dynamic_Items(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{
			"first":  "a",
			"second": "b",
			"third":  "c",
			"fourth": "d",
			"fifth":  "e",
			"sixth":  "f",
		},
		Expect: "exp",
	}
	_ = d.FirstItem()
	_ = d.SecondItem()
	_ = d.ThirdItem()
	_ = d.FourthItem()
	_ = d.FifthItem()
	_ = d.SixthItem()
	_ = d.Expected()
	if !d.HasExpect() {
		t.Fatal("expected HasExpect")
	}
}

func Test_I13_Dynamic_GetVariants(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{
			"key1":   "val1",
			"actual": "act",
		},
	}
	_, _ = d.GetLowerCase("KEY1")
	_ = d.GetDirectLower("KEY1")
	_ = d.Actual()
	_ = d.Arrange()
	_, _ = d.Get("key1")
	v, ok := d.GetAsInt("key1")
	_ = v
	_ = ok
	_ = d.GetAsIntDefault("key1", 0)
	_, _ = d.GetAsString("key1")
	_ = d.GetAsStringDefault("key1")
	_, _ = d.GetAsStrings("key1")
	_, _ = d.GetAsAnyItems("key1")
}

func Test_I13_Dynamic_HasDefined(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{"a": "x"},
	}
	if !d.HasDefined("a") {
		t.Fatal("expected HasDefined")
	}
	if !d.Has("a") {
		t.Fatal("expected Has")
	}
	if d.HasDefinedAll("a", "missing") {
		t.Fatal("expected false for missing key")
	}
	if !d.HasDefinedAll("a") {
		t.Fatal("expected true for present key")
	}
	if d.IsKeyInvalid("a") {
		t.Fatal("expected valid key")
	}
	if !d.IsKeyMissing("missing") {
		t.Fatal("expected missing key")
	}
}

func Test_I13_Dynamic_ValidArgsAndSlice(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{"1": "a", "2": "b"},
		Expect: "exp",
	}
	va := d.ValidArgs()
	if len(va) == 0 {
		t.Fatal("expected valid args")
	}
	args := d.Args("1", "2")
	if len(args) != 2 {
		t.Fatal("expected 2 args")
	}
	s := d.Slice()
	if len(s) == 0 {
		t.Fatal("expected slice")
	}
	str := d.String()
	if str == "" {
		t.Fatal("expected string")
	}
	// cached path
	_ = d.String()

	_ = d.AsArgsMapper()
	_ = d.AsArgFuncNameContractsBinder()
	_ = d.AsArgBaseContractsBinder()
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicFunc — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_DynamicFunc_Branches(t *testing.T) {
	df := &DynamicFunc[func(string) string]{
		Params:   Map{"first": "hello"},
		WorkFunc: sampleStringFunc,
		Expect:   "HELLO",
	}
	if df.ArgsCount() != 1 {
		t.Fatal("expected 1 arg")
	}
	if df.GetWorkFunc() == nil {
		t.Fatal("expected work func")
	}
	if df.Length() != 1 {
		t.Fatal("expected length 1")
	}
	if !df.HasFirst() {
		t.Fatal("expected has first")
	}
	_ = df.GetByIndex(0)
	_ = df.GetByIndex(999) // out of bounds
	_ = df.FirstItem()
	_ = df.SecondItem()
	_ = df.ThirdItem()
	_ = df.FourthItem()
	_ = df.FifthItem()
	_ = df.SixthItem()
	_ = df.Expected()
	if !df.HasFunc() {
		t.Fatal("expected has func")
	}
	if !df.HasExpect() {
		t.Fatal("expected has expect")
	}
	_ = df.GetFuncName()
	_ = df.FuncWrap()
	_, _ = df.Invoke("hello")
	results := df.InvokeMust("hello")
	if len(results) == 0 {
		t.Fatal("expected results")
	}
	_, _ = df.InvokeWithValidArgs()
	_ = df.ValidArgs()
	_ = df.Args("first")
	_ = df.Slice()
	_ = df.String()
	// cached
	_ = df.String()

	_ = df.When()
	_ = df.Title()
	_, _ = df.GetLowerCase("FIRST")
	_ = df.GetDirectLower("FIRST")
	_ = df.Actual()
	_ = df.Arrange()
	_, _ = df.Get("first")
	_, _ = df.GetAsInt("first")
	_, _ = df.GetAsString("first")
	_, _ = df.GetAsStrings("first")
	_, _ = df.GetAsAnyItems("first")

	_ = df.AsArgsMapper()
	_ = df.AsArgFuncNameContractsBinder()
	_ = df.AsArgBaseContractsBinder()
}

func Test_I13_DynamicFunc_NilBranches(t *testing.T) {
	var df *DynamicFunc[any]
	if df.ArgsCount() != 0 {
		t.Fatal("nil ArgsCount should be 0")
	}
	if df.Length() != 0 {
		t.Fatal("nil Length should be 0")
	}
	_, ok := df.Get("x")
	if ok {
		t.Fatal("nil Get should be invalid")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FuncWrap — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_FuncWrap_NewTypedFuncWrap(t *testing.T) {
	fw := NewTypedFuncWrap(sampleStringFunc)
	if fw.IsInvalid() {
		t.Fatal("expected valid")
	}
	if !fw.IsValid() {
		t.Fatal("expected valid")
	}
	if !fw.HasValidFunc() {
		t.Fatal("expected has valid func")
	}
	_ = fw.GetFuncName()
	_ = fw.GetPascalCaseFuncName()
	_ = fw.PkgPath()
	_ = fw.PkgNameOnly()
	_ = fw.FuncDirectInvokeName()
	_ = fw.GetType()
	_ = fw.IsPublicMethod()
	_ = fw.IsPrivateMethod()

	// equality
	fw2 := NewTypedFuncWrap(sampleStringFunc)
	if fw.IsNotEqual(fw2) {
		t.Fatal("expected equal")
	}
	if !fw.IsEqualValue(*fw2) {
		t.Fatal("expected equal value")
	}

	// nil func
	fwNil := NewTypedFuncWrap[any](nil)
	if !fwNil.IsInvalid() {
		t.Fatal("expected invalid for nil")
	}

	// non-func
	fwStr := NewTypedFuncWrap("not a func")
	if !fwStr.IsInvalid() {
		t.Fatal("expected invalid for non-func")
	}
}

func Test_I13_FuncWrap_InvokeTypedHelpers(t *testing.T) {
	boolFn := func(s string) bool { return s == "yes" }
	fw := NewTypedFuncWrap(boolFn)
	if !fw.IsBoolFunc() {
		t.Fatal("expected bool func")
	}
	val, err := fw.InvokeAsBool("yes")
	if err != nil || !val {
		t.Fatal("expected true")
	}

	strFn := func(s string) string { return s + "!" }
	fws := NewTypedFuncWrap(strFn)
	if !fws.IsStringFunc() {
		t.Fatal("expected string func")
	}
	sv, _ := fws.InvokeAsString("hi")
	if sv != "hi!" {
		t.Fatal("expected hi!")
	}

	anyFn := func() string { return "x" }
	fwa := NewTypedFuncWrap(anyFn)
	if !fwa.IsAnyFunc() {
		t.Fatal("expected any func")
	}
	av, _ := fwa.InvokeAsAny()
	if av != "x" {
		t.Fatal("expected x")
	}

	voidFn := func() {}
	fwv := NewTypedFuncWrap(voidFn)
	if !fwv.IsVoidFunc() {
		t.Fatal("expected void func")
	}
	_, _ = fwv.VoidCall()
}

func Test_I13_FuncWrap_Validation(t *testing.T) {
	fw := NewTypedFuncWrap(sampleStringFunc)
	fw.MustBeValid()
	if fw.ValidationError() != nil {
		t.Fatal("expected no validation error")
	}
	if fw.InvalidError() != nil {
		t.Fatal("expected no invalid error")
	}
	err := fw.ValidateMethodArgs([]any{"hello"})
	if err != nil {
		t.Fatal("expected no error:", err)
	}
	// wrong arg count
	err2 := fw.ValidateMethodArgs([]any{"a", "b"})
	if err2 == nil {
		t.Fatal("expected arg count mismatch error")
	}
}

func Test_I13_FuncWrap_Args(t *testing.T) {
	fw := NewTypedFuncWrap(sampleStringFunc)
	if fw.ArgsCount() != 1 {
		t.Fatal("expected 1 arg")
	}
	inTypes := fw.GetInArgsTypes()
	if len(inTypes) != 1 {
		t.Fatal("expected 1 in type")
	}
	outTypes := fw.GetOutArgsTypes()
	if len(outTypes) != 1 {
		t.Fatal("expected 1 out type")
	}
	_ = fw.GetInArgsTypesNames()
	_ = fw.GetOutArgsTypesNames()
	_ = fw.InArgNames()
	_ = fw.InArgNamesEachLine()
	_ = fw.OutArgNames()
	_ = fw.OutArgNamesEachLine()
	if !fw.IsInTypeMatches("hello") {
		t.Fatal("expected in type match")
	}
	if !fw.IsOutTypeMatches("result") {
		t.Fatal("expected out type match")
	}
}

func Test_I13_FuncWrap_NilReceiver(t *testing.T) {
	var fw *FuncWrap[any]
	if fw.GetFuncName() != "" {
		t.Fatal("nil should return empty name")
	}
	if fw.GetPascalCaseFuncName() != "" {
		t.Fatal("nil should return empty pascal name")
	}
	if fw.HasValidFunc() {
		t.Fatal("nil should not have valid func")
	}
	if !fw.IsInvalid() {
		t.Fatal("nil should be invalid")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FuncMap — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_FuncMap_Ops(t *testing.T) {
	fm := NewFuncWrap.Map(sampleStringFunc)
	if fm.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	if !fm.HasAnyItem() {
		t.Fatal("expected has items")
	}
	name := ""
	for k := range fm {
		name = k
		break
	}
	if !fm.Has(name) {
		t.Fatal("expected has")
	}
	if !fm.IsContains(name) {
		t.Fatal("expected contains")
	}
	g := fm.Get(name)
	if g == nil {
		t.Fatal("expected get")
	}
	_ = fm.GetPascalCaseFuncName(name)
	if !fm.IsValidFuncOf(name) {
		t.Fatal("expected valid")
	}
	if fm.IsInvalidFunc(name) {
		t.Fatal("expected not invalid")
	}
	_ = fm.PkgPath(name)
	_ = fm.PkgNameOnly(name)
	_ = fm.FuncDirectInvokeName(name)
	_ = fm.ArgsCount(name)
	_ = fm.ArgsLength(name)
	_ = fm.ReturnLength(name)
	_ = fm.IsPublicMethod(name)
	_ = fm.IsPrivateMethod(name)
	_ = fm.GetType(name)
	_ = fm.GetOutArgsTypes(name)
	_ = fm.GetInArgsTypes(name)
	_ = fm.GetInArgsTypesNames(name)

	_, _ = fm.VerifyInArgs(name, []any{"x"})
	_, _ = fm.VerifyOutArgs(name, []any{"x"})
	_, _ = fm.InArgsVerifyRv(name, nil)
	_, _ = fm.OutArgsVerifyRv(name, nil)
	_ = fm.VoidCallNoReturn(name, "x")
	_ = fm.ValidationError(name)
	_ = fm.ValidateMethodArgs(name, []any{"x"})
	_, _ = fm.Invoke(name, "hello")
	_ = fm.InvokeMust(name, "hello")
	_, _ = fm.VoidCall(name)
	_, _ = fm.GetFirstResponseOfInvoke(name, "test")
	_, _ = fm.InvokeResultOfIndex(name, 0, "test")
	_ = fm.InvalidError()

	// not found branches
	if fm.Has("nonexistent") {
		t.Fatal("expected not found")
	}
	_ = fm.Get("nonexistent")
	_ = fm.ArgsCount("nonexistent")
}

func Test_I13_FuncMap_Add(t *testing.T) {
	fm := FuncMap{}
	fm.Add(sampleStringFunc)
	if fm.IsEmpty() {
		t.Fatal("expected non-empty after Add")
	}
	fm.Adds(sampleStringFunc)
}

// ══════════════════════════════════════════════════════════════════════════════
// Holder — uncovered InvokeWithValidArgs, InvokeArgs, Args
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_Holder_Branches(t *testing.T) {
	h := &Holder[func(string) string]{
		First:    "hello",
		Second:   "world",
		Third:    "!",
		Fourth:   "a",
		Fifth:    "b",
		Sixth:    "c",
		WorkFunc: sampleStringFunc,
		Expect:   "HELLO",
	}
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
	if !h.HasFirst() {
		t.Fatal("expected has first")
	}
	if !h.HasSixth() {
		t.Fatal("expected has sixth")
	}
	if !h.HasFunc() {
		t.Fatal("expected has func")
	}
	if !h.HasExpect() {
		t.Fatal("expected has expect")
	}
	_ = h.GetFuncName()
	_ = h.FuncWrap()
	_, _ = h.InvokeWithValidArgs()
	_, _ = h.InvokeArgs(1)
	_ = h.ValidArgs()
	_ = h.Args(6)
	_ = h.Slice()
	_ = h.GetByIndex(0)
	_ = h.String()
	_ = h.AsSixthParameter()
	_ = h.AsArgFuncContractsBinder()
}

// ══════════════════════════════════════════════════════════════════════════════
// OneFunc — uncovered InvokeArgs, Args branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_OneFunc_Branches(t *testing.T) {
	of := &OneFunc[string]{
		First:    "hello",
		WorkFunc: sampleStringFunc,
		Expect:   "HELLO",
	}
	_ = of.FirstItem()
	_ = of.Expected()
	_ = of.ArgTwo()
	if !of.HasFirst() {
		t.Fatal("expected has first")
	}
	if !of.HasFunc() {
		t.Fatal("expected has func")
	}
	_, _ = of.InvokeWithValidArgs()
	_, _ = of.InvokeArgs(1)
	_ = of.ValidArgs()
	_ = of.Args(1)
	_ = of.Slice()
	_ = of.GetByIndex(0)
	_ = of.String()
	_ = of.LeftRight()
	_ = of.AsOneFuncParameter()
	_ = of.AsArgFuncContractsBinder()
	_ = of.AsArgBaseContractsBinder()
}

// ══════════════════════════════════════════════════════════════════════════════
// TwoFunc — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_TwoFunc_Branches(t *testing.T) {
	tf := &TwoFunc[string, string]{
		First:    "a",
		Second:   "b",
		WorkFunc: sampleStringFunc,
		Expect:   "A",
	}
	_, _ = tf.InvokeWithValidArgs()
	_, _ = tf.InvokeArgs(2)
	_ = tf.ValidArgs()
	_ = tf.Args(2)
	_ = tf.Slice()
	_ = tf.String()
	_ = tf.LeftRight()
	_ = tf.AsTwoFuncParameter()
}

// ══════════════════════════════════════════════════════════════════════════════
// ThreeFunc — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_ThreeFunc_Branches(t *testing.T) {
	tf := &ThreeFunc[string, string, string]{
		First:    "a",
		Second:   "b",
		Third:    "c",
		WorkFunc: sampleStringFunc,
		Expect:   "A",
	}
	_, _ = tf.InvokeWithValidArgs()
	_, _ = tf.InvokeArgs(3)
	_ = tf.ValidArgs()
	_ = tf.Args(3)
	_ = tf.Slice()
	_ = tf.String()
	_ = tf.LeftRight()
	_ = tf.AsThreeFuncParameter()
}

// ══════════════════════════════════════════════════════════════════════════════
// FourFunc — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_FourFunc_Branches(t *testing.T) {
	ff := &FourFunc[string, string, string, string]{
		First:    "a",
		Second:   "b",
		Third:    "c",
		Fourth:   "d",
		WorkFunc: sampleStringFunc,
		Expect:   "A",
	}
	_, _ = ff.InvokeWithValidArgs()
	_, _ = ff.InvokeArgs(4)
	_ = ff.ValidArgs()
	_ = ff.Args(4)
	_ = ff.Slice()
	_ = ff.String()
	_ = ff.AsFourFuncParameter()
}

// ══════════════════════════════════════════════════════════════════════════════
// FiveFunc — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_FiveFunc_Branches(t *testing.T) {
	ff := &FiveFunc[string, string, string, string, string]{
		First:    "a",
		Second:   "b",
		Third:    "c",
		Fourth:   "d",
		Fifth:    "e",
		WorkFunc: sampleStringFunc,
		Expect:   "A",
	}
	_, _ = ff.InvokeWithValidArgs()
	_, _ = ff.InvokeArgs(5)
	_ = ff.ValidArgs()
	_ = ff.Args(5)
	_ = ff.Slice()
	_ = ff.String()
	_ = ff.AsFifthFuncParameter()
}

// ══════════════════════════════════════════════════════════════════════════════
// SixFunc — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_SixFunc_Branches(t *testing.T) {
	sf := &SixFunc[string, string, string, string, string, string]{
		First:    "a",
		Second:   "b",
		Third:    "c",
		Fourth:   "d",
		Fifth:    "e",
		Sixth:    "f",
		WorkFunc: sampleStringFunc,
		Expect:   "A",
	}
	_, _ = sf.InvokeWithValidArgs()
	_, _ = sf.InvokeArgs(6)
	_ = sf.ValidArgs()
	_ = sf.Args(6)
	_ = sf.Slice()
	_ = sf.String()
	_ = sf.AsSixthFuncParameter()
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_LeftRight_Branches(t *testing.T) {
	lr := &LeftRight[string, string]{
		Left:   "a",
		Right:  "b",
		Expect: "exp",
	}
	if !lr.HasLeft() {
		t.Fatal("expected has left")
	}
	if !lr.HasRight() {
		t.Fatal("expected has right")
	}
	_ = lr.ValidArgs()
	_ = lr.Args(2)
	_ = lr.Slice()
	_ = lr.String()
	_ = lr.Clone()
	_ = lr.AsTwoParameter()
	_ = lr.AsArgBaseContractsBinder()
}

// ══════════════════════════════════════════════════════════════════════════════
// String type — exercise
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_String_Branches(t *testing.T) {
	s := String("hello")
	_ = s.Concat(" world")
	_ = s.Join(",", "a", "b")
	_ = s.Split(",")
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
	_ = s.ReplaceAll("hello", "bye")
	_ = s.TrimReplaceMap(map[string]string{"hello": "bye"})
	_ = s.Substring(0, 3)
	_ = s.IsEmpty()
	_ = s.HasCharacter()
	_ = s.IsDefined()
	_ = s.AscIILength()
}

// ══════════════════════════════════════════════════════════════════════════════
// newFuncWrapCreator — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_NewFuncWrap_Invalid(t *testing.T) {
	fw := NewFuncWrap.Invalid()
	if !fw.IsInvalid() {
		t.Fatal("expected invalid")
	}
}

func Test_I13_NewFuncWrap_Many(t *testing.T) {
	many := NewFuncWrap.Many(sampleStringFunc)
	if len(many) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_I13_NewFuncWrap_Default_NonFunc(t *testing.T) {
	fw := NewFuncWrap.Default("not a func")
	if !fw.IsInvalid() {
		t.Fatal("expected invalid for non-func")
	}
}

func Test_I13_NewFuncWrap_Default_FuncWrapPtr(t *testing.T) {
	fw := NewFuncWrap.Default(sampleStringFunc)
	fw2 := NewFuncWrap.Default(fw) // should detect *FuncWrapAny
	if fw2.IsInvalid() {
		t.Fatal("expected valid when passing *FuncWrapAny")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// funcDetector — exercise
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_FuncDetector_Branches(t *testing.T) {
	fd := funcDetector{}
	// direct func
	fw := fd.GetFuncWrap(sampleStringFunc)
	if fw.IsInvalid() {
		t.Fatal("expected valid from func")
	}
	// from Map
	m := Map{"func": sampleStringFunc}
	fw2 := fd.GetFuncWrap(m)
	_ = fw2
	// from *FuncWrapAny
	fw3 := fd.GetFuncWrap(fw)
	_ = fw3
}

// ══════════════════════════════════════════════════════════════════════════════
// Map — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_Map_GetByIndex_OutOfBounds(t *testing.T) {
	m := Map{"a": 1}
	r := m.GetByIndex(999)
	if r != nil {
		t.Fatal("expected nil for out of bounds")
	}
}

func Test_I13_Map_SortedKeys(t *testing.T) {
	m := Map{"b": 1, "a": 2}
	keys, err := m.SortedKeys()
	if err != nil {
		t.Fatal("expected no error:", err)
	}
	if len(keys) != 2 || keys[0] != "a" {
		t.Fatal("expected sorted keys")
	}
}

func Test_I13_Map_InvokeWithValidArgs(t *testing.T) {
	m := Map{"func": sampleStringFunc, "first": "hello"}
	_, _ = m.InvokeWithValidArgs()
}

func Test_I13_Map_InvokeArgs(t *testing.T) {
	m := Map{"func": sampleStringFunc, "first": "hello"}
	_, _ = m.InvokeArgs("first")
}

func Test_I13_Map_GetAsStringSliceFirstOfNames(t *testing.T) {
	m := Map{"items": []string{"a", "b"}}
	r := m.GetAsStringSliceFirstOfNames("items")
	if len(r) != 2 {
		t.Fatal("expected 2 items")
	}
}
