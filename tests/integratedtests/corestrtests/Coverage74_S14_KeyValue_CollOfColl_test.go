package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ========================================
// S14: KeyValuePair, KeyAnyValuePair,
//       KeyValueCollection, CollectionsOfCollection,
//       newKeyValuesCreator, newCollectionsOfCollectionCreator
// ========================================

// --- KeyValuePair ---

func Test_C74_KeyValuePair_KeyName(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_KeyName", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "host", Value: "localhost"}

		// Act
		result := kv.KeyName()

		// Assert
		if result != "host" {
			t.Errorf("KeyName expected 'host', got '%s'", result)
		}
	})
}

func Test_C74_KeyValuePair_VariableName(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_VariableName", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "port", Value: "8080"}

		// Act
		result := kv.VariableName()

		// Assert
		if result != "port" {
			t.Errorf("expected 'port', got '%s'", result)
		}
	})
}

func Test_C74_KeyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueString", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "val123"}

		// Act
		result := kv.ValueString()

		// Assert
		if result != "val123" {
			t.Errorf("expected 'val123', got '%s'", result)
		}
	})
}

func Test_C74_KeyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsVariableNameEqual", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "name", Value: "x"}

		// Act & Assert
		if !kv.IsVariableNameEqual("name") {
			t.Error("expected true for matching name")
		}
		if kv.IsVariableNameEqual("other") {
			t.Error("expected false for non-matching name")
		}
	})
}

func Test_C74_KeyValuePair_IsValueEqual(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsValueEqual", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		if !kv.IsValueEqual("abc") {
			t.Error("expected true")
		}
		if kv.IsValueEqual("xyz") {
			t.Error("expected false")
		}
	})
}

func Test_C74_KeyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_Compile", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act
		result := kv.Compile()

		// Assert
		if result != kv.String() {
			t.Errorf("Compile should equal String(), got '%s'", result)
		}
	})
}

func Test_C74_KeyValuePair_IsKeyEmpty_IsValueEmpty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsKeyEmpty_IsValueEmpty", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "", Value: "x"}

		// Act & Assert
		if !kv.IsKeyEmpty() {
			t.Error("expected key empty")
		}
		if kv.IsValueEmpty() {
			t.Error("expected value not empty")
		}
	})
}

func Test_C74_KeyValuePair_HasKey_HasValue(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_HasKey_HasValue", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: ""}

		// Act & Assert
		if !kv.HasKey() {
			t.Error("expected HasKey true")
		}
		if kv.HasValue() {
			t.Error("expected HasValue false")
		}
	})
}

func Test_C74_KeyValuePair_IsKeyValueEmpty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsKeyValueEmpty", func() {
		// Arrange
		empty := corestr.KeyValuePair{}
		nonEmpty := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act & Assert
		if !empty.IsKeyValueEmpty() {
			t.Error("expected true for empty")
		}
		if nonEmpty.IsKeyValueEmpty() {
			t.Error("expected false for non-empty")
		}
	})
}

func Test_C74_KeyValuePair_TrimKey_TrimValue(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_TrimKey_TrimValue", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: " key ", Value: " val "}

		// Act & Assert
		if kv.TrimKey() != "key" {
			t.Errorf("expected 'key', got '%s'", kv.TrimKey())
		}
		if kv.TrimValue() != "val" {
			t.Errorf("expected 'val', got '%s'", kv.TrimValue())
		}
	})
}

func Test_C74_KeyValuePair_ValueBool(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueBool", func() {
		// Arrange
		trueKv := corestr.KeyValuePair{Key: "k", Value: "true"}
		falseKv := corestr.KeyValuePair{Key: "k", Value: "false"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "xyz"}
		emptyKv := corestr.KeyValuePair{Key: "k", Value: ""}

		// Act & Assert
		if !trueKv.ValueBool() {
			t.Error("expected true")
		}
		if falseKv.ValueBool() {
			t.Error("expected false")
		}
		if invalidKv.ValueBool() {
			t.Error("expected false for invalid")
		}
		if emptyKv.ValueBool() {
			t.Error("expected false for empty")
		}
	})
}

func Test_C74_KeyValuePair_ValueInt(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueInt", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		if kv.ValueInt(0) != 42 {
			t.Errorf("expected 42, got %d", kv.ValueInt(0))
		}
		if invalidKv.ValueInt(99) != 99 {
			t.Errorf("expected default 99, got %d", invalidKv.ValueInt(99))
		}
	})
}

