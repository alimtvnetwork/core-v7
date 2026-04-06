package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// CloneSlice
// ==========================================================================

func Test_Cov5_CloneSlice_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov5_CloneSlice_NonEmpty", func() {
		result := corestr.CloneSlice([]string{"a", "b"})
		actual := args.Map{"len": len(result), "first": result[0]}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "CloneSlice copies -- non-empty", actual)
	})
}

func Test_Cov5_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov5_CloneSlice_Empty", func() {
		result := corestr.CloneSlice([]string{})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice empty -- empty", actual)
	})
}

// ==========================================================================
// ValidValue
// ==========================================================================

func Test_Cov5_ValidValue_Constructors(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_Constructors", func() {
		v1 := corestr.NewValidValue("hello")
		v2 := corestr.NewValidValueEmpty()
		v3 := corestr.InvalidValidValue("err")
		v4 := corestr.InvalidValidValueNoMessage()
		actual := args.Map{
			"v1Valid": v1.IsValid, "v1Value": v1.Value,
			"v2Valid": v2.IsValid, "v2Empty": v2.IsEmpty(),
			"v3Valid": v3.IsValid, "v3Msg": v3.Message,
			"v4Valid": v4.IsValid,
		}
		expected := args.Map{
			"v1Valid": true, "v1Value": "hello",
			"v2Valid": true, "v2Empty": true,
			"v3Valid": false, "v3Msg": "err",
			"v4Valid": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue constructors -- all", actual)
	})
}

func Test_Cov5_ValidValue_StringChecks(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_StringChecks", func() {
		v := corestr.NewValidValue("hello")
		actual := args.Map{
			"isEmpty":       v.IsEmpty(),
			"isWhitespace":  v.IsWhitespace(),
			"trim":          v.Trim(),
			"hasValidNE":    v.HasValidNonEmpty(),
			"hasValidNW":    v.HasValidNonWhitespace(),
			"hasSafeNE":     v.HasSafeNonEmpty(),
			"is":            v.Is("hello"),
			"isContains":    v.IsContains("ell"),
			"isEqualNoCase": v.IsEqualNonSensitive("HELLO"),
		}
		expected := args.Map{
			"isEmpty": false, "isWhitespace": false,
			"trim": "hello", "hasValidNE": true,
			"hasValidNW": true, "hasSafeNE": true,
			"is": true, "isContains": true, "isEqualNoCase": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue string checks -- hello", actual)
	})
}

func Test_Cov5_ValidValue_TypeConversions(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_TypeConversions", func() {
		vBool := corestr.NewValidValue("true")
		vInt := corestr.NewValidValue("42")
		vFloat := corestr.NewValidValue("3.14")
		vByte := corestr.NewValidValue("5")
		actual := args.Map{
			"bool":     vBool.ValueBool(),
			"int":      vInt.ValueInt(0),
			"defInt":   vInt.ValueDefInt(),
			"float":    vFloat.ValueFloat64(0) > 3,
			"defFloat": vFloat.ValueDefFloat64() > 3,
			"byte":     int(vByte.ValueByte(0)),
			"defByte":  int(vByte.ValueDefByte()),
		}
		expected := args.Map{
			"bool": true, "int": 42, "defInt": 42,
			"float": true, "defFloat": true,
			"byte": 5, "defByte": 5,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue type conversions -- various", actual)
	})
}

func Test_Cov5_ValidValue_TypeConversions_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_TypeConversions_Invalid", func() {
		v := corestr.NewValidValue("abc")
		actual := args.Map{
			"bool":   v.ValueBool(),
			"int":    v.ValueInt(99),
			"float":  v.ValueFloat64(1.0),
			"byte":   int(v.ValueByte(0)),
		}
		expected := args.Map{
			"bool": false, "int": 99, "float": 1.0, "byte": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue invalid conversions -- abc", actual)
	})
}

func Test_Cov5_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_IsAnyOf", func() {
		v := corestr.NewValidValue("hello")
		actual := args.Map{
			"match":   v.IsAnyOf("hello", "world"),
			"noMatch": v.IsAnyOf("foo", "bar"),
			"empty":   v.IsAnyOf(),
		}
		expected := args.Map{"match": true, "noMatch": false, "empty": true}
		expected.ShouldBeEqual(t, 0, "ValidValue IsAnyOf -- various", actual)
	})
}

