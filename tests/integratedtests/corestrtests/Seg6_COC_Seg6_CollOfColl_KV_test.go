package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CollectionsOfCollection — Segment 6c
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg6_COC_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_IsEmpty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 0)

		// Act
		actual := args.Map{
			"empty": coc.IsEmpty(),
			"hasItems": coc.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Seg6_COC_Add(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Add", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(c)

		// Act
		actual := args.Map{
			"len": coc.Length(),
			"allLen": coc.AllIndividualItemsLength(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"allLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Add -- 1 collection 2 items", actual)
	})
}

func Test_Seg6_COC_Add_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Add_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Cap(0)
		coc.Add(c)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Add empty collection -- skipped", actual)
	})
}

func Test_Seg6_COC_Adds(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Adds", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c1 := *corestr.New.Collection.Strings([]string{"a"})
		c2 := *corestr.New.Collection.Strings([]string{"b"})
		coc.Adds(c1, c2)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Adds -- 2 collections", actual)
	})
}

func Test_Seg6_COC_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Adds_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.Adds(nil...)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds nil -- no change", actual)
	})
}

func Test_Seg6_COC_AddCollections(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddCollections", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := *corestr.New.Collection.Strings([]string{"a"})
		coc.AddCollections(c)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollections -- 1 collection", actual)
	})
}

func Test_Seg6_COC_AddCollections_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddCollections_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.AddCollections(nil...)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollections nil -- no change", actual)
	})
}

func Test_Seg6_COC_AddStrings(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.AddStrings(false, []string{"a", "b"})

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 1 collection", actual)
	})
}

func Test_Seg6_COC_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddStrings_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.AddStrings(false, []string{})

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty -- no change", actual)
	})
}

func Test_Seg6_COC_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddsStringsOfStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 4)
		coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsStringsOfStrings -- 2 collections", actual)
	})
}

func Test_Seg6_COC_AddsStringsOfStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddsStringsOfStrings_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.AddsStringsOfStrings(false, nil...)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsStringsOfStrings nil -- no change", actual)
	})
}

func Test_Seg6_COC_Items(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Items", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)

		// Act
		actual := args.Map{"len": len(coc.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items -- 1 collection", actual)
	})
}

func Test_Seg6_COC_List(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_List", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 4)
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c"})
		coc.Add(c1).Add(c2)

		// Act
		actual := args.Map{"len": len(coc.List(0))}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "List -- 3 items total", actual)
	})
}

func Test_Seg6_COC_List_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_List_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 0)

		// Act
		actual := args.Map{"len": len(coc.List(0))}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "List empty -- 0", actual)
	})
}

func Test_Seg6_COC_ToCollection(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_ToCollection", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(c)
		result := coc.ToCollection()

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ToCollection -- 2 items", actual)
	})
}

func Test_Seg6_COC_AllIndividualItemsLength_WithEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AllIndividualItemsLength_WithEmpty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 4)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Cap(0)
		coc.Add(c1).Add(c2)

		// Act
		actual := args.Map{"len": coc.AllIndividualItemsLength()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AllIndividualItemsLength -- skips empty", actual)
	})
}

func Test_Seg6_COC_String(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_String", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)

		// Act
		actual := args.Map{"nonEmpty": coc.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_Seg6_COC_Json(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Json", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		j := coc.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg6_COC_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_MarshalJSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		b, err := coc.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg6_COC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_UnmarshalJSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		b, _ := coc.MarshalJSON()
		coc2 := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		err := coc2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg6_COC_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_UnmarshalJSON_Invalid", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		err := coc.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg6_COC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_JsonModel", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)

		// Act
		actual := args.Map{"len": len(coc.JsonModel().Items)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1 item", actual)
	})
}

func Test_Seg6_COC_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_JsonModelAny", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)

		// Act
		actual := args.Map{"notNil": coc.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg6_COC_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_InterfaceCasts", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)

		// Act
		actual := args.Map{
			"jsoner":   coc.AsJsoner() != nil,
			"binder":   coc.AsJsonContractsBinder() != nil,
			"injector": coc.AsJsonParseSelfInjector() != nil,
			"marsh":    coc.AsJsonMarshaller() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner": true,
			"binder": true,
			"injector": true,
			"marsh": true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Seg6_COC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_ParseInjectUsingJson", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		_, err := coc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg6_COC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_ParseInjectUsingJsonMust", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		result := coc2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg6_COC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_JsonParseSelfInject", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		err := coc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValueCollection — Segment 6d
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg6_KVC_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_IsEmpty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{
			"empty": kvc.IsEmpty(),
			"hasAny": kvc.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Seg6_KVC_Add(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Add", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{
			"len": kvc.Length(),
			"count": kvc.Count(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"count": 2,
		}
		expected.ShouldBeEqual(t, 0, "Add -- 2 pairs", actual)
	})
}

func Test_Seg6_KVC_AddIf(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddIf", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(true, "a", "1").AddIf(false, "b", "2")

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddIf -- only true", actual)
	})
}

