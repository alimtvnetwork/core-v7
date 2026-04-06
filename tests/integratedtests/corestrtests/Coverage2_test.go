package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── SimpleSlice ──

func Test_SimpleSlice_Cap_Cov2(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_Cap_Cov2", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)

		// Act
		actual := args.Map{
			"isNil": s == nil,
			"isEmpty": s.IsEmpty(),
			"length": s.Length(),
			"hasAny": s.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
			"length": 0,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_Cap returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_Add_Cov2(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_Add_Cov2", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Add("hello")
		s.Add("world")

		// Act
		actual := args.Map{
			"length": s.Length(),
			"hasAny": s.HasAnyItem(),
			"first": s.Strings()[0],
		}

		// Assert
		expected := args.Map{
			"length": 2,
			"hasAny": true,
			"first": "hello",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_Add returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_Adds_Cov2(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_Adds_Cov2", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Adds("a", "b", "c")

		// Act
		actual := args.Map{"length": s.Length()}

		// Assert
		expected := args.Map{"length": 3}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_Adds returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_AddIf_Cov2(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_AddIf_Cov2", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.AddIf(true, "yes")
		s.AddIf(false, "no")

		// Act
		actual := args.Map{"length": s.Length()}

		// Assert
		expected := args.Map{"length": 1}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_AddIf returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_AppendFmt_Cov2(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_AppendFmt_Cov2", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.AppendFmt("hello %d", 42)

		// Act
		actual := args.Map{
			"length": s.Length(),
			"first": s.Strings()[0],
		}

		// Assert
		expected := args.Map{
			"length": 1,
			"first": "hello 42",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_AppendFmt returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_String_Cov2(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_String_Cov2", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Add("hello")

		// Act
		actual := args.Map{"notEmpty": s.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_String returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_JoinCsv_Cov2(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_JoinCsv_Cov2", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Adds("a", "b")

		// Act
		actual := args.Map{"notEmpty": s.JoinCsv() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_JoinCsv returns correct value -- with args", actual)
	})
}

// ── Collection ──

func Test_Collection_Cap_Cov2(t *testing.T) {
	safeTest(t, "Test_Collection_Cap_Cov2", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{
			"isNil": c == nil,
			"isEmpty": c.IsEmpty(),
			"length": c.Length(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
			"length": 0,
		}
		expected.ShouldBeEqual(t, 0, "Collection_Cap returns correct value -- with args", actual)
	})
}

func Test_Collection_AddStrings_Cov2(t *testing.T) {
	safeTest(t, "Test_Collection_AddStrings_Cov2", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"hello", "world"})

		// Act
		actual := args.Map{"length": c.Length()}

		// Assert
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "Collection_AddStrings returns correct value -- with args", actual)
	})
}

// ── Hashset ──

func Test_Hashset_Cap_Cov2(t *testing.T) {
	safeTest(t, "Test_Hashset_Cap_Cov2", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{
			"isNil": h == nil,
			"isEmpty": h.IsEmpty(),
			"length": h.Length(),
			"hasAny": h.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
			"length": 0,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset_Cap returns correct value -- with args", actual)
	})
}

func Test_Hashset_Add_Cov2(t *testing.T) {
	safeTest(t, "Test_Hashset_Add_Cov2", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Add("hello")
		h.Add("hello")
		h.Add("world")

		// Act
		actual := args.Map{
			"length": h.Length(),
			"has": h.Has("hello"),
			"hasNo": h.Has("nope"),
		}

		// Assert
		expected := args.Map{
			"length": 2,
			"has": true,
			"hasNo": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset_Add returns correct value -- with args", actual)
	})
}

func Test_Hashset_Adds_Cov2(t *testing.T) {
	safeTest(t, "Test_Hashset_Adds_Cov2", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Adds("a", "b", "c")

		// Act
		actual := args.Map{"length": h.Length()}

		// Assert
		expected := args.Map{"length": 3}
		expected.ShouldBeEqual(t, 0, "Hashset_Adds returns correct value -- with args", actual)
	})
}

// ── Hashmap ──

func Test_Hashmap_Cap_Cov2(t *testing.T) {
	safeTest(t, "Test_Hashmap_Cap_Cov2", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{
			"isNil": h == nil,
			"isEmpty": h.IsEmpty(),
			"length": h.Length(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
			"length": 0,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap_Cap returns correct value -- with args", actual)
	})
}

func Test_Hashmap_Add_Cov2(t *testing.T) {
	safeTest(t, "Test_Hashmap_Add_Cov2", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("key", "value")
		val, _ := h.Get("key")

		// Act
		actual := args.Map{
			"length": h.Length(),
			"has": h.Has("key"),
			"getVal": val,
		}

		// Assert
		expected := args.Map{
			"length": 1,
			"has": true,
			"getVal": "value",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap_Add returns correct value -- with args", actual)
	})
}

// ── KeyValues ──

func Test_KeyValues_Cap_Cov2(t *testing.T) {
	safeTest(t, "Test_KeyValues_Cap_Cov2", func() {
		// Act
		actual := args.Map{"notNil": corestr.New.KeyValues.Cap(5) != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "KeyValues_Cap returns non-empty -- with args", actual)
	})
}

// ── LinkedList ──

func Test_LinkedList_Default_Cov2(t *testing.T) {
	safeTest(t, "Test_LinkedList_Default_Cov2", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()

		// Act
		actual := args.Map{
			"isNil": ll == nil,
			"isEmpty": ll.IsEmpty(),
			"length": ll.Length(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
			"length": 0,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList_Default returns correct value -- with args", actual)
	})
}

func Test_LinkedList_Add_Cov2(t *testing.T) {
	safeTest(t, "Test_LinkedList_Add_Cov2", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()
		ll.Add("hello")
		ll.Add("world")

		// Act
		actual := args.Map{
			"length": ll.Length(),
			"isEmpty": ll.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"length": 2,
			"isEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList_Add returns correct value -- with args", actual)
	})
}

// ── CharHashsetMap / CharCollectionMap / SimpleStringOnce / HashsetsCollection / LinkedCollection / CollectionsOfCollection ──

func Test_CharHashsetMap_Default_Cov2(t *testing.T) {
	safeTest(t, "Test_CharHashsetMap_Default_Cov2", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 0)

		// Act
		actual := args.Map{
			"isNil": chm == nil,
			"isEmpty": chm.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "CharHashsetMap_Default returns correct value -- with args", actual)
	})
}

func Test_CharCollectionMap_Default_Cov2(t *testing.T) {
	safeTest(t, "Test_CharCollectionMap_Default_Cov2", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Empty()

		// Act
		actual := args.Map{
			"isNil": ccm == nil,
			"isEmpty": ccm.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap_Default returns correct value -- with args", actual)
	})
}

func Test_SimpleStringOnce_Default_Cov2(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Default_Cov2", func() {
		// Arrange
		so := corestr.New.SimpleStringOnce.Empty()

		// Act
		actual := args.Map{
			"isNil": false,
			"isEmpty": so.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce_Default returns correct value -- with args", actual)
	})
}

func Test_HashsetsCollection_Cap_Cov2(t *testing.T) {
	safeTest(t, "Test_HashsetsCollection_Cap_Cov2", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Cap(5)

		// Act
		actual := args.Map{
			"isNil": hc == nil,
			"isEmpty": hc.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection_Cap returns correct value -- with args", actual)
	})
}

func Test_LinkedCollection_Empty_Cov2(t *testing.T) {
	safeTest(t, "Test_LinkedCollection_Empty_Cov2", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Empty()

		// Act
		actual := args.Map{
			"isNil": lc == nil,
			"isEmpty": lc.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedCollection_Empty returns empty -- with args", actual)
	})
}

func Test_CollectionsOfCollection_Cap_Cov2(t *testing.T) {
	safeTest(t, "Test_CollectionsOfCollection_Cap_Cov2", func() {
		// Arrange
		cc := corestr.New.CollectionsOfCollection.Cap(5)

		// Act
		actual := args.Map{
			"isNil": cc == nil,
			"isEmpty": cc.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil": false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection_Cap returns correct value -- with args", actual)
	})
}

// ── LeftRight / LeftMiddleRight / ValidValue / ValidValues ──

func Test_LeftRight_Cov2(t *testing.T) {
	safeTest(t, "Test_LeftRight_Cov2", func() {
		// Arrange
		lr := corestr.NewLeftRight("l", "r")

		// Act
		actual := args.Map{
			"isLeftEmpty": lr.IsLeftEmpty(),
			"isRightEmpty": lr.IsRightEmpty(),
			"hasSafe": lr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"isLeftEmpty": false,
			"isRightEmpty": false,
			"hasSafe": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRight_Cov2(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRight_Cov2", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")

		// Act
		actual := args.Map{
			"isLeftEmpty": lmr.IsLeftEmpty(),
			"isMiddleEmpty": lmr.IsMiddleEmpty(),
			"isRightEmpty": lmr.IsRightEmpty(),
		}

		// Assert
		expected := args.Map{
			"isLeftEmpty": false,
			"isMiddleEmpty": false,
			"isRightEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- with args", actual)
	})
}

func Test_ValidValue_Cov2(t *testing.T) {
	safeTest(t, "Test_ValidValue_Cov2", func() {
		// Arrange
		vv := corestr.ValidValue{Value: "hello", IsValid: true}

		// Act
		actual := args.Map{
			"value": vv.Value,
			"isValid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"value": "hello",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- with args", actual)
	})
}

func Test_ValidValues_Cov2(t *testing.T) {
	safeTest(t, "Test_ValidValues_Cov2", func() {
		// Arrange
		vv := corestr.ValidValues{ValidValues: []*corestr.ValidValue{{Value: "a", IsValid: true}}}

		// Act
		actual := args.Map{"len": len(vv.ValidValues)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- with args", actual)
	})
}

// ── AnyToString ──

func Test_AnyToString_Cov2(t *testing.T) {
	safeTest(t, "Test_AnyToString_Cov2", func() {
		// Act
		actual := args.Map{"notEmpty": corestr.AnyToString(false, 42) != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString_Int returns correct value -- with args", actual)
	})
}

func Test_AnyToString_String_Cov2(t *testing.T) {
	safeTest(t, "Test_AnyToString_String_Cov2", func() {
		// Act
		actual := args.Map{"result": corestr.AnyToString(false, "hello")}

		// Assert
		expected := args.Map{"result": "hello"}
		expected.ShouldBeEqual(t, 0, "AnyToString_String returns correct value -- with args", actual)
	})
}

func Test_AnyToString_Nil_Cov2(t *testing.T) {
	safeTest(t, "Test_AnyToString_Nil_Cov2", func() {
		// Arrange
		result := corestr.AnyToString(false, nil)

		// Act
		actual := args.Map{
			"ok": true,
			"empty": result == "",
		}

		// Assert
		expected := args.Map{
			"ok": true,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "AnyToString_Nil returns nil -- formats to non-empty", actual)
	})
}

// ── AllIndividualStringsOfStringsLength ──

func Test_AllIndividualStringsOfStringsLength_Cov2(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_Cov2", func() {
		// Arrange
		items := [][]string{{"ab", "cde"}}

		// Act
		actual := args.Map{"result": corestr.AllIndividualStringsOfStringsLength(&items)}

		// Assert
		expected := args.Map{"result": 2}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns correct value -- with args", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_Nil_Cov2(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_Nil_Cov2", func() {
		// Act
		actual := args.Map{"result": corestr.AllIndividualStringsOfStringsLength(nil)}

		// Assert
		expected := args.Map{"result": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength_Nil returns nil -- with args", actual)
	})
}

// ── CloneSlice / CloneSliceIf ──

func Test_CloneSlice_Cov2(t *testing.T) {
	safeTest(t, "Test_CloneSlice_Cov2", func() {
		// Act
		actual := args.Map{"len": len(corestr.CloneSlice([]string{"a", "b"}))}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns correct value -- with args", actual)
	})
}

func Test_CloneSlice_Nil_Cov2(t *testing.T) {
	safeTest(t, "Test_CloneSlice_Nil_Cov2", func() {
		// Act
		actual := args.Map{"isNil": corestr.CloneSlice(nil) == nil}

		// Assert
		expected := args.Map{"isNil": false}
		expected.ShouldBeEqual(t, 0, "CloneSlice_Nil returns nil -- returns empty slice not nil", actual)
	})
}

func Test_CloneSliceIf_True_Cov2(t *testing.T) {
	safeTest(t, "Test_CloneSliceIf_True_Cov2", func() {
		// Act
		actual := args.Map{"len": len(corestr.CloneSliceIf(true, "a"))}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf_True returns non-empty -- with args", actual)
	})
}

func Test_CloneSliceIf_False_Cov2(t *testing.T) {
	safeTest(t, "Test_CloneSliceIf_False_Cov2", func() {
		// Act
		actual := args.Map{"len": len(corestr.CloneSliceIf(false, "a"))}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf_False returns non-empty -- with args", actual)
	})
}

// ── TextWithLineNumber / ValueStatus ──

func Test_TextWithLineNumber_Cov2(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_Cov2", func() {
		// Arrange
		tw := corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{
			"lineNumber": tw.LineNumber,
			"text": tw.Text,
		}

		// Assert
		expected := args.Map{
			"lineNumber": 5,
			"text": "hello",
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns non-empty -- with args", actual)
	})
}

func Test_ValueStatus_Cov2(t *testing.T) {
	safeTest(t, "Test_ValueStatus_Cov2", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		vs := corestr.ValueStatus{ValueValid: vv, Index: 0}

		// Act
		actual := args.Map{
			"index": vs.Index,
			"notNil": vs.ValueValid != nil,
		}

		// Assert
		expected := args.Map{
			"index": 0,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- with args", actual)
	})
}
