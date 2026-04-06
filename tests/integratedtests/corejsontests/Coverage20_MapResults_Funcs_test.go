package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// MapResults — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_MapResults_BasicOps(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	if !m.IsEmpty() || m.HasAnyItem() || m.Length() != 0 || m.LastIndex() != -1 {
		t.Fatal("basic checks failed")
	}
	m.Add("k", corejson.NewResult.Any("v"))
	if m.Length() != 1 || m.IsEmpty() || !m.HasAnyItem() {
		t.Fatal("filled checks failed")
	}
}

func Test_Cov20_MapResults_AddMethods(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.AddSkipOnNil("k", nil)
	r := corejson.NewResult.AnyPtr("x")
	m.AddSkipOnNil("k", r)
	m.AddPtr("k2", nil)
	m.AddPtr("k2", r)
	err := m.AddAny("k3", "hello")
	if err != nil {
		t.Fatal("unexpected error")
	}
	err = m.AddAny("k4", nil)
	if err == nil {
		t.Fatal("expected error for nil")
	}
	err = m.AddAnySkipOnNil("k5", nil)
	if err != nil {
		t.Fatal("expected nil error for skip nil")
	}
	err = m.AddAnySkipOnNil("k5", "hello")
	if err != nil {
		t.Fatal("unexpected error")
	}
	m.AddAnyNonEmptyNonError("k6", nil)
	m.AddAnyNonEmptyNonError("k6", "hello")
	m.AddAnyNonEmpty("k7", nil)
	m.AddAnyNonEmpty("k7", "world")
}

