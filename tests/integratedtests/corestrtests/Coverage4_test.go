package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Hashmap — additional methods ──

func Test_Cov4_Hashmap_AllKeys(t *testing.T) {
	safeTest(t, "Test_Cov4_Hashmap_AllKeys", func() {
		hm := corestr.New.Hashmap.KeyValues(
			corestr.KeyValuePair{Key: "b", Value: "2"},
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)
		keys := hm.AllKeys()
		actual := args.Map{
			"keysLen": len(keys),
		}
		expected := args.Map{
			"keysLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap AllKeys returns expected -- 2 items", actual)
	})
}

func Test_Cov4_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_Cov4_Hashmap_Remove", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k1", "v1")
		hm.AddOrUpdate("k2", "v2")
		hm.Remove("k1")
		actual := args.Map{"length": hm.Length(), "hasK1": hm.Has("k1"), "hasK2": hm.Has("k2")}
		expected := args.Map{"length": 1, "hasK1": false, "hasK2": true}
		expected.ShouldBeEqual(t, 0, "Hashmap Remove deletes key -- removed k1", actual)
	})
}

func Test_Cov4_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_Cov4_Hashmap_String", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		actual := args.Map{"notEmpty": hm.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap String returns non-empty -- one item", actual)
	})
}

func Test_Cov4_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_Cov4_Hashmap_Get", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		val, found := hm.Get("k")
		_, notFound := hm.Get("missing")
		actual := args.Map{"val": val, "found": found, "notFound": notFound}
		expected := args.Map{"val": "v", "found": true, "notFound": false}
		expected.ShouldBeEqual(t, 0, "Hashmap Get returns expected -- hit and miss", actual)
	})
}

// ── Collection — serialization and iteration ──

func Test_Cov4_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_Cov4_Collection_ListStrings", func() {
		col := corestr.New.Collection.Strings([]string{"x", "y"})
		strs := col.ListStrings()
		actual := args.Map{"len": len(strs), "first": strs[0]}
		expected := args.Map{"len": 2, "first": "x"}
		expected.ShouldBeEqual(t, 0, "Collection ListStrings returns expected -- 2 items", actual)
	})
}

func Test_Cov4_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov4_Collection_IndexAt", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{
			"at0": col.IndexAt(0),
		}
		expected := args.Map{"at0": "a"}
		expected.ShouldBeEqual(t, 0, "Collection IndexAt returns expected -- valid index", actual)
	})
}

func Test_Cov4_Collection_First_Last(t *testing.T) {
	safeTest(t, "Test_Cov4_Collection_First_Last", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"first": col.First(), "last": col.Last()}
		expected := args.Map{"first": "a", "last": "c"}
		expected.ShouldBeEqual(t, 0, "Collection First/Last returns expected -- 3 items", actual)
	})
}

func Test_Cov4_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_Cov4_Collection_Filter", func() {
		col := corestr.New.Collection.Strings([]string{"ab", "cd", "abc"})
		filtered := col.Filter(func(str string, index int) (string, bool, bool) {
			return str, len(str) == 2, false
		})
		actual := args.Map{"len": len(filtered)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection Filter returns 2 -- length 2 items", actual)
	})
}

// ── LinkedList — traversal ──

func Test_Cov4_LinkedList_Traversal(t *testing.T) {
	safeTest(t, "Test_Cov4_LinkedList_Traversal", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.Add("a")
		ll.Add("b")
		ll.Add("c")
		actual := args.Map{
			"length":  ll.Length(),
			"isEmpty": ll.IsEmpty(),
			"first":   ll.Head().String(),
			"last":    ll.Tail().String(),
		}
		expected := args.Map{
			"length": 3, "isEmpty": false,
			"first": "a", "last": "c",
		}
		expected.ShouldBeEqual(t, 0, "LinkedList traversal returns expected -- 3 items", actual)
	})
}

