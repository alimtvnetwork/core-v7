package stringcompareastests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func Test_Variant_Name_Verification(t *testing.T) {
	if stringcompareas.Equal.Name() != "Equal" {
		t.Errorf("expected Equal, got %s", stringcompareas.Equal.Name())
	}
	if stringcompareas.StartsWith.Name() != "StartsWith" {
		t.Error("expected StartsWith")
	}
	if stringcompareas.EndsWith.Name() != "EndsWith" {
		t.Error("expected EndsWith")
	}
	if stringcompareas.Anywhere.Name() != "Anywhere" {
		t.Error("expected Anywhere")
	}
	if stringcompareas.Regex.Name() != "Regex" {
		t.Error("expected Regex")
	}
}

func Test_Variant_Is_Methods(t *testing.T) {
	if !stringcompareas.Equal.IsEqual() {
		t.Error("Equal should be IsEqual")
	}
	if !stringcompareas.StartsWith.IsStartsWith() {
		t.Error("StartsWith should be IsStartsWith")
	}
	if !stringcompareas.EndsWith.IsEndsWith() {
		t.Error("EndsWith should be IsEndsWith")
	}
	if !stringcompareas.Anywhere.IsAnywhere() {
		t.Error("Anywhere should be IsAnywhere")
	}
	if !stringcompareas.Contains.IsContains() {
		t.Error("Contains should be IsContains")
	}
	if !stringcompareas.AnyChars.IsAnyChars() {
		t.Error("AnyChars should be IsAnyChars")
	}
	if !stringcompareas.Regex.IsRegex() {
		t.Error("Regex should be IsRegex")
	}
	if !stringcompareas.Glob.IsGlob() {
		t.Error("Glob should be IsGlob")
	}
	if !stringcompareas.NonGlob.IsNonGlob() {
		t.Error("NonGlob should be IsNonGlob")
	}
}

func Test_Variant_Not_Methods(t *testing.T) {
	if !stringcompareas.NotEqual.IsNotEqual() {
		t.Error("NotEqual should be IsNotEqual")
	}
	if !stringcompareas.NotStartsWith.IsNotStartsWith() {
		t.Error("NotStartsWith should be IsNotStartsWith")
	}
	if !stringcompareas.NotEndsWith.IsNotEndsWith() {
		t.Error("NotEndsWith should be IsNotEndsWith")
	}
	if !stringcompareas.NotContains.IsNotContains() {
		t.Error("NotContains should be IsNotContains")
	}
	if !stringcompareas.NotMatchRegex.IsNotMatchRegex() {
		t.Error("NotMatchRegex should be IsNotMatchRegex")
	}
}

func Test_Variant_IsNegativeCondition(t *testing.T) {
	if !stringcompareas.NotEqual.IsNegativeCondition() {
		t.Error("NotEqual should be negative")
	}
	if !stringcompareas.NotStartsWith.IsNegativeCondition() {
		t.Error("NotStartsWith should be negative")
	}
	if !stringcompareas.NonGlob.IsNegativeCondition() {
		t.Error("NonGlob should be negative")
	}
	if stringcompareas.Equal.IsNegativeCondition() {
		t.Error("Equal should not be negative")
	}
}

func Test_Variant_ValidInvalid(t *testing.T) {
	if !stringcompareas.Equal.IsValid() {
		t.Error("Equal should be valid")
	}
	if stringcompareas.Equal.IsInvalid() {
		t.Error("Equal should not be invalid")
	}
	if !stringcompareas.Invalid.IsInvalid() {
		t.Error("Invalid should be invalid")
	}
}

func Test_Variant_IsCompareSuccess(t *testing.T) {
	// Equal
	if !stringcompareas.Equal.IsCompareSuccess(false, "hello", "hello") {
		t.Error("Equal should match")
	}
	if stringcompareas.Equal.IsCompareSuccess(false, "hello", "world") {
		t.Error("Equal should not match different")
	}

	// StartsWith
	if !stringcompareas.StartsWith.IsCompareSuccess(false, "hello world", "hello") {
		t.Error("StartsWith should match")
	}

	// EndsWith
	if !stringcompareas.EndsWith.IsCompareSuccess(false, "hello world", "world") {
		t.Error("EndsWith should match")
	}

	// Anywhere/Contains
	if !stringcompareas.Anywhere.IsCompareSuccess(false, "hello world", "lo wo") {
		t.Error("Anywhere should match")
	}

	// NotEqual
	if !stringcompareas.NotEqual.IsCompareSuccess(false, "hello", "world") {
		t.Error("NotEqual should match different strings")
	}

	// Case insensitive
	if !stringcompareas.Equal.IsCompareSuccess(true, "Hello", "hello") {
		t.Error("Equal ignore case should match")
	}
}

