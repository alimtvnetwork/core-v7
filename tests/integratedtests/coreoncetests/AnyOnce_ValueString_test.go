package coreoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── AnyOnce — uncovered branches ──

func Test_AnyOnce_ValueString_NilValue(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return nil })
	val := ao.ValueString()

	// Act
	actual := args.Map{"containsNil": val != ""}

	// Assert
	expected := args.Map{"containsNil": true}
	expected.ShouldBeEqual(t, 0, "ValueString returns nil bracket -- nil initializer", actual)
}

func Test_AnyOnce_ValueString_Cached_FromAnyOnceValueString(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "hello" })
	_ = ao.ValueString() // first call compiles
	val := ao.ValueString() // second call uses cache

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ValueString returns cached -- second call", actual)
}

func Test_AnyOnce_CastValueString(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "hello" })
	val, ok := ao.CastValueString()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueString returns value -- string value", actual)
}

func Test_AnyOnce_CastValueString_Fail(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return 42 })
	_, ok := ao.CastValueString()

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "CastValueString returns false -- int value", actual)
}

func Test_AnyOnce_CastValueStrings(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return []string{"a", "b"} })
	val, ok := ao.CastValueStrings()

	// Act
	actual := args.Map{
		"len": len(val),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueStrings returns slice -- string slice", actual)
}

func Test_AnyOnce_CastValueHashmapMap(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return map[string]string{"k": "v"} })
	val, ok := ao.CastValueHashmapMap()

	// Act
	actual := args.Map{
		"len": len(val),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueHashmapMap returns map -- map value", actual)
}

func Test_AnyOnce_CastValueMapStringAnyMap(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return map[string]any{"k": 1} })
	val, ok := ao.CastValueMapStringAnyMap()

	// Act
	actual := args.Map{
		"len": len(val),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueMapStringAnyMap returns map -- map any value", actual)
}

func Test_AnyOnce_CastValueBytes(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return []byte{1, 2} })
	val, ok := ao.CastValueBytes()

	// Act
	actual := args.Map{
		"len": len(val),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueBytes returns bytes -- byte slice", actual)
}

func Test_AnyOnce_IsStringEmpty(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return nil })

	// Act
	actual := args.Map{"isEmpty": ao.IsStringEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsStringEmpty returns true -- nil value", actual)
}

func Test_AnyOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "  " })

	// Act
	actual := args.Map{"isEmpty": ao.IsStringEmptyOrWhitespace()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsStringEmptyOrWhitespace returns true -- whitespace value", actual)
}

func Test_AnyOnce_Deserialize_Success_FromAnyOnceValueString(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return map[string]any{"key": "val"} })
	var result map[string]any
	err := ao.Deserialize(&result)

	// Act
	actual := args.Map{"err": err == nil}

	// Assert
	expected := args.Map{"err": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns nil error -- valid json", actual)
}

func Test_AnyOnce_Serialize_MarshalError(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return func() {} })
	_, err := ao.Serialize()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns error -- unmarshallable func", actual)
}

func Test_AnyOnce_SerializeSkipExistingError(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "test" })
	b, err := ao.SerializeSkipExistingError()

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingError returns bytes -- string value", actual)
}

func Test_AnyOnce_SerializeMust(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "test" })
	b := ao.SerializeMust()

	// Act
	actual := args.Map{"hasBytes": len(b) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "SerializeMust returns bytes -- string value", actual)
}

// ── AnyErrorOnce — uncovered branches ──

func Test_AnyErrorOnce_ValueString_NilValue(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	val, err := ao.ValueString()

	// Act
	actual := args.Map{
		"containsNil": val != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"containsNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ValueString returns nil bracket -- nil value no error", actual)
}

func Test_AnyErrorOnce_ValueString_Cached(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hi", nil })
	_, _ = ao.ValueString()
	val, _ := ao.ValueString()

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hi"}
	expected.ShouldBeEqual(t, 0, "ValueString returns cached -- second call", actual)
}

