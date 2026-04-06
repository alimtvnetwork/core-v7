package coreoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
)

// ===== ErrorOnce coverage =====

func Test_I20_ErrorOnce_String_HasError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("test-err") })
	s := o.String()
	if s != "test-err" {
		t.Fatalf("expected 'test-err', got '%s'", s)
	}
}

func Test_I20_ErrorOnce_Message_NilError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	if o.Message() != "" {
		t.Fatal("expected empty message for nil error")
	}
}

func Test_I20_ErrorOnce_Message_HasError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("msg") })
	if o.Message() != "msg" {
		t.Fatal("expected 'msg'")
	}
}

func Test_I20_ErrorOnce_IsMessageEqual(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("hello") })
	if !o.IsMessageEqual("hello") {
		t.Fatal("expected true")
	}
	if o.IsMessageEqual("nope") {
		t.Fatal("expected false")
	}
}

func Test_I20_ErrorOnce_IsMessageEqual_Nil(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	if o.IsMessageEqual("anything") {
		t.Fatal("expected false for nil error")
	}
}

func Test_I20_ErrorOnce_HandleError_NoError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	// Should not panic
	o.HandleError()
}

func Test_I20_ErrorOnce_HandleError_Panic(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("boom") })
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.HandleError()
}

func Test_I20_ErrorOnce_HandleErrorWith_NoError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	o.HandleErrorWith("extra")
}

func Test_I20_ErrorOnce_HandleErrorWith_Panic(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("boom") })
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.HandleErrorWith("context")
}

func Test_I20_ErrorOnce_ConcatNewString_NoError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	s := o.ConcatNewString("a", "b")
	if s == "" {
		t.Fatal("expected non-empty concat")
	}
}

func Test_I20_ErrorOnce_ConcatNewString_HasError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("err") })
	s := o.ConcatNewString("extra")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I20_ErrorOnce_ConcatNew(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("base") })
	err := o.ConcatNew("more")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_ErrorOnce_MarshalJSON(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("test") })
	b, err := o.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if len(b) == 0 {
		t.Fatal("expected non-empty json")
	}
}

func Test_I20_ErrorOnce_MarshalJSON_NilError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	b, err := o.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `""` {
		t.Fatalf("expected empty string json, got %s", string(b))
	}
}

func Test_I20_ErrorOnce_UnmarshalJSON(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	err := o.UnmarshalJSON([]byte(`"test-error"`))
	if err != nil {
		t.Fatal(err)
	}
	if !o.IsMessageEqual("test-error") {
		t.Fatal("expected unmarshalled error message")
	}
}

func Test_I20_ErrorOnce_Predicates(t *testing.T) {
	oErr := coreonce.NewErrorOncePtr(func() error { return errors.New("e") })
	oNil := coreonce.NewErrorOncePtr(func() error { return nil })

	if !oErr.HasError() {
		t.Fatal("expected HasError true")
	}
	if oNil.HasError() {
		t.Fatal("expected HasError false")
	}
	if !oErr.IsInvalid() {
		t.Fatal("expected IsInvalid true")
	}
	if !oNil.IsValid() {
		t.Fatal("expected IsValid true")
	}
	if !oNil.IsSuccess() {
		t.Fatal("expected IsSuccess true")
	}
	if !oErr.IsFailed() {
		t.Fatal("expected IsFailed true")
	}
	if !oErr.IsDefined() {
		t.Fatal("expected IsDefined true")
	}
	if !oErr.HasAnyItem() {
		t.Fatal("expected HasAnyItem true")
	}
	if oNil.HasAnyItem() {
		t.Fatal("expected HasAnyItem false for nil error")
	}
}

func Test_I20_ErrorOnce_Execute(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("ex") })
	if o.Execute() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I20_ErrorOnce_Serialize(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("ser") })
	b, err := o.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

// ===== BytesErrorOnce additional coverage =====

func Test_I20_BytesErrorOnce_MarshalJSON(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`"hello"`), nil
	})
	b, err := o.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_BytesErrorOnce_SerializeMust_Success(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`"ok"`), nil
	})
	b := o.SerializeMust()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_BytesErrorOnce_SerializeMust_Panic(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("fail")
	})
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.SerializeMust()
}

