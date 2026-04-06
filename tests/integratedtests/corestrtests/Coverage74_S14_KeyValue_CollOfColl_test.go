package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": result != "host"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "KeyName expected 'host', got ''", actual)
	})
}

func Test_C74_KeyValuePair_VariableName(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_VariableName", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "port", Value: "8080"}

		// Act
		result := kv.VariableName()

		// Assert
		actual := args.Map{"result": result != "port"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'port', got ''", actual)
	})
}

func Test_C74_KeyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueString", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "val123"}

		// Act
		result := kv.ValueString()

		// Assert
		actual := args.Map{"result": result != "val123"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'val123', got ''", actual)
	})
}

func Test_C74_KeyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsVariableNameEqual", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "name", Value: "x"}

		// Act & Assert
		actual := args.Map{"result": kv.IsVariableNameEqual("name")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for matching name", actual)
		actual := args.Map{"result": kv.IsVariableNameEqual("other")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for non-matching name", actual)
	})
}

func Test_C74_KeyValuePair_IsValueEqual(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsValueEqual", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		actual := args.Map{"result": kv.IsValueEqual("abc")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv.IsValueEqual("xyz")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C74_KeyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_Compile", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act
		result := kv.Compile()

		// Assert
		actual := args.Map{"result": result != kv.String()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Compile should equal String(), got ''", actual)
	})
}

func Test_C74_KeyValuePair_IsKeyEmpty_IsValueEmpty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsKeyEmpty_IsValueEmpty", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "", Value: "x"}

		// Act & Assert
		actual := args.Map{"result": kv.IsKeyEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected key empty", actual)
		actual := args.Map{"result": kv.IsValueEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected value not empty", actual)
	})
}

func Test_C74_KeyValuePair_HasKey_HasValue(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_HasKey_HasValue", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: ""}

		// Act & Assert
		actual := args.Map{"result": kv.HasKey()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasKey true", actual)
		actual := args.Map{"result": kv.HasValue()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasValue false", actual)
	})
}

func Test_C74_KeyValuePair_IsKeyValueEmpty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsKeyValueEmpty", func() {
		// Arrange
		empty := corestr.KeyValuePair{}
		nonEmpty := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act & Assert
		actual := args.Map{"result": empty.IsKeyValueEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
		actual := args.Map{"result": nonEmpty.IsKeyValueEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for non-empty", actual)
	})
}

func Test_C74_KeyValuePair_TrimKey_TrimValue(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_TrimKey_TrimValue", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: " key ", Value: " val "}

		// Act & Assert
		actual := args.Map{"result": kv.TrimKey() != "key"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'key', got ''", actual)
		actual := args.Map{"result": kv.TrimValue() != "val"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'val', got ''", actual)
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
		actual := args.Map{"result": trueKv.ValueBool()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": falseKv.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": invalidKv.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for invalid", actual)
		actual := args.Map{"result": emptyKv.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_C74_KeyValuePair_ValueInt(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueInt", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		actual := args.Map{"result": kv.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual := args.Map{"result": invalidKv.ValueInt(99) != 99}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default 99", actual)
	})
}

func Test_C74_KeyValuePair_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueDefInt", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "10"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "x"}

		// Act & Assert
		actual := args.Map{"result": kv.ValueDefInt() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
		actual := args.Map{"result": invalidKv.ValueDefInt() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for invalid", actual)
	})
}

func Test_C74_KeyValuePair_ValueByte(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueByte", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "200"}
		overflowKv := corestr.KeyValuePair{Key: "k", Value: "999"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		actual := args.Map{"result": kv.ValueByte(0) != 200}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200", actual)
		actual := args.Map{"result": overflowKv.ValueByte(5) != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default for overflow", actual)
		actual := args.Map{"result": invalidKv.ValueByte(7) != 7}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default for invalid", actual)
	})
}

func Test_C74_KeyValuePair_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueDefByte", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "100"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "abc"}
		overflowKv := corestr.KeyValuePair{Key: "k", Value: "300"}

		// Act & Assert
		actual := args.Map{"result": kv.ValueDefByte() != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		actual := args.Map{"result": invalidKv.ValueDefByte() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual := args.Map{"result": overflowKv.ValueDefByte() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for overflow", actual)
	})
}

func Test_C74_KeyValuePair_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueFloat64", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		invalidKv := corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		actual := args.Map{"result": kv.ValueFloat64(0) != 3.14}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		actual := args.Map{"result": invalidKv.ValueFloat64(1.5) != 1.5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default 1.5", actual)
	})
}