func Test_Seg6_KVC_Adds(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Adds", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Adds -- 2 pairs", actual)
	})
}

func Test_Seg6_KVC_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Adds_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds()

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds empty -- no change", actual)
	})
}

func Test_Seg6_KVC_AddMap(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddMap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddMap -- 1 pair", actual)
	})
}

func Test_Seg6_KVC_AddMap_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddMap_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddMap nil -- no change", actual)
	})
}

func Test_Seg6_KVC_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddHashsetMap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(map[string]bool{"a": true})

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashsetMap -- 1 pair", actual)
	})
}

func Test_Seg6_KVC_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddHashsetMap_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetMap nil -- no change", actual)
	})
}

func Test_Seg6_KVC_AddHashset(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddHashset", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		hs := corestr.New.Hashset.Strings([]string{"a"})
		kvc.AddHashset(hs)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashset -- 1 pair", actual)
	})
}

func Test_Seg6_KVC_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddHashset_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashset(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashset nil -- no change", actual)
	})
}

func Test_Seg6_KVC_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddsHashmap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		hm := corestr.New.Hashmap.Cap(2)
		hm.Set("a", "1")
		kvc.AddsHashmap(hm)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsHashmap -- 1 pair", actual)
	})
}

func Test_Seg6_KVC_AddsHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddsHashmap_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmap(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsHashmap nil -- no change", actual)
	})
}

func Test_Seg6_KVC_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddsHashmaps", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		hm1 := corestr.New.Hashmap.Cap(2)
		hm1.Set("a", "1")
		hm2 := corestr.New.Hashmap.Cap(2)
		hm2.Set("b", "2")
		kvc.AddsHashmaps(hm1, hm2)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsHashmaps -- 2 pairs", actual)
	})
}

func Test_Seg6_KVC_AddsHashmaps_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddsHashmaps_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmaps(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsHashmaps nil -- no change", actual)
	})
}

func Test_Seg6_KVC_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddStringBySplit", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplit("=", "key=value")
		val, found := kvc.Get("key")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "value",
		}
		expected.ShouldBeEqual(t, 0, "AddStringBySplit -- key=value", actual)
	})
}

func Test_Seg6_KVC_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddStringBySplitTrim", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplitTrim("=", " key = value ")
		val, found := kvc.Get("key")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "value",
		}
		expected.ShouldBeEqual(t, 0, "AddStringBySplitTrim -- trimmed", actual)
	})
}

// ── Accessors ───────────────────────────────────────────────────────────────

func Test_Seg6_KVC_First_Last(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_First_Last", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{
			"first": kvc.First().Key,
			"last": kvc.Last().Key,
			"lastIdx": kvc.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "b",
			"lastIdx": 1,
		}
		expected.ShouldBeEqual(t, 0, "First/Last -- correct", actual)
	})
}

func Test_Seg6_KVC_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_FirstOrDefault_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"nil": kvc.FirstOrDefault() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault empty -- nil", actual)
	})
}

func Test_Seg6_KVC_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_LastOrDefault_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"nil": kvc.LastOrDefault() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "LastOrDefault empty -- nil", actual)
	})
}

func Test_Seg6_KVC_HasKey(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_HasKey", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{
			"has": kvc.HasKey("a"),
			"miss": kvc.HasKey("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasKey -- found and missing", actual)
	})
}

func Test_Seg6_KVC_IsContains(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_IsContains", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{
			"has": kvc.IsContains("a"),
			"miss": kvc.IsContains("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContains -- found and missing", actual)
	})
}

func Test_Seg6_KVC_HasIndex(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_HasIndex", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{
			"has": kvc.HasIndex(0),
			"miss": kvc.HasIndex(5),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasIndex -- found and missing", actual)
	})
}

