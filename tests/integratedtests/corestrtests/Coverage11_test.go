package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// ValidValue — constructors & basic
// ═══════════════════════════════════════════

func Test_Cov11_ValidValue_New(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_New", func() {
		vv := corestr.NewValidValue("hello")
		vvEmpty := corestr.NewValidValueEmpty()
		vvInvalid := corestr.InvalidValidValue("err-msg")
		vvInvalidNM := corestr.InvalidValidValueNoMessage()
		actual := args.Map{
			"val":        vv.Value,
			"isValid":    vv.IsValid,
			"emptyVal":   vvEmpty.Value,
			"emptyValid": vvEmpty.IsValid,
			"invMsg":     vvInvalid.Message,
			"invValid":   vvInvalid.IsValid,
			"invNMValid": vvInvalidNM.IsValid,
		}
		expected := args.Map{
			"val": "hello", "isValid": true,
			"emptyVal": "", "emptyValid": true,
			"invMsg": "err-msg", "invValid": false,
			"invNMValid": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- constructors", actual)
	})
}

func Test_Cov11_ValidValue_UsingAny(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_UsingAny", func() {
		vv := corestr.NewValidValueUsingAny(false, true, 42)
		vvAuto := corestr.NewValidValueUsingAnyAutoValid(false, "hello")
		actual := args.Map{
			"val":       vv.Value != "",
			"isValid":   vv.IsValid,
			"autoValid": vvAuto.IsValid,
		}
		expected := args.Map{"val": true, "isValid": true, "autoValid": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- UsingAny", actual)
	})
}

func Test_Cov11_ValidValue_State(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_State", func() {
		vv := corestr.NewValidValue("hello")
		vvEmpty := corestr.NewValidValueEmpty()
		actual := args.Map{
			"isEmpty":          vv.IsEmpty(),
			"emptyIsEmpty":     vvEmpty.IsEmpty(),
			"isWhitespace":     vv.IsWhitespace(),
			"trim":             vv.Trim(),
			"hasValidNonEmpty": vv.HasValidNonEmpty(),
			"hasValidNonWS":    vv.HasValidNonWhitespace(),
			"hasSafeNonEmpty":  vv.HasSafeNonEmpty(),
			"is":               vv.Is("hello"),
			"isNot":            vv.Is("world"),
		}
		expected := args.Map{
			"isEmpty": false, "emptyIsEmpty": true,
			"isWhitespace": false, "trim": "hello",
			"hasValidNonEmpty": true, "hasValidNonWS": true,
			"hasSafeNonEmpty": true, "is": true, "isNot": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- state", actual)
	})
}

func Test_Cov11_ValidValue_Conversions(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_Conversions", func() {
		vvBool := corestr.NewValidValue("true")
		vvInt := corestr.NewValidValue("42")
		vvFloat := corestr.NewValidValue("3.14")
		vvByte := corestr.NewValidValue("200")
		vvBadInt := corestr.NewValidValue("abc")
		actual := args.Map{
			"bool":       vvBool.ValueBool(),
			"int":        vvInt.ValueInt(0),
			"defInt":     vvInt.ValueDefInt(),
			"float":      vvFloat.ValueFloat64(0) > 3.0,
			"defFloat":   vvFloat.ValueDefFloat64() > 3.0,
			"byte":       vvByte.ValueByte(0),
			"defByte":    vvByte.ValueDefByte(),
			"badBool":    vvBadInt.ValueBool(),
			"badInt":     vvBadInt.ValueInt(99),
			"badDefInt":  vvBadInt.ValueDefInt(),
			"badByte":    vvBadInt.ValueByte(55),
			"badDefByte": vvBadInt.ValueDefByte(),
		}
		expected := args.Map{
			"bool": true, "int": 42, "defInt": 42,
			"float": true, "defFloat": true,
			"byte": byte(200), "defByte": byte(200),
			"badBool": false, "badInt": 99, "badDefInt": 0,
			"badByte": byte(0), "badDefByte": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- conversions", actual)
	})
}

func Test_Cov11_ValidValue_HighByte(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_HighByte", func() {
		vv := corestr.NewValidValue("300")
		vvNeg := corestr.NewValidValue("-1")
		actual := args.Map{
			"highByte": vv.ValueByte(0),
			"negByte":  vvNeg.ValueByte(0),
		}
		expected := args.Map{"highByte": byte(255), "negByte": byte(0)}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- high/neg byte", actual)
	})
}

func Test_Cov11_ValidValue_Search(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_Search", func() {
		vv := corestr.NewValidValue("hello world")
		re := regexp.MustCompile(`\w+`)
		actual := args.Map{
			"isAnyOf":       vv.IsAnyOf("hello world", "x"),
			"isAnyOfEmpty":  vv.IsAnyOf(),
			"isAnyOfFalse":  vv.IsAnyOf("x", "y"),
			"isContains":    vv.IsContains("world"),
			"isAnyContains": vv.IsAnyContains("xyz", "world"),
			"isAnyContE":    vv.IsAnyContains(),
			"equalFold":     vv.IsEqualNonSensitive("HELLO WORLD"),
			"regexMatch":    vv.IsRegexMatches(re),
			"regexNil":      vv.IsRegexMatches(nil),
			"regexFind":     vv.RegexFindString(re),
			"regexFindNil":  vv.RegexFindString(nil),
		}
		expected := args.Map{
			"isAnyOf": true, "isAnyOfEmpty": true, "isAnyOfFalse": false,
			"isContains": true, "isAnyContains": true, "isAnyContE": true,
			"equalFold": true, "regexMatch": true, "regexNil": false,
			"regexFind": "hello", "regexFindNil": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- search", actual)
	})
}

