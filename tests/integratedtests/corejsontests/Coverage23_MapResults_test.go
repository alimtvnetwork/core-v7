package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

func Test_C23_MR_Length_Nil(t *testing.T) {
	var mr *corejson.MapResults
	if mr.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_LastIndex(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	if mr.LastIndex() != -1 { t.Fatal("expected -1") }
}

func Test_C23_MR_IsEmpty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	if !mr.IsEmpty() { t.Fatal("expected true") }
}

func Test_C23_MR_HasAnyItem(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	if !mr.HasAnyItem() { t.Fatal("expected true") }
}

func Test_C23_MR_AddSkipOnNil_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddSkipOnNil("k", nil)
	if mr.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_AddSkipOnNil_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.New("v")
	mr.AddSkipOnNil("k", r.Ptr())
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C23_MR_GetByKey_Found(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	r := mr.GetByKey("k")
	if r == nil { t.Fatal("expected non-nil") }
}

func Test_C23_MR_GetByKey_NotFound(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := mr.GetByKey("missing")
	if r != nil { t.Fatal("expected nil") }
}

func Test_C23_MR_HasError_Yes(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Create(nil, errors.New("err"), ""))
	if !mr.HasError() { t.Fatal("expected true") }
}

func Test_C23_MR_HasError_No(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	if mr.HasError() { t.Fatal("expected false") }
}

func Test_C23_MR_AllErrors_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	errs, has := mr.AllErrors()
	if has || len(errs) != 0 { t.Fatal("unexpected") }
}

func Test_C23_MR_AllErrors_WithErrors(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Create(nil, errors.New("err"), ""))
	errs, has := mr.AllErrors()
	if !has || len(errs) != 1 { t.Fatal("unexpected") }
}

func Test_C23_MR_GetErrorsStrings_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	s := mr.GetErrorsStrings()
	if len(s) != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_GetErrorsStrings_WithErrors(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Create(nil, errors.New("err"), ""))
	s := mr.GetErrorsStrings()
	if len(s) != 1 { t.Fatal("expected 1") }
}

func Test_C23_MR_GetErrorsStringsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsStringsPtr()
}

func Test_C23_MR_GetErrorsAsSingleString(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsAsSingleString()
}

func Test_C23_MR_GetErrorsAsSingle(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.GetErrorsAsSingle()
	_ = err
}

func Test_C23_MR_Unmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	var s string
	err := mr.Unmarshal("k", &s)
	_ = err
}

func Test_C23_MR_Deserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	var s string
	_ = mr.Deserialize("k", &s)
}

func Test_C23_MR_DeserializeMust(t *testing.T) {
	defer func() { recover() }()
	mr := corejson.NewMapResults.Empty()
	var s string
	mr.DeserializeMust("missing", &s)
}

func Test_C23_MR_UnmarshalMany_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.UnmarshalMany()
	if err != nil { t.Fatal(err) }
}

func Test_C23_MR_UnmarshalMany_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	var s string
	err := mr.UnmarshalMany(corejson.KeyAny{Key: "k", AnyInf: &s})
	_ = err
}

func Test_C23_MR_UnmarshalManySafe_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.UnmarshalManySafe()
	if err != nil { t.Fatal(err) }
}

func Test_C23_MR_SafeUnmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	var s string
	_ = mr.SafeUnmarshal("k", &s)
}

func Test_C23_MR_SafeDeserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.SafeDeserialize("k", nil)
}

func Test_C23_MR_SafeDeserializeMust(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.SafeDeserializeMust("k", nil)
}

func Test_C23_MR_InjectIntoAt(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New(map[string]string{"a": "b"}))
	target := corejson.Empty.MapResults()
	err := mr.InjectIntoAt("k", target)
	_ = err
}

func Test_C23_MR_AddPtr_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddPtr("k", nil)
	if mr.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_AddAny_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAny("k", nil)
	if err == nil { t.Fatal("expected error") }
}

func Test_C23_MR_AddAny_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAny("k", "v")
	if err != nil { t.Fatal(err) }
}

func Test_C23_MR_AddAny_Error(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAny("k", make(chan int))
	if err == nil { t.Fatal("expected error") }
}

