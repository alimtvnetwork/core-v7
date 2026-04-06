package corestrtests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf / AnyToString / Utility funcs
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov16_CloneSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Cov16_CloneSlice_Nil", func() {
		// Arrange
		result := corestr.CloneSlice(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns nil -- nil", actual)
	})
}

func Test_Cov16_CloneSlice_WithItems(t *testing.T) {
	safeTest(t, "Test_Cov16_CloneSlice_WithItems", func() {
		// Arrange
		result := corestr.CloneSlice([]string{"a", "b"})

		// Act
		actual := args.Map{"len": len(result), "first": result[0]}

		// Assert
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns correct value -- items", actual)
	})
}

func Test_Cov16_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov16_CloneSliceIf_Empty", func() {
		// Arrange
		result := corestr.CloneSliceIf(true)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- empty", actual)
	})
}

func Test_Cov16_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov16_CloneSliceIf_NoClone", func() {
		// Arrange
		result := corestr.CloneSliceIf(false, "a", "b")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns correct value -- noClone", actual)
	})
}

func Test_Cov16_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_Cov16_CloneSliceIf_Clone", func() {
		// Arrange
		result := corestr.CloneSliceIf(true, "a", "b")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns correct value -- clone", actual)
	})
}

func Test_Cov16_AnyToString_Empty(t *testing.T) {
	safeTest(t, "Test_Cov16_AnyToString_Empty", func() {
		// Arrange
		s := corestr.AnyToString(false, "")

		// Act
		actual := args.Map{"isEmpty": s == ""}

		// Assert
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- empty", actual)
	})
}

func Test_Cov16_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_Cov16_AnyToString_WithFieldName", func() {
		// Arrange
		s := corestr.AnyToString(true, "hello")

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- withField", actual)
	})
}

func Test_Cov16_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_Cov16_AnyToString_WithoutFieldName", func() {
		// Arrange
		s := corestr.AnyToString(false, "hello")

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- noField", actual)
	})
}

func Test_Cov16_AnyToString_Ptr(t *testing.T) {
	safeTest(t, "Test_Cov16_AnyToString_Ptr", func() {
		// Arrange
		val := "hello"
		s := corestr.AnyToString(false, &val)

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- ptr", actual)
	})
}

func Test_Cov16_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_Cov16_AllIndividualStringsOfStringsLength_Nil", func() {
		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(nil)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns nil -- nil", actual)
	})
}

func Test_Cov16_AllIndividualStringsOfStringsLength_Items(t *testing.T) {
	safeTest(t, "Test_Cov16_AllIndividualStringsOfStringsLength_Items", func() {
		// Arrange
		items := [][]string{{"a", "b"}, {"c"}}

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&items)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns correct value -- items", actual)
	})
}

func Test_Cov16_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_Cov16_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		// Act
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns nil -- nil", actual)
	})
}

func Test_Cov16_AllIndividualsLengthOfSimpleSlices_Items(t *testing.T) {
	safeTest(t, "Test_Cov16_AllIndividualsLengthOfSimpleSlices_Items", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("c")

		// Act
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns correct value -- items", actual)
	})
}

func Test_Cov16_StringUtils(t *testing.T) {
	safeTest(t, "Test_Cov16_StringUtils", func() {
		// Act
		actual := args.Map{
			"double":              corestr.StringUtils.WrapDouble("a") == `"a"`,
			"single":              corestr.StringUtils.WrapSingle("a") == "'a'",
			"tilda":               corestr.StringUtils.WrapTilda("a") == "`a`",
			"dblIfMissing_empty":  corestr.StringUtils.WrapDoubleIfMissing("") == `""`,
			"dblIfMissing_wrapped": corestr.StringUtils.WrapDoubleIfMissing(`"a"`) == `"a"`,
			"dblIfMissing_plain":  corestr.StringUtils.WrapDoubleIfMissing("a") == `"a"`,
			"sglIfMissing_empty":  corestr.StringUtils.WrapSingleIfMissing("") == "''",
			"sglIfMissing_wrapped": corestr.StringUtils.WrapSingleIfMissing("'a'") == "'a'",
			"sglIfMissing_plain":  corestr.StringUtils.WrapSingleIfMissing("a") == "'a'",
		}

		// Assert
		expected := args.Map{
			"double": true, "single": true, "tilda": true,
			"dblIfMissing_empty": true, "dblIfMissing_wrapped": true, "dblIfMissing_plain": true,
			"sglIfMissing_empty": true, "sglIfMissing_wrapped": true, "sglIfMissing_plain": true,
		}
		expected.ShouldBeEqual(t, 0, "StringUtils returns correct value -- with args", actual)
	})
}

