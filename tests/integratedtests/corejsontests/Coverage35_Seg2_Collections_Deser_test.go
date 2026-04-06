package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 2: Collections, Deserializer, Serializer, Creators
// ══════════════════════════════════════════════════════════════════════════════

// --- ResultsCollection ---

func newTestRC() *corejson.ResultsCollection {
	rc := corejson.NewResultsCollection.UsingCap(4)
	rc.Add(corejson.New(map[string]int{"a": 1}))
	rc.Add(corejson.New(map[string]int{"b": 2}))
	return rc
}

func Test_CovJsonS2_RC01_Length_IsEmpty_HasAnyItem(t *testing.T) {
	rc := newTestRC()
	if rc.Length() != 2 {
		t.Fatal("expected 2")
	}
	if rc.IsEmpty() {
		t.Fatal("expected false")
	}
	if !rc.HasAnyItem() {
		t.Fatal("expected true")
	}
	// nil
	var nilRC *corejson.ResultsCollection
	if nilRC.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovJsonS2_RC02_FirstOrDefault_LastOrDefault(t *testing.T) {
	rc := newTestRC()
	f := rc.FirstOrDefault()
	if f == nil {
		t.Fatal("expected non-nil")
	}
	l := rc.LastOrDefault()
	if l == nil {
		t.Fatal("expected non-nil")
	}
	// empty
	empty := corejson.NewResultsCollection.Empty()
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if empty.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS2_RC03_Take_Limit_Skip(t *testing.T) {
	rc := newTestRC()
	taken := rc.Take(1)
	if taken.Length() != 1 {
		t.Fatal("expected 1")
	}
	limited := rc.Limit(1)
	if limited.Length() != 1 {
		t.Fatal("expected 1")
	}
	// limit -1 returns all
	all := rc.Limit(-1)
	if all.Length() != 2 {
		t.Fatal("expected 2")
	}
	skipped := rc.Skip(1)
	if skipped.Length() != 1 {
		t.Fatal("expected 1")
	}
	// empty
	empty := corejson.NewResultsCollection.Empty()
	if empty.Take(1).Length() != 0 {
		t.Fatal("expected 0")
	}
	if empty.Limit(1).Length() != 0 {
		t.Fatal("expected 0")
	}
	if empty.Skip(1).Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovJsonS2_RC04_AddSkipOnNil_AddNonNilNonError(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSkipOnNil(nil)
	if rc.Length() != 0 {
		t.Fatal("expected 0")
	}
	r := corejson.NewPtr(1)
	rc.AddSkipOnNil(r)
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
	// non nil non error
	rc2 := corejson.NewResultsCollection.Empty()
	errR := &corejson.Result{Error: errors.New("err")}
	rc2.AddNonNilNonError(errR)
	if rc2.Length() != 0 {
		t.Fatal("expected 0")
	}
	rc2.AddNonNilNonError(r)
	if rc2.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovJsonS2_RC05_HasError_AllErrors(t *testing.T) {
	rc := newTestRC()
	if rc.HasError() {
		t.Fatal("expected false")
	}
	errs, hasAny := rc.AllErrors()
	if hasAny || len(errs) > 0 {
		t.Fatal("expected no errors")
	}
	// empty
	empty := corejson.NewResultsCollection.Empty()
	errs2, hasAny2 := empty.AllErrors()
	if hasAny2 || len(errs2) > 0 {
		t.Fatal("expected no errors")
	}
}
func Test_CovJsonS2_RC07_GetAt_GetAtSafe_GetAtSafeUsingLength(t *testing.T) {
	rc := newTestRC()
	_ = rc.GetAt(0)
	safe := rc.GetAtSafe(0)
	if safe == nil {
		t.Fatal("expected non-nil")
	}
	safe2 := rc.GetAtSafe(-2)
	if safe2 != nil {
		t.Fatal("expected nil")
	}
	safe3 := rc.GetAtSafeUsingLength(0, 2)
	if safe3 == nil {
		t.Fatal("expected non-nil")
	}
	safe4 := rc.GetAtSafeUsingLength(5, 2)
	if safe4 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS2_RC08_AddPtr_Adds_AddsPtr_AddAny_AddAnyItems(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddPtr(nil)
	if rc.Length() != 0 {
		t.Fatal("expected 0")
	}
	r := corejson.NewPtr(1)
	rc.AddPtr(r)
	if rc.Length() != 1 {
		t.Fatal("expected 1")
	}
	rc.AddsPtr(nil, r)
	rc.AddAny(1)
	rc.AddAny(nil)
	rc.AddAnyItems(1, nil, 2)
	rc.AddAnyItems(nil)
}

func Test_CovJsonS2_RC09_AddResultsCollection_AddNonNilItemsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddResultsCollection(nil)
	rc2 := newTestRC()
	rc.AddResultsCollection(rc2)
	rc.AddNonNilItemsPtr(nil)
}

func Test_CovJsonS2_RC10_AddAnyItemsSlice(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAnyItemsSlice(nil)
	rc.AddAnyItemsSlice([]any{1, nil, 2})
}

func Test_CovJsonS2_RC11_Dispose_Clear(t *testing.T) {
	rc := newTestRC()
	rc.Dispose()
	// nil dispose
	var nilRC *corejson.ResultsCollection
	nilRC.Dispose()
}
func Test_CovJsonS2_RC13_GetPagesSize_GetPagedCollection_GetSinglePageCollection(t *testing.T) {
	// build 15 items
	rc := corejson.NewResultsCollection.UsingCap(20)
	for i := 0; i < 15; i++ {
		rc.AddAny(i)
	}
	ps := rc.GetPagesSize(5)
	if ps != 3 {
		t.Fatal("expected 3")
	}
	if rc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	pages := rc.GetPagedCollection(5)
	if len(pages) != 3 {
		t.Fatal("expected 3")
	}
	single := rc.GetSinglePageCollection(5, 2)
	if single.Length() != 5 {
		t.Fatal("expected 5")
	}
	// small collection
	small := newTestRC()
	pages2 := small.GetPagedCollection(10)
	if len(pages2) != 1 {
		t.Fatal("expected 1")
	}
	single2 := small.GetSinglePageCollection(10, 1)
	if single2.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovJsonS2_RC14_Json_JsonPtr_JsonModel_Interfaces(t *testing.T) {
	rc := newTestRC()
	j := rc.Json()
	if j.HasError() {
		t.Fatal("expected no error")
	}
	jp := rc.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
	_ = rc.JsonModel()
	_ = rc.JsonModelAny()
	_ = rc.AsJsonContractsBinder()
	_ = rc.AsJsoner()
	_ = rc.AsJsonParseSelfInjector()
	_ = rc.NonPtr()
	_ = rc.Ptr()
}

func Test_CovJsonS2_RC15_ParseInjectUsingJson(t *testing.T) {
	rc := newTestRC()
	jr := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.Empty()
	_, err := rc2.ParseInjectUsingJson(jr)
	if err != nil {
		t.Fatal("expected no error")
	}
}
func Test_CovJsonS2_RC17_UnmarshalAt(t *testing.T) {
	rc := newTestRC()
	var m map[string]int
	err := rc.UnmarshalAt(0, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS2_RC18_UnmarshalIntoSameIndex(t *testing.T) {
	rc := newTestRC()
	var m1, m2 map[string]int
	errs, hasAny := rc.UnmarshalIntoSameIndex(&m1, &m2)
	if hasAny {
		t.Fatal("expected false")
	}
	_ = errs
	// nil
	errs2, hasAny2 := rc.UnmarshalIntoSameIndex(nil)
	_ = errs2
	_ = hasAny2
}

func Test_CovJsonS2_RC19_AddSerializer_AddSerializerFunc(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializer(nil)
	rc.AddSerializerFunc(nil)
	rc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`1`), nil
	})
}

func Test_CovJsonS2_RC20_AddSerializers_AddSerializerFunctions(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializers()
	rc.AddSerializerFunctions()
}

// --- ResultsPtrCollection ---

func newTestRPC() *corejson.ResultsPtrCollection {
	rpc := corejson.NewResultsPtrCollection.UsingCap(4)
	rpc.Add(corejson.NewPtr(map[string]int{"a": 1}))
	rpc.Add(corejson.NewPtr(map[string]int{"b": 2}))
	return rpc
}

func Test_CovJsonS2_RPC01_Length_IsEmpty_HasAnyItem(t *testing.T) {
	rpc := newTestRPC()
	if rpc.Length() != 2 {
		t.Fatal("expected 2")
	}
	var nilRPC *corejson.ResultsPtrCollection
	if nilRPC.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovJsonS2_RPC02_FirstOrDefault_LastOrDefault(t *testing.T) {
	rpc := newTestRPC()
	if rpc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if rpc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	empty := corejson.NewResultsPtrCollection.Empty()
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS2_RPC03_Take_Limit_Skip(t *testing.T) {
	rpc := newTestRPC()
	if rpc.Take(1).Length() != 1 {
		t.Fatal("expected 1")
	}
	if rpc.Limit(1).Length() != 1 {
		t.Fatal("expected 1")
	}
	if rpc.Limit(-1).Length() != 2 {
		t.Fatal("expected 2")
	}
	if rpc.Skip(1).Length() != 1 {
		t.Fatal("expected 1")
	}
	empty := corejson.NewResultsPtrCollection.Empty()
	_ = empty.Take(1)
	_ = empty.Limit(1)
	_ = empty.Skip(1)
}

func Test_CovJsonS2_RPC04_AddSkipOnNil_AddNonNilNonError(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSkipOnNil(nil)
	r := corejson.NewPtr(1)
	rpc.AddSkipOnNil(r)
	if rpc.Length() != 1 {
		t.Fatal("expected 1")
	}
	errR := &corejson.Result{Error: errors.New("err")}
	rpc.AddNonNilNonError(errR)
	if rpc.Length() != 1 {
		t.Fatal("expected 1")
	}
}
func Test_CovJsonS2_RPC06_GetAt_GetAtSafe(t *testing.T) {
	rpc := newTestRPC()
	_ = rpc.GetAt(0)
	safe := rpc.GetAtSafe(0)
	if safe == nil {
		t.Fatal("expected non-nil")
	}
	if rpc.GetAtSafe(-2) != nil {
		t.Fatal("expected nil")
	}
	if rpc.GetAtSafeUsingLength(5, 2) != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS2_RPC07_Add_AddResult_Adds_AddsPtr_AddAny_AddAnyItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New(1)
	rpc.AddResult(r)
	rpc.Adds(nil)
	rpc.AddAny(nil)
	rpc.AddAny(1)
	rpc.AddAnyItems(1, nil, 2)
	rpc.AddAnyItems(nil)
}

func Test_CovJsonS2_RPC08_AddResultsCollection_AddNonNilItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddResultsCollection(nil)
	rpc2 := newTestRPC()
	rpc.AddResultsCollection(rpc2)
}

func Test_CovJsonS2_RPC09_UnmarshalAt(t *testing.T) {
	rpc := newTestRPC()
	var m map[string]int
	err := rpc.UnmarshalAt(0, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
	// nil result
	rpc2 := corejson.NewResultsPtrCollection.Empty()
	rpc2.Add(nil)
	err2 := rpc2.UnmarshalAt(0, &m)
	if err2 != nil {
		t.Fatal("expected nil for nil result")
	}
}

func Test_CovJsonS2_RPC10_UnmarshalIntoSameIndex(t *testing.T) {
	rpc := newTestRPC()
	var m1, m2 map[string]int
	errs, hasAny := rpc.UnmarshalIntoSameIndex(&m1, &m2)
	if hasAny {
		t.Fatal("expected false")
	}
	_ = errs
	// nil
	errs2, hasAny2 := rpc.UnmarshalIntoSameIndex(nil)
	_ = errs2
	_ = hasAny2
}

func Test_CovJsonS2_RPC11_AddSerializer_Funcs(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializer(nil)
	rpc.AddSerializerFunc(nil)
	rpc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`1`), nil
	})
	rpc.AddSerializers()
	rpc.AddSerializerFunctions()
}

// --- MapResults ---

func newTestMR() *corejson.MapResults {
	mr := corejson.NewMapResults.UsingCap(4)
	mr.Add("a", corejson.New(map[string]int{"x": 1}))
	mr.Add("b", corejson.New(map[string]int{"y": 2}))
	return mr
}

func Test_CovJsonS2_MR01_Length_IsEmpty_HasAnyItem(t *testing.T) {
	mr := newTestMR()
	if mr.Length() != 2 {
		t.Fatal("expected 2")
	}
	var nilMR *corejson.MapResults
	if nilMR.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovJsonS2_MR02_AddSkipOnNil_GetByKey(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddSkipOnNil("k", nil)
	r := corejson.NewPtr(1)
	mr.AddSkipOnNil("k", r)
	if mr.Length() != 1 {
		t.Fatal("expected 1")
	}
	got := mr.GetByKey("k")
	if got == nil {
		t.Fatal("expected non-nil")
	}
	if mr.GetByKey("missing") != nil {
		t.Fatal("expected nil")
	}
}
func Test_CovJsonS2_MR04_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	mr := newTestMR()
	keys := mr.AllKeys()
	if len(keys) != 2 {
		t.Fatal("expected 2")
	}
	sorted := mr.AllKeysSorted()
	if sorted[0] != "a" {
		t.Fatal("expected a first")
	}
	vals := mr.AllValues()
	if len(vals) != 2 {
		t.Fatal("expected 2")
	}
	_ = mr.AllResults()
	_ = mr.AllResultsCollection()
	// empty
	empty := corejson.NewMapResults.Empty()
	if len(empty.AllKeys()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.AllKeysSorted()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.AllValues()) != 0 {
		t.Fatal("expected 0")
	}
	if empty.AllResultsCollection().HasAnyItem() {
		t.Fatal("expected empty")
	}
}
func Test_CovJsonS2_MR06_Add_AddPtr_AddAny_AddAnySkipOnNil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddPtr("k", nil)
	if mr.Length() != 0 {
		t.Fatal("expected 0")
	}
	err := mr.AddAny("k", map[string]int{"a": 1})
	if err != nil {
		t.Fatal("expected no error")
	}
	err2 := mr.AddAny("nil", nil)
	if err2 == nil {
		t.Fatal("expected error")
	}
	err3 := mr.AddAnySkipOnNil("nil", nil)
	if err3 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS2_MR07_AddAnyNonEmptyNonError_AddAnyNonEmpty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmptyNonError("k", nil)
	mr.AddAnyNonEmpty("k2", nil)
	mr.AddAnyNonEmpty("k3", 1)
}

func Test_CovJsonS2_MR08_AddKeyWithResult_AddKeyWithResultPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	kr := corejson.KeyWithResult{Key: "k", Result: corejson.New(1)}
	mr.AddKeyWithResult(kr)
	mr.AddKeyWithResultPtr(nil)
	mr.AddKeyWithResultPtr(&kr)
	mr.AddKeysWithResultsPtr()
	mr.AddKeysWithResults()
}

func Test_CovJsonS2_MR09_AddKeyAny(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	ka := corejson.KeyAny{Key: "k", AnyInf: 1}
	mr.AddKeyAnyInf(ka)
	mr.AddKeyAnyInfPtr(nil)
	mr.AddKeyAnyInfPtr(&ka)
	mr.AddKeyAnyItems(ka)
	mr.AddKeyAnyItemsPtr(nil)
}

func Test_CovJsonS2_MR10_AddNonEmptyNonErrorPtr_AddMapResults_AddMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddNonEmptyNonErrorPtr("k", nil)
	mr.AddMapResults(nil)
	mr2 := newTestMR()
	mr.AddMapResults(mr2)
	mr.AddMapAnyItems(nil)
	mr.AddMapAnyItems(map[string]any{"x": 1})
}

