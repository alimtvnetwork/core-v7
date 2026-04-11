package coreoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== ErrorOnce coverage =====

func Test_ErrorOnce_String_HasError(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("test-err") })
	s := o.String()

	// Act
	actual := args.Map{"result": s != "test-err"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'test-err', got ''", actual)
}

func Test_ErrorOnce_Message_NilError(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return nil })

	// Act
	actual := args.Map{"result": o.Message() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty message for nil error", actual)
}

func Test_ErrorOnce_Message_HasError(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("msg") })

	// Act
	actual := args.Map{"result": o.Message() != "msg"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'msg'", actual)
}

func Test_ErrorOnce_IsMessageEqual_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("hello") })

	// Act
	actual := args.Map{"result": o.IsMessageEqual("hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsMessageEqual("nope")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_ErrorOnce_IsMessageEqual_Nil(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return nil })

	// Act
	actual := args.Map{"result": o.IsMessageEqual("anything")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil error", actual)
}

func Test_ErrorOnce_HandleError_NoError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	// Should not panic
	o.HandleError()
}

func Test_ErrorOnce_HandleError_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("boom") })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.HandleError()
}

func Test_ErrorOnce_HandleErrorWith_NoError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	o.HandleErrorWith("extra")
}

func Test_ErrorOnce_HandleErrorWith_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("boom") })
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.HandleErrorWith("context")
}

func Test_ErrorOnce_ConcatNewString_NoError(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	s := o.ConcatNewString("a", "b")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty concat", actual)
}

func Test_ErrorOnce_ConcatNewString_HasError(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("err") })
	s := o.ConcatNewString("extra")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_ErrorOnce_ConcatNew_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("base") })
	err := o.ConcatNew("more")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_ErrorOnce_MarshalJSON(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("test") })
	b, err := o.MarshalJSON()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": len(b) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty json", actual)
}

func Test_ErrorOnce_MarshalJSON_NilError(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	b, err := o.MarshalJSON()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": string(b) != `""`}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string json", actual)
}

func Test_ErrorOnce_UnmarshalJSON(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	err := o.UnmarshalJSON([]byte(`"test-error"`))

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": o.IsMessageEqual("test-error")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected unmarshalled error message", actual)
}

func Test_ErrorOnce_Predicates(t *testing.T) {
	// Arrange
	oErr := coreonce.NewErrorOncePtr(func() error { return errors.New("e") })
	oNil := coreonce.NewErrorOncePtr(func() error { return nil })

	// Act
	actual := args.Map{"result": oErr.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasError true", actual)
	actual = args.Map{"result": oNil.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasError false", actual)
	actual = args.Map{"result": oErr.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid true", actual)
	actual = args.Map{"result": oNil.IsValid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsValid true", actual)
	actual = args.Map{"result": oNil.IsSuccess()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsSuccess true", actual)
	actual = args.Map{"result": oErr.IsFailed()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsFailed true", actual)
	actual = args.Map{"result": oErr.IsDefined()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsDefined true", actual)
	actual = args.Map{"result": oErr.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem true", actual)
	actual = args.Map{"result": oNil.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem false for nil error", actual)
}

func Test_ErrorOnce_Execute(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("ex") })

	// Act
	actual := args.Map{"result": o.Execute() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_ErrorOnce_Serialize(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("ser") })
	b, err := o.Serialize()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": len(b) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

// ===== BytesErrorOnce additional coverage =====

func Test_BytesErrorOnce_MarshalJSON_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`"hello"`), nil
	})
	b, err := o.MarshalJSON()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual = args.Map{"result": len(b) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_BytesErrorOnce_SerializeMust_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`"ok"`), nil
	})
	b := o.SerializeMust()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_BytesErrorOnce_SerializeMust_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("fail")
	})
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.SerializeMust()
}

func Test_BytesErrorOnce_DeserializeMust_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`"value"`), nil
	})
	var s string
	o.DeserializeMust(&s)

	// Act
	actual := args.Map{"result": s != "value"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'value', got ''", actual)
}

func Test_BytesErrorOnce_DeserializeMust_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("err")
	})
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	var s string
	o.DeserializeMust(&s)
}