func Test_I20_BytesErrorOnce_DeserializeMust_Success(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`"value"`), nil
	})
	var s string
	o.DeserializeMust(&s)
	if s != "value" {
		t.Fatalf("expected 'value', got '%s'", s)
	}
}

func Test_I20_BytesErrorOnce_DeserializeMust_Panic(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("err")
	})
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	var s string
	o.DeserializeMust(&s)
}

func Test_I20_BytesErrorOnce_Deserialize_UnmarshalError(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`not-valid-json`), nil
	})
	var s string
	err := o.Deserialize(&s)
	if err == nil {
		t.Fatal("expected unmarshal error")
	}
}

func Test_I20_BytesErrorOnce_MustHaveSafeItems_Success(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("data"), nil
	})
	o.MustHaveSafeItems()
}

func Test_I20_BytesErrorOnce_MustHaveSafeItems_PanicOnError(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("err")
	})
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.MustHaveSafeItems()
}

func Test_I20_BytesErrorOnce_MustHaveSafeItems_PanicOnEmpty(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.MustHaveSafeItems()
}

func Test_I20_BytesErrorOnce_MustBeEmptyError_NoError(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("ok"), nil
	})
	o.MustBeEmptyError()
}

func Test_I20_BytesErrorOnce_MustBeEmptyError_Panic(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("e")
	})
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.MustBeEmptyError()
}

func Test_I20_BytesErrorOnce_Predicates(t *testing.T) {
	oOk := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("data"), nil
	})
	oErr := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("e")
	})

	if !oOk.HasSafeItems() {
		t.Fatal("expected HasSafeItems true")
	}
	if oErr.HasSafeItems() {
		t.Fatal("expected HasSafeItems false")
	}
	if !oOk.IsValid() {
		t.Fatal("expected IsValid true")
	}
	if !oErr.IsInvalid() {
		t.Fatal("expected IsInvalid true")
	}
	if !oOk.IsSuccess() {
		t.Fatal("expected IsSuccess true")
	}
	if !oErr.IsFailed() {
		t.Fatal("expected IsFailed true")
	}
	if !oOk.IsDefined() {
		t.Fatal("expected IsDefined true")
	}
	if !oOk.HasAnyItem() {
		t.Fatal("expected HasAnyItem true")
	}
}

func Test_I20_BytesErrorOnce_IsEmptyBytes(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})
	if !o.IsEmptyBytes() {
		t.Fatal("expected IsEmptyBytes true")
	}
}

func Test_I20_BytesErrorOnce_ValueWithError(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("val"), nil
	})
	b, err := o.ValueWithError()
	if err != nil || len(b) == 0 {
		t.Fatal("expected value with no error")
	}
}

// ===== AnyOnce additional coverage =====

func Test_I20_AnyOnce_CastValueString(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "hello" })
	val, ok := o.CastValueString()
	if !ok || val != "hello" {
		t.Fatal("expected successful cast")
	}
}

func Test_I20_AnyOnce_CastValueString_Fail(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return 42 })
	_, ok := o.CastValueString()
	if ok {
		t.Fatal("expected failed cast")
	}
}

func Test_I20_AnyOnce_CastValueStrings(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return []string{"a", "b"} })
	val, ok := o.CastValueStrings()
	if !ok || len(val) != 2 {
		t.Fatal("expected successful cast")
	}
}

func Test_I20_AnyOnce_CastValueHashmapMap(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return map[string]string{"k": "v"} })
	val, ok := o.CastValueHashmapMap()
	if !ok || val["k"] != "v" {
		t.Fatal("expected successful cast")
	}
}

func Test_I20_AnyOnce_CastValueMapStringAnyMap(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return map[string]any{"k": 1} })
	val, ok := o.CastValueMapStringAnyMap()
	if !ok || val["k"] != 1 {
		t.Fatal("expected successful cast")
	}
}

func Test_I20_AnyOnce_CastValueBytes(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return []byte("hi") })
	val, ok := o.CastValueBytes()
	if !ok || string(val) != "hi" {
		t.Fatal("expected successful cast")
	}
}

func Test_I20_AnyOnce_Serialize_Success(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "test" })
	b, err := o.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("expected successful serialize")
	}
}

func Test_I20_AnyOnce_SerializeSkipExistingError(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "test" })
	b, err := o.SerializeSkipExistingError()
	if err != nil || len(b) == 0 {
		t.Fatal("expected success")
	}
}

