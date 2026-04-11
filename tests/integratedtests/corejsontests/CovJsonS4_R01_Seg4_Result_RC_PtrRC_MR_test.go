package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 4: Result deep methods, ResultsCollection,
//                     ResultsPtrCollection, MapResults — all remaining branches
// ══════════════════════════════════════════════════════════════════════════════

// --- Result deep methods ---

func Test_CovJsonS4_R01_Map_WithFields(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	m := r.Map()

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty map", actual)
}

func Test_CovJsonS4_R02_Map_NilResult(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.Map()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map", actual)
}

func Test_CovJsonS4_R03_Map_WithError(t *testing.T) {
	// Arrange
	r := corejson.Result{Error: errors.New("fail"), TypeName: "T"}
	m := r.Map()

	// Act
	actual := args.Map{"result": m["Error"] != "fail"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error in map", actual)
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
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.BytesTypeName() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type name", actual)
	var nr *corejson.Result
	actual := args.Map{"result": nr.BytesTypeName() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJsonS4_R09_SafeBytesTypeName(t *testing.T) {
	r := corejson.New(1)
	_ = r.SafeBytesTypeName()
	var nr *corejson.Result
	_ = nr.SafeBytesTypeName()
}

func Test_CovJsonS4_R10_SafeString(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.SafeString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS4_R11_JsonStringPtr_Cached(t *testing.T) {
	r := corejson.New(1)
	_ = r.JsonStringPtr()
	_ = r.JsonStringPtr() // second call should use cached
}

func Test_CovJsonS4_R12_PrettyJsonBuffer(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	buf, err := r.PrettyJsonBuffer("", "  ")

	// Act
	actual := args.Map{"result": err != nil || buf.Len() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected buffer", actual)
	// empty
	var nr *corejson.Result
	buf2, _ := nr.PrettyJsonBuffer("", "  ")
	_ = buf2
}

func Test_CovJsonS4_R13_PrettyJsonStringOrErrString(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	// nil
	var nr *corejson.Result
	s2 := nr.PrettyJsonStringOrErrString()
	actual := args.Map{"result": s2 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil message", actual)
	// with error
	re := corejson.Result{Error: errors.New("fail")}
	s3 := re.PrettyJsonStringOrErrString()
	actual := args.Map{"result": s3 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error message", actual)
}

func Test_CovJsonS4_R14_ErrorString(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.ErrorString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	re := corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"result": re.ErrorString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error string", actual)
}

func Test_CovJsonS4_R15_IsErrorEqual(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	re := corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"result": re.IsErrorEqual(errors.New("fail"))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": re.IsErrorEqual(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": re.IsErrorEqual(errors.New("other"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovJsonS4_R16_String_WithError(t *testing.T) {
	// Arrange
	re := corejson.Result{Error: errors.New("fail"), TypeName: "T", Bytes: []byte("x")}
	s := re.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	r := corejson.New(1)
	s2 := r.String()
	actual := args.Map{"result": s2 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS4_R17_SafeNonIssueBytes(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": len(r.SafeNonIssueBytes()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	re := corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"result": len(re.SafeNonIssueBytes()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
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
	// Arrange
	r := corejson.Result{TypeName: "T"}
	e := r.MeaningfulError()

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJsonS4_R21_MeaningfulError_WithError(t *testing.T) {
	// Arrange
	r := corejson.Result{Error: errors.New("fail"), Bytes: []byte("x"), TypeName: "T"}
	e := r.MeaningfulError()

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJsonS4_R22_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": r.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovJsonS4_R23_IsEmptyJson_HasJson_HasBytes_HasJsonBytes(t *testing.T) {
	r := corejson.New(1)
	_ = r.IsEmptyJson()
	_ = r.HasJson()
	_ = r.HasBytes()
	_ = r.HasJsonBytes()
}

func Test_CovJsonS4_R24_HasSafeItems(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.HasSafeItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovJsonS4_R25_IsAnyNull_HasIssuesOrEmpty(t *testing.T) {
	r := corejson.New(1)
	_ = r.IsAnyNull()
	_ = r.HasIssuesOrEmpty()
}

func Test_CovJsonS4_R26_MeaningfulErrorMessage(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.MeaningfulErrorMessage() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
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
	// Arrange
	var r *corejson.Result
	var m map[string]int
	err := r.Unmarshal(&m)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJsonS4_R30_Unmarshal_WithExistingError(t *testing.T) {
	// Arrange
	re := corejson.Result{Error: errors.New("fail"), Bytes: []byte("x"), TypeName: "T"}
	var m map[string]int
	err := re.Unmarshal(&m)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJsonS4_R31_SerializeSkipExistingIssues(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	_, _ = r.SerializeSkipExistingIssues()
	// empty
	re := corejson.Result{Error: errors.New("fail")}
	b, e := re.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{"result": b != nil || e != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil,nil", actual)
}

func Test_CovJsonS4_R32_Serialize_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	_, err := r.Serialize()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJsonS4_R33_Serialize_WithError(t *testing.T) {
	// Arrange
	re := corejson.Result{Error: errors.New("fail")}
	_, err := re.Serialize()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJsonS4_R34_Serialize_Success(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	b, err := r.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
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
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.CloneError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	re := corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"result": re.CloneError() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
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
	// Arrange
	r := corejson.New(1)
	r2 := corejson.New(1)

	// Act
	actual := args.Map{"result": r.IsEqualPtr(&r2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	var nr *corejson.Result
	actual := args.Map{"result": nr.IsEqualPtr(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": nr.IsEqualPtr(&r)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// same pointer
	rp := r.Ptr()
	actual := args.Map{"result": rp.IsEqualPtr(rp)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovJsonS4_R45_IsEqual(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	r2 := corejson.New(1)

	// Act
	actual := args.Map{"result": r.IsEqual(r2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovJsonS4_R46_CombineErrorWithRefString_CombineErrorWithRefError(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.CombineErrorWithRefString("ref") != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": r.CombineErrorWithRefError("ref") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	re := corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"result": re.CombineErrorWithRefString("ref") == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	actual := args.Map{"result": re.CombineErrorWithRefError("ref") == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJsonS4_R47_BytesError(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	be := r.BytesError()

	// Act
	actual := args.Map{"result": be == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	var nr *corejson.Result
	actual := args.Map{"result": nr.BytesError() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS4_R48_Dispose(t *testing.T) {
	r := corejson.New(1)
	r.Dispose()
	var nr *corejson.Result
	nr.Dispose()
}

func Test_CovJsonS4_R49_CloneIf_Clone_ClonePtr(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	_ = r.CloneIf(true, true)
	_ = r.CloneIf(false, false)
	_ = r.CloneIf(true, false)
	_ = r.Clone(true)
	_ = r.Clone(false)
	_ = r.ClonePtr(true)
	_ = r.ClonePtr(false)
	var nr *corejson.Result

	// Act
	actual := args.Map{"result": nr.ClonePtr(true) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
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
	// Arrange
	rc := corejson.NewResultsCollection.UsingCap(5)
	rc.AddAny(1)
	rc.AddAny(2)

	// Act
	actual := args.Map{"result": rc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": rc.LastIndex() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": rc.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": rc.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovJsonS4_RC02_FirstOrDefault_LastOrDefault(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)

	// Act
	actual := args.Map{"result": rc.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": rc.LastOrDefault() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	empty := corejson.NewResultsCollection.UsingCap(0)
	actual := args.Map{"result": empty.FirstOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
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
	// Arrange
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)

	// Act
	actual := args.Map{"result": rc.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, has := rc.AllErrors()
	actual := args.Map{"result": has}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
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
	// Arrange
	rc := corejson.NewResultsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		rc.AddAny(i)
	}

	// Act
	actual := args.Map{"result": rc.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": rc.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	pages := rc.GetPagedCollection(3)
	actual := args.Map{"result": len(pages) < 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
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
	// Arrange
	rc := corejson.NewResultsCollection.UsingCap(2)
	rc.AddAny(1)
	_ = rc.ShadowClone()
	_ = rc.Clone(true)
	_ = rc.ClonePtr(true)
	_ = rc.ClonePtr(false)
	var nilRC *corejson.ResultsCollection

	// Act
	actual := args.Map{"result": nilRC.ClonePtr(true) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- ResultsPtrCollection remaining branches ---

func Test_CovJsonS4_RPC01_Basic(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.UsingCap(5)
	rpc.AddAny(1)
	rpc.AddAny(2)

	// Act
	actual := args.Map{"result": rpc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": rpc.LastIndex() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": rpc.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": rpc.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovJsonS4_RPC02_FirstOrDefault_LastOrDefault(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)

	// Act
	actual := args.Map{"result": rpc.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": rpc.LastOrDefault() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	empty := corejson.NewResultsPtrCollection.UsingCap(0)
	actual := args.Map{"result": empty.FirstOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
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
	// Arrange
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)

	// Act
	actual := args.Map{"result": rpc.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, has := rpc.AllErrors()
	actual := args.Map{"result": has}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
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
	// Arrange
	rpc := corejson.NewResultsPtrCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		rpc.AddAny(i)
	}

	// Act
	actual := args.Map{"result": rpc.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": rpc.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	pages := rpc.GetPagedCollection(3)
	actual := args.Map{"result": len(pages) < 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
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
	// Arrange
	rpc := corejson.NewResultsPtrCollection.UsingCap(2)
	rpc.AddAny(1)
	_ = rpc.Clone(true)
	_ = rpc.Clone(false)
	var nilRPC *corejson.ResultsPtrCollection

	// Act
	actual := args.Map{"result": nilRPC.Clone(true) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// --- MapResults remaining branches ---

func Test_CovJsonS4_MR01_Basic(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.UsingCap(5)
	mr.AddAny("a", 1)
	mr.AddAny("b", 2)

	// Act
	actual := args.Map{"result": mr.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": mr.LastIndex() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": mr.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": mr.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
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
	// Arrange
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("k", 1)

	// Act
	actual := args.Map{"result": mr.GetByKey("k") == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": mr.GetByKey("missing") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS4_MR06_HasError_AllErrors_GetErrorsStrings(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("k", 1)

	// Act
	actual := args.Map{"result": mr.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, has := mr.AllErrors()
	actual := args.Map{"result": has}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_ = mr.GetErrorsStrings()
	_ = mr.GetErrorsStringsPtr()
	_ = mr.GetErrorsAsSingleString()
	_ = mr.GetErrorsAsSingle()
	empty := corejson.NewMapResults.UsingCap(0)
	_ = empty.GetErrorsStrings()
	_, _ = empty.AllErrors()
}

func Test_CovJsonS4_MR07_AllKeys_AllKeysSorted_AllValues_AllResults_AllResultsCollection(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.UsingCap(2)
	mr.AddAny("b", 2)
	mr.AddAny("a", 1)
	keys := mr.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sorted := mr.AllKeysSorted()
	actual := args.Map{"result": sorted[0] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a first", actual)
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
	// Arrange
	mr := corejson.NewMapResults.UsingCap(10)
	for i := 0; i < 10; i++ {
		mr.AddAny("k"+string(rune('a'+i)), i)
	}

	// Act
	actual := args.Map{"result": mr.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": mr.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	pages := mr.GetPagedCollection(3)
	actual := args.Map{"result": len(pages) < 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	keys := mr.AllKeysSorted()
	_ = mr.GetSinglePageCollection(3, 2, keys)
	// small
	small := corejson.NewMapResults.UsingCap(2)
	small.AddAny("k", 1)
	_ = small.GetPagedCollection(10)
	_ = small.GetSinglePageCollection(10, 1, small.AllKeysSorted())
}

func Test_CovJsonS4_MR15_GetNewMapUsingKeys(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.UsingCap(3)
	mr.AddAny("a", 1)
	mr.AddAny("b", 2)
	sub := mr.GetNewMapUsingKeys(false, "a")

	// Act
	actual := args.Map{"result": sub.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
