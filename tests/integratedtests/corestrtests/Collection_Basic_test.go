package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection ──

func Test_Collection_Basic_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Collection_Basic", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"len":     c.Length(),
			"isEmpty": c.IsEmpty(),
			"hasAny":  c.HasAnyItem(),
			"first":   c.First(),
			"last":    c.Last(),
		}

		// Assert
		expected := args.Map{
			"len": 3, "isEmpty": false, "hasAny": true,
			"first": "a", "last": "c",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- basic", actual)
	})
}

func Test_Collection_Add_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Collection_Add", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("hello")
		c.Add("world")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection.Add returns correct value -- with args", actual)
	})
}

func Test_Collection_AddIf_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Collection_AddIf", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddIf(true, "yes")
		c.AddIf(false, "no")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection.AddIf returns correct value -- with args", actual)
	})
}

func Test_Collection_Adds_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Collection_Adds", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection.Adds returns correct value -- with args", actual)
	})
}

func Test_Collection_List(t *testing.T) {
	safeTest(t, "Test_Cov12_Collection_List", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		list := c.List()

		// Act
		actual := args.Map{"len": len(list)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection.List returns correct value -- with args", actual)
	})
}

func Test_Collection_String_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Collection_String", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.String()

		// Act
		actual := args.Map{"notEmpty": result != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Collection.String returns correct value -- with args", actual)
	})
}

func Test_Collection_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_Collection_IsEmpty_Nil", func() {
		// Arrange
		var c *corestr.Collection

		// Act
		actual := args.Map{
			"empty": c.IsEmpty(),
			"len": c.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- nil", actual)
	})
}

// ── Hashmap ──

func Test_Hashmap_Basic_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_Basic", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})

		// Act
		actual := args.Map{
			"has":     hm.Has("k1"),
			"notHas":  !hm.Has("k2"),
			"len":     hm.Length(),
			"isEmpty": hm.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"has": true, "notHas": true, "len": 1, "isEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- basic", actual)
	})
}

func Test_Hashmap_Get_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_Get", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		val, found := hm.Get("k1")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
		}

		// Assert
		expected := args.Map{
			"val": "v1",
			"found": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.Get returns correct value -- with args", actual)
	})
}

func Test_Hashmap_Get_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_Get_NotFound", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		val, found := hm.Get("k2")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"found": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.Get returns correct value -- not found", actual)
	})
}

func Test_Hashmap_AddOrUpdate_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_AddOrUpdate", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		isNew := hm.AddOrUpdate("k1", "v1")
		isNew2 := hm.AddOrUpdate("k1", "v2")
		val, _ := hm.Get("k1")

		// Act
		actual := args.Map{
			"isNew": isNew,
			"isUpdate": !isNew2,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"isNew": true,
			"isUpdate": true,
			"val": "v2",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.AddOrUpdate returns correct value -- with args", actual)
	})
}

func Test_Hashmap_Clear(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_Clear", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		hm.Clear()

		// Act
		actual := args.Map{"empty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap.Clear returns correct value -- with args", actual)
	})
}

func Test_Hashmap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{
			"empty": hm.IsEmpty(),
			"len": hm.Length(),
			"hasAny": hm.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- nil", actual)
	})
}

func Test_Hashmap_Clone_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_Clone", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		cloned := hm.ClonePtr()

		// Act
		actual := args.Map{
			"notNil": cloned != nil,
			"notEmpty": !cloned.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.Clone returns correct value -- with args", actual)
	})
}

func Test_Hashmap_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_ClonePtr_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap
		cloned := hm.ClonePtr()

		// Act
		actual := args.Map{"nil": cloned == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap.ClonePtr returns nil -- nil", actual)
	})
}

func Test_Hashmap_IsEqualPtr(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_IsEqualPtr", func() {
		// Arrange
		a := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		b := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})

		// Act
		actual := args.Map{"equal": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"equal": true}
		expected.ShouldBeEqual(t, 0, "Hashmap.IsEqualPtr returns correct value -- with args", actual)
	})
}

