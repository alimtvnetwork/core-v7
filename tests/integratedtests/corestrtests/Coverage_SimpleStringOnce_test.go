package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Segment 16+17 Part 1
// ══════════════════════════════════════════════════════════════════════════════

func newSSO(val string) *corestr.SimpleStringOnce {
	sso := &corestr.SimpleStringOnce{}
	sso.SetOnUninitialized(val)
	return sso
}

func Test_CovSSO_01_Value_IsInitialized_IsDefined(t *testing.T) {
	safeTest(t, "Test_CovSSO_01_Value_IsInitialized_IsDefined", func() {
		sso := &corestr.SimpleStringOnce{}
		if sso.IsInitialized() {
			t.Fatal("expected uninitialized")
		}
		if sso.IsDefined() {
			t.Fatal("expected undefined")
		}
		if !sso.IsUninitialized() {
			t.Fatal("expected uninitialized")
		}
		sso.SetOnUninitialized("hello")
		if sso.Value() != "hello" {
			t.Fatal("expected hello")
		}
		if !sso.IsInitialized() {
			t.Fatal("expected initialized")
		}
		if !sso.IsDefined() {
			t.Fatal("expected defined")
		}
	})
}

func Test_CovSSO_02_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_CovSSO_02_SetOnUninitialized_AlreadyInit", func() {
		sso := newSSO("first")
		err := sso.SetOnUninitialized("second")
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovSSO_03_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_CovSSO_03_GetSetOnce", func() {
		sso := &corestr.SimpleStringOnce{}
		v := sso.GetSetOnce("hello")
		if v != "hello" {
			t.Fatal("expected hello")
		}
		// already init
		v2 := sso.GetSetOnce("world")
		if v2 != "hello" {
			t.Fatal("expected hello")
		}
	})
}

func Test_CovSSO_04_GetOnce(t *testing.T) {
	safeTest(t, "Test_CovSSO_04_GetOnce", func() {
		sso := &corestr.SimpleStringOnce{}
		v := sso.GetOnce()
		if v != "" {
			t.Fatal("expected empty")
		}
		// already init
		v2 := sso.GetOnce()
		if v2 != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovSSO_05_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_CovSSO_05_GetOnceFunc", func() {
		sso := &corestr.SimpleStringOnce{}
		v := sso.GetOnceFunc(func() string { return "computed" })
		if v != "computed" {
			t.Fatal("expected computed")
		}
		// already init
		v2 := sso.GetOnceFunc(func() string { return "other" })
		if v2 != "computed" {
			t.Fatal("expected computed")
		}
	})
}

func Test_CovSSO_06_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_CovSSO_06_SetOnceIfUninitialized", func() {
		sso := &corestr.SimpleStringOnce{}
		ok := sso.SetOnceIfUninitialized("hello")
		if !ok {
			t.Fatal("expected true")
		}
		ok2 := sso.SetOnceIfUninitialized("world")
		if ok2 {
			t.Fatal("expected false")
		}
	})
}

func Test_CovSSO_07_Invalidate_Reset(t *testing.T) {
	safeTest(t, "Test_CovSSO_07_Invalidate_Reset", func() {
		sso := newSSO("hello")
		sso.Invalidate()
		if sso.IsInitialized() {
			t.Fatal("expected uninitialized")
		}
		sso.SetOnUninitialized("new")
		sso.Reset()
		if sso.IsInitialized() {
			t.Fatal("expected uninitialized")
		}
	})
}

func Test_CovSSO_08_SetInitialize_SetUnInit(t *testing.T) {
	safeTest(t, "Test_CovSSO_08_SetInitialize_SetUnInit", func() {
		sso := &corestr.SimpleStringOnce{}
		sso.SetInitialize()
		if !sso.IsInitialized() {
			t.Fatal("expected initialized")
		}
		sso.SetUnInit()
		if sso.IsInitialized() {
			t.Fatal("expected uninitialized")
		}
	})
}

func Test_CovSSO_09_IsInvalid(t *testing.T) {
	safeTest(t, "Test_CovSSO_09_IsInvalid", func() {
		sso := &corestr.SimpleStringOnce{}
		if !sso.IsInvalid() {
			t.Fatal("expected invalid")
		}
		sso.SetOnUninitialized("")
		if !sso.IsInvalid() {
			t.Fatal("expected invalid for empty value")
		}
		sso2 := newSSO("hello")
		if sso2.IsInvalid() {
			t.Fatal("expected valid")
		}
	})
}