func Test_I20_AnyOnce_SerializeMust_Success(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "val" })
	b := o.SerializeMust()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_AnyOnce_SerializeMust_Panic(t *testing.T) {
	ch := make(chan int)
	o := coreonce.NewAnyOncePtr(func() any { return ch })
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.SerializeMust()
}

func Test_I20_AnyOnce_ValueStringMust(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "abc" })
	s := o.ValueStringMust()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I20_AnyOnce_SafeString(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "abc" })
	s := o.SafeString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I20_AnyOnce_ValueString_Nil(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return nil })
	s := o.ValueString()
	if s == "" {
		t.Fatal("expected nil bracket string")
	}
}

func Test_I20_AnyOnce_ValueString_Cached(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "cached" })
	_ = o.ValueString()
	s2 := o.ValueString() // should hit cache
	if s2 == "" {
		t.Fatal("expected cached value")
	}
}

func Test_I20_AnyOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "  " })
	// String() returns formatted, not just spaces
	_ = o.IsStringEmptyOrWhitespace()
}

func Test_I20_AnyOnce_ValueOnly(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return 42 })
	if o.ValueOnly() != 42 {
		t.Fatal("expected 42")
	}
}

func Test_I20_AnyOnce_IsInitialized(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return 1 })
	if o.IsInitialized() {
		t.Fatal("expected not initialized")
	}
	o.Value()
	if !o.IsInitialized() {
		t.Fatal("expected initialized")
	}
}

// ===== AnyErrorOnce additional coverage =====

func Test_I20_AnyErrorOnce_ExecuteMust_Success(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "ok", nil })
	v := o.ExecuteMust()
	if v != "ok" {
		t.Fatal("expected 'ok'")
	}
}

func Test_I20_AnyErrorOnce_ExecuteMust_Panic(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.ExecuteMust()
}

func Test_I20_AnyErrorOnce_ValueMust_Success(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return 99, nil })
	if o.ValueMust() != 99 {
		t.Fatal("expected 99")
	}
}

func Test_I20_AnyErrorOnce_ValueMust_Panic(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.ValueMust()
}

func Test_I20_AnyErrorOnce_ValueStringMust_Success(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "val", nil })
	s := o.ValueStringMust()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I20_AnyErrorOnce_ValueStringMust_Panic(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("err") })
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.ValueStringMust()
}

func Test_I20_AnyErrorOnce_ValueString_Nil(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	s, err := o.ValueString()
	if err != nil {
		t.Fatal("expected no error")
	}
	if s == "" {
		t.Fatal("expected nil bracket")
	}
}

func Test_I20_AnyErrorOnce_ValueString_Cached(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "x", nil })
	_, _ = o.ValueString()
	s2, _ := o.ValueString() // cached
	if s2 == "" {
		t.Fatal("expected cached")
	}
}

func Test_I20_AnyErrorOnce_CastValueString(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hi", nil })
	val, err, ok := o.CastValueString()
	if !ok || err != nil || val != "hi" {
		t.Fatal("expected success")
	}
}

func Test_I20_AnyErrorOnce_CastValueStrings(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []string{"a"}, nil })
	val, err, ok := o.CastValueStrings()
	if !ok || err != nil || len(val) != 1 {
		t.Fatal("expected success")
	}
}

func Test_I20_AnyErrorOnce_CastValueHashmapMap(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]string{"k": "v"}, nil })
	val, err, ok := o.CastValueHashmapMap()
	if !ok || err != nil || val["k"] != "v" {
		t.Fatal("expected success")
	}
}

func Test_I20_AnyErrorOnce_CastValueMapStringAnyMap(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]any{"k": 1}, nil })
	val, err, ok := o.CastValueMapStringAnyMap()
	if !ok || err != nil || val["k"] != 1 {
		t.Fatal("expected success")
	}
}

func Test_I20_AnyErrorOnce_CastValueBytes(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []byte("b"), nil })
	val, err, ok := o.CastValueBytes()
	if !ok || err != nil || string(val) != "b" {
		t.Fatal("expected success")
	}
}