func Test_Cov11_ValidValue_RegexFindAll(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_RegexFindAll", func() {
		vv := corestr.NewValidValue("abc 123 def")
		re := regexp.MustCompile(`\d+`)
		items, hasAny := vv.RegexFindAllStringsWithFlag(re, -1)
		itemsNil, hasAnyNil := vv.RegexFindAllStringsWithFlag(nil, -1)
		allItems := vv.RegexFindAllStrings(re, -1)
		allNil := vv.RegexFindAllStrings(nil, -1)
		actual := args.Map{
			"items":    len(items),
			"hasAny":   hasAny,
			"nilItems": len(itemsNil),
			"nilHas":   hasAnyNil,
			"allItems": len(allItems),
			"allNil":   len(allNil),
		}
		expected := args.Map{
			"items": 1, "hasAny": true,
			"nilItems": 0, "nilHas": false,
			"allItems": 1, "allNil": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- regex find all", actual)
	})
}

func Test_Cov11_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,,c")
		actual := args.Map{
			"splitLen":    len(vv.Split(",")),
			"nonEmptyLen": len(vv.SplitNonEmpty(",")),
			"trimNonWS":  len(vv.SplitTrimNonWhitespace(",")),
		}
		expected := args.Map{"splitLen": 4, "nonEmptyLen": 4, "trimNonWS": 4}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- split", actual)
	})
}
func Test_Cov11_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_Clone", func() {
		vv := corestr.NewValidValue("hello")
		cloned := vv.Clone()
		var nilVV *corestr.ValidValue
		nilClone := nilVV.Clone()
		actual := args.Map{
			"clonedVal":   cloned.Value,
			"clonedValid": cloned.IsValid,
			"nilClone":    nilClone == nil,
		}
		expected := args.Map{"clonedVal": "hello", "clonedValid": true, "nilClone": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- clone", actual)
	})
}

func Test_Cov11_ValidValue_StringMethods(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_StringMethods", func() {
		vv := corestr.NewValidValue("hello")
		var nilVV *corestr.ValidValue
		actual := args.Map{
			"str":      vv.String(),
			"fullStr":  vv.FullString() != "",
			"nilStr":   nilVV.String(),
			"nilFull":  nilVV.FullString(),
		}
		expected := args.Map{"str": "hello", "fullStr": true, "nilStr": "", "nilFull": ""}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- string methods", actual)
	})
}

func Test_Cov11_ValidValue_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_ClearDispose", func() {
		vv := corestr.NewValidValue("hello")
		vv.Clear()
		var nilVV *corestr.ValidValue
		nilVV.Clear()    // should not panic
		nilVV.Dispose()  // should not panic
		actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
		expected := args.Map{"val": "", "valid": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- clear/dispose", actual)
	})
}

func Test_Cov11_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_Json", func() {
		vv := corestr.NewValidValue("hello")
		j := vv.Json()
		jp := vv.JsonPtr()
		actual := args.Map{
			"jsonLen":  j.Length() > 0,
			"jpNotNil": jp != nil,
		}
		expected := args.Map{"jsonLen": true, "jpNotNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- json", actual)
	})
}

func Test_Cov11_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValue_Serialize", func() {
		vv := corestr.NewValidValue("hello")
		bytes, err := vv.Serialize()
		actual := args.Map{"hasBytes": len(bytes) > 0, "errNil": err == nil}
		expected := args.Map{"hasBytes": true, "errNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- serialize", actual)
	})
}

// ═══════════════════════════════════════════
// ValidValues
// ═══════════════════════════════════════════

func Test_Cov11_ValidValues_Basic(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_Basic", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		vvs.AddFull(false, "c", "err")
		empty := corestr.EmptyValidValues()
		actual := args.Map{
			"len":      vvs.Length(),
			"count":    vvs.Count(),
			"hasAny":   vvs.HasAnyItem(),
			"isEmpty":  vvs.IsEmpty(),
			"emptyLen": empty.Length(),
			"lastIdx":  vvs.LastIndex(),
			"hasIdx":   vvs.HasIndex(0),
		}
		expected := args.Map{
			"len": 3, "count": 3, "hasAny": true,
			"isEmpty": false, "emptyLen": 0, "lastIdx": 2,
			"hasIdx": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- basic", actual)
	})
}

func Test_Cov11_ValidValues_UsingValues(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_UsingValues", func() {
		v1 := corestr.ValidValue{Value: "x", IsValid: true}
		vvs := corestr.NewValidValuesUsingValues(v1)
		emptyVVS := corestr.NewValidValuesUsingValues()
		actual := args.Map{
			"len":      vvs.Length(),
			"emptyLen": emptyVVS.Length(),
		}
		expected := args.Map{"len": 1, "emptyLen": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- using values", actual)
	})
}

func Test_Cov11_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_SafeValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").AddFull(true, "b", "")
		actual := args.Map{
			"at0":       vvs.SafeValueAt(0),
			"at1":       vvs.SafeValueAt(1),
			"atOOB":     vvs.SafeValueAt(99),
			"validAt0":  vvs.SafeValidValueAt(0),
			"validAt1":  vvs.SafeValidValueAt(1),
			"validOOB":  vvs.SafeValidValueAt(99),
		}
		expected := args.Map{
			"at0": "a", "at1": "b", "atOOB": "",
			"validAt0": "a", "validAt1": "b", "validOOB": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- safe value at", actual)
	})
}

