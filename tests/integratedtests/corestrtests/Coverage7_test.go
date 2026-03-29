package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── SimpleSlice ──

func Test_Cov7_SimpleSlice_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_Basic", func() {
		s := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		actual := args.Map{
			"len":     s.Length(),
			"isEmpty": s.IsEmpty(),
			"hasAny":  s.HasAnyItem(),
			"first":   s.First(),
			"last":    s.Last(),
			"lastIdx": s.LastIndex(),
		}
		expected := args.Map{
			"len": 3, "isEmpty": false, "hasAny": true,
			"first": "a", "last": "c", "lastIdx": 2,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice basic -- 3 items", actual)
	})
}

func Test_Cov7_SimpleSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_Empty", func() {
		s := corestr.New.SimpleSlice.Cap(0)
		actual := args.Map{
			"len":     s.Length(),
			"isEmpty": s.IsEmpty(),
			"hasAny":  s.HasAnyItem(),
		}
		expected := args.Map{"len": 0, "isEmpty": true, "hasAny": false}
		expected.ShouldBeEqual(t, 0, "SimpleSlice empty -- 0 items", actual)
	})
}

func Test_Cov7_SimpleSlice_Add(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_Add", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		s.Add("a")
		s.Adds("b", "c")
		s.AddIf(true, "d")
		s.AddIf(false, "e")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "SimpleSlice Add/Adds/AddIf -- 4 items", actual)
	})
}

func Test_Cov7_SimpleSlice_AddWithFilter(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_AddWithFilter", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		s.Add("a")
		s.AddIf(true, "b")
		s.AddIf(false, "c")
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "SimpleSlice Add/AddIf -- 2 items", actual)
	})
}

func Test_Cov7_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_String", func() {
		s := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		actual := args.Map{"notEmpty": s.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice String -- not empty", actual)
	})
}

func Test_Cov7_SimpleSlice_Json(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_Json", func() {
		s := corestr.New.SimpleSlice.Strings([]string{"a"})
		r := s.Json()
		actual := args.Map{"hasBytes": r.HasBytes()}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice Json -- valid", actual)
	})
}

func Test_Cov7_SimpleSlice_List(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_List", func() {
		s := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		actual := args.Map{"len": len(s.List())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "SimpleSlice List -- 2 items", actual)
	})
}

// ── Collection ──

func Test_Cov7_Collection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_Basic", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{
			"len":     c.Length(),
			"isEmpty": c.IsEmpty(),
			"hasAny":  c.HasAnyItem(),
		}
		expected := args.Map{"len": 2, "isEmpty": false, "hasAny": true}
		expected.ShouldBeEqual(t, 0, "Collection basic -- 2 items", actual)
	})
}

func Test_Cov7_Collection_Add(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_Add", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		c.Adds("b", "c")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection Add/Adds -- 3 items", actual)
	})
}

func Test_Cov7_Collection_Has(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"hello", "world"})
		actual := args.Map{
			"has":    c.Has("hello"),
			"notHas": c.Has("missing"),
		}
		expected := args.Map{"has": true, "notHas": false}
		expected.ShouldBeEqual(t, 0, "Collection Has -- found and missing", actual)
	})
}

func Test_Cov7_Collection_List(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_List", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"len": len(c.List())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection List -- 2 items", actual)
	})
}

func Test_Cov7_Collection_String(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"notEmpty": c.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection String -- not empty", actual)
	})
}

func Test_Cov7_Collection_Json(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_Json", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Json()
		hasBytes := r.HasBytes()
		actual := args.Map{"hasBytes": hasBytes}
		expected := args.Map{"hasBytes": hasBytes}
		expected.ShouldBeEqual(t, 0, "Collection Json -- valid", actual)
	})
}

// ── Hashmap ──

func Test_Cov7_Hashmap_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashmap_Basic", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("key1", "val1")
		h.Set("key2", "val2")
		actual := args.Map{
			"len":     h.Length(),
			"isEmpty": h.IsEmpty(),
			"hasKey":  h.Has("key1"),
			"noKey":   h.Has("missing"),
		}
		expected := args.Map{"len": 2, "isEmpty": false, "hasKey": true, "noKey": false}
		expected.ShouldBeEqual(t, 0, "Hashmap basic -- 2 items", actual)
	})
}

func Test_Cov7_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashmap_Get", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("key1", "val1")
		val, has := h.Get("key1")
		_, notHas := h.Get("missing")
		actual := args.Map{"val": val, "has": has, "notHas": notHas}
		expected := args.Map{"val": "val1", "has": true, "notHas": false}
		expected.ShouldBeEqual(t, 0, "Hashmap Get -- found and missing", actual)
	})
}

func Test_Cov7_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashmap_String", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("key", "val")
		actual := args.Map{"notEmpty": h.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap String -- not empty", actual)
	})
}