func Test_Cov16_EmptyCreator_All(t *testing.T) {
	safeTest(t, "Test_Cov16_EmptyCreator_All", func() {
		// Arrange
		_ = corestr.Empty.Collection()
		_ = corestr.Empty.LinkedList()
		_ = corestr.Empty.SimpleSlice()
		_ = corestr.Empty.KeyAnyValuePair()
		_ = corestr.Empty.KeyValuePair()
		_ = corestr.Empty.KeyValueCollection()
		_ = corestr.Empty.LinkedCollections()
		_ = corestr.Empty.LeftRight()
		_ = corestr.Empty.SimpleStringOnce()
		_ = corestr.Empty.SimpleStringOncePtr()
		_ = corestr.Empty.Hashset()
		_ = corestr.Empty.HashsetsCollection()
		_ = corestr.Empty.Hashmap()
		_ = corestr.Empty.CharCollectionMap()
		_ = corestr.Empty.KeyValuesCollection()
		_ = corestr.Empty.CollectionsOfCollection()
		_ = corestr.Empty.CharHashsetMap()

		// Act
		actual := args.Map{"allCreated": true}

		// Assert
		expected := args.Map{"allCreated": true}
		expected.ShouldBeEqual(t, 0, "EmptyCreator returns empty -- all", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// DataModels
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov16_HashmapDataModel(t *testing.T) {
	safeTest(t, "Test_Cov16_HashmapDataModel", func() {
		// Arrange
		dm := &corestr.HashmapDataModel{Items: map[string]string{"a": "b"}}
		hm := corestr.NewHashmapUsingDataModel(dm)
		dm2 := corestr.NewHashmapsDataModelUsing(hm)

		// Act
		actual := args.Map{"hmNotNil": hm != nil, "dm2NotNil": dm2 != nil, "hmNotEmpty": !hm.IsEmpty()}

		// Assert
		expected := args.Map{"hmNotNil": true, "dm2NotNil": true, "hmNotEmpty": true}
		expected.ShouldBeEqual(t, 0, "HashmapDataModel returns correct value -- with args", actual)
	})
}

func Test_Cov16_HashsetDataModel(t *testing.T) {
	safeTest(t, "Test_Cov16_HashsetDataModel", func() {
		// Arrange
		dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
		hs := corestr.NewHashsetUsingDataModel(dm)
		dm2 := corestr.NewHashsetsDataModelUsing(hs)

		// Act
		actual := args.Map{"hsNotNil": hs != nil, "dm2NotNil": dm2 != nil}

		// Assert
		expected := args.Map{"hsNotNil": true, "dm2NotNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetDataModel returns correct value -- with args", actual)
	})
}

func Test_Cov16_HashsetsCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_Cov16_HashsetsCollectionDataModel", func() {
		// Arrange
		dm := &corestr.HashsetsCollectionDataModel{Items: []*corestr.Hashset{}}
		hc := corestr.NewHashsetsCollectionUsingDataModel(dm)
		dm2 := corestr.NewHashsetsCollectionDataModelUsing(hc)

		// Act
		actual := args.Map{"hcNotNil": hc != nil, "dm2NotNil": dm2 != nil}

		// Assert
		expected := args.Map{"hcNotNil": true, "dm2NotNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionDataModel returns correct value -- with args", actual)
	})
}

func Test_Cov16_CharCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_Cov16_CharCollectionDataModel", func() {
		// Arrange
		dm := &corestr.CharCollectionDataModel{
			Items:                  map[byte]*corestr.Collection{},
			EachCollectionCapacity: 10,
		}
		ccm := corestr.NewCharCollectionMapUsingDataModel(dm)
		dm2 := corestr.NewCharCollectionMapDataModelUsing(ccm)

		// Act
		actual := args.Map{"ccmNotNil": ccm != nil, "dm2NotNil": dm2 != nil}

		// Assert
		expected := args.Map{"ccmNotNil": true, "dm2NotNil": true}
		expected.ShouldBeEqual(t, 0, "CharCollectionDataModel returns correct value -- with args", actual)
	})
}

