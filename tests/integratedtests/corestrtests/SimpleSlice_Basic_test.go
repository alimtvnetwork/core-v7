package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── SimpleSlice ──

func Test_SimpleSlice_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_Basic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"len":     s.Length(),
			"isEmpty": s.IsEmpty(),
			"hasAny":  s.HasAnyItem(),
			"first":   s.First(),
			"last":    s.Last(),
			"lastIdx": s.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"len": 3, "isEmpty": false, "hasAny": true,
			"first": "a", "last": "c", "lastIdx": 2,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice basic -- 3 items", actual)
	})
}

func Test_SimpleSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(0)

		// Act
		actual := args.Map{
			"len":     s.Length(),
			"isEmpty": s.IsEmpty(),
			"hasAny":  s.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"isEmpty": true,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice empty -- 0 items", actual)
	})
}

func Test_SimpleSlice_Add(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_Add", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Add("a")
		s.Adds("b", "c")
		s.AddIf(true, "d")
		s.AddIf(false, "e")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "SimpleSlice Add/Adds/AddIf -- 4 items", actual)
	})
}

func Test_SimpleSlice_AddWithFilter(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_AddWithFilter", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Add("a")
		s.AddIf(true, "b")
		s.AddIf(false, "c")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "SimpleSlice Add/AddIf -- 2 items", actual)
	})
}

func Test_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_String", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"notEmpty": s.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice String -- not empty", actual)
	})
}

func Test_SimpleSlice_Json(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_Json", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Strings([]string{"a"})
		r := s.Json()

		// Act
		actual := args.Map{"hasBytes": r.HasBytes()}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice Json -- valid", actual)
	})
}

func Test_SimpleSlice_List(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleSlice_List", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": len(s.List())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "SimpleSlice List -- 2 items", actual)
	})
}

// ── Collection ──

func Test_Collection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_Basic", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"len":     c.Length(),
			"isEmpty": c.IsEmpty(),
			"hasAny":  c.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"isEmpty": false,
			"hasAny": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection basic -- 2 items", actual)
	})
}

func Test_Collection_Add(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_Add", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		c.Adds("b", "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection Add/Adds -- 3 items", actual)
	})
}

func Test_Collection_Has(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_Has", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"hello", "world"})

		// Act
		actual := args.Map{
			"has":    c.Has("hello"),
			"notHas": c.Has("missing"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"notHas": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection Has -- found and missing", actual)
	})
}

func Test_Collection_List(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_List", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": len(c.List())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection List -- 2 items", actual)
	})
}

func Test_Collection_String(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_String", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": c.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection String -- not empty", actual)
	})
}

func Test_Collection_Json(t *testing.T) {
	safeTest(t, "Test_Cov7_Collection_Json", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Json()
		hasBytes := r.HasBytes()

		// Act
		actual := args.Map{"hasBytes": hasBytes}

		// Assert
		expected := args.Map{"hasBytes": hasBytes}
		expected.ShouldBeEqual(t, 0, "Collection Json -- valid", actual)
	})
}

// ── Hashmap ──

func Test_Hashmap_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashmap_Basic", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("key1", "val1")
		h.Set("key2", "val2")

		// Act
		actual := args.Map{
			"len":     h.Length(),
			"isEmpty": h.IsEmpty(),
			"hasKey":  h.Has("key1"),
			"noKey":   h.Has("missing"),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"isEmpty": false,
			"hasKey": true,
			"noKey": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap basic -- 2 items", actual)
	})
}

func Test_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashmap_Get", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("key1", "val1")
		val, has := h.Get("key1")
		_, notHas := h.Get("missing")

		// Act
		actual := args.Map{
			"val": val,
			"has": has,
			"notHas": notHas,
		}

		// Assert
		expected := args.Map{
			"val": "val1",
			"has": true,
			"notHas": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap Get -- found and missing", actual)
	})
}

func Test_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashmap_String", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("key", "val")

		// Act
		actual := args.Map{"notEmpty": h.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap String -- not empty", actual)
	})
}

func Test_Hashmap_Json(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashmap_Json", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("key", "val")
		r := h.Json()
		hasBytes := r.HasBytes()

		// Act
		actual := args.Map{"hasBytes": hasBytes}

		// Assert
		expected := args.Map{"hasBytes": hasBytes}
		expected.ShouldBeEqual(t, 0, "Hashmap Json -- valid", actual)
	})
}

