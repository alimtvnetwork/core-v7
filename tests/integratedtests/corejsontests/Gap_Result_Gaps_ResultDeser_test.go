package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Result.go uncovered branches ──

func Test_Gap_Result_HandleErrorWithMsg_NoPanic(t *testing.T) {
	r := corejson.NewResult.Any("x")
	r.HandleErrorWithMsg("prefix msg")
}

func Test_Gap_Result_ParseInjectUsingJsonMust_NoPanic(t *testing.T) {
	inner := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"hello"`), TypeName: "test"})
	target := &corejson.Result{}
	_ = target.ParseInjectUsingJsonMust(inner.Ptr())
}

func Test_Gap_Result_Unmarshal_BadPayload(t *testing.T) {
	// Arrange
	// Valid result with bytes that don't unmarshal to target type
	r := corejson.NewResult.AnyPtr("hello")
	var out int
	err := r.Unmarshal(&out)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error for type mismatch", actual)
}

func Test_Gap_Result_UnmarshalSkipExistingIssues_BadPayload(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	var out int
	err := r.UnmarshalSkipExistingIssues(&out)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for bad payload", actual)
}

func Test_Gap_Result_MeaningfulError_EmptyBytesWithError(t *testing.T) {
	// Arrange
	// Has error AND empty bytes
	r := &corejson.Result{
		Bytes:    []byte{},
		Error:    errors.New("some err"),
		TypeName: "TestType",
	}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Gap_Result_MeaningfulError_HasErrorAndPayload(t *testing.T) {
	// Arrange
	// Has error AND has payload
	r := &corejson.Result{
		Bytes:    []byte(`"payload"`),
		Error:    errors.New("some err"),
		TypeName: "TestType",
	}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Gap_Result_String_WithNilBytes(t *testing.T) {
	// Arrange
	r := corejson.Result{}
	s := r.String()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil bytes", actual)
}

// ── deserializerLogic uncovered methods ──

func Test_Gap_Deserialize_UsingSerializerFuncTo(t *testing.T) {
	// Arrange
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var s string
	err := corejson.Deserialize.UsingSerializerFuncTo(fn, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_Gap_Deserialize_UsingSerializerFuncTo_Nil(t *testing.T) {
	var s string
	// nil func returns nil result
	r := corejson.Deserialize.UsingSerializerFuncTo(nil, &s)
	_ = r
}

// ── deserializeFromBytesTo uncovered ──

func Test_Gap_BytesTo_ResultCollection(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc, err := corejson.Deserialize.BytesTo.ResultCollection(b)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": rc.Length() == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected items", actual)
}

func Test_Gap_BytesTo_ResultCollectionMust(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc := corejson.Deserialize.BytesTo.ResultCollectionMust(b)

	// Act
	actual := args.Map{"result": rc.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected items", actual)
}

func Test_Gap_BytesTo_ResultsPtrCollection(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc, err := corejson.Deserialize.BytesTo.ResultsPtrCollection(b)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	_ = rc
}

func Test_Gap_BytesTo_ResultsPtrCollectionMust(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc := corejson.Deserialize.BytesTo.ResultsPtrCollectionMust(b)
	_ = rc
}

func Test_Gap_BytesTo_MapResults(t *testing.T) {
	// Arrange
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	b, _ := corejson.Serialize.Raw(m)
	mr, err := corejson.Deserialize.BytesTo.MapResults(b)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	_ = mr
}

func Test_Gap_BytesTo_MapResultsMust(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	b, _ := corejson.Serialize.Raw(m)
	mr := corejson.Deserialize.BytesTo.MapResultsMust(b)
	_ = mr
}

// ── deserializeFromResultTo uncovered ──

func Test_Gap_ResultTo_ResultCollection(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(c)
	rc, err := corejson.Deserialize.ResultTo.ResultCollection(jr)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": rc.Length() == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected items", actual)
}

func Test_Gap_ResultTo_ResultCollectionMust(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(c)
	rc := corejson.Deserialize.ResultTo.ResultCollectionMust(jr)

	// Act
	actual := args.Map{"result": rc.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected items", actual)
}

func Test_Gap_ResultTo_ResultsPtrCollection(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	jr := corejson.NewResult.AnyPtr(c)
	_, err := corejson.Deserialize.ResultTo.ResultsPtrCollection(jr)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Gap_ResultTo_ResultsPtrCollectionMust(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	jr := corejson.NewResult.AnyPtr(c)
	_ = corejson.Deserialize.ResultTo.ResultsPtrCollectionMust(jr)
}

func Test_Gap_ResultTo_Result(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_, err := corejson.Deserialize.ResultTo.Result(jr)
	_ = err
}

func Test_Gap_ResultTo_ResultMust(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_ = corejson.Deserialize.ResultTo.ResultMust(jr)
}

func Test_Gap_ResultTo_ResultPtr(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_, err := corejson.Deserialize.ResultTo.ResultPtr(jr)
	_ = err
}

func Test_Gap_ResultTo_ResultPtrMust(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_ = corejson.Deserialize.ResultTo.ResultPtrMust(jr)
}

func Test_Gap_ResultTo_MapResults(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	jr := corejson.NewResult.AnyPtr(m)
	_, err := corejson.Deserialize.ResultTo.MapResults(jr)
	_ = err
}

func Test_Gap_ResultTo_MapResultsMust(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	jr := corejson.NewResult.AnyPtr(m)
	_ = corejson.Deserialize.ResultTo.MapResultsMust(jr)
}

func Test_Gap_ResultTo_Bytes(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_, err := corejson.Deserialize.ResultTo.Bytes(jr)
	_ = err
}

func Test_Gap_ResultTo_BytesMust(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_ = corejson.Deserialize.ResultTo.BytesMust(jr)
}

// ── newResultCreator uncovered ──

func Test_Gap_NewResult_UsingBytesError_NonNil(t *testing.T) {
	// Arrange
	// import coredata for BytesError is needed - use creator with valid data
	r := corejson.NewResult.Any("hello")
	be := r.BytesError()

	// Act
	actual := args.Map{"result": be == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil BytesError", actual)
	r2 := corejson.NewResult.UsingBytesError(be)
	actual = args.Map{"result": r2.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_Gap_NewResult_UsingBytesError_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingBytesError(nil)

	// Act
	actual := args.Map{"result": r.Bytes != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil bytes", actual)
}

func Test_Gap_NewResult_DeserializeUsingResult_WithIssues(t *testing.T) {
	// Arrange
	errResult := corejson.NewResult.ErrorPtr(errors.New("bad"))
	r := corejson.NewResult.DeserializeUsingResult(errResult)

	// Act
	actual := args.Map{"result": r == nil || !r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error result", actual)
}

func Test_Gap_NewResult_FromStringer(t *testing.T) {
	// Exercise serializer pretty string (FromStringer requires fmt.Stringer, not error)
	_ = corejson.Serialize.ToPrettyStringIncludingErr("hello")
}

// ── castingAny uncovered ──

func Test_Gap_CastAny_FromToOption_NilFrom(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToOption(true, nil, &out)
	// nil from should return an error (not applicable) and fall through
	_ = err
}

func Test_Gap_CastAny_FromToOption_NilTo(t *testing.T) {
	err := corejson.CastAny.FromToOption(true, "hello", nil)
	_ = err
}

func Test_Gap_CastAny_FromToOption_Error(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.FromToOption(false, errors.New(`"hello"`), &out)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for error-to-string deserialization", actual)
	actual = args.Map{"result": out}
	expected = args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_Gap_CastAny_FromToOption_ErrorInvalidJson(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.FromToOption(false, errors.New("not json"), &out)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid json in error message", actual)
}

func Test_Gap_CastAny_FromToOption_NilError(t *testing.T) {
	var nilErr error
	var out string
	err := corejson.CastAny.FromToOption(false, nilErr, &out)
	// nil error case goes to fallback serialization
	_ = err
}

// ── Deserialize.Result / ResultPtr / ResultMust / ResultPtrMust (on deserializerLogic struct) ──

func Test_Gap_Deserialize_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	b, _ := r.Serialize()
	_, err := corejson.Deserialize.Result(b)
	_ = err
}

func Test_Gap_Deserialize_ResultPtr(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	b, _ := r.Serialize()
	_, err := corejson.Deserialize.ResultPtr(b)
	_ = err
}

// ── newResultsCollectionCreator uncovered ──

func Test_Gap_NewResultsCollection_DeserializeUsingResult(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(c)
	rc, err := corejson.NewResultsCollection.DeserializeUsingResult(jr)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": rc.Length() == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected items", actual)
}

func Test_Gap_NewResultsCollection_DeserializeUsingResult_Issues(t *testing.T) {
	// Arrange
	errResult := corejson.NewResult.ErrorPtr(errors.New("bad"))
	_, err := corejson.NewResultsCollection.DeserializeUsingResult(errResult)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Gap_NewResultsCollection_UnmarshalUsingBytes(t *testing.T) {
	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc, err := corejson.NewResultsCollection.UnmarshalUsingBytes(b)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": rc.Length() == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected items", actual)
}

// ── newResultsPtrCollectionCreator uncovered ──

func Test_Gap_NewResultsPtrCollection_DeserializeUsingResult(t *testing.T) {
	// Arrange
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	jr := corejson.NewResult.AnyPtr(c)
	_, err := corejson.NewResultsPtrCollection.DeserializeUsingResult(jr)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Gap_NewResultsPtrCollection_Jsoners(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Jsoners()
}

func Test_Gap_NewResultsPtrCollection_JsonersPlusCap(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.JsonersPlusCap(true, 5)
}

// ── newMapResultsCreator uncovered ──

func Test_Gap_NewMapResults_DeserializeUsingResult(t *testing.T) {
	// Arrange
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	jr := corejson.NewResult.AnyPtr(m)
	_, err := corejson.NewMapResults.DeserializeUsingResult(jr)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Gap_NewMapResults_DeserializeUsingResult_Issues(t *testing.T) {
	// Arrange
	errResult := corejson.NewResult.ErrorPtr(errors.New("bad"))
	_, err := corejson.NewMapResults.DeserializeUsingResult(errResult)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── newBytesCollectionCreator uncovered ──

func Test_Gap_NewBytesCollection_DeserializeUsingResult(t *testing.T) {
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"x"`))
	jr := corejson.NewResult.AnyPtr(c)
	_, err := corejson.NewBytesCollection.DeserializeUsingResult(jr)
	_ = err
}

func Test_Gap_NewBytesCollection_Jsoners(t *testing.T) {
	_ = corejson.NewBytesCollection.Jsoners()
}

// ── Serializer uncovered: FromStringer ──

func Test_Gap_Serialize_Apply_MarshalError(t *testing.T) {
	// Arrange
	ch := make(chan int)
	r := corejson.Serialize.Apply(ch)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error for channel", actual)
}
