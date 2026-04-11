package coreoncetests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// IntegerOnce — comprehensive method coverage
// ==========================================================================

func Test_IntegerOnce_AllMethods(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return 5 })

	// Act
	actual := args.Map{
		"value": io.Value(), "execute": io.Execute(),
		"isEmpty": io.IsEmpty(), "isZero": io.IsZero(),
		"isAboveZero": io.IsAboveZero(), "isAboveEqualZero": io.IsAboveEqualZero(),
		"isLessThanZero": io.IsLessThanZero(), "isLessThanEqualZero": io.IsLessThanEqualZero(),
		"isAbove3": io.IsAbove(3), "isAboveEqual5": io.IsAboveEqual(5),
		"isLessThan10": io.IsLessThan(10), "isLessThanEqual5": io.IsLessThanEqual(5),
		"isInvalidIndex": io.IsInvalidIndex(), "isValidIndex": io.IsValidIndex(),
		"isNegative": io.IsNegative(), "isPositive": io.IsPositive(),
		"string": io.String(),
	}

	// Assert
	expected := args.Map{
		"value": 5, "execute": 5,
		"isEmpty": false, "isZero": false,
		"isAboveZero": true, "isAboveEqualZero": true,
		"isLessThanZero": false, "isLessThanEqualZero": false,
		"isAbove3": true, "isAboveEqual5": true,
		"isLessThan10": true, "isLessThanEqual5": true,
		"isInvalidIndex": false, "isValidIndex": true,
		"isNegative": false, "isPositive": true,
		"string": "5",
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce AllMethods returns expected -- value 5", actual)
}

func Test_IntegerOnce_NegativeValue(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return -3 })

	// Act
	actual := args.Map{
		"isNegative": io.IsNegative(), "isPositive": io.IsPositive(),
		"isAboveZero": io.IsAboveZero(), "isLessThanZero": io.IsLessThanZero(),
		"isInvalidIndex": io.IsInvalidIndex(),
	}

	// Assert
	expected := args.Map{
		"isNegative": true, "isPositive": false,
		"isAboveZero": false, "isLessThanZero": true,
		"isInvalidIndex": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce negative returns expected -- value -3", actual)
}

func Test_IntegerOnce_MarshalUnmarshal(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return 42 })
	marshalledBytes, marshalErr := io.MarshalJSON()
	unmarshalErr := io.UnmarshalJSON(marshalledBytes)
	_, serErr := io.Serialize()

	// Act
	actual := args.Map{
		"marshalOk": marshalErr == nil, "unmarshalOk": unmarshalErr == nil,
		"serializeOk": serErr == nil,
	}

	// Assert
	expected := args.Map{
		"marshalOk": true,
		"unmarshalOk": true,
		"serializeOk": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce Marshal/Unmarshal returns no error -- value 42", actual)
}

// ==========================================================================
// BoolOnce — comprehensive method coverage
// ==========================================================================

func Test_BoolOnce_AllMethods(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return true })

	// Act
	actual := args.Map{
		"value": bo.Value(), "execute": bo.Execute(),
		"string": bo.String(),
	}

	// Assert
	expected := args.Map{
		"value": true,
		"execute": true,
		"string": "true",
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce AllMethods returns expected -- value true", actual)
}

func Test_BoolOnce_MarshalUnmarshal(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return false })
	marshalledBytes, marshalErr := bo.MarshalJSON()
	unmarshalErr := bo.UnmarshalJSON(marshalledBytes)
	_, serErr := bo.Serialize()

	// Act
	actual := args.Map{
		"marshalOk": marshalErr == nil, "unmarshalOk": unmarshalErr == nil,
		"serializeOk": serErr == nil, "valAfterUnmarshal": bo.Value(),
	}

	// Assert
	expected := args.Map{
		"marshalOk": true, "unmarshalOk": true,
		"serializeOk": true, "valAfterUnmarshal": false,
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce Marshal/Unmarshal returns no error -- value false", actual)
}

// ==========================================================================
// ByteOnce — comprehensive method coverage
// ==========================================================================