func Test_C74_KeyValuePair_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueDefFloat64", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "2.5"}

		// Act & Assert
		actual := args.Map{"result": kv.ValueDefFloat64() != 2.5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_C74_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueValid", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "myval"}

		// Act
		vv := kv.ValueValid()

		// Assert
		actual := args.Map{"result": vv.Value != "myval" || !vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid value 'myval'", actual)
	})
}

func Test_C74_KeyValuePair_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_ValueValidOptions", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		vv := kv.ValueValidOptions(false, "err msg")

		// Assert
		actual := args.Map{"result": vv.IsValid || vv.Message != "err msg"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid with message", actual)
	})
}

func Test_C74_KeyValuePair_Is(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_Is", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act & Assert
		actual := args.Map{"result": kv.Is("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv.Is("a", "c")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C74_KeyValuePair_IsKey_IsVal(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_IsKey_IsVal", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "x", Value: "y"}

		// Act & Assert
		actual := args.Map{"result": kv.IsKey("x")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected IsKey true", actual)
		actual := args.Map{"result": kv.IsVal("y")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected IsVal true", actual)
		actual := args.Map{"result": kv.IsKey("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected IsKey false", actual)
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
		actual := args.Map{"result": full.IsKeyValueAnyEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for full", actual)
		actual := args.Map{"result": emptyKey.IsKeyValueAnyEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty key", actual)
		actual := args.Map{"result": emptyVal.IsKeyValueAnyEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty val", actual)
		actual := args.Map{"result": nilPtr.IsKeyValueAnyEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
	})
}

func Test_C74_KeyValuePair_FormatString(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_FormatString", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "host", Value: "localhost"}

		// Act
		result := kv.FormatString("%s=%s")

		// Assert
		actual := args.Map{"result": result != "host=localhost"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'host=localhost', got ''", actual)
	})
}

func Test_C74_KeyValuePair_String(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_String", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act
		result := kv.String()

		// Assert
		actual := args.Map{"result": result != "{a:b}"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '{a:b}', got ''", actual)
	})
}

func Test_C74_KeyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_Clear_Dispose", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		kv.Clear()

		// Assert
		actual := args.Map{"result": kv.Key != "" || kv.Value != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cleared", actual)
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
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "json error", actual)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "serialize error:", actual)
		actual := args.Map{"result": len(bytes) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty bytes", actual)
	})
}

func Test_C74_KeyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C74_KeyValuePair_SerializeMust", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act
		bytes := kv.SerializeMust()

		// Assert
		actual := args.Map{"result": len(bytes) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty bytes", actual)
	})
}

// --- KeyAnyValuePair ---

func Test_C74_KeyAnyValuePair_KeyName_VariableName(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_KeyName_VariableName", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "mykey", Value: 42}

		// Act & Assert
		actual := args.Map{"result": kav.KeyName() != "mykey"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "KeyName mismatch", actual)
		actual := args.Map{"result": kav.VariableName() != "mykey"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "VariableName mismatch", actual)
	})
}

func Test_C74_KeyAnyValuePair_ValueAny(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_ValueAny", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "hello"}

		// Act
		result := kav.ValueAny()

		// Assert
		actual := args.Map{"result": result != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ValueAny mismatch", actual)
	})
}

func Test_C74_KeyAnyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsVariableNameEqual", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "name", Value: nil}

		// Act & Assert
		actual := args.Map{"result": kav.IsVariableNameEqual("name")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kav.IsVariableNameEqual("other")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C74_KeyAnyValuePair_IsValueNull_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsValueNull_Nil", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act & Assert
		actual := args.Map{"result": kav.IsValueNull()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected null for nil value", actual)
	})
}

func Test_C74_KeyAnyValuePair_IsValueNull_NilReceiver(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsValueNull_NilReceiver", func() {
		// Arrange
		var kav *corestr.KeyAnyValuePair

		// Act & Assert
		actual := args.Map{"result": kav.IsValueNull()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected null for nil receiver", actual)
	})
}

func Test_C74_KeyAnyValuePair_HasNonNull(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_HasNonNull", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		kavNil := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act & Assert
		actual := args.Map{"result": kav.HasNonNull()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for non-nil value", actual)
		actual := args.Map{"result": kavNil.HasNonNull()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil value", actual)
	})
}

