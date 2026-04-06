package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// LinkedList — comprehensive
// ═══════════════════════════════════════════

func Test_Cov10_LinkedList_Basic(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_Basic", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")
		actual := args.Map{
			"len":      ll.Length(),
			"lenLock":  ll.LengthLock(),
			"isEmpty":  ll.IsEmpty(),
			"hasItems": ll.HasItems(),
			"headNN":   ll.Head() != nil,
			"tailNN":   ll.Tail() != nil,
		}
		expected := args.Map{
			"len": 3, "lenLock": 3, "isEmpty": false, "hasItems": true,
			"headNN": true, "tailNN": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- basic", actual)
	})
}

func Test_Cov10_LinkedList_AddVariations(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_AddVariations", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("a")
		ll.AddNonEmpty("")
		ll.AddNonEmptyWhitespace("b")
		ll.AddNonEmptyWhitespace("  ")
		ll.AddIf(true, "c")
		ll.AddIf(false, "x")
		ll.AddsIf(true, "d", "e")
		ll.AddsIf(false, "y")
		ll.AddFunc(func() string { return "f" })
		ll.Push("g")
		ll.PushBack("h")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 8} // a,b,c,d,e,f,g,h
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- add variations", actual)
	})
}

func Test_Cov10_LinkedList_AddFront(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_AddFront", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.AddFront("a")
		ll.PushFront("z")
		actual := args.Map{"len": ll.Length(), "head": ll.Head().Element}
		expected := args.Map{"len": 3, "head": "z"}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddFront", actual)
	})
}

func Test_Cov10_LinkedList_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_AddLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("a")
		actual := args.Map{"len": ll.Length(), "emptyLock": ll.IsEmptyLock()}
		expected := args.Map{"len": 1, "emptyLock": false}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddLock", actual)
	})
}

func Test_Cov10_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_IsEquals", func() {
		ll1 := corestr.New.LinkedList.Create()
		ll1.Add("a").Add("b")
		ll2 := corestr.New.LinkedList.Create()
		ll2.Add("a").Add("b")
		ll3 := corestr.New.LinkedList.Create()
		ll3.Add("a").Add("c")
		var nilLL *corestr.LinkedList
		actual := args.Map{
			"equal":    ll1.IsEquals(ll2),
			"notEqual": ll1.IsEquals(ll3),
			"nilBoth":  nilLL.IsEquals(nil),
			"nilOne":   nilLL.IsEquals(ll1),
		}
		expected := args.Map{"equal": true, "notEqual": false, "nilBoth": true, "nilOne": false}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- IsEquals", actual)
	})
}

func Test_Cov10_LinkedList_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_InsertAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("c")
		ll.InsertAt(1, "b")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- InsertAt", actual)
	})
}

func Test_Cov10_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_Loop", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{"count": count}
		expected := args.Map{"count": 3}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- Loop", actual)
	})
}

func Test_Cov10_LinkedList_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_Loop_Break", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true // break on first
		})
		actual := args.Map{"count": count}
		expected := args.Map{"count": 1}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- Loop break", actual)
	})
}

func Test_Cov10_LinkedList_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_Loop_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			called = true
			return false
		})
		actual := args.Map{"called": called}
		expected := args.Map{"called": false}
		expected.ShouldBeEqual(t, 0, "LinkedList returns empty -- Loop empty", actual)
	})
}

func Test_Cov10_LinkedList_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov10_LinkedList_AddItemsMap", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})
		ll.AddItemsMap(nil)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddItemsMap", actual)
	})
}

// ═══════════════════════════════════════════
// ValidValue — comprehensive
// ═══════════════════════════════════════════

func Test_Cov10_ValidValue_Constructors(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_Constructors", func() {
		vv := corestr.NewValidValue("hello")
		vvEmpty := corestr.NewValidValueEmpty()
		inv := corestr.InvalidValidValue("err")
		invNo := corestr.InvalidValidValueNoMessage()
		actual := args.Map{
			"vvVal":    vv.Value,
			"vvValid":  vv.IsValid,
			"emptyVal": vvEmpty.Value,
			"invValid": inv.IsValid,
			"invMsg":   inv.Message,
			"invNoMsg": invNo.Message,
		}
		expected := args.Map{
			"vvVal": "hello", "vvValid": true, "emptyVal": "",
			"invValid": false, "invMsg": "err", "invNoMsg": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- constructors", actual)
	})
}