func Test_I20_AnyErrorOnce_SerializeSkipExistingError(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "v", nil })
	b, err := o.SerializeSkipExistingError()
	if err != nil || len(b) == 0 {
		t.Fatal("expected success")
	}
}

func Test_I20_AnyErrorOnce_SerializeMust_Success(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "v", nil })
	b := o.SerializeMust()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I20_AnyErrorOnce_SerializeMust_Panic(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.SerializeMust()
}

func Test_I20_AnyErrorOnce_Serialize_MarshalError(t *testing.T) {
	ch := make(chan int)
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return ch, nil })
	_, err := o.Serialize()
	if err == nil {
		t.Fatal("expected marshal error")
	}
}

func Test_I20_AnyErrorOnce_ValueOnly_Initialized(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "v", nil })
	o.Value() // initialize
	v := o.ValueOnly()
	if v != "v" {
		t.Fatal("expected 'v'")
	}
}

func Test_I20_AnyErrorOnce_IsStringEmpty(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	if !o.IsStringEmpty() {
		t.Fatal("expected true")
	}
}

func Test_I20_AnyErrorOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	if !o.IsStringEmptyOrWhitespace() {
		t.Fatal("expected true")
	}
}

func Test_I20_AnyErrorOnce_Error_AlreadyInitialized(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	o.Value() // initialize
	err := o.Error()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I20_AnyErrorOnce_IsEmpty_NilValue(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	if !o.IsEmpty() {
		t.Fatal("expected empty")
	}
}

// ===== IntegerOnce additional coverage =====

func Test_I20_IntegerOnce_Comparisons(t *testing.T) {
	o := coreonce.NewIntegerOncePtr(func() int { return 5 })
	if !o.IsAbove(3) {
		t.Fatal("expected true")
	}
	if !o.IsAboveEqual(5) {
		t.Fatal("expected true")
	}
	if !o.IsLessThan(10) {
		t.Fatal("expected true")
	}
	if !o.IsLessThanEqual(5) {
		t.Fatal("expected true")
	}
	if !o.IsAboveZero() {
		t.Fatal("expected true")
	}
	if !o.IsAboveEqualZero() {
		t.Fatal("expected true")
	}
	if o.IsLessThanZero() {
		t.Fatal("expected false")
	}
	if o.IsLessThanEqualZero() {
		t.Fatal("expected false")
	}
	if !o.IsPositive() {
		t.Fatal("expected true")
	}
	if o.IsNegative() {
		t.Fatal("expected false")
	}
	if !o.IsValidIndex() {
		t.Fatal("expected true")
	}
	if o.IsInvalidIndex() {
		t.Fatal("expected false")
	}
}

func Test_I20_IntegerOnce_NegativeComparisons(t *testing.T) {
	o := coreonce.NewIntegerOncePtr(func() int { return -1 })
	if !o.IsNegative() {
		t.Fatal("expected true")
	}
	if !o.IsInvalidIndex() {
		t.Fatal("expected true")
	}
	if !o.IsLessThanZero() {
		t.Fatal("expected true")
	}
	if !o.IsLessThanEqualZero() {
		t.Fatal("expected true")
	}
}

// ===== ByteOnce additional coverage =====

func Test_I20_ByteOnce_Methods(t *testing.T) {
	o := coreonce.NewByteOncePtr(func() byte { return 5 })
	if o.Int() != 5 {
		t.Fatal("expected 5")
	}
	if !o.IsPositive() {
		t.Fatal("expected positive")
	}
	if o.IsNegative() {
		t.Fatal("byte is unsigned, should not be negative")
	}
	if o.IsZero() {
		t.Fatal("expected non-zero")
	}
	if o.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

func Test_I20_ByteOnce_Zero(t *testing.T) {
	o := coreonce.NewByteOncePtr(func() byte { return 0 })
	if !o.IsZero() {
		t.Fatal("expected zero")
	}
	if !o.IsEmpty() {
		t.Fatal("expected empty")
	}
}

// ===== BytesOnce additional coverage =====

func Test_I20_BytesOnce_NilInitializer(t *testing.T) {
	o := &coreonce.BytesOnce{}
	// initializerFunc is nil
	if !o.IsEmpty() {
		t.Fatal("expected empty")
	}
	if o.Length() != 0 {
		t.Fatal("expected 0 length")
	}
}

// ===== StringOnce additional coverage =====

func Test_I20_StringOnce_SplitLeftRightTrim(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return " left : right " })
	l, r := o.SplitLeftRightTrim(":")
	if l != "left" || r != "right" {
		t.Fatalf("expected 'left','right', got '%s','%s'", l, r)
	}
}