func Test_C74_KeyAnyValuePair_HasValue(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_HasValue", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "x"}

		// Act & Assert
		actual := args.Map{"result": kav.HasValue()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C74_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsValueEmptyString", func() {
		// Arrange
		kavEmpty := corestr.KeyAnyValuePair{Key: "k", Value: ""}
		kavNonEmpty := corestr.KeyAnyValuePair{Key: "k", Value: "abc"}

		// Act & Assert
		actual := args.Map{"result": kavEmpty.IsValueEmptyString()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty string value", actual)
		actual := args.Map{"result": kavNonEmpty.IsValueEmptyString()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for non-empty", actual)
	})
}

func Test_C74_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_IsValueWhitespace", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "  "}

		// Act & Assert
		actual := args.Map{"result": kav.IsValueWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for whitespace value string", actual)
	})
}

func Test_C74_KeyAnyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_ValueString", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 123}

		// Act
		result := kav.ValueString()

		// Assert
		actual := args.Map{"result": result != "123"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '123', got ''", actual)
	})
}

func Test_C74_KeyAnyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_Compile", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "a", Value: "b"}

		// Act
		result := kav.Compile()

		// Assert
		actual := args.Map{"result": result != kav.String()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Compile should equal String()", actual)
	})
}

func Test_C74_KeyAnyValuePair_String(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_String", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "x", Value: "y"}

		// Act
		result := kav.String()

		// Assert
		actual := args.Map{"result": result != "{x:y}"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '{x:y}', got ''", actual)
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
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "json error", actual)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "serialize error:", actual)
		actual := args.Map{"result": len(bytes) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty bytes", actual)
	})
}

func Test_C74_KeyAnyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_SerializeMust", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		bytes := kav.SerializeMust()

		// Assert
		actual := args.Map{"result": len(bytes) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty bytes", actual)
	})
}

func Test_C74_KeyAnyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_Clear_Dispose", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		kav.Clear()

		// Assert
		actual := args.Map{"result": kav.Key != "" || kav.Value != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cleared", actual)
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
		actual := args.Map{"result": binder == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil binder", actual)
	})
}

func Test_C74_KeyAnyValuePair_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_AsJsoner", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		jsoner := kav.AsJsoner()

		// Assert
		actual := args.Map{"result": jsoner == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil jsoner", actual)
	})
}

func Test_C74_KeyAnyValuePair_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C74_KeyAnyValuePair_AsJsonParseSelfInjector", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		injector := kav.AsJsonParseSelfInjector()

		// Assert
		actual := args.Map{"result": injector == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil injector", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
		actual := args.Map{"result": result.Key != "k"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected key 'k', got ''", actual)
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
		actual := args.Map{"result": result.Key != "test"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'test', got ''", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
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
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C74_KeyValueCollection_IsEmpty_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_IsEmpty_HasAnyItem", func() {
		// Arrange
		empty := corestr.New.KeyValues.Empty()
		nonEmpty := corestr.New.KeyValues.Empty()
		nonEmpty.Add("k", "v")

		// Act & Assert
		actual := args.Map{"result": empty.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": empty.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		actual := args.Map{"result": nonEmpty.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual := args.Map{"result": nonEmpty.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
	})
}

func Test_C74_KeyValueCollection_Count(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Count", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1")

		// Act & Assert
		actual := args.Map{"result": kvc.Count() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C74_KeyValueCollection_First_Last(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_First_Last", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2").Add("c", "3")

		// Act & Assert
		actual := args.Map{"result": kvc.First().Key != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "First key should be 'a'", actual)
		actual := args.Map{"result": kvc.Last().Key != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Last key should be 'c'", actual)
	})
}

func Test_C74_KeyValueCollection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_FirstOrDefault_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		result := kvc.FirstOrDefault()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil for empty collection", actual)
	})
}

func Test_C74_KeyValueCollection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_LastOrDefault_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		result := kvc.LastOrDefault()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil for empty collection", actual)
	})
}

func Test_C74_KeyValueCollection_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_LastIndex_HasIndex", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act & Assert
		actual := args.Map{"result": kvc.LastIndex() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": kvc.HasIndex(0)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for index 0", actual)
		actual := args.Map{"result": kvc.HasIndex(1)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for index 1", actual)
		actual := args.Map{"result": kvc.HasIndex(2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for index 2", actual)
	})
}