func Test_Cov11_ValidValues_Indexes(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_Indexes", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b").Add("c")
		vals := vvs.SafeValuesAtIndexes(0, 2)
		validVals := vvs.SafeValidValuesAtIndexes(0, 2)
		emptyVals := vvs.SafeValuesAtIndexes()
		actual := args.Map{
			"valsLen":      len(vals),
			"val0":         vals[0],
			"val1":         vals[1],
			"validLen":     len(validVals),
			"emptyValsLen": len(emptyVals),
		}
		expected := args.Map{
			"valsLen": 2, "val0": "a", "val1": "c",
			"validLen": 2, "emptyValsLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- indexes", actual)
	})
}

func Test_Cov11_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_Strings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		strs := vvs.Strings()
		fullStrs := vvs.FullStrings()
		str := vvs.String()
		actual := args.Map{
			"strsLen":     len(strs),
			"fullStrsLen": len(fullStrs),
			"strNE":       str != "",
		}
		expected := args.Map{"strsLen": 2, "fullStrsLen": 2, "strNE": true}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- strings", actual)
	})
}

func Test_Cov11_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_Find", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b").Add("c")
		found := vvs.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, v.Value == "b", false
		})
		actual := args.Map{"foundLen": len(found)}
		expected := args.Map{"foundLen": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- find", actual)
	})
}

func Test_Cov11_ValidValues_Adds(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_Adds", func() {
		vvs := corestr.NewValidValues(5)
		v1 := corestr.ValidValue{Value: "x", IsValid: true}
		vvs.Adds(v1)
		vvs.Adds() // empty
		vvs.AddsPtr(&v1)
		vvs.AddsPtr() // empty
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- adds", actual)
	})
}

func Test_Cov11_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_ConcatNew", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		concat := vvs.ConcatNew(true, vvs2)
		concatEmpty := vvs.ConcatNew(true)
		concatNoClone := vvs.ConcatNew(false)
		actual := args.Map{
			"concatLen":      concat.Length(),
			"concatEmptyLen": concatEmpty.Length(),
			"noCloneLen":     concatNoClone.Length(),
		}
		expected := args.Map{"concatLen": 2, "concatEmptyLen": 1, "noCloneLen": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- concat new", actual)
	})
}

func Test_Cov11_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_AddValidValues", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		other := corestr.NewValidValues(5)
		other.Add("b")
		vvs.AddValidValues(other)
		vvs.AddValidValues(nil)
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- add valid values", actual)
	})
}

func Test_Cov11_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_AddHashset", func() {
		vvs := corestr.NewValidValues(5)
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		vvs.AddHashset(hs)
		vvs.AddHashset(nil)
		vvs.AddHashsetMap(map[string]bool{"c": true, "d": false})
		vvs.AddHashsetMap(nil)
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- add hashset", actual)
	})
}

func Test_Cov11_ValidValues_HashmapMap(t *testing.T) {
	safeTest(t, "Test_Cov11_ValidValues_HashmapMap", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").AddFull(true, "b", "msg")
		hm := vvs.Hashmap()
		m := vvs.Map()
		actual := args.Map{"hmLen": hm.Length(), "mapLen": len(m)}
		expected := args.Map{"hmLen": 2, "mapLen": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- hashmap/map", actual)
	})
}

// ═══════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════

func Test_Cov11_ValueStatus_Basic(t *testing.T) {
	safeTest(t, "Test_Cov11_ValueStatus_Basic", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		vs2 := corestr.InvalidValueStatus("err")
		cloned := vs2.Clone()
		actual := args.Map{
			"vsValid":     vs.ValueValid.IsValid,
			"vs2Msg":      vs2.ValueValid.Message,
			"clonedMsg":   cloned.ValueValid.Message,
			"clonedIdx":   cloned.Index,
		}
		expected := args.Map{
			"vsValid": false, "vs2Msg": "err",
			"clonedMsg": "err", "clonedIdx": -1,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- basic", actual)
	})
}

// ═══════════════════════════════════════════
// KeyValuePair — comprehensive
// ═══════════════════════════════════════════

func Test_Cov11_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValuePair_Basic", func() {
		kv := corestr.KeyValuePair{Key: "name", Value: "test"}
		actual := args.Map{
			"key":        kv.KeyName(),
			"varName":    kv.VariableName(),
			"valStr":     kv.ValueString(),
			"isVarName":  kv.IsVariableNameEqual("name"),
			"isValEqual": kv.IsValueEqual("test"),
			"isKeyEmpty": kv.IsKeyEmpty(),
			"isValEmpty": kv.IsValueEmpty(),
			"hasKey":     kv.HasKey(),
			"hasVal":     kv.HasValue(),
			"isKVEmpty":  kv.IsKeyValueEmpty(),
			"compile":    kv.Compile() != "",
			"str":        kv.String() != "",
		}
		expected := args.Map{
			"key": "name", "varName": "name", "valStr": "test",
			"isVarName": true, "isValEqual": true,
			"isKeyEmpty": false, "isValEmpty": false,
			"hasKey": true, "hasVal": true, "isKVEmpty": false,
			"compile": true, "str": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- basic", actual)
	})
}

func Test_Cov11_KeyValuePair_Conversions(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValuePair_Conversions", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		kvBool := corestr.KeyValuePair{Key: "k", Value: "true"}
		kvBad := corestr.KeyValuePair{Key: "k", Value: "abc"}
		actual := args.Map{
			"int":       kv.ValueInt(0),
			"defInt":    kv.ValueDefInt(),
			"bool":      kvBool.ValueBool(),
			"badBool":   kvBad.ValueBool(),
			"badInt":    kvBad.ValueInt(99),
			"byte":      kv.ValueByte(0),
			"defByte":   kv.ValueDefByte(),
			"badByte":   kvBad.ValueByte(55),
			"float":     kv.ValueFloat64(0) > 0,
			"defFloat":  kv.ValueDefFloat64() > 0,
			"trimKey":   kv.TrimKey(),
			"trimVal":   kv.TrimValue(),
		}
		expected := args.Map{
			"int": 42, "defInt": 42, "bool": true,
			"badBool": false, "badInt": 99,
			"byte": byte(42), "defByte": byte(42), "badByte": byte(55),
			"float": true, "defFloat": true,
			"trimKey": "k", "trimVal": "42",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- conversions", actual)
	})
}