func Test_BytesErrorOnce_Deserialize_UnmarshalError_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`not-valid-json`), nil
	})
	var s string
	err := o.Deserialize(&s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error", actual)
}

func Test_BytesErrorOnce_MustHaveSafeItems_Success(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("data"), nil
	})
	o.MustHaveSafeItems()
}

func Test_BytesErrorOnce_MustHaveSafeItems_PanicOnError_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("err")
	})
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.MustHaveSafeItems()
}

func Test_BytesErrorOnce_MustHaveSafeItems_PanicOnEmpty_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.MustHaveSafeItems()
}

func Test_BytesErrorOnce_MustBeEmptyError_NoError_FromErrorOnceStringItera(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("ok"), nil
	})
	o.MustBeEmptyError()
}

func Test_BytesErrorOnce_MustBeEmptyError_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("e")
	})
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.MustBeEmptyError()
}

func Test_BytesErrorOnce_Predicates(t *testing.T) {
	// Arrange
	oOk := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("data"), nil
	})
	oErr := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("e")
	})

	// Act
	actual := args.Map{"result": oOk.HasSafeItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasSafeItems true", actual)
	actual = args.Map{"result": oErr.HasSafeItems()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasSafeItems false", actual)
	actual = args.Map{"result": oOk.IsValid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsValid true", actual)
	actual = args.Map{"result": oErr.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid true", actual)
	actual = args.Map{"result": oOk.IsSuccess()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsSuccess true", actual)
	actual = args.Map{"result": oErr.IsFailed()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsFailed true", actual)
	actual = args.Map{"result": oOk.IsDefined()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsDefined true", actual)
	actual = args.Map{"result": oOk.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem true", actual)
}

func Test_BytesErrorOnce_IsEmptyBytes_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})

	// Act
	actual := args.Map{"result": o.IsEmptyBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEmptyBytes true", actual)
}

func Test_BytesErrorOnce_ValueWithError_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("val"), nil
	})
	b, err := o.ValueWithError()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected value with no error", actual)
}

// ===== AnyOnce additional coverage =====

func Test_AnyOnce_CastValueString_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return "hello" })
	val, ok := o.CastValueString()

	// Act
	actual := args.Map{"result": ok || val != "hello"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_AnyOnce_CastValueString_Fail_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return 42 })
	_, ok := o.CastValueString()

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected failed cast", actual)
}

func Test_AnyOnce_CastValueStrings_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return []string{"a", "b"} })
	val, ok := o.CastValueStrings()

	// Act
	actual := args.Map{"result": ok || len(val) != 2}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_AnyOnce_CastValueHashmapMap_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return map[string]string{"k": "v"} })
	val, ok := o.CastValueHashmapMap()

	// Act
	actual := args.Map{"result": ok || val["k"] != "v"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_AnyOnce_CastValueMapStringAnyMap_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return map[string]any{"k": 1} })
	val, ok := o.CastValueMapStringAnyMap()

	// Act
	actual := args.Map{"result": ok || val["k"] != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_AnyOnce_CastValueBytes_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return []byte("hi") })
	val, ok := o.CastValueBytes()

	// Act
	actual := args.Map{"result": ok || string(val) != "hi"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_AnyOnce_Serialize_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return "test" })
	b, err := o.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected successful serialize", actual)
}

func Test_AnyOnce_SerializeSkipExistingError_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return "test" })
	b, err := o.SerializeSkipExistingError()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_AnyOnce_SerializeMust_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return "val" })
	b := o.SerializeMust()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_AnyOnce_SerializeMust_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	ch := make(chan int)
	o := coreonce.NewAnyOncePtr(func() any { return ch })
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.SerializeMust()
}

func Test_AnyOnce_ValueStringMust(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return "abc" })
	s := o.ValueStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_AnyOnce_SafeString(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return "abc" })
	s := o.SafeString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_AnyOnce_ValueString_Nil(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return nil })
	s := o.ValueString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil bracket string", actual)
}

func Test_AnyOnce_ValueString_Cached_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return "cached" })
	_ = o.ValueString()
	s2 := o.ValueString() // should hit cache

	// Act
	actual := args.Map{"result": s2 == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached value", actual)
}