func Test_C74_KeyValueCollection_HasKey_IsContains(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_HasKey_IsContains", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("host", "localhost")

		// Act & Assert
		actual := args.Map{"result": kvc.HasKey("host")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected HasKey true", actual)
		actual := args.Map{"result": kvc.HasKey("port")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasKey false", actual)
		actual := args.Map{"result": kvc.IsContains("host")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected IsContains true", actual)
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
		actual := args.Map{"result": found || val != "8080"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found with val '8080'", actual)
		actual := args.Map{"result": notFound}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not found", actual)
	})
}

func Test_C74_KeyValueCollection_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_SafeValueAt", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act & Assert
		actual := args.Map{"result": kvc.SafeValueAt(0) != "1"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '1'", actual)
		actual := args.Map{"result": kvc.SafeValueAt(5) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for out of range", actual)
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
		actual := args.Map{"result": len(vals) != 2 || vals[0] != "1" || vals[1] != "3"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected ['1','3']", actual)
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
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 string", actual)
	})
}

func Test_C74_KeyValueCollection_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Strings_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		result := kvc.Strings()

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
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
		actual := args.Map{"result": len(result) != 1 || result[0] != "host=localhost"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'host=localhost'", actual)
	})
}

func Test_C74_KeyValueCollection_StringsUsingFormat_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_StringsUsingFormat_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		result := kvc.StringsUsingFormat("%s=%s")

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
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
		actual := args.Map{"result": len(keys) != 2 || keys[0] != "a" || keys[1] != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "keys mismatch:", actual)
		actual := args.Map{"result": len(values) != 2 || values[0] != "1" || values[1] != "2"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "values mismatch:", actual)
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
		actual := args.Map{"result": keys[0] != "a" || keys[1] != "b" || keys[2] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted keys", actual)
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

		actual := args.Map{"result": joinKeys != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
		actual := args.Map{"result": joinValues != "1,2"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '1,2', got ''", actual)
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
		actual := args.Map{"result": len(found) != 1 || found[0].Key != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected to find 'b'", actual)
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
		actual := args.Map{"result": len(found) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
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
		actual := args.Map{"result": len(found) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C74_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddStringBySplit", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddStringBySplit("=", "host=localhost")

		// Assert
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 item", actual)
	})
}

func Test_C74_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddStringBySplitTrim", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddStringBySplitTrim("=", " host = localhost ")

		// Assert
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 item", actual)
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
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_KeyValueCollection_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_Adds_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.Adds()

		// Assert
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C74_KeyValueCollection_AddMap(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddMap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddMap(map[string]string{"a": "1", "b": "2"})

		// Assert
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_KeyValueCollection_AddMap_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddMap_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddMap(nil)

		// Assert
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C74_KeyValueCollection_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddHashsetMap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddHashsetMap(map[string]bool{"x": true, "y": true})

		// Assert
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_KeyValueCollection_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddHashsetMap_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddHashsetMap(nil)

		// Assert
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_KeyValueCollection_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddHashset_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddHashset(nil)

		// Assert
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C74_KeyValueCollection_AddsHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddsHashmap_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddsHashmap(nil)

		// Assert
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_KeyValueCollection_AddsHashmaps_Nil(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AddsHashmaps_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		kvc.AddsHashmaps()

		// Assert
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hashmap length 2", actual)
		actual := args.Map{"result": len(m) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected map length 2", actual)
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
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "json error", actual)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "serialize error:", actual)
		actual := args.Map{"result": len(bytes) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty bytes", actual)
	})
}

func Test_C74_KeyValueCollection_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_MarshalUnmarshalJSON", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act
		bytes, err := kvc.MarshalJSON()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal error:", actual)

		target := corestr.New.KeyValues.Empty()
		err = target.UnmarshalJSON(bytes)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal error:", actual)
		actual := args.Map{"result": target.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1 item", actual)
	})
}

func Test_C74_KeyValueCollection_AsJsoner_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AsJsoner_AsJsonContractsBinder", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act & Assert
		actual := args.Map{"result": kvc.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil jsoner", actual)
		actual := args.Map{"result": kvc.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil binder", actual)
	})
}

