package args

import (
	"errors"
	"fmt"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// DynamicFunc — Get nil, GetAsInt/String/Strings/AnyItems invalid, Invoke, InvokeMust
// Covers DynamicFunc.go L212, L219-221, L232-234, L245-247, L258-260, L287-289, L291-298, L303-308, L312-317
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_DynamicFunc_Get_NilReturn(t *testing.T) {
	df := &DynamicFunc[string]{
		Params: Map{"key1": "value1"},
	}
	item, isValid := df.Get("nonexistent")
	if item != nil || isValid {
		t.Error("expected nil, false for nonexistent key")
	}
}

func Test_Cov9_DynamicFunc_GetAsInt_Invalid(t *testing.T) {
	df := &DynamicFunc[string]{
		Params: Map{"key1": "not-int"},
	}
	_, isValid := df.GetAsInt("nonexistent")
	if isValid {
		t.Error("expected false for nonexistent key")
	}
}

func Test_Cov9_DynamicFunc_GetAsString_Invalid(t *testing.T) {
	df := &DynamicFunc[string]{
		Params: Map{"key1": 123},
	}
	_, isValid := df.GetAsString("nonexistent")
	if isValid {
		t.Error("expected false for nonexistent key")
	}
}

func Test_Cov9_DynamicFunc_GetAsStrings_Invalid(t *testing.T) {
	df := &DynamicFunc[string]{
		Params: Map{"key1": 123},
	}
	_, isValid := df.GetAsStrings("nonexistent")
	if isValid {
		t.Error("expected false for nonexistent key")
	}
}

func Test_Cov9_DynamicFunc_GetAsAnyItems_Invalid(t *testing.T) {
	df := &DynamicFunc[string]{
		Params: Map{"key1": 123},
	}
	_, isValid := df.GetAsAnyItems("nonexistent")
	if isValid {
		t.Error("expected false for nonexistent key")
	}
}

func Test_Cov9_DynamicFunc_Invoke(t *testing.T) {
	df := &DynamicFunc[string]{
		WorkFunc: sampleGreet,
		Params:   Map{},
	}
	results, err := df.Invoke("World")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != "Hello World" {
		t.Errorf("got %v", results[0])
	}
}

func Test_Cov9_DynamicFunc_InvokeMust(t *testing.T) {
	df := &DynamicFunc[string]{
		WorkFunc: sampleGreet,
		Params:   Map{},
	}
	results := df.InvokeMust("Test")
	if results[0] != "Hello Test" {
		t.Errorf("got %v", results[0])
	}
}

func Test_Cov9_DynamicFunc_InvokeWithValidArgs(t *testing.T) {
	df := &DynamicFunc[string]{
		WorkFunc: sampleGreet,
		Params:   Map{"input-String1": "ValidArg"},
	}
	results, err := df.InvokeWithValidArgs()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != "Hello ValidArg" {
		t.Errorf("got %v", results[0])
	}
}

func Test_Cov9_DynamicFunc_InvokeArgs(t *testing.T) {
	df := &DynamicFunc[string]{
		WorkFunc: sampleGreet,
		Params:   Map{"name": "ArgsVal"},
	}
	results, err := df.InvokeArgs("name")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != "Hello ArgsVal" {
		t.Errorf("got %v", results[0])
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Dynamic — GetWorkFunc, Invoke, InvokeMust, InvokeWithValidArgs, InvokeArgs
// Covers Dynamic.go L51, L76-78, L80-82, L86-88, L92-94
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Dynamic_GetWorkFunc(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{"func": sampleGreet},
	}
	wf := d.GetWorkFunc()
	if wf == nil {
		t.Error("expected non-nil WorkFunc")
	}
}

func Test_Cov9_Dynamic_Invoke(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{"func": sampleGreet},
	}
	results, err := d.Invoke("World")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != "Hello World" {
		t.Errorf("got %v", results[0])
	}
}

func Test_Cov9_Dynamic_InvokeMust(t *testing.T) {
	d := &Dynamic[string]{
		Params: Map{"func": sampleGreet},
	}
	results := d.InvokeMust("Test")
	if results[0] != "Hello Test" {
		t.Errorf("got %v", results[0])
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FuncWrap — IsEqual branches, InvokeMust panic, InvokeSkip panic,
// InvokeResultOfIndex, InvokeFirstAndError
// Covers FuncWrap.go L215-245, FuncWrapInvoke.go L27-28, L61-71, L85-87, L99-101, L107-137
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_FuncWrap_IsEqual_DiffInvalid(t *testing.T) {
	fw1 := NewFuncWrap.Default(sampleAdd)
	fw2 := NewFuncWrap.Invalid()
	result := fw1.IsEqual(fw2)
	if result {
		t.Error("expected false — different isInvalid")
	}
}

func Test_Cov9_FuncWrap_IsEqual_DiffName(t *testing.T) {
	fw1 := NewFuncWrap.Default(sampleAdd)
	fw2 := NewFuncWrap.Default(sampleGreet)
	result := fw1.IsEqual(fw2)
	if result {
		t.Error("expected false — different Name")
	}
}

func Test_Cov9_FuncWrap_IsEqual_DiffArgCount(t *testing.T) {
	fw1 := NewFuncWrap.Default(sampleAdd)
	fw3 := NewFuncWrap.Default(sampleThreeArgs)
	result := fw1.IsEqual(fw3)
	if result {
		t.Error("expected false — different ArgsCount")
	}
}

func Test_Cov9_FuncWrap_IsEqual_DiffReturnLength(t *testing.T) {
	fw1 := NewFuncWrap.Default(sampleAdd)       // returns (int)
	fw2 := NewFuncWrap.Default(sampleMultiReturn) // returns (string, error)
	result := fw1.IsEqual(fw2)
	if result {
		t.Error("expected false — different ReturnLength")
	}
}

func Test_Cov9_FuncWrap_InvokeMust_Panic(t *testing.T) {
	fw := NewFuncWrap.Invalid()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic from InvokeMust on invalid func")
		}
	}()
	fw.InvokeMust()
}

func Test_Cov9_FuncWrap_InvokeSkip_Panic(t *testing.T) {
	// Pass wrong args to trigger Call panic
	fw := NewFuncWrap.Default(sampleAdd)
	_, err := fw.Invoke("wrong", "types")
	if err == nil {
		t.Error("expected error from type mismatch invoke")
	}
}

func Test_Cov9_FuncWrap_InvokeResultOfIndex(t *testing.T) {
	fw := NewFuncWrap.Default(sampleAdd)
	result, err := fw.InvokeResultOfIndex(0, 3, 4)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if result != 7 {
		t.Errorf("got %v, want 7", result)
	}
}

func Test_Cov9_FuncWrap_InvokeResultOfIndex_Error(t *testing.T) {
	fw := NewFuncWrap.Invalid()
	_, err := fw.InvokeResultOfIndex(0)
	if err == nil {
		t.Error("expected error from invalid func")
	}
}

func Test_Cov9_FuncWrap_InvokeFirstAndError(t *testing.T) {
	fw := NewFuncWrap.Default(sampleMultiReturn)
	first, funcErr, procErr := fw.InvokeFirstAndError(42)
	if procErr != nil {
		t.Fatalf("processing error: %v", procErr)
	}
	if funcErr != nil {
		t.Fatalf("func error: %v", funcErr)
	}
	if first != "42" {
		t.Errorf("got %v, want '42'", first)
	}
}

func Test_Cov9_FuncWrap_InvokeFirstAndError_ProcessingError(t *testing.T) {
	fw := NewFuncWrap.Invalid()
	_, _, procErr := fw.InvokeFirstAndError()
	if procErr == nil {
		t.Error("expected processing error from invalid func")
	}
}

func Test_Cov9_FuncWrap_InvokeFirstAndError_TooFewReturns(t *testing.T) {
	fw := NewFuncWrap.Default(sampleGreet) // returns only (string)
	_, _, procErr := fw.InvokeFirstAndError("x")
	if procErr == nil {
		t.Error("expected error for function with < 2 returns")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FuncWrapTypedHelpers — InvokeAsBool, InvokeAsError, InvokeAsString, InvokeAsAny, InvokeAsAnyError
// Covers FuncWrapTypedHelpers.go L93-104, L114-131, L138-149, L158-164, L174-190
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_FuncWrap_InvokeAsBool(t *testing.T) {
	fw := NewFuncWrap.Default(sampleBool)
	result, err := fw.InvokeAsBool(5)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !result {
		t.Error("expected true")
	}
}

func Test_Cov9_FuncWrap_InvokeAsBool_NotBool(t *testing.T) {
	fw := NewFuncWrap.Default(sampleGreet)
	result, err := fw.InvokeAsBool("x")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if result {
		t.Error("expected false for non-bool result")
	}
}

func Test_Cov9_FuncWrap_InvokeAsBool_Empty(t *testing.T) {
	fw := NewFuncWrap.Default(sampleNoArgs)
	_, _ = fw.InvokeAsBool() // sampleNoArgs returns string, not bool
}

func Test_Cov9_FuncWrap_InvokeAsError(t *testing.T) {
	fw := NewFuncWrap.Default(sampleError)
	funcErr, procErr := fw.InvokeAsError("bad")
	if procErr != nil {
		t.Fatalf("processing error: %v", procErr)
	}
	if funcErr == nil {
		t.Error("expected func error")
	}
}

func Test_Cov9_FuncWrap_InvokeAsError_NilReturn(t *testing.T) {
	nilErrFunc := func() error { return nil }
	fw := NewFuncWrap.Default(nilErrFunc)
	funcErr, procErr := fw.InvokeAsError()
	if procErr != nil {
		t.Fatalf("processing error: %v", procErr)
	}
	if funcErr != nil {
		t.Error("expected nil func error")
	}
}

func Test_Cov9_FuncWrap_InvokeAsString(t *testing.T) {
	fw := NewFuncWrap.Default(sampleGreet)
	result, err := fw.InvokeAsString("World")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if result != "Hello World" {
		t.Errorf("got %q", result)
	}
}

func Test_Cov9_FuncWrap_InvokeAsString_NotString(t *testing.T) {
	fw := NewFuncWrap.Default(sampleBool)
	result, err := fw.InvokeAsString(1)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if result != "" {
		t.Error("expected empty string for non-string result")
	}
}

func Test_Cov9_FuncWrap_InvokeAsAny(t *testing.T) {
	fw := NewFuncWrap.Default(sampleGreet)
	result, err := fw.InvokeAsAny("X")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if result != "Hello X" {
		t.Errorf("got %v", result)
	}
}

func Test_Cov9_FuncWrap_InvokeAsAnyError(t *testing.T) {
	fw := NewFuncWrap.Default(sampleMultiReturn)
	result, funcErr, procErr := fw.InvokeAsAnyError(-1)
	if procErr != nil {
		t.Fatalf("processing error: %v", procErr)
	}
	if funcErr == nil {
		t.Error("expected func error")
	}
	_ = result
}

func Test_Cov9_FuncWrap_InvokeAsAnyError_NoError(t *testing.T) {
	fw := NewFuncWrap.Default(sampleMultiReturn)
	result, funcErr, procErr := fw.InvokeAsAnyError(42)
	if procErr != nil {
		t.Fatalf("processing error: %v", procErr)
	}
	if funcErr != nil {
		t.Error("expected no func error")
	}
	if result != "42" {
		t.Errorf("got %v", result)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FuncWrapArgs — InArgNames, InArgNamesEachLine, OutArgNames, OutArgNamesEachLine
// Covers FuncWrapArgs.go L115-117, L148-150, L189-191, L253-255
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_FuncWrap_InArgNames_Cached(t *testing.T) {
	fw := NewFuncWrap.Default(sampleAdd)
	names1 := fw.InArgNames()
	names2 := fw.InArgNames() // cached
	if len(names1) != len(names2) {
		t.Error("expected same result on cached call")
	}
}

func Test_Cov9_FuncWrap_InArgNamesEachLine(t *testing.T) {
	fw := NewFuncWrap.Default(sampleAdd) // 2 args
	lines := fw.InArgNamesEachLine()
	if len(lines) < 2 {
		t.Errorf("expected multi-line output, got %d lines", len(lines))
	}
}

func Test_Cov9_FuncWrap_OutArgNames(t *testing.T) {
	fw := NewFuncWrap.Default(sampleMultiReturn)
	names := fw.OutArgNames()
	if len(names) != 2 {
		t.Errorf("expected 2 out arg names, got %d", len(names))
	}
}

func Test_Cov9_FuncWrap_GetOutArgsTypesNames(t *testing.T) {
	fw := NewFuncWrap.Default(sampleMultiReturn)
	names := fw.GetOutArgsTypesNames()
	if len(names) != 2 {
		t.Errorf("expected 2 out arg type names, got %d", len(names))
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FuncWrapValidation — InvalidError branches
// Covers FuncWrapValidation.go L69-71
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_FuncWrap_InvalidError_NoValidFunc(t *testing.T) {
	fw := NewFuncWrap.Default("not-a-func")
	err := fw.InvalidError()
	if err == nil {
		t.Error("expected InvalidError for non-func")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FuncMap — Add nil, Adds, AddStructFunctions, InArgsVerifyRv, OutArgsVerifyRv,
// MustBeValid, InvokeMust error, InvokeError, InvokeFirstAndError
// Covers FuncMap.go L69-71, L84-86, L100-121, L322, L336, L361, L382-383, L464-471, L485
// ══════════════════════════════════════════════════════════════════════════════

type sampleStruct struct{}

func (s sampleStruct) PublicMethod() string { return "public" }

func Test_Cov9_FuncMap_Add_Nil(t *testing.T) {
	fm := FuncMap(nil)
	result := fm.Add(sampleGreet)
	if result == nil {
		t.Error("expected non-nil FuncMap after Add")
	}
}

func Test_Cov9_FuncMap_Adds_Nil(t *testing.T) {
	fm := FuncMap(nil)
	result := fm.Adds(sampleGreet, sampleAdd)
	if result == nil {
		t.Error("expected non-nil FuncMap after Adds")
	}
}

func Test_Cov9_FuncMap_AddStructFunctions(t *testing.T) {
	fm := FuncMap(nil)
	err := fm.AddStructFunctions(sampleStruct{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_Cov9_FuncMap_InArgsVerifyRv(t *testing.T) {
	fm := NewFuncWrap.Map(sampleAdd)
	ok, err := fm.InArgsVerifyRv("sampleAdd", NewFuncWrap.Default(sampleAdd).GetInArgsTypes())
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !ok {
		t.Error("expected ok")
	}
}

func Test_Cov9_FuncMap_OutArgsVerifyRv(t *testing.T) {
	fm := NewFuncWrap.Map(sampleAdd)
	ok, err := fm.OutArgsVerifyRv("sampleAdd", NewFuncWrap.Default(sampleAdd).GetOutArgsTypes())
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !ok {
		t.Error("expected ok")
	}
}

func Test_Cov9_FuncMap_MustBeValid(t *testing.T) {
	fm := NewFuncWrap.Map(sampleAdd)
	// Should not panic
	fm.MustBeValid("sampleAdd")
}

func Test_Cov9_FuncMap_InvokeMust_Error(t *testing.T) {
	fm := NewFuncWrap.Map(sampleAdd)
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic from InvokeMust with wrong args")
		}
	}()
	fm.InvokeMust("sampleAdd", "wrong", "types")
}

func Test_Cov9_FuncMap_InvokeError(t *testing.T) {
	fm := NewFuncWrap.Map(sampleError)
	funcErr, procErr := fm.InvokeError("sampleError", "test-msg")
	if procErr != nil {
		t.Fatalf("processing error: %v", procErr)
	}
	if funcErr == nil {
		t.Error("expected func error")
	}
}

func Test_Cov9_FuncMap_InvokeFirstAndError(t *testing.T) {
	fm := NewFuncWrap.Map(sampleMultiReturn)
	first, funcErr, procErr := fm.InvokeFirstAndError("sampleMultiReturn", 42)
	if procErr != nil {
		t.Fatalf("processing error: %v", procErr)
	}
	if funcErr != nil {
		t.Fatalf("func error: %v", funcErr)
	}
	if first != "42" {
		t.Errorf("got %v, want '42'", first)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Map — GetFuncName nil, SortedKeysMust panic, InvokeMust, InvokeWithValidArgs
// Covers Map.go L107, L207-208, L422-423, L434-439
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Map_GetFuncName_NilFunc(t *testing.T) {
	m := Map{"key": "value"}
	name := m.GetFuncName()
	if name != "" {
		t.Errorf("expected empty name, got %q", name)
	}
}

func Test_Cov9_Map_SortedKeysMust_Success(t *testing.T) {
	m := Map{"b": 2, "a": 1}
	keys := m.SortedKeysMust()
	if len(keys) != 2 {
		t.Errorf("expected 2 keys, got %d", len(keys))
	}
}

func Test_Cov9_Map_InvokeMust(t *testing.T) {
	m := Map{"func": sampleGreet}
	results := m.InvokeMust("MapTest")
	if results[0] != "Hello MapTest" {
		t.Errorf("got %v", results[0])
	}
}

func Test_Cov9_Map_InvokeMust_Panic(t *testing.T) {
	m := Map{"func": sampleGreet}
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic from InvokeMust with wrong args")
		}
	}()
	m.InvokeMust(123) // wrong arg type
}

func Test_Cov9_Map_InvokeWithValidArgs(t *testing.T) {
	m := Map{
		"func":          sampleGreet,
		"input-String1": "ValidArg",
	}
	results, err := m.InvokeWithValidArgs()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != "Hello ValidArg" {
		t.Errorf("got %v", results[0])
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Holder — Invoke, InvokeMust, InvokeWithValidArgs, InvokeArgs
// Covers Holder.go L188-190, L195-197, L202-204
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Holder_Invoke(t *testing.T) {
	h := &Holder[string]{
		First:    "World",
		WorkFunc: sampleGreet,
	}
	results, err := h.Invoke("World")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != "Hello World" {
		t.Errorf("got %v", results[0])
	}
}

func Test_Cov9_Holder_InvokeMust(t *testing.T) {
	h := &Holder[string]{First: "Test", WorkFunc: sampleGreet}
	results := h.InvokeMust("Test")
	if results[0] != "Hello Test" {
		t.Errorf("got %v", results[0])
	}
}

func Test_Cov9_Holder_InvokeWithValidArgs(t *testing.T) {
	h := &Holder[string]{First: "VA", WorkFunc: sampleGreet}
	results, err := h.InvokeWithValidArgs()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != "Hello VA" {
		t.Errorf("got %v", results[0])
	}
}

func Test_Cov9_Holder_InvokeArgs(t *testing.T) {
	h := &Holder[string]{First: "Args", WorkFunc: sampleGreet}
	results, err := h.InvokeArgs(1)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != "Hello Args" {
		t.Errorf("got %v", results[0])
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// newFuncWrapCreator — MethodToFunc, StructToMap
// Covers newFuncWrapCreator.go L27-28, L101-118, L122-141
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_NewFuncWrap_MethodToFunc_Nil(t *testing.T) {
	result, err := NewFuncWrap.MethodToFunc(nil)
	if err == nil {
		t.Error("expected error for nil method")
	}
	if result == nil || !result.IsInvalid() {
		t.Error("expected invalid FuncWrap")
	}
}

func Test_Cov9_NewFuncWrap_StructToMap(t *testing.T) {
	fm, err := NewFuncWrap.StructToMap(sampleStruct{})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if len(fm) == 0 {
		t.Error("expected at least one method in map")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// funcDetector — GetFuncWrap with different types
// Covers funcDetector.go L16-22
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_FuncDetector_Map(t *testing.T) {
	m := Map{"func": sampleGreet}
	fd := funcDetector{}
	fw := fd.GetFuncWrap(m)
	if fw == nil {
		t.Error("expected non-nil FuncWrap from Map")
	}
}

func Test_Cov9_FuncDetector_FuncWrapPtr(t *testing.T) {
	fwPtr := NewFuncWrap.Default(sampleGreet)
	fd := funcDetector{}
	fw := fd.GetFuncWrap(fwPtr)
	if fw != fwPtr {
		t.Error("expected same pointer back")
	}
}

func Test_Cov9_FuncDetector_Default(t *testing.T) {
	fd := funcDetector{}
	fw := fd.GetFuncWrap(sampleAdd)
	if fw == nil || fw.IsInvalid() {
		t.Error("expected valid FuncWrap from raw func")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// argsHelper — invokeMustHelper panic
// Covers argsHelper.go L31-33, L62-69
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_InvokeMustHelper_Panic(t *testing.T) {
	fw := NewFuncWrap.Invalid()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic from invokeMustHelper")
		}
	}()
	invokeMustHelper(fw)
}

func Test_Cov9_BuildToString_Cached(t *testing.T) {
	// Exercise cached path by calling String() twice on any Func type
	tf := &TwoFunc[int, int]{First: 1, Second: 2}
	s1 := tf.String()
	s2 := tf.String() // cached
	if s1 != s2 {
		t.Error("expected same result on cached call")
	}
	_ = fmt.Sprint(s1, s2)
}