func Test_C23_MR_AddAnySkipOnNil_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAnySkipOnNil("k", nil)
	if err != nil { t.Fatal(err) }
}

func Test_C23_MR_AddAnySkipOnNil_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAnySkipOnNil("k", "v")
	if err != nil { t.Fatal(err) }
}

func Test_C23_MR_AddAnyNonEmptyNonError(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmptyNonError("k", nil)
	mr.AddAnyNonEmptyNonError("k2", "v")
	_ = mr
}

func Test_C23_MR_AddAnyNonEmpty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmpty("k", nil)
	mr.AddAnyNonEmpty("k2", "v")
}

func Test_C23_MR_AddKeyWithResult(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithResult(corejson.KeyWithResult{Key: "k", Result: corejson.New("v")})
}

func Test_C23_MR_AddKeyWithResultPtr_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithResultPtr(nil)
	if mr.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_AddKeysWithResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResults(corejson.KeyWithResult{Key: "k", Result: corejson.New("v")})
}

func Test_C23_MR_AddKeysWithResults_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResults()
}

func Test_C23_MR_AddKeysWithResultsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResultsPtr()
}

func Test_C23_MR_AddKeyAnyInf(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInf(corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func Test_C23_MR_AddKeyAnyInfPtr_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInfPtr(nil)
}

func Test_C23_MR_AddKeyAnyInfPtr_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInfPtr(&corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func Test_C23_MR_AddKeyAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItems(corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func Test_C23_MR_AddKeyAnyItems_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItems()
}

func Test_C23_MR_AddKeyAnyItemsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItemsPtr()
}

func Test_C23_MR_AddNonEmptyNonErrorPtr_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddNonEmptyNonErrorPtr("k", nil)
}

func Test_C23_MR_AddNonEmptyNonErrorPtr_Error(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.NewResult.ErrorPtr(errors.New("err"))
	mr.AddNonEmptyNonErrorPtr("k", r)
	if mr.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_AddNonEmptyNonErrorPtr_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.New("v")
	mr.AddNonEmptyNonErrorPtr("k", r.Ptr())
}

func Test_C23_MR_AddMapResults_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResults(nil)
}

func Test_C23_MR_AddMapResults_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	other := corejson.NewMapResults.Empty()
	other.Add("k", corejson.New("v"))
	mr.AddMapResults(other)
}

func Test_C23_MR_AddMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapAnyItems(map[string]any{"k": "v"})
}

func Test_C23_MR_AddMapAnyItems_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapAnyItems(map[string]any{})
}

func Test_C23_MR_AllKeys(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.New("1"))
	mr.Add("b", corejson.New("2"))
	keys := mr.AllKeys()
	if len(keys) != 2 { t.Fatal("expected 2") }
}

func Test_C23_MR_AllKeys_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	keys := mr.AllKeys()
	if len(keys) != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_AllKeysSorted(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("b", corejson.New("2"))
	mr.Add("a", corejson.New("1"))
	keys := mr.AllKeysSorted()
	if keys[0] != "a" { t.Fatal("expected sorted") }
}

func Test_C23_MR_AllKeysSorted_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	keys := mr.AllKeysSorted()
	if len(keys) != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_AllValues(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	vals := mr.AllValues()
	if len(vals) != 1 { t.Fatal("expected 1") }
}

func Test_C23_MR_AllValues_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	vals := mr.AllValues()
	if len(vals) != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_AllResultsCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	rc := mr.AllResultsCollection()
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C23_MR_AllResultsCollection_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	rc := mr.AllResultsCollection()
	if rc.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_AllResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AllResults()
}

func Test_C23_MR_GetStrings(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	s := mr.GetStrings()
	if len(s) != 1 { t.Fatal("expected 1") }
}

func Test_C23_MR_GetStrings_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	s := mr.GetStrings()
	if len(s) != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_GetStringsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetStringsPtr()
}

func Test_C23_MR_AddJsoner_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddJsoner("k", nil)
	if mr.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_AddKeyWithJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.New("v")
	mr.AddKeyWithJsoner(corejson.KeyWithJsoner{Key: "k", Jsoner: &r})
}