func Test_C74_KeyValuePair_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueDefInt", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "10"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "x"}

		// Act & Assert
		if kv.ValueDefInt() != 10 {
			t.Errorf("expected 10")
		}
		if invalidKv.ValueDefInt() != 0 {
			t.Errorf("expected 0 for invalid")
		}
	})
}

func Test_C74_KeyValuePair_ValueByte(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueByte", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "200"}
		overflowKv := corestr.KeyValuePair{Key: "k", Value: "999"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		if kv.ValueByte(0) != 200 {
			t.Errorf("expected 200")
		}
		if overflowKv.ValueByte(5) != 5 {
			t.Error("expected default for overflow")
		}
		if invalidKv.ValueByte(7) != 7 {
			t.Error("expected default for invalid")
		}
	})
}

func Test_C74_KeyValuePair_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueDefByte", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "100"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "abc"}
		overflowKv := corestr.KeyValuePair{Key: "k", Value: "300"}

		// Act & Assert
		if kv.ValueDefByte() != 100 {
			t.Errorf("expected 100")
		}
		if invalidKv.ValueDefByte() != 0 {
			t.Error("expected 0")
		}
		if overflowKv.ValueDefByte() != 0 {
			t.Error("expected 0 for overflow")
		}
	})
}

func Test_C74_KeyValuePair_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueFloat64", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		if kv.ValueFloat64(0) != 3.14 {
			t.Errorf("expected 3.14")
		}
		if invalidKv.ValueFloat64(1.5) != 1.5 {
			t.Error("expected default 1.5")
		}
	})
}

func Test_C74_KeyValuePair_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueDefFloat64", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "2.5"}

		// Act & Assert
		if kv.ValueDefFloat64() != 2.5 {
			t.Errorf("expected 2.5")
		}
	})
}

func Test_C74_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueValid", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "myval"}

		// Act
		vv := kv.ValueValid()

		// Assert
		if vv.Value != "myval" || !vv.IsValid {
			t.Error("expected valid value 'myval'")
		}
	})
}

func Test_C74_KeyValuePair_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueValidOptions", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		vv := kv.ValueValidOptions(false, "err msg")

		// Assert
		if vv.IsValid || vv.Message != "err msg" {
			t.Error("expected invalid with message")
		}
	})
}

func Test_C74_KeyValuePair_Is(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_Is", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act & Assert
		if !kv.Is("a", "b") {
			t.Error("expected true")
		}
		if kv.Is("a", "c") {
			t.Error("expected false")
		}
	})
}

func Test_C74_KeyValuePair_IsKey_IsVal(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsKey_IsVal", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "x", Value: "y"}

		// Act & Assert
		if !kv.IsKey("x") {
			t.Error("expected IsKey true")
		}
		if !kv.IsVal("y") {
			t.Error("expected IsVal true")
		}
		if kv.IsKey("z") {
			t.Error("expected IsKey false")
		}
	})
}

func Test_C74_KeyValuePair_IsKeyValueAnyEmpty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsKeyValueAnyEmpty", func() {
		// Arrange
		full := corestr.KeyValuePair{Key: "k", Value: "v"}
		emptyKey := corestr.KeyValuePair{Key: "", Value: "v"}
		emptyVal := corestr.KeyValuePair{Key: "k", Value: ""}
		var nilPtr *corestr.KeyValuePair

		// Act & Assert
		if full.IsKeyValueAnyEmpty() {
			t.Error("expected false for full")
		}
		if !emptyKey.IsKeyValueAnyEmpty() {
			t.Error("expected true for empty key")
		}
		if !emptyVal.IsKeyValueAnyEmpty() {
			t.Error("expected true for empty val")
		}
		if !nilPtr.IsKeyValueAnyEmpty() {
			t.Error("expected true for nil")
		}
	})
}

func Test_C74_KeyValuePair_FormatString(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_FormatString", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "host", Value: "localhost"}

		// Act
		result := kv.FormatString("%s=%s")

		// Assert
		if result != "host=localhost" {
			t.Errorf("expected 'host=localhost', got '%s'", result)
		}
	})
}

func Test_C74_KeyValuePair_String(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_String", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act
		result := kv.String()

		// Assert
		if result != "{a:b}" {
			t.Errorf("expected '{a:b}', got '%s'", result)
		}
	})
}

func Test_C74_KeyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_Clear_Dispose", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		kv.Clear()

		// Assert
		if kv.Key != "" || kv.Value != "" {
			t.Error("expected cleared")
		}
	})
}