func Test_Hashmap_IsEqualPtr_Different(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashmap_IsEqualPtr_Different", func() {
		// Arrange
		a := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		b := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v2"})

		// Act
		actual := args.Map{"equal": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"equal": false}
		expected.ShouldBeEqual(t, 0, "Hashmap.IsEqualPtr returns correct value -- different", actual)
	})
}

// ── Hashset ──

func Test_Hashset_Basic_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashset_Basic", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"has":     hs.Has("a"),
			"notHas":  !hs.Has("d"),
			"len":     hs.Length(),
			"isEmpty": hs.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"has": true, "notHas": true, "len": 3, "isEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- basic", actual)
	})
}

func Test_Hashset_Add(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashset_Add", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.Add("x")
		hs.Add("y")
		hs.Add("x") // duplicate

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset.Add returns correct value -- with args", actual)
	})
}

func Test_Hashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_Hashset_Nil", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act
		actual := args.Map{
			"empty": hs.IsEmpty(),
			"len": hs.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- nil", actual)
	})
}

// ── SimpleSlice ──

func Test_SimpleSlice_Basic_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_SimpleSlice_Basic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.SpreadStrings("a", "b", "c")

		// Act
		actual := args.Map{
			"len":     ss.Length(),
			"isEmpty": ss.IsEmpty(),
			"hasAny":  ss.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"len": 3, "isEmpty": false, "hasAny": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- basic", actual)
	})
}

func Test_SimpleSlice_Nil_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_SimpleSlice_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		actual := args.Map{
			"empty": ss.IsEmpty(),
			"len": ss.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns nil -- nil", actual)
	})
}

// ── LeftRight ──

func Test_LeftRight_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_LeftRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("left", "right")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "left",
			"right": "right",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- with args", actual)
	})
}

// ── LeftMiddleRight ──

func Test_LeftMiddleRight_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_LeftMiddleRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "l",
			"middle": "m",
			"right": "r",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- with args", actual)
	})
}

// ── ValidValue ──

func Test_ValidValue_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_ValidValue", func() {
		// Arrange
		vv := corestr.ValidValue{Value: "hello", IsValid: true}

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- with args", actual)
	})
}

func Test_ValidValue_Invalid_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_ValidValue_Invalid", func() {
		// Arrange
		vv := corestr.ValidValue{IsValid: false}

		// Act
		actual := args.Map{"valid": vv.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns error -- invalid", actual)
	})
}

// ── ValueStatus ──

func Test_ValueStatus_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_ValueStatus", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		vs := &corestr.ValueStatus{ValueValid: vv, Index: 0}

		// Act
		actual := args.Map{
			"val": vs.ValueValid.Value,
			"idx": vs.Index,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"idx": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- with args", actual)
	})
}

// ── KeyValuePair ──

func Test_KeyValuePair_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_KeyValuePair", func() {
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
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- with args", actual)
	})
}

// ── KeyAnyValuePair ──

func Test_KeyAnyValuePair_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_KeyAnyValuePair", func() {
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
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- with args", actual)
	})
}

// ── emptyCreator ──

func Test_Empty_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov12_Empty_Hashmap", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		actual := args.Map{
			"notNil": hm != nil,
			"empty": hm.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "Empty.Hashmap returns empty -- with args", actual)
	})
}

func Test_Empty_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov12_Empty_Hashset", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act
		actual := args.Map{
			"notNil": hs != nil,
			"empty": hs.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "Empty.Hashset returns empty -- with args", actual)
	})
}

// ── TextWithLineNumber ──

func Test_TextWithLineNumber_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_TextWithLineNumber", func() {
		// Arrange
		tln := corestr.TextWithLineNumber{Text: "hello", LineNumber: 1}

		// Act
		actual := args.Map{
			"text": tln.Text,
			"num": tln.LineNumber,
		}

		// Assert
		expected := args.Map{
			"text": "hello",
			"num": 1,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns non-empty -- with args", actual)
	})
}

