package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── StringOnce ──

func Test_StringOnce_SplitLeftRight(t *testing.T) {
	// Arrange
	so := coreonce.NewStringOnce(func() string { return "key=value" })
	l, r := so.SplitLeftRight("=")
	lt, rt := so.SplitLeftRightTrim("=")

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
		"leftTrim": lt,
		"rightTrim": rt,
	}

	// Assert
	expected := args.Map{
		"left": "key",
		"right": "value",
		"leftTrim": "key",
		"rightTrim": "value",
	}
	expected.ShouldBeEqual(t, 0, "StringOnce returns correct value -- SplitLeftRight", actual)
}

func Test_StringOnce_SplitLeftRight_NoSep(t *testing.T) {
	// Arrange
	so := coreonce.NewStringOnce(func() string { return "nosep" })
	l, r := so.SplitLeftRight("=")

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": "nosep",
		"right": "",
	}
	expected.ShouldBeEqual(t, 0, "StringOnce returns empty -- SplitLeftRight no separator", actual)
}

func Test_StringOnce_Helpers(t *testing.T) {
	// Arrange
	so := coreonce.NewStringOnce(func() string { return "Hello World" })

	// Act
	actual := args.Map{
		"isEqual":      so.IsEqual("Hello World"),
		"hasPrefix":    so.HasPrefix("Hello"),
		"hasSuffix":    so.HasSuffix("World"),
		"startsWith":   so.IsStartsWith("Hello"),
		"endsWith":     so.IsEndsWith("World"),
		"contains":     so.IsContains("lo Wo"),
		"isEmpty":      so.IsEmpty(),
		"isEmptyOrWS":  so.IsEmptyOrWhitespace(),
		"bytesLen":     len(so.Bytes()) > 0,
		"errorNotNil":  so.Error() != nil,
		"valuePtrNN":   so.ValuePtr() != nil,
	}

	// Assert
	expected := args.Map{
		"isEqual": true, "hasPrefix": true, "hasSuffix": true,
		"startsWith": true, "endsWith": true, "contains": true,
		"isEmpty": false, "isEmptyOrWS": false,
		"bytesLen": true, "errorNotNil": true, "valuePtrNN": true,
	}
	expected.ShouldBeEqual(t, 0, "StringOnce returns correct value -- helpers", actual)
}

func Test_StringOnce_SplitBy_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	so := coreonce.NewStringOnce(func() string { return "a,b,c" })
	result := so.SplitBy(",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "StringOnce returns correct value -- SplitBy", actual)
}

func Test_StringOnce_Execute(t *testing.T) {
	// Arrange
	so := coreonce.NewStringOnce(func() string { return "exec" })

	// Act
	actual := args.Map{"val": so.Execute()}

	// Assert
	expected := args.Map{"val": "exec"}
	expected.ShouldBeEqual(t, 0, "StringOnce returns correct value -- Execute", actual)
}

func Test_StringOnce_Serialize(t *testing.T) {
	// Arrange
	so := coreonce.NewStringOnce(func() string { return "test" })
	data, err := so.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": len(data) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "StringOnce returns correct value -- Serialize", actual)
}

func Test_StringOnce_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	so := coreonce.NewStringOnce(func() string { return "hello" })
	data, err := so.MarshalJSON()
	var so2 coreonce.StringOnce
	err2 := so2.UnmarshalJSON(data)

	// Act
	actual := args.Map{
		"noErr": err == nil, "noErr2": err2 == nil,
		"value": so2.Value(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"noErr2": true,
		"value": "hello",
	}
	expected.ShouldBeEqual(t, 0, "StringOnce returns correct value -- MarshalUnmarshalJSON", actual)
}

// ── BoolOnce ──

func Test_BoolOnce_Execute(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return true })

	// Act
	actual := args.Map{
		"val":    bo.Value(),
		"exec":   bo.Execute(),
		"string": bo.String(),
	}

	// Assert
	expected := args.Map{
		"val": true,
		"exec": true,
		"string": "true",
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce returns correct value -- Execute", actual)
}

func Test_BoolOnce_Serialize_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return false })
	data, err := bo.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": len(data) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce returns correct value -- Serialize", actual)
}

func Test_BoolOnce_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return true })
	data, err := bo.MarshalJSON()
	var bo2 coreonce.BoolOnce
	err2 := bo2.UnmarshalJSON(data)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"noErr2": err2 == nil,
		"val": bo2.Value(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"noErr2": true,
		"val": true,
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce returns correct value -- MarshalUnmarshalJSON", actual)
}