func Test_I20_StringOnce_HasPrefix(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "hello-world" })
	if !o.HasPrefix("hello") {
		t.Fatal("expected true")
	}
	if !o.IsStartsWith("hello") {
		t.Fatal("expected true")
	}
}

func Test_I20_StringOnce_HasSuffix(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "hello-world" })
	if !o.HasSuffix("world") {
		t.Fatal("expected true")
	}
	if !o.IsEndsWith("world") {
		t.Fatal("expected true")
	}
}

func Test_I20_StringOnce_IsContains(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "hello-world" })
	if !o.IsContains("lo-wo") {
		t.Fatal("expected true")
	}
}

func Test_I20_StringOnce_IsEmptyOrWhitespace(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "  " })
	if !o.IsEmptyOrWhitespace() {
		t.Fatal("expected true")
	}
}

func Test_I20_StringOnce_Bytes(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "abc" })
	if string(o.Bytes()) != "abc" {
		t.Fatal("expected 'abc'")
	}
}

func Test_I20_StringOnce_Error(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "err-msg" })
	if o.Error().Error() != "err-msg" {
		t.Fatal("expected 'err-msg'")
	}
}

func Test_I20_StringOnce_SplitBy(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "a,b,c" })
	parts := o.SplitBy(",")
	if len(parts) != 3 {
		t.Fatal("expected 3 parts")
	}
}

func Test_I20_StringOnce_ValuePtr(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "ptr" })
	p := o.ValuePtr()
	if *p != "ptr" {
		t.Fatal("expected 'ptr'")
	}
}

// ===== MapStringStringOnce additional coverage =====

func Test_I20_MapStringStringOnce_AllValuesSorted(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"b": "z", "a": "y"}
	})
	vs := o.AllValuesSorted()
	if len(vs) != 2 {
		t.Fatal("expected 2")
	}
	// call again to hit cache
	vs2 := o.AllValuesSorted()
	if len(vs2) != 2 {
		t.Fatal("expected cached 2")
	}
}

func Test_I20_MapStringStringOnce_GetValueWithStatus(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	val, has := o.GetValueWithStatus("k")
	if !has || val != "v" {
		t.Fatal("expected found")
	}
	_, has2 := o.GetValueWithStatus("missing")
	if has2 {
		t.Fatal("expected not found")
	}
}

func Test_I20_MapStringStringOnce_ValuesPtr(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	p := o.ValuesPtr()
	if p == nil {
		t.Fatal("expected non-nil ptr")
	}
}

func Test_I20_MapStringStringOnce_Strings_Cached(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	_ = o.Strings()
	s2 := o.Strings() // cached
	if len(s2) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_I20_MapStringStringOnce_IsMissing(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	if !o.IsMissing("nope") {
		t.Fatal("expected true")
	}
	if o.IsMissing("k") {
		t.Fatal("expected false")
	}
}

func Test_I20_MapStringStringOnce_String(t *testing.T) {
	o := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	s := o.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ===== StringsOnce additional coverage =====

func Test_I20_StringsOnce_HasAll_Missing(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	if o.HasAll("a", "c") {
		t.Fatal("expected false")
	}
}

func Test_I20_StringsOnce_UniqueMapLock(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	m := o.UniqueMapLock()
	if len(m) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_I20_StringsOnce_CsvLines(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	lines := o.CsvLines()
	if len(lines) == 0 {
		t.Fatal("expected non-empty")
	}
}

func Test_I20_StringsOnce_RangesMap(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"x", "y"} })
	m := o.RangesMap()
	if len(m) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_I20_StringsOnce_SafeStrings(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a"} })
	if len(o.SafeStrings()) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_I20_StringsOnce_SafeStrings_Empty(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{} })
	if len(o.SafeStrings()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_I20_StringsOnce_String(t *testing.T) {
	o := coreonce.NewStringsOnce(func() []string { return []string{"a", "b"} })
	s := o.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I20_StringsOnce_Length_NilValues(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return nil })
	if o.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_I20_StringsOnce_UniqueMap_NilValues(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return nil })
	m := o.UniqueMap()
	if len(m) != 0 {
		t.Fatal("expected empty map")
	}
}