// ── Hashset ──

func Test_Hashset_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_Basic", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b", "a"})

		// Act
		actual := args.Map{
			"len":     h.Length(),
			"isEmpty": h.IsEmpty(),
			"has":     h.Has("a"),
			"notHas":  h.Has("c"),
			"hasAll":  h.HasAll("a", "b"),
			"notAll":  h.HasAll("a", "c"),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"isEmpty": false,
			"has": true,
			"notHas": false,
			"hasAll": true,
			"notAll": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset basic -- dedup 2 items", actual)
	})
}

func Test_Hashset_List(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_List", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": len(h.List())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset List -- 2 items", actual)
	})
}

func Test_Hashset_Add(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_Add", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Add("a")
		h.Adds("b", "c")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashset Add/Adds -- 3 items", actual)
	})
}

func Test_Hashset_String(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_String", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": h.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset String -- not empty", actual)
	})
}

func Test_Hashset_Json(t *testing.T) {
	safeTest(t, "Test_Cov7_Hashset_Json", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		r := h.Json()

		// Act
		actual := args.Map{"hasBytes": r.HasBytes()}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Hashset Json -- valid", actual)
	})
}

// ── LeftRight / LeftMiddleRight ──

func Test_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov7_LeftRightFromSplit", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("hello=world", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "hello",
			"right": "world",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit -- equals split", actual)
	})
}

func Test_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov7_LeftMiddleRightFromSplit", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a:b:c", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit -- colon split", actual)
	})
}

// ── LinkedList ──

func Test_LinkedList_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_LinkedList_Basic", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()
		ll.Add("a")
		ll.Add("b")

		// Act
		actual := args.Map{
			"len":     ll.Length(),
			"isEmpty": ll.IsEmpty(),
			"hasAny":  ll.HasItems(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"isEmpty": false,
			"hasAny": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList basic -- 2 items", actual)
	})
}

func Test_LinkedList_String(t *testing.T) {
	safeTest(t, "Test_Cov7_LinkedList_String", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()
		ll.Add("a")

		// Act
		actual := args.Map{"notEmpty": ll.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "LinkedList String -- not empty", actual)
	})
}

// ── AnyToString ──

func Test_AnyToString(t *testing.T) {
	safeTest(t, "Test_Cov7_AnyToString", func() {
		// Act
		actual := args.Map{
			"str": corestr.AnyToString(false, "hello"),
			"int": corestr.AnyToString(false, 42) != "",
		}

		// Assert
		expected := args.Map{
			"str": "hello",
			"int": true,
		}
		expected.ShouldBeEqual(t, 0, "AnyToString -- all types", actual)
	})
}

// ── AllIndividualStringsOfStringsLength ──

func Test_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_Cov7_AllIndividualStringsOfStringsLength", func() {
		// Arrange
		items := [][]string{{"a", "b"}, {"c"}}
		result := corestr.AllIndividualStringsOfStringsLength(&items)

		// Act
		actual := args.Map{"len": result}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength -- 3 items", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_Cov7_AllIndividualStringsOfStringsLength_Nil", func() {
		// Arrange
		result := corestr.AllIndividualStringsOfStringsLength(nil)

		// Act
		actual := args.Map{"len": result}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength nil -- 0", actual)
	})
}

// ── AllIndividualsLengthOfSimpleSlices ──

func Test_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov7_AllIndividualsLengthOfSimpleSlices", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		s2 := corestr.New.SimpleSlice.Strings([]string{"c"})
		result := corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)

		// Act
		actual := args.Map{"len": result}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices -- 3 items", actual)
	})
}

// ── CloneSlice / CloneSliceIf ──

func Test_CloneSlice(t *testing.T) {
	safeTest(t, "Test_Cov7_CloneSlice", func() {
		// Arrange
		result := corestr.CloneSlice([]string{"a", "b"})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSlice -- 2 items", actual)
	})
}

func Test_CloneSliceIf(t *testing.T) {
	safeTest(t, "Test_Cov7_CloneSliceIf", func() {
		// Arrange
		result := corestr.CloneSliceIf(true, []string{"a"}...)
		noClone := corestr.CloneSliceIf(false, []string{"a"}...)

		// Act
		actual := args.Map{
			"cloneLen": len(result),
			"noCloneLen": len(noClone),
		}

		// Assert
		expected := args.Map{
			"cloneLen": 1,
			"noCloneLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf -- clone and no clone", actual)
	})
}