// ── IntegerOnce ──

func Test_IntegerOnce_Comparisons_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return 5 })

	// Act
	actual := args.Map{
		"value":        io.Value(),
		"execute":      io.Execute(),
		"isEmpty":      io.IsEmpty(),
		"isZero":       io.IsZero(),
		"isAboveZero":  io.IsAboveZero(),
		"isAboveEqZ":   io.IsAboveEqualZero(),
		"isLtZero":     io.IsLessThanZero(),
		"isLtEqZ":      io.IsLessThanEqualZero(),
		"isAbove3":     io.IsAbove(3),
		"isAboveEq5":   io.IsAboveEqual(5),
		"isLt10":       io.IsLessThan(10),
		"isLtEq5":      io.IsLessThanEqual(5),
		"invalidIndex": io.IsInvalidIndex(),
		"validIndex":   io.IsValidIndex(),
		"isNegative":   io.IsNegative(),
		"isPositive":   io.IsPositive(),
		"string":       io.String(),
	}

	// Assert
	expected := args.Map{
		"value": 5, "execute": 5,
		"isEmpty": false, "isZero": false,
		"isAboveZero": true, "isAboveEqZ": true,
		"isLtZero": false, "isLtEqZ": false,
		"isAbove3": true, "isAboveEq5": true,
		"isLt10": true, "isLtEq5": true,
		"invalidIndex": false, "validIndex": true,
		"isNegative": false, "isPositive": true,
		"string": "5",
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce returns correct value -- Comparisons", actual)
}

func Test_IntegerOnce_Serialize_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return 42 })
	data, err := io.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": len(data) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce returns correct value -- Serialize", actual)
}

// ── ErrorOnce ──

func Test_ErrorOnce_NilError(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })

	// Act
	actual := args.Map{
		"hasError": eo.HasError(),
		"isEmpty":  eo.IsEmpty(),
		"isValid":  eo.IsValid(),
		"isNull":   eo.IsNull(),
		"message":  eo.Message(),
		"isSuccess": eo.IsSuccess(),
		"isFailed":  eo.IsFailed(),
		"isDefined": eo.IsDefined(),
		"hasAny":    eo.HasAnyItem(),
		"isInvalid": eo.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"hasError": false, "isEmpty": true, "isValid": true,
		"isNull": true, "message": "",
		"isSuccess": true, "isFailed": false,
		"isDefined": false, "hasAny": false, "isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce returns nil -- nil error", actual)
}

func Test_ErrorOnce_IsMessageEqual_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })

	// Act
	actual := args.Map{"msgEq": eo.IsMessageEqual("test")}

	// Assert
	expected := args.Map{"msgEq": false}
	expected.ShouldBeEqual(t, 0, "ErrorOnce returns nil -- IsMessageEqual nil", actual)
}

func Test_ErrorOnce_ConcatNew_NilError(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })
	result := eo.ConcatNewString("extra")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorOnce returns nil -- ConcatNewString nil error", actual)
}

// ── MapStringStringOnce ──

func Test_MapStringStringOnce_AllKeys_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	m := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})

	// Act
	actual := args.Map{
		"length":       m.Length(),
		"hasAnyItem":   m.HasAnyItem(),
		"isEmpty":      m.IsEmpty(),
		"hasA":         m.Has("a"),
		"containsA":    m.IsContains("a"),
		"missingC":     m.IsMissing("c"),
		"getA":         m.GetValue("a"),
		"hasAll":       m.HasAll("a", "b"),
		"keysLen":      len(m.AllKeys()),
		"valuesLen":    len(m.AllValues()),
		"keysSrtLen":   len(m.AllKeysSorted()),
		"valuesSrtLen": len(m.AllValuesSorted()),
		"stringsLen":   len(m.Strings()),
	}

	// Assert
	expected := args.Map{
		"length": 2, "hasAnyItem": true, "isEmpty": false,
		"hasA": true, "containsA": true, "missingC": true,
		"getA": "1", "hasAll": true,
		"keysLen": 2, "valuesLen": 2,
		"keysSrtLen": 2, "valuesSrtLen": 2,
		"stringsLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce returns correct value -- AllKeys", actual)
}