func Test_AnyErrorOnce_ValueStringMust_Panic(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("fail") })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"recovered": r != nil}

	// Assert
		expected := args.Map{"recovered": true}
		expected.ShouldBeEqual(t, 0, "ValueStringMust panics -- error present", actual)
	}()
	ao.ValueStringMust()
}

func Test_AnyErrorOnce_ExecuteMust_Panic(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("fail") })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"recovered": r != nil}

	// Assert
		expected := args.Map{"recovered": true}
		expected.ShouldBeEqual(t, 0, "ExecuteMust panics -- error present", actual)
	}()
	ao.ExecuteMust()
}

func Test_AnyErrorOnce_CastValueString(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hello", nil })
	val, err, ok := ao.CastValueString()

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"noErr": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueString returns value -- string", actual)
}

func Test_AnyErrorOnce_CastValueStrings(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []string{"a"}, nil })
	val, _, ok := ao.CastValueStrings()

	// Act
	actual := args.Map{
		"len": len(val),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueStrings returns slice -- string slice", actual)
}

func Test_AnyErrorOnce_CastValueHashmapMap(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]string{"k": "v"}, nil })
	val, _, ok := ao.CastValueHashmapMap()

	// Act
	actual := args.Map{
		"len": len(val),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueHashmapMap returns map -- map value", actual)
}

func Test_AnyErrorOnce_CastValueMapStringAnyMap(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]any{"k": 1}, nil })
	val, _, ok := ao.CastValueMapStringAnyMap()

	// Act
	actual := args.Map{
		"len": len(val),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueMapStringAnyMap returns map -- map any", actual)
}

func Test_AnyErrorOnce_CastValueBytes(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []byte{1}, nil })
	val, _, ok := ao.CastValueBytes()

	// Act
	actual := args.Map{
		"len": len(val),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CastValueBytes returns bytes -- byte slice", actual)
}

func Test_AnyErrorOnce_IsStringEmpty(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })

	// Act
	actual := args.Map{"isEmpty": ao.IsStringEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsStringEmpty returns true -- nil value", actual)
}

func Test_AnyErrorOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "  ", nil })

	// Act
	actual := args.Map{"isEmpty": ao.IsStringEmptyOrWhitespace()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsStringEmptyOrWhitespace returns true -- whitespace value", actual)
}

func Test_AnyErrorOnce_Serialize_ExistingError(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("pre-err") })
	_, err := ao.Serialize()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns error -- existing error", actual)
}

func Test_AnyErrorOnce_Serialize_MarshalError(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return func() {}, nil })
	_, err := ao.Serialize()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns error -- unmarshallable func", actual)
}

func Test_AnyErrorOnce_SerializeSkipExistingError(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "ok", errors.New("err") })
	b, err := ao.SerializeSkipExistingError()

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingError returns bytes -- ignores existing error", actual)
}

func Test_AnyErrorOnce_SerializeMust_Panic(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("fail") })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"recovered": r != nil}

	// Assert
		expected := args.Map{"recovered": true}
		expected.ShouldBeEqual(t, 0, "SerializeMust panics -- error present", actual)
	}()
	ao.SerializeMust()
}

func Test_AnyErrorOnce_IsEmpty_Nil(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })

	// Act
	actual := args.Map{"isEmpty": ao.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- nil value nil error", actual)
}

func Test_AnyErrorOnce_HasAnyItem(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "x", nil })

	// Act
	actual := args.Map{"has": ao.HasAnyItem()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns true -- non-nil value", actual)
}

func Test_AnyErrorOnce_IsValid_IsInvalid(t *testing.T) {
	// Arrange
	aoOk := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "x", nil })
	aoErr := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })

	// Act
	actual := args.Map{
		"valid":   aoOk.IsValid(),
		"success": aoOk.IsSuccess(),
		"invalid": aoErr.IsInvalid(),
		"failed":  aoErr.IsFailed(),
	}

	// Assert
	expected := args.Map{
		"valid":   true,
		"success": true,
		"invalid": true,
		"failed":  true,
	}
	expected.ShouldBeEqual(t, 0, "IsValid IsInvalid correct -- success and failure", actual)
}

// ── BytesErrorOnce — uncovered branches ──