func Test_I20_StringsOnce_UniqueMap_Cached(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a"} })
	_ = o.UniqueMap()
	m2 := o.UniqueMap() // cached
	if len(m2) != 1 {
		t.Fatal("expected 1")
	}
}

// ===== IntegersOnce additional coverage =====

func Test_I20_IntegersOnce_RangesBoolMap(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2, 3} })
	m := o.RangesBoolMap()
	if len(m) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_I20_IntegersOnce_RangesBoolMap_Empty(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := o.RangesBoolMap()
	if len(m) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_I20_IntegersOnce_UniqueMap(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 1, 2} })
	m := o.UniqueMap()
	if len(m) != 2 {
		t.Fatal("expected 2 unique")
	}
}

func Test_I20_IntegersOnce_UniqueMap_Empty(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := o.UniqueMap()
	if len(m) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_I20_IntegersOnce_Sorted_Cached(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{3, 1, 2} })
	_ = o.Sorted()
	s2 := o.Sorted() // cached
	if s2[0] != 1 {
		t.Fatal("expected sorted")
	}
}

func Test_I20_IntegersOnce_RangesMap_Empty(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := o.RangesMap()
	if len(m) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_I20_IntegersOnce_Aliases(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })
	if len(o.Values()) != 2 {
		t.Fatal("expected 2")
	}
	if len(o.Integers()) != 2 {
		t.Fatal("expected 2")
	}
	if len(o.Slice()) != 2 {
		t.Fatal("expected 2")
	}
	if len(o.List()) != 2 {
		t.Fatal("expected 2")
	}
}

// ===== MapStringStringOnce.HasAll =====

func Test_I20_MapStringStringOnce_HasAll(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})
	if !o.HasAll("a", "b") {
		t.Fatal("expected true")
	}
	if o.HasAll("a", "c") {
		t.Fatal("expected false")
	}
}

// ===== MapStringStringOnce.AllKeys cached =====

func Test_I20_MapStringStringOnce_AllKeys_Cached(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	_ = o.AllKeys()
	k2 := o.AllKeys() // cached
	if len(k2) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_I20_MapStringStringOnce_AllValues_Cached(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	_ = o.AllValues()
	v2 := o.AllValues() // cached
	if len(v2) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_I20_MapStringStringOnce_AllKeysSorted_Cached(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"b": "2", "a": "1"}
	})
	_ = o.AllKeysSorted()
	k2 := o.AllKeysSorted() // cached
	if k2[0] != "a" {
		t.Fatal("expected sorted")
	}
}

// ===== NewAnyOnce / NewAnyErrorOnce (non-ptr constructors) =====

func Test_I20_AnyOnce_NonPtr(t *testing.T) {
	o := coreonce.NewAnyOnce(func() any { return "np" })
	if o.Value() != "np" {
		t.Fatal("expected 'np'")
	}
}

func Test_I20_AnyErrorOnce_NonPtr(t *testing.T) {
	o := coreonce.NewAnyErrorOnce(func() (any, error) { return "np", nil })
	v, err := o.Value()
	if err != nil || v != "np" {
		t.Fatal("expected 'np'")
	}
}

func Test_I20_ErrorOnce_NonPtr(t *testing.T) {
	o := coreonce.NewErrorOnce(func() error { return nil })
	if o.Value() != nil {
		t.Fatal("expected nil")
	}
}

func Test_I20_BytesErrorOnce_NonPtr(t *testing.T) {
	o := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("x"), nil })
	v, err := o.Value()
	if err != nil || string(v) != "x" {
		t.Fatal("expected 'x'")
	}
}

// ===== StringsOnce.Sorted cached =====

func Test_I20_StringsOnce_Sorted_Cached(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"c", "a", "b"} })
	_ = o.Sorted()
	s2 := o.Sorted() // cached
	if s2[0] != "a" {
		t.Fatal("expected sorted")
	}
}

// ===== IntegersOnce.IsEqual nil paths =====

func Test_I20_IntegersOnce_IsEqual_BothNil(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return nil })
	if !o.IsEqual() {
		t.Fatal("expected true - both nil/empty")
	}
}