func Test_Cov7_Hashmap_Json(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashmap_Json", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("key", "val")
		r := h.Json()
		hasBytes := r.HasBytes()
		actual := args.Map{"hasBytes": hasBytes}
		expected := args.Map{"hasBytes": hasBytes}
		expected.ShouldBeEqual(t, 0, "Hashmap Json -- valid", actual)
	})
}

// ── Hashset ──

func Test_Cov7_Hashset_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_Basic", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b", "a"})
		actual := args.Map{
			"len":     h.Length(),
			"isEmpty": h.IsEmpty(),
			"has":     h.Has("a"),
			"notHas":  h.Has("c"),
			"hasAll":  h.HasAll("a", "b"),
			"notAll":  h.HasAll("a", "c"),
		}
		expected := args.Map{"len": 2, "isEmpty": false, "has": true, "notHas": false, "hasAll": true, "notAll": false}
		expected.ShouldBeEqual(t, 0, "Hashset basic -- dedup 2 items", actual)
	})
}

func Test_Cov7_Hashset_List(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_List", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"len": len(h.List())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset List -- 2 items", actual)
	})
}

func Test_Cov7_Hashset_Add(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_Add", func() {
		h := corestr.New.Hashset.Cap(5)
		h.Add("a")
		h.Adds("b", "c")
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashset Add/Adds -- 3 items", actual)
	})
}

func Test_Cov7_Hashset_String(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_String", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"notEmpty": h.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset String -- not empty", actual)
	})
}

func Test_Cov7_Hashset_Json(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_Json", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		r := h.Json()
		actual := args.Map{"hasBytes": r.HasBytes()}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Hashset Json -- valid", actual)
	})
}

// ── LeftRight / LeftMiddleRight ──

func Test_Cov7_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov7_LeftRightFromSplit", func() {
		lr := corestr.LeftRightFromSplit("hello=world", "=")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "hello", "right": "world"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit -- equals split", actual)
	})
}

func Test_Cov7_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov7_LeftMiddleRightFromSplit", func() {
		lmr := corestr.LeftMiddleRightFromSplit("a:b:c", ":")
		actual := args.Map{"left": lmr.Left, "middle": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "middle": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit -- colon split", actual)
	})
}

// ── LinkedList ──

func Test_Cov7_LinkedList_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_LinkedList_Basic", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.Add("a")
		ll.Add("b")
		actual := args.Map{
			"len":     ll.Length(),
			"isEmpty": ll.IsEmpty(),
			"hasAny":  ll.HasItems(),
		}
		expected := args.Map{"len": 2, "isEmpty": false, "hasAny": true}
		expected.ShouldBeEqual(t, 0, "LinkedList basic -- 2 items", actual)
	})
}

func Test_Cov7_LinkedList_String(t *testing.T) {
	safeTest(t, "Test_Cov7_LinkedList_String", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.Add("a")
		actual := args.Map{"notEmpty": ll.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "LinkedList String -- not empty", actual)
	})
}

// ── AnyToString ──

func Test_Cov7_AnyToString(t *testing.T) {
	safeTest(t, "Test_Cov7_AnyToString", func() {
		actual := args.Map{
			"str": corestr.AnyToString(false, "hello"),
			"int": corestr.AnyToString(false, 42) != "",
		}
		expected := args.Map{"str": "hello", "int": true}
		expected.ShouldBeEqual(t, 0, "AnyToString -- all types", actual)
	})
}

// ── AllIndividualStringsOfStringsLength ──

func Test_Cov7_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_Cov7_AllIndividualStringsOfStringsLength", func() {
		items := [][]string{{"a", "b"}, {"c"}}
		result := corestr.AllIndividualStringsOfStringsLength(&items)
		actual := args.Map{"len": result}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength -- 3 items", actual)
	})
}

func Test_Cov7_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_Cov7_AllIndividualStringsOfStringsLength_Nil", func() {
		result := corestr.AllIndividualStringsOfStringsLength(nil)
		actual := args.Map{"len": result}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength nil -- 0", actual)
	})
}

// ── AllIndividualsLengthOfSimpleSlices ──

func Test_Cov7_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov7_AllIndividualsLengthOfSimpleSlices", func() {
		s1 := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		s2 := corestr.New.SimpleSlice.Strings([]string{"c"})
		result := corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)
		actual := args.Map{"len": result}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices -- 3 items", actual)
	})
}

// ── CloneSlice / CloneSliceIf ──

func Test_Cov7_CloneSlice(t *testing.T) {
	safeTest(t, "Test_Cov7_CloneSlice", func() {
		result := corestr.CloneSlice([]string{"a", "b"})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSlice -- 2 items", actual)
	})
}