func Test_BytesErrorOnce_HasIssuesOrEmpty_NilBytes(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })

	// Act
	actual := args.Map{"hasIssues": bo.HasIssuesOrEmpty()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "HasIssuesOrEmpty returns true -- nil bytes", actual)
}

func Test_BytesErrorOnce_HasSafeItems(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte{1}, nil })

	// Act
	actual := args.Map{"hasSafe": bo.HasSafeItems()}

	// Assert
	expected := args.Map{"hasSafe": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems returns true -- has bytes no error", actual)
}

func Test_BytesErrorOnce_Deserialize_ExistingError(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("fail") })
	var result string
	err := bo.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- existing error", actual)
}

func Test_BytesErrorOnce_Deserialize_UnmarshalError(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("not-json"), nil })
	var result int
	err := bo.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- invalid json", actual)
}

func Test_BytesErrorOnce_Deserialize_Success(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte(`"hello"`), nil })
	var result string
	err := bo.Deserialize(&result)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": result,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns value -- valid json string", actual)
}

func Test_BytesErrorOnce_DeserializeMust_Panic(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("fail") })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"recovered": r != nil}

	// Assert
		expected := args.Map{"recovered": true}
		expected.ShouldBeEqual(t, 0, "DeserializeMust panics -- error present", actual)
	}()
	bo.DeserializeMust(nil)
}

func Test_BytesErrorOnce_MustHaveSafeItems_PanicOnEmpty(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"recovered": r != nil}

	// Assert
		expected := args.Map{"recovered": true}
		expected.ShouldBeEqual(t, 0, "MustHaveSafeItems panics -- empty bytes no error", actual)
	}()
	bo.MustHaveSafeItems()
}

func Test_BytesErrorOnce_MustHaveSafeItems_PanicOnError(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("e") })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"recovered": r != nil}

	// Assert
		expected := args.Map{"recovered": true}
		expected.ShouldBeEqual(t, 0, "MustHaveSafeItems panics -- has error", actual)
	}()
	bo.MustHaveSafeItems()
}

func Test_BytesErrorOnce_IsEmptyBytes(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })

	// Act
	actual := args.Map{"isEmpty": bo.IsEmptyBytes()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyBytes returns true -- nil bytes", actual)
}

func Test_BytesErrorOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("  "), nil })

	// Act
	actual := args.Map{"isEmpty": bo.IsStringEmptyOrWhitespace()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsStringEmptyOrWhitespace returns true -- whitespace bytes", actual)
}

func Test_BytesErrorOnce_SerializeMust_Panic(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("fail") })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"recovered": r != nil}

	// Assert
		expected := args.Map{"recovered": true}
		expected.ShouldBeEqual(t, 0, "SerializeMust panics -- error present", actual)
	}()
	bo.SerializeMust()
}

// ── ErrorOnce — uncovered branches ──

func Test_ErrorOnce_HandleErrorWith_Panic(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return errors.New("oops") })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"recovered": r != nil}

	// Assert
		expected := args.Map{"recovered": true}
		expected.ShouldBeEqual(t, 0, "HandleErrorWith panics -- has error with message", actual)
	}()
	eo.HandleErrorWith("extra", "context")
}

func Test_ErrorOnce_ConcatNew(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return errors.New("base") })
	err := eo.ConcatNew("msg1", "msg2")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ConcatNew returns error -- has base error", actual)
}

func Test_ErrorOnce_ConcatNewString_NilError(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return nil })
	result := eo.ConcatNewString("msg1")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ConcatNewString returns messages only -- nil error", actual)
}

func Test_ErrorOnce_IsMessageEqual(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return errors.New("exact") })

	// Act
	actual := args.Map{
		"match":   eo.IsMessageEqual("exact"),
		"noMatch": eo.IsMessageEqual("other"),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsMessageEqual returns correct -- exact and mismatch", actual)
}

func Test_ErrorOnce_IsMessageEqual_NilError(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return nil })

	// Act
	actual := args.Map{"match": eo.IsMessageEqual("any")}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMessageEqual returns false -- nil error", actual)
}

// ── IntegersOnce — uncovered branches ──

