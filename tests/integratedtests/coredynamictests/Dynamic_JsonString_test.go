package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Dynamic — DynamicJson.go ──

func Test_Dynamic_JsonString_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	jsonStr, err := d.JsonString()

	// Act
	actual := args.Map{
		"notEmpty": jsonStr != "",
		"noErr":    err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonString", actual)
}

func Test_Dynamic_JsonBytes_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(42)
	jsonBytes, err := d.JsonBytes()

	// Act
	actual := args.Map{
		"notEmpty": len(jsonBytes) > 0,
		"noErr":    err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonBytes", actual)
}

func Test_Dynamic_JsonBytesPtr_Null_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(nil)
	jsonBytes, err := d.JsonBytesPtr()

	// Act
	actual := args.Map{
		"emptyBytes": len(jsonBytes) == 0,
		"noErr":      err == nil,
	}

	// Assert
	expected := args.Map{
		"emptyBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonBytesPtr null", actual)
}

func Test_Dynamic_ValueMarshal_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("test")
	data, err := d.ValueMarshal()

	// Act
	actual := args.Map{
		"notEmpty": len(data) > 0,
		"noErr":    err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ValueMarshal", actual)
}

func Test_Dynamic_JsonModel_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("model")

	// Act
	actual := args.Map{
		"model":    d.JsonModel(),
		"modelAny": d.JsonModelAny(),
	}

	// Assert
	expected := args.Map{
		"model": "model", "modelAny": "model",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonModel", actual)
}

func Test_Dynamic_Json_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	j := d.Json()
	jp := d.JsonPtr()
	jsonNotEmpty := j.JsonString() != ""

	// Act
	actual := args.Map{
		"jsonNotEmpty": jsonNotEmpty,
		"jsonPtrNN":    jp != nil,
	}

	// Assert
	expected := args.Map{
		"jsonNotEmpty": jsonNotEmpty,
		"jsonPtrNN": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Json/JsonPtr", actual)
}

func Test_Dynamic_MarshalJSON_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("test")
	data, err := d.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"notEmpty": len(data) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- MarshalJSON", actual)
}

// ── Dynamic — DynamicStatus.go ──

func Test_DynamicStatus_Invalid_FromDynamicJsonString(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("err msg")
	dsNoMsg := coredynamic.InvalidDynamicStatusNoMessage()

	// Act
	actual := args.Map{
		"message":     ds.Message,
		"noMsgEmpty":  dsNoMsg.Message == "",
		"isNull":      ds.IsNull(),
	}

	// Assert
	expected := args.Map{
		"message": "err msg", "noMsgEmpty": true, "isNull": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns error -- Invalid", actual)
}

func Test_DynamicStatus_Clone_FromDynamicJsonString(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("msg")
	cloned := ds.Clone()
	clonedPtr := ds.ClonePtr()
	var nilDS *coredynamic.DynamicStatus

	// Act
	actual := args.Map{
		"clonedMsg":    cloned.Message,
		"clonedPtrMsg": clonedPtr.Message,
		"nilClone":     nilDS.ClonePtr() == nil,
	}

	// Assert
	expected := args.Map{
		"clonedMsg": "msg", "clonedPtrMsg": "msg", "nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- Clone", actual)
}

// ── Dynamic — ValueStatus ──

func Test_ValueStatus_Invalid_FromDynamicJsonString(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatus("err")
	vsNoMsg := coredynamic.InvalidValueStatusNoMessage()

	// Act
	actual := args.Map{
		"isValid":      vs.IsValid,
		"message":      vs.Message,
		"noMsgMessage": vsNoMsg.Message,
	}

	// Assert
	expected := args.Map{
		"isValid": false, "message": "err", "noMsgMessage": "",
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns error -- Invalid", actual)
}

// ── Dynamic — SafeTypeName ──

func Test_SafeTypeName_FromDynamicJsonString(t *testing.T) {
	// Act
	actual := args.Map{
		"stringType": coredynamic.SafeTypeName("hello"),
		"intType":    coredynamic.SafeTypeName(42),
		"nilType":    coredynamic.SafeTypeName(nil),
	}

	// Assert
	expected := args.Map{
		"stringType": "string", "intType": "int", "nilType": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns correct value -- with args", actual)
}

// ── Dynamic — LengthOfReflect ──

func Test_LengthOfReflect_Struct(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(struct{}{})

	// Act
	actual := args.Map{"length": d.Length()}

	// Assert
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- struct", actual)
}

func Test_LengthOfReflect_Map_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2})

	// Act
	actual := args.Map{"length": d.Length()}

	// Assert
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- map", actual)
}

// ── Dynamic — DynamicReflect.go ──

func Test_Dynamic_ReflectType_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"typeNotNil":  d.ReflectType() != nil,
		"typeNameNE":  d.ReflectTypeName() != "",
		"reflectValNN": d.ReflectValue() != nil,
	}

	// Assert
	expected := args.Map{
		"typeNotNil": true, "typeNameNE": true, "reflectValNN": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ReflectType", actual)
}

func Test_Dynamic_Loop_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	count := 0
	d.Loop(func(index int, item any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Loop", actual)
}

func Test_Dynamic_Loop_Invalid_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()
	called := d.Loop(func(index int, item any) bool { return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns error -- Loop invalid", actual)
}

func Test_Dynamic_LoopMap_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	count := 0
	d.LoopMap(func(index int, key, value any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- LoopMap", actual)
}

// ── Dynamic — Constructors ──

func Test_Dynamic_InvalidDynamic_FromDynamicJsonString(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()

	// Act
	actual := args.Map{
		"isValid":  d.IsValid(),
		"isNull":   d.IsNull(),
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"isNull": true,
	}
	expected.ShouldBeEqual(t, 0, "InvalidDynamic returns error -- with args", actual)
}

func Test_Dynamic_NonPtr_Ptr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("test")
	nonPtr := d.NonPtr()
	ptr := d.Ptr()

	// Act
	actual := args.Map{
		"nonPtrValid": nonPtr.IsValid(),
		"ptrNotNil":   ptr != nil,
	}

	// Assert
	expected := args.Map{
		"nonPtrValid": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- NonPtr/Ptr", actual)
}

func Test_Dynamic_ClonePtr_Nil_FromDynamicJsonString(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"isNil": d.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- ClonePtr nil", actual)
}