func Test_C74_KeyValuePair_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_Dispose_Nil", func() {
		// Arrange
		var kv *corestr.KeyValuePair

		// Act — should not panic
		kv.Clear()
		kv.Dispose()
	})
}

func Test_C74_KeyValuePair_Json_Serialize(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_Json_Serialize", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		jsonResult := kv.Json()
		bytes, err := kv.Serialize()

		// Assert
		if jsonResult.HasError() {
			t.Error("json error")
		}
		if err != nil {
			t.Errorf("serialize error: %v", err)
		}
		if len(bytes) == 0 {
			t.Error("empty bytes")
		}
	})
}

func Test_C74_KeyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_SerializeMust", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act
		bytes := kv.SerializeMust()

		// Assert
		if len(bytes) == 0 {
			t.Error("expected non-empty bytes")
		}
	})
}

// --- KeyAnyValuePair ---

func Test_C74_KeyAnyValuePair_KeyName_VariableName(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_KeyName_VariableName", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "mykey", Value: 42}

		// Act & Assert
		if kav.KeyName() != "mykey" {
			t.Error("KeyName mismatch")
		}
		if kav.VariableName() != "mykey" {
			t.Error("VariableName mismatch")
		}
	})
}

func Test_C74_KeyAnyValuePair_ValueAny(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_ValueAny", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "hello"}

		// Act
		result := kav.ValueAny()

		// Assert
		if result != "hello" {
			t.Error("ValueAny mismatch")
		}
	})
}

func Test_C74_KeyAnyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsVariableNameEqual", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "name", Value: nil}

		// Act & Assert
		if !kav.IsVariableNameEqual("name") {
			t.Error("expected true")
		}
		if kav.IsVariableNameEqual("other") {
			t.Error("expected false")
		}
	})
}

func Test_C74_KeyAnyValuePair_IsValueNull_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsValueNull_Nil", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act & Assert
		if !kav.IsValueNull() {
			t.Error("expected null for nil value")
		}
	})
}

func Test_C74_KeyAnyValuePair_IsValueNull_NilReceiver(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsValueNull_NilReceiver", func() {
		// Arrange
		var kav *corestr.KeyAnyValuePair

		// Act & Assert
		if !kav.IsValueNull() {
			t.Error("expected null for nil receiver")
		}
	})
}

func Test_C74_KeyAnyValuePair_HasNonNull(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_HasNonNull", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		kavNil := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act & Assert
		if !kav.HasNonNull() {
			t.Error("expected true for non-nil value")
		}
		if kavNil.HasNonNull() {
			t.Error("expected false for nil value")
		}
	})
}

func Test_C74_KeyAnyValuePair_HasValue(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_HasValue", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "x"}

		// Act & Assert
		if !kav.HasValue() {
			t.Error("expected true")
		}
	})
}

func Test_C74_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsValueEmptyString", func() {
		// Arrange
		kavEmpty := corestr.KeyAnyValuePair{Key: "k", Value: ""}
		kavNonEmpty := corestr.KeyAnyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		if !kavEmpty.IsValueEmptyString() {
			t.Error("expected true for empty string value")
		}
		if kavNonEmpty.IsValueEmptyString() {
			t.Error("expected false for non-empty")
		}
	})
}

func Test_C74_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsValueWhitespace", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "  "}

		// Act & Assert
		if !kav.IsValueWhitespace() {
			t.Error("expected true for whitespace value string")
		}
	})
}

func Test_C74_KeyAnyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_ValueString", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 123}

		// Act
		result := kav.ValueString()

		// Assert
		if result != "123" {
			t.Errorf("expected '123', got '%s'", result)
		}
	})
}

func Test_C74_KeyAnyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_Compile", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "a", Value: "b"}

		// Act
		result := kav.Compile()

		// Assert
		if result != kav.String() {
			t.Error("Compile should equal String()")
		}
	})
}

func Test_C74_KeyAnyValuePair_String(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_String", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "x", Value: "y"}

		// Act
		result := kav.String()

		// Assert
		if result != "{x:y}" {
			t.Errorf("expected '{x:y}', got '%s'", result)
		}
	})
}

func Test_C74_KeyAnyValuePair_Json_Serialize(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_Json_Serialize", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		jsonResult := kav.Json()
		bytes, err := kav.Serialize()

		// Assert
		if jsonResult.HasError() {
			t.Error("json error")
		}
		if err != nil {
			t.Errorf("serialize error: %v", err)
		}
		if len(bytes) == 0 {
			t.Error("empty bytes")
		}
	})
}