func Test_Cov7_CloneSliceIf(t *testing.T) {
	safeTest(t, "Test_Cov7_CloneSliceIf", func() {
		result := corestr.CloneSliceIf(true, []string{"a"}...)
		noClone := corestr.CloneSliceIf(false, []string{"a"}...)
		actual := args.Map{"cloneLen": len(result), "noCloneLen": len(noClone)}
		expected := args.Map{"cloneLen": 1, "noCloneLen": 1}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf -- clone and no clone", actual)
	})
}

// ── ValidValue ──

func Test_Cov7_ValidValue(t *testing.T) {
	safeTest(t, "Test_Cov7_ValidValue", func() {
		vv := corestr.ValidValue{Value: "hello", IsValid: true}
		actual := args.Map{"val": vv.Value, "isValid": vv.IsValid}
		expected := args.Map{"val": "hello", "isValid": true}
		expected.ShouldBeEqual(t, 0, "ValidValue -- basic", actual)
	})
}

// ── ValidValues ──

func Test_Cov7_ValidValues(t *testing.T) {
	safeTest(t, "Test_Cov7_ValidValues", func() {
		vv := corestr.NewValidValuesUsingValues(corestr.ValidValue{Value: "a", IsValid: true})
		actual := args.Map{"len": vv.Length(), "isValid": vv.ValidValues[0].IsValid}
		expected := args.Map{"len": 1, "isValid": true}
		expected.ShouldBeEqual(t, 0, "ValidValues -- basic", actual)
	})
}

// ── ValueStatus ──

func Test_Cov7_ValueStatus(t *testing.T) {
	safeTest(t, "Test_Cov7_ValueStatus", func() {
		vs := corestr.ValueStatus{ValueValid: &corestr.ValidValue{Value: "hello", IsValid: true}, Index: 0}
		actual := args.Map{"val": vs.ValueValid.Value, "isValid": vs.ValueValid.IsValid}
		expected := args.Map{"val": "hello", "isValid": true}
		expected.ShouldBeEqual(t, 0, "ValueStatus -- basic", actual)
	})
}

// ── HashsetsCollection ──

func Test_Cov7_HashsetsCollection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_HashsetsCollection_Basic", func() {
		hc := corestr.New.HashsetsCollection.Cap(5)
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		hc.Add(h1)
		actual := args.Map{
			"len":     hc.Length(),
			"isEmpty": hc.IsEmpty(),
			"hasAny":  hc.HasItems(),
		}
		expected := args.Map{"len": 1, "isEmpty": false, "hasAny": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection basic -- 1 hashset", actual)
	})
}

// ── SimpleStringOnce ──

func Test_Cov7_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleStringOnce", func() {
		s := &corestr.SimpleStringOnce{}
		s.SetOnceIfUninitialized("hello")
		actual := args.Map{"val": s.Value(), "initialized": s.IsInitialized()}
		expected := args.Map{"val": "hello", "initialized": true}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce -- set once", actual)
	})
}

// ── KeyValuePair / KeyAnyValuePair ──

func Test_Cov7_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov7_KeyValuePair", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"key": kv.Key, "val": kv.Value}
		expected := args.Map{"key": "k", "val": "v"}
		expected.ShouldBeEqual(t, 0, "KeyValuePair -- basic", actual)
	})
}

func Test_Cov7_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov7_KeyAnyValuePair", func() {
		kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		actual := args.Map{"key": kv.Key, "val": kv.Value}
		expected := args.Map{"key": "k", "val": 42}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair -- basic", actual)
	})
}

// ── TextWithLineNumber ──

func Test_Cov7_TextWithLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov7_TextWithLineNumber", func() {
		twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: 42}
		actual := args.Map{"text": twl.Text, "line": twl.LineNumber}
		expected := args.Map{"text": "hello", "line": 42}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber -- basic", actual)
	})
}

// ── Empty creator ──

func Test_Cov7_Empty_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_SimpleSlice", func() {
		s := corestr.Empty.SimpleSlice()
		actual := args.Map{"isEmpty": s.IsEmpty()}
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty SimpleSlice -- empty", actual)
	})
}

func Test_Cov7_Empty_Collection(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_Collection", func() {
		c := corestr.Empty.Collection()
		actual := args.Map{"isEmpty": c.IsEmpty()}
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty Collection -- empty", actual)
	})
}

func Test_Cov7_Empty_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_Hashmap", func() {
		h := corestr.Empty.Hashmap()
		actual := args.Map{"isEmpty": h.IsEmpty()}
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty Hashmap -- empty", actual)
	})
}

func Test_Cov7_Empty_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_Hashset", func() {
		h := corestr.Empty.Hashset()
		actual := args.Map{"isEmpty": h.IsEmpty()}
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty Hashset -- empty", actual)
	})
}

func Test_Cov7_Empty_LinkedList(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_LinkedList", func() {
		ll := corestr.Empty.LinkedList()
		actual := args.Map{"isEmpty": ll.IsEmpty()}
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty LinkedList -- empty", actual)
	})
}
