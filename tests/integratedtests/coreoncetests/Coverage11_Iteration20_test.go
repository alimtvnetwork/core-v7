package coreoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== ErrorOnce coverage =====

func Test_I20_ErrorOnce_String_HasError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("test-err") })
	s := o.String()
	actual := args.Map{"result": s != "test-err"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'test-err', got ''", actual)
}

func Test_I20_ErrorOnce_Message_NilError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	actual := args.Map{"result": o.Message() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty message for nil error", actual)
}

func Test_I20_ErrorOnce_Message_HasError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("msg") })
	actual := args.Map{"result": o.Message() != "msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'msg'", actual)
}

func Test_I20_ErrorOnce_IsMessageEqual(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("hello") })
	actual := args.Map{"result": o.IsMessageEqual("hello")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsMessageEqual("nope")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_I20_ErrorOnce_IsMessageEqual_Nil(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	actual := args.Map{"result": o.IsMessageEqual("anything")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil error", actual)
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
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
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
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.HandleErrorWith("context")
}

func Test_I20_ErrorOnce_ConcatNewString_NoError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	s := o.ConcatNewString("a", "b")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty concat", actual)
}

func Test_I20_ErrorOnce_ConcatNewString_HasError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("err") })
	s := o.ConcatNewString("extra")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_I20_ErrorOnce_ConcatNew(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("base") })
	err := o.ConcatNew("more")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_ErrorOnce_MarshalJSON(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("test") })
	b, err := o.MarshalJSON()
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty json", actual)
}

func Test_I20_ErrorOnce_MarshalJSON_NilError(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	b, err := o.MarshalJSON()
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": string(b) != `""`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string json", actual)
}

func Test_I20_ErrorOnce_UnmarshalJSON(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return nil })
	err := o.UnmarshalJSON([]byte(`"test-error"`))
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": o.IsMessageEqual("test-error")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected unmarshalled error message", actual)
}

func Test_I20_ErrorOnce_Predicates(t *testing.T) {
	oErr := coreonce.NewErrorOncePtr(func() error { return errors.New("e") })
	oNil := coreonce.NewErrorOncePtr(func() error { return nil })

	actual := args.Map{"result": oErr.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasError true", actual)
	actual := args.Map{"result": oNil.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasError false", actual)
	actual := args.Map{"result": oErr.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid true", actual)
	actual := args.Map{"result": oNil.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsValid true", actual)
	actual := args.Map{"result": oNil.IsSuccess()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsSuccess true", actual)
	actual := args.Map{"result": oErr.IsFailed()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsFailed true", actual)
	actual := args.Map{"result": oErr.IsDefined()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsDefined true", actual)
	actual := args.Map{"result": oErr.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem true", actual)
	actual := args.Map{"result": oNil.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem false for nil error", actual)
}

func Test_I20_ErrorOnce_Execute(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("ex") })
	actual := args.Map{"result": o.Execute() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_I20_ErrorOnce_Serialize(t *testing.T) {
	o := coreonce.NewErrorOncePtr(func() error { return errors.New("ser") })
	b, err := o.Serialize()
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

// ===== BytesErrorOnce additional coverage =====

func Test_I20_BytesErrorOnce_MarshalJSON(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`"hello"`), nil
	})
	b, err := o.MarshalJSON()
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_BytesErrorOnce_SerializeMust_Success(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`"ok"`), nil
	})
	b := o.SerializeMust()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_BytesErrorOnce_SerializeMust_Panic(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("fail")
	})
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.SerializeMust()
}

func Test_I20_BytesErrorOnce_DeserializeMust_Success(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`"value"`), nil
	})
	var s string
	o.DeserializeMust(&s)
	actual := args.Map{"result": s != "value"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'value', got ''", actual)
}

func Test_I20_BytesErrorOnce_DeserializeMust_Panic(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("err")
	})
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
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
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error", actual)
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
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.MustHaveSafeItems()
}