func Test_CovSSO_10_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_CovSSO_10_IsEmpty_IsWhitespace", func() {
		sso := &corestr.SimpleStringOnce{}
		if !sso.IsEmpty() {
			t.Fatal("expected empty")
		}
		if !sso.IsWhitespace() {
			t.Fatal("expected whitespace")
		}
		sso2 := newSSO("hello")
		if sso2.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_CovSSO_11_Trim(t *testing.T) {
	safeTest(t, "Test_CovSSO_11_Trim", func() {
		sso := newSSO("  hello  ")
		if sso.Trim() != "hello" {
			t.Fatal("expected trimmed")
		}
	})
}

func Test_CovSSO_12_HasValidNonEmpty_HasValidNonWhitespace_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovSSO_12_HasValidNonEmpty_HasValidNonWhitespace_HasSafeNonEmpty", func() {
		sso := &corestr.SimpleStringOnce{}
		if sso.HasValidNonEmpty() {
			t.Fatal("expected false")
		}
		if sso.HasValidNonWhitespace() {
			t.Fatal("expected false")
		}
		if sso.HasSafeNonEmpty() {
			t.Fatal("expected false")
		}
		sso2 := newSSO("hello")
		if !sso2.HasValidNonEmpty() {
			t.Fatal("expected true")
		}
		if !sso2.HasValidNonWhitespace() {
			t.Fatal("expected true")
		}
		if !sso2.HasSafeNonEmpty() {
			t.Fatal("expected true")
		}
	})
}

func Test_CovSSO_13_SafeValue(t *testing.T) {
	safeTest(t, "Test_CovSSO_13_SafeValue", func() {
		sso := &corestr.SimpleStringOnce{}
		if sso.SafeValue() != "" {
			t.Fatal("expected empty")
		}
		sso2 := newSSO("hello")
		if sso2.SafeValue() != "hello" {
			t.Fatal("expected hello")
		}
	})
}

func Test_CovSSO_14_ValueBytes_ValueBytesPtr(t *testing.T) {
	safeTest(t, "Test_CovSSO_14_ValueBytes_ValueBytesPtr", func() {
		sso := newSSO("hi")
		if len(sso.ValueBytes()) != 2 {
			t.Fatal("expected 2")
		}
		if len(sso.ValueBytesPtr()) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovSSO_15_Int_ValueInt_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_CovSSO_15_Int_ValueInt_ValueDefInt", func() {
		sso := newSSO("42")
		if sso.Int() != 42 {
			t.Fatal("expected 42")
		}
		if sso.ValueInt(0) != 42 {
			t.Fatal("expected 42")
		}
		if sso.ValueDefInt() != 42 {
			t.Fatal("expected 42")
		}
		// invalid
		sso2 := newSSO("abc")
		if sso2.Int() != 0 {
			t.Fatal("expected 0")
		}
		if sso2.ValueInt(99) != 99 {
			t.Fatal("expected 99")
		}
		if sso2.ValueDefInt() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovSSO_16_Byte_ValueByte_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_CovSSO_16_Byte_ValueByte_ValueDefByte", func() {
		sso := newSSO("100")
		if sso.Byte() != 100 {
			t.Fatal("expected 100")
		}
		if sso.ValueByte(0) != 100 {
			t.Fatal("expected 100")
		}
		if sso.ValueDefByte() != 100 {
			t.Fatal("expected 100")
		}
		// out of range
		sso2 := newSSO("999")
		if sso2.Byte() != 0 {
			t.Fatal("expected 0")
		}
		if sso2.ValueByte(5) != 5 {
			t.Fatal("expected 5")
		}
		// invalid
		sso3 := newSSO("abc")
		if sso3.Byte() != 0 {
			t.Fatal("expected 0")
		}
		if sso3.ValueByte(7) != 7 {
			t.Fatal("expected 7")
		}
	})
}