func Test_IntegersOnce_IsEqual_DiffContent(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2, 3} })

	// Act
	actual := args.Map{
		"same":     io.IsEqual(1, 2, 3),
		"diff":     io.IsEqual(1, 2, 4),
		"diffLen":  io.IsEqual(1, 2),
		"sameNil":  io.IsEqual(1, 2, 3),
	}

	// Assert
	expected := args.Map{
		"same":     true,
		"diff":     false,
		"diffLen":  false,
		"sameNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct -- various comparisons", actual)
}

func Test_IntegersOnce_Sorted_Cached(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{3, 1, 2} })
	sorted1 := io.Sorted()
	sorted2 := io.Sorted()

	// Act
	actual := args.Map{
		"same": len(sorted1) == len(sorted2),
		"first": sorted1[0],
	}

	// Assert
	expected := args.Map{
		"same": true,
		"first": 1,
	}
	expected.ShouldBeEqual(t, 0, "Sorted returns cached -- second call same result", actual)
}

func Test_IntegersOnce_RangesMap(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{10, 20} })
	m := io.RangesMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangesMap returns correct length -- two items", actual)
}

func Test_IntegersOnce_RangesBoolMap(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{5} })
	m := io.RangesBoolMap()

	// Act
	actual := args.Map{"has5": m[5]}

	// Assert
	expected := args.Map{"has5": true}
	expected.ShouldBeEqual(t, 0, "RangesBoolMap returns true -- value 5 present", actual)
}

func Test_IntegersOnce_RangesMap_Empty(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := io.RangesMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangesMap returns empty -- no items", actual)
}

// ── StringsOnce — uncovered branches ──

func Test_StringsOnce_UniqueMapLock(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b", "a"} })
	m := so.UniqueMapLock()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UniqueMapLock returns deduped map -- duplicates present", actual)
}

func Test_StringsOnce_UniqueMap_NilValues(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return nil })
	m := so.UniqueMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "UniqueMap returns empty -- nil initializer", actual)
}

func Test_StringsOnce_UniqueMap_Cached(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"x"} })
	_ = so.UniqueMap()
	m := so.UniqueMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "UniqueMap returns cached -- second call", actual)
}

func Test_StringsOnce_IsEqual_DiffContent(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })

	// Act
	actual := args.Map{
		"same":    so.IsEqual("a", "b"),
		"diff":    so.IsEqual("a", "c"),
		"diffLen": so.IsEqual("a"),
	}

	// Assert
	expected := args.Map{
		"same":    true,
		"diff":    false,
		"diffLen": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct -- various comparisons", actual)
}

func Test_StringsOnce_Sorted_Cached(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"b", "a"} })
	_ = so.Sorted()
	sorted := so.Sorted()

	// Act
	actual := args.Map{"first": sorted[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Sorted returns cached sorted -- second call", actual)
}

func Test_StringsOnce_RangesMap_Empty(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return []string{} })
	m := so.RangesMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangesMap returns empty -- no items", actual)
}

func Test_StringsOnce_Length_NilValues(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return nil })

	// Act
	actual := args.Map{"len": so.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil values", actual)
}

func Test_StringsOnce_HasAll_Missing(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })

	// Act
	actual := args.Map{
		"allPresent": so.HasAll("a", "b"),
		"oneMissing": so.HasAll("a", "c"),
	}

	// Assert
	expected := args.Map{
		"allPresent": true,
		"oneMissing": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAll returns correct -- present and missing", actual)
}

// ── MapStringStringOnce — uncovered branches ──

func Test_MapStringStringOnce_Strings_Cached(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"k": "v"} })
	_ = mo.Strings()
	s := mo.Strings()

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Strings returns cached -- second call", actual)
}

func Test_MapStringStringOnce_Strings_Empty(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })
	s := mo.Strings()

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Strings returns empty -- empty map", actual)
}

func Test_MapStringStringOnce_AllKeys_Empty(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })

	// Act
	actual := args.Map{"len": len(mo.AllKeys())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllKeys returns empty -- empty map", actual)
}

