package coreoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// AnyOnce — remaining uncovered branches
// ==========================================================================

// SerializeMust panic path: unmarshallable value triggers panic
func Test_AnyOnce_SerializeMust_Panic(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return func() {} })
	panicked := callPanics(func() { ao.SerializeMust() })

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "SerializeMust panics -- unmarshallable func value", actual)
}

// Deserialize when Serialize fails (value is func → Marshal error → early return)
func Test_AnyOnce_Deserialize_SerializeError_FromAnyOnceSerializeMust(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return func() {} })
	var result string
	err := ao.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- Serialize fails on func", actual)
}

// String() on non-nil value — covers the else branch in String()
func Test_AnyOnce_String_NonNil(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return 42 })

	// Act
	actual := args.Map{"notEmpty": ao.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns non-empty -- int value", actual)
}

// IsStringEmptyOrWhitespace on non-nil non-whitespace value
func Test_AnyOnce_IsStringEmptyOrWhitespace_NonEmpty(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "hello" })

	// Act
	actual := args.Map{"isEmpty": ao.IsStringEmptyOrWhitespace()}

	// Assert
	expected := args.Map{"isEmpty": false}
	expected.ShouldBeEqual(t, 0, "IsStringEmptyOrWhitespace returns false -- non-empty value", actual)
}

// ==========================================================================
// AnyErrorOnce — remaining uncovered branches
// ==========================================================================

// Error() when already initialized — covers the isInitialized early return (line 37-39)
func Test_AnyErrorOnce_Error_AlreadyInitialized(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "ok", nil })
	// First call initializes
	_, _ = aeo.Value()
	// Second call to Error() uses the isInitialized path
	err := aeo.Error()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Error returns nil -- already initialized no error", actual)
}

// Error() already initialized with error
func Test_AnyErrorOnce_Error_AlreadyInitialized_WithError(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("fail") })
	_, _ = aeo.Value()
	err := aeo.Error()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Error returns error -- already initialized with error", actual)
}

// ValueOnly() when already initialized — covers the isInitialized early return (line 246-248)
func Test_AnyErrorOnce_ValueOnly_AlreadyInitialized(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "cached", nil })
	_, _ = aeo.Value()
	val := aeo.ValueOnly()

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "cached"}
	expected.ShouldBeEqual(t, 0, "ValueOnly returns cached -- already initialized", actual)
}

// Deserialize when value has existing error — covers early return at Serialize
func Test_AnyErrorOnce_Deserialize_ExistingError_FromAnyOnceSerializeMust(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("pre") })
	var result string
	err := aeo.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- existing error from Serialize", actual)
}

// Deserialize with marshal error (value is func)
func Test_AnyErrorOnce_Deserialize_MarshalError_FromAnyOnceSerializeMust(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return func() {}, nil })
	var result string
	err := aeo.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- unmarshallable func value", actual)
}

// String() with nil value — covers the IsNull early return
func Test_AnyErrorOnce_String_Nil(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })

	// Act
	actual := args.Map{"isEmpty": aeo.String() == ""}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns empty -- nil value", actual)
}

// ==========================================================================
// ErrorOnce — remaining uncovered branches
// ==========================================================================

// HandleError with error — panics
func Test_ErrorOnce_HandleError_Panic(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return errors.New("boom") })
	panicked := callPanics(func() { eo.HandleError() })

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleError panics -- has error", actual)
}

// String() on nil error — panics because it.Value().Error() on nil
func Test_ErrorOnce_String_NilError_Panics(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return nil })
	panicked := callPanics(func() { _ = eo.String() })

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "String panics -- nil error calling Error()", actual)
}

// ErrorOnce.MarshalJSON with error value — covers the non-nil path
func Test_ErrorOnce_MarshalJSON_WithError(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return errors.New("err-msg") })
	mb, err := eo.MarshalJSON()

	// Act
	actual := args.Map{
		"hasBytes": len(mb) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns error message JSON -- has error", actual)
}