func Test_ByteOnce_AllMethods(t *testing.T) {
	// Arrange
	bo := coreonce.NewByteOnce(func() byte { return 42 })

	// Act
	actual := args.Map{
		"value": int(bo.Value()), "execute": int(bo.Execute()),
		"int": bo.Int(), "isEmpty": bo.IsEmpty(),
		"isZero": bo.IsZero(), "isNegative": bo.IsNegative(),
		"isPositive": bo.IsPositive(), "string": bo.String(),
	}

	// Assert
	expected := args.Map{
		"value": 42, "execute": 42,
		"int": 42, "isEmpty": false,
		"isZero": false, "isNegative": false,
		"isPositive": true, "string": "42",
	}
	expected.ShouldBeEqual(t, 0, "ByteOnce AllMethods returns expected -- value 42", actual)
}

func Test_ByteOnce_MarshalUnmarshal(t *testing.T) {
	// Arrange
	bo := coreonce.NewByteOnce(func() byte { return 10 })
	marshalledBytes, marshalErr := bo.MarshalJSON()
	unmarshalErr := bo.UnmarshalJSON(marshalledBytes)
	_, serErr := bo.Serialize()

	// Act
	actual := args.Map{
		"marshalOk": marshalErr == nil,
		"unmarshalOk": unmarshalErr == nil,
		"serializeOk": serErr == nil,
	}

	// Assert
	expected := args.Map{
		"marshalOk": true,
		"unmarshalOk": true,
		"serializeOk": true,
	}
	expected.ShouldBeEqual(t, 0, "ByteOnce Marshal/Unmarshal returns no error -- value 10", actual)
}

// ==========================================================================
// BytesOnce — comprehensive method coverage
// ==========================================================================

func Test_BytesOnce_AllMethods(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesOnce(func() []byte { return []byte("hello") })

	// Act
	actual := args.Map{
		"length": bo.Length(), "isEmpty": bo.IsEmpty(),
		"string": bo.String(), "execute": string(bo.Execute()),
	}

	// Assert
	expected := args.Map{
		"length": 5,
		"isEmpty": false,
		"string": "hello",
		"execute": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce AllMethods returns expected -- value hello", actual)
}

func Test_BytesOnce_NilFunc(t *testing.T) {
	// Arrange
	bo := coreonce.BytesOnce{}

	// Act
	actual := args.Map{
		"isEmpty": bo.IsEmpty(),
		"length": bo.Length(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"length": 0,
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce nil func returns empty -- no initializer", actual)
}

func Test_BytesOnce_MarshalUnmarshal(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesOnce(func() []byte { return []byte("test") })
	marshalledBytes, marshalErr := bo.MarshalJSON()
	unmarshalErr := bo.UnmarshalJSON(marshalledBytes)
	_, serErr := bo.Serialize()

	// Act
	actual := args.Map{
		"marshalOk": marshalErr == nil,
		"unmarshalOk": unmarshalErr == nil,
		"serializeOk": serErr == nil,
	}

	// Assert
	expected := args.Map{
		"marshalOk": true,
		"unmarshalOk": true,
		"serializeOk": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce Marshal/Unmarshal returns no error -- test", actual)
}

// ==========================================================================
// ErrorOnce — comprehensive method coverage
// ==========================================================================

func Test_ErrorOnce_WithNilError(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })

	// Act
	actual := args.Map{
		"hasError": eo.HasError(), "isEmpty": eo.IsEmpty(),
		"isEmptyError": eo.IsEmptyError(), "hasAnyItem": eo.HasAnyItem(),
		"isDefined": eo.IsDefined(), "isInvalid": eo.IsInvalid(),
		"isValid": eo.IsValid(), "isSuccess": eo.IsSuccess(),
		"isFailed": eo.IsFailed(), "isNull": eo.IsNull(),
		"isNullOrEmpty": eo.IsNullOrEmpty(), "message": eo.Message(),
		"isMessageEqual": eo.IsMessageEqual("test"),
	}

	// Assert
	expected := args.Map{
		"hasError": false, "isEmpty": true, "isEmptyError": true,
		"hasAnyItem": false, "isDefined": false, "isInvalid": false,
		"isValid": true, "isSuccess": true, "isFailed": false,
		"isNull": true, "isNullOrEmpty": true, "message": "",
		"isMessageEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce nil returns all safe -- nil error", actual)
}

func Test_ErrorOnce_ConcatNewString(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })
	result := eo.ConcatNewString("extra")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorOnce ConcatNewString returns extra -- nil error", actual)
}

func Test_ErrorOnce_ConcatNew_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })

	// Act
	actual := args.Map{"hasErr": eo.ConcatNew("msg") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ErrorOnce ConcatNew returns error -- always", actual)
}