func Test_C74_KeyAnyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_SerializeMust", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		bytes := kav.SerializeMust()

		// Assert
		if len(bytes) == 0 {
			t.Error("expected non-empty bytes")
		}
	})
}

func Test_C74_KeyAnyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_Clear_Dispose", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		kav.Clear()

		// Assert
		if kav.Key != "" || kav.Value != nil {
			t.Error("expected cleared")
		}
	})
}

func Test_C74_KeyAnyValuePair_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_Dispose_Nil", func() {
		// Arrange
		var kav *corestr.KeyAnyValuePair

		// Act — should not panic
		kav.Clear()
		kav.Dispose()
	})
}

func Test_C74_KeyAnyValuePair_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_AsJsonContractsBinder", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		binder := kav.AsJsonContractsBinder()

		// Assert
		if binder == nil {
			t.Error("expected non-nil binder")
		}
	})
}

func Test_C74_KeyAnyValuePair_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_AsJsoner", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		jsoner := kav.AsJsoner()

		// Assert
		if jsoner == nil {
			t.Error("expected non-nil jsoner")
		}
	})
}

func Test_C74_KeyAnyValuePair_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_AsJsonParseSelfInjector", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		injector := kav.AsJsonParseSelfInjector()

		// Assert
		if injector == nil {
			t.Error("expected non-nil injector")
		}
	})
}

func Test_C74_KeyAnyValuePair_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_ParseInjectUsingJson", func() {
		// Arrange
		original := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jsonResult := original.JsonPtr()
		target := &corestr.KeyAnyValuePair{}

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result.Key != "k" {
			t.Errorf("expected key 'k', got '%s'", result.Key)
		}
	})
}

func Test_C74_KeyAnyValuePair_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_ParseInjectUsingJsonMust", func() {
		// Arrange
		original := corestr.KeyAnyValuePair{Key: "test", Value: "data"}
		jsonResult := original.JsonPtr()
		target := &corestr.KeyAnyValuePair{}

		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Assert
		if result.Key != "test" {
			t.Errorf("expected 'test', got '%s'", result.Key)
		}
	})
}

func Test_C74_KeyAnyValuePair_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_JsonParseSelfInject", func() {
		// Arrange
		original := corestr.KeyAnyValuePair{Key: "a", Value: "b"}
		jsonResult := original.JsonPtr()
		target := &corestr.KeyAnyValuePair{}

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

// --- KeyValueCollection ---

func Test_C74_KeyValueCollection_Add_Length(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Add_Length", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.Add("k1", "v1").Add("k2", "v2")

		// Assert
		if kvc.Length() != 2 {
			t.Errorf("expected 2, got %d", kvc.Length())
		}
	})
}

func Test_C74_KeyValueCollection_AddIf(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddIf", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddIf(true, "k1", "v1")
		kvc.AddIf(false, "k2", "v2")

		// Assert
		if kvc.Length() != 1 {
			t.Errorf("expected 1, got %d", kvc.Length())
		}
	})
}

func Test_C74_KeyValueCollection_IsEmpty_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_IsEmpty_HasAnyItem", func() {
		// Arrange
		empty := corestr.New.KeyValues.Empty()
		nonEmpty := corestr.New.KeyValues.Empty()
		nonEmpty.Add("k", "v")

		// Act & Assert
		if !empty.IsEmpty() {
			t.Error("expected empty")
		}
		if empty.HasAnyItem() {
			t.Error("expected no items")
		}
		if nonEmpty.IsEmpty() {
			t.Error("expected not empty")
		}
		if !nonEmpty.HasAnyItem() {
			t.Error("expected has items")
		}
	})
}

func Test_C74_KeyValueCollection_Count(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Count", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1")

		// Act & Assert
		if kvc.Count() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_C74_KeyValueCollection_First_Last(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_First_Last", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2").Add("c", "3")

		// Act & Assert
		if kvc.First().Key != "a" {
			t.Error("First key should be 'a'")
		}
		if kvc.Last().Key != "c" {
			t.Error("Last key should be 'c'")
		}
	})
}

func Test_C74_KeyValueCollection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_FirstOrDefault_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		result := kvc.FirstOrDefault()

		// Assert
		if result != nil {
			t.Error("expected nil for empty collection")
		}
	})
}

func Test_C74_KeyValueCollection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_LastOrDefault_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		result := kvc.LastOrDefault()

		// Assert
		if result != nil {
			t.Error("expected nil for empty collection")
		}
	})
}