// ErrorOnce.Serialize with error value
func Test_ErrorOnce_Serialize_WithError(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return errors.New("err-val") })
	b, err := eo.Serialize()

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
	expected.ShouldBeEqual(t, 0, "Serialize returns bytes -- error value serializes", actual)
}

// ==========================================================================
// BytesErrorOnce — remaining uncovered branches
// ==========================================================================

// MustBeEmptyError with no error — should not panic
func Test_BytesErrorOnce_MustBeEmptyError_NoError(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("ok"), nil })
	panicked := callPanics(func() { beo.MustBeEmptyError() })

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError does not panic -- no error", actual)
}

// MustBeEmptyError with error — panics
func Test_BytesErrorOnce_MustBeEmptyError_Panic(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("e") })
	panicked := callPanics(func() { beo.MustBeEmptyError() })

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError panics -- has error", actual)
}

// HandleError with no error — should not panic
func Test_BytesErrorOnce_HandleError_NoError(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("ok"), nil })
	panicked := callPanics(func() { beo.HandleError() })

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "HandleError does not panic -- no error", actual)
}

// HandleError with error — panics
func Test_BytesErrorOnce_HandleError_Panic(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("e") })
	panicked := callPanics(func() { beo.HandleError() })

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleError panics -- has error", actual)
}

// Error() when already initialized — covers isInitialized early return (line 63-65)
func Test_BytesErrorOnce_Error_AlreadyInitialized(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("ok"), nil })
	_, _ = beo.Value()
	err := beo.Error()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Error returns nil -- already initialized no error", actual)
}

// ValueOnly() when already initialized — covers isInitialized early return (line 221-222)
func Test_BytesErrorOnce_ValueOnly_AlreadyInitialized(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("data"), nil })
	_, _ = beo.Value()
	val := beo.ValueOnly()

	// Act
	actual := args.Map{"len": len(val)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ValueOnly returns cached -- already initialized", actual)
}

// MustHaveSafeItems with valid data — should not panic
func Test_BytesErrorOnce_MustHaveSafeItems_NoError(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("ok"), nil })
	panicked := callPanics(func() { beo.MustHaveSafeItems() })

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "MustHaveSafeItems does not panic -- has bytes no error", actual)
}

// DeserializeMust with success — no panic
func Test_BytesErrorOnce_DeserializeMust_NoError(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte(`"hello"`), nil })
	var result string
	panicked := callPanics(func() { beo.DeserializeMust(&result) })

	// Act
	actual := args.Map{
		"panicked": panicked,
		"val": result,
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "DeserializeMust does not panic -- valid json", actual)
}

// Deserialize with nil toPtr — covers typeNameString empty path
func Test_BytesErrorOnce_Deserialize_NilToPtr(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("not-json"), nil })
	err := beo.Deserialize(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- nil toPtr with invalid json", actual)
}

// Deserialize with existing error and nil toPtr — covers typeNameString nil path
func Test_BytesErrorOnce_Deserialize_ExistingError_NilToPtr(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("fail") })
	err := beo.Deserialize(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- existing error nil toPtr", actual)
}

// Deserialize with existing error and non-nil data — covers valString path
func Test_BytesErrorOnce_Deserialize_ExistingError_WithData(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("some-data"), errors.New("fail") })
	var result string
	err := beo.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- existing error with data bytes", actual)
}

// SerializeMust with no error — no panic
func Test_BytesErrorOnce_SerializeMust_NoError(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("ok"), nil })
	var result []byte
	panicked := callPanics(func() { result = beo.SerializeMust() })

	// Act
	actual := args.Map{
		"panicked": panicked,
		"hasBytes": len(result) > 0,
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializeMust does not panic -- valid bytes", actual)
}

// IsEmpty on nil receiver scenario — covers it == nil check
func Test_BytesErrorOnce_IsEmpty_NilData(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })

	// Act
	actual := args.Map{"isEmpty": beo.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- nil data nil error", actual)
}

// HasIssuesOrEmpty with error — covers err != nil path
func Test_BytesErrorOnce_HasIssuesOrEmpty_WithError(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("data"), errors.New("e") })

	// Act
	actual := args.Map{"hasIssues": beo.HasIssuesOrEmpty()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "HasIssuesOrEmpty returns true -- has error despite data", actual)
}

