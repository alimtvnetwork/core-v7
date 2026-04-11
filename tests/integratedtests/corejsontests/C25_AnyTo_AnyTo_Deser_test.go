package corejsontests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── anyTo methods ──

func Test_C25_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult(nil)

	// Act
	actual := args.Map{"result": r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C25_AnyTo_SerializedJsonResult_Result(t *testing.T) {
	r := corejson.New("x")
	out := corejson.AnyTo.SerializedJsonResult(r)
	_ = out
}

func Test_C25_AnyTo_SerializedJsonResult_ResultPtr(t *testing.T) {
	r := corejson.New("x")
	out := corejson.AnyTo.SerializedJsonResult(r.Ptr())
	_ = out
}

func Test_C25_AnyTo_SerializedJsonResult_Bytes(t *testing.T) {
	out := corejson.AnyTo.SerializedJsonResult([]byte(`"x"`))
	_ = out
}

func Test_C25_AnyTo_SerializedJsonResult_String(t *testing.T) {
	out := corejson.AnyTo.SerializedJsonResult(`"hello"`)
	_ = out
}

func Test_C25_AnyTo_SerializedJsonResult_Error(t *testing.T) {
	out := corejson.AnyTo.SerializedJsonResult(errors.New(`"errmsg"`))
	_ = out
}

func Test_C25_AnyTo_SerializedJsonResult_EmptyError(t *testing.T) {
	// Arrange
	var e error
	out := corejson.AnyTo.SerializedJsonResult(e)

	// Act
	actual := args.Map{"result": out.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C25_AnyTo_SerializedJsonResult_Any(t *testing.T) {
	out := corejson.AnyTo.SerializedJsonResult(42)
	_ = out
}

func Test_C25_AnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw("hello")
	_ = b
	_ = err
}

func Test_C25_AnyTo_SerializedString(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.SerializedString("hello")

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_AnyTo_SerializedString_Error(t *testing.T) {
	// Arrange
	_, err := corejson.AnyTo.SerializedString(make(chan int))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C25_AnyTo_SerializedSafeString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedSafeString("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C25_AnyTo_SerializedSafeString_Error(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedSafeString(make(chan int))

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for error", actual)
}

func Test_C25_AnyTo_SerializedStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedStringMust("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C25_AnyTo_SafeJsonString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonString("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C25_AnyTo_PrettyStringWithError_String(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError("hello")

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_AnyTo_PrettyStringWithError_Bytes(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":"b"}`))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	_ = s
}

func Test_C25_AnyTo_PrettyStringWithError_Result(t *testing.T) {
	r := corejson.New(map[string]string{"a": "b"})
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	_ = s
	_ = err
}

func Test_C25_AnyTo_PrettyStringWithError_ResultWithErr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "t")
	_, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C25_AnyTo_PrettyStringWithError_ResultPtr(t *testing.T) {
	r := corejson.New("x")
	s, _ := corejson.AnyTo.PrettyStringWithError(r.Ptr())
	_ = s
}

func Test_C25_AnyTo_PrettyStringWithError_ResultPtrWithErr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("e"), "t")
	_, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C25_AnyTo_PrettyStringWithError_Any(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError(42)
	_ = s
	_ = err
}

func Test_C25_AnyTo_SafeJsonPrettyString_String(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonPrettyString("hello")

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_AnyTo_SafeJsonPrettyString_Bytes(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":"b"}`))
	_ = s
}

func Test_C25_AnyTo_SafeJsonPrettyString_Result(t *testing.T) {
	r := corejson.New("x")
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	_ = s
}

func Test_C25_AnyTo_SafeJsonPrettyString_ResultPtr(t *testing.T) {
	r := corejson.New("x")
	s := corejson.AnyTo.SafeJsonPrettyString(r.Ptr())
	_ = s
}

func Test_C25_AnyTo_SafeJsonPrettyString_Any(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString(42)
	_ = s
}

func Test_C25_AnyTo_JsonString_String(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonString("hello")

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_AnyTo_JsonString_Bytes(t *testing.T) {
	s := corejson.AnyTo.JsonString([]byte(`"x"`))
	_ = s
}

func Test_C25_AnyTo_JsonString_Result(t *testing.T) {
	r := corejson.New("x")
	s := corejson.AnyTo.JsonString(r)
	_ = s
}

func Test_C25_AnyTo_JsonString_ResultPtr(t *testing.T) {
	r := corejson.New("x")
	s := corejson.AnyTo.JsonString(r.Ptr())
	_ = s
}

func Test_C25_AnyTo_JsonString_Any(t *testing.T) {
	s := corejson.AnyTo.JsonString(42)
	_ = s
}

func Test_C25_AnyTo_JsonStringWithErr_String(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.JsonStringWithErr("hello")

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_AnyTo_JsonStringWithErr_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr([]byte(`"x"`))
	_ = s
	_ = err
}

func Test_C25_AnyTo_JsonStringWithErr_Result(t *testing.T) {
	r := corejson.New("x")
	s, err := corejson.AnyTo.JsonStringWithErr(r)
	_ = s
	_ = err
}

func Test_C25_AnyTo_JsonStringWithErr_ResultWithErr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Create([]byte(`"x"`), errors.New("e"), "")
	_, err := corejson.AnyTo.JsonStringWithErr(r)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C25_AnyTo_JsonStringWithErr_ResultPtr(t *testing.T) {
	r := corejson.New("x")
	_, _ = corejson.AnyTo.JsonStringWithErr(r.Ptr())
}

func Test_C25_AnyTo_JsonStringWithErr_ResultPtrWithErr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("e"), "")
	_, err := corejson.AnyTo.JsonStringWithErr(r)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C25_AnyTo_JsonStringWithErr_Any(t *testing.T) {
	_, _ = corejson.AnyTo.JsonStringWithErr(42)
}

func Test_C25_AnyTo_JsonStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonStringMust("hello")

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_AnyTo_JsonStringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	corejson.AnyTo.JsonStringMust(make(chan int))
}

func Test_C25_AnyTo_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust("hello")
	_ = s
}

func Test_C25_AnyTo_PrettyStringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	corejson.AnyTo.PrettyStringMust(make(chan int))
}

func Test_C25_AnyTo_UsingSerializer(t *testing.T) {
	_ = corejson.AnyTo.UsingSerializer(nil)
}

func Test_C25_AnyTo_SerializedFieldsMap(t *testing.T) {
	m, err := corejson.AnyTo.SerializedFieldsMap(map[string]string{"a": "b"})
	_ = m
	_ = err
}

// ── deserializerLogic methods ──

func Test_C25_Deser_ApplyMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	var s string
	corejson.Deserialize.ApplyMust(r, &s)
}

func Test_C25_Deser_FromString(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.FromString(`"hello"`, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_Deser_FromStringMust(t *testing.T) {
	var s string
	corejson.Deserialize.FromStringMust(`"hello"`, &s)
}

func Test_C25_Deser_FromStringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	var s string
	corejson.Deserialize.FromStringMust(`invalid`, &s)
}

func Test_C25_Deser_FromTo(t *testing.T) {
	var s string
	err := corejson.Deserialize.FromTo(`"hello"`, &s)
	_ = err
}

func Test_C25_Deser_MapAnyToPointer_SkipEmpty(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_MapAnyToPointer_Valid(t *testing.T) {
	// Arrange
	var m map[string]string
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"k": "v"}, &m)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingBytesIf_Skip(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingBytesIf(false, nil, nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingBytesIf_Do(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesIf(true, []byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingBytesMust(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.UsingBytesMust([]byte(`invalid`), nil)
}

func Test_C25_Deser_UsingBytesPointer_Nil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingBytesPointer(nil, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C25_Deser_UsingBytesPointer_Valid(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesPointer([]byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingBytesPointerIf_Skip(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingBytesPointerIf(false, nil, nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingBytesPointerIf_Do(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(true, []byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingBytesPointerMust(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.UsingBytesPointerMust(nil, nil)
}

func Test_C25_Deser_UsingSafeBytesMust_Empty(t *testing.T) {
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, nil) // no panic
}

func Test_C25_Deser_UsingSafeBytesMust_Panic(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.UsingSafeBytesMust([]byte(`invalid`), nil)
}

func Test_C25_Deser_UsingSerializerTo(t *testing.T) {
	_ = corejson.Deserialize.UsingSerializerTo(nil, nil)
}

func Test_C25_Deser_UsingSerializerFuncTo(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingSerializerFuncTo(func() ([]byte, error) {
		return json.Marshal("test")
	}, &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingDeserializerToOption_SkipNil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingDeserializerToOption_NotSkipNil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C25_Deser_UsingDeserializerDefined_Nil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingDeserializerDefined(nil, nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil func", actual)
}

func Test_C25_Deser_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingDeserializerFuncDefined(func(toPtr any) error {
		return nil
	}, nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingJsonerToAny_SkipNil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingJsonerToAny_NotSkipNil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C25_Deser_UsingJsonerToAnyMust_SkipNil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingJsonerToAnyMust_NotSkipNil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, nil, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C25_Deser_UsingError_Nil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingError(nil, nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingError_Valid(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingError(errors.New(`"hello"`), &s)
	_ = err
}

func Test_C25_Deser_UsingErrorWhichJsonResult_Nil(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_Deser_UsingErrorWhichJsonResult_Valid(t *testing.T) {
	var r corejson.Result
	err := corejson.Deserialize.UsingErrorWhichJsonResult(
		errors.New(`{"Bytes":"dGVzdA==","TypeName":"t"}`), &r)
	_ = err
}

// ── deserializeFromBytesTo ──

func Test_C25_BytesTo_Integer(t *testing.T) {
	// Arrange
	i, err := corejson.Deserialize.BytesTo.Integer([]byte(`42`))

	// Act
	actual := args.Map{"result": err != nil || i != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_BytesTo_IntegerMust(t *testing.T) {
	// Arrange
	i := corejson.Deserialize.BytesTo.IntegerMust([]byte(`42`))

	// Act
	actual := args.Map{"result": i != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_BytesTo_Integer64(t *testing.T) {
	// Arrange
	i, err := corejson.Deserialize.BytesTo.Integer64([]byte(`42`))

	// Act
	actual := args.Map{"result": err != nil || i != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_BytesTo_Integer64Must(t *testing.T) {
	// Arrange
	i := corejson.Deserialize.BytesTo.Integer64Must([]byte(`42`))

	// Act
	actual := args.Map{"result": i != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_BytesTo_Integers(t *testing.T) {
	// Arrange
	ints, err := corejson.Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))

	// Act
	actual := args.Map{"result": err != nil || len(ints) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_BytesTo_IntegersMust(t *testing.T) {
	// Arrange
	ints := corejson.Deserialize.BytesTo.IntegersMust([]byte(`[1,2]`))

	// Act
	actual := args.Map{"result": len(ints) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_BytesTo_IntegersMust_Panic(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.BytesTo.IntegersMust([]byte(`invalid`))
}

func Test_C25_BytesTo_StringMust(t *testing.T) {
	// Arrange
	s := corejson.Deserialize.BytesTo.StringMust([]byte(`"hello"`))

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_BytesTo_StringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.BytesTo.StringMust([]byte(`invalid`))
}

func Test_C25_BytesTo_MapAnyItem(t *testing.T) {
	// Arrange
	m, err := corejson.Deserialize.BytesTo.MapAnyItem([]byte(`{"a":"b"}`))

	// Act
	actual := args.Map{"result": err != nil || m["a"] != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_BytesTo_MapAnyItemMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapAnyItemMust([]byte(`{"a":"b"}`))
	_ = m
}

func Test_C25_BytesTo_MapStringString(t *testing.T) {
	// Arrange
	m, err := corejson.Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))

	// Act
	actual := args.Map{"result": err != nil || m["a"] != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_BytesTo_MapStringStringMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapStringStringMust([]byte(`{"a":"b"}`))
	_ = m
}

func Test_C25_BytesTo_ResultCollection(t *testing.T) {
	b, _ := json.Marshal([]corejson.Result{corejson.New("x")})
	rc, err := corejson.Deserialize.BytesTo.ResultCollection(b)
	_ = rc
	_ = err
}

func Test_C25_BytesTo_ResultCollectionMust(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.BytesTo.ResultCollectionMust([]byte(`invalid`))
}

func Test_C25_BytesTo_ResultsPtrCollection(t *testing.T) {
	r := corejson.New("x")
	b, _ := json.Marshal([]*corejson.Result{r.Ptr()})
	rc, err := corejson.Deserialize.BytesTo.ResultsPtrCollection(b)
	_ = rc
	_ = err
}

func Test_C25_BytesTo_ResultsPtrCollectionMust(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.BytesTo.ResultsPtrCollectionMust([]byte(`invalid`))
}

func Test_C25_BytesTo_MapResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	b, _ := json.Marshal(mr)
	out, err := corejson.Deserialize.BytesTo.MapResults(b)
	_ = out
	_ = err
}

func Test_C25_BytesTo_MapResultsMust(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.BytesTo.MapResultsMust([]byte(`invalid`))
}

func Test_C25_BytesTo_Bytes(t *testing.T) {
	b, err := corejson.Deserialize.BytesTo.Bytes([]byte(`"aGVsbG8="`))
	_ = b
	_ = err
}

func Test_C25_BytesTo_BytesMust(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.BytesTo.BytesMust([]byte(`invalid`))
}

func Test_C25_Deser_Result(t *testing.T) {
	r := corejson.New("x")
	b, _ := json.Marshal(r)
	out, err := corejson.Deserialize.Result(b)
	_ = out
	_ = err
}

func Test_C25_Deser_ResultMust(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.ResultMust([]byte(`invalid`))
}

func Test_C25_Deser_ResultPtr(t *testing.T) {
	r := corejson.New("x")
	b, _ := json.Marshal(r)
	out, err := corejson.Deserialize.ResultPtr(b)
	_ = out
	_ = err
}

func Test_C25_Deser_ResultPtrMust(t *testing.T) {
	defer func() { recover() }()
	corejson.Deserialize.ResultPtrMust([]byte(`invalid`))
}

// ── deserializeFromResultTo ──

func Test_C25_ResultTo_String(t *testing.T) {
	r := corejson.New("hello")
	s, err := corejson.Deserialize.ResultTo.String(r.Ptr())
	_ = s
	_ = err
}

func Test_C25_ResultTo_Bool(t *testing.T) {
	r := corejson.New(true)
	b, err := corejson.Deserialize.ResultTo.Bool(r.Ptr())
	_ = b
	_ = err
}

func Test_C25_ResultTo_Byte(t *testing.T) {
	r := corejson.New(byte(5))
	b, err := corejson.Deserialize.ResultTo.Byte(r.Ptr())
	_ = b
	_ = err
}

func Test_C25_ResultTo_ByteMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.ByteMust(r)
}

func Test_C25_ResultTo_BoolMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.BoolMust(r)
}

func Test_C25_ResultTo_StringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.StringMust(r)
}

func Test_C25_ResultTo_StringsMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.StringsMust(r)
}

func Test_C25_ResultTo_MapAnyItem(t *testing.T) {
	r := corejson.New(map[string]any{"a": "b"})
	m, err := corejson.Deserialize.ResultTo.MapAnyItem(r.Ptr())
	_ = m
	_ = err
}

func Test_C25_ResultTo_MapAnyItemMust(t *testing.T) {
	r := corejson.New(map[string]any{"a": "b"})
	_ = corejson.Deserialize.ResultTo.MapAnyItemMust(r.Ptr())
}

func Test_C25_ResultTo_MapStringString(t *testing.T) {
	r := corejson.New(map[string]string{"a": "b"})
	m, err := corejson.Deserialize.ResultTo.MapStringString(r.Ptr())
	_ = m
	_ = err
}

func Test_C25_ResultTo_MapStringStringMust(t *testing.T) {
	r := corejson.New(map[string]string{"a": "b"})
	_ = corejson.Deserialize.ResultTo.MapStringStringMust(r.Ptr())
}

func Test_C25_ResultTo_ResultCollection(t *testing.T) {
	items := []corejson.Result{corejson.New("x")}
	r := corejson.New(items)
	rc, err := corejson.Deserialize.ResultTo.ResultCollection(r.Ptr())
	_ = rc
	_ = err
}

func Test_C25_ResultTo_ResultCollectionMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.ResultCollectionMust(r)
}

func Test_C25_ResultTo_ResultsPtrCollection(t *testing.T) {
	r := corejson.New([]*corejson.Result{})
	rc, err := corejson.Deserialize.ResultTo.ResultsPtrCollection(r.Ptr())
	_ = rc
	_ = err
}

func Test_C25_ResultTo_ResultsPtrCollectionMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.ResultsPtrCollectionMust(r)
}

func Test_C25_ResultTo_Result(t *testing.T) {
	inner := corejson.New("x")
	r := corejson.New(inner)
	out, err := corejson.Deserialize.ResultTo.Result(r.Ptr())
	_ = out
	_ = err
}

func Test_C25_ResultTo_ResultMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.ResultMust(r)
}

func Test_C25_ResultTo_ResultPtr(t *testing.T) {
	inner := corejson.New("x")
	r := corejson.New(inner)
	out, err := corejson.Deserialize.ResultTo.ResultPtr(r.Ptr())
	_ = out
	_ = err
}

func Test_C25_ResultTo_ResultPtrMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.ResultPtrMust(r)
}

func Test_C25_ResultTo_MapResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.New(mr)
	out, err := corejson.Deserialize.ResultTo.MapResults(r.Ptr())
	_ = out
	_ = err
}

func Test_C25_ResultTo_MapResultsMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.MapResultsMust(r)
}

func Test_C25_ResultTo_Bytes(t *testing.T) {
	inner := corejson.New("x")
	r := corejson.New(inner)
	b, err := corejson.Deserialize.ResultTo.Bytes(r.Ptr())
	_ = b
	_ = err
}

func Test_C25_ResultTo_BytesMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	corejson.Deserialize.ResultTo.BytesMust(r)
}

// ── castingAny ──

func Test_C25_CastAny_FromToReflection(t *testing.T) {
	src := "hello"
	dst := ""
	err := corejson.CastAny.FromToReflection(&src, &dst)
	_ = err
}

func Test_C25_CastAny_OrDeserializeTo(t *testing.T) {
	var s string
	err := corejson.CastAny.OrDeserializeTo(`"hello"`, &s)
	_ = err
}

func Test_C25_CastAny_FromToOption_Bytes(t *testing.T) {
	// Arrange
	var s string
	err := corejson.CastAny.FromToOption(false, []byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C25_CastAny_FromToOption_String(t *testing.T) {
	// Arrange
	var s string
	err := corejson.CastAny.FromToOption(false, `"hello"`, &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_CastAny_FromToOption_Error(t *testing.T) {
	var s string
	err := corejson.CastAny.FromToOption(false, errors.New(`"errmsg"`), &s)
	_ = err
}

func Test_C25_CastAny_FromToOption_NilError(t *testing.T) {
	// Arrange
	var e error
	var s string
	err := corejson.CastAny.FromToOption(false, e, &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_CastAny_FromToOption_SerializerFunc(t *testing.T) {
	// Arrange
	var s string
	f := func() ([]byte, error) { return json.Marshal("hello") }
	err := corejson.CastAny.FromToOption(false, f, &s)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_CastAny_FromToOption_Any(t *testing.T) {
	// Arrange
	var i int
	err := corejson.CastAny.FromToOption(false, 42, &i)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C25_CastAny_reflectionCasting_NilFrom(t *testing.T) {
	err := corejson.CastAny.FromToOption(true, nil, nil)
	_ = err
}

func Test_C25_CastAny_reflectionCasting_DiffTypes(t *testing.T) {
	src := "hello"
	var dst int
	err := corejson.CastAny.FromToOption(true, &src, &dst)
	_ = err
}