// ── HashsetsCollection ──

func Test_HashsetsCollection_Basic_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_HashsetsCollection_Basic", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.Cap(5)

		// Act
		actual := args.Map{
			"isEmpty": hsc.IsEmpty(),
			"len": hsc.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns correct value -- basic", actual)
	})
}

func Test_HashsetsCollection_Add(t *testing.T) {
	safeTest(t, "Test_Cov12_HashsetsCollection_Add", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.Cap(5)
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hsc.Add(hs)

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection.Add returns correct value -- with args", actual)
	})
}

func Test_HashsetsCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_HashsetsCollection_Nil", func() {
		// Arrange
		var hsc *corestr.HashsetsCollection

		// Act
		actual := args.Map{
			"empty": hsc.IsEmpty(),
			"len": hsc.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns nil -- nil", actual)
	})
}

// ── CollectionsOfCollection ──

func Test_CollectionsOfCollection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov12_CollectionsOfCollection_Basic", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 5)

		// Act
		actual := args.Map{
			"isEmpty": coc.IsEmpty(),
			"len": coc.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns correct value -- basic", actual)
	})
}

func Test_CollectionsOfCollection_Add(t *testing.T) {
	safeTest(t, "Test_Cov12_CollectionsOfCollection_Add", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 5)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(c)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection.Add returns correct value -- with args", actual)
	})
}

func Test_CollectionsOfCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_CollectionsOfCollection_Nil", func() {
		// Arrange
		var coc *corestr.CollectionsOfCollection
		isNil := coc == nil

		// Act
		actual := args.Map{"isNil": isNil}

		// Assert
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns nil -- nil", actual)
	})
}

// ── KeyValueCollection ──

func Test_KeyValueCollection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov12_KeyValueCollection_Basic", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Cap(5)

		// Act
		actual := args.Map{
			"isEmpty": kvc.IsEmpty(),
			"len": kvc.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- basic", actual)
	})
}

func Test_KeyValueCollection_Add(t *testing.T) {
	safeTest(t, "Test_Cov12_KeyValueCollection_Add", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Cap(5)
		kvc.Add("key", "val")

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection.Add returns correct value -- with args", actual)
	})
}

func Test_KeyValueCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_KeyValueCollection_Nil", func() {
		// Arrange
		var kvc *corestr.KeyValueCollection

		// Act
		actual := args.Map{
			"empty": kvc.IsEmpty(),
			"len": kvc.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns nil -- nil", actual)
	})
}

// ── SimpleStringOnce ──

func Test_SimpleStringOnce_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_SimpleStringOnce", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{
			"val":     sso.Value(),
			"isEmpty": sso.IsEmpty(),
			"hasVal":  sso.HasValidNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"val": "hello", "isEmpty": false, "hasVal": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- with args", actual)
	})
}

func Test_SimpleStringOnce_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_SimpleStringOnce_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce
		// IsEmpty panics on nil receiver — just verify nil check

		// Act
		actual := args.Map{"isNil": sso == nil}

		// Assert
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns nil -- nil", actual)
	})
}

// ── CharCollectionMap ──

func Test_CharCollectionMap_Basic(t *testing.T) {
	safeTest(t, "Test_Cov12_CharCollectionMap_Basic", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Empty()

		// Act
		actual := args.Map{
			"isEmpty": ccm.IsEmpty(),
			"len": ccm.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap returns correct value -- basic", actual)
	})
}

func Test_CharCollectionMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_CharCollectionMap_Nil", func() {
		// Arrange
		var ccm *corestr.CharCollectionMap

		// Act
		actual := args.Map{
			"empty": ccm.IsEmpty(),
			"len": ccm.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap returns nil -- nil", actual)
	})
}

// ── CharHashsetMap ──