func Test_C74_KeyValueCollection_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_LastIndex_HasIndex", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act & Assert
		if kvc.LastIndex() != 1 {
			t.Errorf("expected 1")
		}
		if !kvc.HasIndex(0) {
			t.Error("expected true for index 0")
		}
		if !kvc.HasIndex(1) {
			t.Error("expected true for index 1")
		}
		if kvc.HasIndex(2) {
			t.Error("expected false for index 2")
		}
	})
}

func Test_C74_KeyValueCollection_HasKey_IsContains(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_HasKey_IsContains", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("host", "localhost")

		// Act & Assert
		if !kvc.HasKey("host") {
			t.Error("expected HasKey true")
		}
		if kvc.HasKey("port") {
			t.Error("expected HasKey false")
		}
		if !kvc.IsContains("host") {
			t.Error("expected IsContains true")
		}
	})
}

func Test_C74_KeyValueCollection_Get(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Get", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("port", "8080")

		// Act
		val, found := kvc.Get("port")
		_, notFound := kvc.Get("missing")

		// Assert
		if !found || val != "8080" {
			t.Error("expected found with val '8080'")
		}
		if notFound {
			t.Error("expected not found")
		}
	})
}

func Test_C74_KeyValueCollection_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_SafeValueAt", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act & Assert
		if kvc.SafeValueAt(0) != "1" {
			t.Error("expected '1'")
		}
		if kvc.SafeValueAt(5) != "" {
			t.Error("expected empty for out of range")
		}
	})
}

func Test_C74_KeyValueCollection_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_SafeValuesAtIndexes", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2").Add("c", "3")

		// Act
		vals := kvc.SafeValuesAtIndexes(0, 2)

		// Assert
		if len(vals) != 2 || vals[0] != "1" || vals[1] != "3" {
			t.Errorf("expected ['1','3'], got %v", vals)
		}
	})
}

func Test_C74_KeyValueCollection_Strings(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Strings", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1")

		// Act
		result := kvc.Strings()

		// Assert
		if len(result) != 1 {
			t.Errorf("expected 1 string")
		}
	})
}

func Test_C74_KeyValueCollection_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Strings_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		result := kvc.Strings()

		// Assert
		if len(result) != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C74_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_StringsUsingFormat", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("host", "localhost")

		// Act
		result := kvc.StringsUsingFormat("%s=%s")

		// Assert
		if len(result) != 1 || result[0] != "host=localhost" {
			t.Errorf("expected 'host=localhost', got %v", result)
		}
	})
}

func Test_C74_KeyValueCollection_StringsUsingFormat_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_StringsUsingFormat_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		result := kvc.StringsUsingFormat("%s=%s")

		// Assert
		if len(result) != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C74_KeyValueCollection_AllKeys_AllValues(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AllKeys_AllValues", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act
		keys := kvc.AllKeys()
		values := kvc.AllValues()

		// Assert
		if len(keys) != 2 || keys[0] != "a" || keys[1] != "b" {
			t.Errorf("keys mismatch: %v", keys)
		}
		if len(values) != 2 || values[0] != "1" || values[1] != "2" {
			t.Errorf("values mismatch: %v", values)
		}
	})
}

func Test_C74_KeyValueCollection_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AllKeysSorted", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("c", "3").Add("a", "1").Add("b", "2")

		// Act
		keys := kvc.AllKeysSorted()

		// Assert
		if keys[0] != "a" || keys[1] != "b" || keys[2] != "c" {
			t.Errorf("expected sorted keys, got %v", keys)
		}
	})
}

func Test_C74_KeyValueCollection_Join_JoinKeys_JoinValues(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Join_JoinKeys_JoinValues", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act & Assert
		joinKeys := kvc.JoinKeys(",")
		joinValues := kvc.JoinValues(",")

		if joinKeys != "a,b" {
			t.Errorf("expected 'a,b', got '%s'", joinKeys)
		}
		if joinValues != "1,2" {
			t.Errorf("expected '1,2', got '%s'", joinValues)
		}
	})
}

func Test_C74_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Find", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2").Add("c", "3")

		// Act
		found := kvc.Find(func(index int, current corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return current, current.Key == "b", false
		})

		// Assert
		if len(found) != 1 || found[0].Key != "b" {
			t.Errorf("expected to find 'b', got %v", found)
		}
	})
}

