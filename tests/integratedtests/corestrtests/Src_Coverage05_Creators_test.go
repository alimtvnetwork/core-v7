package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC05_CollectionCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_CollectionCreator_Verification", func() {
		// Arrange
		tc := srcC05CollectionCreatorTestCase

		// Act
		c := corestr.New.Collection.Empty()
		c2 := corestr.New.Collection.Cap(10)
		c3 := corestr.New.Collection.Strings([]string{"a", "b"})
		c4 := corestr.New.Collection.Create([]string{"a"})
		c5 := corestr.New.Collection.CloneStrings([]string{"a"})
		c6 := corestr.New.Collection.StringsOptions(true, []string{"a"})
		c7 := corestr.New.Collection.StringsOptions(false, []string{"a"})
		c8 := corestr.New.Collection.StringsOptions(false, nil)
		c9 := corestr.New.Collection.LineUsingSep(",", "a,b")
		_ = corestr.New.Collection.LineDefault("a\nb")
		c11 := corestr.New.Collection.StringsPlusCap(5, []string{"a"})
		_ = corestr.New.Collection.StringsPlusCap(0, []string{"a"})
		_ = corestr.New.Collection.CapStrings(5, []string{"a"})
		_ = corestr.New.Collection.CapStrings(0, []string{"a"})
		c15 := corestr.New.Collection.LenCap(2, 5)

		actual := args.Map{
			"emptyIsEmpty":  c.IsEmpty(),
			"capHasCap":     c2.Capacity() >= 10,
			"stringsLen":    c3.Length(),
			"createLen":     c4.Length(),
			"cloneLen":      c5.Length(),
			"optCloneLen":   c6.Length(),
			"optNoCloneLen": c7.Length(),
			"optNilEmpty":   c8.IsEmpty(),
			"sepLen":        c9.Length(),
			"plusCapLen":    c11.Length(),
			"lenCapLen":     c15.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC05_SimpleStringOnceCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_SimpleStringOnceCreator_Verification", func() {
		// Arrange
		tc := srcC05SimpleStringOnceCreatorTestCase

		// Act
		s := corestr.New.SimpleStringOnce.Init("hello")
		s2 := corestr.New.SimpleStringOnce.InitPtr("hello")
		s3 := corestr.New.SimpleStringOnce.Uninitialized("val")
		_ = corestr.New.SimpleStringOnce.Create("val", true)
		_ = corestr.New.SimpleStringOnce.CreatePtr("val", false)
		_ = corestr.New.SimpleStringOnce.Empty()
		_ = corestr.New.SimpleStringOnce.Any(false, "hello", true)

		actual := args.Map{
			"initValue":         s.Value(),
			"initIsInitialized": s.IsInitialized(),
			"initPtrValue":      s2.Value(),
			"uninitIsInit":      s3.IsInitialized(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC05_HashsetCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_HashsetCreator_Verification", func() {
		// Arrange
		tc := srcC05HashsetCreatorTestCase

		// Act
		h := corestr.New.Hashset.Empty()
		h3 := corestr.New.Hashset.Strings([]string{"a", "b"})
		noPanic := !callPanicsSrcC05(func() {
			_ = corestr.New.Hashset.Cap(10)
			_ = corestr.New.Hashset.Strings(nil)
			_ = corestr.New.Hashset.StringsSpreadItems("a")
			_ = corestr.New.Hashset.StringsSpreadItems()
			_ = corestr.New.Hashset.UsingMap(map[string]bool{"a": true})
			_ = corestr.New.Hashset.UsingMap(nil)
			_ = corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{"a": true})
			_ = corestr.New.Hashset.UsingMapOption(5, false, map[string]bool{"a": true})
			_ = corestr.New.Hashset.UsingMapOption(5, false, nil)
			_ = corestr.New.Hashset.StringsOption(5, true, "a")
			_ = corestr.New.Hashset.StringsOption(0, false)
			_ = corestr.New.Hashset.StringsOption(5, false)
			_ = corestr.New.Hashset.UsingCollection(corestr.New.Collection.Strings([]string{"a"}))
			_ = corestr.New.Hashset.UsingCollection(nil)
			ss := corestr.New.SimpleSlice.Lines("a")
			_ = corestr.New.Hashset.SimpleSlice(ss)
			emptySlice := corestr.New.SimpleSlice.Empty()
			_ = corestr.New.Hashset.SimpleSlice(emptySlice)
		})

		actual := args.Map{
			"emptyIsEmpty": h.IsEmpty(),
			"stringsLen":   h3.Length(),
			"noPanic":      noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC05_HashmapCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_HashmapCreator_Verification", func() {
		// Arrange
		tc := srcC05HashmapCreatorTestCase

		// Act
		h := corestr.New.Hashmap.Empty()
		h3 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "b"})
		noPanic := !callPanicsSrcC05(func() {
			_ = corestr.New.Hashmap.Cap(10)
			_ = corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"a": "b"})
			_ = corestr.New.Hashmap.UsingMapOptions(false, 5, map[string]string{"a": "b"})
			_ = corestr.New.Hashmap.UsingMapOptions(false, 5, nil)
			_ = corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "b"})
			_ = corestr.New.Hashmap.MapWithCap(0, map[string]string{"a": "b"})
			_ = corestr.New.Hashmap.MapWithCap(5, nil)
			_ = corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
			_ = corestr.New.Hashmap.KeyValues()
			_ = corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
			_ = corestr.New.Hashmap.KeyAnyValues()
			_ = corestr.New.Hashmap.KeyValuesStrings([]string{"k"}, []string{"v"})
			_ = corestr.New.Hashmap.KeyValuesStrings(nil, nil)
			keys := corestr.New.Collection.Strings([]string{"k"})
			vals := corestr.New.Collection.Strings([]string{"v"})
			_ = corestr.New.Hashmap.KeyValuesCollection(keys, vals)
			_ = corestr.New.Hashmap.KeyValuesCollection(nil, nil)
		})

		actual := args.Map{
			"emptyIsEmpty":  h.IsEmpty(),
			"usingMapEmpty": h3.IsEmpty(),
			"noPanic":       noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC05_LinkedListCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_LinkedListCreator_Verification", func() {
		// Arrange
		tc := srcC05LinkedListCreatorTestCase

		// Act
		ll := corestr.New.LinkedList.Create()
		ll3 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		noPanic := !callPanicsSrcC05(func() {
			_ = corestr.New.LinkedList.Empty()
			_ = corestr.New.LinkedList.Strings(nil)
			_ = corestr.New.LinkedList.SpreadStrings("a")
			_ = corestr.New.LinkedList.SpreadStrings()
			_ = corestr.New.LinkedList.UsingMap(map[string]bool{"a": true})
			_ = corestr.New.LinkedList.UsingMap(nil)
		})

		actual := args.Map{
			"createLen":  ll.Length(),
			"stringsLen": ll3.Length(),
			"noPanic":    noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC05_LinkedCollectionCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_LinkedCollectionCreator_Verification", func() {
		// Arrange
		tc := srcC05LinkedCollectionCreatorTestCase

		// Act
		lc := corestr.New.LinkedCollection.Create()
		lc3 := corestr.New.LinkedCollection.Strings("a", "b")
		noPanic := !callPanicsSrcC05(func() {
			_ = corestr.New.LinkedCollection.Empty()
			_ = corestr.New.LinkedCollection.Strings()
			_ = corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}))
			_ = corestr.New.LinkedCollection.UsingCollections()
		})

		actual := args.Map{
			"createLen":  lc.Length(),
			"stringsLen": lc3.Length(),
			"noPanic":    noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC05_KeyValuesCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_KeyValuesCreator_Verification", func() {
		// Arrange
		tc := srcC05KeyValuesCreatorTestCase

		// Act
		kv := corestr.New.KeyValues.Empty()
		kv3 := corestr.New.KeyValues.UsingMap(map[string]string{"k": "v"})
		noPanic := !callPanicsSrcC05(func() {
			_ = corestr.New.KeyValues.Cap(10)
			_ = corestr.New.KeyValues.UsingMap(nil)
			_ = corestr.New.KeyValues.UsingKeyValuePairs(corestr.KeyValuePair{Key: "k", Value: "v"})
			_ = corestr.New.KeyValues.UsingKeyValuePairs()
			_ = corestr.New.KeyValues.UsingKeyValueStrings([]string{"k"}, []string{"v"})
			_ = corestr.New.KeyValues.UsingKeyValueStrings(nil, nil)
		})

		actual := args.Map{
			"emptyIsEmpty": kv.IsEmpty(),
			"usingMapLen":  kv3.Length(),
			"noPanic":      noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC05_CollOfCollCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_CollOfCollCreator_Verification", func() {
		// Arrange
		tc := srcC05CollOfCollCreatorTestCase

		// Act
		c := corestr.New.CollectionsOfCollection.Empty()
		noPanic := !callPanicsSrcC05(func() {
			_ = corestr.New.CollectionsOfCollection.Cap(5)
			_ = corestr.New.CollectionsOfCollection.Strings([]string{"a"})
			_ = corestr.New.CollectionsOfCollection.CloneStrings([]string{"a"})
			_ = corestr.New.CollectionsOfCollection.StringsOption(true, 5, []string{"a"})
			_ = corestr.New.CollectionsOfCollection.StringsOptions(false, 0, []string{"a"})
			_ = corestr.New.CollectionsOfCollection.SpreadStrings(true, "a", "b")
			_ = corestr.New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})
			_ = corestr.New.CollectionsOfCollection.LenCap(0, 5)
		})

		actual := args.Map{
			"emptyIsEmpty": c.IsEmpty(),
			"noPanic":      noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC05_HashsetsCollCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_HashsetsCollCreator_Verification", func() {
		// Arrange
		tc := srcC05HashsetsCollCreatorTestCase

		// Act
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hc4 := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)
		noPanic := !callPanicsSrcC05(func() {
			_ = corestr.New.HashsetsCollection.Cap(5)
			_ = corestr.New.HashsetsCollection.LenCap(0, 5)
			_ = corestr.New.HashsetsCollection.UsingHashsetsPointers()
		})

		actual := args.Map{
			"emptyIsEmpty":   hc.IsEmpty(),
			"usingPtrsEmpty": hc4.IsEmpty(),
			"noPanic":        noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC05_CharMapCreators_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC05_CharMapCreators_Verification", func() {
		// Arrange
		tc := srcC05CharMapCreatorsTestCase

		// Act
		ccm := corestr.New.CharCollectionMap.Empty()
		noPanic := !callPanicsSrcC05(func() {
			_ = corestr.New.CharCollectionMap.CapSelfCap(20, 20)
			_ = corestr.New.CharCollectionMap.CapSelfCap(1, 1)
			_ = corestr.New.CharCollectionMap.Items([]string{"a", "b"})
			_ = corestr.New.CharCollectionMap.Items(nil)
			_ = corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 5, []string{"a"})
			_ = corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 5, nil)
			_ = corestr.New.CharHashsetMap.Cap(20, 20)
			_ = corestr.New.CharHashsetMap.Cap(1, 1)
			_ = corestr.New.CharHashsetMap.CapItems(20, 20, "a", "b")
			_ = corestr.New.CharHashsetMap.Strings(20, []string{"a"})
			_ = corestr.New.CharHashsetMap.Strings(20, nil)
		})

		actual := args.Map{
			"ccmEmptyIsEmpty": ccm.IsEmpty(),
			"noPanic":         noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC05(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