func Test_I20_BytesErrorOnce_MustHaveSafeItems_PanicOnEmpty(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
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
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
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

	actual := args.Map{"result": oOk.HasSafeItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasSafeItems true", actual)
	actual := args.Map{"result": oErr.HasSafeItems()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasSafeItems false", actual)
	actual := args.Map{"result": oOk.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsValid true", actual)
	actual := args.Map{"result": oErr.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid true", actual)
	actual := args.Map{"result": oOk.IsSuccess()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsSuccess true", actual)
	actual := args.Map{"result": oErr.IsFailed()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsFailed true", actual)
	actual := args.Map{"result": oOk.IsDefined()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsDefined true", actual)
	actual := args.Map{"result": oOk.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem true", actual)
}

func Test_I20_BytesErrorOnce_IsEmptyBytes(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})
	actual := args.Map{"result": o.IsEmptyBytes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEmptyBytes true", actual)
}

func Test_I20_BytesErrorOnce_ValueWithError(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("val"), nil
	})
	b, err := o.ValueWithError()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected value with no error", actual)
}

// ===== AnyOnce additional coverage =====

func Test_I20_AnyOnce_CastValueString(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "hello" })
	val, ok := o.CastValueString()
	actual := args.Map{"result": ok || val != "hello"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_I20_AnyOnce_CastValueString_Fail(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return 42 })
	_, ok := o.CastValueString()
	actual := args.Map{"result": ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected failed cast", actual)
}

func Test_I20_AnyOnce_CastValueStrings(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return []string{"a", "b"} })
	val, ok := o.CastValueStrings()
	actual := args.Map{"result": ok || len(val) != 2}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_I20_AnyOnce_CastValueHashmapMap(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return map[string]string{"k": "v"} })
	val, ok := o.CastValueHashmapMap()
	actual := args.Map{"result": ok || val["k"] != "v"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_I20_AnyOnce_CastValueMapStringAnyMap(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return map[string]any{"k": 1} })
	val, ok := o.CastValueMapStringAnyMap()
	actual := args.Map{"result": ok || val["k"] != 1}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_I20_AnyOnce_CastValueBytes(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return []byte("hi") })
	val, ok := o.CastValueBytes()
	actual := args.Map{"result": ok || string(val) != "hi"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected successful cast", actual)
}

func Test_I20_AnyOnce_Serialize_Success(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "test" })
	b, err := o.Serialize()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected successful serialize", actual)
}

func Test_I20_AnyOnce_SerializeSkipExistingError(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "test" })
	b, err := o.SerializeSkipExistingError()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_I20_AnyOnce_SerializeMust_Success(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "val" })
	b := o.SerializeMust()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_AnyOnce_SerializeMust_Panic(t *testing.T) {
	ch := make(chan int)
	o := coreonce.NewAnyOncePtr(func() any { return ch })
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.SerializeMust()
}

func Test_I20_AnyOnce_ValueStringMust(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "abc" })
	s := o.ValueStringMust()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_I20_AnyOnce_SafeString(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "abc" })
	s := o.SafeString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_I20_AnyOnce_ValueString_Nil(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return nil })
	s := o.ValueString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil bracket string", actual)
}

func Test_I20_AnyOnce_ValueString_Cached(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "cached" })
	_ = o.ValueString()
	s2 := o.ValueString() // should hit cache
	actual := args.Map{"result": s2 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached value", actual)
}

func Test_I20_AnyOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return "  " })
	// String() returns formatted, not just spaces
	_ = o.IsStringEmptyOrWhitespace()
}

func Test_I20_AnyOnce_ValueOnly(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return 42 })
	actual := args.Map{"result": o.ValueOnly() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_I20_AnyOnce_IsInitialized(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return 1 })
	actual := args.Map{"result": o.IsInitialized()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not initialized", actual)
	o.Value()
	actual := args.Map{"result": o.IsInitialized()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected initialized", actual)
}

// ===== AnyErrorOnce additional coverage =====

func Test_I20_AnyErrorOnce_ExecuteMust_Success(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "ok", nil })
	v := o.ExecuteMust()
	actual := args.Map{"result": v != "ok"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'ok'", actual)
}

func Test_I20_AnyErrorOnce_ExecuteMust_Panic(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.ExecuteMust()
}