func Test_Cov10_ValidValue_Checks(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_Checks", func() {
		vv := corestr.NewValidValue("hello")
		vvEmpty := corestr.NewValidValue("")
		actual := args.Map{
			"isEmpty":        vv.IsEmpty(),
			"isWS":           vv.IsWhitespace(),
			"hasValidNE":     vv.HasValidNonEmpty(),
			"hasValidNWS":    vv.HasValidNonWhitespace(),
			"hasSafe":        vv.HasSafeNonEmpty(),
			"emptyIsEmpty":   vvEmpty.IsEmpty(),
			"trim":           vv.Trim(),
			"is":             vv.Is("hello"),
			"isNot":          vv.Is("world"),
			"isContains":     vv.IsContains("ell"),
			"isEqInsensitive": vv.IsEqualNonSensitive("HELLO"),
		}
		expected := args.Map{
			"isEmpty": false, "isWS": false, "hasValidNE": true,
			"hasValidNWS": true, "hasSafe": true, "emptyIsEmpty": true,
			"trim": "hello", "is": true, "isNot": false,
			"isContains": true, "isEqInsensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- checks", actual)
	})
}

func Test_Cov10_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_IsAnyOf", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{
			"found":    vv.IsAnyOf("world", "hello"),
			"notFound": vv.IsAnyOf("world", "foo"),
			"empty":    vv.IsAnyOf(),
		}
		expected := args.Map{"found": true, "notFound": false, "empty": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsAnyOf", actual)
	})
}

func Test_Cov10_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_IsAnyContains", func() {
		vv := corestr.NewValidValue("hello world")
		actual := args.Map{
			"found":    vv.IsAnyContains("xyz", "world"),
			"notFound": vv.IsAnyContains("xyz", "abc"),
			"empty":    vv.IsAnyContains(),
		}
		expected := args.Map{"found": true, "notFound": false, "empty": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsAnyContains", actual)
	})
}

func Test_Cov10_ValidValue_TypeConversions(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_TypeConversions", func() {
		vvBool := corestr.NewValidValue("true")
		vvInt := corestr.NewValidValue("42")
		vvFloat := corestr.NewValidValue("3.14")
		vvByte := corestr.NewValidValue("200")
		vvBad := corestr.NewValidValue("abc")
		actual := args.Map{
			"bool":      vvBool.ValueBool(),
			"int":       vvInt.ValueInt(0),
			"defInt":    vvInt.ValueDefInt(),
			"float":     vvFloat.ValueFloat64(0),
			"defFloat":  vvFloat.ValueDefFloat64(),
			"byte":      vvByte.ValueByte(0),
			"defByte":   vvByte.ValueDefByte(),
			"badBool":   vvBad.ValueBool(),
			"badInt":    vvBad.ValueInt(99),
			"emptyBool": corestr.NewValidValue("").ValueBool(),
		}
		expected := args.Map{
			"bool": true, "int": 42, "defInt": 42,
			"float": 3.14, "defFloat": 3.14,
			"byte": byte(200), "defByte": byte(200),
			"badBool": false, "badInt": 99, "emptyBool": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- type conversions", actual)
	})
}
func Test_Cov10_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")
		nonEmpty := vv.SplitNonEmpty(",")
		trimNWS := vv.SplitTrimNonWhitespace(",")
		actual := args.Map{
			"partsLen":  len(parts),
			"neLen":     len(nonEmpty),
			"trimLen":   len(trimNWS),
		}
		expected := args.Map{"partsLen": 3, "neLen": 3, "trimLen": 3}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Split", actual)
	})
}

func Test_Cov10_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_Clone", func() {
		vv := corestr.NewValidValue("hello")
		cloned := vv.Clone()
		var nilVV *corestr.ValidValue
		actual := args.Map{
			"cloneVal": cloned.Value,
			"nilClone": nilVV.Clone() == nil,
		}
		expected := args.Map{"cloneVal": "hello", "nilClone": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Clone", actual)
	})
}

func Test_Cov10_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_String", func() {
		vv := corestr.NewValidValue("hello")
		var nilVV *corestr.ValidValue
		actual := args.Map{
			"str":      vv.String(),
			"fullStr":  vv.FullString() != "",
			"nilStr":   nilVV.String(),
			"nilFull":  nilVV.FullString(),
		}
		expected := args.Map{"str": "hello", "fullStr": true, "nilStr": "", "nilFull": ""}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- String", actual)
	})
}