func Test_C74_KeyValueCollection_Find_WithBreak(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Find_WithBreak", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2").Add("c", "3")

		// Act
		found := kvc.Find(func(index int, current corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return current, true, index == 0
		})

		// Assert
		if len(found) != 1 {
			t.Errorf("expected 1 due to break, got %d", len(found))
		}
	})
}

func Test_C74_KeyValueCollection_Find_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Find_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		found := kvc.Find(func(index int, current corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return current, true, false
		})

		// Assert
		if len(found) != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C74_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddStringBySplit", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddStringBySplit("=", "host=localhost")

		// Assert
		if kvc.Length() != 1 {
			t.Error("expected 1 item")
		}
	})
}

func Test_C74_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddStringBySplitTrim", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddStringBySplitTrim("=", " host = localhost ")

		// Assert
		if kvc.Length() != 1 {
			t.Error("expected 1 item")
		}
	})
}

func Test_C74_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Adds", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Assert
		if kvc.Length() != 2 {
			t.Errorf("expected 2, got %d", kvc.Length())
		}
	})
}

func Test_C74_KeyValueCollection_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Adds_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.Adds()

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_KeyValueCollection_AddMap(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddMap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddMap(map[string]string{"a": "1", "b": "2"})

		// Assert
		if kvc.Length() != 2 {
			t.Errorf("expected 2, got %d", kvc.Length())
		}
	})
}

func Test_C74_KeyValueCollection_AddMap_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddMap_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddMap(nil)

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_KeyValueCollection_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddHashsetMap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddHashsetMap(map[string]bool{"x": true, "y": true})

		// Assert
		if kvc.Length() != 2 {
			t.Errorf("expected 2, got %d", kvc.Length())
		}
	})
}

func Test_C74_KeyValueCollection_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddHashsetMap_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddHashsetMap(nil)

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_KeyValueCollection_AddHashset(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddHashset", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		kvc.AddHashset(hs)

		// Assert
		if kvc.Length() != 2 {
			t.Errorf("expected 2, got %d", kvc.Length())
		}
	})
}

func Test_C74_KeyValueCollection_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddHashset_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddHashset(nil)

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_KeyValueCollection_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddsHashmap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		kvc.AddsHashmap(hm)

		// Assert
		if kvc.Length() != 1 {
			t.Errorf("expected 1, got %d", kvc.Length())
		}
	})
}

func Test_C74_KeyValueCollection_AddsHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddsHashmap_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddsHashmap(nil)

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_KeyValueCollection_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddsHashmaps", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		hm1 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2"})

		// Act
		kvc.AddsHashmaps(hm1, hm2)

		// Assert
		if kvc.Length() != 2 {
			t.Errorf("expected 2, got %d", kvc.Length())
		}
	})
}

func Test_C74_KeyValueCollection_AddsHashmaps_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddsHashmaps_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddsHashmaps()

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_KeyValueCollection_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Hashmap_Map", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act
		hm := kvc.Hashmap()
		m := kvc.Map()

		// Assert
		if hm.Length() != 2 {
			t.Errorf("expected hashmap length 2")
		}
		if len(m) != 2 {
			t.Errorf("expected map length 2")
		}
	})
}

func Test_C74_KeyValueCollection_Json_Serialize(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Json_Serialize", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		jsonResult := kvc.Json()
		bytes, err := kvc.Serialize()

		// Assert
		if jsonResult.HasError() {
			t.Error("json error")
		}
		if err != nil {
			t.Errorf("serialize error: %v", err)
		}
		if len(bytes) == 0 {
			t.Error("empty bytes")
		}
	})
}

func Test_C74_KeyValueCollection_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_MarshalUnmarshalJSON", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act
		bytes, err := kvc.MarshalJSON()
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		target := corestr.New.KeyValues.Empty()
		err = target.UnmarshalJSON(bytes)

		// Assert
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		if target.Length() != 2 {
			t.Errorf("expected 2, got %d", target.Length())
		}
	})
}

func Test_C74_KeyValueCollection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_ParseInjectUsingJson", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("x", "y")
		jsonResult := kvc.JsonPtr()
		target := corestr.New.KeyValues.Empty()

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result.Length() < 1 {
			t.Error("expected at least 1 item")
		}
	})
}

func Test_C74_KeyValueCollection_AsJsoner_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AsJsoner_AsJsonContractsBinder", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act & Assert
		if kvc.AsJsoner() == nil {
			t.Error("expected non-nil jsoner")
		}
		if kvc.AsJsonContractsBinder() == nil {
			t.Error("expected non-nil binder")
		}
	})
}

