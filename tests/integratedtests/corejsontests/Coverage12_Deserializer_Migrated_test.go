package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Migrated from Coverage03_Deserializer_test.go ──

func Test_C03_Deserializer_UsingBytes(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingBytes([]byte(`"hello"`), &out)

	// Act
	actual := args.Map{"result": err != nil || out != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	err2 := corejson.Deserialize.UsingBytes([]byte(`invalid`), &out)
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C03_Deserializer_UsingString(t *testing.T) {
	// Arrange
	var out int
	err := corejson.Deserialize.UsingString("42", &out)

	// Act
	actual := args.Map{"result": err != nil || out != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_FromString(t *testing.T) {
	// Arrange
	var out int
	err := corejson.Deserialize.FromString("42", &out)

	// Act
	actual := args.Map{"result": err != nil || out != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_FromStringMust(t *testing.T) {
	// Arrange
	var out int
	corejson.Deserialize.FromStringMust("42", &out)

	// Act
	actual := args.Map{"result": out != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_UsingStringPtr(t *testing.T) {
	// Arrange
	s := `"hello"`
	var out string
	err := corejson.Deserialize.UsingStringPtr(&s, &out)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	err2 := corejson.Deserialize.UsingStringPtr(nil, &out)
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C03_Deserializer_UsingError(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingError(nil, nil)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil error", actual)
	var out string
	err2 := corejson.Deserialize.UsingError(errors.New(`"hello"`), &out)
	actual := args.Map{"result": err2}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err2", actual)
}

func Test_C03_Deserializer_UsingStringOption(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingStringOption(true, "", &out)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty string skip", actual)
	err2 := corejson.Deserialize.UsingStringOption(false, `"x"`, &out)
	actual := args.Map{"result": err2}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err2", actual)
}

func Test_C03_Deserializer_UsingStringIgnoreEmpty(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &out)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C03_Deserializer_UsingBytesPointer(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingBytesPointer(nil, &out)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	err2 := corejson.Deserialize.UsingBytesPointer([]byte(`"x"`), &out)
	actual := args.Map{"result": err2}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err2", actual)
}

func Test_C03_Deserializer_UsingBytesPointerMust(t *testing.T) {
	// Arrange
	var out string
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"x"`), &out)

	// Act
	actual := args.Map{"result": out != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_UsingBytesMust(t *testing.T) {
	// Arrange
	var out int
	corejson.Deserialize.UsingBytesMust([]byte("42"), &out)

	// Act
	actual := args.Map{"result": out != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_UsingSafeBytesMust(t *testing.T) {
	// Arrange
	var out int
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &out)
	corejson.Deserialize.UsingSafeBytesMust([]byte("42"), &out)

	// Act
	actual := args.Map{"result": out != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_UsingBytesIf(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &out)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil when skip", actual)
	err2 := corejson.Deserialize.UsingBytesIf(true, []byte(`"x"`), &out)
	actual := args.Map{"result": err2}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err2", actual)
}

func Test_C03_Deserializer_UsingBytesPointerIf(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &out)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil when skip", actual)
}

func Test_C03_Deserializer_Apply(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")
	var out string
	err := corejson.Deserialize.Apply(r.Ptr(), &out)

	// Act
	actual := args.Map{"result": err != nil || out != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_ApplyMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any(42)
	var out int
	corejson.Deserialize.ApplyMust(r.Ptr(), &out)

	// Act
	actual := args.Map{"result": out != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_UsingResult(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hi")
	var out string
	err := corejson.Deserialize.UsingResult(r.Ptr(), &out)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C03_Deserializer_MapAnyToPointer(t *testing.T) {
	// Arrange
	m := map[string]any{"key": "val"}
	var out map[string]any
	err := corejson.Deserialize.MapAnyToPointer(false, m, &out)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	err2 := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &out)
	actual := args.Map{"result": err2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty skip", actual)
}

func Test_C03_Deserializer_AnyToFieldsMap(t *testing.T) {
	_, _ = corejson.Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
}

func Test_C03_Deserializer_BytesTo_Strings(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw([]string{"a", "b"})
	lines, err := corejson.Deserialize.BytesTo.Strings(b)

	// Act
	actual := args.Map{"result": err != nil || len(lines) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_StringsMust(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw([]string{"a"})
	lines := corejson.Deserialize.BytesTo.StringsMust(b)

	// Act
	actual := args.Map{"result": len(lines) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_String(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw("hello")
	s, err := corejson.Deserialize.BytesTo.String(b)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_StringMust(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw("x")
	s := corejson.Deserialize.BytesTo.StringMust(b)

	// Act
	actual := args.Map{"result": s != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_Integer(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(42)
	i, err := corejson.Deserialize.BytesTo.Integer(b)

	// Act
	actual := args.Map{"result": err != nil || i != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_IntegerMust(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(42)
	i := corejson.Deserialize.BytesTo.IntegerMust(b)

	// Act
	actual := args.Map{"result": i != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_Integer64(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(64)
	i, err := corejson.Deserialize.BytesTo.Integer64(b)

	// Act
	actual := args.Map{"result": err != nil || i != 64}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_Integer64Must(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(64)
	i := corejson.Deserialize.BytesTo.Integer64Must(b)

	// Act
	actual := args.Map{"result": i != 64}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_Integers(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw([]int{1, 2})
	ints, err := corejson.Deserialize.BytesTo.Integers(b)

	// Act
	actual := args.Map{"result": err != nil || len(ints) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_IntegersMust(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw([]int{1})
	ints := corejson.Deserialize.BytesTo.IntegersMust(b)

	// Act
	actual := args.Map{"result": len(ints) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_Bool(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(true)
	v, err := corejson.Deserialize.BytesTo.Bool(b)

	// Act
	actual := args.Map{"result": err != nil || !v}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_BoolMust(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(false)
	v := corejson.Deserialize.BytesTo.BoolMust(b)

	// Act
	actual := args.Map{"result": v}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_MapAnyItem(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(map[string]any{"k": "v"})
	m, err := corejson.Deserialize.BytesTo.MapAnyItem(b)

	// Act
	actual := args.Map{"result": err != nil || len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_MapAnyItemMust(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(map[string]any{"k": "v"})
	m := corejson.Deserialize.BytesTo.MapAnyItemMust(b)

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_MapStringString(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(map[string]string{"k": "v"})
	m, err := corejson.Deserialize.BytesTo.MapStringString(b)

	// Act
	actual := args.Map{"result": err != nil || m["k"] != "v"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_MapStringStringMust(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(map[string]string{"k": "v"})
	m := corejson.Deserialize.BytesTo.MapStringStringMust(b)

	// Act
	actual := args.Map{"result": m["k"] != "v"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_BytesTo_Bytes(t *testing.T) {
	input := []byte(`"aGVsbG8="`)
	_, _ = corejson.Deserialize.BytesTo.Bytes(input)
}

func Test_C03_Deserializer_ResultTo_String(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	s, err := corejson.Deserialize.ResultTo.String(r)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_ResultTo_StringMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("x")
	s := corejson.Deserialize.ResultTo.StringMust(r)

	// Act
	actual := args.Map{"result": s != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_ResultTo_Bool(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(true)
	v, err := corejson.Deserialize.ResultTo.Bool(r)

	// Act
	actual := args.Map{"result": err != nil || !v}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_ResultTo_BoolMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(true)
	v := corejson.Deserialize.ResultTo.BoolMust(r)

	// Act
	actual := args.Map{"result": v}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_ResultTo_Byte(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(byte(65))
	_, err := corejson.Deserialize.ResultTo.Byte(r)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C03_Deserializer_ResultTo_MapAnyItem(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(map[string]any{"k": "v"})
	m, err := corejson.Deserialize.ResultTo.MapAnyItem(r)

	// Act
	actual := args.Map{"result": err != nil || len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_ResultTo_MapAnyItemMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(map[string]any{"k": "v"})
	m := corejson.Deserialize.ResultTo.MapAnyItemMust(r)

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_ResultTo_MapStringString(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(map[string]string{"k": "v"})
	m, err := corejson.Deserialize.ResultTo.MapStringString(r)

	// Act
	actual := args.Map{"result": err != nil || m["k"] != "v"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_ResultTo_MapStringStringMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(map[string]string{"k": "v"})
	m := corejson.Deserialize.ResultTo.MapStringStringMust(r)

	// Act
	actual := args.Map{"result": m["k"] != "v"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C03_Deserializer_ResultTo_StringsMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr([]string{"a", "b"})
	lines := corejson.Deserialize.ResultTo.StringsMust(r)

	// Act
	actual := args.Map{"result": len(lines) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// ── Additional Deserializer methods from Coverage11, 15 ──

func Test_C11_Deserialize_UsingErrorWhichJsonResult(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &struct{}{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C11_Deserialize_FromTo(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.FromTo([]byte(`"hello"`), &out)

	// Act
	actual := args.Map{"result": err != nil || out != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C11_Deserialize_UsingDeserializerToOption(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &struct{}{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err2 := corejson.Deserialize.UsingDeserializerToOption(false, nil, &struct{}{})
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C11_Deserialize_UsingDeserializerDefined(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &struct{}{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C11_Deserialize_UsingDeserializerFuncDefined(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &struct{}{})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err2 := corejson.Deserialize.UsingDeserializerFuncDefined(func(toPtr any) error { return nil }, &struct{}{})
	actual := args.Map{"result": err2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C11_Deserialize_UsingJsonerToAny(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &struct{}{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err2 := corejson.Deserialize.UsingJsonerToAny(false, nil, &struct{}{})
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C11_Deserialize_UsingJsonerToAnyMust(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &struct{}{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C15_Deserialize_UsingSerializerFuncTo(t *testing.T) {
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

func Test_C15_Deserialize_ResultTo_ByteMust(t *testing.T) {
	// Arrange
	r := corejson.Serialize.Apply(byte(65))
	b := corejson.Deserialize.ResultTo.ByteMust(r)

	// Act
	actual := args.Map{"result": b != 65}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C15_Deserialize_BytesTo_BytesMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.BytesMust([]byte(`"aGVsbG8="`))
}