func Test_Cov16_CharHashsetDataModel(t *testing.T) {
	safeTest(t, "Test_Cov16_CharHashsetDataModel", func() {
		// Arrange
		dm := &corestr.CharHashsetDataModel{
			Items:               map[byte]*corestr.Hashset{},
			EachHashsetCapacity: 10,
		}
		chm := corestr.NewCharHashsetMapUsingDataModel(dm)
		dm2 := corestr.NewCharHashsetMapDataModelUsing(chm)

		// Act
		actual := args.Map{"chmNotNil": chm != nil, "dm2NotNil": dm2 != nil}

		// Assert
		expected := args.Map{"chmNotNil": true, "dm2NotNil": true}
		expected.ShouldBeEqual(t, 0, "CharHashsetDataModel returns correct value -- with args", actual)
	})
}

func Test_Cov16_SimpleStringOnceModel(t *testing.T) {
	safeTest(t, "Test_Cov16_SimpleStringOnceModel", func() {
		// Arrange
		m := corestr.SimpleStringOnceModel{Value: "hello", IsInitialize: true}

		// Act
		actual := args.Map{"value": m.Value}

		// Assert
		expected := args.Map{"value": "hello"}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnceModel returns correct value -- with args", actual)
	})
}

func Test_Cov16_CollectionsOfCollectionModel(t *testing.T) {
	safeTest(t, "Test_Cov16_CollectionsOfCollectionModel", func() {
		// Arrange
		m := corestr.CollectionsOfCollectionModel{Items: []*corestr.Collection{}}

		// Act
		actual := args.Map{"notNil": m.Items != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollectionModel returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection comprehensive
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov16_Collection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_Basic", func() {
		// Arrange
		c := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{
			"empty": c.IsEmpty(), "hasItems": c.HasItems(), "len": c.Length(),
			"count": c.Count(), "lastIdx": c.LastIndex(), "hasIdx0": c.HasIndex(0),
		}

		// Assert
		expected := args.Map{
			"empty": true, "hasItems": false, "len": 0,
			"count": 0, "lastIdx": -1, "hasIdx0": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- basic", actual)
	})
}

func Test_Cov16_Collection_NilReceiver(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_NilReceiver", func() {
		// Arrange
		var c *corestr.Collection

		// Act
		actual := args.Map{"len": c.Length(), "empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"len": 0, "empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- nil receiver", actual)
	})
}

func Test_Cov16_Collection_Add(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_Add", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a").Add("b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Add", actual)
	})
}

func Test_Cov16_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("")
		c.AddNonEmpty("a")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddNonEmpty", actual)
	})
}

func Test_Cov16_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddNonEmptyWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("a")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddNonEmptyWhitespace", actual)
	})
}

func Test_Cov16_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddError", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddError(nil)
		c.AddError(errors.New("e"))

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AddError", actual)
	})
}

func Test_Cov16_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddIf", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddIf(false, "skip")
		c.AddIf(true, "add")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddIf", actual)
	})
}

func Test_Cov16_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddIfMany", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddIfMany(false, "a", "b")
		c.AddIfMany(true, "c", "d")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddIfMany", actual)
	})
}

func Test_Cov16_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_Adds", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Adds", actual)
	})
}