func Test_I20_IntegersOnce_IsEqual_OneSideNil(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return nil })
	if o.IsEqual(1) {
		t.Fatal("expected false")
	}
}

func Test_I20_IntegersOnce_IsEqual_LengthMismatch(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })
	if o.IsEqual(1) {
		t.Fatal("expected false")
	}
}

// ===== StringsOnce.IsEqual nil paths =====

func Test_I20_StringsOnce_IsEqual_BothNil(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return nil })
	if !o.IsEqual() {
		t.Fatal("expected true")
	}
}

func Test_I20_StringsOnce_IsEqual_OneSideNil(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return nil })
	if o.IsEqual("a") {
		t.Fatal("expected false")
	}
}

func Test_I20_StringsOnce_IsEqual_LengthMismatch(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	if o.IsEqual("a") {
		t.Fatal("expected false")
	}
}

// ===== MapStringStringOnce.IsEqual nil paths =====

func Test_I20_MapStringStringOnce_IsEqual_BothNil(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })
	if !o.IsEqual(nil) {
		t.Fatal("expected true")
	}
}

func Test_I20_MapStringStringOnce_IsEqual_OneSideNil(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })
	if o.IsEqual(map[string]string{"a": "1"}) {
		t.Fatal("expected false")
	}
}

func Test_I20_MapStringStringOnce_IsEqual_LengthMismatch(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})
	if o.IsEqual(map[string]string{"a": "1"}) {
		t.Fatal("expected false")
	}
}

// ===== BytesErrorOnce.IsEmpty =====

func Test_I20_BytesErrorOnce_IsEmpty_True(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })
	if !o.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I20_BytesErrorOnce_IsStringEmpty(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })
	if !o.IsStringEmpty() {
		t.Fatal("expected empty string")
	}
}

func Test_I20_BytesErrorOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("  "), nil })
	if !o.IsStringEmptyOrWhitespace() {
		t.Fatal("expected whitespace-empty")
	}
}

func Test_I20_BytesErrorOnce_HandleError_NoError(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("ok"), nil })
	o.HandleError()
}

func Test_I20_BytesErrorOnce_HandleError_Panic(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("e") })
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	o.HandleError()
}

// ===== AnyOnce.IsStringEmpty & IsNull =====

func Test_I20_AnyOnce_IsStringEmpty(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return nil })
	if !o.IsStringEmpty() {
		t.Fatal("expected true")
	}
}

func Test_I20_AnyOnce_IsNull(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return nil })
	if !o.IsNull() {
		t.Fatal("expected true")
	}
}

// ===== Bug-fix tests: Deserialize unmarshal error paths (previously unreachable) =====

func Test_I20_AnyOnce_Deserialize_UnmarshalError(t *testing.T) {
	// Value is valid JSON ("hello") but cannot unmarshal into *int
	o := coreonce.NewAnyOncePtr(func() any { return "hello" })
	var result int
	err := o.Deserialize(&result)
	if err == nil {
		t.Fatal("expected unmarshal error, got nil")
	}
	expected := "deserializing failed:"
	if len(err.Error()) < len(expected) {
		t.Fatalf("error too short: %s", err.Error())
	}
	if err.Error()[:len(expected)] != expected {
		t.Fatalf("expected error starting with '%s', got '%s'", expected, err.Error())
	}
}

func Test_I20_AnyOnce_Deserialize_UnmarshalError_NilToPtr(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return map[string]any{"a": 1} })
	err := o.Deserialize(nil)
	if err == nil {
		t.Fatal("expected unmarshal error for nil toPtr")
	}
}

func Test_I20_AnyErrorOnce_Deserialize_UnmarshalError(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hello", nil })
	var result int
	err := o.Deserialize(&result)
	if err == nil {
		t.Fatal("expected unmarshal error, got nil")
	}
	expected := "deserializing failed:"
	if err.Error()[:len(expected)] != expected {
		t.Fatalf("expected error starting with '%s', got '%s'", expected, err.Error())
	}
}

func Test_I20_AnyErrorOnce_Deserialize_UnmarshalError_NilToPtr(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]any{"a": 1}, nil })
	err := o.Deserialize(nil)
	if err == nil {
		t.Fatal("expected unmarshal error for nil toPtr")
	}
}