func Test_Seg6_KVC_Get(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Get", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		val, found := kvc.Get("a")
		val2, found2 := kvc.Get("z")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
			"val2": val2,
			"found2": found2,
		}

		// Assert
		expected := args.Map{
			"val": "1",
			"found": true,
			"val2": "",
			"found2": false,
		}
		expected.ShouldBeEqual(t, 0, "Get -- found and missing", actual)
	})
}

func Test_Seg6_KVC_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SafeValueAt", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{
			"val": kvc.SafeValueAt(0),
			"miss": kvc.SafeValueAt(5),
		}

		// Assert
		expected := args.Map{
			"val": "1",
			"miss": "",
		}
		expected.ShouldBeEqual(t, 0, "SafeValueAt -- found and out of bounds", actual)
	})
}

func Test_Seg6_KVC_SafeValueAt_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SafeValueAt_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"val": kvc.SafeValueAt(0)}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "SafeValueAt empty -- empty string", actual)
	})
}

func Test_Seg6_KVC_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SafeValuesAtIndexes", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")
		result := kvc.SafeValuesAtIndexes(0, 1, 5)

		// Act
		actual := args.Map{
			"len": len(result),
			"first": result[0],
			"last": result[2],
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"first": "1",
			"last": "",
		}
		expected.ShouldBeEqual(t, 0, "SafeValuesAtIndexes -- values", actual)
	})
}

func Test_Seg6_KVC_SafeValuesAtIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SafeValuesAtIndexes_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		result := kvc.SafeValuesAtIndexes()

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeValuesAtIndexes empty -- 0", actual)
	})
}

func Test_Seg6_KVC_Find(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Find", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")
		result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Find -- 1 match", actual)
	})
}

func Test_Seg6_KVC_Find_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Find_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Find empty -- 0", actual)
	})
}

func Test_Seg6_KVC_Find_Break(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Find_Break", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2").Add("c", "3")
		result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, true
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Find break -- 1 item", actual)
	})
}

// ── Keys / Values ───────────────────────────────────────────────────────────

func Test_Seg6_KVC_AllKeys(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllKeys", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"len": len(kvc.AllKeys())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AllKeys -- 2 keys", actual)
	})
}

func Test_Seg6_KVC_AllKeys_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllKeys_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"len": len(kvc.AllKeys())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllKeys empty -- 0", actual)
	})
}

func Test_Seg6_KVC_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllKeysSorted", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("b", "2").Add("a", "1")
		sorted := kvc.AllKeysSorted()

		// Act
		actual := args.Map{"first": sorted[0]}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "AllKeysSorted -- sorted", actual)
	})
}

func Test_Seg6_KVC_AllValues(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllValues", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"len": len(kvc.AllValues())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AllValues -- 2 values", actual)
	})
}

func Test_Seg6_KVC_AllValues_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllValues_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"len": len(kvc.AllValues())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllValues empty -- 0", actual)
	})
}

// ── String / Join ───────────────────────────────────────────────────────────

func Test_Seg6_KVC_Strings(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Strings", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"len": len(kvc.Strings())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Strings -- 1 string", actual)
	})
}

func Test_Seg6_KVC_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Strings_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"len": len(kvc.Strings())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Strings empty -- 0", actual)
	})
}

func Test_Seg6_KVC_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_StringsUsingFormat", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		result := kvc.StringsUsingFormat("%v=%v")

		// Act
		actual := args.Map{"val": result[0]}

		// Assert
		expected := args.Map{"val": "a=1"}
		expected.ShouldBeEqual(t, 0, "StringsUsingFormat -- formatted", actual)
	})
}

func Test_Seg6_KVC_StringsUsingFormat_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_StringsUsingFormat_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		result := kvc.StringsUsingFormat("%v=%v")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "StringsUsingFormat empty -- 0", actual)
	})
}

func Test_Seg6_KVC_String(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_String", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"nonEmpty": kvc.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg6_KVC_Compile(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Compile", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"nonEmpty": kvc.Compile() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Compile -- delegates to String", actual)
	})
}

func Test_Seg6_KVC_Join(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Join", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"nonEmpty": kvc.Join(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Join -- non-empty", actual)
	})
}

func Test_Seg6_KVC_JoinKeys(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JoinKeys", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"nonEmpty": kvc.JoinKeys(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinKeys -- non-empty", actual)
	})
}

func Test_Seg6_KVC_JoinValues(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JoinValues", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"nonEmpty": kvc.JoinValues(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinValues -- non-empty", actual)
	})
}