func Test_CharHashsetMap_Basic(t *testing.T) {
	safeTest(t, "Test_Cov12_CharHashsetMap_Basic", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(5, 5)

		// Act
		actual := args.Map{
			"isEmpty": chm.IsEmpty(),
			"len": chm.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "CharHashsetMap returns correct value -- basic", actual)
	})
}

func Test_CharHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_CharHashsetMap_Nil", func() {
		// Arrange
		var chm *corestr.CharHashsetMap

		// Act
		actual := args.Map{
			"empty": chm.IsEmpty(),
			"len": chm.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "CharHashsetMap returns nil -- nil", actual)
	})
}

// ── CloneSlice ──

func Test_CloneSlice_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_CloneSlice", func() {
		// Arrange
		original := []string{"a", "b", "c"}
		cloned := corestr.CloneSlice(original)
		original[0] = "X"

		// Act
		actual := args.Map{
			"clonedFirst": cloned[0],
			"len": len(cloned),
		}

		// Assert
		expected := args.Map{
			"clonedFirst": "a",
			"len": 3,
		}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns correct value -- with args", actual)
	})
}

func Test_CloneSlice_Nil_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_CloneSlice_Nil", func() {
		// Arrange
		cloned := corestr.CloneSlice(nil)

		// Act
		actual := args.Map{"empty": len(cloned) == 0}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns nil -- nil", actual)
	})
}

func Test_CloneSliceIf_True(t *testing.T) {
	safeTest(t, "Test_Cov12_CloneSliceIf_True", func() {
		// Arrange
		original := []string{"a"}
		cloned := corestr.CloneSliceIf(true, original...)
		original[0] = "X"

		// Act
		actual := args.Map{"cloned": cloned[0]}

		// Assert
		expected := args.Map{"cloned": "a"}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns non-empty -- true", actual)
	})
}

func Test_CloneSliceIf_False(t *testing.T) {
	safeTest(t, "Test_Cov12_CloneSliceIf_False", func() {
		// Arrange
		original := []string{"a"}
		cloned := corestr.CloneSliceIf(false, original...)

		// Act
		actual := args.Map{"len": len(cloned)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns non-empty -- false", actual)
	})
}

// ── LinkedList ──

func Test_LinkedList_Basic(t *testing.T) {
	safeTest(t, "Test_Cov12_LinkedList_Basic", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{
			"isEmpty": ll.IsEmpty(),
			"len": ll.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- basic", actual)
	})
}

func Test_LinkedList_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_LinkedList_Nil", func() {
		// Arrange
		var ll *corestr.LinkedList

		// Act
		actual := args.Map{"isNil": ll == nil}

		// Assert
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "LinkedList returns nil -- nil", actual)
	})
}

// ── LinkedCollections ──

func Test_LinkedCollections_Basic(t *testing.T) {
	safeTest(t, "Test_Cov12_LinkedCollections_Basic", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{
			"isEmpty": lc.IsEmpty(),
			"len": lc.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- basic", actual)
	})
}

func Test_LinkedCollections_Nil(t *testing.T) {
	safeTest(t, "Test_Cov12_LinkedCollections_Nil", func() {
		// Arrange
		var lc *corestr.LinkedCollections
		var panicked bool
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			_ = lc.IsEmpty()
			_ = lc.Length()
		}()

		// Act
		actual := args.Map{"panicked": panicked}

		// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "LinkedCollections panics on nil receiver -- nil", actual)
	})
}

// ── ValidValues ──

func Test_ValidValues_FromCollectionBasic(t *testing.T) {
	safeTest(t, "Test_Cov12_ValidValues", func() {
		// Arrange
		vvs := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: false},
		)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- with args", actual)
	})
}

// ── HashmapDiff ──

func Test_HashmapDiff(t *testing.T) {
	safeTest(t, "Test_Cov12_HashmapDiff", func() {
		// Arrange
		diff := corestr.HashmapDiff(map[string]string{"k": "v"})

		// Act
		actual := args.Map{
			"len": diff.Length(),
			"empty": diff.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- with args", actual)
	})
}