func Test_Cov11_KeyValuePair_Is(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValuePair_Is", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{
			"is":        kv.Is("k", "v"),
			"isKey":     kv.IsKey("k"),
			"isVal":     kv.IsVal("v"),
			"isAnyE":    kv.IsKeyValueAnyEmpty(),
			"formatStr": kv.FormatString("%s=%s") != "",
		}
		expected := args.Map{
			"is": true, "isKey": true, "isVal": true,
			"isAnyE": false, "formatStr": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- is", actual)
	})
}

func Test_Cov11_KeyValuePair_ValValid(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValuePair_ValValid", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()
		vvOpt := kv.ValueValidOptions(false, "msg")
		actual := args.Map{
			"vvVal":   vv.Value,
			"vvValid": vv.IsValid,
			"optVal":  vvOpt.Value,
			"optMsg":  vvOpt.Message,
		}
		expected := args.Map{
			"vvVal": "v", "vvValid": true,
			"optVal": "v", "optMsg": "msg",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns non-empty -- value valid", actual)
	})
}

func Test_Cov11_KeyValuePair_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValuePair_ClearDispose", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		var nilKV *corestr.KeyValuePair
		nilKV.Clear()
		nilKV.Dispose()
		actual := args.Map{"key": kv.Key, "val": kv.Value}
		expected := args.Map{"key": "", "val": ""}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- clear/dispose", actual)
	})
}

func Test_Cov11_KeyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValuePair_Json", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		j := kv.Json()
		jp := kv.JsonPtr()
		bytes, err := kv.Serialize()
		must := kv.SerializeMust()
		actual := args.Map{
			"jsonLen":  j.Length() > 0,
			"jpNotNil": jp != nil,
			"bytesLen": len(bytes) > 0,
			"errNil":   err == nil,
			"mustLen":  len(must) > 0,
		}
		expected := args.Map{
			"jsonLen": true, "jpNotNil": true,
			"bytesLen": true, "errNil": true, "mustLen": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- json", actual)
	})
}

func Test_Cov11_KeyValuePair_HighByte(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValuePair_HighByte", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "300"}
		kvBad := corestr.KeyValuePair{Key: "k", Value: "abc"}
		actual := args.Map{
			"highByte":   kv.ValueByte(0),
			"highDefB":   kv.ValueDefByte(),
			"badByte":    kvBad.ValueByte(5),
			"badDefByte": kvBad.ValueDefByte(),
		}
		expected := args.Map{
			"highByte": byte(0), "highDefB": byte(0),
			"badByte": byte(5), "badDefByte": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- high byte", actual)
	})
}

// ═══════════════════════════════════════════
// KeyAnyValuePair
// ═══════════════════════════════════════════

func Test_Cov11_KeyAnyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyAnyValuePair_Basic", func() {
		kav := corestr.KeyAnyValuePair{Key: "name", Value: 42}
		actual := args.Map{
			"key":       kav.KeyName(),
			"varName":   kav.VariableName(),
			"isVarName": kav.IsVariableNameEqual("name"),
			"hasNonNil": kav.HasNonNull(),
			"hasValue":  kav.HasValue(),
			"isValNull": kav.IsValueNull(),
			"valStr":    kav.ValueString() != "",
			"compile":   kav.Compile() != "",
			"str":       kav.String() != "",
		}
		expected := args.Map{
			"key": "name", "varName": "name", "isVarName": true,
			"hasNonNil": true, "hasValue": true, "isValNull": false,
			"valStr": true, "compile": true, "str": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- basic", actual)
	})
}

func Test_Cov11_KeyAnyValuePair_NilValue(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyAnyValuePair_NilValue", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}
		actual := args.Map{
			"isNull":    kav.IsValueNull(),
			"isEmptyS":  kav.IsValueEmptyString(),
			"isWS":      kav.IsValueWhitespace(),
			"valStr":    kav.ValueString(),
		}
		expected := args.Map{
			"isNull": true, "isEmptyS": true, "isWS": true, "valStr": "",
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns nil -- nil value", actual)
	})
}

func Test_Cov11_KeyAnyValuePair_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyAnyValuePair_ClearDispose", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		kav.Clear()
		var nilKAV *corestr.KeyAnyValuePair
		nilKAV.Clear()
		nilKAV.Dispose()
		actual := args.Map{"key": kav.Key, "valNil": kav.Value == nil}
		expected := args.Map{"key": "", "valNil": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- clear/dispose", actual)
	})
}

func Test_Cov11_KeyAnyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyAnyValuePair_Json", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kav.Json()
		jp := kav.JsonPtr()
		bytes, err := kav.Serialize()
		must := kav.SerializeMust()
		binder := kav.AsJsonContractsBinder()
		jsoner := kav.AsJsoner()
		injector := kav.AsJsonParseSelfInjector()
		actual := args.Map{
			"jsonLen": j.Length() > 0, "jpNotNil": jp != nil,
			"bytesLen": len(bytes) > 0, "errNil": err == nil,
			"mustLen": len(must) > 0,
			"binderNN": binder != nil, "jsonerNN": jsoner != nil,
			"injectorNN": injector != nil,
		}
		expected := args.Map{
			"jsonLen": true, "jpNotNil": true,
			"bytesLen": true, "errNil": true, "mustLen": true,
			"binderNN": true, "jsonerNN": true, "injectorNN": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- json", actual)
	})
}