// ── Hashmap / Map ───────────────────────────────────────────────────────────

func Test_Seg6_KVC_Hashmap(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Hashmap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		hm := kvc.Hashmap()

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap -- 1 item", actual)
	})
}

func Test_Seg6_KVC_Hashmap_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Hashmap_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		hm := kvc.Hashmap()

		// Act
		actual := args.Map{"empty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap empty -- empty", actual)
	})
}

func Test_Seg6_KVC_Map(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Map", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"len": len(kvc.Map())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Map -- 1 item", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_Seg6_KVC_Json(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Json", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		j := kvc.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg6_KVC_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_MarshalJSON", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, err := kvc.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg6_KVC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_UnmarshalJSON", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, _ := kvc.MarshalJSON()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.UnmarshalJSON(b)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": kvc2.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg6_KVC_UnmarshalJSON_EmptyArray(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_UnmarshalJSON_EmptyArray", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		err := kvc.UnmarshalJSON([]byte(`[]`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": kvc.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON empty array -- empty", actual)
	})
}

func Test_Seg6_KVC_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Serialize", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, err := kvc.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg6_KVC_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SerializeMust", func() {
		// Arrange
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b := kvc.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "SerializeMust -- success", actual)
	})
}

func Test_Seg6_KVC_Deserialize(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Deserialize", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		var dest interface{}
		err := kvc.Deserialize(&dest)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_Seg6_KVC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JsonModel", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"len": len(kvc.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1 item", actual)
	})
}

func Test_Seg6_KVC_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JsonModelAny", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"notNil": kvc.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg6_KVC_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_InterfaceCasts", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{
			"jsoner":   kvc.AsJsoner() != nil,
			"binder":   kvc.AsJsonContractsBinder() != nil,
			"injector": kvc.AsJsonParseSelfInjector() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner": true,
			"binder": true,
			"injector": true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Seg6_KVC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_ParseInjectUsingJson", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jr := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		_, err := kvc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg6_KVC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JsonParseSelfInject", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jr := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── Clear / Dispose ─────────────────────────────────────────────────────────

func Test_Seg6_KVC_Clear(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Clear", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Clear()

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg6_KVC_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Dispose", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Dispose()

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyValuePair — Segment 6e
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg6_KAVP_Basic(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Basic", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "name", Value: "hello"}

		// Act
		actual := args.Map{
			"key":     kv.KeyName(),
			"varName": kv.VariableName(),
			"valAny":  kv.ValueAny() != nil,
			"isEqual": kv.IsVariableNameEqual("name"),
			"notEq":   kv.IsVariableNameEqual("other"),
		}

		// Assert
		expected := args.Map{
			"key": "name",
			"varName": "name",
			"valAny": true,
			"isEqual": true,
			"notEq": false,
		}
		expected.ShouldBeEqual(t, 0, "Basic accessors -- correct", actual)
	})
}

func Test_Seg6_KAVP_ValueChecks(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ValueChecks", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "hello"}
		kvNil := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act
		actual := args.Map{
			"hasVal":        kv.HasValue(),
			"hasNonNull":    kv.HasNonNull(),
			"isNull":        kv.IsValueNull(),
			"nilIsNull":     kvNil.IsValueNull(),
			"nilHasVal":     kvNil.HasValue(),
			"emptyStr":      kv.IsValueEmptyString(),
			"whitespace":    kv.IsValueWhitespace(),
			"nilEmptyStr":   kvNil.IsValueEmptyString(),
			"nilWhitespace": kvNil.IsValueWhitespace(),
		}

		// Assert
		expected := args.Map{
			"hasVal": true, "hasNonNull": true, "isNull": false,
			"nilIsNull": true, "nilHasVal": false,
			"emptyStr": false, "whitespace": false,
			"nilEmptyStr": true, "nilWhitespace": true,
		}
		expected.ShouldBeEqual(t, 0, "Value checks -- correct", actual)
	})
}

func Test_Seg6_KAVP_ValueString(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ValueString", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{"val": kv.ValueString()}

		// Assert
		expected := args.Map{"val": "42"}
		expected.ShouldBeEqual(t, 0, "ValueString -- formatted", actual)
	})
}

func Test_Seg6_KAVP_ValueString_Cached(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ValueString_Cached", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		_ = kv.ValueString() // init

		// Act
		actual := args.Map{"val": kv.ValueString()}

		// Assert
		expected := args.Map{"val": "42"}
		expected.ShouldBeEqual(t, 0, "ValueString cached -- same value", actual)
	})
}