func Test_C74_KeyValueCollection_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AsJsonParseSelfInjector", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		injector := kvc.AsJsonParseSelfInjector()

		// Assert
		if injector == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_KeyValueCollection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_JsonParseSelfInject", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		jsonResult := kvc.JsonPtr()
		target := corestr.New.KeyValues.Empty()

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func Test_C74_KeyValueCollection_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Clear_Dispose", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1")

		// Act
		kvc.Clear()

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0 after clear")
		}
	})
}

func Test_C74_KeyValueCollection_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Dispose_Nil", func() {
		// Arrange
		var kvc *corestr.KeyValueCollection

		// Act — should not panic
		kvc.Clear()
		kvc.Dispose()
	})
}

func Test_C74_KeyValueCollection_Deserialize(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Deserialize", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		var target []corestr.KeyValuePair
		err := kvc.Deserialize(&target)

		// Assert
		if err != nil {
			t.Errorf("deserialize error: %v", err)
		}
	})
}

func Test_C74_KeyValueCollection_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_SerializeMust", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		bytes := kvc.SerializeMust()

		// Assert
		if len(bytes) == 0 {
			t.Error("expected non-empty bytes")
		}
	})
}

// --- newKeyValuesCreator ---

func Test_C74_NewKeyValues_Cap(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_Cap", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.Cap(5)

		// Assert
		if kvc == nil || kvc.Length() != 0 {
			t.Error("expected empty with capacity")
		}
	})
}

func Test_C74_NewKeyValues_UsingMap(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingMap", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Assert
		if kvc.Length() != 2 {
			t.Errorf("expected 2, got %d", kvc.Length())
		}
	})
}

func Test_C74_NewKeyValues_UsingMap_Empty(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingMap_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{})

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_NewKeyValues_UsingKeyValuePairs(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingKeyValuePairs", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValuePairs(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Assert
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_C74_NewKeyValues_UsingKeyValuePairs_Empty(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingKeyValuePairs_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValuePairs()

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_NewKeyValues_UsingKeyValueStrings(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingKeyValueStrings", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValueStrings(
			[]string{"a", "b"},
			[]string{"1", "2"},
		)

		// Assert
		if kvc.Length() != 2 {
			t.Errorf("expected 2, got %d", kvc.Length())
		}
	})
}

func Test_C74_NewKeyValues_UsingKeyValueStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingKeyValueStrings_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValueStrings([]string{}, []string{})

		// Assert
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

// --- CollectionsOfCollection ---

func Test_C74_CollOfColl_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_IsEmpty_HasItems", func() {
		// Arrange
		empty := corestr.New.CollectionsOfCollection.Empty()
		nonEmpty := corestr.New.CollectionsOfCollection.SpreadStrings(false, "a", "b")

		// Act & Assert
		if !empty.IsEmpty() {
			t.Error("expected empty")
		}
		if empty.HasItems() {
			t.Error("expected no items")
		}
		if nonEmpty.IsEmpty() {
			t.Error("expected not empty")
		}
	})
}

func Test_C74_CollOfColl_Length(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_Length", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		if coc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_CollOfColl_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AllIndividualItemsLength", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c"})
		coc.Add(c1).Add(c2)

		// Act
		total := coc.AllIndividualItemsLength()

		// Assert
		if total != 3 {
			t.Errorf("expected 3, got %d", total)
		}
	})
}

func Test_C74_CollOfColl_AllIndividualItemsLength_Empty(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AllIndividualItemsLength_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		if coc.AllIndividualItemsLength() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_CollOfColl_Items(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_Items", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := corestr.New.Collection.Strings([]string{"x"})
		coc.Add(c)

		// Act
		items := coc.Items()

		// Assert
		if len(items) != 1 {
			t.Errorf("expected 1, got %d", len(items))
		}
	})
}

func Test_C74_CollOfColl_List(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_List", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c"})
		coc.Add(c1).Add(c2)

		// Act
		list := coc.List(0)

		// Assert
		if len(list) != 3 {
			t.Errorf("expected 3, got %d", len(list))
		}
	})
}

func Test_C74_CollOfColl_List_Empty(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_List_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		list := coc.List(5)

		// Assert
		if len(list) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_CollOfColl_ToCollection(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_ToCollection", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := corestr.New.Collection.Strings([]string{"x", "y"})
		coc.Add(c)

		// Act
		col := coc.ToCollection()

		// Assert
		if col.Length() != 2 {
			t.Errorf("expected 2, got %d", col.Length())
		}
	})
}

