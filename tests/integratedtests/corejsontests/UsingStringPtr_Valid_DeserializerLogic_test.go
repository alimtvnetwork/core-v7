package corejsontests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── UsingStringPtr ──

func Test_UsingStringPtr_Valid(t *testing.T) {
	// Arrange
	s := `"hello"`
	var result string

	// Act
	err := corejson.Deserialize.UsingStringPtr(&s, &result)

	// Assert
	actual := args.Map{
		"val": result,
		"noErr": err == nil,
	}
	expected := args.Map{
		"val": "hello",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingStringPtr returns value -- valid string", actual)
}

func Test_UsingStringPtr_Nil(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.UsingStringPtr(nil, &result)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingStringPtr returns error -- nil ptr", actual)
}

// ── UsingError ──

func Test_UsingError_Valid(t *testing.T) {
	// Arrange
	errJson := errors.New(`"error_msg"`)
	var result string

	// Act
	err := corejson.Deserialize.UsingError(errJson, &result)

	// Assert
	actual := args.Map{
		"val": result,
		"noErr": err == nil,
	}
	expected := args.Map{
		"val": "error_msg",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingError returns value -- valid json error", actual)
}

func Test_UsingError_Nil(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.UsingError(nil, &result)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UsingError returns nil -- nil error", actual)
}

// ── UsingBytesPointer ──

func Test_UsingBytesPointer_Valid(t *testing.T) {
	// Arrange
	bytes := []byte(`"test"`)
	var result string

	// Act
	err := corejson.Deserialize.UsingBytesPointer(bytes, &result)

	// Assert
	actual := args.Map{
		"val": result,
		"noErr": err == nil,
	}
	expected := args.Map{
		"val": "test",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingBytesPointer returns value -- valid", actual)
}

func Test_UsingBytesPointer_Nil(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.UsingBytesPointer(nil, &result)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingBytesPointer returns error -- nil bytes", actual)
}

func Test_UsingBytesPointerMust_Panics(t *testing.T) {
	// Arrange
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		var result string
		corejson.Deserialize.UsingBytesPointerMust(nil, &result)
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "UsingBytesPointerMust panics -- nil bytes", actual)
}

// ── UsingBytesIf / UsingBytesPointerIf ──

func Test_UsingBytesIf_Skip(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"test"`), &result)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"empty": result == "",
	}
	expected := args.Map{
		"noErr": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingBytesIf skips -- false flag", actual)
}

func Test_UsingBytesIf_Process(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.UsingBytesIf(true, []byte(`"ok"`), &result)

	// Assert
	actual := args.Map{
		"val": result,
		"noErr": err == nil,
	}
	expected := args.Map{
		"val": "ok",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingBytesIf processes -- true flag", actual)
}

func Test_UsingBytesPointerIf_Skip(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &result)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"empty": result == "",
	}
	expected := args.Map{
		"noErr": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingBytesPointerIf skips -- false flag", actual)
}

// ── UsingSafeBytesMust ──

func Test_UsingSafeBytesMust_Empty(t *testing.T) {
	// Arrange
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		var result string
		corejson.Deserialize.UsingSafeBytesMust([]byte{}, &result)
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "UsingSafeBytesMust no panic -- empty bytes", actual)
}

func Test_UsingSafeBytesMust_Valid(t *testing.T) {
	// Arrange
	var result string

	// Act
	corejson.Deserialize.UsingSafeBytesMust([]byte(`"hello"`), &result)

	// Assert
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "UsingSafeBytesMust sets value -- valid bytes", actual)
}

// ── UsingStringOption / UsingStringIgnoreEmpty ──

func Test_UsingStringOption_IgnoreEmpty(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.UsingStringOption(true, "", &result)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"empty": result == "",
	}
	expected := args.Map{
		"noErr": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingStringOption skips -- empty+ignore", actual)
}

func Test_UsingStringIgnoreEmpty_Empty(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &result)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UsingStringIgnoreEmpty skips -- empty", actual)
}

func Test_UsingStringIgnoreEmpty_Valid(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.UsingStringIgnoreEmpty(`"hello"`, &result)

	// Assert
	actual := args.Map{
		"val": result,
		"noErr": err == nil,
	}
	expected := args.Map{
		"val": "hello",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingStringIgnoreEmpty parses -- valid", actual)
}

// ── MapAnyToPointer ──

func Test_MapAnyToPointer_Valid(t *testing.T) {
	// Arrange
	type simple struct {
		Name string `json:"Name"`
	}
	m := map[string]any{"Name": "Alice"}
	var result simple

	// Act
	err := corejson.Deserialize.MapAnyToPointer(false, m, &result)

	// Assert
	actual := args.Map{
		"name": result.Name,
		"noErr": err == nil,
	}
	expected := args.Map{
		"name": "Alice",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyToPointer returns struct -- valid map", actual)
}

func Test_MapAnyToPointer_SkipEmpty(t *testing.T) {
	// Arrange
	type simple struct{ Name string }
	var result simple

	// Act
	err := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &result)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"empty": result.Name == "",
	}
	expected := args.Map{
		"noErr": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyToPointer skips -- empty map+skipFlag", actual)
}

// ── AnyToFieldsMap ──

func Test_AnyToFieldsMap_Valid(t *testing.T) {
	// Arrange
	type simple struct {
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}

	// Act
	result, err := corejson.Deserialize.AnyToFieldsMap(simple{Name: "Bob", Age: 30})

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"name": result["Name"],
	}
	expected := args.Map{
		"noErr": true,
		"name": "Bob",
	}
	expected.ShouldBeEqual(t, 0, "AnyToFieldsMap returns map -- valid struct", actual)
}

// ── FromString / FromStringMust ──

func Test_FromString_Valid(t *testing.T) {
	// Arrange
	var result string

	// Act
	err := corejson.Deserialize.FromString(`"hi"`, &result)

	// Assert
	actual := args.Map{
		"val": result,
		"noErr": err == nil,
	}
	expected := args.Map{
		"val": "hi",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FromString returns value -- valid", actual)
}

func Test_FromStringMust_Panics(t *testing.T) {
	// Arrange
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		var result int
		corejson.Deserialize.FromStringMust("not-json", &result)
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "FromStringMust panics -- invalid json", actual)
}

// ── Serializer methods ──

func Test_Serialize_FromStringer(t *testing.T) {
	// Arrange
	stringer := fmt.Stringer(simpleStringer{val: "hello"})

	// Act
	result := corejson.Serialize.FromStringer(stringer)

	// Assert
	actual := args.Map{
		"hasBytes": result.Length() > 0,
		"noErr": !result.HasError(),
	}
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FromStringer returns result -- valid stringer", actual)
}

type simpleStringer struct{ val string }

func (s simpleStringer) String() string { return s.val }

func Test_Serialize_ToBytesSwallowErr_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result := corejson.Serialize.ToBytesSwallowErr("test")

	// Assert
	actual := args.Map{"hasBytes": len(result) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "ToBytesSwallowErr returns bytes -- valid", actual)
}

func Test_Serialize_ToSafeBytesSwallowErr_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result := corejson.Serialize.ToSafeBytesSwallowErr("test")

	// Assert
	actual := args.Map{"hasBytes": len(result) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "ToSafeBytesSwallowErr returns bytes -- valid", actual)
}

func Test_Serialize_ToBytesErr_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result, err := corejson.Serialize.ToBytesErr("test")

	// Assert
	actual := args.Map{
		"hasBytes": len(result) > 0,
		"noErr": err == nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToBytesErr returns bytes -- valid", actual)
}

func Test_Serialize_ToStringErr_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result, err := corejson.Serialize.ToStringErr("test")

	// Assert
	actual := args.Map{
		"hasContent": len(result) > 0,
		"noErr": err == nil,
	}
	expected := args.Map{
		"hasContent": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStringErr returns string -- valid", actual)
}

func Test_Serialize_ToPrettyStringErr_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange
	m := map[string]string{"a": "b"}

	// Act
	result, err := corejson.Serialize.ToPrettyStringErr(m)

	// Assert
	actual := args.Map{
		"hasContent": len(result) > 0,
		"noErr": err == nil,
	}
	expected := args.Map{
		"hasContent": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToPrettyStringErr returns string -- valid", actual)
}

func Test_Serialize_ToPrettyStringIncludingErr_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result := corejson.Serialize.ToPrettyStringIncludingErr(map[string]string{"a": "b"})

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "ToPrettyStringIncludingErr returns string -- valid", actual)
}

func Test_Serialize_Pretty_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result := corejson.Serialize.Pretty(map[string]string{"a": "b"})

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Pretty returns formatted string -- valid", actual)
}

func Test_Serialize_FromInteger64_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result := corejson.Serialize.FromInteger64(42)

	// Assert
	actual := args.Map{
		"hasBytes": result.Length() > 0,
		"noErr": !result.HasError(),
	}
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FromInteger64 returns result -- valid int", actual)
}

func Test_Serialize_FromBool_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result := corejson.Serialize.FromBool(true)

	// Assert
	actual := args.Map{"hasBytes": result.Length() > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "FromBool returns result -- true", actual)
}

func Test_Serialize_FromIntegers_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result := corejson.Serialize.FromIntegers([]int{1, 2, 3})

	// Assert
	actual := args.Map{"hasBytes": result.Length() > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "FromIntegers returns result -- slice", actual)
}

func Test_Serialize_FromBytes_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result := corejson.Serialize.FromBytes([]byte("test"))

	// Assert
	actual := args.Map{"hasBytes": result.Length() > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "FromBytes returns result -- bytes", actual)
}

func Test_Serialize_FromStringsSpread_FromUsingStringPtrValidD(t *testing.T) {
	// Arrange & Act
	result := corejson.Serialize.FromStringsSpread("a", "b")

	// Assert
	actual := args.Map{"hasBytes": result.Length() > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "FromStringsSpread returns result -- spread", actual)
}

// ── JsonString function ──

func Test_JsonString_Valid(t *testing.T) {
	// Arrange
	type simple struct {
		Name string `json:"Name"`
	}

	// Act
	result, err := corejson.JsonString(simple{Name: "Alice"})

	// Assert
	actual := args.Map{
		"hasContent": len(result) > 0,
		"noErr": err == nil,
	}
	expected := args.Map{
		"hasContent": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns string -- valid struct", actual)
}