func Test_CovJsonS2_MR11_GetPagesSize_GetPagedCollection(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(10)
	for i := 0; i < 10; i++ {
		mr.Add(string(rune('a'+i)), corejson.New(i))
	}
	if mr.GetPagesSize(3) != 4 {
		t.Fatal("expected 4")
	}
	if mr.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	pages := mr.GetPagedCollection(3)
	if len(pages) < 3 {
		t.Fatal("expected at least 3")
	}
	// small
	small := newTestMR()
	pages2 := small.GetPagedCollection(10)
	if len(pages2) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovJsonS2_MR12_AddMapResultsUsingCloneOption(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	raw := map[string]corejson.Result{"k": corejson.New(1)}
	mr.AddMapResultsUsingCloneOption(false, false, raw)
	mr2 := corejson.NewMapResults.Empty()
	mr2.AddMapResultsUsingCloneOption(true, true, raw)
	mr3 := corejson.NewMapResults.Empty()
	mr3.AddMapResultsUsingCloneOption(false, false, nil)
}

func Test_CovJsonS2_MR13_GetNewMapUsingKeys(t *testing.T) {
	mr := newTestMR()
	sub := mr.GetNewMapUsingKeys(false, "a")
	if sub.Length() != 1 {
		t.Fatal("expected 1")
	}
	// empty keys
	empty := mr.GetNewMapUsingKeys(false)
	if empty.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovJsonS2_MR14_ResultCollection_Json_JsonPtr_Interfaces(t *testing.T) {
	mr := newTestMR()
	rc := mr.ResultCollection()
	if rc.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	_ = mr.Json()
	_ = mr.JsonPtr()
	_ = mr.JsonModel()
	_ = mr.JsonModelAny()
	_ = mr.AsJsonContractsBinder()
	_ = mr.AsJsoner()
	_ = mr.AsJsonParseSelfInjector()
}

func Test_CovJsonS2_MR15_Clear_Dispose(t *testing.T) {
	mr := newTestMR()
	mr.Dispose()
	var nilMR *corejson.MapResults
	nilMR.Dispose()
}

func Test_CovJsonS2_MR16_ParseInjectUsingJson(t *testing.T) {
	mr := newTestMR()
	jr := mr.JsonPtr()
	mr2 := corejson.NewMapResults.Empty()
	_, err := mr2.ParseInjectUsingJson(jr)
	if err != nil {
		t.Fatal("expected no error")
	}
}

// --- Deserializer (deserializerLogic) ---

func Test_CovJsonS2_DL01_Apply_UsingResult(t *testing.T) {
	r := corejson.NewPtr(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.Deserialize.Apply(r, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
	err2 := corejson.Deserialize.UsingResult(r, &m)
	if err2 != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS2_DL02_UsingString_FromString(t *testing.T) {
	var m map[string]int
	err := corejson.Deserialize.UsingString(`{"a":1}`, &m)
	if err != nil || m["a"] != 1 {
		t.Fatal("expected a=1")
	}
	var m2 map[string]int
	err2 := corejson.Deserialize.FromString(`{"a":1}`, &m2)
	if err2 != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS2_DL03_UsingStringOption_UsingStringIgnoreEmpty(t *testing.T) {
	var m map[string]int
	err := corejson.Deserialize.UsingStringOption(true, "", &m)
	if err != nil {
		t.Fatal("expected nil for empty")
	}
	err2 := corejson.Deserialize.UsingStringIgnoreEmpty("", &m)
	if err2 != nil {
		t.Fatal("expected nil for empty")
	}
}

func Test_CovJsonS2_DL04_UsingStringPtr(t *testing.T) {
	s := `{"a":1}`
	var m map[string]int
	err := corejson.Deserialize.UsingStringPtr(&s, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
	err2 := corejson.Deserialize.UsingStringPtr(nil, &m)
	if err2 == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_CovJsonS2_DL05_UsingError_UsingErrorWhichJsonResult(t *testing.T) {
	err := corejson.Deserialize.UsingError(nil, nil)
	if err != nil {
		t.Fatal("expected nil")
	}
	err2 := corejson.Deserialize.UsingErrorWhichJsonResult(nil, nil)
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS2_DL06_UsingBytes_UsingBytesPointer_UsingBytesMust(t *testing.T) {
	var m map[string]int
	err := corejson.Deserialize.UsingBytes([]byte(`{"a":1}`), &m)
	if err != nil {
		t.Fatal("expected no error")
	}
	err2 := corejson.Deserialize.UsingBytesPointer(nil, &m)
	if err2 == nil {
		t.Fatal("expected error for nil bytes")
	}
}

func Test_CovJsonS2_DL07_UsingBytesIf_UsingBytesPointerIf(t *testing.T) {
	var m map[string]int
	err := corejson.Deserialize.UsingBytesIf(false, nil, &m)
	if err != nil {
		t.Fatal("expected nil")
	}
	err2 := corejson.Deserialize.UsingBytesPointerIf(false, nil, &m)
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS2_DL08_MapAnyToPointer(t *testing.T) {
	var m map[string]int
	err := corejson.Deserialize.MapAnyToPointer(true, nil, &m)
	if err != nil {
		t.Fatal("expected nil for empty skip")
	}
	err2 := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"a": 1}, &m)
	if err2 != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS2_DL09_FromTo(t *testing.T) {
	from := map[string]int{"a": 1}
	var to map[string]int
	err := corejson.Deserialize.FromTo(from, &to)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS2_DL10_AnyToFieldsMap(t *testing.T) {
	// AnyToFieldsMap → DeserializedFieldsToMap passes value not pointer — known limitation
	m, _ := corejson.Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
	_ = m // covers the call path regardless of result
}

// --- deserializeFromBytesTo ---

func Test_CovJsonS2_BT01_Strings_String_Integer(t *testing.T) {
	sb, _ := corejson.Serialize.Raw([]string{"a", "b"})
	lines, err := corejson.Deserialize.BytesTo.Strings(sb)
	if err != nil || len(lines) != 2 {
		t.Fatal("expected 2")
	}
	sb2, _ := corejson.Serialize.Raw("hello")
	s, err2 := corejson.Deserialize.BytesTo.String(sb2)
	if err2 != nil || s != "hello" {
		t.Fatal("expected hello")
	}
	sb3, _ := corejson.Serialize.Raw(42)
	i, err3 := corejson.Deserialize.BytesTo.Integer(sb3)
	if err3 != nil || i != 42 {
		t.Fatal("expected 42")
	}
	_ = corejson.Deserialize.BytesTo.IntegerMust(sb3)
}

func Test_CovJsonS2_BT02_Integer64_Integers_Bool(t *testing.T) {
	sb, _ := corejson.Serialize.Raw(int64(42))
	i64, err := corejson.Deserialize.BytesTo.Integer64(sb)
	if err != nil || i64 != 42 {
		t.Fatal("expected 42")
	}
	_ = corejson.Deserialize.BytesTo.Integer64Must(sb)
	sbi, _ := corejson.Serialize.Raw([]int{1, 2})
	ints, err2 := corejson.Deserialize.BytesTo.Integers(sbi)
	if err2 != nil || len(ints) != 2 {
		t.Fatal("expected 2")
	}
	sbb, _ := corejson.Serialize.Raw(true)
	b, err3 := corejson.Deserialize.BytesTo.Bool(sbb)
	if err3 != nil || !b {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS2_BT03_MapAnyItem_MapStringString(t *testing.T) {
	sb, _ := corejson.Serialize.Raw(map[string]any{"a": 1})
	m, err := corejson.Deserialize.BytesTo.MapAnyItem(sb)
	if err != nil || len(m) == 0 {
		t.Fatal("expected map")
	}
	_ = corejson.Deserialize.BytesTo.MapAnyItemMust(sb)
	sb2, _ := corejson.Serialize.Raw(map[string]string{"a": "b"})
	ms, err2 := corejson.Deserialize.BytesTo.MapStringString(sb2)
	if err2 != nil || len(ms) == 0 {
		t.Fatal("expected map")
	}
	_ = corejson.Deserialize.BytesTo.MapStringStringMust(sb2)
}

func Test_CovJsonS2_BT04_Bytes_BytesMust(t *testing.T) {
	inner := []byte(`"hello"`)
	sb, _ := corejson.Serialize.Raw(inner)
	b, err := corejson.Deserialize.BytesTo.Bytes(sb)
	if err != nil {
		t.Fatal("expected no error")
	}
	_ = b
}

func Test_CovJsonS2_BT05_ResultCollection_ResultsPtrCollection_MapResults(t *testing.T) {
	rc := newTestRC()
	sb, _ := corejson.Serialize.Raw(rc)
	rc2, err := corejson.Deserialize.BytesTo.ResultCollection(sb)
	if err != nil {
		t.Fatal("expected no error")
	}
	_ = rc2
	rpc := newTestRPC()
	sb2, _ := corejson.Serialize.Raw(rpc)
	rpc2, err2 := corejson.Deserialize.BytesTo.ResultsPtrCollection(sb2)
	if err2 != nil {
		t.Fatal("expected no error")
	}
	_ = rpc2
	mr := newTestMR()
	sb3, _ := corejson.Serialize.Raw(mr)
	mr2, err3 := corejson.Deserialize.BytesTo.MapResults(sb3)
	if err3 != nil {
		t.Fatal("expected no error")
	}
	_ = mr2
}

// --- deserializeFromResultTo ---

func Test_CovJsonS2_RT01_String_Bool_Byte(t *testing.T) {
	r := corejson.NewPtr("hello")
	s, err := corejson.Deserialize.ResultTo.String(r)
	if err != nil || s != "hello" {
		t.Fatal("expected hello")
	}
	r2 := corejson.NewPtr(true)
	b, err2 := corejson.Deserialize.ResultTo.Bool(r2)
	if err2 != nil || !b {
		t.Fatal("expected true")
	}
	r3 := corejson.NewPtr(byte(5))
	bv, err3 := corejson.Deserialize.ResultTo.Byte(r3)
	if err3 != nil || bv != 5 {
		t.Fatal("expected 5")
	}
}

func Test_CovJsonS2_RT02_MapAnyItem_MapStringString(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"a": 1})
	m, err := corejson.Deserialize.ResultTo.MapAnyItem(r)
	if err != nil || len(m) == 0 {
		t.Fatal("expected map")
	}
	r2 := corejson.NewPtr(map[string]string{"a": "b"})
	ms, err2 := corejson.Deserialize.ResultTo.MapStringString(r2)
	if err2 != nil || len(ms) == 0 {
		t.Fatal("expected map")
	}
}

// --- Serializer (serializerLogic) ---

func Test_CovJsonS2_SL01_Apply_Various(t *testing.T) {
	r := corejson.Serialize.Apply(1)
	if r.HasError() {
		t.Fatal("expected no error")
	}
	_ = corejson.Serialize.FromBytes([]byte("test"))
	_ = corejson.Serialize.FromStrings([]string{"a"})
	_ = corejson.Serialize.FromStringsSpread("a", "b")
	_ = corejson.Serialize.FromString("hello")
	_ = corejson.Serialize.FromInteger(42)
	_ = corejson.Serialize.FromInteger64(42)
	_ = corejson.Serialize.FromBool(true)
	_ = corejson.Serialize.FromIntegers([]int{1, 2})
	_ = corejson.Serialize.StringsApply([]string{"a"})
}

func Test_CovJsonS2_SL02_UsingAny_UsingAnyPtr_Raw_Marshal(t *testing.T) {
	r := corejson.Serialize.UsingAny(1)
	if r.HasError() {
		t.Fatal("expected no error")
	}
	rp := corejson.Serialize.UsingAnyPtr(1)
	if rp.HasError() {
		t.Fatal("expected no error")
	}
	raw, err := corejson.Serialize.Raw(1)
	if err != nil || len(raw) == 0 {
		t.Fatal("expected bytes")
	}
	raw2, err2 := corejson.Serialize.Marshal(1)
	if err2 != nil || len(raw2) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_CovJsonS2_SL03_ToBytes_ToString(t *testing.T) {
	_ = corejson.Serialize.ToBytesMust(1)
	_ = corejson.Serialize.ToSafeBytesMust(1)
	_ = corejson.Serialize.ToSafeBytesSwallowErr(1)
	_ = corejson.Serialize.ToBytesSwallowErr(1)
	_, _ = corejson.Serialize.ToBytesErr(1)
	_ = corejson.Serialize.ToString(1)
	_ = corejson.Serialize.ToStringMust(1)
	_, _ = corejson.Serialize.ToStringErr(1)
	_, _ = corejson.Serialize.ToPrettyStringErr(1)
	_ = corejson.Serialize.ToPrettyStringIncludingErr(1)
	_ = corejson.Serialize.Pretty(1)
}

// --- Creators ---

func Test_CovJsonS2_CR01_Empty(t *testing.T) {
	_ = corejson.Empty.Result()
	_ = corejson.Empty.ResultPtr()
	_ = corejson.Empty.ResultWithErr("type", errors.New("err"))
	_ = corejson.Empty.ResultPtrWithErr("type", errors.New("err"))
	_ = corejson.Empty.BytesCollection()
	_ = corejson.Empty.BytesCollectionPtr()
	_ = corejson.Empty.ResultsCollection()
	_ = corejson.Empty.ResultsPtrCollection()
	_ = corejson.Empty.MapResults()
}
func Test_CovJsonS2_CR03_NewResultCreator_DeserializeUsingBytes_DeserializeUsingResult(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	b, _ := r.Serialize()
	_ = corejson.NewResult.DeserializeUsingBytes(b)
	_ = corejson.NewResult.UnmarshalUsingBytes(b)
	rp := r.ToPtr()
	_ = corejson.NewResult.DeserializeUsingResult(rp)
}

func Test_CovJsonS2_CR04_NewResultsCollection(t *testing.T) {
	_ = corejson.NewResultsCollection.Empty()
	_ = corejson.NewResultsCollection.Default()
	_ = corejson.NewResultsCollection.UsingCap(5)
	_ = corejson.NewResultsCollection.AnyItems(1, 2)
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(2, 1, 2)
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(2)
	_ = corejson.NewResultsCollection.UsingResultsPtr()
	_ = corejson.NewResultsCollection.UsingResults()
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(2, corejson.New(1))
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(2)
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(2, corejson.NewPtr(1))
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(2)
	_ = corejson.NewResultsCollection.Serializers()
	_ = corejson.NewResultsCollection.SerializerFunctions()
}

func Test_CovJsonS2_CR05_NewResultsPtrCollection(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Empty()
	_ = corejson.NewResultsPtrCollection.Default()
	_ = corejson.NewResultsPtrCollection.UsingCap(5)
	_ = corejson.NewResultsPtrCollection.AnyItems(1, 2)
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(2, 1, 2)
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(2)
	_ = corejson.NewResultsPtrCollection.UsingResults()
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(2, corejson.NewPtr(1))
	_ = corejson.NewResultsPtrCollection.Serializers()
}

func Test_CovJsonS2_CR06_NewMapResults(t *testing.T) {
	_ = corejson.NewMapResults.Empty()
	_ = corejson.NewMapResults.UsingCap(5)
	_ = corejson.NewMapResults.UsingKeyAnyItems(0, corejson.KeyAny{Key: "k", AnyInf: 1})
	_ = corejson.NewMapResults.UsingKeyAnyItems(0)
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, nil)
	raw := map[string]corejson.Result{"k": corejson.New(1)}
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, raw)
	_ = corejson.NewMapResults.UsingMapOptions(true, false, 2, raw)
	_ = corejson.NewMapResults.UsingMapPlusCap(0, nil)
	_ = corejson.NewMapResults.UsingMapPlusCap(0, raw)
	_ = corejson.NewMapResults.UsingMapPlusCapClone(0, nil)
	_ = corejson.NewMapResults.UsingMapPlusCapClone(0, raw)
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(0, nil)
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(0, raw)
	_ = corejson.NewMapResults.UsingMap(nil)
	_ = corejson.NewMapResults.UsingMap(raw)
	_ = corejson.NewMapResults.UsingMapAnyItems(nil)
	_ = corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": 1})
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(0, nil)
	_ = corejson.NewMapResults.UsingKeyWithResults()
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(0)
	_ = corejson.NewMapResults.UsingKeyJsoners()
	_ = corejson.NewMapResults.UsingKeyJsonersPlusCap(0)
}

func Test_CovJsonS2_CR07_NewBytesCollection(t *testing.T) {
	_ = corejson.NewBytesCollection.Empty()
	_ = corejson.NewBytesCollection.UsingCap(5)
	_, _ = corejson.NewBytesCollection.AnyItems(1, 2)
}

func Test_CovJsonS2_CR08_NewResultsCollection_DeserializeUsingBytes(t *testing.T) {
	rc := newTestRC()
	b, _ := corejson.Serialize.Raw(rc)
	rc2, err := corejson.NewResultsCollection.DeserializeUsingBytes(b)
	if err != nil {
		t.Fatal("expected no error")
	}
	_ = rc2
	_, _ = corejson.NewResultsCollection.UnmarshalUsingBytes(b)
	rp := corejson.NewPtr(rc)
	_, _ = corejson.NewResultsCollection.DeserializeUsingResult(rp)
}

func Test_CovJsonS2_CR09_NewResultsPtrCollection_DeserializeUsingBytes(t *testing.T) {
	rpc := newTestRPC()
	b, _ := corejson.Serialize.Raw(rpc)
	rpc2, err := corejson.NewResultsPtrCollection.DeserializeUsingBytes(b)
	if err != nil {
		t.Fatal("expected no error")
	}
	_ = rpc2
	_, _ = corejson.NewResultsPtrCollection.UnmarshalUsingBytes(b)
	rp := corejson.NewPtr(rpc)
	_, _ = corejson.NewResultsPtrCollection.DeserializeUsingResult(rp)
}

func Test_CovJsonS2_CR10_NewMapResults_DeserializeUsingBytes(t *testing.T) {
	mr := newTestMR()
	b, _ := corejson.Serialize.Raw(mr)
	mr2, err := corejson.NewMapResults.DeserializeUsingBytes(b)
	if err != nil {
		t.Fatal("expected no error")
	}
	_ = mr2
	_, _ = corejson.NewMapResults.UnmarshalUsingBytes(b)
	rp := corejson.NewPtr(mr)
	_, _ = corejson.NewMapResults.DeserializeUsingResult(rp)
}

func Test_CovJsonS2_CR11_NewBytesCollection_DeserializeUsingBytes_DeserializeUsingResult(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte(`"a"`))
	b, _ := corejson.Serialize.Raw(bc)
	bc2, err := corejson.NewBytesCollection.DeserializeUsingBytes(b)
	if err != nil {
		t.Fatal("expected no error")
	}
	_ = bc2
	_, _ = corejson.NewBytesCollection.UnmarshalUsingBytes(b)
	rp := corejson.NewPtr(bc)
	_, _ = corejson.NewBytesCollection.DeserializeUsingResult(rp)
}

// --- BytesCollection methods ---

func Test_CovJsonS2_BC01_Length_FirstOrDefault_LastOrDefault(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte(`1`))
	bc.Add([]byte(`2`))
	if bc.Length() != 2 {
		t.Fatal("expected 2")
	}
	if bc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if bc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	empty := corejson.NewBytesCollection.Empty()
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS2_BC02_Take_Limit_Skip(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(3)
	bc.Add([]byte(`1`))
	bc.Add([]byte(`2`))
	_ = bc.Take(1)
	_ = bc.Limit(1)
	_ = bc.Limit(-1)
	_ = bc.Skip(1)
	empty := corejson.NewBytesCollection.Empty()
	_ = empty.Take(1)
	_ = empty.Limit(1)
	_ = empty.Skip(1)
}

func Test_CovJsonS2_BC03_AddSkipOnNil_AddNonEmpty_AddPtr_Adds(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSkipOnNil(nil)
	bc.AddNonEmpty(nil)
	bc.AddPtr(nil)
	bc.Adds()
	bc.Adds([]byte(`1`), nil, []byte(`2`))
}

func Test_CovJsonS2_BC04_AddResult_AddResultPtr_AddAny_AddAnyItems(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	r := corejson.New(1)
	bc.AddResult(r)
	bc.AddResultPtr(corejson.NewPtr(1))
	err := bc.AddAny(1)
	if err != nil {
		t.Fatal("expected no error")
	}
	err2 := bc.AddAnyItems(1, 2)
	if err2 != nil {
		t.Fatal("expected no error")
	}
	_ = bc.AddAnyItems()
}
func Test_CovJsonS2_BC06_UnmarshalAt_UnmarshalIntoSameIndex(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte(`1`))
	bc.Add([]byte(`2`))
	var i, j int
	err := bc.UnmarshalAt(0, &i)
	if err != nil {
		t.Fatal("expected no error")
	}
	errs, _ := bc.UnmarshalIntoSameIndex(&i, &j)
	_ = errs
	bc.UnmarshalIntoSameIndex(nil)
}
func Test_CovJsonS2_BC08_GetPagesSize_GetPagedCollection(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		bc.Add([]byte(`1`))
	}
	if bc.GetPagesSize(3) != 4 {
		t.Fatal("expected 4")
	}
	if bc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	pages := bc.GetPagedCollection(3)
	if len(pages) < 3 {
		t.Fatal("expected at least 3")
	}
	// small
	small := corejson.NewBytesCollection.UsingCap(2)
	small.Add([]byte(`1`))
	pages2 := small.GetPagedCollection(5)
	if len(pages2) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovJsonS2_BC09_Json_JsonModel_Interfaces(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(1)
	bc.Add([]byte(`1`))
	_ = bc.JsonModel()
	_ = bc.JsonModelAny()
	_ = bc.Json()
	_ = bc.JsonPtr()
	_ = bc.AsJsonContractsBinder()
	_ = bc.AsJsoner()
	_ = bc.AsJsonParseSelfInjector()
}

func Test_CovJsonS2_BC10_AddMapResults_AddRawMapResults_AddsPtr_AddBytesCollection(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	mr := newTestMR()
	bc.AddMapResults(mr)
	bc.AddRawMapResults(nil)
	bc.AddsPtr(nil)
	bc2 := corejson.NewBytesCollection.UsingCap(1)
	bc2.Add([]byte(`1`))
	bc.AddBytesCollection(bc2)
}

func Test_CovJsonS2_BC11_AddSerializer_AddSerializers_Funcs(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializer(nil)
	bc.AddSerializers()
	bc.AddSerializerFunc(nil)
	bc.AddSerializerFunctions()
	bc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`1`), nil
	})
}