func Test_Variant_CompareSuccessCaseSensitive(t *testing.T) {
	v := stringcompareas.Equal
	if !v.IsCompareSuccessCaseSensitive("hello", "hello") {
		t.Error("case sensitive should match")
	}
	if v.IsCompareSuccessCaseSensitive("Hello", "hello") {
		t.Error("case sensitive should not match different case")
	}
}

func Test_Variant_CompareSuccessNonCaseSensitive(t *testing.T) {
	v := stringcompareas.Equal
	if !v.IsCompareSuccessNonCaseSensitive("Hello", "hello") {
		t.Error("case insensitive should match")
	}
}

func Test_Variant_VerifyMessage(t *testing.T) {
	msg := stringcompareas.Equal.VerifyMessage(false, "hello", "hello")
	if msg != "" {
		t.Error("matching should return empty message")
	}

	msg = stringcompareas.Equal.VerifyMessage(false, "hello", "world")
	if msg == "" {
		t.Error("non-matching should return error message")
	}

	// negative case
	msg = stringcompareas.NotEqual.VerifyMessage(false, "hello", "hello")
	if msg == "" {
		t.Error("NotEqual same values should return error message")
	}
}

func Test_Variant_VerifyError(t *testing.T) {
	err := stringcompareas.Equal.VerifyError(false, "hello", "hello")
	if err != nil {
		t.Error("matching should return nil error")
	}

	err = stringcompareas.Equal.VerifyError(false, "hello", "world")
	if err == nil {
		t.Error("non-matching should return error")
	}
}

func Test_Variant_VerifyMessageCaseSensitive(t *testing.T) {
	msg := stringcompareas.Equal.VerifyMessageCaseSensitive("hello", "hello")
	if msg != "" {
		t.Error("matching should return empty")
	}
}

func Test_Variant_VerifyErrorCaseSensitive(t *testing.T) {
	err := stringcompareas.Equal.VerifyErrorCaseSensitive("hello", "hello")
	if err != nil {
		t.Error("matching should return nil")
	}
}

func Test_Variant_MarshalJSON(t *testing.T) {
	bytes, err := stringcompareas.Equal.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}
	if len(bytes) == 0 {
		t.Error("MarshalJSON should return bytes")
	}
}

func Test_Variant_UnmarshalJSON(t *testing.T) {
	v := stringcompareas.Invalid
	data, _ := json.Marshal("Equal")
	err := v.UnmarshalJSON(data)
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	}
}

func Test_Variant_ValueMethods(t *testing.T) {
	v := stringcompareas.Equal
	if v.Value() != 0 {
		t.Error("Equal should be 0")
	}
	if v.ValueInt() != 0 {
		t.Error("ValueInt should be 0")
	}
	if v.ValueInt8() != 0 {
		t.Error("ValueInt8 should be 0")
	}
	if v.ValueInt16() != 0 {
		t.Error("ValueInt16 should be 0")
	}
	if v.ValueInt32() != 0 {
		t.Error("ValueInt32 should be 0")
	}
	if v.ValueUInt16() != 0 {
		t.Error("ValueUInt16 should be 0")
	}
}

func Test_Variant_String(t *testing.T) {
	if stringcompareas.Equal.String() != "Equal" {
		t.Error("String should be Equal")
	}
}

func Test_Variant_Is(t *testing.T) {
	if !stringcompareas.Equal.Is(stringcompareas.Equal) {
		t.Error("Equal should Is Equal")
	}
	if stringcompareas.Equal.Is(stringcompareas.StartsWith) {
		t.Error("Equal should not Is StartsWith")
	}
}

func Test_Variant_NameValue(t *testing.T) {
	nv := stringcompareas.Equal.NameValue()
	if nv == "" {
		t.Error("NameValue should not be empty")
	}
}

func Test_Variant_ValueString(t *testing.T) {
	vs := stringcompareas.Equal.ValueString()
	if vs == "" {
		t.Error("ValueString should not be empty")
	}
}

