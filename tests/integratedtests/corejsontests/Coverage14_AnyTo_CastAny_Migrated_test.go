package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Migrated from Coverage06_AnyTo_CastAny_test.go and Coverage11 ──

func Test_C06_AnyTo_SerializedJsonResult(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("x")
	jr := corejson.AnyTo.SerializedJsonResult(r)

	// Act
	actual := args.Map{"hasError": jr.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)

	rp := corejson.NewResult.AnyPtr("x")
	jr2 := corejson.AnyTo.SerializedJsonResult(rp)
	actual2 := args.Map{"hasError": jr2.HasError()}
		expected2 := args.Map{"hasError": false}
		expected2.ShouldBeEqual(t, 0, "result2 has no error", actual2)

	jr3 := corejson.AnyTo.SerializedJsonResult([]byte(`"x"`))
	actual3 := args.Map{"hasError": jr3.HasError()}
		expected3 := args.Map{"hasError": false}
		expected3.ShouldBeEqual(t, 0, "result3 has no error", actual3)

	jr4 := corejson.AnyTo.SerializedJsonResult("hello")
	_ = jr4

	jr5 := corejson.AnyTo.SerializedJsonResult(nil)
	actual := args.Map{"result": jr5.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)

	jr6 := corejson.AnyTo.SerializedJsonResult(42)
	actual6 := args.Map{"hasError": jr6.HasError()}
		expected6 := args.Map{"hasError": false}
		expected6.ShouldBeEqual(t, 0, "result6 has no error", actual6)

	jr7 := corejson.AnyTo.SerializedJsonResult(errors.New("oops"))
	actual := args.Map{"result": jr7 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C06_AnyTo_SerializedRaw(t *testing.T) {
	// Arrange
	b, err := corejson.AnyTo.SerializedRaw("hello")

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_AnyTo_SerializedString(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.SerializedString("hello")

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	_, err2 := corejson.AnyTo.SerializedString(nil)
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C06_AnyTo_SerializedSafeString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedSafeString("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	s2 := corejson.AnyTo.SerializedSafeString(nil)
	actual := args.Map{"result": s2 != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C06_AnyTo_SerializedStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedStringMust("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C06_AnyTo_SafeJsonString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonString("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C06_AnyTo_PrettyStringWithError(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError("hello")

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	s2, err2 := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	actual := args.Map{"result": err2 != nil || s2 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	r := corejson.NewResult.Any(42)
	s3, err3 := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"result": err3 != nil || s3 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	rp := corejson.NewResult.AnyPtr(42)
	s4, err4 := corejson.AnyTo.PrettyStringWithError(rp)
	actual := args.Map{"result": err4 != nil || s4 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)

	re := corejson.Result{Error: errors.New("e")}
	_, err5 := corejson.AnyTo.PrettyStringWithError(re)
	actual := args.Map{"result": err5 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	rep := &corejson.Result{Error: errors.New("e")}
	_, err6 := corejson.AnyTo.PrettyStringWithError(rep)
	actual := args.Map{"result": err6 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	s5, err5b := corejson.AnyTo.PrettyStringWithError(map[string]int{"a": 1})
	actual := args.Map{"result": err5b != nil || s5 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_AnyTo_SafeJsonPrettyString(t *testing.T) {
	_ = corejson.AnyTo.SafeJsonPrettyString("hello")
	_ = corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	_ = corejson.AnyTo.SafeJsonPrettyString(corejson.NewResult.Any(1))
	_ = corejson.AnyTo.SafeJsonPrettyString(corejson.NewResult.AnyPtr(1))
	_ = corejson.AnyTo.SafeJsonPrettyString(42)
}

func Test_C06_AnyTo_JsonString(t *testing.T) {
	_ = corejson.AnyTo.JsonString("hello")
	_ = corejson.AnyTo.JsonString([]byte("test"))
	_ = corejson.AnyTo.JsonString(corejson.NewResult.Any(1))
	_ = corejson.AnyTo.JsonString(corejson.NewResult.AnyPtr(1))
	_ = corejson.AnyTo.JsonString(42)
}

func Test_C06_AnyTo_JsonStringWithErr(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.JsonStringWithErr("hello")

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	_, _ = corejson.AnyTo.JsonStringWithErr([]byte("test"))
	_, _ = corejson.AnyTo.JsonStringWithErr(corejson.NewResult.Any(1))
	_, _ = corejson.AnyTo.JsonStringWithErr(corejson.NewResult.AnyPtr(1))
	_, _ = corejson.AnyTo.JsonStringWithErr(42)

	_, err2 := corejson.AnyTo.JsonStringWithErr(corejson.Result{Error: errors.New("e")})
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	_, err3 := corejson.AnyTo.JsonStringWithErr(&corejson.Result{Error: errors.New("e")})
	actual := args.Map{"result": err3 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C06_AnyTo_JsonStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonStringMust("hello")

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_AnyTo_PrettyStringMust(t *testing.T) {
	_ = corejson.AnyTo.PrettyStringMust("hello")
}

func Test_C06_AnyTo_SerializedFieldsMap(t *testing.T) {
	_, _ = corejson.AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
}

func Test_C06_CastAny_FromToDefault(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.FromToDefault([]byte(`"hello"`), &out)

	// Act
	actual := args.Map{"result": err != nil || out != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_CastAny_FromToOption_Bytes(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.FromToOption(false, []byte(`"hello"`), &out)

	// Act
	actual := args.Map{"result": err != nil || out != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_CastAny_FromToOption_String(t *testing.T) {
	// Arrange
	var out int
	err := corejson.CastAny.FromToOption(false, "42", &out)

	// Act
	actual := args.Map{"result": err != nil || out != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_CastAny_FromToOption_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	var out string
	_ = corejson.CastAny.FromToOption(false, r, &out)
}

func Test_C06_CastAny_FromToOption_ResultPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var out string
	_ = corejson.CastAny.FromToOption(false, r, &out)
}

func Test_C06_CastAny_FromToOption_SerializerFunc(t *testing.T) {
	// Arrange
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var out string
	err := corejson.CastAny.FromToOption(false, fn, &out)

	// Act
	actual := args.Map{"result": err != nil || out != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_CastAny_FromToOption_AnyFallback(t *testing.T) {
	// Arrange
	type simple struct{ Name string }
	src := simple{Name: "test"}
	var dst simple
	err := corejson.CastAny.FromToOption(false, src, &dst)

	// Act
	actual := args.Map{"result": err != nil || dst.Name != "test"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_CastAny_FromToOption_WithReflection(t *testing.T) {
	var out string
	_ = corejson.CastAny.FromToOption(true, "hello", &out)
}

func Test_C06_CastAny_OrDeserializeTo(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.OrDeserializeTo([]byte(`"hi"`), &out)

	// Act
	actual := args.Map{"result": err != nil || out != "hi"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_CastAny_FromToReflection(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.FromToReflection([]byte(`"hello"`), &out)

	// Act
	actual := args.Map{"result": err != nil || out != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

type testSerializer14 struct{}
func (testSerializer14) Serialize() ([]byte, error) { return []byte(`"x"`), nil }

func Test_C06_AnyTo_UsingSerializer_Alt(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.UsingSerializer(testSerializer14{})

	// Act
	actual := args.Map{"result": r == nil || r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C06_AnyTo_PrettyStringMust_Map(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.PrettyStringMust(map[string]string{"a": "1"})

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}