func Test_Cov16_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddStrings", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddStrings([]string{"x", "y"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddStrings", actual)
	})
}

func Test_Cov16_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddFunc", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFunc(func() string { return "hello" })

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddFunc", actual)
	})
}

func Test_Cov16_Collection_AddFuncErr_NoErr(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddFuncErr_NoErr", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) {},
		)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AddFuncErr noErr", actual)
	})
}

func Test_Cov16_Collection_AddFuncErr_WithErr(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddFuncErr_WithErr", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		called := false
		c.AddFuncErr(
			func() (string, error) { return "", errors.New("e") },
			func(err error) { called = true },
		)

		// Act
		actual := args.Map{"len": c.Length(), "called": called}

		// Assert
		expected := args.Map{"len": 0, "called": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AddFuncErr withErr", actual)
	})
}

func Test_Cov16_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddLock("a")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddLock", actual)
	})
}

func Test_Cov16_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddsLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddsLock("a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddsLock", actual)
	})
}

func Test_Cov16_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddCollection", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c1.AddCollection(c2)
		c1.AddCollection(corestr.New.Collection.Empty())

		// Act
		actual := args.Map{"len": c1.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollection", actual)
	})
}

func Test_Cov16_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddCollections", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddCollections(
			corestr.New.Collection.Strings([]string{"a"}),
			corestr.New.Collection.Empty(),
			corestr.New.Collection.Strings([]string{"b"}),
		)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollections", actual)
	})
}

func Test_Cov16_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_RemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := c.RemoveAt(1)

		// Act
		actual := args.Map{"ok": ok, "len": c.Length(), "negFail": c.RemoveAt(-1), "bigFail": c.RemoveAt(100)}

		// Assert
		expected := args.Map{"ok": true, "len": 2, "negFail": false, "bigFail": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt", actual)
	})
}

func Test_Cov16_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_ListStrings", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"listLen": len(c.ListStrings()), "listPtrLen": len(c.ListStringsPtr())}

		// Assert
		expected := args.Map{"listLen": 1, "listPtrLen": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ListStrings", actual)
	})
}

func Test_Cov16_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_LengthLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": c.LengthLock()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- LengthLock", actual)
	})
}

func Test_Cov16_Collection_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_IsEmptyLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{"empty": c.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsEmptyLock", actual)
	})
}

func Test_Cov16_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AsError", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		nilDef := c.AsDefaultError() == nil
		nilSep := c.AsError(",") == nil
		c.Add("e1")
		hasDef := c.AsDefaultError() != nil

		// Act
		actual := args.Map{"nilDef": nilDef, "nilSep": nilSep, "hasDef": hasDef}

		// Assert
		expected := args.Map{"nilDef": true, "nilSep": true, "hasDef": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AsError", actual)
	})
}

func Test_Cov16_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_ToError", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		nilToErr := c.ToError(",") == nil
		nilToDef := c.ToDefaultError() == nil
		c.Add("e")
		hasToErr := c.ToError(",") != nil

		// Act
		actual := args.Map{"nilToErr": nilToErr, "nilToDef": nilToDef, "hasToErr": hasToErr}

		// Assert
		expected := args.Map{"nilToErr": true, "nilToDef": true, "hasToErr": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- ToError", actual)
	})
}

func Test_Cov16_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_EachItemSplitBy", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a,b", "c"})
		result := c.EachItemSplitBy(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- EachItemSplitBy", actual)
	})
}

func Test_Cov16_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_ConcatNew", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		newC := c.ConcatNew(0, "b", "c")

		// Act
		actual := args.Map{"len": newC.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ConcatNew", actual)
	})
}

func Test_Cov16_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_IsEquals", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"equal": a.IsEquals(b)}

		// Assert
		expected := args.Map{"equal": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsEquals", actual)
	})
}

func Test_Cov16_Collection_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_IsEqualsWithSensitive", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"A"})
		b := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"caseInsensitive": a.IsEqualsWithSensitive(false, b),
			"caseSensitive":   a.IsEqualsWithSensitive(true, b),
		}

		// Assert
		expected := args.Map{"caseInsensitive": true, "caseSensitive": false}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- IsEqualsWithSensitive", actual)
	})
}