func Test_Variant_ToNumberString(t *testing.T) {
	ns := stringcompareas.Equal.ToNumberString()
	if ns == "" {
		t.Error("ToNumberString should not be empty")
	}
}

func Test_Variant_RangeNamesCsv(t *testing.T) {
	csv := stringcompareas.Equal.RangeNamesCsv()
	if csv == "" {
		t.Error("RangeNamesCsv should not be empty")
	}
}

func Test_Variant_IsAnyMethod(t *testing.T) {
	if !stringcompareas.Equal.IsAnyMethod("Equal") {
		t.Error("should match Equal")
	}
	if stringcompareas.Equal.IsAnyMethod("StartsWith") {
		t.Error("should not match StartsWith")
	}
}

func Test_Variant_IsNameEqual(t *testing.T) {
	if !stringcompareas.Equal.IsNameEqual("Equal") {
		t.Error("should be name equal")
	}
}

func Test_Variant_IsAnyNamesOf(t *testing.T) {
	if !stringcompareas.Equal.IsAnyNamesOf("Equal", "StartsWith") {
		t.Error("should match Equal")
	}
	if stringcompareas.Equal.IsAnyNamesOf("StartsWith", "EndsWith") {
		t.Error("should not match")
	}
}

func Test_Variant_IsValueEqual(t *testing.T) {
	if !stringcompareas.Equal.IsValueEqual(0) {
		t.Error("Equal value should be 0")
	}
}

func Test_Variant_IsAnyValuesEqual(t *testing.T) {
	if !stringcompareas.Equal.IsAnyValuesEqual(0, 1) {
		t.Error("should match 0")
	}
	if stringcompareas.Equal.IsAnyValuesEqual(1, 2) {
		t.Error("should not match")
	}
}

func Test_Variant_IsByteValueEqual(t *testing.T) {
	if !stringcompareas.Equal.IsByteValueEqual(0) {
		t.Error("should match")
	}
}

func Test_Variant_MaxMinByte(t *testing.T) {
	v := stringcompareas.Equal
	if v.MaxByte() == 0 {
		t.Error("max should not be 0")
	}
	if v.MinByte() != 0 {
		t.Error("min should be 0")
	}
}

func Test_Variant_RangesByte(t *testing.T) {
	v := stringcompareas.Equal
	ranges := v.RangesByte()
	if len(ranges) == 0 {
		t.Error("ranges should not be empty")
	}
}

func Test_Variant_AllNameValues(t *testing.T) {
	nv := stringcompareas.Equal.AllNameValues()
	if len(nv) == 0 {
		t.Error("AllNameValues should not be empty")
	}
}

func Test_Variant_IntegerEnumRanges(t *testing.T) {
	r := stringcompareas.Equal.IntegerEnumRanges()
	if len(r) == 0 {
		t.Error("IntegerEnumRanges should not be empty")
	}
}

func Test_Variant_MinMaxAny(t *testing.T) {
	min, max := stringcompareas.Equal.MinMaxAny()
	if min == nil || max == nil {
		t.Error("MinMaxAny should not be nil")
	}
}

func Test_Variant_MinMaxValueString(t *testing.T) {
	if stringcompareas.Equal.MinValueString() == "" {
		t.Error("MinValueString should not be empty")
	}
	if stringcompareas.Equal.MaxValueString() == "" {
		t.Error("MaxValueString should not be empty")
	}
}

func Test_Variant_MaxMinInt(t *testing.T) {
	if stringcompareas.Equal.MaxInt() == 0 {
		t.Error("MaxInt should not be 0")
	}
}

func Test_Variant_RangesDynamicMap(t *testing.T) {
	m := stringcompareas.Equal.RangesDynamicMap()
	if len(m) == 0 {
		t.Error("RangesDynamicMap should not be empty")
	}
}

func Test_Variant_Format(t *testing.T) {
	f := stringcompareas.Equal.Format("type: %s")
	if f == "" {
		t.Error("Format should not be empty")
	}
}

func Test_Variant_ToPtr(t *testing.T) {
	v := stringcompareas.Equal
	ptr := v.ToPtr()
	if ptr == nil {
		t.Error("ToPtr should not be nil")
	}
}