func Test_CovSSO_17_Int16_Int32(t *testing.T) {
	safeTest(t, "Test_CovSSO_17_Int16_Int32", func() {
		sso := newSSO("100")
		if sso.Int16() != 100 {
			t.Fatal("expected 100")
		}
		if sso.Int32() != 100 {
			t.Fatal("expected 100")
		}
		// out of range for int16
		sso2 := newSSO("99999")
		if sso2.Int16() != 0 {
			t.Fatal("expected 0")
		}
		// invalid
		sso3 := newSSO("abc")
		if sso3.Int16() != 0 {
			t.Fatal("expected 0")
		}
		if sso3.Int32() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovSSO_18_Uint16_Uint32(t *testing.T) {
	safeTest(t, "Test_CovSSO_18_Uint16_Uint32", func() {
		sso := newSSO("100")
		v16, ok16 := sso.Uint16()
		if v16 != 100 || !ok16 {
			t.Fatal("expected 100 in range")
		}
		v32, ok32 := sso.Uint32()
		if v32 != 100 || !ok32 {
			t.Fatal("expected 100 in range")
		}
	})
}

func Test_CovSSO_19_WithinRange_WithinRangeDefault(t *testing.T) {
	safeTest(t, "Test_CovSSO_19_WithinRange_WithinRangeDefault", func() {
		sso := newSSO("50")
		v, ok := sso.WithinRange(true, 0, 100)
		if v != 50 || !ok {
			t.Fatal("expected 50 in range")
		}
		// out of range with boundary
		v2, ok2 := sso.WithinRange(true, 60, 100)
		if v2 != 60 || ok2 {
			t.Fatal("expected min boundary")
		}
		// out of range, above max
		sso2 := newSSO("200")
		v3, ok3 := sso2.WithinRange(true, 0, 100)
		if v3 != 100 || ok3 {
			t.Fatal("expected max boundary")
		}
		// no boundary
		v4, ok4 := sso2.WithinRange(false, 0, 100)
		if v4 != 200 || ok4 {
			t.Fatal("expected 200 out of range")
		}
		// invalid
		sso3 := newSSO("abc")
		v5, ok5 := sso3.WithinRange(true, 0, 100)
		if v5 != 0 || ok5 {
			t.Fatal("expected 0")
		}
		// WithinRangeDefault
		v6, ok6 := sso.WithinRangeDefault(0, 100)
		if v6 != 50 || !ok6 {
			t.Fatal("expected 50")
		}
	})
}

func Test_CovSSO_20_ValueFloat64_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_CovSSO_20_ValueFloat64_ValueDefFloat64", func() {
		sso := newSSO("3.14")
		if sso.ValueFloat64(0) != 3.14 {
			t.Fatal("expected 3.14")
		}
		if sso.ValueDefFloat64() != 3.14 {
			t.Fatal("expected 3.14")
		}
		// invalid
		sso2 := newSSO("abc")
		if sso2.ValueFloat64(1.5) != 1.5 {
			t.Fatal("expected 1.5")
		}
	})
}

func Test_CovSSO_21_Boolean_BooleanDefault_IsValueBool(t *testing.T) {
	safeTest(t, "Test_CovSSO_21_Boolean_BooleanDefault_IsValueBool", func() {
		cases := []struct {
			val    string
			expect bool
		}{
			{"yes", true}, {"y", true}, {"1", true}, {"YES", true}, {"Y", true},
			{"true", true}, {"false", false}, {"", false}, {"abc", false},
		}
		for _, c := range cases {
			sso := newSSO(c.val)
			if sso.Boolean(false) != c.expect {
				t.Fatalf("Boolean(%s) expected %v", c.val, c.expect)
			}
		}
		// with consider init
		uninit := &corestr.SimpleStringOnce{}
		if uninit.Boolean(true) {
			t.Fatal("expected false for uninitialized")
		}
		// BooleanDefault
		sso := newSSO("true")
		if !sso.BooleanDefault() {
			t.Fatal("expected true")
		}
		// IsValueBool
		if !sso.IsValueBool() {
			t.Fatal("expected true")
		}
	})
}

func Test_CovSSO_22_IsSetter(t *testing.T) {
	safeTest(t, "Test_CovSSO_22_IsSetter", func() {
		sso := newSSO("yes")
		v := sso.IsSetter(false)
		if v.String() != "True" {
			t.Fatalf("expected True, got %s", v.String())
		}
		// invalid
		sso2 := newSSO("abc")
		v2 := sso2.IsSetter(false)
		if v2.String() != "Uninitialized" {
			t.Fatal("expected Uninitialized")
		}
		// consider init, uninitialized
		uninit := &corestr.SimpleStringOnce{}
		v3 := uninit.IsSetter(true)
		if v3.String() != "False" {
			t.Fatal("expected False")
		}
		// false value
		sso3 := newSSO("false")
		v4 := sso3.IsSetter(false)
		if v4.String() != "False" {
			t.Fatal("expected False")
		}
	})
}

