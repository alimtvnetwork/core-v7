package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 4: Result deep methods, ResultsCollection,
//                     ResultsPtrCollection, MapResults — all remaining branches
// ══════════════════════════════════════════════════════════════════════════════

// --- Result deep methods ---

func Test_CovJsonS4_R01_Map_WithFields(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	m := r.Map()
	if len(m) == 0 {
		t.Fatal("expected non-empty map")
	}
}

func Test_CovJsonS4_R02_Map_NilResult(t *testing.T) {
	var r *corejson.Result
	m := r.Map()
	if len(m) != 0 {
		t.Fatal("expected empty map")
	}
}

func Test_CovJsonS4_R03_Map_WithError(t *testing.T) {
	r := corejson.Result{Error: errors.New("fail"), TypeName: "T"}
	m := r.Map()
	if m["Error"] != "fail" {
		t.Fatal("expected error in map")
	}
}

func Test_CovJsonS4_R04_DeserializedFieldsToMap(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	_, err := r.DeserializedFieldsToMap()
	_ = err
}

func Test_CovJsonS4_R05_SafeDeserializedFieldsToMap(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	_ = r.SafeDeserializedFieldsToMap()
}

func Test_CovJsonS4_R06_FieldsNames(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	names, _ := r.FieldsNames()
	_ = names
}

func Test_CovJsonS4_R07_SafeFieldsNames(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	_ = r.SafeFieldsNames()
}