func Test_Cov4_LinkedList_HeadList(t *testing.T) {
	safeTest(t, "Test_Cov4_LinkedList_HeadList", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.Add("x")
		ll.Add("y")
		strs := ll.Head().List()
		actual := args.Map{"len": len(strs)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedList Head.List returns expected -- 2 items", actual)
	})
}

// ── SimpleStringOnce ──

func Test_Cov4_SimpleStringOnce_Value(t *testing.T) {
	safeTest(t, "Test_Cov4_SimpleStringOnce_Value", func() {
		so := corestr.New.SimpleStringOnce.Init("hello")
		actual := args.Map{
			"isEmpty": so.IsEmpty(),
			"value":   so.Value(),
		}
		expected := args.Map{
			"isEmpty": false, "value": "hello",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce Init returns expected -- hello", actual)
	})
}

func Test_Cov4_SimpleStringOnce_Empty(t *testing.T) {
	safeTest(t, "Test_Cov4_SimpleStringOnce_Empty", func() {
		so := corestr.New.SimpleStringOnce.Empty()
		actual := args.Map{"isEmpty": so.IsEmpty()}
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce Empty returns true -- empty", actual)
	})
}

// ── CharHashsetMap ──

func Test_Cov4_CharHashsetMap_AddAndHas(t *testing.T) {
	safeTest(t, "Test_Cov4_CharHashsetMap_AddAndHas", func() {
		chm := corestr.New.CharHashsetMap.Cap(5, 5)
		chm.Add("alpha")
		chm.Add("also")
		chm.Add("beta")
		actual := args.Map{
			"isEmpty": chm.IsEmpty(),
			"length":  chm.Length(),
		}
		expected := args.Map{"isEmpty": false, "length": 2}
		expected.ShouldBeEqual(t, 0, "CharHashsetMap Add returns expected -- 2 chars", actual)
	})
}

// ── CharCollectionMap ──

func Test_Cov4_CharCollectionMap_AddAndGet(t *testing.T) {
	safeTest(t, "Test_Cov4_CharCollectionMap_AddAndGet", func() {
		ccm := corestr.New.CharCollectionMap.Empty()
		ccm.Add("val1")
		ccm.Add("val2")
		actual := args.Map{"isEmpty": ccm.IsEmpty(), "length": ccm.Length()}
		expected := args.Map{"isEmpty": false, "length": 1}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap Add returns expected -- 1 char", actual)
	})
}

// ── HashsetsCollection ──

func Test_Cov4_HashsetsCollection_Add(t *testing.T) {
	safeTest(t, "Test_Cov4_HashsetsCollection_Add", func() {
		hc := corestr.New.HashsetsCollection.Cap(5)
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc.Add(hs)
		actual := args.Map{"isEmpty": hc.IsEmpty(), "length": hc.Length()}
		expected := args.Map{"isEmpty": false, "length": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection Add returns expected -- 1 hashset", actual)
	})
}

// ── CollectionsOfCollection ──

func Test_Cov4_CollectionsOfCollection_Add(t *testing.T) {
	safeTest(t, "Test_Cov4_CollectionsOfCollection_Add", func() {
		cc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a"})
		cc.Add(col)
		actual := args.Map{"isEmpty": cc.IsEmpty(), "length": cc.Length()}
		expected := args.Map{"isEmpty": false, "length": 1}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection Add returns expected -- 1 collection", actual)
	})
}

// ── LeftRight — HasSafeNonEmpty full coverage ──

func Test_Cov4_LeftRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov4_LeftRight_HasSafeNonEmpty", func() {
		lr := corestr.NewLeftRight("l", "r")
		actual := args.Map{"hasSafe": lr.HasSafeNonEmpty()}
		expected := args.Map{"hasSafe": true}
		expected.ShouldBeEqual(t, 0, "LeftRight HasSafeNonEmpty returns true -- both set", actual)
	})
}

// ── LeftMiddleRight — HasSafeNonEmpty ──