// ═══════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════

func Test_Cov11_TextWithLineNumber_Basic(t *testing.T) {
	safeTest(t, "Test_Cov11_TextWithLineNumber_Basic", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
		twEmpty := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		var nilTW *corestr.TextWithLineNumber
		actual := args.Map{
			"hasLine":   tw.HasLineNumber(),
			"invalid":   tw.IsInvalidLineNumber(),
			"len":       tw.Length(),
			"isEmpty":   tw.IsEmpty(),
			"isEmptyT":  tw.IsEmptyText(),
			"isEmptyBo": tw.IsEmptyTextLineBoth(),
			"emptyHas":  twEmpty.HasLineNumber(),
			"emptyInv":  twEmpty.IsInvalidLineNumber(),
			"emptyIsE":  twEmpty.IsEmpty(),
			"nilLen":    nilTW.Length(),
			"nilIsE":    nilTW.IsEmpty(),
			"nilText":   nilTW.IsEmptyText(),
		}
		expected := args.Map{
			"hasLine": true, "invalid": false, "len": 5,
			"isEmpty": false, "isEmptyT": false, "isEmptyBo": false,
			"emptyHas": false, "emptyInv": true, "emptyIsE": true,
			"nilLen": 0, "nilIsE": true, "nilText": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns non-empty -- with args", actual)
	})
}

// ═══════════════════════════════════════════
// HashmapDiff
// ═══════════════════════════════════════════

func Test_Cov11_HashmapDiff_Basic(t *testing.T) {
	safeTest(t, "Test_Cov11_HashmapDiff_Basic", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})
		var nilHD *corestr.HashmapDiff
		actual := args.Map{
			"len":      hd.Length(),
			"isEmpty":  hd.IsEmpty(),
			"hasAny":   hd.HasAnyItem(),
			"lastIdx":  hd.LastIndex(),
			"nilLen":   nilHD.Length(),
			"rawLen":   len(hd.Raw()),
			"nilRaw":   len(nilHD.Raw()),
			"mapAny":   len(hd.MapAnyItems()),
			"nilMap":   len(nilHD.MapAnyItems()),
			"keysLen":  len(hd.AllKeysSorted()),
		}
		expected := args.Map{
			"len": 2, "isEmpty": false, "hasAny": true, "lastIdx": 1,
			"nilLen": 0, "rawLen": 2, "nilRaw": 0, "mapAny": 2,
			"nilMap": 0, "keysLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- basic", actual)
	})
}

func Test_Cov11_HashmapDiff_Compare(t *testing.T) {
	safeTest(t, "Test_Cov11_HashmapDiff_Compare", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})
		same := map[string]string{"a": "1", "b": "2"}
		diff := map[string]string{"a": "1", "b": "3"}
		actual := args.Map{
			"isEqual":    hd.IsRawEqual(same),
			"isNotEqual": hd.IsRawEqual(diff),
			"hasChanges": hd.HasAnyChanges(diff),
			"noChanges":  hd.HasAnyChanges(same),
		}
		expected := args.Map{
			"isEqual": true, "isNotEqual": false,
			"hasChanges": true, "noChanges": false,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- compare", actual)
	})
}

func Test_Cov11_HashmapDiff_Diff(t *testing.T) {
	safeTest(t, "Test_Cov11_HashmapDiff_Diff", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})
		diff := map[string]string{"a": "1", "b": "3"}
		diffRaw := hd.DiffRaw(diff)
		diffHM := hd.HashmapDiffUsingRaw(diff)
		diffJson := hd.DiffJsonMessage(diff)
		diffSlice := hd.ToStringsSliceOfDiffMap(diffRaw)
		shouldMsg := hd.ShouldDiffMessage("test", diff)
		logMsg := hd.LogShouldDiffMessage("test", diff)
		rawMapDiff := hd.RawMapStringAnyDiff()
		actual := args.Map{
			"diffRawLen":   len(diffRaw),
			"diffHMLen":    diffHM.Length(),
			"diffJsonNE":   diffJson != "",
			"diffSliceLen": len(diffSlice) > 0,
			"shouldMsgNE":  shouldMsg != "",
			"logMsgNE":     logMsg != "",
			"rawMapLen":    len(rawMapDiff),
		}
		expected := args.Map{
			"diffRawLen": 1, "diffHMLen": 1,
			"diffJsonNE": true, "diffSliceLen": true,
			"shouldMsgNE": true, "logMsgNE": true,
			"rawMapLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- diff", actual)
	})
}

func Test_Cov11_HashmapDiff_NoDiff(t *testing.T) {
	safeTest(t, "Test_Cov11_HashmapDiff_NoDiff", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		same := map[string]string{"a": "1"}
		diffHM := hd.HashmapDiffUsingRaw(same)
		actual := args.Map{"diffLen": diffHM.Length()}
		expected := args.Map{"diffLen": 0}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns empty -- no diff", actual)
	})
}

func Test_Cov11_HashmapDiff_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov11_HashmapDiff_Serialize", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		bytes, err := hd.Serialize()
		actual := args.Map{"hasBytes": len(bytes) > 0, "errNil": err == nil}
		expected := args.Map{"hasBytes": true, "errNil": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- serialize", actual)
	})
}

// ═══════════════════════════════════════════
// LeftRight — string specific
// ═══════════════════════════════════════════