func Test_CovJsonS4_R08_BytesTypeName(t *testing.T) {
	r := corejson.New(1)
	if r.BytesTypeName() == "" {
		t.Fatal("expected type name")
	}
	var nr *corejson.Result
	if nr.BytesTypeName() != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovJsonS4_R09_SafeBytesTypeName(t *testing.T) {
	r := corejson.New(1)
	_ = r.SafeBytesTypeName()
	var nr *corejson.Result
	_ = nr.SafeBytesTypeName()
}

func Test_CovJsonS4_R10_SafeString(t *testing.T) {
	r := corejson.New(1)
	if r.SafeString() == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS4_R11_JsonStringPtr_Cached(t *testing.T) {
	r := corejson.New(1)
	_ = r.JsonStringPtr()
	_ = r.JsonStringPtr() // second call should use cached
}

func Test_CovJsonS4_R12_PrettyJsonBuffer(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	buf, err := r.PrettyJsonBuffer("", "  ")
	if err != nil || buf.Len() == 0 {
		t.Fatal("expected buffer")
	}
	// empty
	var nr *corejson.Result
	buf2, _ := nr.PrettyJsonBuffer("", "  ")
	_ = buf2
}

func Test_CovJsonS4_R13_PrettyJsonStringOrErrString(t *testing.T) {
	r := corejson.New(1)
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected string")
	}
	// nil
	var nr *corejson.Result
	s2 := nr.PrettyJsonStringOrErrString()
	if s2 == "" {
		t.Fatal("expected nil message")
	}
	// with error
	re := corejson.Result{Error: errors.New("fail")}
	s3 := re.PrettyJsonStringOrErrString()
	if s3 == "" {
		t.Fatal("expected error message")
	}
}

func Test_CovJsonS4_R14_ErrorString(t *testing.T) {
	r := corejson.New(1)
	if r.ErrorString() != "" {
		t.Fatal("expected empty")
	}
	re := corejson.Result{Error: errors.New("fail")}
	if re.ErrorString() == "" {
		t.Fatal("expected error string")
	}
}

func Test_CovJsonS4_R15_IsErrorEqual(t *testing.T) {
	r := corejson.New(1)
	if !r.IsErrorEqual(nil) {
		t.Fatal("expected true")
	}
	re := corejson.Result{Error: errors.New("fail")}
	if !re.IsErrorEqual(errors.New("fail")) {
		t.Fatal("expected true")
	}
	if re.IsErrorEqual(nil) {
		t.Fatal("expected false")
	}
	if re.IsErrorEqual(errors.New("other")) {
		t.Fatal("expected false")
	}
}

func Test_CovJsonS4_R16_String_WithError(t *testing.T) {
	re := corejson.Result{Error: errors.New("fail"), TypeName: "T", Bytes: []byte("x")}
	s := re.String()
	if s == "" {
		t.Fatal("expected string")
	}
	r := corejson.New(1)
	s2 := r.String()
	if s2 == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS4_R17_SafeNonIssueBytes(t *testing.T) {
	r := corejson.New(1)
	if len(r.SafeNonIssueBytes()) == 0 {
		t.Fatal("expected bytes")
	}
	re := corejson.Result{Error: errors.New("fail")}
	if len(re.SafeNonIssueBytes()) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_CovJsonS4_R18_SafeBytes_Values_SafeValues(t *testing.T) {
	r := corejson.New(1)
	_ = r.SafeBytes()
	_ = r.Values()
	_ = r.SafeValues()
	_ = r.SafeValuesPtr()
}

func Test_CovJsonS4_R19_Raw_RawMust_RawString_RawStringMust_RawErrString_RawPrettyString(t *testing.T) {
	r := corejson.New(1)
	_, _ = r.Raw()
	_ = r.RawMust()
	_, _ = r.RawString()
	_ = r.RawStringMust()
	_, _ = r.RawErrString()
	_, _ = r.RawPrettyString()
}

func Test_CovJsonS4_R20_MeaningfulError_NilBytes(t *testing.T) {
	r := corejson.Result{TypeName: "T"}
	e := r.MeaningfulError()
	if e == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJsonS4_R21_MeaningfulError_WithError(t *testing.T) {
	r := corejson.Result{Error: errors.New("fail"), Bytes: []byte("x"), TypeName: "T"}
	e := r.MeaningfulError()
	if e == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJsonS4_R22_IsEmpty_HasAnyItem(t *testing.T) {
	r := corejson.New(1)
	if r.IsEmpty() {
		t.Fatal("expected false")
	}
	if !r.HasAnyItem() {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS4_R23_IsEmptyJson_HasJson_HasBytes_HasJsonBytes(t *testing.T) {
	r := corejson.New(1)
	_ = r.IsEmptyJson()
	_ = r.HasJson()
	_ = r.HasBytes()
	_ = r.HasJsonBytes()
}

func Test_CovJsonS4_R24_HasSafeItems(t *testing.T) {
	r := corejson.New(1)
	if !r.HasSafeItems() {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS4_R25_IsAnyNull_HasIssuesOrEmpty(t *testing.T) {
	r := corejson.New(1)
	_ = r.IsAnyNull()
	_ = r.HasIssuesOrEmpty()
}

func Test_CovJsonS4_R26_MeaningfulErrorMessage(t *testing.T) {
	r := corejson.New(1)
	if r.MeaningfulErrorMessage() != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovJsonS4_R27_InjectInto(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	jr := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.UsingCap(1)
	_ = jr.InjectInto(rc2)
}

func Test_CovJsonS4_R28_Deserialize_DeserializeMust_UnmarshalMust(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	var m map[string]int
	_ = r.Deserialize(&m)
	r2 := corejson.New(map[string]int{"b": 2})
	var m2 map[string]int
	r2.DeserializeMust(&m2)
	r3 := corejson.New(map[string]int{"c": 3})
	var m3 map[string]int
	r3.UnmarshalMust(&m3)
}

func Test_CovJsonS4_R29_Unmarshal_NilResult(t *testing.T) {
	var r *corejson.Result
	var m map[string]int
	err := r.Unmarshal(&m)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJsonS4_R30_Unmarshal_WithExistingError(t *testing.T) {
	re := corejson.Result{Error: errors.New("fail"), Bytes: []byte("x"), TypeName: "T"}
	var m map[string]int
	err := re.Unmarshal(&m)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJsonS4_R31_SerializeSkipExistingIssues(t *testing.T) {
	r := corejson.New(1)
	_, _ = r.SerializeSkipExistingIssues()
	// empty
	re := corejson.Result{Error: errors.New("fail")}
	b, e := re.SerializeSkipExistingIssues()
	if b != nil || e != nil {
		t.Fatal("expected nil,nil")
	}
}

func Test_CovJsonS4_R32_Serialize_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Serialize()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJsonS4_R33_Serialize_WithError(t *testing.T) {
	re := corejson.Result{Error: errors.New("fail")}
	_, err := re.Serialize()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJsonS4_R34_Serialize_Success(t *testing.T) {
	r := corejson.New(1)
	b, err := r.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_CovJsonS4_R35_SerializeMust(t *testing.T) {
	r := corejson.New(1)
	_ = r.SerializeMust()
}

func Test_CovJsonS4_R36_UnmarshalSkipExistingIssues(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	var m map[string]int
	_ = r.UnmarshalSkipExistingIssues(&m)
	// empty
	re := corejson.Result{Error: errors.New("fail")}
	_ = re.UnmarshalSkipExistingIssues(&m)
}

func Test_CovJsonS4_R37_UnmarshalResult(t *testing.T) {
	r := corejson.New(1)
	jr := r.JsonPtr()
	serialized, _ := jr.Serialize()
	r2 := corejson.Result{Bytes: serialized, TypeName: "Result"}
	_, _ = r2.UnmarshalResult()
}

func Test_CovJsonS4_R38_JsonModel_JsonModelAny(t *testing.T) {
	r := corejson.New(1)
	_ = r.JsonModel()
	_ = r.JsonModelAny()
	var nr *corejson.Result
	_ = nr.JsonModel()
	_ = nr.JsonModelAny()
}

func Test_CovJsonS4_R39_Json_JsonPtr(t *testing.T) {
	r := corejson.New(1)
	_ = r.Json()
	_ = r.JsonPtr()
}

func Test_CovJsonS4_R40_ParseInjectUsingJson(t *testing.T) {
	r := corejson.New(1)
	jr := r.JsonPtr()
	b, _ := jr.Serialize()
	r2 := corejson.Result{Bytes: b, TypeName: "Result"}
	empty := corejson.Empty.ResultPtr()
	_, _ = empty.ParseInjectUsingJson(&r2)
}

func Test_CovJsonS4_R41_ParseInjectUsingJsonMust(t *testing.T) {
	r := corejson.New(1)
	jr := r.JsonPtr()
	b, _ := jr.Serialize()
	r2 := corejson.Result{Bytes: b, TypeName: "Result"}
	empty := corejson.Empty.ResultPtr()
	_ = empty.ParseInjectUsingJsonMust(&r2)
}

func Test_CovJsonS4_R42_CloneError(t *testing.T) {
	r := corejson.New(1)
	if r.CloneError() != nil {
		t.Fatal("expected nil")
	}
	re := corejson.Result{Error: errors.New("fail")}
	if re.CloneError() == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJsonS4_R43_Ptr_NonPtr_ToPtr_ToNonPtr(t *testing.T) {
	r := corejson.New(1)
	_ = r.Ptr()
	_ = r.NonPtr()
	_ = r.ToPtr()
	_ = r.ToNonPtr()
	var nr *corejson.Result
	_ = nr.NonPtr()
}

func Test_CovJsonS4_R44_IsEqualPtr(t *testing.T) {
	r := corejson.New(1)
	r2 := corejson.New(1)
	if !r.IsEqualPtr(&r2) {
		t.Fatal("expected true")
	}
	var nr *corejson.Result
	if !nr.IsEqualPtr(nil) {
		t.Fatal("expected true")
	}
	if nr.IsEqualPtr(&r) {
		t.Fatal("expected false")
	}
	// same pointer
	rp := r.Ptr()
	if !rp.IsEqualPtr(rp) {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS4_R45_IsEqual(t *testing.T) {
	r := corejson.New(1)
	r2 := corejson.New(1)
	if !r.IsEqual(r2) {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS4_R46_CombineErrorWithRefString_CombineErrorWithRefError(t *testing.T) {
	r := corejson.New(1)
	if r.CombineErrorWithRefString("ref") != "" {
		t.Fatal("expected empty")
	}
	if r.CombineErrorWithRefError("ref") != nil {
		t.Fatal("expected nil")
	}
	re := corejson.Result{Error: errors.New("fail")}
	if re.CombineErrorWithRefString("ref") == "" {
		t.Fatal("expected string")
	}
	if re.CombineErrorWithRefError("ref") == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJsonS4_R47_BytesError(t *testing.T) {
	r := corejson.New(1)
	be := r.BytesError()
	if be == nil {
		t.Fatal("expected non-nil")
	}
	var nr *corejson.Result
	if nr.BytesError() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS4_R48_Dispose(t *testing.T) {
	r := corejson.New(1)
	r.Dispose()
	var nr *corejson.Result
	nr.Dispose()
}

func Test_CovJsonS4_R49_CloneIf_Clone_ClonePtr(t *testing.T) {
	r := corejson.New(1)
	_ = r.CloneIf(true, true)
	_ = r.CloneIf(false, false)
	_ = r.CloneIf(true, false)
	_ = r.Clone(true)
	_ = r.Clone(false)
	_ = r.ClonePtr(true)
	_ = r.ClonePtr(false)
	var nr *corejson.Result
	if nr.ClonePtr(true) != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS4_R50_AsJsonContractsBinder_AsJsoner_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.New(1)
	_ = r.AsJsonContractsBinder()
	_ = r.AsJsoner()
	_ = r.AsJsonParseSelfInjector()
}

func Test_CovJsonS4_R51_JsonParseSelfInject(t *testing.T) {
	r := corejson.New(1)
	jr := r.JsonPtr()
	b, _ := jr.Serialize()
	r2 := corejson.Result{Bytes: b, TypeName: "Result"}
	_ = r.JsonParseSelfInject(&r2)
}

// --- ResultsCollection remaining branches ---

func Test_CovJsonS4_RC01_Basic(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(5)
	rc.AddAny(1)
	rc.AddAny(2)
	if rc.Length() != 2 {
		t.Fatal("expected 2")
	}
	if rc.LastIndex() != 1 {
		t.Fatal("expected 1")
	}
	if rc.IsEmpty() {
		t.Fatal("expected false")
	}
	if !rc.HasAnyItem() {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS4_RC02_FirstOrDefault_LastOrDefault(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)
	if rc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if rc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	empty := corejson.NewResultsCollection.UsingCap(0)
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS4_RC03_Take_Limit_Skip(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		rc.AddAny(i)
	}
	_ = rc.Take(3)
	_ = rc.Limit(3)
	_ = rc.Limit(-1)
	_ = rc.Skip(2)
	// empty
	empty := corejson.NewResultsCollection.UsingCap(0)
	_ = empty.Take(1)
	_ = empty.Limit(1)
	_ = empty.Skip(1)
}

func Test_CovJsonS4_RC04_AddMethods(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(5)
	r := corejson.New(1)
	rc.AddSkipOnNil(nil)
	rc.AddSkipOnNil(&r)
	rc.AddNonNilNonError(nil)
	rc.AddNonNilNonError(&r)
	rc.Add(r)
	rc.Adds(r)
	rc.AddPtr(&r)
	rc.AddPtr(nil)
	rc.AddsPtr(&r)
	rc.AddsPtr(nil)
	rc.AddAny(1)
	rc.AddAny(nil)
	rc.AddAnyItems(1, nil, 2)
	rc.AddAnyItemsSlice([]any{1, nil, 2})
	rc.AddNonNilItemsPtr(&r, nil)
	rc.AddNonNilItemsPtr()
	rc.AddResultsCollection(nil)
	rc2 := corejson.NewResultsCollection.UsingCap(1)
	rc2.AddAny(1)
	rc.AddResultsCollection(rc2)
}

func Test_CovJsonS4_RC05_GetAt_GetAtSafe_GetAtSafeUsingLength(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)
	_ = rc.GetAt(0)
	_ = rc.GetAtSafe(0)
	_ = rc.GetAtSafe(-1)
	_ = rc.GetAtSafe(10)
	_ = rc.GetAtSafeUsingLength(0, 1)
	_ = rc.GetAtSafeUsingLength(10, 1)
}

func Test_CovJsonS4_RC06_HasError_AllErrors_GetErrorsStrings(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)
	if rc.HasError() {
		t.Fatal("expected false")
	}
	_, has := rc.AllErrors()
	if has {
		t.Fatal("expected false")
	}
	_ = rc.GetErrorsStrings()
	_ = rc.GetErrorsStringsPtr()
	_ = rc.GetErrorsAsSingleString()
	_ = rc.GetErrorsAsSingle()
	// empty
	empty := corejson.NewResultsCollection.UsingCap(0)
	_ = empty.GetErrorsStrings()
	_, _ = empty.AllErrors()
}

func Test_CovJsonS4_RC07_UnmarshalAt_InjectIntoAt(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(map[string]int{"a": 1})
	var m map[string]int
	_ = rc.UnmarshalAt(0, &m)
	rc2 := corejson.NewResultsCollection.UsingCap(1)
	_ = rc.InjectIntoAt(0, rc2)
}

func Test_CovJsonS4_RC08_InjectIntoSameIndex(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)
	_, _ = rc.InjectIntoSameIndex(nil)
}

func Test_CovJsonS4_RC09_UnmarshalIntoSameIndex(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(map[string]int{"a": 1})
	var m map[string]int
	_, _ = rc.UnmarshalIntoSameIndex(&m)
	_, _ = rc.UnmarshalIntoSameIndex(nil)
}

func Test_CovJsonS4_RC10_AddSerializers(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddSerializerFunc(nil)
	rc.AddSerializerFunc(func() ([]byte, error) { return []byte("1"), nil })
	rc.AddSerializerFunctions()
	rc.AddSerializerFunctions(func() ([]byte, error) { return []byte("1"), nil })
	rc.AddSerializer(nil)
	rc.AddSerializers()
}

func Test_CovJsonS4_RC11_AddMapResults_AddRawMapResults(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	mr := corejson.NewMapResults.UsingCap(1)
	mr.AddAny("k", 1)
	rc.AddMapResults(mr)
	rc.AddRawMapResults(nil)
	rc.AddRawMapResults(mr.Items)
}

func Test_CovJsonS4_RC12_AddJsoners(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc2 := corejson.NewResultsCollection.UsingCap(1)
	rc2.AddAny(1)
	rc.AddJsoners(false, rc2)
	rc.AddJsoners(true, nil)
}

func Test_CovJsonS4_RC13_GetStrings_GetStringsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)
	_ = rc.GetStrings()
	_ = rc.GetStringsPtr()
	empty := corejson.NewResultsCollection.UsingCap(0)
	_ = empty.GetStrings()
}

func Test_CovJsonS4_RC14_NonPtr_Ptr(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(1)
	_ = rc.NonPtr()
	_ = rc.Ptr()
}

func Test_CovJsonS4_RC15_Clear_Dispose(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)
	rc.Clear()
	rc2 := corejson.NewResultsCollection.UsingCap(2)
	rc2.AddAny(1)
	rc2.Dispose()
	var nilRC *corejson.ResultsCollection
	nilRC.Clear()
	nilRC.Dispose()
}

func Test_CovJsonS4_RC16_Paging(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		rc.AddAny(i)
	}
	if rc.GetPagesSize(3) != 4 {
		t.Fatal("expected 4")
	}
	if rc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	pages := rc.GetPagedCollection(3)
	if len(pages) < 3 {
		t.Fatal("expected at least 3")
	}
	_ = rc.GetSinglePageCollection(3, 2)
	// small
	small := corejson.NewResultsCollection.UsingCap(2)
	small.AddAny(1)
	_ = small.GetPagedCollection(10)
	_ = small.GetSinglePageCollection(10, 1)
}

func Test_CovJsonS4_RC17_Json_JsonPtr_JsonModel_JsonModelAny(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	_ = rc.Json()
	_ = rc.JsonPtr()
	_ = rc.JsonModel()
	_ = rc.JsonModelAny()
}

func Test_CovJsonS4_RC18_ParseInjectUsingJson_ParseInjectUsingJsonMust(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	jr := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.UsingCap(0)
	_, _ = rc2.ParseInjectUsingJson(jr)
	rc3 := corejson.NewResultsCollection.UsingCap(0)
	_ = rc3.ParseInjectUsingJsonMust(jr)
}

func Test_CovJsonS4_RC19_AsInterfaces(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(1)
	_ = rc.AsJsonContractsBinder()
	_ = rc.AsJsoner()
	_ = rc.AsJsonParseSelfInjector()
	_ = rc.JsonParseSelfInject(rc.JsonPtr())
}

func Test_CovJsonS4_RC20_ShadowClone_Clone_ClonePtr(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)
	_ = rc.ShadowClone()
	_ = rc.Clone(true)
	_ = rc.ClonePtr(true)
	_ = rc.ClonePtr(false)
	var nilRC *corejson.ResultsCollection
	if nilRC.ClonePtr(true) != nil {
		t.Fatal("expected nil")
	}
}

// --- ResultsPtrCollection remaining branches ---

func Test_CovJsonS4_RPC01_Basic(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(5)
	rpc.AddAny(1)
	rpc.AddAny(2)
	if rpc.Length() != 2 {
		t.Fatal("expected 2")
	}
	if rpc.LastIndex() != 1 {
		t.Fatal("expected 1")
	}
	if rpc.IsEmpty() {
		t.Fatal("expected false")
	}
	if !rpc.HasAnyItem() {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS4_RPC02_FirstOrDefault_LastOrDefault(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)
	if rpc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if rpc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	empty := corejson.NewResultsPtrCollection.UsingCap(0)
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS4_RPC03_Take_Limit_Skip(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		rpc.AddAny(i)
	}
	_ = rpc.Take(3)
	_ = rpc.Limit(3)
	_ = rpc.Limit(-1)
	_ = rpc.Skip(2)
	empty := corejson.NewResultsPtrCollection.UsingCap(0)
	_ = empty.Take(1)
	_ = empty.Limit(1)
	_ = empty.Skip(1)
}

func Test_CovJsonS4_RPC04_AddMethods(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(5)
	r := corejson.NewPtr(1)
	rpc.AddSkipOnNil(nil)
	rpc.AddSkipOnNil(r)
	rpc.AddNonNilNonError(nil)
	rpc.AddNonNilNonError(r)
	rpc.Add(r)
	rpc.Adds(r)
	rpc.AddAny(1)
	rpc.AddAny(nil)
	rpc.AddAnyItems(1, nil, 2)
	rpc.AddResult(*r)
	rpc.AddNonNilItems(r, nil)
	rpc.AddNonNilItemsPtr(r, nil)
	rpc.AddNonNilItemsPtr()
	rpc.AddResultsCollection(nil)
	rpc2 := corejson.NewResultsPtrCollection.UsingCap(1)
	rpc2.AddAny(1)
	rpc.AddResultsCollection(rpc2)
}

func Test_CovJsonS4_RPC05_GetAt_GetAtSafe_GetAtSafeUsingLength(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)
	_ = rpc.GetAt(0)
	_ = rpc.GetAtSafe(0)
	_ = rpc.GetAtSafe(-1)
	_ = rpc.GetAtSafe(10)
	_ = rpc.GetAtSafeUsingLength(0, 1)
	_ = rpc.GetAtSafeUsingLength(10, 1)
}

func Test_CovJsonS4_RPC06_HasError_AllErrors_GetErrorsStrings(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)
	if rpc.HasError() {
		t.Fatal("expected false")
	}
	_, has := rpc.AllErrors()
	if has {
		t.Fatal("expected false")
	}
	_ = rpc.GetErrorsStrings()
	_ = rpc.GetErrorsStringsPtr()
	_ = rpc.GetErrorsAsSingleString()
	_ = rpc.GetErrorsAsSingle()
	empty := corejson.NewResultsPtrCollection.UsingCap(0)
	_ = empty.GetErrorsStrings()
	_, _ = empty.AllErrors()
}

func Test_CovJsonS4_RPC07_UnmarshalAt_InjectIntoAt(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(map[string]int{"a": 1})
	var m map[string]int
	_ = rpc.UnmarshalAt(0, &m)
	rc := corejson.NewResultsCollection.UsingCap(1)
	_ = rpc.InjectIntoAt(0, rc)
}

func Test_CovJsonS4_RPC08_InjectIntoSameIndex(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)
	_, _ = rpc.InjectIntoSameIndex(nil)
}

func Test_CovJsonS4_RPC09_UnmarshalIntoSameIndex(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(map[string]int{"a": 1})
	var m map[string]int
	_, _ = rpc.UnmarshalIntoSameIndex(&m)
	_, _ = rpc.UnmarshalIntoSameIndex(nil)
}

func Test_CovJsonS4_RPC10_AddSerializers(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddSerializerFunc(nil)
	rpc.AddSerializerFunc(func() ([]byte, error) { return []byte("1"), nil })
	rpc.AddSerializerFunctions()
	rpc.AddSerializerFunctions(func() ([]byte, error) { return []byte("1"), nil })
	rpc.AddSerializer(nil)
	rpc.AddSerializers()
}

func Test_CovJsonS4_RPC11_AddJsoners(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	rpc.AddJsoners(false, rc)
	rpc.AddJsoners(true, nil)
}

func Test_CovJsonS4_RPC12_GetStrings_GetStringsPtr(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)
	_ = rpc.GetStrings()
	_ = rpc.GetStringsPtr()
	empty := corejson.NewResultsPtrCollection.UsingCap(0)
	_ = empty.GetStrings()
}

func Test_CovJsonS4_RPC13_NonPtr_Ptr(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(1)
	_ = rpc.NonPtr()
	_ = rpc.Ptr()
}

func Test_CovJsonS4_RPC14_Clear_Dispose(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)
	rpc.Clear()
	rpc2 := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc2.AddAny(1)
	rpc2.Dispose()
	var nilRPC *corejson.ResultsPtrCollection
	nilRPC.Clear()
	nilRPC.Dispose()
}

func Test_CovJsonS4_RPC15_Paging(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		rpc.AddAny(i)
	}
	if rpc.GetPagesSize(3) != 4 {
		t.Fatal("expected 4")
	}
	if rpc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	pages := rpc.GetPagedCollection(3)
	if len(pages) < 3 {
		t.Fatal("expected at least 3")
	}
	_ = rpc.GetSinglePageCollection(3, 2)
	small := corejson.NewResultsPtrCollection.UsingCap(2)
	small.AddAny(1)
	_ = small.GetPagedCollection(10)
	_ = small.GetSinglePageCollection(10, 1)
}

func Test_CovJsonS4_RPC16_Json_JsonPtr_JsonModel_JsonModelAny(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(1)
	rpc.AddAny(1)
	_ = rpc.Json()
	_ = rpc.JsonPtr()
	_ = rpc.JsonModel()
	_ = rpc.JsonModelAny()
}

func Test_CovJsonS4_RPC17_ParseInject_AsInterfaces(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(1)
	rpc.AddAny(1)
	jr := rpc.JsonPtr()
	rpc2 := corejson.NewResultsPtrCollection.UsingCap(0)
	_, _ = rpc2.ParseInjectUsingJson(jr)
	rpc3 := corejson.NewResultsPtrCollection.UsingCap(0)
	_ = rpc3.ParseInjectUsingJsonMust(jr)
	_ = rpc.AsJsonContractsBinder()
	_ = rpc.AsJsoner()
	_ = rpc.AsJsonParseSelfInjector()
	_ = rpc.JsonParseSelfInject(jr)
}

func Test_CovJsonS4_RPC18_Clone(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)
	_ = rpc.Clone(true)
	_ = rpc.Clone(false)
	var nilRPC *corejson.ResultsPtrCollection
	if nilRPC.Clone(true) != nil {
		t.Fatal("expected nil")
	}
}

// --- MapResults remaining branches ---

func Test_CovJsonS4_MR01_Basic(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(5)
	mr.AddAny("a", 1)
	mr.AddAny("b", 2)
	if mr.Length() != 2 {
		t.Fatal("expected 2")
	}
	if mr.LastIndex() != 1 {
		t.Fatal("expected 1")
	}
	if mr.IsEmpty() {
		t.Fatal("expected false")
	}
	if !mr.HasAnyItem() {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS4_MR02_AddMethods(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(5)
	r := corejson.New(1)
	mr.Add("k", r)
	mr.AddPtr("k2", &r)
	mr.AddPtr("k3", nil)
	mr.AddSkipOnNil("k4", nil)
	mr.AddSkipOnNil("k5", &r)
	_ = mr.AddAny("k6", 1)
	_ = mr.AddAny("nil", nil)
	_ = mr.AddAnySkipOnNil("k7", 1)
	_ = mr.AddAnySkipOnNil("k8", nil)
	mr.AddAnyNonEmptyNonError("k9", 1)
	mr.AddAnyNonEmptyNonError("k10", nil)
	mr.AddAnyNonEmpty("k11", 1)
	mr.AddAnyNonEmpty("k12", nil)
	mr.AddNonEmptyNonErrorPtr("k13", &r)
	mr.AddNonEmptyNonErrorPtr("k14", nil)
}

func Test_CovJsonS4_MR03_AddKeyWith(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(5)
	r := corejson.New(1)
	mr.AddKeyWithResult(corejson.KeyWithResult{Key: "k", Result: r})
	mr.AddKeyWithResultPtr(nil)
	mr.AddKeyWithResultPtr(&corejson.KeyWithResult{Key: "k2", Result: r})
	mr.AddKeysWithResultsPtr()
	mr.AddKeysWithResultsPtr(&corejson.KeyWithResult{Key: "k3", Result: r})
	mr.AddKeysWithResults(corejson.KeyWithResult{Key: "k4", Result: r})
	mr.AddKeysWithResults()
	mr.AddKeyAnyInf(corejson.KeyAny{Key: "k5", AnyInf: 1})
	mr.AddKeyAnyInfPtr(nil)
	mr.AddKeyAnyInfPtr(&corejson.KeyAny{Key: "k6", AnyInf: 1})
	mr.AddKeyAnyItems(corejson.KeyAny{Key: "k7", AnyInf: 1})
	mr.AddKeyAnyItems()
	mr.AddKeyAnyItemsPtr(&corejson.KeyAny{Key: "k8", AnyInf: 1})
	mr.AddKeyAnyItemsPtr()
}

func Test_CovJsonS4_MR04_AddMapResults_AddMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(5)
	mr2 := corejson.NewMapResults.UsingCap(1)
	mr2.AddAny("k", 1)
	mr.AddMapResults(mr2)
	mr.AddMapResults(nil)
	mr.AddMapResults(corejson.NewMapResults.UsingCap(0))
	mr.AddMapAnyItems(map[string]any{"k2": 2})
	mr.AddMapAnyItems(nil)
}

func Test_CovJsonS4_MR05_GetByKey(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("k", 1)
	if mr.GetByKey("k") == nil {
		t.Fatal("expected non-nil")
	}
	if mr.GetByKey("missing") != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS4_MR06_HasError_AllErrors_GetErrorsStrings(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("k", 1)
	if mr.HasError() {
		t.Fatal("expected false")
	}
	_, has := mr.AllErrors()
	if has {
		t.Fatal("expected false")
	}
	_ = mr.GetErrorsStrings()
	_ = mr.GetErrorsStringsPtr()
	_ = mr.GetErrorsAsSingleString()
	_ = mr.GetErrorsAsSingle()
	empty := corejson.NewMapResults.UsingCap(0)
	_ = empty.GetErrorsStrings()
	_, _ = empty.AllErrors()
}

func Test_CovJsonS4_MR07_AllKeys_AllKeysSorted_AllValues_AllResults_AllResultsCollection(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("b", 2)
	mr.AddAny("a", 1)
	keys := mr.AllKeys()
	if len(keys) != 2 {
		t.Fatal("expected 2")
	}
	sorted := mr.AllKeysSorted()
	if sorted[0] != "a" {
		t.Fatal("expected a first")
	}
	_ = mr.AllValues()
	_ = mr.AllResults()
	_ = mr.AllResultsCollection()
	_ = mr.ResultCollection()
	empty := corejson.NewMapResults.UsingCap(0)
	_ = empty.AllKeys()
	_ = empty.AllKeysSorted()
	_ = empty.AllValues()
	_ = empty.AllResultsCollection()
	_ = empty.ResultCollection()
}

func Test_CovJsonS4_MR08_GetStrings_GetStringsPtr(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("k", 1)
	_ = mr.GetStrings()
	_ = mr.GetStringsPtr()
	empty := corejson.NewMapResults.UsingCap(0)
	_ = empty.GetStrings()
}

func Test_CovJsonS4_MR09_AddJsoner_AddKeyWithJsoner_AddKeysWithJsoners(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(5)
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	mr.AddJsoner("k", rc)
	mr.AddJsoner("k2", nil)
	mr.AddKeyWithJsoner(corejson.KeyWithJsoner{Key: "k3", Jsoner: rc})
	mr.AddKeysWithJsoners(corejson.KeyWithJsoner{Key: "k4", Jsoner: rc})
	mr.AddKeysWithJsoners()
	mr.AddKeyWithJsonerPtr(nil)
	mr.AddKeyWithJsonerPtr(&corejson.KeyWithJsoner{Key: "k5", Jsoner: rc})
}

func Test_CovJsonS4_MR10_Unmarshal_SafeUnmarshal_Deserialize(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("k", map[string]int{"a": 1})
	var m map[string]int
	_ = mr.Unmarshal("k", &m)
	_ = mr.Deserialize("k", &m)
	_ = mr.SafeUnmarshal("k", &m)
	_ = mr.SafeDeserialize("k", &m)
}

func Test_CovJsonS4_MR11_UnmarshalMany_UnmarshalManySafe(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("k", map[string]int{"a": 1})
	var m map[string]int
	_ = mr.UnmarshalMany(corejson.KeyAny{Key: "k", AnyInf: &m})
	_ = mr.UnmarshalMany()
	_ = mr.UnmarshalManySafe(corejson.KeyAny{Key: "k", AnyInf: &m})
	_ = mr.UnmarshalManySafe()
}

func Test_CovJsonS4_MR12_InjectIntoAt(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(2)
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	mr.Add("k", rc.Json())
	rc2 := corejson.NewResultsCollection.UsingCap(0)
	_ = mr.InjectIntoAt("k", rc2)
}

func Test_CovJsonS4_MR13_DeserializeMust_SafeDeserializeMust(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("k", map[string]int{"a": 1})
	var m map[string]int
	mr.DeserializeMust("k", &m)
	mr.SafeDeserializeMust("k", &m)
}

func Test_CovJsonS4_MR14_Paging(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(10)
	for i := 0; i < 10; i++ {
		mr.AddAny("k"+string(rune('a'+i)), i)
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
	keys := mr.AllKeysSorted()
	_ = mr.GetSinglePageCollection(3, 2, keys)
	// small
	small := corejson.NewMapResults.UsingCap(2)
	small.AddAny("k", 1)
	_ = small.GetPagedCollection(10)
	_ = small.GetSinglePageCollection(10, 1, small.AllKeysSorted())
}

func Test_CovJsonS4_MR15_GetNewMapUsingKeys(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(3)
	mr.AddAny("a", 1)
	mr.AddAny("b", 2)
	sub := mr.GetNewMapUsingKeys(false, "a")
	if sub.Length() != 1 {
		t.Fatal("expected 1")
	}
	_ = mr.GetNewMapUsingKeys(false)
	_ = mr.GetNewMapUsingKeys(false, "missing")
}

func Test_CovJsonS4_MR16_AddMapResultsUsingCloneOption(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(3)
	mr2 := corejson.NewMapResults.UsingCap(1)
	r := corejson.New(1)
	mr2.Add("k", r)
	mr.AddMapResultsUsingCloneOption(false, false, mr2.Items)
	mr.AddMapResultsUsingCloneOption(true, true, mr2.Items)
	mr.AddMapResultsUsingCloneOption(false, false, nil)
}

func Test_CovJsonS4_MR17_Clear_Dispose(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("k", 1)
	mr.Clear()
	mr2 := corejson.NewMapResults.UsingCap(2)
	mr2.AddAny("k", 1)
	mr2.Dispose()
	var nilMR *corejson.MapResults
	nilMR.Clear()
	nilMR.Dispose()
}

func Test_CovJsonS4_MR18_Json_JsonPtr_JsonModel_JsonModelAny(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(1)
	mr.AddAny("k", 1)
	_ = mr.Json()
	_ = mr.JsonPtr()
	_ = mr.JsonModel()
	_ = mr.JsonModelAny()
}

func Test_CovJsonS4_MR19_ParseInject_AsInterfaces(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(1)
	mr.AddAny("k", 1)
	jr := mr.JsonPtr()
	mr2 := corejson.NewMapResults.UsingCap(0)
	_, _ = mr2.ParseInjectUsingJson(jr)
	mr3 := corejson.NewMapResults.UsingCap(0)
	_ = mr3.ParseInjectUsingJsonMust(jr)
	_ = mr.AsJsonContractsBinder()
	_ = mr.AsJsoner()
	_ = mr.AsJsonParseSelfInjector()
	_ = mr.JsonParseSelfInject(jr)
}