func Test_Cov10_ValidValue_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_ClearDispose", func() {
		vv := corestr.NewValidValue("hello")
		vv.Clear()
		var nilVV *corestr.ValidValue
		nilVV.Clear()   // should not panic
		nilVV.Dispose() // should not panic
		actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
		expected := args.Map{"val": "", "valid": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Clear/Dispose", actual)
	})
}

func Test_Cov10_ValidValue_JSON(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_JSON", func() {
		vv := corestr.NewValidValue("hello")
		j := vv.Json()
		jp := vv.JsonPtr()
		b, err := vv.Serialize()
		actual := args.Map{
			"jHas": j.HasBytes(), "jpNN": jp != nil,
			"bLen": len(b) > 0, "noErr": err == nil,
		}
		expected := args.Map{"jHas": true, "jpNN": true, "bLen": true, "noErr": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- JSON", actual)
	})
}

func Test_Cov10_ValidValue_Regex(t *testing.T) {
	safeTest(t, "Test_Cov10_ValidValue_Regex", func() {
		vv := corestr.NewValidValue("hello123")
		actual := args.Map{
			"nilRegex": vv.IsRegexMatches(nil),
			"nilFind":  vv.RegexFindString(nil),
			"nilAll":   len(vv.RegexFindAllStrings(nil, -1)),
		}
		expected := args.Map{"nilRegex": false, "nilFind": "", "nilAll": 0}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- Regex nil", actual)
	})
}

// ═══════════════════════════════════════════
// KeyValuePair — comprehensive
// ═══════════════════════════════════════════

func Test_Cov10_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_Cov10_KeyValuePair_Basic", func() {
		kv := &corestr.KeyValuePair{Key: "name", Value: "alice"}
		actual := args.Map{
			"keyName":   kv.KeyName(),
			"varName":   kv.VariableName(),
			"valStr":    kv.ValueString(),
			"isVarEq":   kv.IsVariableNameEqual("name"),
			"isValEq":   kv.IsValueEqual("alice"),
			"hasKey":    kv.HasKey(),
			"hasVal":    kv.HasValue(),
			"isKeyEmpty": kv.IsKeyEmpty(),
			"isValEmpty": kv.IsValueEmpty(),
			"isKVEmpty":  kv.IsKeyValueEmpty(),
			"isKVAnyE":   kv.IsKeyValueAnyEmpty(),
			"isKey":     kv.IsKey("name"),
			"isVal":     kv.IsVal("alice"),
			"is":        kv.Is("name", "alice"),
			"compile":   kv.Compile() != "",
			"str":       kv.String() != "",
			"trimKey":   kv.TrimKey(),
			"trimVal":   kv.TrimValue(),
		}
		expected := args.Map{
			"keyName": "name", "varName": "name", "valStr": "alice",
			"isVarEq": true, "isValEq": true, "hasKey": true, "hasVal": true,
			"isKeyEmpty": false, "isValEmpty": false, "isKVEmpty": false,
			"isKVAnyE": false, "isKey": true, "isVal": true, "is": true,
			"compile": true, "str": true, "trimKey": "name", "trimVal": "alice",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- basic", actual)
	})
}

func Test_Cov10_KeyValuePair_TypeConversions(t *testing.T) {
	safeTest(t, "Test_Cov10_KeyValuePair_TypeConversions", func() {
		kvBool := &corestr.KeyValuePair{Key: "k", Value: "true"}
		kvInt := &corestr.KeyValuePair{Key: "k", Value: "42"}
		kvFloat := &corestr.KeyValuePair{Key: "k", Value: "3.14"}
		kvByte := &corestr.KeyValuePair{Key: "k", Value: "100"}
		kvBad := &corestr.KeyValuePair{Key: "k", Value: "abc"}
		actual := args.Map{
			"bool":      kvBool.ValueBool(),
			"int":       kvInt.ValueInt(0),
			"defInt":    kvInt.ValueDefInt(),
			"float":     kvFloat.ValueFloat64(0),
			"defFloat":  kvFloat.ValueDefFloat64(),
			"byte":      kvByte.ValueByte(0),
			"defByte":   kvByte.ValueDefByte(),
			"badBool":   kvBad.ValueBool(),
			"badInt":    kvBad.ValueInt(99),
			"emptyBool": (&corestr.KeyValuePair{Key: "k", Value: ""}).ValueBool(),
		}
		expected := args.Map{
			"bool": true, "int": 42, "defInt": 42,
			"float": 3.14, "defFloat": 3.14,
			"byte": byte(100), "defByte": byte(100),
			"badBool": false, "badInt": 99, "emptyBool": false,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- type conversions", actual)
	})
}

func Test_Cov10_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_Cov10_KeyValuePair_ValueValid", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()
		vvo := kv.ValueValidOptions(false, "msg")
		actual := args.Map{
			"vvVal":     vv.Value,
			"vvValid":   vv.IsValid,
			"vvoValid":  vvo.IsValid,
			"vvoMsg":    vvo.Message,
		}
		expected := args.Map{
			"vvVal": "v", "vvValid": true, "vvoValid": false, "vvoMsg": "msg",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns non-empty -- ValueValid", actual)
	})
}