func Test_C23_MR_AddKeysWithJsoners(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithJsoners()
}

func Test_C23_MR_AddKeyWithJsonerPtr_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithJsonerPtr(nil)
}

func Test_C23_MR_AddKeyWithJsonerPtr_NilJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithJsonerPtr(&corejson.KeyWithJsoner{Key: "k", Jsoner: nil})
}

func Test_C23_MR_AddMapResultsUsingCloneOption_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResultsUsingCloneOption(false, false, map[string]corejson.Result{})
}

func Test_C23_MR_AddMapResultsUsingCloneOption_NoClone(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResultsUsingCloneOption(false, false, map[string]corejson.Result{
		"k": corejson.New("v"),
	})
}

func Test_C23_MR_AddMapResultsUsingCloneOption_Clone(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResultsUsingCloneOption(true, true, map[string]corejson.Result{
		"k": corejson.New("v"),
	})
}

func Test_C23_MR_GetPagesSize_Zero(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	if mr.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_GetPagesSize_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 5; i++ { mr.Add(string(rune('a'+i)), corejson.New(i)) }
	p := mr.GetPagesSize(2)
	if p != 3 { t.Fatalf("expected 3, got %d", p) }
}

func Test_C23_MR_GetPagedCollection_Small(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.New("1"))
	result := mr.GetPagedCollection(5)
	if len(result) != 1 { t.Fatal("expected 1") }
}

func Test_C23_MR_GetPagedCollection_Multi(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 5; i++ { mr.Add(string(rune('a'+i)), corejson.New(i)) }
	result := mr.GetPagedCollection(2)
	if len(result) != 3 { t.Fatalf("expected 3, got %d", len(result)) }
}

func Test_C23_MR_GetNewMapUsingKeys_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	result := mr.GetNewMapUsingKeys(false)
	if result.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_GetNewMapUsingKeys_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.New("1"))
	mr.Add("b", corejson.New("2"))
	result := mr.GetNewMapUsingKeys(false, "a")
	if result.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C23_MR_ResultCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	rc := mr.ResultCollection()
	_ = rc
}

func Test_C23_MR_ResultCollection_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	rc := mr.ResultCollection()
	if rc.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C23_MR_JsonModel(t *testing.T) { _ = corejson.NewMapResults.Empty().JsonModel() }
func Test_C23_MR_JsonModelAny(t *testing.T) { _ = corejson.NewMapResults.Empty().JsonModelAny() }

func Test_C23_MR_Clear(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	mr.Clear()
}

func Test_C23_MR_Clear_Nil(t *testing.T) {
	var mr *corejson.MapResults
	mr.Clear()
}

func Test_C23_MR_Dispose(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Dispose()
}

func Test_C23_MR_Dispose_Nil(t *testing.T) {
	var mr *corejson.MapResults
	mr.Dispose()
}

func Test_C23_MR_Json(t *testing.T) { _ = corejson.NewMapResults.Empty().Json() }
func Test_C23_MR_JsonPtr(t *testing.T) { _ = corejson.NewMapResults.Empty().JsonPtr() }

func Test_C23_MR_ParseInjectUsingJson_Error(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	_, err := mr.ParseInjectUsingJson(bad)
	if err == nil { t.Fatal("expected error") }
}

func Test_C23_MR_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	mr := corejson.NewMapResults.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	mr.ParseInjectUsingJsonMust(bad)
}

func Test_C23_MR_AsJsonContractsBinder(t *testing.T) { _ = corejson.NewMapResults.Empty().AsJsonContractsBinder() }
func Test_C23_MR_AsJsoner(t *testing.T) { _ = corejson.NewMapResults.Empty().AsJsoner() }
func Test_C23_MR_AsJsonParseSelfInjector(t *testing.T) { _ = corejson.NewMapResults.Empty().AsJsonParseSelfInjector() }
func Test_C23_MR_JsonParseSelfInject(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.New(*mr)
	_ = mr.JsonParseSelfInject(&r)
}