func Test_Seg6_KAVP_String(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_String", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kv.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg6_KAVP_Compile(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Compile", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kv.Compile() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Compile -- delegates to String", actual)
	})
}

func Test_Seg6_KAVP_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_SerializeMust", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "SerializeMust -- success", actual)
	})
}

func Test_Seg6_KAVP_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Serialize", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b, err := kv.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg6_KAVP_Json(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Json", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kv.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg6_KAVP_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_InterfaceCasts", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"jsoner":   kv.AsJsoner() != nil,
			"binder":   kv.AsJsonContractsBinder() != nil,
			"injector": kv.AsJsonParseSelfInjector() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner": true,
			"binder": true,
			"injector": true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Seg6_KAVP_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ParseInjectUsingJson", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		_, err := kv2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg6_KAVP_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ParseInjectUsingJsonMust", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		result := kv2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg6_KAVP_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_JsonParseSelfInject", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		err := kv2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_Seg6_KAVP_Clear(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Clear", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Clear()

		// Act
		actual := args.Map{
			"key": kv.Key,
			"isNull": kv.IsValueNull(),
		}

		// Assert
		expected := args.Map{
			"key": "",
			"isNull": true,
		}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg6_KAVP_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Dispose", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Dispose()

		// Act
		actual := args.Map{"key": kv.Key}

		// Assert
		expected := args.Map{"key": ""}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_Seg6_KAVP_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Clear_Nil", func() {
		var kv *corestr.KeyAnyValuePair
		kv.Clear() // should not panic
	})
}

func Test_Seg6_KAVP_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Dispose_Nil", func() {
		var kv *corestr.KeyAnyValuePair
		kv.Dispose() // should not panic
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair — Segment 6f
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg6_KVP_Basic(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Basic", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "name", Value: "hello"}

		// Act
		actual := args.Map{
			"key":      kv.KeyName(),
			"varName":  kv.VariableName(),
			"valStr":   kv.ValueString(),
			"isEqual":  kv.IsVariableNameEqual("name"),
			"isValEq":  kv.IsValueEqual("hello"),
			"isValNeq": kv.IsValueEqual("other"),
		}

		// Assert
		expected := args.Map{
			"key": "name", "varName": "name", "valStr": "hello",
			"isEqual": true, "isValEq": true, "isValNeq": false,
		}
		expected.ShouldBeEqual(t, 0, "Basic accessors -- correct", actual)
	})
}

func Test_Seg6_KVP_Empty_Checks(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Empty_Checks", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kvEmpty := &corestr.KeyValuePair{}

		// Act
		actual := args.Map{
			"keyEmpty":     kv.IsKeyEmpty(),
			"valEmpty":     kv.IsValueEmpty(),
			"hasKey":       kv.HasKey(),
			"hasVal":       kv.HasValue(),
			"kvEmpty":      kv.IsKeyValueEmpty(),
			"emptyKeyE":    kvEmpty.IsKeyEmpty(),
			"emptyValE":    kvEmpty.IsValueEmpty(),
			"emptyKVE":     kvEmpty.IsKeyValueEmpty(),
			"anyEmpty":     kv.IsKeyValueAnyEmpty(),
			"anyEmptyFull": kvEmpty.IsKeyValueAnyEmpty(),
		}

		// Assert
		expected := args.Map{
			"keyEmpty": false, "valEmpty": false,
			"hasKey": true, "hasVal": true,
			"kvEmpty": false, "emptyKeyE": true,
			"emptyValE": true, "emptyKVE": true,
			"anyEmpty": false, "anyEmptyFull": true,
		}
		expected.ShouldBeEqual(t, 0, "Empty checks -- correct", actual)
	})
}

func Test_Seg6_KVP_Is_IsKey_IsVal(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Is_IsKey_IsVal", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"is":    kv.Is("k", "v"),
			"isNot": kv.Is("k", "x"),
			"isKey": kv.IsKey("k"),
			"isVal": kv.IsVal("v"),
		}

		// Assert
		expected := args.Map{
			"is": true,
			"isNot": false,
			"isKey": true,
			"isVal": true,
		}
		expected.ShouldBeEqual(t, 0, "Is/IsKey/IsVal -- correct", actual)
	})
}