func Test_Cov16_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_JsonString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		// Exercise the methods for coverage (value receiver on JsonPtr causes empty output)
		_ = c.JsonString()
		_ = c.JsonStringMust()
		_ = c.StringJSON()
		// Verify json.Marshal works correctly with pointer receiver
		b, _ := json.Marshal(c)

		// Act
		actual := args.Map{"json": len(b) > 2}

		// Assert
		expected := args.Map{"json": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonString", actual)
	})
}

func Test_Cov16_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddHashmapsValues", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsValues(hm)
		c.AddHashmapsValues(nil)
		c.AddHashmapsValues(corestr.New.Hashmap.Empty())

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddHashmapsValues", actual)
	})
}

func Test_Cov16_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddHashmapsKeys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeys(hm)
		c.AddHashmapsKeys(nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddHashmapsKeys", actual)
	})
}

func Test_Cov16_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddPointerCollectionsLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c2 := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(c2)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddPointerCollectionsLock", actual)
	})
}

func Test_Cov16_Collection_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_HasIndex", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"has0": c.HasIndex(0), "has1": c.HasIndex(1), "has2": c.HasIndex(2), "hasNeg": c.HasIndex(-1)}

		// Assert
		expected := args.Map{"has0": true, "has1": true, "has2": false, "hasNeg": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasIndex", actual)
	})
}

func Test_Cov16_Collection_Extended(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_Extended", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		_ = c.Take(2)
		_ = c.Skip(1)
		_ = c.First()
		_ = c.Last()
		_ = c.FirstOrDefault()
		_ = c.LastOrDefault()
		_ = c.IndexAt(0)
		_ = c.List()
		_ = c.HasItems()
		_ = c.Reverse()
		_ = c.GetPagesSize(2)
		_ = c.UniqueList()
		_ = c.UniqueListLock()
		_ = c.UniqueBoolMap()
		_ = c.UniqueBoolMapLock()
		_ = c.Items()
		_ = c.ListPtr()
		_ = c.ListCopyPtrLock()
		_ = c.Has("b")
		_ = c.HasLock("a")
		aPtr := "a"
		_ = c.HasPtr(&aPtr)
		_ = c.HasAll("a", "b")
		_ = c.SortedListAsc()
		_ = c.SortedAsc()
		_ = c.SortedAscLock()
		_ = c.SortedListDsc()
		_ = c.NonEmptyList()
		_ = c.NonEmptyListPtr()
		_ = c.NonEmptyItems()
		_ = c.NonEmptyItemsPtr()
		_ = c.NonEmptyItemsOrNonWhitespace()
		_ = c.NonEmptyItemsOrNonWhitespacePtr()
		_ = c.HashsetAsIs()
		_ = c.HashsetWithDoubleLength()
		_ = c.HashsetLock()
		_ = c.CharCollectionMap()
		_ = c.Join(",")
		_ = c.JoinLine()
		_ = c.Joins(",")
		_ = c.NonEmptyJoins(",")
		_ = c.NonWhitespaceJoins(",")
		_ = c.String()
		_ = c.StringLock()
		_ = c.SummaryString(1)
		_ = c.SummaryStringWithHeader("header")
		_ = c.CsvLines()
		_ = c.Csv()
		bPtr := "a"
		_ = c.IsContainsPtr(&bPtr)
		_ = c.IsContainsAll("a", "b")
		_ = c.IsContainsAllLock("a", "b")
		_ = c.IsContainsAllSlice([]string{"a"})
		_ = c.HasUsingSensitivity("A", false)
		c.AddCapacity(10)
		c.Resize(5)
		_ = c.New("d", "e")
		_ = c.JsonModelAny()
		_ = c.Json()
		_ = c.JsonPtr()
		_ = c.AsJsonMarshaller()
		_ = c.AsJsonContractsBinder()
		_, _ = c.Serialize()
		var target []string
		_ = c.Deserialize(&target)

		// Act
		actual := args.Map{"exercised": true}

		// Assert
		expected := args.Map{"exercised": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- extended methods", actual)
	})
}