func Test_Cov11_LeftRight_Constructors(t *testing.T) {
	safeTest(t, "Test_Cov11_LeftRight_Constructors", func() {
		lr := corestr.NewLeftRight("left", "right")
		lrInv := corestr.InvalidLeftRight("err")
		lrInvNM := corestr.InvalidLeftRightNoMessage()
		lrSlice := corestr.LeftRightUsingSlice([]string{"a", "b"})
		lrSlice1 := corestr.LeftRightUsingSlice([]string{"a"})
		lrSlice0 := corestr.LeftRightUsingSlice([]string{})
		lrSlicePtr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		lrSlicePtrE := corestr.LeftRightUsingSlicePtr([]string{})
		lrTrim := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		lrTrim1 := corestr.LeftRightTrimmedUsingSlice([]string{"a"})
		lrTrimNil := corestr.LeftRightTrimmedUsingSlice(nil)
		actual := args.Map{
			"lrLeft":      lr.Left,
			"lrRight":     lr.Right,
			"lrValid":     lr.IsValid,
			"invValid":    lrInv.IsValid,
			"invNMValid":  lrInvNM.IsValid,
			"sliceLeft":   lrSlice.Left,
			"sliceRight":  lrSlice.Right,
			"sliceValid":  lrSlice.IsValid,
			"slice1Left":  lrSlice1.Left,
			"slice1Valid": lrSlice1.IsValid,
			"slice0Valid": lrSlice0.IsValid,
			"ptrLeft":     lrSlicePtr.Left,
			"ptrEValid":   lrSlicePtrE.IsValid,
			"trimLeft":    lrTrim.Left,
			"trimRight":   lrTrim.Right,
			"trim1Left":   lrTrim1.Left,
			"trimNilValid": lrTrimNil.IsValid,
		}
		expected := args.Map{
			"lrLeft": "left", "lrRight": "right", "lrValid": true,
			"invValid": false, "invNMValid": false,
			"sliceLeft": "a", "sliceRight": "b", "sliceValid": true,
			"slice1Left": "a", "slice1Valid": false,
			"slice0Valid": false,
			"ptrLeft": "a", "ptrEValid": false,
			"trimLeft": "a", "trimRight": "b",
			"trim1Left": "a", "trimNilValid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- constructors", actual)
	})
}

func Test_Cov11_LeftRight_Methods(t *testing.T) {
	safeTest(t, "Test_Cov11_LeftRight_Methods", func() {
		lr := corestr.NewLeftRight("left", "right")
		re := regexp.MustCompile(`left`)
		actual := args.Map{
			"leftBytes":   len(lr.LeftBytes()) > 0,
			"rightBytes":  len(lr.RightBytes()) > 0,
			"leftTrim":    lr.LeftTrim(),
			"rightTrim":   lr.RightTrim(),
			"isLeftEmpty": lr.IsLeftEmpty(),
			"isRightE":    lr.IsRightEmpty(),
			"isLeftWS":    lr.IsLeftWhitespace(),
			"isRightWS":   lr.IsRightWhitespace(),
			"hasVNEL":     lr.HasValidNonEmptyLeft(),
			"hasVNER":     lr.HasValidNonEmptyRight(),
			"hasVNWSL":    lr.HasValidNonWhitespaceLeft(),
			"hasVNWSR":    lr.HasValidNonWhitespaceRight(),
			"hasSafe":     lr.HasSafeNonEmpty(),
			"isLeft":      lr.IsLeft("left"),
			"isRight":     lr.IsRight("right"),
			"is":          lr.Is("left", "right"),
			"reLeft":      lr.IsLeftRegexMatch(re),
			"reRight":     lr.IsRightRegexMatch(re),
			"reNil":       lr.IsLeftRegexMatch(nil),
		}
		expected := args.Map{
			"leftBytes": true, "rightBytes": true,
			"leftTrim": "left", "rightTrim": "right",
			"isLeftEmpty": false, "isRightE": false,
			"isLeftWS": false, "isRightWS": false,
			"hasVNEL": true, "hasVNER": true,
			"hasVNWSL": true, "hasVNWSR": true,
			"hasSafe": true, "isLeft": true, "isRight": true,
			"is": true, "reLeft": true, "reRight": false, "reNil": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- methods", actual)
	})
}

func Test_Cov11_LeftRight_Equal(t *testing.T) {
	safeTest(t, "Test_Cov11_LeftRight_Equal", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("x", "y")
		var nilLR *corestr.LeftRight
		actual := args.Map{
			"equal":       lr1.IsEqual(lr2),
			"notEqual":    lr1.IsEqual(lr3),
			"bothNil":     nilLR.IsEqual(nil),
			"oneNil":      lr1.IsEqual(nil),
		}
		expected := args.Map{
			"equal": true, "notEqual": false,
			"bothNil": true, "oneNil": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- equal", actual)
	})
}

func Test_Cov11_LeftRight_ClonePtrNonPtr(t *testing.T) {
	safeTest(t, "Test_Cov11_LeftRight_ClonePtrNonPtr", func() {
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()
		nonPtr := lr.NonPtr()
		ptr := lr.Ptr()
		actual := args.Map{
			"clonedLeft": cloned.Left,
			"nonPtrLeft": nonPtr.Left,
			"ptrNotNil":  ptr != nil,
		}
		expected := args.Map{"clonedLeft": "a", "nonPtrLeft": "a", "ptrNotNil": true}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- clone/ptr", actual)
	})
}

func Test_Cov11_LeftRight_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov11_LeftRight_ClearDispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()
		var nilLR *corestr.LeftRight
		nilLR.Clear()
		nilLR.Dispose()
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "", "right": ""}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- clear/dispose", actual)
	})
}

// ═══════════════════════════════════════════
// LeftMiddleRight — string specific
// ═══════════════════════════════════════════