func Test_Variant_TypeName(t *testing.T) {
	tn := stringcompareas.Equal.TypeName()
	if tn == "" {
		t.Error("TypeName should not be empty")
	}
}

func Test_Variant_EnumType(t *testing.T) {
	et := stringcompareas.Equal.EnumType()
	if et == nil {
		t.Error("EnumType should not be nil")
	}
}

func Test_Variant_AsInterfaces(t *testing.T) {
	if stringcompareas.Equal.AsBasicEnumContractsBinder() == nil {
		t.Error("AsBasicEnumContractsBinder should not be nil")
	}
	if stringcompareas.Equal.AsStringCompareTyper() == nil {
		t.Error("AsStringCompareTyper should not be nil")
	}
	if stringcompareas.Equal.AsBasicByteEnumContractsBinder() == nil {
		t.Error("AsBasicByteEnumContractsBinder should not be nil")
	}
}

func Test_Variant_UnmarshallEnumToValue(t *testing.T) {
	data, _ := json.Marshal("Equal")
	val, err := stringcompareas.Equal.UnmarshallEnumToValue(data)
	if err != nil {
		t.Errorf("UnmarshallEnumToValue error: %v", err)
	}
	if val != byte(stringcompareas.Equal) {
		t.Error("should return Equal value")
	}
}

func Test_Variant_IsEnumEqual(t *testing.T) {
	a := stringcompareas.Equal
	b := stringcompareas.Equal
	if !a.IsEnumEqual(&b) {
		t.Error("same enum should be equal")
	}
}

func Test_Variant_IsAnyEnumsEqual(t *testing.T) {
	a := stringcompareas.Equal
	b := stringcompareas.StartsWith
	c := stringcompareas.Equal
	if !a.IsAnyEnumsEqual(&b, &c) {
		t.Error("should match Equal")
	}
}

func Test_Variant_OnlySupportedErr(t *testing.T) {
	// Passing all names as supported → no unsupported → nil error
	allNames := []string{
		"Equal", "StartsWith", "EndsWith", "Anywhere",
		"IsContains", "AnyChars", "Regex",
		"NotEqual", "NotStartsWith", "NotEndsWith",
		"NotContains", "NotAnyChars", "NotMatchRegex",
		"Glob", "NonGlob", "Invalid",
	}
	err := stringcompareas.Equal.OnlySupportedErr(allNames...)
	if err != nil {
		t.Errorf("all names supported should not error, got: %v", err)
	}

	// Passing only "Equal" → all others unsupported → error expected
	err2 := stringcompareas.Equal.OnlySupportedErr("Equal")
	if err2 == nil {
		t.Error("partial support should return error for unsupported names")
	}
}

func Test_Variant_DynamicCompare(t *testing.T) {
	v := stringcompareas.Equal
	dynFunc := func(index int, content string, compareAs stringcompareas.Variant) bool {
		return compareAs == stringcompareas.Equal && content == "hello"
	}
	if !v.DynamicCompare(dynFunc, 0, "hello") {
		t.Error("dynamic compare should return true")
	}
}

func Test_Variant_IsLineCompareFunc(t *testing.T) {
	fn := stringcompareas.Equal.IsLineCompareFunc()
	if fn == nil {
		t.Error("IsLineCompareFunc should not be nil")
	}
	if !fn("hello", "hello", false) {
		t.Error("Equal compare func should match same strings")
	}
}

func Test_Variant_GlobCompare(t *testing.T) {
	if !stringcompareas.Glob.IsCompareSuccess(false, "hello.txt", "*.txt") {
		t.Error("Glob should match *.txt")
	}
	if stringcompareas.NonGlob.IsCompareSuccess(false, "hello.txt", "*.txt") {
		t.Error("NonGlob should not match *.txt")
	}
}

func Test_Variant_AnyCharsCompare(t *testing.T) {
	if !stringcompareas.AnyChars.IsCompareSuccess(false, "hello", "hlo") {
		t.Error("AnyChars should match when chars exist")
	}
	if stringcompareas.NotAnyChars.IsCompareSuccess(false, "hello", "hlo") {
		t.Error("NotAnyChars should not match when chars exist")
	}
}

func Test_Variant_RegexCompare(t *testing.T) {
	if !stringcompareas.Regex.IsCompareSuccess(false, "hello123", `\d+`) {
		t.Error("Regex should match digits")
	}
}