func Test_Cov10_KeyValuePair_FormatString(t *testing.T) {
	safeTest(t, "Test_Cov10_KeyValuePair_FormatString", func() {
		kv := &corestr.KeyValuePair{Key: "name", Value: "alice"}
		formatted := kv.FormatString("%s=%s")
		actual := args.Map{"formatted": formatted}
		expected := args.Map{"formatted": "name=alice"}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- FormatString", actual)
	})
}

func Test_Cov10_KeyValuePair_JSON(t *testing.T) {
	safeTest(t, "Test_Cov10_KeyValuePair_JSON", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		j := kv.Json()
		jp := kv.JsonPtr()
		b, err := kv.Serialize()
		actual := args.Map{
			"jHas": j.HasBytes(), "jpNN": jp != nil,
			"bLen": len(b) > 0, "noErr": err == nil,
		}
		expected := args.Map{"jHas": true, "jpNN": true, "bLen": true, "noErr": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- JSON", actual)
	})
}

func Test_Cov10_KeyValuePair_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov10_KeyValuePair_ClearDispose", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		var nilKV *corestr.KeyValuePair
		nilKV.Clear()   // should not panic
		nilKV.Dispose() // should not panic
		actual := args.Map{"key": kv.Key, "val": kv.Value}
		expected := args.Map{"key": "", "val": ""}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Clear/Dispose", actual)
	})
}

func Test_Cov10_KeyValuePair_NilChecks(t *testing.T) {
	safeTest(t, "Test_Cov10_KeyValuePair_NilChecks", func() {
		var nilKV *corestr.KeyValuePair
		actual := args.Map{
			"nilIsKVAnyEmpty": nilKV.IsKeyValueAnyEmpty(),
		}
		expected := args.Map{"nilIsKVAnyEmpty": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns nil -- nil checks", actual)
	})
}

// ═══════════════════════════════════════════
// LeftMiddleRight
// ═══════════════════════════════════════════

func Test_Cov10_LeftMiddleRight(t *testing.T) {
	safeTest(t, "Test_Cov10_LeftMiddleRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{
			"isAll": lmr.IsAll("a", "b", "c"),
		}
		expected := args.Map{"isAll": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- IsAll", actual)
	})
}

// ═══════════════════════════════════════════
// LeftRight (corestr)
// ═══════════════════════════════════════════

func Test_Cov10_LeftRight(t *testing.T) {
	safeTest(t, "Test_Cov10_LeftRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "b"}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- with args", actual)
	})
}

// ═══════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════

func Test_Cov10_ValueStatus(t *testing.T) {
	safeTest(t, "Test_Cov10_ValueStatus", func() {
		vs := corestr.ValueStatus{
			ValueValid: &corestr.ValidValue{Value: "hello"},
			Index:      0,
		}
		actual := args.Map{"val": vs.ValueValid.Value, "idx": vs.Index}
		expected := args.Map{"val": "hello", "idx": 0}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- with args", actual)
	})
}

// ═══════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════

func Test_Cov10_TextWithLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov10_TextWithLineNumber", func() {
		tln := corestr.TextWithLineNumber{Text: "hello", LineNumber: 1}
		actual := args.Map{"text": tln.Text, "lineNum": tln.LineNumber}
		expected := args.Map{"text": "hello", "lineNum": 1}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns non-empty -- with args", actual)
	})
}