func Test_Cov11_LeftMiddleRight_Constructors(t *testing.T) {
	safeTest(t, "Test_Cov11_LeftMiddleRight_Constructors", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		lmrInv := corestr.InvalidLeftMiddleRight("err")
		lmrInvNM := corestr.InvalidLeftMiddleRightNoMessage()
		actual := args.Map{
			"left":      lmr.Left,
			"mid":       lmr.Middle,
			"right":     lmr.Right,
			"valid":     lmr.IsValid,
			"invValid":  lmrInv.IsValid,
			"invNMVal":  lmrInvNM.IsValid,
		}
		expected := args.Map{
			"left": "l", "mid": "m", "right": "r", "valid": true,
			"invValid": false, "invNMVal": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- constructors", actual)
	})
}

func Test_Cov11_LeftMiddleRight_Methods(t *testing.T) {
	safeTest(t, "Test_Cov11_LeftMiddleRight_Methods", func() {
		lmr := corestr.NewLeftMiddleRight("left", "mid", "right")
		actual := args.Map{
			"leftBytes":  len(lmr.LeftBytes()) > 0,
			"rightBytes": len(lmr.RightBytes()) > 0,
			"midBytes":   len(lmr.MiddleBytes()) > 0,
			"leftTrim":   lmr.LeftTrim(),
			"rightTrim":  lmr.RightTrim(),
			"midTrim":    lmr.MiddleTrim(),
			"isLeftE":    lmr.IsLeftEmpty(),
			"isRightE":   lmr.IsRightEmpty(),
			"isMidE":     lmr.IsMiddleEmpty(),
			"isLeftWS":   lmr.IsLeftWhitespace(),
			"isRightWS":  lmr.IsRightWhitespace(),
			"isMidWS":    lmr.IsMiddleWhitespace(),
			"hasVNEL":    lmr.HasValidNonEmptyLeft(),
			"hasVNER":    lmr.HasValidNonEmptyRight(),
			"hasVNEM":    lmr.HasValidNonEmptyMiddle(),
			"hasVNWSL":   lmr.HasValidNonWhitespaceLeft(),
			"hasVNWSR":   lmr.HasValidNonWhitespaceRight(),
			"hasVNWSM":   lmr.HasValidNonWhitespaceMiddle(),
			"hasSafe":    lmr.HasSafeNonEmpty(),
			"isAll":      lmr.IsAll("left", "mid", "right"),
			"is":         lmr.Is("left", "right"),
		}
		expected := args.Map{
			"leftBytes": true, "rightBytes": true, "midBytes": true,
			"leftTrim": "left", "rightTrim": "right", "midTrim": "mid",
			"isLeftE": false, "isRightE": false, "isMidE": false,
			"isLeftWS": false, "isRightWS": false, "isMidWS": false,
			"hasVNEL": true, "hasVNER": true, "hasVNEM": true,
			"hasVNWSL": true, "hasVNWSR": true, "hasVNWSM": true,
			"hasSafe": true, "isAll": true, "is": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- methods", actual)
	})
}

func Test_Cov11_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_Cov11_LeftMiddleRight_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		cloned := lmr.Clone()
		lr := lmr.ToLeftRight()
		actual := args.Map{
			"clonedLeft": cloned.Left,
			"lrLeft":     lr.Left,
			"lrRight":    lr.Right,
		}
		expected := args.Map{"clonedLeft": "l", "lrLeft": "l", "lrRight": "r"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- clone", actual)
	})
}

func Test_Cov11_LeftMiddleRight_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov11_LeftMiddleRight_ClearDispose", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		lmr.Clear()
		var nilLMR *corestr.LeftMiddleRight
		nilLMR.Clear()
		nilLMR.Dispose()
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle}
		expected := args.Map{"left": "", "mid": ""}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- clear/dispose", actual)
	})
}

// ═══════════════════════════════════════════
// CollectionsOfCollection
// ═══════════════════════════════════════════

func Test_Cov11_CollectionsOfCollection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov11_CollectionsOfCollection_Basic", func() {
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 5)
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c"})
		coc.Add(c1)
		coc.Add(c2)
		actual := args.Map{
			"len":       coc.Length(),
			"isEmpty":   coc.IsEmpty(),
			"hasItems":  coc.HasItems(),
			"allLen":    coc.AllIndividualItemsLength(),
			"itemsLen":  len(coc.Items()),
			"listLen":   len(coc.List(0)),
			"toCollLen": coc.ToCollection().Length(),
		}
		expected := args.Map{
			"len": 2, "isEmpty": false, "hasItems": true,
			"allLen": 3, "itemsLen": 2, "listLen": 3, "toCollLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns correct value -- basic", actual)
	})
}
func Test_Cov11_KeyValueCollection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValueCollection_Basic", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1").Add("k2", "v2")
		actual := args.Map{
			"len":      kvc.Length(),
			"count":    kvc.Count(),
			"isEmpty":  kvc.IsEmpty(),
			"hasAny":   kvc.HasAnyItem(),
			"lastIdx":  kvc.LastIndex(),
			"hasIdx":   kvc.HasIndex(0),
			"hasKey":   kvc.HasKey("k1"),
			"isContains": kvc.IsContains("k1"),
		}
		expected := args.Map{
			"len": 2, "count": 2, "isEmpty": false,
			"hasAny": true, "lastIdx": 1, "hasIdx": true,
			"hasKey": true, "isContains": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- basic", actual)
	})
}