func Test_CovSSO_23_Is_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_CovSSO_23_Is_IsAnyOf", func() {
		sso := newSSO("hello")
		if !sso.Is("hello") {
			t.Fatal("expected true")
		}
		if sso.Is("world") {
			t.Fatal("expected false")
		}
		if !sso.IsAnyOf("a", "hello") {
			t.Fatal("expected true")
		}
		if sso.IsAnyOf("a", "b") {
			t.Fatal("expected false")
		}
		// empty values → true
		if !sso.IsAnyOf() {
			t.Fatal("expected true")
		}
	})
}

func Test_CovSSO_24_IsContains_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_CovSSO_24_IsContains_IsAnyContains", func() {
		sso := newSSO("hello world")
		if !sso.IsContains("world") {
			t.Fatal("expected true")
		}
		if !sso.IsAnyContains("xyz", "world") {
			t.Fatal("expected true")
		}
		if sso.IsAnyContains("xyz", "abc") {
			t.Fatal("expected false")
		}
		if !sso.IsAnyContains() {
			t.Fatal("expected true for empty")
		}
	})
}

func Test_CovSSO_25_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_CovSSO_25_IsEqualNonSensitive", func() {
		sso := newSSO("Hello")
		if !sso.IsEqualNonSensitive("hello") {
			t.Fatal("expected true")
		}
	})
}

func Test_CovSSO_26_IsRegexMatches_RegexFind(t *testing.T) {
	safeTest(t, "Test_CovSSO_26_IsRegexMatches_RegexFind", func() {
		sso := newSSO("hello123")
		re := regexp.MustCompile(`\d+`)
		if !sso.IsRegexMatches(re) {
			t.Fatal("expected true")
		}
		if sso.IsRegexMatches(nil) {
			t.Fatal("expected false for nil")
		}
		found := sso.RegexFindString(re)
		if found != "123" {
			t.Fatal("expected 123")
		}
		if sso.RegexFindString(nil) != "" {
			t.Fatal("expected empty")
		}
		items := sso.RegexFindAllStrings(re, -1)
		if len(items) != 1 {
			t.Fatal("expected 1")
		}
		if len(sso.RegexFindAllStrings(nil, -1)) != 0 {
			t.Fatal("expected 0")
		}
		items2, has := sso.RegexFindAllStringsWithFlag(re, -1)
		if !has || len(items2) != 1 {
			t.Fatal("expected found")
		}
		_, has2 := sso.RegexFindAllStringsWithFlag(nil, -1)
		if has2 {
			t.Fatal("expected false")
		}
	})
}

func Test_CovSSO_27_Split_SplitLeftRight_SplitLeftRightTrim(t *testing.T) {
	safeTest(t, "Test_CovSSO_27_Split_SplitLeftRight_SplitLeftRightTrim", func() {
		sso := newSSO("key=value")
		parts := sso.Split("=")
		if len(parts) != 2 {
			t.Fatal("expected 2")
		}
		l, r := sso.SplitLeftRight("=")
		if l != "key" || r != "value" {
			t.Fatal("expected key,value")
		}
		// no right
		sso2 := newSSO("onlykey")
		l2, r2 := sso2.SplitLeftRight("=")
		if l2 != "onlykey" || r2 != "" {
			t.Fatal("expected onlykey,empty")
		}
		// trim
		sso3 := newSSO(" key = value ")
		l3, r3 := sso3.SplitLeftRightTrim("=")
		if l3 != "key" || r3 != "value" {
			t.Fatalf("expected key,value got '%s','%s'", l3, r3)
		}
		// no right trim
		sso4 := newSSO("onlykey")
		l4, r4 := sso4.SplitLeftRightTrim("=")
		if l4 != "onlykey" || r4 != "" {
			t.Fatal("expected onlykey,empty")
		}
	})
}

