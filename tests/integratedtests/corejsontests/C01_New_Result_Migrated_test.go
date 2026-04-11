package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Migrated from Coverage01_Result_test.go ──

func Test_C01_New(t *testing.T) {
	// Arrange
	r := corejson.New("hello")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	actual = args.Map{"result": r.JsonString() != `"hello"`}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected json", actual)
	actual = args.Map{"result": r.TypeName == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type name", actual)
}

func Test_C01_New_MarshalError(t *testing.T) {
	// Arrange
	ch := make(chan int)
	r := corejson.New(ch)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error for channel", actual)
}

func Test_C01_NewPtr(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(42)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": r.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	actual = args.Map{"result": r.JsonString() != "42"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected json", actual)
}

func Test_C01_NewPtr_MarshalError(t *testing.T) {
	// Arrange
	ch := make(chan int)
	r := corejson.NewPtr(ch)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error for channel", actual)
}

func Test_C01_Result_Map(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("test")
	m := r.Map()

	// Act
	actual := args.Map{
		"result": _,
		ok := m["Bytes"]; !ok,
	}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes key", actual)

	rErr := corejson.NewResult.Error(errors.New("fail"))
	m2 := rErr.Map()
	actual = args.Map{
		"result": _,
		ok := m2["Error"]; !ok,
	}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error key", actual)

	var nilR *corejson.Result
	m3 := nilR.Map()
	actual = args.Map{"result": len(m3) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map for nil", actual)
}

func Test_C01_Result_JsonStringPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	s1 := r.JsonStringPtr()

	// Act
	actual := args.Map{"result": s1 == nil || *s1 == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
	s2 := r.JsonStringPtr()
	actual = args.Map{"result": *s1 != *s2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache miss", actual)

	var nilR *corejson.Result
	s3 := nilR.JsonStringPtr()
	actual = args.Map{"result": s3 == nil || *s3 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string for nil", actual)

	emptyR := corejson.Result{}
	s4 := emptyR.JsonStringPtr()
	actual = args.Map{"result": s4 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C01_Result_SafeString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(123)
	s := r.SafeString()

	// Act
	actual := args.Map{"result": s != "123"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected safe string", actual)
}

func Test_C01_Result_PrettyJsonBuffer(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	buf, err := r.PrettyJsonBuffer("", "  ")

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": buf.Len() == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty buffer", actual)

	emptyR := corejson.Result{}
	buf2, _ := emptyR.PrettyJsonBuffer("", "  ")
	actual = args.Map{"result": buf2.Len() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty buffer for empty result", actual)
}

func Test_C01_Result_PrettyJsonString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	s := r.PrettyJsonString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty pretty string", actual)

	var nilR *corejson.Result
	s2 := nilR.PrettyJsonString()
	actual = args.Map{"result": s2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)

	emptyR := &corejson.Result{}
	s3 := emptyR.PrettyJsonString()
	actual = args.Map{"result": s3 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for empty", actual)
}

func Test_C01_Result_PrettyJsonStringOrErrString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(42)
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	rErr := corejson.NewResult.Error(errors.New("boom"))
	s2 := rErr.PrettyJsonStringOrErrString()
	actual = args.Map{"result": s2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error message", actual)

	var nilR *corejson.Result
	s3 := nilR.PrettyJsonStringOrErrString()
	actual = args.Map{"result": s3 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil message", actual)
}

func Test_C01_Result_Length(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hi")

	// Act
	actual := args.Map{"result": r.Length() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero length", actual)
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected zero for nil", actual)
}

func Test_C01_Result_HasError_ErrorString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have error", actual)
	actual = args.Map{"result": r.ErrorString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty error string", actual)

	rErr := corejson.NewResult.Error(errors.New("fail"))
	actual = args.Map{"result": rErr.HasError()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have error", actual)
	actual = args.Map{"result": rErr.ErrorString() != "fail"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error string", actual)
}

func Test_C01_Result_IsErrorEqual(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)

	rErr := corejson.NewResult.Error(errors.New("boom"))
	actual = args.Map{"result": rErr.IsErrorEqual(nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error vs nil should not be equal", actual)
	actual = args.Map{"result": rErr.IsErrorEqual(errors.New("boom"))}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same error message should be equal", actual)
	actual = args.Map{"result": rErr.IsErrorEqual(errors.New("other"))}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different errors should not be equal", actual)
}

func Test_C01_Result_String(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	_ = r.String() // may return empty depending on IsAnyNull
	rErr := corejson.NewResult.Error(errors.New("err"))
	_ = rErr.String()
}

func Test_C01_Result_SafeNonIssueBytes(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(42)

	// Act
	actual := args.Map{"result": len(r.SafeNonIssueBytes()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty bytes", actual)
	rErr := corejson.NewResult.Error(errors.New("e"))
	actual = args.Map{"result": len(rErr.SafeNonIssueBytes()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty bytes for error result", actual)
}

func Test_C01_Result_SafeBytes_Values_SafeValues(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(1)

	// Act
	actual := args.Map{"result": len(r.SafeBytes()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	actual = args.Map{"result": len(r.Values()) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected values", actual)
	actual = args.Map{"result": len(r.SafeValues()) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected safe values", actual)
	actual = args.Map{"result": len(r.SafeValuesPtr()) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected safe values ptr", actual)

	var nilR *corejson.Result
	actual = args.Map{"result": len(nilR.SafeBytes()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	actual = args.Map{"result": len(nilR.SafeValues()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_C01_Result_Raw(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	b, err := r.Raw()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected raw bytes", actual)
	var nilR *corejson.Result
	_, err2 := nilR.Raw()
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C01_Result_RawMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("y")
	b := r.RawMust()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C01_Result_RawString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("z")
	s, err := r.RawString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected raw string", actual)
}

func Test_C01_Result_RawStringMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("a")

	// Act
	actual := args.Map{"result": r.RawStringMust() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_C01_Result_RawErrString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("b")
	b, errStr := r.RawErrString()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	_ = errStr
}

func Test_C01_Result_RawPrettyString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(map[string]int{"k": 1})
	s, err := r.RawPrettyString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)
}

func Test_C01_Result_MeaningfulError(t *testing.T) {
	// Arrange
	var nilR *corejson.Result

	// Act
	actual := args.Map{"result": nilR.MeaningfulError() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	r := corejson.NewResult.Any("good")
	actual = args.Map{"result": r.MeaningfulError() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil error", actual)
	emptyR := &corejson.Result{}
	actual = args.Map{"result": emptyR.MeaningfulError() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for empty bytes", actual)
	rErr := corejson.NewResult.Error(errors.New("boom"))
	actual = args.Map{"result": rErr.MeaningfulError() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C01_Result_MeaningfulErrorMessage(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("ok")

	// Act
	actual := args.Map{"result": r.MeaningfulErrorMessage() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rErr := corejson.NewResult.Error(errors.New("x"))
	actual = args.Map{"result": rErr.MeaningfulErrorMessage() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected message", actual)
}

func Test_C01_Result_IsEmptyError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.IsEmptyError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty error", actual)
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.IsEmptyError()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty error for nil", actual)
}

func Test_C01_Result_HasSafeItems(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.HasSafeItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected safe items", actual)
	rErr := corejson.NewResult.Error(errors.New("e"))
	actual = args.Map{"result": rErr.HasSafeItems()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have safe items", actual)
}

func Test_C01_Result_IsAnyNull(t *testing.T) {
	// Arrange
	var nilR *corejson.Result

	// Act
	actual := args.Map{"result": nilR.IsAnyNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected null for nil", actual)
	r := corejson.Result{}
	actual = args.Map{"result": r.IsAnyNull()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected null for empty bytes", actual)
	r2 := corejson.NewResult.Any(1)
	actual = args.Map{"result": r2.IsAnyNull()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be null", actual)
}

func Test_C01_Result_HasIssuesOrEmpty(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.HasIssuesOrEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have issues", actual)
	r2 := corejson.NewResult.Error(errors.New("e"))
	actual = args.Map{"result": r2.HasIssuesOrEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have issues", actual)
}

func Test_C01_Result_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	r := corejson.Result{}

	// Act
	actual := args.Map{"result": r.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": r.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have items", actual)
	r2 := corejson.NewResult.Any("x")
	actual = args.Map{"result": r2.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	actual = args.Map{"result": r2.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_C01_Result_IsEmptyJson_HasJson_HasBytes_HasJsonBytes(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.IsEmptyJson()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty json", actual)
	actual = args.Map{"result": r.HasJson()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have json", actual)
	actual = args.Map{"result": r.HasBytes()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have bytes", actual)
	actual = args.Map{"result": r.HasJsonBytes()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have json bytes", actual)
	empty := corejson.Result{Bytes: []byte("{}")}
	actual = args.Map{"result": empty.IsEmptyJsonBytes()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty json for {}", actual)
}

func Test_C01_Result_Deserialize_Unmarshal(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(map[string]string{"k": "v"})
	var out map[string]string
	err := r.Deserialize(&out)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": out["k"] != "v"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected value", actual)
	r.DeserializeMust(&out)

	var nilR *corejson.Result
	err2 := nilR.Unmarshal(&out)
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)

	rErr := corejson.NewResult.Error(errors.New("e"))
	err3 := rErr.Unmarshal(&out)
	actual = args.Map{"result": err3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C01_Result_UnmarshalSkipExistingIssues(t *testing.T) {
	// Arrange
	rErr := corejson.NewResult.Error(errors.New("e"))
	var out string
	err := rErr.UnmarshalSkipExistingIssues(&out)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should skip and return nil", actual)
	r := corejson.NewResult.Any("hello")
	err2 := r.UnmarshalSkipExistingIssues(&out)
	actual = args.Map{"result": err2}
	expected = args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err2", actual)
	actual = args.Map{"result": out != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected value", actual)
}

func Test_C01_Result_Serialize(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(42)
	b, err := r.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected serialized bytes", actual)
	var nilR *corejson.Result
	_, err2 := nilR.Serialize()
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	rErr := corejson.NewResult.Error(errors.New("e"))
	_, err3 := rErr.Serialize()
	actual = args.Map{"result": err3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C01_Result_SerializeSkipExistingIssues(t *testing.T) {
	// Arrange
	rErr := corejson.NewResult.Error(errors.New("e"))
	b, err := rErr.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{"result": b != nil || err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return nil,nil for issues", actual)
	r := corejson.NewResult.Any(42)
	b2, err2 := r.SerializeSkipExistingIssues()
	actual = args.Map{"result": err2 != nil || len(b2) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C01_Result_SerializeMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(42)

	// Act
	actual := args.Map{"result": len(r.SerializeMust()) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C01_Result_UnmarshalResult(t *testing.T) {
	r := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"x"`), TypeName: "test"})
	_, _ = r.UnmarshalResult()
}

func Test_C01_Result_JsonModel_JsonModelAny(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	_ = r.JsonModel()
	var nilR *corejson.Result
	m2 := nilR.JsonModel()

	// Act
	actual := args.Map{"result": m2.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	r2 := corejson.NewResult.Any("y")
	_ = r2.JsonModelAny()
}

func Test_C01_Result_Json_JsonPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	j := r.Json()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	actual = args.Map{"result": r.JsonPtr() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C01_Result_ParseInjectUsingJson(t *testing.T) {
	r := corejson.NewResult.Any("x")
	target := &corejson.Result{}
	_, _ = target.ParseInjectUsingJson(r.Ptr())
}

func Test_C01_Result_CloneError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.CloneError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be nil", actual)
	rErr := corejson.NewResult.Error(errors.New("e"))
	actual = args.Map{"result": rErr.CloneError() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have error", actual)
}

func Test_C01_Result_Ptr_NonPtr_ToPtr_ToNonPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	p := r.Ptr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ptr", actual)
	_ = p.NonPtr()
	var nilR *corejson.Result
	np2 := nilR.NonPtr()
	actual = args.Map{"result": np2.Error == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	r2 := corejson.NewResult.Any("y")
	_ = r2.ToPtr()
	_ = r2.ToNonPtr()
}

func Test_C01_Result_IsEqualPtr(t *testing.T) {
	// Arrange
	r1 := corejson.NewResult.AnyPtr("x")
	r2 := corejson.NewResult.AnyPtr("x")

	// Act
	actual := args.Map{"result": r1.IsEqualPtr(r2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.IsEqualPtr(nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": nilR.IsEqualPtr(r1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)
	actual = args.Map{"result": r1.IsEqualPtr(nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil should not be equal", actual)
	r3 := corejson.NewResult.AnyPtr("y")
	actual = args.Map{"result": r1.IsEqualPtr(r3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different should not be equal", actual)
}

func Test_C01_Result_IsEqual(t *testing.T) {
	// Arrange
	r1 := corejson.NewResult.Any("x")
	r2 := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r1.IsEqual(r2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
	r3 := corejson.NewResult.Any("y")
	actual = args.Map{"result": r1.IsEqual(r3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be equal", actual)
}

func Test_C01_Result_CombineErrorWithRefString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.CombineErrorWithRefString("ref1") != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for no error", actual)
	rErr := corejson.NewResult.Error(errors.New("e"))
	actual = args.Map{"result": rErr.CombineErrorWithRefString("ref1") == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected combined string", actual)
}

func Test_C01_Result_CombineErrorWithRefError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.CombineErrorWithRefError("ref") != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	rErr := corejson.NewResult.Error(errors.New("e"))
	actual = args.Map{"result": rErr.CombineErrorWithRefError("ref") == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C01_Result_BytesError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.BytesError() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.BytesError() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C01_Result_Dispose(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	r.Dispose()

	// Act
	actual := args.Map{"result": r.Bytes != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil bytes after dispose", actual)
	var nilR *corejson.Result
	nilR.Dispose()
}

func Test_C01_Result_Clone(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("test")
	c := r.Clone(false)

	// Act
	actual := args.Map{"result": c.JsonString() != r.JsonString()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "shallow clone mismatch", actual)
	c2 := r.Clone(true)
	actual = args.Map{"result": c2.JsonString() != r.JsonString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "deep clone mismatch", actual)
	empty := corejson.Result{}
	_ = empty.Clone(true)
}

func Test_C01_Result_CloneIf(t *testing.T) {
	r := corejson.NewResult.Any("x")
	_ = r.CloneIf(true, false)
	_ = r.CloneIf(false, false)
}

func Test_C01_Result_ClonePtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")
	c := r.ClonePtr(true)

	// Act
	actual := args.Map{"result": c == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil clone", actual)
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.ClonePtr(true) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil", actual)
}

func Test_C01_Result_InjectInto(t *testing.T) {
	r := corejson.NewResult.Any("x")
	target := &corejson.Result{}
	_ = r.InjectInto(target)
}

func Test_C01_Result_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.AsJsonContractsBinder() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C01_Result_AsJsoner(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.AsJsoner() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C01_Result_JsonParseSelfInject(t *testing.T) {
	r := corejson.NewResult.Any("x")
	target := corejson.NewResult.Any("y")
	_ = target.JsonParseSelfInject(r.Ptr())
}

func Test_C01_Result_AsJsonParseSelfInjector(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.AsJsonParseSelfInjector() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C01_Result_DeserializedFieldsToMap(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	_, _ = r.DeserializedFieldsToMap()
	var nilR *corejson.Result
	fm, err := nilR.DeserializedFieldsToMap()

	// Act
	actual := args.Map{"result": err != nil || len(fm) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_C01_Result_SafeDeserializedFieldsToMap(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	_ = r.SafeDeserializedFieldsToMap()
}

func Test_C01_Result_FieldsNames(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	_, _ = r.FieldsNames()
}

func Test_C01_Result_SafeFieldsNames(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1})
	_ = r.SafeFieldsNames()
}

func Test_C01_Result_BytesTypeName(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.BytesTypeName() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type name", actual)
	var nilR *corejson.Result
	actual = args.Map{"result": nilR.BytesTypeName() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_C01_Result_SafeBytesTypeName(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")

	// Act
	actual := args.Map{"result": r.SafeBytesTypeName() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type name", actual)
	emptyR := &corejson.Result{}
	actual = args.Map{"result": emptyR.SafeBytesTypeName() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for empty result", actual)
}

func Test_C01_Result_HandleError_NoPanic(t *testing.T) {
	r := corejson.NewResult.Any("x")
	r.HandleError()
}

func Test_C01_Result_MustBeSafe_NoPanic(t *testing.T) {
	r := corejson.NewResult.Any("x")
	r.MustBeSafe()
}

func Test_C01_Result_UnmarshalMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(42)
	var i int
	r.UnmarshalMust(&i)

	// Act
	actual := args.Map{"result": i != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}