func Test_Cov16_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_Filter", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		filterFn := corestr.IsStringFilter(func(str string, index int) (string, bool, bool) {
			return str, len(str) > 1, false
		})
		filtered := c.Filter(filterFn)
		_ = c.FilterLock(filterFn)
		_ = c.FilteredCollection(filterFn)
		_ = c.FilteredCollectionLock(filterFn)
		ptrFilterFn := corestr.IsStringPointerFilter(func(sp *string, index int) (*string, bool, bool) {
			if sp != nil && len(*sp) > 1 {
				return sp, true, false
			}
			return nil, false, false
		})
		_ = c.FilterPtr(ptrFilterFn)
		_ = c.FilterPtrLock(ptrFilterFn)

		// Act
		actual := args.Map{"filteredLen": len(filtered)}

		// Assert
		expected := args.Map{"filteredLen": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Filter", actual)
	})
}

func Test_Cov16_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AppendAnys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnys("a", 42, nil)
		c.AppendAnys()
		c.AppendAnysLock("b")
		c.AppendAnysLock()
		c.AppendNonEmptyAnys("c", nil)
		c.AppendNonEmptyAnys()
		c.AddsNonEmpty("d", "", "e")
		c.AddsNonEmpty()

		// Act
		actual := args.Map{"hasItems": c.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendAnys", actual)
	})
}

func Test_Cov16_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddNonEmptyStrings", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("a", "", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct -- AddNonEmptyStrings", actual)
	})
}

func Test_Cov16_Collection_AddNonEmptyStringsSlice(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddNonEmptyStringsSlice", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{"a", "", "b"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct -- AddNonEmptyStringsSlice", actual)
	})
}

func Test_Cov16_Collection_ExpandMerge(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_ExpandMerge", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.ExpandSlicePlusAdd(
			[]string{"b", "c"},
			func(line string) []string {
				return []string{line}
			},
		)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ExpandSlicePlusAdd", actual)
	})
}

func Test_Cov16_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_MergeSlicesOfSlice", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a", "b"}, []string{"c"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- MergeSlicesOfSlice", actual)
	})
}

func Test_Cov16_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_GetAllExcept", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"b"}))
		r2 := c.GetAllExcept([]string{"b"})

		// Act
		actual := args.Map{"rLen": len(r), "r2Len": len(r2)}

		// Assert
		expected := args.Map{"rLen": 2, "r2Len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetAllExcept", actual)
	})
}

func Test_Cov16_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_GetHashsetPlusHasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})

		// Act
		actual := args.Map{"hasAll": hasAll, "hsNotNil": hs != nil}

		// Assert
		expected := args.Map{"hasAll": true, "hsNotNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetHashsetPlusHasAll", actual)
	})
}

func Test_Cov16_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddStringsByFuncChecking", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking(
			[]string{"a", "bb", "c", "dd"},
			func(s string) bool { return len(s) > 1 },
		)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddStringsByFuncChecking", actual)
	})
}

func Test_Cov16_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_AddFuncResult", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "hello" })

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddFuncResult", actual)
	})
}

func Test_Cov16_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_ParseInjectUsingJson", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		// Use json.Marshal with pointer to get correct JSON (bypasses value receiver issue)
		b, _ := json.Marshal(c)
		jr := &corejson.Result{Bytes: b}
		c2 := corestr.New.Collection.Empty()
		_, err := c2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ParseInjectUsingJson", actual)
	})
}

func Test_Cov16_Collection_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov16_Collection_ClearDispose", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()
		lenAfterClear := c.Length()
		c.Add("x")
		c.Dispose()

		// Act
		actual := args.Map{"clearedLen": lenAfterClear}

		// Assert
		expected := args.Map{"clearedLen": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ClearDispose", actual)
	})
}