// ==========================================================================
// StringsOnce — remaining uncovered branches
// ==========================================================================

// IsEmpty with nil initializerFunc — covers the initializerFunc == nil path
func Test_StringsOnce_IsEmpty_NilInitializerFunc(t *testing.T) {
	// Arrange
	so := &coreonce.StringsOnce{}

	// Act
	actual := args.Map{"isEmpty": so.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- nil initializerFunc", actual)
}

// ==========================================================================
// MapStringStringOnce — remaining uncovered branches
// ==========================================================================

// IsEmpty with nil initializerFunc — covers the initializerFunc == nil path
func Test_MapStringStringOnce_IsEmpty_NilInitializerFunc(t *testing.T) {
	// Arrange
	mo := &coreonce.MapStringStringOnce{}

	// Act
	actual := args.Map{"isEmpty": mo.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- nil initializerFunc", actual)
}

// ==========================================================================
// IntegersOnce — remaining uncovered branches
// ==========================================================================

// IsEqual with nil items vs nil items — covers both-nil path
func Test_IntegersOnce_IsEqual_BothNil(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return nil })

	// Act
	actual := args.Map{"isEqual": io.IsEqual(nil...)}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- both nil", actual)
}

// ==========================================================================
// StringsOnce — IsEqual nil branches
// ==========================================================================

// IsEqual with nil values and nil comparison
func Test_StringsOnce_IsEqual_BothNil(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return nil })

	// Act
	actual := args.Map{"isEqual": so.IsEqual(nil...)}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- both nil", actual)
}

// IsEqual with nil values and non-nil comparison — covers currentItems == nil path
func Test_StringsOnce_IsEqual_NilVsNonNil(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return nil })

	// Act
	actual := args.Map{"isEqual": so.IsEqual("a")}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- nil vs non-nil", actual)
}

// ==========================================================================
// IntegersOnce — IsEqual nil vs non-nil
// ==========================================================================

func Test_IntegersOnce_IsEqual_NilVsNonNil(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return nil })

	// Act
	actual := args.Map{"isEqual": io.IsEqual(1)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- nil vs non-nil", actual)
}

// ==========================================================================
// MapStringStringOnce — IsEqual nil edge cases
// ==========================================================================

// IsEqual right nil, left non-nil — covers currentItems non-nil, rightMap nil
func Test_MapStringStringOnce_IsEqual_LeftNonNilRightNil(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })

	// Act
	actual := args.Map{"isEqual": mo.IsEqual(nil)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- non-nil vs nil", actual)
}

// ==========================================================================
// AnyErrorOnce — Deserialize dead code exploration
// ==========================================================================

// Deserialize with nil toPtr — covers the toPtr == nil → typeSafeName empty path
func Test_AnyErrorOnce_Deserialize_NilToPtr(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hello", nil })
	err := aeo.Deserialize(nil)
	// Deserialize: Serialize succeeds, then if err == nil → return err (nil)
	// The Unmarshal to nil may fail but the code returns err (nil) due to the bug

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": false}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- nil toPtr", actual)
}

// ==========================================================================
// BytesOnce — IsEmpty nil pointer check
// ==========================================================================

func Test_BytesOnce_IsEmpty_NilPtr(t *testing.T) {
	// Arrange
	var bo *coreonce.BytesOnce

	// Act
	actual := args.Map{"isEmpty": bo.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- nil pointer", actual)
}

func Test_BytesOnce_Length_NilPtr(t *testing.T) {
	// Arrange
	var bo *coreonce.BytesOnce

	// Act
	actual := args.Map{"length": bo.Length()}

	// Assert
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil pointer", actual)
}

// BytesOnce nil func with Value() call — covers initializerFunc == nil path
func Test_BytesOnce_Value_NilFunc(t *testing.T) {
	// Arrange
	bo := &coreonce.BytesOnce{}
	val := bo.Value()

	// Act
	actual := args.Map{"len": len(val)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Value returns empty -- nil initializerFunc", actual)
}