// ═══════════════════════════════════════════
// SimpleSlice — additional methods
// ═══════════════════════════════════════════

func Test_Cov10_SimpleSlice_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov10_SimpleSlice_HasIndex", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.Add("a").Add("b")
		actual := args.Map{
			"hasIdx0":  ss.HasIndex(0),
			"hasIdx1":  ss.HasIndex(1),
			"hasIdx5":  ss.HasIndex(5),
		}
		expected := args.Map{"hasIdx0": true, "hasIdx1": true, "hasIdx5": false}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- HasIndex", actual)
	})
}

func Test_Cov10_SimpleSlice_FirstLastOrDefault(t *testing.T) {
	safeTest(t, "Test_Cov10_SimpleSlice_FirstLastOrDefault", func() {
		ss := corestr.New.SimpleSlice.Cap(3)
		ss.Add("a").Add("b").Add("c")
		empty := corestr.New.SimpleSlice.Cap(0)
		actual := args.Map{
			"first":      ss.FirstOrDefault(),
			"last":       ss.LastOrDefault(),
			"emptyFirst": empty.FirstOrDefault(),
			"emptyLast":  empty.LastOrDefault(),
		}
		expected := args.Map{
			"first": "a", "last": "c", "emptyFirst": "", "emptyLast": "",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- FirstLastOrDefault", actual)
	})
}

// ═══════════════════════════════════════════
// Hashset — additional methods
// ═══════════════════════════════════════════

func Test_Cov10_Hashset_SortedList(t *testing.T) {
	safeTest(t, "Test_Cov10_Hashset_SortedList", func() {
		h := corestr.New.Hashset.Cap(5)
		h.Adds("c", "a", "b")
		sorted := h.SortedList()
		actual := args.Map{"first": sorted[0], "last": sorted[2]}
		expected := args.Map{"first": "a", "last": "c"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- SortedList", actual)
	})
}

func Test_Cov10_Hashset_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov10_Hashset_IsEqual", func() {
		h1 := corestr.New.Hashset.Cap(3)
		h1.Adds("a", "b")
		h2 := corestr.New.Hashset.Cap(3)
		h2.Adds("b", "a")
		h3 := corestr.New.Hashset.Cap(3)
		h3.Adds("a", "c")
		actual := args.Map{
			"equal":    h1.IsEqual(h2),
			"notEqual": h1.IsEqual(h3),
		}
		expected := args.Map{"equal": true, "notEqual": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEqual", actual)
	})
}

// ═══════════════════════════════════════════
// Hashmap — additional methods
// ═══════════════════════════════════════════

func Test_Cov10_Hashmap_Keys(t *testing.T) {
	safeTest(t, "Test_Cov10_Hashmap_Keys", func() {
		h := corestr.New.Hashmap.Cap(3)
		h.Set("b", "2")
		h.Set("a", "1")
		keys := h.Keys()
		actual := args.Map{"count": len(keys)}
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Keys", actual)
	})
}

func Test_Cov10_Hashmap_GetValue(t *testing.T) {
	safeTest(t, "Test_Cov10_Hashmap_GetValue", func() {
		h := corestr.New.Hashmap.Cap(3)
		h.Set("k", "v")
		v, found := h.GetValue("k")
		_, notFound := h.GetValue("x")
		actual := args.Map{"v": v, "found": found, "notFound": notFound}
		expected := args.Map{"v": "v", "found": true, "notFound": false}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetValue", actual)
	})
}

// ═══════════════════════════════════════════
// Collection — Join (comma)
// ═══════════════════════════════════════════

func Test_Cov10_Collection_JoinComma(t *testing.T) {
	safeTest(t, "Test_Cov10_Collection_JoinComma", func() {
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c")
		csv := c.Join(", ")
		actual := args.Map{"csv": csv}
		expected := args.Map{"csv": "a, b, c"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Join comma", actual)
	})
}

// ═══════════════════════════════════════════
// CharHashsetMap
// ═══════════════════════════════════════════

func Test_Cov10_CharHashsetMap(t *testing.T) {
	safeTest(t, "Test_Cov10_CharHashsetMap", func() {
		chm := corestr.New.CharHashsetMap.Cap(3, 5)
		chm.Add("hello")
		chm.Add("help")
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": actual["len"]}
		expected.ShouldBeEqual(t, 0, "CharHashsetMap returns correct value -- with args", actual)
	})
}