func Test_C74_CollOfColl_AddStrings(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AddStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddStrings(false, []string{"a", "b"})

		// Assert
		if coc.Length() != 1 {
			t.Errorf("expected 1, got %d", coc.Length())
		}
	})
}

func Test_C74_CollOfColl_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AddStrings_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddStrings(false, []string{})

		// Assert
		if coc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_CollOfColl_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AddsStringsOfStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b", "c"})

		// Assert
		if coc.Length() != 2 {
			t.Errorf("expected 2, got %d", coc.Length())
		}
	})
}

func Test_C74_CollOfColl_AddsStringsOfStrings_Nil(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AddsStringsOfStrings_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddsStringsOfStrings(false)

		// Assert
		if coc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_CollOfColl_Adds_AddCollections(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_Adds_AddCollections", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c1 := *corestr.New.Collection.Strings([]string{"a"})

		// Act
		coc.Adds(c1)

		// Assert
		if coc.Length() != 1 {
			t.Errorf("expected 1, got %d", coc.Length())
		}
	})
}

func Test_C74_CollOfColl_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_Adds_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddCollections()

		// Assert
		if coc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C74_CollOfColl_Add_EmptyCollection(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_Add_EmptyCollection", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		emptyCol := corestr.New.Collection.Strings([]string{})

		// Act
		coc.Add(emptyCol)

		// Assert — empty collection should be skipped
		if coc.Length() != 0 {
			t.Errorf("expected 0, got %d", coc.Length())
		}
	})
}

func Test_C74_CollOfColl_String(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_String", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(c)

		// Act
		result := coc.String()

		// Assert
		if result == "" {
			t.Error("expected non-empty string")
		}
	})
}

func Test_C74_CollOfColl_Json_Serialize(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_Json_Serialize", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)

		// Act
		jsonResult := coc.Json()

		// Assert
		if jsonResult.HasError() {
			t.Errorf("json error: %v", jsonResult.Error)
		}
	})
}

func Test_C74_CollOfColl_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_MarshalUnmarshalJSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(c)

		// Act
		bytes, err := coc.MarshalJSON()
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		target := corestr.New.CollectionsOfCollection.Empty()
		err = target.UnmarshalJSON(bytes)

		// Assert
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
	})
}

func Test_C74_CollOfColl_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_ParseInjectUsingJson", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := corestr.New.Collection.Strings([]string{"x"})
		coc.Add(c)
		jsonResult := coc.JsonPtr()
		target := corestr.New.CollectionsOfCollection.Empty()

		// Act
		_, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func Test_C74_CollOfColl_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AsJsoner", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		if coc.AsJsoner() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_CollOfColl_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AsJsonParseSelfInjector", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		if coc.AsJsonParseSelfInjector() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_CollOfColl_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AsJsonMarshaller", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		if coc.AsJsonMarshaller() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_CollOfColl_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AsJsonContractsBinder", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		if coc.AsJsonContractsBinder() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_CollOfColl_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_JsonParseSelfInject", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		jsonResult := coc.JsonPtr()
		target := corestr.New.CollectionsOfCollection.Empty()

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

// --- newCollectionsOfCollectionCreator ---

func Test_C74_NewCollOfColl_Cap(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_Cap", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Cap(5)

		// Assert
		if coc == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_NewCollOfColl_StringsOfStrings(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_StringsOfStrings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})

		// Assert
		if coc.Length() != 2 {
			t.Errorf("expected 2, got %d", coc.Length())
		}
	})
}

func Test_C74_NewCollOfColl_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_SpreadStrings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.SpreadStrings(false, "x", "y")

		// Assert
		if coc.Length() != 1 {
			t.Errorf("expected 1, got %d", coc.Length())
		}
	})
}

func Test_C74_NewCollOfColl_CloneStrings(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_CloneStrings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.CloneStrings([]string{"a", "b"})

		// Assert
		if coc == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_NewCollOfColl_Strings(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_Strings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Strings([]string{"a"})

		// Assert
		if coc == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_NewCollOfColl_StringsOption(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_StringsOption", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOption(true, 5, []string{"a"})

		// Assert
		if coc == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_NewCollOfColl_StringsOptions(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_StringsOptions", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOptions(false, 3, []string{"x"})

		// Assert
		if coc == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C74_NewCollOfColl_LenCap(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_LenCap", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 10)

		// Assert
		if coc == nil {
			t.Error("expected non-nil")
		}
	})
}