func Test_C74_KeyValueCollection_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C74_KeyValueCollection_AsJsonParseSelfInjector", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		injector := kvc.AsJsonParseSelfInjector()

		// Assert
		actual := args.Map{"result": injector == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
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
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 after clear", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "deserialize error:", actual)
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
		actual := args.Map{"result": len(bytes) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty bytes", actual)
	})
}

// --- newKeyValuesCreator ---

func Test_C74_NewKeyValues_Cap(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_Cap", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.Cap(5)

		// Assert
		actual := args.Map{"result": kvc == nil || kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty with capacity", actual)
	})
}

func Test_C74_NewKeyValues_UsingMap(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingMap", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Assert
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_NewKeyValues_UsingMap_Empty(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingMap_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{})

		// Assert
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C74_NewKeyValues_UsingKeyValuePairs(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingKeyValuePairs", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValuePairs(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Assert
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C74_NewKeyValues_UsingKeyValuePairs_Empty(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingKeyValuePairs_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValuePairs()

		// Assert
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_NewKeyValues_UsingKeyValueStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C74_NewKeyValues_UsingKeyValueStrings_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValueStrings([]string{}, []string{})

		// Assert
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// --- CollectionsOfCollection ---

func Test_C74_CollOfColl_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_IsEmpty_HasItems", func() {
		// Arrange
		empty := corestr.New.CollectionsOfCollection.Empty()
		nonEmpty := corestr.New.CollectionsOfCollection.SpreadStrings(false, "a", "b")

		// Act & Assert
		actual := args.Map{"result": empty.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": empty.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		actual := args.Map{"result": nonEmpty.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_C74_CollOfColl_Length(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_Length", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		actual := args.Map{"result": coc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": total != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C74_CollOfColl_AllIndividualItemsLength_Empty(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AllIndividualItemsLength_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		actual := args.Map{"result": coc.AllIndividualItemsLength() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": len(list) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C74_CollOfColl_List_Empty(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_List_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		list := coc.List(5)

		// Assert
		actual := args.Map{"result": len(list) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_CollOfColl_AddStrings(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AddStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddStrings(false, []string{"a", "b"})

		// Assert
		actual := args.Map{"result": coc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C74_CollOfColl_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AddStrings_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddStrings(false, []string{})

		// Assert
		actual := args.Map{"result": coc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C74_CollOfColl_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AddsStringsOfStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b", "c"})

		// Assert
		actual := args.Map{"result": coc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_CollOfColl_AddsStringsOfStrings_Nil(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AddsStringsOfStrings_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddsStringsOfStrings(false)

		// Assert
		actual := args.Map{"result": coc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": coc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C74_CollOfColl_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_Adds_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		coc.AddCollections()

		// Assert
		actual := args.Map{"result": coc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": coc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
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
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "json error:", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal error:", actual)

		target := corestr.New.CollectionsOfCollection.Empty()
		err = target.UnmarshalJSON(bytes)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal error:", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	})
}

func Test_C74_CollOfColl_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AsJsoner", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		actual := args.Map{"result": coc.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C74_CollOfColl_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AsJsonParseSelfInjector", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		actual := args.Map{"result": coc.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C74_CollOfColl_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AsJsonMarshaller", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		actual := args.Map{"result": coc.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C74_CollOfColl_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C74_CollOfColl_AsJsonContractsBinder", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act & Assert
		actual := args.Map{"result": coc.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	})
}

// --- newCollectionsOfCollectionCreator ---

func Test_C74_NewCollOfColl_Cap(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_Cap", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Cap(5)

		// Assert
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C74_NewCollOfColl_StringsOfStrings(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_StringsOfStrings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})

		// Assert
		actual := args.Map{"result": coc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C74_NewCollOfColl_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_SpreadStrings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.SpreadStrings(false, "x", "y")

		// Assert
		actual := args.Map{"result": coc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C74_NewCollOfColl_CloneStrings(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_CloneStrings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.CloneStrings([]string{"a", "b"})

		// Assert
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C74_NewCollOfColl_Strings(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_Strings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Strings([]string{"a"})

		// Assert
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C74_NewCollOfColl_StringsOption(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_StringsOption", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOption(true, 5, []string{"a"})

		// Assert
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C74_NewCollOfColl_StringsOptions(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_StringsOptions", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOptions(false, 3, []string{"x"})

		// Assert
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C74_NewCollOfColl_LenCap(t *testing.T) {
	safeTest(t, "Test_C74_NewCollOfColl_LenCap", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 10)

		// Assert
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}