func Test_Cov4_LeftMiddleRight_HasSafe(t *testing.T) {
	safeTest(t, "Test_Cov4_LeftMiddleRight_HasSafe", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		actual := args.Map{
			"hasSafe": lmr.HasSafeNonEmpty(),
			"isAll":   lmr.IsAll("l", "m", "r"),
		}
		expected := args.Map{"hasSafe": true, "isAll": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight HasSafe/IsAll returns true -- all set", actual)
	})
}

// ── Hashset — remove and HasAny ──

func Test_Cov4_Hashset_HasAny(t *testing.T) {
	safeTest(t, "Test_Cov4_Hashset_HasAny", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		actual := args.Map{
			"hasAnyTrue":  hs.HasAny("x", "b"),
			"hasAnyFalse": hs.HasAny("x", "y"),
			"hasAll":      hs.HasAll("a", "b"),
			"hasAllFalse": hs.HasAll("a", "z"),
		}
		expected := args.Map{
			"hasAnyTrue": true, "hasAnyFalse": false,
			"hasAll": true, "hasAllFalse": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset HasAny/HasAll returns expected -- various", actual)
	})
}

// ── Hashmap — HasAny and HasAll ──

func Test_Cov4_Hashmap_HasAny(t *testing.T) {
	safeTest(t, "Test_Cov4_Hashmap_HasAny", func() {
		hm := corestr.New.Hashmap.KeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		actual := args.Map{
			"hasAnyTrue":  hm.HasAny("x", "a"),
			"hasAnyFalse": hm.HasAny("x", "y"),
			"hasAll":      hm.HasAll("a", "b"),
			"hasAllFalse": hm.HasAll("a", "z"),
		}
		expected := args.Map{
			"hasAnyTrue": true, "hasAnyFalse": false,
			"hasAll": true, "hasAllFalse": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap HasAny/HasAll returns expected -- various", actual)
	})
}

// ── Hashset — String ──

func Test_Cov4_Hashset_String(t *testing.T) {
	safeTest(t, "Test_Cov4_Hashset_String", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"notEmpty": hs.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset String returns non-empty -- 1 item", actual)
	})
}

// ── Collection — Clone ──

func Test_Cov4_Collection_Clone(t *testing.T) {
	safeTest(t, "Test_Cov4_Collection_Clone", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		cloned := corestr.New.Collection.CloneStrings(col.ListStrings())
		actual := args.Map{"length": cloned.Length()}
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "Collection Clone returns expected -- 2 items", actual)
	})
}

// ── SimpleSlice — FirstOrDefault / LastOrDefault ──

func Test_Cov4_SimpleSlice_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_Cov4_SimpleSlice_FirstOrDefault", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		empty := corestr.New.SimpleSlice.Empty()
		actual := args.Map{
			"first":        ss.FirstOrDefault(),
			"emptyDefault": empty.FirstOrDefault(),
			"lastOrDef":    ss.LastOrDefault(),
		}
		expected := args.Map{"first": "a", "emptyDefault": "", "lastOrDef": "b"}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- FirstOrDefault/LastOrDefault returns expected", actual)
	})
}

// ── SimpleSlice — SafeAt ──

func Test_Cov4_SimpleSlice_SafeAt(t *testing.T) {
	safeTest(t, "Test_Cov4_SimpleSlice_SafeAt", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		hasIdx0 := ss.HasIndex(0)
		hasIdx99 := ss.HasIndex(99)
		hasIdxNeg := ss.HasIndex(-1)
		actual := args.Map{
			"hasIdx0":   hasIdx0,
			"hasIdx99":  hasIdx99,
			"hasIdxNeg": hasIdxNeg,
		}
		expected := args.Map{"hasIdx0": true, "hasIdx99": false, "hasIdxNeg": false}
		expected.ShouldBeEqual(t, 0, "SimpleSlice HasIndex returns expected -- valid and out of bounds", actual)
	})
}