func Test_Cov5_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_IsAnyContains", func() {
		v := corestr.NewValidValue("hello world")
		actual := args.Map{
			"match":   v.IsAnyContains("world", "xyz"),
			"noMatch": v.IsAnyContains("xyz", "abc"),
			"empty":   v.IsAnyContains(),
		}
		expected := args.Map{"match": true, "noMatch": false, "empty": true}
		expected.ShouldBeEqual(t, 0, "ValidValue IsAnyContains -- various", actual)
	})
}

func Test_Cov5_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_Split", func() {
		v := corestr.NewValidValue("a,b,c")
		result := v.Split(",")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValue Split -- comma", actual)
	})
}

func Test_Cov5_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_Clone", func() {
		v := corestr.NewValidValue("hello")
		cloned := v.Clone()
		var nilV *corestr.ValidValue
		nilClone := nilV.Clone()
		actual := args.Map{
			"val":     cloned.Value,
			"nilNil":  nilClone == nil,
		}
		expected := args.Map{"val": "hello", "nilNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue Clone -- valid and nil", actual)
	})
}

func Test_Cov5_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_String", func() {
		v := corestr.NewValidValue("hello")
		var nilV *corestr.ValidValue
		actual := args.Map{
			"string": v.String(),
			"full":   v.FullString() != "",
			"nil":    nilV.String(),
		}
		expected := args.Map{"string": "hello", "full": true, "nil": ""}
		expected.ShouldBeEqual(t, 0, "ValidValue String -- valid and nil", actual)
	})
}
func Test_Cov5_ValidValue_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_ClearDispose", func() {
		v := corestr.NewValidValue("hello")
		v.Clear()
		actual := args.Map{"empty": v.IsEmpty(), "valid": v.IsValid}
		expected := args.Map{"empty": true, "valid": false}
		expected.ShouldBeEqual(t, 0, "ValidValue Clear -- cleared", actual)
	})
}

func Test_Cov5_ValidValue_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValue_Dispose_Nil", func() {
		var v *corestr.ValidValue
		v.Dispose() // should not panic
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "ValidValue Dispose nil -- no panic", actual)
	})
}

// ==========================================================================
// ValidValues
// ==========================================================================

func Test_Cov5_ValidValues_Basics(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValues_Basics", func() {
		vv := corestr.EmptyValidValues()
		vv.Add("a").Add("b")
		actual := args.Map{
			"len":    vv.Length(),
			"count":  vv.Count(),
			"hasAny": vv.HasAnyItem(),
			"empty":  vv.IsEmpty(),
			"last":   vv.LastIndex(),
			"hasIdx": vv.HasIndex(0),
		}
		expected := args.Map{
			"len": 2, "count": 2, "hasAny": true,
			"empty": false, "last": 1, "hasIdx": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues basics -- 2 items", actual)
	})
}

func Test_Cov5_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValues_SafeValueAt", func() {
		vv := corestr.EmptyValidValues()
		vv.Add("a")
		actual := args.Map{
			"valid":   vv.SafeValueAt(0),
			"invalid": vv.SafeValueAt(99),
		}
		expected := args.Map{"valid": "a", "invalid": ""}
		expected.ShouldBeEqual(t, 0, "ValidValues SafeValueAt -- valid and invalid", actual)
	})
}

func Test_Cov5_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValues_Strings", func() {
		vv := corestr.EmptyValidValues()
		vv.Add("a").Add("b")
		actual := args.Map{
			"stringsLen": len(vv.Strings()),
			"string":     vv.String() != "",
		}
		expected := args.Map{"stringsLen": 2, "string": true}
		expected.ShouldBeEqual(t, 0, "ValidValues Strings -- 2 items", actual)
	})
}