func Test_ErrorOnce_MarshalUnmarshal(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })
	marshalledBytes, marshalErr := eo.MarshalJSON()
	unmarshalErr := eo.UnmarshalJSON(marshalledBytes)
	_, serErr := eo.Serialize()

	// Act
	actual := args.Map{
		"marshalOk": marshalErr == nil,
		"unmarshalOk": unmarshalErr == nil,
		"serializeOk": serErr == nil,
	}

	// Assert
	expected := args.Map{
		"marshalOk": true,
		"unmarshalOk": true,
		"serializeOk": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce Marshal/Unmarshal returns no error -- nil error", actual)
}

func Test_ErrorOnce_HandleError_NoError_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })
	eo.HandleError() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ErrorOnce HandleError no panic -- nil error", actual)
}

func Test_ErrorOnce_HandleErrorWith_NoError_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })
	eo.HandleErrorWith("msg") // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ErrorOnce HandleErrorWith no panic -- nil error", actual)
}

// ==========================================================================
// AnyOnce — comprehensive method coverage
// ==========================================================================

func Test_AnyOnce_ValueString_Nil_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return nil })

	// Act
	actual := args.Map{
		"isNull": ao.IsNull(), "isEmpty": ao.IsStringEmpty(),
		"isEmptyWs": ao.IsStringEmptyOrWhitespace(),
		"string": ao.String(), "isInit": ao.IsInitialized(),
	}

	// Assert
	expected := args.Map{
		"isNull": true, "isEmpty": true, "isEmptyWs": true,
		"string": "", "isInit": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce nil value returns empty -- nil initializer", actual)
}

func Test_AnyOnce_CastMethods(t *testing.T) {
	// Arrange
	aoStr := coreonce.NewAnyOnce(func() any { return "hello" })
	valStr, okStr := aoStr.CastValueString()
	aoStrings := coreonce.NewAnyOnce(func() any { return []string{"a"} })
	valStrings, okStrings := aoStrings.CastValueStrings()
	aoMap := coreonce.NewAnyOnce(func() any { return map[string]string{"k": "v"} })
	valMap, okMap := aoMap.CastValueHashmapMap()
	aoMapAny := coreonce.NewAnyOnce(func() any { return map[string]any{"k": 1} })
	valMapAny, okMapAny := aoMapAny.CastValueMapStringAnyMap()
	aoBytes := coreonce.NewAnyOnce(func() any { return []byte("hi") })
	valBytes, okBytes := aoBytes.CastValueBytes()

	// Act
	actual := args.Map{
		"str": valStr, "okStr": okStr,
		"stringsLen": len(valStrings), "okStrings": okStrings,
		"mapLen": len(valMap), "okMap": okMap,
		"mapAnyLen": len(valMapAny), "okMapAny": okMapAny,
		"bytesLen": len(valBytes), "okBytes": okBytes,
	}

	// Assert
	expected := args.Map{
		"str": "hello", "okStr": true,
		"stringsLen": 1, "okStrings": true,
		"mapLen": 1, "okMap": true,
		"mapAnyLen": 1, "okMapAny": true,
		"bytesLen": 2, "okBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce CastMethods return expected -- various types", actual)
}

func Test_AnyOnce_Serialize(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return "hello" })
	_, serErr := ao.Serialize()
	serMust := ao.SerializeMust()
	_, skipErr := ao.SerializeSkipExistingError()

	// Act
	actual := args.Map{
		"serOk": serErr == nil,
		"mustLen": len(serMust) > 0,
		"skipOk": skipErr == nil,
	}

	// Assert
	expected := args.Map{
		"serOk": true,
		"mustLen": true,
		"skipOk": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce Serialize returns no error -- valid value", actual)
}

func Test_AnyOnce_Deserialize(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return "hello" })
	var result string
	err := ao.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "AnyOnce Deserialize returns no error -- valid", actual)
}

// ==========================================================================
// AnyErrorOnce — comprehensive method coverage
// ==========================================================================

func Test_AnyErrorOnce_AllMethods(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })

	// Act
	actual := args.Map{
		"hasError": aeo.HasError(), "isEmptyError": aeo.IsEmptyError(),
		"isEmpty": aeo.IsEmpty(), "hasAnyItem": aeo.HasAnyItem(),
		"isDefined": aeo.IsDefined(), "isInvalid": aeo.IsInvalid(),
		"isValid": aeo.IsValid(), "isSuccess": aeo.IsSuccess(),
		"isFailed": aeo.IsFailed(), "isNull": aeo.IsNull(),
		"isInit": aeo.IsInitialized(), "isStringEmpty": aeo.IsStringEmpty(),
		"isStringEmptyWs": aeo.IsStringEmptyOrWhitespace(),
		"stringNotEmpty": aeo.String() != "",
	}

	// Assert
	expected := args.Map{
		"hasError": false, "isEmptyError": true, "isEmpty": false,
		"hasAnyItem": true, "isDefined": true, "isInvalid": false,
		"isValid": true, "isSuccess": true, "isFailed": false,
		"isNull": false, "isInit": true, "isStringEmpty": false,
		"isStringEmptyWs": false, "stringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce AllMethods returns expected -- valid value", actual)
}