func Test_AnyOnce_IsStringEmptyOrWhitespace_FromErrorOnceStringItera(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "  " })
	// String() returns formatted, not just spaces
	_ = o.IsStringEmptyOrWhitespace()
}

func Test_AnyOnce_ValueOnly(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return 42 })

	// Act
	actual := args.Map{"result": o.ValueOnly() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_AnyOnce_IsInitialized(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return 1 })

	// Act
	actual := args.Map{"result": o.IsInitialized()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not initialized", actual)
	o.Value()
	actual = args.Map{"result": o.IsInitialized()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected initialized", actual)
}

// ===== AnyErrorOnce additional coverage =====

func Test_AnyErrorOnce_ExecuteMust_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "ok", nil })
	v := o.ExecuteMust()

	// Act
	actual := args.Map{"result": v != "ok"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'ok'", actual)
}

func Test_AnyErrorOnce_ExecuteMust_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.ExecuteMust()
}

func Test_AnyErrorOnce_ValueMust_Success_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return 99, nil })

	// Act
	actual := args.Map{"result": o.ValueMust() != 99}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
}

func Test_AnyErrorOnce_ValueMust_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.ValueMust()
}

func Test_AnyErrorOnce_ValueStringMust_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "val", nil })
	s := o.ValueStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_AnyErrorOnce_ValueStringMust_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("err") })
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.ValueStringMust()
}

func Test_AnyErrorOnce_ValueString_Nil_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	s, err := o.ValueString()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil bracket", actual)
}

func Test_AnyErrorOnce_ValueString_Cached_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "x", nil })
	_, _ = o.ValueString()
	s2, _ := o.ValueString() // cached

	// Act
	actual := args.Map{"result": s2 == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached", actual)
}

func Test_AnyErrorOnce_CastValueString_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hi", nil })
	val, err, ok := o.CastValueString()

	// Act
	actual := args.Map{"result": ok || err != nil || val != "hi"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_AnyErrorOnce_CastValueStrings_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []string{"a"}, nil })
	val, err, ok := o.CastValueStrings()

	// Act
	actual := args.Map{"result": ok || err != nil || len(val) != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_AnyErrorOnce_CastValueHashmapMap_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]string{"k": "v"}, nil })
	val, err, ok := o.CastValueHashmapMap()

	// Act
	actual := args.Map{"result": ok || err != nil || val["k"] != "v"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_AnyErrorOnce_CastValueMapStringAnyMap_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]any{"k": 1}, nil })
	val, err, ok := o.CastValueMapStringAnyMap()

	// Act
	actual := args.Map{"result": ok || err != nil || val["k"] != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_AnyErrorOnce_CastValueBytes_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []byte("b"), nil })
	val, err, ok := o.CastValueBytes()

	// Act
	actual := args.Map{"result": ok || err != nil || string(val) != "b"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_AnyErrorOnce_SerializeSkipExistingError_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "v", nil })
	b, err := o.SerializeSkipExistingError()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_AnyErrorOnce_SerializeMust_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "v", nil })
	b := o.SerializeMust()

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_AnyErrorOnce_SerializeMust_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.SerializeMust()
}

func Test_AnyErrorOnce_Serialize_MarshalError_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	ch := make(chan int)
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return ch, nil })
	_, err := o.Serialize()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected marshal error", actual)
}

func Test_AnyErrorOnce_ValueOnly_Initialized(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "v", nil })
	o.Value() // initialize
	v := o.ValueOnly()

	// Act
	actual := args.Map{"result": v != "v"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'v'", actual)
}