func Test_Cov11_KeyValueCollection_Access(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValueCollection_Access", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1").Add("k2", "v2")
		first := kvc.First()
		firstOD := kvc.FirstOrDefault()
		last := kvc.Last()
		lastOD := kvc.LastOrDefault()
		val, found := kvc.Get("k1")
		_, notFound := kvc.Get("missing")
		actual := args.Map{
			"firstKey":   first.Key,
			"firstODKey": firstOD.Key,
			"lastKey":    last.Key,
			"lastODKey":  lastOD.Key,
			"val":        val,
			"found":      found,
			"notFound":   notFound,
		}
		expected := args.Map{
			"firstKey": "k1", "firstODKey": "k1",
			"lastKey": "k2", "lastODKey": "k2",
			"val": "v1", "found": true, "notFound": false,
		}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- access", actual)
	})
}

func Test_Cov11_KeyValueCollection_EmptyAccess(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValueCollection_EmptyAccess", func() {
		kvc := &corestr.KeyValueCollection{}
		firstOD := kvc.FirstOrDefault()
		lastOD := kvc.LastOrDefault()
		safeVal := kvc.SafeValueAt(0)
		actual := args.Map{
			"firstNil": firstOD == nil,
			"lastNil":  lastOD == nil,
			"safeVal":  safeVal,
		}
		expected := args.Map{"firstNil": true, "lastNil": true, "safeVal": ""}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns empty -- empty access", actual)
	})
}

func Test_Cov11_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValueCollection_Adds", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(true, "k1", "v1")
		kvc.AddIf(false, "k2", "v2")
		kvc.Adds(corestr.KeyValuePair{Key: "k3", Value: "v3"})
		kvc.Adds()
		kvc.AddMap(map[string]string{"k4": "v4"})
		kvc.AddMap(nil)
		kvc.AddHashsetMap(map[string]bool{"k5": true})
		kvc.AddHashsetMap(nil)
		hs := corestr.New.Hashset.Strings([]string{"k6"})
		kvc.AddHashset(hs)
		kvc.AddHashset(nil)
		hm := corestr.New.Hashmap.Cap(1)
		hm.AddOrUpdate("k7", "v7")
		kvc.AddsHashmap(hm)
		kvc.AddsHashmap(nil)
		kvc.AddsHashmaps(hm)
		kvc.AddsHashmaps()
		actual := args.Map{"len": kvc.Length()}
		expected := args.Map{"len": 7}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- adds", actual)
	})
}

func Test_Cov11_KeyValueCollection_Strings(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValueCollection_Strings", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1").Add("k2", "v2")
		strs := kvc.Strings()
		fmtStrs := kvc.StringsUsingFormat("%s=%s")
		str := kvc.String()
		keys := kvc.AllKeys()
		vals := kvc.AllValues()
		sortedKeys := kvc.AllKeysSorted()
		join := kvc.Join(", ")
		joinKeys := kvc.JoinKeys(", ")
		joinVals := kvc.JoinValues(", ")
		actual := args.Map{
			"strsLen":   len(strs),
			"fmtLen":    len(fmtStrs),
			"strNE":     str != "",
			"keysLen":   len(keys),
			"valsLen":   len(vals),
			"sortedLen": len(sortedKeys),
			"joinNE":    join != "",
			"joinKNE":   joinKeys != "",
			"joinVNE":   joinVals != "",
		}
		expected := args.Map{
			"strsLen": 2, "fmtLen": 2, "strNE": true,
			"keysLen": 2, "valsLen": 2, "sortedLen": 2,
			"joinNE": true, "joinKNE": true, "joinVNE": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- strings", actual)
	})
}

func Test_Cov11_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValueCollection_Find", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1").Add("k2", "v2")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "k2", false
		})
		actual := args.Map{"foundLen": len(found)}
		expected := args.Map{"foundLen": 1}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- find", actual)
	})
}

func Test_Cov11_KeyValueCollection_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValueCollection_SafeValueAt", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1")
		safe0 := kvc.SafeValueAt(0)
		safeOOB := kvc.SafeValueAt(99)
		vals := kvc.SafeValuesAtIndexes(0)
		valsEmpty := kvc.SafeValuesAtIndexes()
		actual := args.Map{
			"safe0":    safe0,
			"safeOOB":  safeOOB,
			"valsLen":  len(vals),
			"emptyLen": len(valsEmpty),
		}
		expected := args.Map{"safe0": "v1", "safeOOB": "", "valsLen": 1, "emptyLen": 0}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- safe value at", actual)
	})
}

func Test_Cov11_KeyValueCollection_HashmapMap(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValueCollection_HashmapMap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1")
		hm := kvc.Hashmap()
		m := kvc.Map()
		actual := args.Map{"hmLen": hm.Length(), "mapLen": len(m)}
		expected := args.Map{"hmLen": 1, "mapLen": 1}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- hashmap/map", actual)
	})
}

func Test_Cov11_KeyValueCollection_Json(t *testing.T) {
	safeTest(t, "Test_Cov11_KeyValueCollection_Json", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1")
		j := kvc.Json()
		jp := kvc.JsonPtr()
		model := kvc.JsonModel()
		modelAny := kvc.JsonModelAny()
		bytes, err := kvc.Serialize()
		must := kvc.SerializeMust()
		actual := args.Map{
			"jsonLen": j.Length() > 0, "jpNN": jp != nil,
			"modelLen": len(model), "modelAnyNN": modelAny != nil,
			"bytesLen": len(bytes) > 0, "errNil": err == nil,
			"mustLen": len(must) > 0,
		}
		expected := args.Map{
			"jsonLen": true, "jpNN": true,
			"modelLen": 1, "modelAnyNN": true,
			"bytesLen": true, "errNil": true, "mustLen": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- json", actual)
	})
}
