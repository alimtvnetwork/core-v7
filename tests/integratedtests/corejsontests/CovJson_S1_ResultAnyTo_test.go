package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 1: Result methods, anyTo, BytesCloneIf, etc.
// ══════════════════════════════════════════════════════════════════════════════

// --- BytesCloneIf / BytesDeepClone / BytesToString ---

func Test_CovJson_S1_01_BytesCloneIf(t *testing.T) {
	// Arrange
	b := []byte(`{"a":1}`)
	r := corejson.BytesCloneIf(true, b)

	// Act
	actual := args.Map{"result": len(r) != len(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same len", actual)
	// not deep clone
	r2 := corejson.BytesCloneIf(false, b)
	actual = args.Map{"result": len(r2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	// empty input
	r3 := corejson.BytesCloneIf(true, []byte{})
	actual = args.Map{"result": len(r3) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJson_S1_02_BytesDeepClone(t *testing.T) {
	// Arrange
	b := []byte(`{"a":1}`)
	r := corejson.BytesDeepClone(b)

	// Act
	actual := args.Map{"result": len(r) != len(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same len", actual)
	r2 := corejson.BytesDeepClone(nil)
	actual = args.Map{"result": len(r2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJson_S1_03_BytesToString(t *testing.T) {
	// Arrange
	r := corejson.BytesToString([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := corejson.BytesToString(nil)
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJson_S1_04_BytesToPrettyString(t *testing.T) {
	// Arrange
	r := corejson.BytesToPrettyString([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := corejson.BytesToPrettyString(nil)
	actual = args.Map{"result": r2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// --- New / NewPtr ---

func Test_CovJson_S1_05_New(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	actual = args.Map{"result": r.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)
}

func Test_CovJson_S1_06_NewPtr(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": r == nil || r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

// --- JsonString / JsonStringOrErrMsg ---

func Test_CovJson_S1_07_JsonString(t *testing.T) {
	// Arrange
	s, err := corejson.JsonString(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json string", actual)
}

func Test_CovJson_S1_08_JsonStringOrErrMsg(t *testing.T) {
	// Arrange
	s := corejson.JsonStringOrErrMsg(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// --- Result methods ---

func newTestResult() *corejson.Result {
	r := corejson.New(map[string]int{"a": 1, "b": 2})
	return r.ToPtr()
}

func Test_CovJson_S1_09_Result_Map(t *testing.T) {
	// Arrange
	r := newTestResult()
	m := r.Map()

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil
	var nilR *corejson.Result
	m2 := nilR.Map()
	actual = args.Map{"result": len(m2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJson_S1_10_Result_BytesTypeName_SafeBytesTypeName(t *testing.T) {
	// Arrange
	r := newTestResult()
	_ = r.BytesTypeName()
	_ = r.SafeBytesTypeName()
	// nil
	var nilR *corejson.Result

	// Act
	actual := args.Map{"result": nilR.BytesTypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJson_S1_11_Result_SafeString_JsonString_JsonStringPtr(t *testing.T) {
	// Arrange
	r := newTestResult()
	s := r.SafeString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	js := r.JsonString()
	actual = args.Map{"result": js == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	jsp := r.JsonStringPtr()
	actual = args.Map{"result": jsp == nil || *jsp == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil
	var nilR *corejson.Result
	nsp := nilR.JsonStringPtr()
	actual = args.Map{"result": nsp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil ptr", actual)
}

func Test_CovJson_S1_12_Result_PrettyJsonString(t *testing.T) {
	// Arrange
	r := newTestResult()
	pj := r.PrettyJsonString()

	// Act
	actual := args.Map{"result": pj == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.PrettyJsonString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJson_S1_13_Result_PrettyJsonStringOrErrString(t *testing.T) {
	// Arrange
	r := newTestResult()
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil
	var nilR *corejson.Result
	s2 := nilR.PrettyJsonStringOrErrString()
	actual = args.Map{"result": s2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty error string", actual)
	// with error
	errR := &corejson.Result{Error: errors.New("test")}
	s3 := errR.PrettyJsonStringOrErrString()
	actual = args.Map{"result": s3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovJson_S1_14_Result_Length(t *testing.T) {
	// Arrange
	r := newTestResult()

	// Act
	actual := args.Map{"result": r.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero", actual)
	// nil
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJson_S1_15_Result_HasError_ErrorString_IsEmptyError(t *testing.T) {
	// Arrange
	r := newTestResult()

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	es := r.ErrorString()
	actual = args.Map{"result": es != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": r.IsEmptyError()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovJson_S1_16_Result_IsErrorEqual(t *testing.T) {
	// Arrange
	r := newTestResult()

	// Act
	actual := args.Map{"result": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	errR := &corejson.Result{Error: errors.New("test")}
	actual = args.Map{"result": errR.IsErrorEqual(errors.New("test"))}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": errR.IsErrorEqual(nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": errR.IsErrorEqual(errors.New("other"))}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovJson_S1_17_Result_String(t *testing.T) {
	// Arrange
	r := newTestResult()
	s := r.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovJson_S1_18_Result_SafeBytes_SafeValues_SafeNonIssueBytes_Values(t *testing.T) {
	// Arrange
	r := newTestResult()
	sb := r.SafeBytes()

	// Act
	actual := args.Map{"result": len(sb) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	sv := r.SafeValues()
	actual = args.Map{"result": len(sv) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	_ = r.SafeNonIssueBytes()
	_ = r.Values()
	_ = r.SafeValuesPtr()
}

func Test_CovJson_S1_19_Result_Raw_RawMust(t *testing.T) {
	// Arrange
	r := newTestResult()
	b, err := r.Raw()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	_ = r.RawMust()
}

func Test_CovJson_S1_20_Result_Raw_Nil(t *testing.T) {
	// Arrange
	var nilR *corejson.Result
	_, err := nilR.Raw()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJson_S1_21_Result_RawString_RawStringMust_RawErrString_RawPrettyString(t *testing.T) {
	// Arrange
	r := newTestResult()
	s, err := r.RawString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	sm := r.RawStringMust()
	actual = args.Map{"result": sm == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	rb, re := r.RawErrString()
	actual = args.Map{"result": len(rb) == 0 || re != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	ps, perr := r.RawPrettyString()
	actual = args.Map{"result": perr != nil || ps == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)
}

func Test_CovJson_S1_22_Result_MeaningfulError_MeaningfulErrorMessage(t *testing.T) {
	// Arrange
	r := newTestResult()

	// Act
	actual := args.Map{"result": r.MeaningfulError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": r.MeaningfulErrorMessage() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	// nil result
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.MeaningfulError() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// empty bytes
	emptyR := &corejson.Result{}
	actual = args.Map{"result": emptyR.MeaningfulError() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for empty bytes", actual)
	// error + bytes
	errR := &corejson.Result{Bytes: []byte(`{"a":1}`), Error: errors.New("some err")}
	actual = args.Map{"result": errR.MeaningfulError() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJson_S1_23_Result_HasIssuesOrEmpty_HasSafeItems_IsAnyNull(t *testing.T) {
	// Arrange
	r := newTestResult()

	// Act
	actual := args.Map{"result": r.HasIssuesOrEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": r.HasSafeItems()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": r.IsAnyNull()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovJson_S1_24_Result_HasBytes_HasJsonBytes_HasJson_IsEmptyJson_IsEmptyJsonBytes(t *testing.T) {
	// Arrange
	r := newTestResult()

	// Act
	actual := args.Map{"result": r.HasBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": r.HasJsonBytes()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": r.HasJson()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": r.IsEmptyJson()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": r.IsEmptyJsonBytes()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// empty json "{}"
	emptyJson := &corejson.Result{Bytes: []byte("{}")}
	actual = args.Map{"result": emptyJson.IsEmptyJsonBytes()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for {}", actual)
}

func Test_CovJson_S1_25_Result_Serialize_SerializeMust(t *testing.T) {
	// Arrange
	r := newTestResult()
	b, err := r.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected serialization", actual)
	_ = r.SerializeMust()
	// nil
	var nilR *corejson.Result
	_, err2 := nilR.Serialize()
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// with error
	errR := &corejson.Result{Error: errors.New("test")}
	_, err3 := errR.Serialize()
	actual = args.Map{"result": err3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJson_S1_26_Result_SerializeSkipExistingIssues(t *testing.T) {
	// Arrange
	r := newTestResult()
	b, err := r.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected serialization", actual)
	// with issues
	errR := &corejson.Result{Error: errors.New("test")}
	b2, err2 := errR.SerializeSkipExistingIssues()
	actual = args.Map{"result": err2 != nil || b2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJson_S1_27_Result_Unmarshal_Deserialize(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	rp := r.ToPtr()
	var m map[string]int
	err := rp.Deserialize(&m)

	// Act
	actual := args.Map{"result": err != nil || m["a"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected deserialization", actual)
	// nil result
	var nilR *corejson.Result
	err2 := nilR.Unmarshal(&m)
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// result with error
	errR := &corejson.Result{Error: errors.New("test")}
	err3 := errR.Unmarshal(&m)
	actual = args.Map{"result": err3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJson_S1_28_Result_UnmarshalSkipExistingIssues(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1}).ToPtr()
	var m map[string]int
	err := r.UnmarshalSkipExistingIssues(&m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// with issues
	errR := &corejson.Result{Error: errors.New("test")}
	err2 := errR.UnmarshalSkipExistingIssues(&m)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJson_S1_29_Result_UnmarshalResult(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1}).ToPtr()
	_, _ = r.UnmarshalResult()
}

func Test_CovJson_S1_30_Result_JsonModel_JsonModelAny(t *testing.T) {
	// Arrange
	r := newTestResult()
	jm := r.JsonModel()

	// Act
	actual := args.Map{"result": jm.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	jma := r.JsonModelAny()
	actual = args.Map{"result": jma == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// nil
	var nilR *corejson.Result
	njm := nilR.JsonModel()
	actual = args.Map{"result": njm.Error == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJson_S1_31_Result_Json_JsonPtr(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	j := r.Json()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	jp := r.JsonPtr()
	actual = args.Map{"result": jp == nil || jp.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJson_S1_32_Result_CloneError(t *testing.T) {
	// Arrange
	r := newTestResult()

	// Act
	actual := args.Map{"result": r.CloneError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	errR := &corejson.Result{Error: errors.New("test")}
	actual = args.Map{"result": errR.CloneError() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJson_S1_33_Result_Ptr_NonPtr_ToPtr_ToNonPtr(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	p := r.Ptr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	np := p.NonPtr()
	actual = args.Map{"result": np.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	tp := r.ToPtr()
	actual = args.Map{"result": tp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	tnp := r.ToNonPtr()
	actual = args.Map{"result": tnp.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil NonPtr
	var nilR *corejson.Result
	nnp := nilR.NonPtr()
	actual = args.Map{"result": nnp.Error == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJson_S1_34_Result_IsEqualPtr(t *testing.T) {
	// Arrange
	r1 := newTestResult()
	r2 := newTestResult()

	// Act
	actual := args.Map{"result": r1.IsEqualPtr(r2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	// nil both
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.IsEqualPtr(nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil==nil", actual)
	actual = args.Map{"result": nilR.IsEqualPtr(r1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// different
	r3 := corejson.New(map[string]int{"c": 3}).ToPtr()
	actual = args.Map{"result": r1.IsEqualPtr(r3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovJson_S1_35_Result_IsEqual(t *testing.T) {
	// Arrange
	r1 := corejson.New(map[string]int{"a": 1})
	r2 := corejson.New(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": r1.IsEqual(r2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_CovJson_S1_36_Result_CombineErrorWithRefString_Error(t *testing.T) {
	// Arrange
	r := newTestResult()
	s := r.CombineErrorWithRefString("ref")

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	e := r.CombineErrorWithRefError("ref")
	actual = args.Map{"result": e != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	errR := &corejson.Result{Error: errors.New("test")}
	s2 := errR.CombineErrorWithRefString("ref")
	actual = args.Map{"result": s2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	e2 := errR.CombineErrorWithRefError("ref")
	actual = args.Map{"result": e2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJson_S1_37_Result_BytesError(t *testing.T) {
	// Arrange
	r := newTestResult()
	be := r.BytesError()

	// Act
	actual := args.Map{"result": be == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// nil
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.BytesError() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJson_S1_38_Result_Dispose(t *testing.T) {
	// Arrange
	r := newTestResult()
	r.Dispose()

	// Act
	actual := args.Map{"result": r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	// nil dispose
	var nilR *corejson.Result
	nilR.Dispose()
}

func Test_CovJson_S1_39_Result_CloneIf_Clone_ClonePtr(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	c := r.CloneIf(true, true)

	// Act
	actual := args.Map{"result": c.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	c2 := r.CloneIf(false, false)
	actual = args.Map{"result": c2.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	c3 := r.Clone(true)
	actual = args.Map{"result": c3.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	c4 := r.Clone(false)
	actual = args.Map{"result": c4.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	cp := r.ToPtr().ClonePtr(true)
	actual = args.Map{"result": cp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// nil clone
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.ClonePtr(true) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	// empty result clone
	emptyR := corejson.Result{}
	ec := emptyR.Clone(true)
	actual = args.Map{"result": ec.Error != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJson_S1_40_Result_AsJsonContractsBinder_AsJsoner_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	_ = r.AsJsonContractsBinder()
	_ = r.AsJsoner()
	_ = r.AsJsonParseSelfInjector()
}

func Test_CovJson_S1_41_Result_PrettyJsonBuffer(t *testing.T) {
	// Arrange
	r := newTestResult()
	buf, err := r.PrettyJsonBuffer("", "  ")

	// Act
	actual := args.Map{"result": err != nil || buf.Len() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected buffer", actual)
	// empty
	emptyR := &corejson.Result{}
	buf2, _ := emptyR.PrettyJsonBuffer("", "  ")
	actual = args.Map{"result": buf2.Len() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJson_S1_42_Result_HasAnyItem_IsEmpty(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": r.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": r.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// --- AnyTo ---

func Test_CovJson_S1_43_AnyTo_SerializedJsonResult(t *testing.T) {
	// Arrange
	// from map
	r := corejson.AnyTo.SerializedJsonResult(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// from string
	r2 := corejson.AnyTo.SerializedJsonResult(`{"a":1}`)
	actual = args.Map{"result": r2.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// from bytes
	r3 := corejson.AnyTo.SerializedJsonResult([]byte(`{"a":1}`))
	actual = args.Map{"result": r3.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// from Result
	r4 := corejson.AnyTo.SerializedJsonResult(corejson.New(1))
	actual = args.Map{"result": r4.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// from *Result
	rp := corejson.New(1).ToPtr()
	r5 := corejson.AnyTo.SerializedJsonResult(rp)
	actual = args.Map{"result": r5.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// from nil
	r6 := corejson.AnyTo.SerializedJsonResult(nil)
	actual = args.Map{"result": r6.Error == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// from error
	r7 := corejson.AnyTo.SerializedJsonResult(errors.New("test"))
	_ = r7
}

func Test_CovJson_S1_44_AnyTo_SerializedRaw(t *testing.T) {
	// Arrange
	b, err := corejson.AnyTo.SerializedRaw(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_CovJson_S1_45_AnyTo_SerializedString(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.SerializedString(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJson_S1_46_AnyTo_SerializedSafeString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedSafeString(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovJson_S1_47_AnyTo_SafeJsonString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonString(1)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovJson_S1_48_AnyTo_SafeJsonPrettyString(t *testing.T) {
	// Arrange
	// string passthrough
	s := corejson.AnyTo.SafeJsonPrettyString("hello")

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	// bytes
	s2 := corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	actual = args.Map{"result": s2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// Result
	r := corejson.New(map[string]int{"a": 1})
	s3 := corejson.AnyTo.SafeJsonPrettyString(r)
	actual = args.Map{"result": s3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// *Result
	s4 := corejson.AnyTo.SafeJsonPrettyString(r.ToPtr())
	actual = args.Map{"result": s4 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// any
	s5 := corejson.AnyTo.SafeJsonPrettyString(42)
	actual = args.Map{"result": s5 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovJson_S1_49_AnyTo_JsonString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonString("hello")

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	s2 := corejson.AnyTo.JsonString([]byte(`test`))
	actual = args.Map{"result": s2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r := corejson.New(1)
	s3 := corejson.AnyTo.JsonString(r)
	actual = args.Map{"result": s3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	s4 := corejson.AnyTo.JsonString(r.ToPtr())
	actual = args.Map{"result": s4 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	s5 := corejson.AnyTo.JsonString(42)
	actual = args.Map{"result": s5 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovJson_S1_50_AnyTo_JsonStringWithErr(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.JsonStringWithErr("hello")

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	s2, err2 := corejson.AnyTo.JsonStringWithErr([]byte(`test`))
	actual = args.Map{"result": err2 != nil || s2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	r := corejson.New(1)
	s3, err3 := corejson.AnyTo.JsonStringWithErr(r)
	actual = args.Map{"result": err3 != nil || s3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	s4, err4 := corejson.AnyTo.JsonStringWithErr(r.ToPtr())
	actual = args.Map{"result": err4 != nil || s4 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	s5, err5 := corejson.AnyTo.JsonStringWithErr(42)
	actual = args.Map{"result": err5 != nil || s5 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJson_S1_51_AnyTo_PrettyStringWithError(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError("hello")

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	s2, err2 := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	actual = args.Map{"result": err2 != nil || s2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	s3, err3 := corejson.AnyTo.PrettyStringWithError(42)
	actual = args.Map{"result": err3 != nil || s3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJson_S1_52_AnyTo_SerializedFieldsMap(t *testing.T) {
	fm, err := corejson.AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
	_ = fm
	_ = err
}

// --- CastingAny ---

func Test_CovJson_S1_53_CastingAny_FromToDefault(t *testing.T) {
	// Arrange
	from := map[string]int{"a": 1}
	var to map[string]int
	err := corejson.CastAny.FromToDefault(from, &to)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJson_S1_54_CastingAny_FromToReflection(t *testing.T) {
	// Arrange
	from := map[string]int{"a": 1}
	var to map[string]int
	err := corejson.CastAny.FromToReflection(from, &to)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJson_S1_55_CastingAny_OrDeserializeTo(t *testing.T) {
	// Arrange
	from := map[string]int{"a": 1}
	var to map[string]int
	err := corejson.CastAny.OrDeserializeTo(from, &to)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJson_S1_56_CastingAny_FromBytes(t *testing.T) {
	// Arrange
	var to map[string]int
	err := corejson.CastAny.FromToDefault([]byte(`{"a":1}`), &to)

	// Act
	actual := args.Map{"result": err != nil || to["a"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a=1", actual)
}

func Test_CovJson_S1_57_CastingAny_FromString(t *testing.T) {
	// Arrange
	var to map[string]int
	err := corejson.CastAny.FromToDefault(`{"a":1}`, &to)

	// Act
	actual := args.Map{"result": err != nil || to["a"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a=1", actual)
}