func Test_Cov20_MapResults_GetByKey(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	if m.GetByKey("k") == nil {
		t.Fatal("expected result")
	}
	if m.GetByKey("missing") != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov20_MapResults_Errors(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("ok", corejson.NewResult.Any("v"))
	m.Add("err", corejson.NewResult.Error(errors.New("e")))
	if !m.HasError() {
		t.Fatal("expected error")
	}
	errs, hasErr := m.AllErrors()
	if !hasErr || len(errs) == 0 {
		t.Fatal("expected errors")
	}
	strs := m.GetErrorsStrings()
	if len(strs) == 0 {
		t.Fatal("expected strings")
	}
	_ = m.GetErrorsStringsPtr()
	_ = m.GetErrorsAsSingleString()
	_ = m.GetErrorsAsSingle()
}

func Test_Cov20_MapResults_Keys(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("b", corejson.NewResult.Any("2"))
	m.Add("a", corejson.NewResult.Any("1"))
	keys := m.AllKeys()
	if len(keys) != 2 {
		t.Fatal("expected 2 keys")
	}
	sorted := m.AllKeysSorted()
	if sorted[0] != "a" {
		t.Fatal("expected sorted")
	}
}

func Test_Cov20_MapResults_AllValues(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	vals := m.AllValues()
	if len(vals) != 1 {
		t.Fatal("expected 1")
	}
	_ = m.AllResults()
	_ = m.AllResultsCollection()
}

func Test_Cov20_MapResults_GetStrings(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	strs := m.GetStrings()
	if len(strs) != 1 {
		t.Fatal("expected 1")
	}
	_ = m.GetStringsPtr()
}

func Test_Cov20_MapResults_AddKeyWithResult(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.AddKeyWithResult(corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
	m.AddKeyWithResultPtr(nil)
	m.AddKeyWithResultPtr(&corejson.KeyWithResult{Key: "k2", Result: corejson.NewResult.Any("v2")})
	m.AddKeysWithResults(corejson.KeyWithResult{Key: "k3", Result: corejson.NewResult.Any("v3")})
	m.AddKeysWithResultsPtr(&corejson.KeyWithResult{Key: "k4", Result: corejson.NewResult.Any("v4")}, nil)
	if m.Length() < 4 {
		t.Fatal("expected at least 4")
	}
}

func Test_Cov20_MapResults_AddKeyAny(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.AddKeyAnyInf(corejson.KeyAny{Key: "k", AnyInf: "v"})
	m.AddKeyAnyInfPtr(nil)
	m.AddKeyAnyInfPtr(&corejson.KeyAny{Key: "k2", AnyInf: "v2"})
	m.AddKeyAnyItems(corejson.KeyAny{Key: "k3", AnyInf: "v3"})
	m.AddKeyAnyItemsPtr(&corejson.KeyAny{Key: "k4", AnyInf: "v4"}, nil)
}

func Test_Cov20_MapResults_AddNonEmptyNonErrorPtr(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.AddNonEmptyNonErrorPtr("k", nil)
	m.AddNonEmptyNonErrorPtr("k", corejson.NewResult.ErrorPtr(errors.New("e")))
	r := corejson.NewResult.AnyPtr("v")
	m.AddNonEmptyNonErrorPtr("k", r)
	if m.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov20_MapResults_AddMapResults(t *testing.T) {
	m1 := corejson.NewMapResults.Empty()
	m1.Add("k", corejson.NewResult.Any("v"))
	m2 := corejson.NewMapResults.Empty()
	m2.AddMapResults(m1)
	if m2.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov20_MapResults_AddMapAnyItems(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.AddMapAnyItems(map[string]any{"k": "v"})
	if m.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov20_MapResults_AddMapResultsUsingCloneOption(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	items := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	m.AddMapResultsUsingCloneOption(false, false, items)
	m.AddMapResultsUsingCloneOption(true, true, items)
}

func Test_Cov20_MapResults_Paging(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	for i := 0; i < 10; i++ {
		m.Add("k"+string(rune('a'+i)), corejson.NewResult.Any("v"))
	}
	if m.GetPagesSize(3) != 4 {
		t.Fatal("expected 4")
	}
	pages := m.GetPagedCollection(3)
	if len(pages) != 4 {
		t.Fatalf("expected 4, got %d", len(pages))
	}
}

func Test_Cov20_MapResults_GetNewMapUsingKeys(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("a", corejson.NewResult.Any("1"))
	m.Add("b", corejson.NewResult.Any("2"))
	sub := m.GetNewMapUsingKeys(false, "a")
	if sub.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov20_MapResults_ResultCollection(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	rc := m.ResultCollection()
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov20_MapResults_ClearDispose(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	m.Clear()
	m.Add("k2", corejson.NewResult.Any("v2"))
	m.Dispose()
}

func Test_Cov20_MapResults_JsonOps(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	_ = m.JsonModel()
	_ = m.JsonModelAny()
	_ = m.Json()
	_ = m.JsonPtr()
	_ = m.AsJsonContractsBinder()
	_ = m.AsJsoner()
	_ = m.AsJsonParseSelfInjector()
}

func Test_Cov20_MapResults_AddJsoner(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	r := corejson.NewResult.Any("x")
	m.AddJsoner("k", &r)
	m.AddJsoner("k2", nil) // skip
	if m.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov20_MapResults_AddKeyWithJsoner(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	r := corejson.NewResult.Any("x")
	m.AddKeyWithJsoner(corejson.KeyWithJsoner{Key: "k", Jsoner: &r})
}

func Test_Cov20_MapResults_AddKeyWithJsonerPtr(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.AddKeyWithJsonerPtr(nil)
	r := corejson.NewResult.Any("x")
	m.AddKeyWithJsonerPtr(&corejson.KeyWithJsoner{Key: "k", Jsoner: &r})
}

func Test_Cov20_MapResults_InjectIntoAt(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any([]string{"a"}))
	r := corejson.NewResult.Any("x")
	err := m.InjectIntoAt("k", &r)
	_ = err
}

func Test_Cov20_MapResults_Unmarshal(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("hello"))
	var s string
	err := m.Unmarshal("missing", &s)
	_ = err
	_ = m.Deserialize("k", &s)
}

func Test_Cov20_MapResults_SafeUnmarshal(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("hello"))
	var s string
	err := m.SafeUnmarshal("k", &s)
	_ = err
	_ = m.SafeDeserialize("k", &s)
}

func Test_Cov20_MapResults_ParseInjectUsingJson(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	jr := m.JsonPtr()
	target := corejson.NewMapResults.Empty()
	_, err := target.ParseInjectUsingJson(jr)
	_ = err
}

// ══════════════════════════════════════════════════════════════════════════════
// Package-level functions — BytesCloneIf, BytesDeepClone, BytesToString, etc.
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_BytesCloneIf_DeepClone(t *testing.T) {
	b := []byte("hello")
	c := corejson.BytesCloneIf(true, b)
	if string(c) != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_Cov20_BytesCloneIf_NoClone(t *testing.T) {
	b := []byte("hello")
	c := corejson.BytesCloneIf(false, b)
	if len(c) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_Cov20_BytesCloneIf_Empty(t *testing.T) {
	c := corejson.BytesCloneIf(true, []byte{})
	if len(c) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_Cov20_BytesDeepClone(t *testing.T) {
	c := corejson.BytesDeepClone([]byte("hello"))
	if string(c) != "hello" {
		t.Fatal("expected hello")
	}
	e := corejson.BytesDeepClone(nil)
	if len(e) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_Cov20_BytesToString(t *testing.T) {
	s := corejson.BytesToString([]byte("hello"))
	if s != "hello" {
		t.Fatal("expected hello")
	}
	e := corejson.BytesToString(nil)
	if e != "" {
		t.Fatal("expected empty")
	}
}

func Test_Cov20_BytesToPrettyString(t *testing.T) {
	s := corejson.BytesToPrettyString([]byte(`{"a":1}`))
	if s == "" {
		t.Fatal("expected non-empty")
	}
	e := corejson.BytesToPrettyString(nil)
	if e != "" {
		t.Fatal("expected empty")
	}
}

func Test_Cov20_JsonString_Func(t *testing.T) {
	s, err := corejson.JsonString("hello")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_Cov20_JsonStringOrErrMsg_Valid(t *testing.T) {
	s := corejson.JsonStringOrErrMsg("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov20_JsonStringOrErrMsg_Error(t *testing.T) {
	ch := make(chan int)
	s := corejson.JsonStringOrErrMsg(ch)
	if s == "" {
		t.Fatal("expected error message")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// anyTo — additional coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_AnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw("hello")
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func Test_Cov20_AnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString("hello")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func Test_Cov20_AnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov20_AnyTo_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov20_AnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov20_AnyTo_PrettyStringWithError(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
	// bytes
	s2, err2 := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	if err2 != nil || s2 == "" {
		t.Fatal("unexpected")
	}
	// Result
	r := corejson.NewResult.Any("x")
	s3, err3 := corejson.AnyTo.PrettyStringWithError(r)
	if err3 != nil || s3 == "" {
		t.Fatal("unexpected")
	}
	// *Result
	rp := corejson.NewResult.AnyPtr("x")
	s4, err4 := corejson.AnyTo.PrettyStringWithError(rp)
	if err4 != nil || s4 == "" {
		t.Fatal("unexpected")
	}
	// any
	s5, err5 := corejson.AnyTo.PrettyStringWithError(map[string]int{"a": 1})
	if err5 != nil || s5 == "" {
		t.Fatal("unexpected")
	}
}

func Test_Cov20_AnyTo_SafeJsonPrettyString(t *testing.T) {
	_ = corejson.AnyTo.SafeJsonPrettyString("hello")
	_ = corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	r := corejson.NewResult.Any("x")
	_ = corejson.AnyTo.SafeJsonPrettyString(r)
	rp := corejson.NewResult.AnyPtr("x")
	_ = corejson.AnyTo.SafeJsonPrettyString(rp)
	_ = corejson.AnyTo.SafeJsonPrettyString(42)
}

func Test_Cov20_AnyTo_JsonString(t *testing.T) {
	_ = corejson.AnyTo.JsonString("hello")
	_ = corejson.AnyTo.JsonString([]byte(`"x"`))
	r := corejson.NewResult.Any("x")
	_ = corejson.AnyTo.JsonString(r)
	rp := corejson.NewResult.AnyPtr("x")
	_ = corejson.AnyTo.JsonString(rp)
	_ = corejson.AnyTo.JsonString(42)
}

func Test_Cov20_AnyTo_JsonStringWithErr(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("hello")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
	s2, err2 := corejson.AnyTo.JsonStringWithErr([]byte(`"x"`))
	if err2 != nil || s2 == "" {
		t.Fatal("unexpected")
	}
	r := corejson.NewResult.Any("x")
	_, _ = corejson.AnyTo.JsonStringWithErr(r)
	rp := corejson.NewResult.AnyPtr("x")
	_, _ = corejson.AnyTo.JsonStringWithErr(rp)
	_, _ = corejson.AnyTo.JsonStringWithErr(42)
}

func Test_Cov20_AnyTo_JsonStringMust(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov20_AnyTo_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov20_AnyTo_UsingSerializer(t *testing.T) {
	r := corejson.AnyTo.UsingSerializer(nil)
	if r != nil {
		t.Fatal("expected nil for nil serializer")
	}
}

func Test_Cov20_AnyTo_SerializedFieldsMap(t *testing.T) {
	fm, err := corejson.AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
	_ = fm
	_ = err
}

func Test_Cov20_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(nil)
	if r == nil || !r.HasError() {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov20_AnyTo_SerializedJsonResult_Jsoner(t *testing.T) {
	inner := corejson.NewResult.Any("x")
	r := corejson.AnyTo.SerializedJsonResult(&inner)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov20_AnyTo_SerializedJsonResult_Error(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(errors.New("hello"))
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov20_AnyTo_SerializedJsonResult_EmptyError(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(errors.New(""))
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Serializer — additional coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_Serializer_Methods(t *testing.T) {
	_ = corejson.Serialize.StringsApply([]string{"a"})
	_ = corejson.Serialize.FromBytes([]byte(`"x"`))
	_ = corejson.Serialize.FromStrings([]string{"a"})
	_ = corejson.Serialize.FromStringsSpread("a", "b")
	_ = corejson.Serialize.FromString("hello")
	_ = corejson.Serialize.FromInteger(42)
	_ = corejson.Serialize.FromInteger64(42)
	_ = corejson.Serialize.FromBool(true)
	_ = corejson.Serialize.FromIntegers([]int{1, 2})
	_ = corejson.Serialize.UsingAnyPtr("hello")
	_ = corejson.Serialize.UsingAny("hello")
	_, _ = corejson.Serialize.Raw("hello")
	_, _ = corejson.Serialize.Marshal("hello")
	_ = corejson.Serialize.ApplyMust("hello")
	_ = corejson.Serialize.ToBytesMust("hello")
	_ = corejson.Serialize.ToSafeBytesMust("hello")
	_ = corejson.Serialize.ToSafeBytesSwallowErr("hello")
	_ = corejson.Serialize.ToBytesSwallowErr("hello")
	_, _ = corejson.Serialize.ToBytesErr("hello")
	_ = corejson.Serialize.ToString("hello")
	_ = corejson.Serialize.ToStringMust("hello")
	_, _ = corejson.Serialize.ToStringErr("hello")
	_, _ = corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	_ = corejson.Serialize.ToPrettyStringIncludingErr("hello")
	_ = corejson.Serialize.Pretty("hello")
}

// ══════════════════════════════════════════════════════════════════════════════
// Deserializer — additional coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_Deserializer_Methods(t *testing.T) {
	var s string
	_ = corejson.Deserialize.UsingStringPtr(nil, &s)
	str := `"hello"`
	_ = corejson.Deserialize.UsingStringPtr(&str, &s)
	_ = corejson.Deserialize.UsingError(nil, &s)
	_ = corejson.Deserialize.UsingError(errors.New(`"world"`), &s)
	_ = corejson.Deserialize.UsingResult(corejson.NewResult.AnyPtr("x"), &s)
	corejson.Deserialize.ApplyMust(corejson.NewResult.AnyPtr("hello"), &s)
	_ = corejson.Deserialize.FromString(`"x"`, &s)
	corejson.Deserialize.FromStringMust(`"y"`, &s)
	_ = corejson.Deserialize.UsingStringOption(true, "", &s)
	_ = corejson.Deserialize.UsingStringOption(false, `"x"`, &s)
	_ = corejson.Deserialize.UsingStringIgnoreEmpty("", &s)
	_ = corejson.Deserialize.UsingStringIgnoreEmpty(`"x"`, &s)
	corejson.Deserialize.UsingBytesMust([]byte(`"hello"`), &s)
	_ = corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &s)
	_ = corejson.Deserialize.UsingBytesIf(true, []byte(`"x"`), &s)
	_ = corejson.Deserialize.UsingBytesPointer([]byte(`"x"`), &s)
	_ = corejson.Deserialize.UsingBytesPointer(nil, &s)
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"hello"`), &s)
	_ = corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &s)
	_ = corejson.Deserialize.UsingBytesPointerIf(true, []byte(`"x"`), &s)
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &s)
	corejson.Deserialize.UsingSafeBytesMust([]byte(`"safe"`), &s)
	_, _ = corejson.Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
	_ = corejson.Deserialize.MapAnyToPointer(true, nil, &s)
	_ = corejson.Deserialize.MapAnyToPointer(false, map[string]any{"a": "b"}, &s)
	_ = corejson.Deserialize.UsingDeserializerToOption(true, nil, &s)
	_ = corejson.Deserialize.UsingDeserializerToOption(false, nil, &s)
	_ = corejson.Deserialize.UsingDeserializerDefined(nil, &s)
	_ = corejson.Deserialize.UsingDeserializerFuncDefined(nil, &s)
	_ = corejson.Deserialize.UsingDeserializerFuncDefined(func(toPtr any) error { return nil }, &s)
	_ = corejson.Deserialize.UsingJsonerToAny(true, nil, &s)
	_ = corejson.Deserialize.UsingJsonerToAny(false, nil, &s)
	_ = corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &s)
	_ = corejson.Deserialize.FromTo("hello", &s)
	_ = corejson.Deserialize.UsingErrorWhichJsonResult(nil, &s)
}

// ══════════════════════════════════════════════════════════════════════════════
// NewResult creator — additional coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_NewResult_Various(t *testing.T) {
	_ = corejson.NewResult.UsingBytes([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesType([]byte(`"x"`), "T")
	_ = corejson.NewResult.UsingBytesTypePtr([]byte(`"x"`), "T")
	_ = corejson.NewResult.UsingTypeBytesPtr("T", []byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtr(nil)
	_ = corejson.NewResult.UsingBytesPtr([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "T")
	_ = corejson.NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "T")
	_ = corejson.NewResult.UsingBytesErrPtr(nil, errors.New("e"), "T")
	_ = corejson.NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "T")
	_ = corejson.NewResult.PtrUsingStringPtr(nil, "T")
	str := `"hello"`
	_ = corejson.NewResult.PtrUsingStringPtr(&str, "T")
	_ = corejson.NewResult.UsingErrorStringPtr(nil, &str, "T")
	_ = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "T")
	_ = corejson.NewResult.Ptr([]byte(`"x"`), nil, "T")
	_ = corejson.NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "T")
	_ = corejson.NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = corejson.NewResult.UsingTypePlusString("T", `"x"`)
	_ = corejson.NewResult.UsingTypePlusStringPtr("T", nil)
	_ = corejson.NewResult.UsingTypePlusStringPtr("T", &str)
	_ = corejson.NewResult.UsingStringWithType(`"x"`, "T")
	_ = corejson.NewResult.UsingString(`"x"`)
	_ = corejson.NewResult.UsingStringPtr(nil)
	_ = corejson.NewResult.UsingStringPtr(&str)
	_ = corejson.NewResult.CreatePtr([]byte(`"x"`), nil, "T")
	_ = corejson.NewResult.NonPtr([]byte(`"x"`), nil, "T")
	_ = corejson.NewResult.Create([]byte(`"x"`), nil, "T")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "T")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, nil, "T")
	_ = corejson.NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "T")
	_ = corejson.NewResult.CastingAny("hello")
	_ = corejson.NewResult.Error(errors.New("e"))
	_ = corejson.NewResult.ErrorPtr(errors.New("e"))
	_ = corejson.NewResult.Empty()
	_ = corejson.NewResult.EmptyPtr()
	_ = corejson.NewResult.TypeName("T")
	_ = corejson.NewResult.TypeNameBytes("T")
	_ = corejson.NewResult.Many("a", "b")
	_ = corejson.NewResult.Serialize("hello")
	_ = corejson.NewResult.Marshal("hello")
	_ = corejson.NewResult.UsingSerializer(nil)
	_ = corejson.NewResult.UsingSerializerFunc(nil)
	_ = corejson.NewResult.UsingJsoner(nil)
	_ = corejson.NewResult.AnyToCastingResult("hello")
	_ = corejson.NewResult.UnmarshalUsingBytes([]byte(`{}`))
	_ = corejson.NewResult.DeserializeUsingBytes([]byte(`{}`))
}

// ══════════════════════════════════════════════════════════════════════════════
// CastAny — additional coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_CastAny_FromToDefault(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToDefault([]byte(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_Cov20_CastAny_FromToReflection(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToReflection([]byte(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_Cov20_CastAny_OrDeserializeTo(t *testing.T) {
	var out string
	err := corejson.CastAny.OrDeserializeTo([]byte(`"hello"`), &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_Cov20_CastAny_FromToOption_Jsoner(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	var out string
	err := corejson.CastAny.FromToOption(false, &r, &out)
	_ = err
}

func Test_Cov20_CastAny_FromToOption_SerializerFunc(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var out string
	err := corejson.CastAny.FromToOption(false, fn, &out)
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func Test_Cov20_CastAny_FromToOption_Serializer(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	var out string
	err := corejson.CastAny.FromToOption(false, r, &out)
	_ = err
}

func Test_Cov20_CastAny_FromToOption_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var out string
	err := corejson.CastAny.FromToOption(false, r, &out)
	_ = err
}

func Test_Cov20_CastAny_FromToOption_String(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToOption(false, `"hello"`, &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Empty creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov20_EmptyCreator(t *testing.T) {
	_ = corejson.Empty.Result()
	_ = corejson.Empty.ResultWithErr("T", errors.New("e"))
	_ = corejson.Empty.ResultPtrWithErr("T", errors.New("e"))
	_ = corejson.Empty.ResultPtr()
	_ = corejson.Empty.BytesCollection()
	_ = corejson.Empty.BytesCollectionPtr()
	_ = corejson.Empty.ResultsCollection()
	_ = corejson.Empty.ResultsPtrCollection()
	_ = corejson.Empty.MapResults()
}