// ── ValidValue ──

func Test_ValidValue(t *testing.T) {
	safeTest(t, "Test_Cov7_ValidValue", func() {
		// Arrange
		vv := corestr.ValidValue{Value: "hello", IsValid: true}

		// Act
		actual := args.Map{
			"val": vv.Value,
			"isValid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue -- basic", actual)
	})
}

// ── ValidValues ──

func Test_ValidValues(t *testing.T) {
	safeTest(t, "Test_Cov7_ValidValues", func() {
		// Arrange
		vv := corestr.NewValidValuesUsingValues(corestr.ValidValue{Value: "a", IsValid: true})

		// Act
		actual := args.Map{
			"len": vv.Length(),
			"isValid": vv.ValidValues[0].IsValid,
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues -- basic", actual)
	})
}

// ── ValueStatus ──

func Test_ValueStatus(t *testing.T) {
	safeTest(t, "Test_Cov7_ValueStatus", func() {
		// Arrange
		vs := corestr.ValueStatus{ValueValid: &corestr.ValidValue{Value: "hello", IsValid: true}, Index: 0}

		// Act
		actual := args.Map{
			"val": vs.ValueValid.Value,
			"isValid": vs.ValueValid.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus -- basic", actual)
	})
}

// ── HashsetsCollection ──

func Test_HashsetsCollection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov7_HashsetsCollection_Basic", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Cap(5)
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		hc.Add(h1)

		// Act
		actual := args.Map{
			"len":     hc.Length(),
			"isEmpty": hc.IsEmpty(),
			"hasAny":  hc.HasItems(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"isEmpty": false,
			"hasAny": true,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection basic -- 1 hashset", actual)
	})
}

// ── SimpleStringOnce ──

func Test_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_Cov7_SimpleStringOnce", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}
		s.SetOnceIfUninitialized("hello")

		// Act
		actual := args.Map{
			"val": s.Value(),
			"initialized": s.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"initialized": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce -- set once", actual)
	})
}

// ── KeyValuePair / KeyAnyValuePair ──

func Test_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov7_KeyValuePair", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"key": kv.Key,
			"val": kv.Value,
		}

		// Assert
		expected := args.Map{
			"key": "k",
			"val": "v",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair -- basic", actual)
	})
}

func Test_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov7_KeyAnyValuePair", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{
			"key": kv.Key,
			"val": kv.Value,
		}

		// Assert
		expected := args.Map{
			"key": "k",
			"val": 42,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair -- basic", actual)
	})
}

// ── TextWithLineNumber ──

func Test_TextWithLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov7_TextWithLineNumber", func() {
		// Arrange
		twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: 42}

		// Act
		actual := args.Map{
			"text": twl.Text,
			"line": twl.LineNumber,
		}

		// Assert
		expected := args.Map{
			"text": "hello",
			"line": 42,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber -- basic", actual)
	})
}

// ── Empty creator ──

func Test_Empty_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_SimpleSlice", func() {
		// Arrange
		s := corestr.Empty.SimpleSlice()

		// Act
		actual := args.Map{"isEmpty": s.IsEmpty()}

		// Assert
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty SimpleSlice -- empty", actual)
	})
}

func Test_Empty_Collection(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_Collection", func() {
		// Arrange
		c := corestr.Empty.Collection()

		// Act
		actual := args.Map{"isEmpty": c.IsEmpty()}

		// Assert
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty Collection -- empty", actual)
	})
}

func Test_Empty_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_Hashmap", func() {
		// Arrange
		h := corestr.Empty.Hashmap()

		// Act
		actual := args.Map{"isEmpty": h.IsEmpty()}

		// Assert
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty Hashmap -- empty", actual)
	})
}

func Test_Empty_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_Hashset", func() {
		// Arrange
		h := corestr.Empty.Hashset()

		// Act
		actual := args.Map{"isEmpty": h.IsEmpty()}

		// Assert
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty Hashset -- empty", actual)
	})
}

func Test_Empty_LinkedList(t *testing.T) {
	safeTest(t, "Test_Cov7_Empty_LinkedList", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()

		// Act
		actual := args.Map{"isEmpty": ll.IsEmpty()}

		// Assert
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Empty LinkedList -- empty", actual)
	})
}