func Test_AnyErrorOnce_CastMethods(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })
	valStr, err, ok := aeo.CastValueString()

	// Act
	actual := args.Map{
		"val": valStr,
		"hasErr": err != nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"hasErr": false,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce CastValueString returns value -- valid", actual)
}

func Test_AnyErrorOnce_CastStrings(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return []string{"a"}, nil })
	vals, err, ok := aeo.CastValueStrings()

	// Act
	actual := args.Map{
		"len": len(vals),
		"hasErr": err != nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasErr": false,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce CastValueStrings returns values -- valid", actual)
}

func Test_AnyErrorOnce_CastHashmapMap(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return map[string]string{"k": "v"}, nil })
	vals, err, ok := aeo.CastValueHashmapMap()

	// Act
	actual := args.Map{
		"len": len(vals),
		"hasErr": err != nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasErr": false,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce CastValueHashmapMap returns values -- valid", actual)
}

func Test_AnyErrorOnce_CastMapAny(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return map[string]any{"k": 1}, nil })
	vals, err, ok := aeo.CastValueMapStringAnyMap()

	// Act
	actual := args.Map{
		"len": len(vals),
		"hasErr": err != nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasErr": false,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce CastValueMapStringAnyMap returns values -- valid", actual)
}

func Test_AnyErrorOnce_CastBytes(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return []byte("hi"), nil })
	vals, err, ok := aeo.CastValueBytes()

	// Act
	actual := args.Map{
		"len": len(vals),
		"hasErr": err != nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasErr": false,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce CastValueBytes returns values -- valid", actual)
}

func Test_AnyErrorOnce_ValueString(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })
	val, err := aeo.ValueString()
	safeStr := aeo.SafeString()
	valStrOnly := aeo.ValueStringOnly()

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"notEmpty": val != "",
		"safe": safeStr != "",
		"only": valStrOnly != "",
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"notEmpty": true,
		"safe": true,
		"only": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueString returns value -- valid", actual)
}

func Test_AnyErrorOnce_ValueStringMust(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })
	val := aeo.ValueStringMust()

	// Act
	actual := args.Map{"notEmpty": val != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueStringMust returns value -- valid", actual)
}

func Test_AnyErrorOnce_ExecuteMust(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })
	val := aeo.ExecuteMust()

	// Act
	actual := args.Map{"notNil": val != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ExecuteMust returns value -- valid", actual)
}

func Test_AnyErrorOnce_ValueMust(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })
	val := aeo.ValueMust()

	// Act
	actual := args.Map{"notNil": val != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueMust returns value -- valid", actual)
}

func Test_AnyErrorOnce_Serialize(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })
	_, serErr := aeo.Serialize()
	serMust := aeo.SerializeMust()
	_, skipErr := aeo.SerializeSkipExistingError()

	// Act
	actual := args.Map{
		"serOk": serErr == nil,
		"mustLen": len(serMust) > 0,
		"skipOk": skipErr == nil,
	}

	// Assert
	expected := args.Map{
		"serOk": true,
		"mustLen": true,
		"skipOk": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce Serialize returns no error -- valid", actual)
}

func Test_AnyErrorOnce_Deserialize(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })
	var result string
	err := aeo.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce Deserialize returns no error -- valid", actual)
}

// ==========================================================================
// BytesErrorOnce — comprehensive method coverage
// ==========================================================================