func Test_I20_AnyErrorOnce_ValueMust_Success(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return 99, nil })
	actual := args.Map{"result": o.ValueMust() != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
}

func Test_I20_AnyErrorOnce_ValueMust_Panic(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.ValueMust()
}

func Test_I20_AnyErrorOnce_ValueStringMust_Success(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "val", nil })
	s := o.ValueStringMust()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_I20_AnyErrorOnce_ValueStringMust_Panic(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("err") })
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.ValueStringMust()
}

func Test_I20_AnyErrorOnce_ValueString_Nil(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	s, err := o.ValueString()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil bracket", actual)
}

func Test_I20_AnyErrorOnce_ValueString_Cached(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "x", nil })
	_, _ = o.ValueString()
	s2, _ := o.ValueString() // cached
	actual := args.Map{"result": s2 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached", actual)
}

func Test_I20_AnyErrorOnce_CastValueString(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hi", nil })
	val, err, ok := o.CastValueString()
	actual := args.Map{"result": ok || err != nil || val != "hi"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_I20_AnyErrorOnce_CastValueStrings(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []string{"a"}, nil })
	val, err, ok := o.CastValueStrings()
	actual := args.Map{"result": ok || err != nil || len(val) != 1}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_I20_AnyErrorOnce_CastValueHashmapMap(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]string{"k": "v"}, nil })
	val, err, ok := o.CastValueHashmapMap()
	actual := args.Map{"result": ok || err != nil || val["k"] != "v"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_I20_AnyErrorOnce_CastValueMapStringAnyMap(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]any{"k": 1}, nil })
	val, err, ok := o.CastValueMapStringAnyMap()
	actual := args.Map{"result": ok || err != nil || val["k"] != 1}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_I20_AnyErrorOnce_CastValueBytes(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []byte("b"), nil })
	val, err, ok := o.CastValueBytes()
	actual := args.Map{"result": ok || err != nil || string(val) != "b"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_I20_AnyErrorOnce_SerializeSkipExistingError(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "v", nil })
	b, err := o.SerializeSkipExistingError()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_I20_AnyErrorOnce_SerializeMust_Success(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "v", nil })
	b := o.SerializeMust()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_I20_AnyErrorOnce_SerializeMust_Panic(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.SerializeMust()
}

func Test_I20_AnyErrorOnce_Serialize_MarshalError(t *testing.T) {
	ch := make(chan int)
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return ch, nil })
	_, err := o.Serialize()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected marshal error", actual)
}

func Test_I20_AnyErrorOnce_ValueOnly_Initialized(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "v", nil })
	o.Value() // initialize
	v := o.ValueOnly()
	actual := args.Map{"result": v != "v"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'v'", actual)
}