func Test_MapStringStringOnce_AllKeys_Cached(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })
	_ = mo.AllKeys()
	k := mo.AllKeys()

	// Act
	actual := args.Map{"len": len(k)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AllKeys returns cached -- second call", actual)
}

func Test_MapStringStringOnce_AllValues_Empty(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })

	// Act
	actual := args.Map{"len": len(mo.AllValues())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllValues returns empty -- empty map", actual)
}

func Test_MapStringStringOnce_AllValues_Cached(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })
	_ = mo.AllValues()
	v := mo.AllValues()

	// Act
	actual := args.Map{"len": len(v)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AllValues returns cached -- second call", actual)
}

func Test_MapStringStringOnce_AllKeysSorted_Empty(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })

	// Act
	actual := args.Map{"len": len(mo.AllKeysSorted())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllKeysSorted returns empty -- empty map", actual)
}

func Test_MapStringStringOnce_AllKeysSorted_Cached(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"b": "2", "a": "1"} })
	_ = mo.AllKeysSorted()
	k := mo.AllKeysSorted()

	// Act
	actual := args.Map{"first": k[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "AllKeysSorted returns cached sorted -- second call", actual)
}

func Test_MapStringStringOnce_AllValuesSorted_Empty(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })

	// Act
	actual := args.Map{"len": len(mo.AllValuesSorted())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllValuesSorted returns empty -- empty map", actual)
}

func Test_MapStringStringOnce_AllValuesSorted_Cached(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "z", "b": "a"} })
	_ = mo.AllValuesSorted()
	v := mo.AllValuesSorted()

	// Act
	actual := args.Map{"first": v[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "AllValuesSorted returns cached sorted -- second call", actual)
}

func Test_MapStringStringOnce_IsEqual_MissingKey_FromAnyOnceValueString(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1", "b": "2"} })

	// Act
	actual := args.Map{
		"exact":    mo.IsEqual(map[string]string{"a": "1", "b": "2"}),
		"missing":  mo.IsEqual(map[string]string{"a": "1", "c": "2"}),
		"diffVal":  mo.IsEqual(map[string]string{"a": "1", "b": "3"}),
		"diffLen":  mo.IsEqual(map[string]string{"a": "1"}),
	}

	// Assert
	expected := args.Map{
		"exact":    true,
		"missing":  false,
		"diffVal":  false,
		"diffLen":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct -- various comparisons", actual)
}

func Test_MapStringStringOnce_IsEqual_BothNil(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })

	// Act
	actual := args.Map{"isEqual": mo.IsEqual(nil)}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- both nil", actual)
}

func Test_MapStringStringOnce_IsEqual_OneNil(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })

	// Act
	actual := args.Map{"isEqual": mo.IsEqual(map[string]string{"a": "1"})}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- left nil right not", actual)
}

func Test_MapStringStringOnce_Length_NilValues(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })

	// Act
	actual := args.Map{"len": mo.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil map", actual)
}

func Test_MapStringStringOnce_HasAll_Missing(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1", "b": "2"} })

	// Act
	actual := args.Map{
		"allPresent": mo.HasAll("a", "b"),
		"oneMissing": mo.HasAll("a", "c"),
	}

	// Assert
	expected := args.Map{
		"allPresent": true,
		"oneMissing": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAll returns correct -- present and missing", actual)
}

func Test_MapStringStringOnce_IsMissing(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })

	// Act
	actual := args.Map{
		"missingB": mo.IsMissing("b"),
		"hasA":     !mo.IsMissing("a"),
	}

	// Assert
	expected := args.Map{
		"missingB": true,
		"hasA":     true,
	}
	expected.ShouldBeEqual(t, 0, "IsMissing returns correct -- present and absent", actual)
}

func Test_MapStringStringOnce_GetValueWithStatus(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"k": "v"} })
	val, has := mo.GetValueWithStatus("k")
	_, miss := mo.GetValueWithStatus("x")

	// Act
	actual := args.Map{
		"val": val,
		"has": has,
		"miss": miss,
	}

	// Assert
	expected := args.Map{
		"val": "v",
		"has": true,
		"miss": false,
	}
	expected.ShouldBeEqual(t, 0, "GetValueWithStatus returns correct -- hit and miss", actual)
}