func Test_AnyErrorOnce_IsStringEmpty_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })

	// Act
	actual := args.Map{"result": o.IsStringEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_AnyErrorOnce_IsStringEmptyOrWhitespace_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })

	// Act
	actual := args.Map{"result": o.IsStringEmptyOrWhitespace()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_AnyErrorOnce_Error_AlreadyInitialized_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	o.Value() // initialize
	err := o.Error()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_AnyErrorOnce_IsEmpty_NilValue(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })

	// Act
	actual := args.Map{"result": o.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ===== IntegerOnce additional coverage =====

func Test_IntegerOnce_Comparisons(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegerOncePtr(func() int { return 5 })

	// Act
	actual := args.Map{"result": o.IsAbove(3)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsAboveEqual(5)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsLessThan(10)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsLessThanEqual(5)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsAboveZero()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsAboveEqualZero()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsLessThanZero()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": o.IsLessThanEqualZero()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": o.IsPositive()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsNegative()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": o.IsValidIndex()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsInvalidIndex()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IntegerOnce_NegativeComparisons(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegerOncePtr(func() int { return -1 })

	// Act
	actual := args.Map{"result": o.IsNegative()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsInvalidIndex()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsLessThanZero()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsLessThanEqualZero()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ===== ByteOnce additional coverage =====

func Test_ByteOnce_Methods(t *testing.T) {
	// Arrange
	o := coreonce.NewByteOncePtr(func() byte { return 5 })

	// Act
	actual := args.Map{"result": o.Int() != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual = args.Map{"result": o.IsPositive()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected positive", actual)
	actual = args.Map{"result": o.IsNegative()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "byte is unsigned, should not be negative", actual)
	actual = args.Map{"result": o.IsZero()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero", actual)
	actual = args.Map{"result": o.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_ByteOnce_Zero_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewByteOncePtr(func() byte { return 0 })

	// Act
	actual := args.Map{"result": o.IsZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected zero", actual)
	actual = args.Map{"result": o.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ===== BytesOnce additional coverage =====

func Test_BytesOnce_NilInitializer(t *testing.T) {
	// Arrange
	o := &coreonce.BytesOnce{}
	// initializerFunc is nil

	// Act
	actual := args.Map{"result": o.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": o.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 length", actual)
}

// ===== StringOnce additional coverage =====

func Test_StringOnce_SplitLeftRightTrim(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return " left : right " })
	l, r := o.SplitLeftRightTrim(":")

	// Act
	actual := args.Map{"result": l != "left" || r != "right"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'left','right', got '',''", actual)
}

func Test_StringOnce_HasPrefix(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "hello-world" })

	// Act
	actual := args.Map{"result": o.HasPrefix("hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsStartsWith("hello")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_StringOnce_HasSuffix(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "hello-world" })

	// Act
	actual := args.Map{"result": o.HasSuffix("world")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsEndsWith("world")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_StringOnce_IsContains(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "hello-world" })

	// Act
	actual := args.Map{"result": o.IsContains("lo-wo")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_StringOnce_IsEmptyOrWhitespace(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "  " })

	// Act
	actual := args.Map{"result": o.IsEmptyOrWhitespace()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_StringOnce_Bytes(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "abc" })

	// Act
	actual := args.Map{"result": string(o.Bytes()) != "abc"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'abc'", actual)
}

func Test_StringOnce_Error(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "err-msg" })

	// Act
	actual := args.Map{"result": o.Error().Error() != "err-msg"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'err-msg'", actual)
}

func Test_StringOnce_SplitBy(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "a,b,c" })
	parts := o.SplitBy(",")

	// Act
	actual := args.Map{"result": len(parts) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 parts", actual)
}

func Test_StringOnce_ValuePtr(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "ptr" })
	p := o.ValuePtr()

	// Act
	actual := args.Map{"result": *p != "ptr"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'ptr'", actual)
}

// ===== MapStringStringOnce additional coverage =====

func Test_MapStringStringOnce_AllValuesSorted(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"b": "z", "a": "y"}
	})
	vs := o.AllValuesSorted()

	// Act
	actual := args.Map{"result": len(vs) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	// call again to hit cache
	vs2 := o.AllValuesSorted()
	actual = args.Map{"result": len(vs2) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached 2", actual)
}

func Test_MapStringStringOnce_GetValueWithStatus_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	val, has := o.GetValueWithStatus("k")

	// Act
	actual := args.Map{"result": has || val != "v"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected found", actual)
	_, has2 := o.GetValueWithStatus("missing")
	actual = args.Map{"result": has2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_MapStringStringOnce_ValuesPtr(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	p := o.ValuesPtr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil ptr", actual)
}

func Test_MapStringStringOnce_Strings_Cached_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	_ = o.Strings()
	s2 := o.Strings() // cached

	// Act
	actual := args.Map{"result": len(s2) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_MapStringStringOnce_IsMissing_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})

	// Act
	actual := args.Map{"result": o.IsMissing("nope")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.IsMissing("k")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_MapStringStringOnce_String(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	s := o.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ===== StringsOnce additional coverage =====

func Test_StringsOnce_HasAll_Missing_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })

	// Act
	actual := args.Map{"result": o.HasAll("a", "c")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_StringsOnce_UniqueMapLock_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	m := o.UniqueMapLock()

	// Act
	actual := args.Map{"result": len(m) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_StringsOnce_CsvLines(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	lines := o.CsvLines()

	// Act
	actual := args.Map{"result": len(lines) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_StringsOnce_RangesMap(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"x", "y"} })
	m := o.RangesMap()

	// Act
	actual := args.Map{"result": len(m) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_StringsOnce_SafeStrings(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a"} })

	// Act
	actual := args.Map{"result": len(o.SafeStrings()) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_StringsOnce_SafeStrings_Empty(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{} })

	// Act
	actual := args.Map{"result": len(o.SafeStrings()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_StringsOnce_String(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOnce(func() []string { return []string{"a", "b"} })
	s := o.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_StringsOnce_Length_NilValues_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return nil })

	// Act
	actual := args.Map{"result": o.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_StringsOnce_UniqueMap_NilValues_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return nil })
	m := o.UniqueMap()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map", actual)
}

func Test_StringsOnce_UniqueMap_Cached_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a"} })
	_ = o.UniqueMap()
	m2 := o.UniqueMap() // cached

	// Act
	actual := args.Map{"result": len(m2) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ===== IntegersOnce additional coverage =====

func Test_IntegersOnce_RangesBoolMap_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2, 3} })
	m := o.RangesBoolMap()

	// Act
	actual := args.Map{"result": len(m) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_IntegersOnce_RangesBoolMap_Empty(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := o.RangesBoolMap()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_IntegersOnce_UniqueMap(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 1, 2} })
	m := o.UniqueMap()

	// Act
	actual := args.Map{"result": len(m) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 unique", actual)
}

func Test_IntegersOnce_UniqueMap_Empty(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := o.UniqueMap()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_IntegersOnce_Sorted_Cached_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{3, 1, 2} })
	_ = o.Sorted()
	s2 := o.Sorted() // cached

	// Act
	actual := args.Map{"result": s2[0] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_IntegersOnce_RangesMap_Empty_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := o.RangesMap()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_IntegersOnce_Aliases(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })

	// Act
	actual := args.Map{"result": len(o.Values()) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": len(o.Integers()) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": len(o.Slice()) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": len(o.List()) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ===== MapStringStringOnce.HasAll =====

func Test_MapStringStringOnce_HasAll(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})

	// Act
	actual := args.Map{"result": o.HasAll("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": o.HasAll("a", "c")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ===== MapStringStringOnce.AllKeys cached =====

func Test_MapStringStringOnce_AllKeys_Cached_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	_ = o.AllKeys()
	k2 := o.AllKeys() // cached

	// Act
	actual := args.Map{"result": len(k2) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_MapStringStringOnce_AllValues_Cached_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	_ = o.AllValues()
	v2 := o.AllValues() // cached

	// Act
	actual := args.Map{"result": len(v2) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_MapStringStringOnce_AllKeysSorted_Cached_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"b": "2", "a": "1"}
	})
	_ = o.AllKeysSorted()
	k2 := o.AllKeysSorted() // cached

	// Act
	actual := args.Map{"result": k2[0] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

// ===== NewAnyOnce / NewAnyErrorOnce (non-ptr constructors) =====

func Test_AnyOnce_NonPtr(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOnce(func() any { return "np" })

	// Act
	actual := args.Map{"result": o.Value() != "np"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'np'", actual)
}

func Test_AnyErrorOnce_NonPtr(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOnce(func() (any, error) { return "np", nil })
	v, err := o.Value()

	// Act
	actual := args.Map{"result": err != nil || v != "np"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'np'", actual)
}

func Test_ErrorOnce_NonPtr(t *testing.T) {
	// Arrange
	o := coreonce.NewErrorOnce(func() error { return nil })

	// Act
	actual := args.Map{"result": o.Value() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_BytesErrorOnce_NonPtr(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("x"), nil })
	v, err := o.Value()

	// Act
	actual := args.Map{"result": err != nil || string(v) != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
}

// ===== StringsOnce.Sorted cached =====

func Test_StringsOnce_Sorted_Cached_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"c", "a", "b"} })
	_ = o.Sorted()
	s2 := o.Sorted() // cached

	// Act
	actual := args.Map{"result": s2[0] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

// ===== IntegersOnce.IsEqual nil paths =====

func Test_IntegersOnce_IsEqual_BothNil_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return nil })

	// Act
	actual := args.Map{"result": o.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true - both nil/empty", actual)
}

func Test_IntegersOnce_IsEqual_OneSideNil(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return nil })

	// Act
	actual := args.Map{"result": o.IsEqual(1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IntegersOnce_IsEqual_LengthMismatch(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })

	// Act
	actual := args.Map{"result": o.IsEqual(1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ===== StringsOnce.IsEqual nil paths =====

func Test_StringsOnce_IsEqual_BothNil_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return nil })

	// Act
	actual := args.Map{"result": o.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_StringsOnce_IsEqual_OneSideNil(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return nil })

	// Act
	actual := args.Map{"result": o.IsEqual("a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_StringsOnce_IsEqual_LengthMismatch(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })

	// Act
	actual := args.Map{"result": o.IsEqual("a")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ===== MapStringStringOnce.IsEqual nil paths =====

func Test_MapStringStringOnce_IsEqual_BothNil_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })

	// Act
	actual := args.Map{"result": o.IsEqual(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_MapStringStringOnce_IsEqual_OneSideNil(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })

	// Act
	actual := args.Map{"result": o.IsEqual(map[string]string{"a": "1"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_MapStringStringOnce_IsEqual_LengthMismatch(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})

	// Act
	actual := args.Map{"result": o.IsEqual(map[string]string{"a": "1"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ===== BytesErrorOnce.IsEmpty =====

func Test_BytesErrorOnce_IsEmpty_True(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })

	// Act
	actual := args.Map{"result": o.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_BytesErrorOnce_IsStringEmpty(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })

	// Act
	actual := args.Map{"result": o.IsStringEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty string", actual)
}

func Test_BytesErrorOnce_IsStringEmptyOrWhitespace_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("  "), nil })

	// Act
	actual := args.Map{"result": o.IsStringEmptyOrWhitespace()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected whitespace-empty", actual)
}

func Test_BytesErrorOnce_HandleError_NoError_FromErrorOnceStringItera(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("ok"), nil })
	o.HandleError()
}

func Test_BytesErrorOnce_HandleError_Panic_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("e") })
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.HandleError()
}

// ===== AnyOnce.IsStringEmpty & IsNull =====

func Test_AnyOnce_IsStringEmpty_FromErrorOnceStringItera(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return nil })

	// Act
	actual := args.Map{"result": o.IsStringEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_AnyOnce_IsNull(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return nil })

	// Act
	actual := args.Map{"result": o.IsNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ===== Bug-fix tests: Deserialize unmarshal error paths (previously unreachable) =====

func Test_AnyOnce_Deserialize_UnmarshalError(t *testing.T) {
	// Arrange
	// Value is valid JSON ("hello") but cannot unmarshal into *int
	o := coreonce.NewAnyOncePtr(func() any { return "hello" })
	var result int
	err := o.Deserialize(&result)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error, got nil", actual)
	expected = "deserializing failed:"
	actual = args.Map{"result": len(err.Error()) < len(expected)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error too short:", actual)
	actual = args.Map{"result": err.Error()[:len(expected)] != expected}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error starting with '', got ''", actual)
}

func Test_AnyOnce_Deserialize_UnmarshalError_NilToPtr(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return map[string]any{"a": 1} })
	err := o.Deserialize(nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error for nil toPtr", actual)
}

func Test_AnyErrorOnce_Deserialize_UnmarshalError(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hello", nil })
	var result int
	err := o.Deserialize(&result)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error, got nil", actual)
	expected = "deserializing failed:"
	actual = args.Map{"result": err.Error()[:len(expected)] != expected}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error starting with '', got ''", actual)
}

func Test_AnyErrorOnce_Deserialize_UnmarshalError_NilToPtr(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]any{"a": 1}, nil })
	err := o.Deserialize(nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error for nil toPtr", actual)
}