func Test_BytesErrorOnce_AllMethods(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("hello"), nil })

	// Act
	actual := args.Map{
		"hasError": beo.HasError(), "isEmptyError": beo.IsEmptyError(),
		"isEmpty": beo.IsEmpty(), "isEmptyBytes": beo.IsEmptyBytes(),
		"length": beo.Length(), "hasAnyItem": beo.HasAnyItem(),
		"isDefined": beo.IsDefined(), "isInvalid": beo.IsInvalid(),
		"isValid": beo.IsValid(), "isSuccess": beo.IsSuccess(),
		"isFailed": beo.IsFailed(), "isInit": beo.IsInitialized(),
		"isBytesEmpty": beo.IsBytesEmpty(), "isNull": beo.IsNull(),
		"isStringEmpty": beo.IsStringEmpty(), "isStringEmptyWs": beo.IsStringEmptyOrWhitespace(),
		"string": beo.String(), "hasSafe": beo.HasSafeItems(),
		"hasIssues": beo.HasIssuesOrEmpty(),
	}

	// Assert
	expected := args.Map{
		"hasError": false, "isEmptyError": true, "isEmpty": false,
		"isEmptyBytes": false, "length": 5, "hasAnyItem": true,
		"isDefined": true, "isInvalid": false, "isValid": true,
		"isSuccess": true, "isFailed": false, "isInit": true,
		"isBytesEmpty": false, "isNull": false, "isStringEmpty": false,
		"isStringEmptyWs": false, "string": "hello", "hasSafe": true,
		"hasIssues": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce AllMethods returns expected -- valid bytes", actual)
}

func Test_BytesErrorOnce_ValueWithError_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("hi"), nil })
	val, err := beo.ValueWithError()

	// Act
	actual := args.Map{
		"len": len(val),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce ValueWithError returns value -- valid", actual)
}

func Test_BytesErrorOnce_Execute_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("hi"), nil })
	val, err := beo.Execute()

	// Act
	actual := args.Map{
		"len": len(val),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce Execute returns value -- valid", actual)
}

func Test_BytesErrorOnce_HandleError_NoError_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return nil, nil })
	beo.HandleError() // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce HandleError no panic -- nil error", actual)
}

func Test_BytesErrorOnce_MustBeEmptyError_NoError_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return nil, nil })
	beo.MustBeEmptyError() // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce MustBeEmptyError no panic -- nil error", actual)
}

func Test_BytesErrorOnce_MustHaveSafeItems(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("data"), nil })
	beo.MustHaveSafeItems() // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce MustHaveSafeItems no panic -- has data", actual)
}

func Test_BytesErrorOnce_MarshalJSON_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) {
		return json.Marshal("hello")
	})
	marshalBytes, marshalErr := beo.MarshalJSON()
	_, serErr := beo.Serialize()

	// Act
	actual := args.Map{
		"marshalOk": marshalErr == nil, "marshalLen": len(marshalBytes) > 0,
		"serializeOk": serErr == nil,
	}

	// Assert
	expected := args.Map{
		"marshalOk": true,
		"marshalLen": true,
		"serializeOk": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce MarshalJSON returns bytes -- valid", actual)
}

func Test_BytesErrorOnce_SerializeMust_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) {
		return json.Marshal("hello")
	})
	result := beo.SerializeMust()

	// Act
	actual := args.Map{"hasData": len(result) > 0}

	// Assert
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce SerializeMust returns bytes -- valid", actual)
}

func Test_BytesErrorOnce_Deserialize_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) {
		return json.Marshal("hello")
	})
	var result string
	err := beo.Deserialize(&result)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"result": result,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce Deserialize returns no error -- valid json", actual)
}

// ==========================================================================
// StringsOnce — UnmarshalJSON
// ==========================================================================

func Test_StringsOnce_UnmarshalJSON(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return nil })
	data, _ := json.Marshal([]string{"a", "b"})
	err := so.UnmarshalJSON(data)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": so.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce UnmarshalJSON populates values -- valid json", actual)
}

// ==========================================================================
// IntegersOnce — UnmarshalJSON, IsZero
// ==========================================================================

func Test_IntegersOnce_UnmarshalJSON(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOnce(func() []int { return nil })
	data, _ := json.Marshal([]int{1, 2, 3})
	err := io.UnmarshalJSON(data)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"isEmpty": io.IsEmpty(),
		"isZero": io.IsZero(),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"isEmpty": false,
		"isZero": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce UnmarshalJSON populates values -- valid json", actual)
}

// ==========================================================================
// MapStringStringOnce — UnmarshalJSON
// ==========================================================================

func Test_MapStringStringOnce_UnmarshalJSON_FromIntegerOnceAllMethod(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string { return nil })
	data, _ := json.Marshal(map[string]string{"a": "1"})
	err := mso.UnmarshalJSON(data)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": mso.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce UnmarshalJSON populates -- valid json", actual)
}