func Test_Seg6_KVP_Trim(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Trim", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: " k ", Value: " v "}

		// Act
		actual := args.Map{
			"trimKey": kv.TrimKey(),
			"trimVal": kv.TrimValue(),
		}

		// Assert
		expected := args.Map{
			"trimKey": "k",
			"trimVal": "v",
		}
		expected.ShouldBeEqual(t, 0, "Trim -- trimmed", actual)
	})
}

func Test_Seg6_KVP_ValueBool(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueBool", func() {
		// Arrange
		kvTrue := &corestr.KeyValuePair{Key: "k", Value: "true"}
		kvFalse := &corestr.KeyValuePair{Key: "k", Value: "invalid"}
		kvEmpty := &corestr.KeyValuePair{Key: "k", Value: ""}

		// Act
		actual := args.Map{
			"true": kvTrue.ValueBool(),
			"invalid": kvFalse.ValueBool(),
			"empty": kvEmpty.ValueBool(),
		}

		// Assert
		expected := args.Map{
			"true": true,
			"invalid": false,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "ValueBool -- various", actual)
	})
}

func Test_Seg6_KVP_ValueInt(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueInt", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "42"}
		kvBad := &corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act
		actual := args.Map{
			"val": kv.ValueInt(0),
			"def": kvBad.ValueInt(99),
			"defInt": kv.ValueDefInt(),
			"badDef": kvBad.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"val": 42,
			"def": 99,
			"defInt": 42,
			"badDef": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValueInt -- various", actual)
	})
}

func Test_Seg6_KVP_ValueByte(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueByte", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "65"}
		kvBad := &corestr.KeyValuePair{Key: "k", Value: "abc"}
		kvBig := &corestr.KeyValuePair{Key: "k", Value: "999"}

		// Act
		actual := args.Map{
			"val":    kv.ValueByte(0),
			"bad":    kvBad.ValueByte(1),
			"big":    kvBig.ValueByte(2),
			"defVal": kv.ValueDefByte(),
			"defBad": kvBad.ValueDefByte(),
			"defBig": kvBig.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"val": byte(65),
			"bad": byte(1),
			"big": byte(2),
			"defVal": byte(65),
			"defBad": byte(0),
			"defBig": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte -- various", actual)
	})
}

func Test_Seg6_KVP_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueFloat64", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "1.5"}
		kvBad := &corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act
		actual := args.Map{
			"val": kv.ValueFloat64(0),
			"bad": kvBad.ValueFloat64(9.9),
			"def": kv.ValueDefFloat64(),
		}

		// Assert
		expected := args.Map{
			"val": 1.5,
			"bad": 9.9,
			"def": 1.5,
		}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 -- various", actual)
	})
}

func Test_Seg6_KVP_ValueValid(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueValid", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "hello"}
		vv := kv.ValueValid()

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
		expected.ShouldBeEqual(t, 0, "ValueValid -- valid", actual)
	})
}

func Test_Seg6_KVP_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueValidOptions", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "hello"}
		vv := kv.ValueValidOptions(false, "err")

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
			"msg": vv.Message,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "ValueValidOptions -- custom", actual)
	})
}

func Test_Seg6_KVP_FormatString(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_FormatString", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"val": kv.FormatString("%v=%v")}

		// Assert
		expected := args.Map{"val": "k=v"}
		expected.ShouldBeEqual(t, 0, "FormatString -- formatted", actual)
	})
}

func Test_Seg6_KVP_String(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_String", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kv.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg6_KVP_Compile(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Compile", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kv.Compile() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Compile -- delegates to String", actual)
	})
}

func Test_Seg6_KVP_Json(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Json", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		j := kv.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg6_KVP_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Serialize", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		b, err := kv.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg6_KVP_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_SerializeMust", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "SerializeMust -- success", actual)
	})
}

func Test_Seg6_KVP_Clear(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Clear", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()

		// Act
		actual := args.Map{
			"key": kv.Key,
			"val": kv.Value,
		}

		// Assert
		expected := args.Map{
			"key": "",
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg6_KVP_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Dispose", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Dispose()

		// Act
		actual := args.Map{"key": kv.Key}

		// Assert
		expected := args.Map{"key": ""}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_Seg6_KVP_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Clear_Nil", func() {
		var kv *corestr.KeyValuePair
		kv.Clear() // should not panic
	})
}

func Test_Seg6_KVP_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Dispose_Nil", func() {
		var kv *corestr.KeyValuePair
		kv.Dispose() // should not panic
	})
}