func Test_I20_AnyErrorOnce_IsStringEmpty(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	actual := args.Map{"result": o.IsStringEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_I20_AnyErrorOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	actual := args.Map{"result": o.IsStringEmptyOrWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_I20_AnyErrorOnce_Error_AlreadyInitialized(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("e") })
	o.Value() // initialize
	err := o.Error()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_I20_AnyErrorOnce_IsEmpty_NilValue(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	actual := args.Map{"result": o.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ===== IntegerOnce additional coverage =====

func Test_I20_IntegerOnce_Comparisons(t *testing.T) {
	o := coreonce.NewIntegerOncePtr(func() int { return 5 })
	actual := args.Map{"result": o.IsAbove(3)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsAboveEqual(5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsLessThan(10)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsLessThanEqual(5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsAboveZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsAboveEqualZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsLessThanZero()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": o.IsLessThanEqualZero()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": o.IsPositive()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsNegative()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": o.IsValidIndex()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsInvalidIndex()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_I20_IntegerOnce_NegativeComparisons(t *testing.T) {
	o := coreonce.NewIntegerOncePtr(func() int { return -1 })
	actual := args.Map{"result": o.IsNegative()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsInvalidIndex()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsLessThanZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsLessThanEqualZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ===== ByteOnce additional coverage =====

func Test_I20_ByteOnce_Methods(t *testing.T) {
	o := coreonce.NewByteOncePtr(func() byte { return 5 })
	actual := args.Map{"result": o.Int() != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual := args.Map{"result": o.IsPositive()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected positive", actual)
	actual := args.Map{"result": o.IsNegative()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "byte is unsigned, should not be negative", actual)
	actual := args.Map{"result": o.IsZero()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-zero", actual)
	actual := args.Map{"result": o.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_I20_ByteOnce_Zero(t *testing.T) {
	o := coreonce.NewByteOncePtr(func() byte { return 0 })
	actual := args.Map{"result": o.IsZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected zero", actual)
	actual := args.Map{"result": o.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ===== BytesOnce additional coverage =====

func Test_I20_BytesOnce_NilInitializer(t *testing.T) {
	o := &coreonce.BytesOnce{}
	// initializerFunc is nil
	actual := args.Map{"result": o.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": o.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 length", actual)
}

// ===== StringOnce additional coverage =====

func Test_I20_StringOnce_SplitLeftRightTrim(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return " left : right " })
	l, r := o.SplitLeftRightTrim(":")
	actual := args.Map{"result": l != "left" || r != "right"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'left','right', got '',''", actual)
}

func Test_I20_StringOnce_HasPrefix(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "hello-world" })
	actual := args.Map{"result": o.HasPrefix("hello")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsStartsWith("hello")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_I20_StringOnce_HasSuffix(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "hello-world" })
	actual := args.Map{"result": o.HasSuffix("world")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsEndsWith("world")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_I20_StringOnce_IsContains(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "hello-world" })
	actual := args.Map{"result": o.IsContains("lo-wo")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_I20_StringOnce_IsEmptyOrWhitespace(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "  " })
	actual := args.Map{"result": o.IsEmptyOrWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_I20_StringOnce_Bytes(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "abc" })
	actual := args.Map{"result": string(o.Bytes()) != "abc"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'abc'", actual)
}

func Test_I20_StringOnce_Error(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "err-msg" })
	actual := args.Map{"result": o.Error().Error() != "err-msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'err-msg'", actual)
}

func Test_I20_StringOnce_SplitBy(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "a,b,c" })
	parts := o.SplitBy(",")
	actual := args.Map{"result": len(parts) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 parts", actual)
}

func Test_I20_StringOnce_ValuePtr(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "ptr" })
	p := o.ValuePtr()
	actual := args.Map{"result": *p != "ptr"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'ptr'", actual)
}

// ===== MapStringStringOnce additional coverage =====

func Test_I20_MapStringStringOnce_AllValuesSorted(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"b": "z", "a": "y"}
	})
	vs := o.AllValuesSorted()
	actual := args.Map{"result": len(vs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	// call again to hit cache
	vs2 := o.AllValuesSorted()
	actual := args.Map{"result": len(vs2) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached 2", actual)
}

func Test_I20_MapStringStringOnce_GetValueWithStatus(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	val, has := o.GetValueWithStatus("k")
	actual := args.Map{"result": has || val != "v"}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected found", actual)
	_, has2 := o.GetValueWithStatus("missing")
	actual := args.Map{"result": has2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_I20_MapStringStringOnce_ValuesPtr(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	p := o.ValuesPtr()
	actual := args.Map{"result": p == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil ptr", actual)
}

func Test_I20_MapStringStringOnce_Strings_Cached(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	_ = o.Strings()
	s2 := o.Strings() // cached
	actual := args.Map{"result": len(s2) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_I20_MapStringStringOnce_IsMissing(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	actual := args.Map{"result": o.IsMissing("nope")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.IsMissing("k")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_I20_MapStringStringOnce_String(t *testing.T) {
	o := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	s := o.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ===== StringsOnce additional coverage =====

func Test_I20_StringsOnce_HasAll_Missing(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	actual := args.Map{"result": o.HasAll("a", "c")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_I20_StringsOnce_UniqueMapLock(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	m := o.UniqueMapLock()
	actual := args.Map{"result": len(m) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_I20_StringsOnce_CsvLines(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	lines := o.CsvLines()
	actual := args.Map{"result": len(lines) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_I20_StringsOnce_RangesMap(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"x", "y"} })
	m := o.RangesMap()
	actual := args.Map{"result": len(m) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_I20_StringsOnce_SafeStrings(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a"} })
	actual := args.Map{"result": len(o.SafeStrings()) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_I20_StringsOnce_SafeStrings_Empty(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{} })
	actual := args.Map{"result": len(o.SafeStrings()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_I20_StringsOnce_String(t *testing.T) {
	o := coreonce.NewStringsOnce(func() []string { return []string{"a", "b"} })
	s := o.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_I20_StringsOnce_Length_NilValues(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return nil })
	actual := args.Map{"result": o.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_I20_StringsOnce_UniqueMap_NilValues(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return nil })
	m := o.UniqueMap()
	actual := args.Map{"result": len(m) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map", actual)
}

func Test_I20_StringsOnce_UniqueMap_Cached(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a"} })
	_ = o.UniqueMap()
	m2 := o.UniqueMap() // cached
	actual := args.Map{"result": len(m2) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ===== IntegersOnce additional coverage =====

func Test_I20_IntegersOnce_RangesBoolMap(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2, 3} })
	m := o.RangesBoolMap()
	actual := args.Map{"result": len(m) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_I20_IntegersOnce_RangesBoolMap_Empty(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := o.RangesBoolMap()
	actual := args.Map{"result": len(m) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_I20_IntegersOnce_UniqueMap(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 1, 2} })
	m := o.UniqueMap()
	actual := args.Map{"result": len(m) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 unique", actual)
}

func Test_I20_IntegersOnce_UniqueMap_Empty(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := o.UniqueMap()
	actual := args.Map{"result": len(m) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_I20_IntegersOnce_Sorted_Cached(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{3, 1, 2} })
	_ = o.Sorted()
	s2 := o.Sorted() // cached
	actual := args.Map{"result": s2[0] != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_I20_IntegersOnce_RangesMap_Empty(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := o.RangesMap()
	actual := args.Map{"result": len(m) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_I20_IntegersOnce_Aliases(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })
	actual := args.Map{"result": len(o.Values()) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": len(o.Integers()) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": len(o.Slice()) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": len(o.List()) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ===== MapStringStringOnce.HasAll =====

func Test_I20_MapStringStringOnce_HasAll(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})
	actual := args.Map{"result": o.HasAll("a", "b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": o.HasAll("a", "c")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ===== MapStringStringOnce.AllKeys cached =====

func Test_I20_MapStringStringOnce_AllKeys_Cached(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	_ = o.AllKeys()
	k2 := o.AllKeys() // cached
	actual := args.Map{"result": len(k2) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_I20_MapStringStringOnce_AllValues_Cached(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	_ = o.AllValues()
	v2 := o.AllValues() // cached
	actual := args.Map{"result": len(v2) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_I20_MapStringStringOnce_AllKeysSorted_Cached(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"b": "2", "a": "1"}
	})
	_ = o.AllKeysSorted()
	k2 := o.AllKeysSorted() // cached
	actual := args.Map{"result": k2[0] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

// ===== NewAnyOnce / NewAnyErrorOnce (non-ptr constructors) =====

func Test_I20_AnyOnce_NonPtr(t *testing.T) {
	o := coreonce.NewAnyOnce(func() any { return "np" })
	actual := args.Map{"result": o.Value() != "np"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'np'", actual)
}

func Test_I20_AnyErrorOnce_NonPtr(t *testing.T) {
	o := coreonce.NewAnyErrorOnce(func() (any, error) { return "np", nil })
	v, err := o.Value()
	actual := args.Map{"result": err != nil || v != "np"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'np'", actual)
}

func Test_I20_ErrorOnce_NonPtr(t *testing.T) {
	o := coreonce.NewErrorOnce(func() error { return nil })
	actual := args.Map{"result": o.Value() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_I20_BytesErrorOnce_NonPtr(t *testing.T) {
	o := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("x"), nil })
	v, err := o.Value()
	actual := args.Map{"result": err != nil || string(v) != "x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
}

// ===== StringsOnce.Sorted cached =====

func Test_I20_StringsOnce_Sorted_Cached(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"c", "a", "b"} })
	_ = o.Sorted()
	s2 := o.Sorted() // cached
	actual := args.Map{"result": s2[0] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

// ===== IntegersOnce.IsEqual nil paths =====

func Test_I20_IntegersOnce_IsEqual_BothNil(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return nil })
	actual := args.Map{"result": o.IsEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true - both nil/empty", actual)
}

func Test_I20_IntegersOnce_IsEqual_OneSideNil(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return nil })
	actual := args.Map{"result": o.IsEqual(1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_I20_IntegersOnce_IsEqual_LengthMismatch(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })
	actual := args.Map{"result": o.IsEqual(1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ===== StringsOnce.IsEqual nil paths =====

func Test_I20_StringsOnce_IsEqual_BothNil(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return nil })
	actual := args.Map{"result": o.IsEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_I20_StringsOnce_IsEqual_OneSideNil(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return nil })
	actual := args.Map{"result": o.IsEqual("a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_I20_StringsOnce_IsEqual_LengthMismatch(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	actual := args.Map{"result": o.IsEqual("a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ===== MapStringStringOnce.IsEqual nil paths =====

func Test_I20_MapStringStringOnce_IsEqual_BothNil(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })
	actual := args.Map{"result": o.IsEqual(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_I20_MapStringStringOnce_IsEqual_OneSideNil(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })
	actual := args.Map{"result": o.IsEqual(map[string]string{"a": "1"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_I20_MapStringStringOnce_IsEqual_LengthMismatch(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})
	actual := args.Map{"result": o.IsEqual(map[string]string{"a": "1"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ===== BytesErrorOnce.IsEmpty =====

func Test_I20_BytesErrorOnce_IsEmpty_True(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })
	actual := args.Map{"result": o.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_I20_BytesErrorOnce_IsStringEmpty(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })
	actual := args.Map{"result": o.IsStringEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty string", actual)
}

func Test_I20_BytesErrorOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("  "), nil })
	actual := args.Map{"result": o.IsStringEmptyOrWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected whitespace-empty", actual)
}

func Test_I20_BytesErrorOnce_HandleError_NoError(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("ok"), nil })
	o.HandleError()
}

func Test_I20_BytesErrorOnce_HandleError_Panic(t *testing.T) {
	o := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("e") })
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	o.HandleError()
}

// ===== AnyOnce.IsStringEmpty & IsNull =====

func Test_I20_AnyOnce_IsStringEmpty(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return nil })
	actual := args.Map{"result": o.IsStringEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_I20_AnyOnce_IsNull(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return nil })
	actual := args.Map{"result": o.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ===== Bug-fix tests: Deserialize unmarshal error paths (previously unreachable) =====

func Test_I20_AnyOnce_Deserialize_UnmarshalError(t *testing.T) {
	// Value is valid JSON ("hello") but cannot unmarshal into *int
	o := coreonce.NewAnyOncePtr(func() any { return "hello" })
	var result int
	err := o.Deserialize(&result)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error, got nil", actual)
	expected := "deserializing failed:"
	actual := args.Map{"result": len(err.Error()) < len(expected)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error too short:", actual)
	actual := args.Map{"result": err.Error()[:len(expected)] != expected}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error starting with '', got ''", actual)
}

func Test_I20_AnyOnce_Deserialize_UnmarshalError_NilToPtr(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return map[string]any{"a": 1} })
	err := o.Deserialize(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error for nil toPtr", actual)
}

func Test_I20_AnyErrorOnce_Deserialize_UnmarshalError(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hello", nil })
	var result int
	err := o.Deserialize(&result)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error, got nil", actual)
	expected := "deserializing failed:"
	actual := args.Map{"result": err.Error()[:len(expected)] != expected}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error starting with '', got ''", actual)
}

func Test_I20_AnyErrorOnce_Deserialize_UnmarshalError_NilToPtr(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) { return map[string]any{"a": 1}, nil })
	err := o.Deserialize(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error for nil toPtr", actual)
}