func Test_MapStringStringOnce_GetValueWithStatus_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	m := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"key": "val"}
	})
	v, has := m.GetValueWithStatus("key")
	v2, has2 := m.GetValueWithStatus("missing")

	// Act
	actual := args.Map{
		"val": v, "has": has,
		"val2": v2, "has2": has2,
	}

	// Assert
	expected := args.Map{
		"val": "val", "has": true,
		"val2": "", "has2": false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce returns non-empty -- GetValueWithStatus", actual)
}

func Test_MapStringStringOnce_IsEqual_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	m := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1"}
	})

	// Act
	actual := args.Map{
		"equal":    m.IsEqual(map[string]string{"a": "1"}),
		"notEqual": m.IsEqual(map[string]string{"a": "2"}),
		"missing":  m.IsEqual(map[string]string{"b": "1"}),
		"diffLen":  m.IsEqual(map[string]string{"a": "1", "b": "2"}),
	}

	// Assert
	expected := args.Map{
		"equal": true, "notEqual": false,
		"missing": false, "diffLen": false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce returns correct value -- IsEqual", actual)
}

func Test_MapStringStringOnce_String_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	m := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1"}
	})

	// Act
	actual := args.Map{"notEmpty": m.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce returns correct value -- String", actual)
}

func Test_MapStringStringOnce_JsonStringMust_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	m := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	result := m.JsonStringMust()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce returns correct value -- JsonStringMust", actual)
}

// ── AnyOnce ──

func Test_AnyOnce_CastMethods_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	aoStr := coreonce.NewAnyOnce(func() any { return "hello" })
	aoStrings := coreonce.NewAnyOnce(func() any { return []string{"a", "b"} })
	aoMap := coreonce.NewAnyOnce(func() any { return map[string]string{"k": "v"} })
	aoMapAny := coreonce.NewAnyOnce(func() any { return map[string]any{"k": 1} })
	aoBytes := coreonce.NewAnyOnce(func() any { return []byte("hi") })

	vStr, okStr := aoStr.CastValueString()
	vStrings, okStrings := aoStrings.CastValueStrings()
	vMap, okMap := aoMap.CastValueHashmapMap()
	vMapAny, okMapAny := aoMapAny.CastValueMapStringAnyMap()
	vBytes, okBytes := aoBytes.CastValueBytes()

	// Act
	actual := args.Map{
		"str": vStr, "okStr": okStr,
		"stringsLen": len(vStrings), "okStrings": okStrings,
		"mapLen": len(vMap), "okMap": okMap,
		"mapAnyLen": len(vMapAny), "okMapAny": okMapAny,
		"bytesLen": len(vBytes), "okBytes": okBytes,
	}

	// Assert
	expected := args.Map{
		"str": "hello", "okStr": true,
		"stringsLen": 2, "okStrings": true,
		"mapLen": 1, "okMap": true,
		"mapAnyLen": 1, "okMapAny": true,
		"bytesLen": 2, "okBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce returns correct value -- CastMethods", actual)
}

func Test_AnyOnce_Null(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return nil })
	isInit := ao.IsInitialized()

	// Act
	actual := args.Map{
		"isNull":           ao.IsNull(),
		"isStringEmpty":    ao.IsStringEmpty(),
		"isStringEmptyWS": ao.IsStringEmptyOrWhitespace(),
		"isInitialized":   isInit,
	}

	// Assert
	expected := args.Map{
		"isNull": true, "isStringEmpty": true,
		"isStringEmptyWS": true, "isInitialized": isInit,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce returns correct value -- null value", actual)
}

func Test_AnyOnce_SerializeMust_FromStringOnceSplitLeftR(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return "test" })
	result := ao.SerializeMust()

	// Act
	actual := args.Map{"notEmpty": len(result) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyOnce returns correct value -- SerializeMust", actual)
}

func Test_AnyOnce_ValueString(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return 42 })
	v1 := ao.ValueString()
	v2 := ao.ValueStringOnly()
	v3 := ao.SafeString()
	v4 := ao.ValueStringMust()

	// Act
	actual := args.Map{
		"v1NotEmpty": v1 != "",
		"v2Eq":       v1 == v2,
		"v3Eq":       v1 == v3,
		"v4Eq":       v1 == v4,
	}

	// Assert
	expected := args.Map{
		"v1NotEmpty": true, "v2Eq": true, "v3Eq": true, "v4Eq": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce returns non-empty -- ValueString variants", actual)
}
