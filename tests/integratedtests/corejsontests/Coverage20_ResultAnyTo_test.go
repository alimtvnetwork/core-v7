package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 1: Result methods, anyTo, BytesCloneIf, etc.
// ══════════════════════════════════════════════════════════════════════════════

// --- BytesCloneIf / BytesDeepClone / BytesToString ---

func Test_CovJson_S1_01_BytesCloneIf(t *testing.T) {
	b := []byte(`{"a":1}`)
	r := corejson.BytesCloneIf(true, b)
	if len(r) != len(b) {
		t.Fatal("expected same len")
	}
	// not deep clone
	r2 := corejson.BytesCloneIf(false, b)
	if len(r2) != 0 {
		t.Fatal("expected empty")
	}
	// empty input
	r3 := corejson.BytesCloneIf(true, []byte{})
	if len(r3) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_CovJson_S1_02_BytesDeepClone(t *testing.T) {
	b := []byte(`{"a":1}`)
	r := corejson.BytesDeepClone(b)
	if len(r) != len(b) {
		t.Fatal("expected same len")
	}
	r2 := corejson.BytesDeepClone(nil)
	if len(r2) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_CovJson_S1_03_BytesToString(t *testing.T) {
	r := corejson.BytesToString([]byte(`{"a":1}`))
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := corejson.BytesToString(nil)
	if r2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovJson_S1_04_BytesToPrettyString(t *testing.T) {
	r := corejson.BytesToPrettyString([]byte(`{"a":1}`))
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := corejson.BytesToPrettyString(nil)
	if r2 != "" {
		t.Fatal("expected empty")
	}
}

// --- New / NewPtr ---

func Test_CovJson_S1_05_New(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	if r.HasError() {
		t.Fatal("expected no error")
	}
	if r.IsEmpty() {
		t.Fatal("expected not empty")
	}
}

func Test_CovJson_S1_06_NewPtr(t *testing.T) {
	r := corejson.NewPtr(map[string]int{"a": 1})
	if r == nil || r.HasError() {
		t.Fatal("expected no error")
	}
}

// --- JsonString / JsonStringOrErrMsg ---

func Test_CovJson_S1_07_JsonString(t *testing.T) {
	s, err := corejson.JsonString(map[string]int{"a": 1})
	if err != nil || s == "" {
		t.Fatal("expected json string")
	}
}

func Test_CovJson_S1_08_JsonStringOrErrMsg(t *testing.T) {
	s := corejson.JsonStringOrErrMsg(map[string]int{"a": 1})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// --- Result methods ---

func newTestResult() *corejson.Result {
	r := corejson.New(map[string]int{"a": 1, "b": 2})
	return r.ToPtr()
}

func Test_CovJson_S1_09_Result_Map(t *testing.T) {
	r := newTestResult()
	m := r.Map()
	if len(m) == 0 {
		t.Fatal("expected non-empty")
	}
	// nil
	var nilR *corejson.Result
	m2 := nilR.Map()
	if len(m2) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_CovJson_S1_10_Result_BytesTypeName_SafeBytesTypeName(t *testing.T) {
	r := newTestResult()
	_ = r.BytesTypeName()
	_ = r.SafeBytesTypeName()
	// nil
	var nilR *corejson.Result
	if nilR.BytesTypeName() != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovJson_S1_11_Result_SafeString_JsonString_JsonStringPtr(t *testing.T) {
	r := newTestResult()
	s := r.SafeString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	js := r.JsonString()
	if js == "" {
		t.Fatal("expected non-empty")
	}
	jsp := r.JsonStringPtr()
	if jsp == nil || *jsp == "" {
		t.Fatal("expected non-empty")
	}
	// nil
	var nilR *corejson.Result
	nsp := nilR.JsonStringPtr()
	if nsp == nil {
		t.Fatal("expected non-nil ptr")
	}
}

func Test_CovJson_S1_12_Result_PrettyJsonString(t *testing.T) {
	r := newTestResult()
	pj := r.PrettyJsonString()
	if pj == "" {
		t.Fatal("expected non-empty")
	}
	// nil
	var nilR *corejson.Result
	if nilR.PrettyJsonString() != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovJson_S1_13_Result_PrettyJsonStringOrErrString(t *testing.T) {
	r := newTestResult()
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	// nil
	var nilR *corejson.Result
	s2 := nilR.PrettyJsonStringOrErrString()
	if s2 == "" {
		t.Fatal("expected non-empty error string")
	}
	// with error
	errR := &corejson.Result{Error: errors.New("test")}
	s3 := errR.PrettyJsonStringOrErrString()
	if s3 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovJson_S1_14_Result_Length(t *testing.T) {
	r := newTestResult()
	if r.Length() == 0 {
		t.Fatal("expected non-zero")
	}
	// nil
	var nilR *corejson.Result
	if nilR.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovJson_S1_15_Result_HasError_ErrorString_IsEmptyError(t *testing.T) {
	r := newTestResult()
	if r.HasError() {
		t.Fatal("expected no error")
	}
	es := r.ErrorString()
	if es != "" {
		t.Fatal("expected empty")
	}
	if !r.IsEmptyError() {
		t.Fatal("expected true")
	}
}

func Test_CovJson_S1_16_Result_IsErrorEqual(t *testing.T) {
	r := newTestResult()
	if !r.IsErrorEqual(nil) {
		t.Fatal("expected true")
	}
	errR := &corejson.Result{Error: errors.New("test")}
	if !errR.IsErrorEqual(errors.New("test")) {
		t.Fatal("expected true")
	}
	if errR.IsErrorEqual(nil) {
		t.Fatal("expected false")
	}
	if errR.IsErrorEqual(errors.New("other")) {
		t.Fatal("expected false")
	}
}

func Test_CovJson_S1_17_Result_String(t *testing.T) {
	r := newTestResult()
	s := r.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}
func Test_CovJson_S1_19_Result_Raw_RawMust(t *testing.T) {
	r := newTestResult()
	b, err := r.Raw()
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
	_ = r.RawMust()
}

func Test_CovJson_S1_20_Result_Raw_Nil(t *testing.T) {
	var nilR *corejson.Result
	_, err := nilR.Raw()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJson_S1_21_Result_RawString_RawStringMust_RawErrString_RawPrettyString(t *testing.T) {
	r := newTestResult()
	s, err := r.RawString()
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
	sm := r.RawStringMust()
	if sm == "" {
		t.Fatal("expected non-empty")
	}
	rb, re := r.RawErrString()
	if len(rb) == 0 || re != "" {
		t.Fatal("expected bytes")
	}
	ps, perr := r.RawPrettyString()
	if perr != nil || ps == "" {
		t.Fatal("expected pretty string")
	}
}

func Test_CovJson_S1_22_Result_MeaningfulError_MeaningfulErrorMessage(t *testing.T) {
	r := newTestResult()
	if r.MeaningfulError() != nil {
		t.Fatal("expected nil")
	}
	if r.MeaningfulErrorMessage() != "" {
		t.Fatal("expected empty")
	}
	// nil result
	var nilR *corejson.Result
	if nilR.MeaningfulError() == nil {
		t.Fatal("expected error")
	}
	// empty bytes
	emptyR := &corejson.Result{}
	if emptyR.MeaningfulError() == nil {
		t.Fatal("expected error for empty bytes")
	}
	// error + bytes
	errR := &corejson.Result{Bytes: []byte(`{"a":1}`), Error: errors.New("some err")}
	if errR.MeaningfulError() == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJson_S1_23_Result_HasIssuesOrEmpty_HasSafeItems_IsAnyNull(t *testing.T) {
	r := newTestResult()
	if r.HasIssuesOrEmpty() {
		t.Fatal("expected false")
	}
	if !r.HasSafeItems() {
		t.Fatal("expected true")
	}
	if r.IsAnyNull() {
		t.Fatal("expected false")
	}
}

func Test_CovJson_S1_24_Result_HasBytes_HasJsonBytes_HasJson_IsEmptyJson_IsEmptyJsonBytes(t *testing.T) {
	r := newTestResult()
	if !r.HasBytes() {
		t.Fatal("expected true")
	}
	if !r.HasJsonBytes() {
		t.Fatal("expected true")
	}
	if !r.HasJson() {
		t.Fatal("expected true")
	}
	if r.IsEmptyJson() {
		t.Fatal("expected false")
	}
	if r.IsEmptyJsonBytes() {
		t.Fatal("expected false")
	}
	// empty json "{}"
	emptyJson := &corejson.Result{Bytes: []byte("{}")}
	if !emptyJson.IsEmptyJsonBytes() {
		t.Fatal("expected true for {}")
	}
}

func Test_CovJson_S1_25_Result_Serialize_SerializeMust(t *testing.T) {
	r := newTestResult()
	b, err := r.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("expected serialization")
	}
	_ = r.SerializeMust()
	// nil
	var nilR *corejson.Result
	_, err2 := nilR.Serialize()
	if err2 == nil {
		t.Fatal("expected error")
	}
	// with error
	errR := &corejson.Result{Error: errors.New("test")}
	_, err3 := errR.Serialize()
	if err3 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJson_S1_26_Result_SerializeSkipExistingIssues(t *testing.T) {
	r := newTestResult()
	b, err := r.SerializeSkipExistingIssues()
	if err != nil || len(b) == 0 {
		t.Fatal("expected serialization")
	}
	// with issues
	errR := &corejson.Result{Error: errors.New("test")}
	b2, err2 := errR.SerializeSkipExistingIssues()
	if err2 != nil || b2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJson_S1_27_Result_Unmarshal_Deserialize(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	rp := r.ToPtr()
	var m map[string]int
	err := rp.Deserialize(&m)
	if err != nil || m["a"] != 1 {
		t.Fatal("expected deserialization")
	}
	// nil result
	var nilR *corejson.Result
	err2 := nilR.Unmarshal(&m)
	if err2 == nil {
		t.Fatal("expected error")
	}
	// result with error
	errR := &corejson.Result{Error: errors.New("test")}
	err3 := errR.Unmarshal(&m)
	if err3 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJson_S1_28_Result_UnmarshalSkipExistingIssues(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1}).ToPtr()
	var m map[string]int
	err := r.UnmarshalSkipExistingIssues(&m)
	if err != nil {
		t.Fatal("expected no error")
	}
	// with issues
	errR := &corejson.Result{Error: errors.New("test")}
	err2 := errR.UnmarshalSkipExistingIssues(&m)
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJson_S1_29_Result_UnmarshalResult(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1}).ToPtr()
	_, _ = r.UnmarshalResult()
}

func Test_CovJson_S1_30_Result_JsonModel_JsonModelAny(t *testing.T) {
	r := newTestResult()
	jm := r.JsonModel()
	if jm.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	jma := r.JsonModelAny()
	if jma == nil {
		t.Fatal("expected non-nil")
	}
	// nil
	var nilR *corejson.Result
	njm := nilR.JsonModel()
	if njm.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJson_S1_31_Result_Json_JsonPtr(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	j := r.Json()
	if j.HasError() {
		t.Fatal("expected no error")
	}
	jp := r.JsonPtr()
	if jp == nil || jp.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CovJson_S1_32_Result_CloneError(t *testing.T) {
	r := newTestResult()
	if r.CloneError() != nil {
		t.Fatal("expected nil")
	}
	errR := &corejson.Result{Error: errors.New("test")}
	if errR.CloneError() == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJson_S1_33_Result_Ptr_NonPtr_ToPtr_ToNonPtr(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	p := r.Ptr()
	if p == nil {
		t.Fatal("expected non-nil")
	}
	np := p.NonPtr()
	if np.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	tp := r.ToPtr()
	if tp == nil {
		t.Fatal("expected non-nil")
	}
	tnp := r.ToNonPtr()
	if tnp.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	// nil NonPtr
	var nilR *corejson.Result
	nnp := nilR.NonPtr()
	if nnp.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJson_S1_34_Result_IsEqualPtr(t *testing.T) {
	r1 := newTestResult()
	r2 := newTestResult()
	if !r1.IsEqualPtr(r2) {
		t.Fatal("expected equal")
	}
	// nil both
	var nilR *corejson.Result
	if !nilR.IsEqualPtr(nil) {
		t.Fatal("expected nil==nil")
	}
	if nilR.IsEqualPtr(r1) {
		t.Fatal("expected false")
	}
	// different
	r3 := corejson.New(map[string]int{"c": 3}).ToPtr()
	if r1.IsEqualPtr(r3) {
		t.Fatal("expected false")
	}
}

func Test_CovJson_S1_35_Result_IsEqual(t *testing.T) {
	r1 := corejson.New(map[string]int{"a": 1})
	r2 := corejson.New(map[string]int{"a": 1})
	if !r1.IsEqual(r2) {
		t.Fatal("expected equal")
	}
}

func Test_CovJson_S1_36_Result_CombineErrorWithRefString_Error(t *testing.T) {
	r := newTestResult()
	s := r.CombineErrorWithRefString("ref")
	if s != "" {
		t.Fatal("expected empty")
	}
	e := r.CombineErrorWithRefError("ref")
	if e != nil {
		t.Fatal("expected nil")
	}
	errR := &corejson.Result{Error: errors.New("test")}
	s2 := errR.CombineErrorWithRefString("ref")
	if s2 == "" {
		t.Fatal("expected non-empty")
	}
	e2 := errR.CombineErrorWithRefError("ref")
	if e2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJson_S1_37_Result_BytesError(t *testing.T) {
	r := newTestResult()
	be := r.BytesError()
	if be == nil {
		t.Fatal("expected non-nil")
	}
	// nil
	var nilR *corejson.Result
	if nilR.BytesError() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJson_S1_38_Result_Dispose(t *testing.T) {
	r := newTestResult()
	r.Dispose()
	if r.Length() != 0 {
		t.Fatal("expected 0")
	}
	// nil dispose
	var nilR *corejson.Result
	nilR.Dispose()
}
func Test_CovJson_S1_40_Result_AsJsonContractsBinder_AsJsoner_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	_ = r.AsJsonContractsBinder()
	_ = r.AsJsoner()
	_ = r.AsJsonParseSelfInjector()
}

func Test_CovJson_S1_41_Result_PrettyJsonBuffer(t *testing.T) {
	r := newTestResult()
	buf, err := r.PrettyJsonBuffer("", "  ")
	if err != nil || buf.Len() == 0 {
		t.Fatal("expected buffer")
	}
	// empty
	emptyR := &corejson.Result{}
	buf2, _ := emptyR.PrettyJsonBuffer("", "  ")
	if buf2.Len() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_CovJson_S1_42_Result_HasAnyItem_IsEmpty(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	if !r.HasAnyItem() {
		t.Fatal("expected true")
	}
	if r.IsEmpty() {
		t.Fatal("expected false")
	}
}

// --- AnyTo ---

func Test_CovJson_S1_43_AnyTo_SerializedJsonResult(t *testing.T) {
	// from map
	r := corejson.AnyTo.SerializedJsonResult(map[string]int{"a": 1})
	if r.HasError() {
		t.Fatal("expected no error")
	}
	// from string
	r2 := corejson.AnyTo.SerializedJsonResult(`{"a":1}`)
	if r2.HasError() {
		t.Fatal("expected no error")
	}
	// from bytes
	r3 := corejson.AnyTo.SerializedJsonResult([]byte(`{"a":1}`))
	if r3.HasError() {
		t.Fatal("expected no error")
	}
	// from Result
	r4 := corejson.AnyTo.SerializedJsonResult(corejson.New(1))
	if r4.HasError() {
		t.Fatal("expected no error")
	}
	// from *Result
	rp := corejson.New(1).ToPtr()
	r5 := corejson.AnyTo.SerializedJsonResult(rp)
	if r5.HasError() {
		t.Fatal("expected no error")
	}
	// from nil
	r6 := corejson.AnyTo.SerializedJsonResult(nil)
	if r6.Error == nil {
		t.Fatal("expected error")
	}
	// from error
	r7 := corejson.AnyTo.SerializedJsonResult(errors.New("test"))
	_ = r7
}

func Test_CovJson_S1_44_AnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw(map[string]int{"a": 1})
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_CovJson_S1_45_AnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString(map[string]int{"a": 1})
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJson_S1_46_AnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString(map[string]int{"a": 1})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovJson_S1_47_AnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString(1)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovJson_S1_48_AnyTo_SafeJsonPrettyString(t *testing.T) {
	// string passthrough
	s := corejson.AnyTo.SafeJsonPrettyString("hello")
	if s != "hello" {
		t.Fatal("expected hello")
	}
	// bytes
	s2 := corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	if s2 == "" {
		t.Fatal("expected non-empty")
	}
	// Result
	r := corejson.New(map[string]int{"a": 1})
	s3 := corejson.AnyTo.SafeJsonPrettyString(r)
	if s3 == "" {
		t.Fatal("expected non-empty")
	}
	// *Result
	s4 := corejson.AnyTo.SafeJsonPrettyString(r.ToPtr())
	if s4 == "" {
		t.Fatal("expected non-empty")
	}
	// any
	s5 := corejson.AnyTo.SafeJsonPrettyString(42)
	if s5 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovJson_S1_49_AnyTo_JsonString(t *testing.T) {
	s := corejson.AnyTo.JsonString("hello")
	if s != "hello" {
		t.Fatal("expected hello")
	}
	s2 := corejson.AnyTo.JsonString([]byte(`test`))
	if s2 == "" {
		t.Fatal("expected non-empty")
	}
	r := corejson.New(1)
	s3 := corejson.AnyTo.JsonString(r)
	if s3 == "" {
		t.Fatal("expected non-empty")
	}
	s4 := corejson.AnyTo.JsonString(r.ToPtr())
	if s4 == "" {
		t.Fatal("expected non-empty")
	}
	s5 := corejson.AnyTo.JsonString(42)
	if s5 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovJson_S1_50_AnyTo_JsonStringWithErr(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("hello")
	if err != nil || s != "hello" {
		t.Fatal("expected hello")
	}
	s2, err2 := corejson.AnyTo.JsonStringWithErr([]byte(`test`))
	if err2 != nil || s2 == "" {
		t.Fatal("expected string")
	}
	r := corejson.New(1)
	s3, err3 := corejson.AnyTo.JsonStringWithErr(r)
	if err3 != nil || s3 == "" {
		t.Fatal("expected string")
	}
	s4, err4 := corejson.AnyTo.JsonStringWithErr(r.ToPtr())
	if err4 != nil || s4 == "" {
		t.Fatal("expected string")
	}
	s5, err5 := corejson.AnyTo.JsonStringWithErr(42)
	if err5 != nil || s5 == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJson_S1_51_AnyTo_PrettyStringWithError(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	if err != nil || s != "hello" {
		t.Fatal("expected hello")
	}
	s2, err2 := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	if err2 != nil || s2 == "" {
		t.Fatal("expected string")
	}
	s3, err3 := corejson.AnyTo.PrettyStringWithError(42)
	if err3 != nil || s3 == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJson_S1_52_AnyTo_SerializedFieldsMap(t *testing.T) {
	fm, err := corejson.AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
	_ = fm
	_ = err
}

// --- CastingAny ---

func Test_CovJson_S1_53_CastingAny_FromToDefault(t *testing.T) {
	from := map[string]int{"a": 1}
	var to map[string]int
	err := corejson.CastAny.FromToDefault(from, &to)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJson_S1_54_CastingAny_FromToReflection(t *testing.T) {
	from := map[string]int{"a": 1}
	var to map[string]int
	err := corejson.CastAny.FromToReflection(from, &to)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJson_S1_55_CastingAny_OrDeserializeTo(t *testing.T) {
	from := map[string]int{"a": 1}
	var to map[string]int
	err := corejson.CastAny.OrDeserializeTo(from, &to)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJson_S1_56_CastingAny_FromBytes(t *testing.T) {
	var to map[string]int
	err := corejson.CastAny.FromToDefault([]byte(`{"a":1}`), &to)
	if err != nil || to["a"] != 1 {
		t.Fatal("expected a=1")
	}
}

func Test_CovJson_S1_57_CastingAny_FromString(t *testing.T) {
	var to map[string]int
	err := corejson.CastAny.FromToDefault(`{"a":1}`, &to)
	if err != nil || to["a"] != 1 {
		t.Fatal("expected a=1")
	}
}