func Test_Cov5_ValidValues_NilLength(t *testing.T) {
	safeTest(t, "Test_Cov5_ValidValues_NilLength", func() {
		var vv *corestr.ValidValues
		actual := args.Map{"len": vv.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues nil Length -- 0", actual)
	})
}

// ==========================================================================
// KeyValuePair
// ==========================================================================

func Test_Cov5_KeyValuePair_Basics(t *testing.T) {
	safeTest(t, "Test_Cov5_KeyValuePair_Basics", func() {
		kv := &corestr.KeyValuePair{Key: "name", Value: "alice"}
		actual := args.Map{
			"key":      kv.KeyName(),
			"varName":  kv.VariableName(),
			"val":      kv.ValueString(),
			"isKeyEq":  kv.IsVariableNameEqual("name"),
			"isValEq":  kv.IsValueEqual("alice"),
			"hasKey":   kv.HasKey(),
			"hasVal":   kv.HasValue(),
			"keyEmpty": kv.IsKeyEmpty(),
			"valEmpty": kv.IsValueEmpty(),
			"kvEmpty":  kv.IsKeyValueEmpty(),
			"kvAnyEm": kv.IsKeyValueAnyEmpty(),
			"compile":  kv.Compile() != "",
			"string":   kv.String() != "",
			"trimKey":  kv.TrimKey(),
			"trimVal":  kv.TrimValue(),
		}
		expected := args.Map{
			"key": "name", "varName": "name", "val": "alice",
			"isKeyEq": true, "isValEq": true,
			"hasKey": true, "hasVal": true,
			"keyEmpty": false, "valEmpty": false,
			"kvEmpty": false, "kvAnyEm": false,
			"compile": true, "string": true,
			"trimKey": "name", "trimVal": "alice",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair basics -- name:alice", actual)
	})
}

func Test_Cov5_KeyValuePair_TypeConversions(t *testing.T) {
	safeTest(t, "Test_Cov5_KeyValuePair_TypeConversions", func() {
		kvBool := &corestr.KeyValuePair{Key: "k", Value: "true"}
		kvInt := &corestr.KeyValuePair{Key: "k", Value: "42"}
		kvFloat := &corestr.KeyValuePair{Key: "k", Value: "3.14"}
		kvByte := &corestr.KeyValuePair{Key: "k", Value: "5"}
		actual := args.Map{
			"bool":     kvBool.ValueBool(),
			"int":      kvInt.ValueInt(0),
			"defInt":   kvInt.ValueDefInt(),
			"float":    kvFloat.ValueFloat64(0) > 3,
			"defFloat": kvFloat.ValueDefFloat64() > 3,
			"byte":     int(kvByte.ValueByte(0)),
			"defByte":  int(kvByte.ValueDefByte()),
		}
		expected := args.Map{
			"bool": true, "int": 42, "defInt": 42,
			"float": true, "defFloat": true,
			"byte": 5, "defByte": 5,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair type conversions -- various", actual)
	})
}

func Test_Cov5_KeyValuePair_Is(t *testing.T) {
	safeTest(t, "Test_Cov5_KeyValuePair_Is", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{
			"is":    kv.Is("k", "v"),
			"isKey": kv.IsKey("k"),
			"isVal": kv.IsVal("v"),
		}
		expected := args.Map{"is": true, "isKey": true, "isVal": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair Is methods -- k:v", actual)
	})
}

func Test_Cov5_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_Cov5_KeyValuePair_ValueValid", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()
		actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
		expected := args.Map{"val": "v", "valid": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair ValueValid -- v", actual)
	})
}

func Test_Cov5_KeyValuePair_Clear(t *testing.T) {
	safeTest(t, "Test_Cov5_KeyValuePair_Clear", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		actual := args.Map{"key": kv.Key, "val": kv.Value}
		expected := args.Map{"key": "", "val": ""}
		expected.ShouldBeEqual(t, 0, "KeyValuePair Clear -- cleared", actual)
	})
}

func Test_Cov5_KeyValuePair_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov5_KeyValuePair_Dispose_Nil", func() {
		var kv *corestr.KeyValuePair
		kv.Dispose()
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair Dispose nil -- no panic", actual)
	})
}

func Test_Cov5_KeyValuePair_FormatString(t *testing.T) {
	safeTest(t, "Test_Cov5_KeyValuePair_FormatString", func() {
		kv := &corestr.KeyValuePair{Key: "name", Value: "alice"}
		result := kv.FormatString("%s=%s")
		actual := args.Map{"result": result}
		expected := args.Map{"result": "name=alice"}
		expected.ShouldBeEqual(t, 0, "KeyValuePair FormatString -- name=alice", actual)
	})
}

func Test_Cov5_KeyValuePair_BoolEmpty(t *testing.T) {
	safeTest(t, "Test_Cov5_KeyValuePair_BoolEmpty", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: ""}
		actual := args.Map{"bool": kv.ValueBool()}
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "KeyValuePair ValueBool empty -- false", actual)
	})
}

func Test_Cov5_KeyValuePair_ByteOverflow(t *testing.T) {
	safeTest(t, "Test_Cov5_KeyValuePair_ByteOverflow", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "999"}
		actual := args.Map{"byte": int(kv.ValueByte(0))}
		expected := args.Map{"byte": 0}
		expected.ShouldBeEqual(t, 0, "KeyValuePair ValueByte overflow -- default", actual)
	})
}