func Test_CovSSO_28_SplitNonEmpty_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_CovSSO_28_SplitNonEmpty_SplitTrimNonWhitespace", func() {
		sso := newSSO("a,,b")
		parts := sso.SplitNonEmpty(",")
		if len(parts) < 2 {
			t.Fatal("expected at least 2")
		}
		sso2 := newSSO("a, ,b")
		parts2 := sso2.SplitTrimNonWhitespace(",")
		if len(parts2) < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_CovSSO_29_LinesSimpleSlice_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_CovSSO_29_LinesSimpleSlice_SimpleSlice", func() {
		sso := newSSO("a\nb")
		ss := sso.LinesSimpleSlice()
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
		ss2 := sso.SimpleSlice(",")
		if ss2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSSO_30_ConcatNew_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_CovSSO_30_ConcatNew_ConcatNewUsingStrings", func() {
		sso := newSSO("hello")
		c := sso.ConcatNew(" world")
		if c.Value() != "hello world" {
			t.Fatal("expected hello world")
		}
		c2 := sso.ConcatNewUsingStrings("-", "a", "b")
		if c2.Value() != "hello-a-b" {
			t.Fatalf("expected 'hello-a-b', got '%s'", c2.Value())
		}
	})
}

func Test_CovSSO_31_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_CovSSO_31_NonPtr_Ptr", func() {
		sso := newSSO("hello")
		np := sso.NonPtr()
		if np.Value() != "hello" {
			t.Fatal("expected hello")
		}
		p := sso.Ptr()
		if p.Value() != "hello" {
			t.Fatal("expected hello")
		}
	})
}

func Test_CovSSO_32_Clone_ClonePtr_CloneUsingNewVal(t *testing.T) {
	safeTest(t, "Test_CovSSO_32_Clone_ClonePtr_CloneUsingNewVal", func() {
		sso := newSSO("hello")
		c := sso.Clone()
		if c.Value() != "hello" {
			t.Fatal("expected hello")
		}
		cp := sso.ClonePtr()
		if cp.Value() != "hello" {
			t.Fatal("expected hello")
		}
		cv := sso.CloneUsingNewVal("new")
		if cv.Value() != "new" {
			t.Fatal("expected new")
		}
	})
}

func Test_CovSSO_33_String_StringPtr(t *testing.T) {
	safeTest(t, "Test_CovSSO_33_String_StringPtr", func() {
		sso := newSSO("hello")
		if sso.String() != "hello" {
			t.Fatal("expected hello")
		}
		sp := sso.StringPtr()
		if *sp != "hello" {
			t.Fatal("expected hello")
		}
	})
}

func Test_CovSSO_34_Dispose(t *testing.T) {
	safeTest(t, "Test_CovSSO_34_Dispose", func() {
		sso := newSSO("hello")
		sso.Dispose()
	})
}

func Test_CovSSO_35_JsonModel_MarshalUnmarshal(t *testing.T) {
	safeTest(t, "Test_CovSSO_35_JsonModel_MarshalUnmarshal", func() {
		sso := newSSO("hello")
		_ = sso.JsonModel()
		_ = sso.JsonModelAny()
		data, err := sso.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		sso2 := &corestr.SimpleStringOnce{}
		err2 := sso2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		// invalid
		err3 := sso2.UnmarshalJSON([]byte("bad"))
		if err3 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovSSO_36_Json_JsonPtr_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovSSO_36_Json_JsonPtr_ParseInject", func() {
		sso := newSSO("hello")
		_ = sso.Json()
		jr := sso.JsonPtr()
		sso2 := &corestr.SimpleStringOnce{}
		r, err := sso2.ParseInjectUsingJson(jr)
		if err != nil || r == nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovSSO_37_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovSSO_37_ParseInjectUsingJsonMust", func() {
		sso := newSSO("hello")
		jr := sso.JsonPtr()
		sso2 := &corestr.SimpleStringOnce{}
		r := sso2.ParseInjectUsingJsonMust(jr)
		if r == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovSSO_38_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovSSO_38_JsonParseSelfInject", func() {
		sso := newSSO("hello")
		jr := sso.JsonPtr()
		sso2 := &corestr.SimpleStringOnce{}
		err := sso2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovSSO_39_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovSSO_39_AsInterfaces", func() {
		sso := newSSO("hello")
		_ = sso.AsJsonContractsBinder()
		_ = sso.AsJsoner()
		_ = sso.AsJsonParseSelfInjector()
		_ = sso.AsJsonMarshaller()
	})
}

func Test_CovSSO_40_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovSSO_40_Serialize_Deserialize", func() {
		sso := newSSO("hello")
		_, err := sso.Serialize()
		if err != nil {
			t.Fatal("unexpected error")
		}
		target := &corestr.SimpleStringOnce{}
		err2 := sso.Deserialize(target)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
	})
}
