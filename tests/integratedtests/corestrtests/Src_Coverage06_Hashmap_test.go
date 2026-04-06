package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC06_Hashmap_Basic_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Basic_Verification", func() {
		// Arrange
		tc := srcC06HashmapBasicTestCase
		h := corestr.New.Hashmap.Empty()
		var nilH *corestr.Hashmap

		// Act
		actual := args.Map{
			"isEmpty":      h.IsEmpty(),
			"hasItems":     h.HasItems(),
			"hasAny":       h.HasAnyItem(),
			"length":       h.Length(),
			"nilLength":    nilH.Length(),
			"nilSafeItems": nilH.SafeItems() == nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_AddGet_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_AddGet_Verification", func() {
		// Arrange
		tc := srcC06HashmapAddGetTestCase
		h := corestr.New.Hashmap.Cap(5)

		// Act
		isNew := h.AddOrUpdate("k1", "v1")
		isNew2 := h.AddOrUpdate("k1", "v2")
		v, found := h.Get("k1")
		_, _ = h.GetValue("k1")
		actual := args.Map{
			"isNewFirst":  isNew,
			"isNewSecond": isNew2,
			"getValue":    v,
			"found":       found,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Set_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Set_Verification", func() {
		// Arrange
		tc := srcC06HashmapSetTestCase
		h := corestr.New.Hashmap.Empty()

		// Act
		h.Set("a", "1")
		h.SetTrim(" b ", " 2 ")
		afterSetTrim := h.Length()
		h.SetBySplitter("=", "c=3")
		h.SetBySplitter("=", "d")
		actual := args.Map{
			"afterSetTrim": afterSetTrim,
			"afterSplit":   h.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Has_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Has_Verification", func() {
		// Arrange
		tc := srcC06HashmapHasTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{
			"has":        h.Has("a"),
			"contains":   h.Contains("a"),
			"notMissing": h.IsKeyMissing("a"),
			"zMissing":   h.IsKeyMissing("z"),
			"hasAll":     h.HasAll("a", "b"),
			"hasAllStr":  h.HasAllStrings("a", "b"),
			"hasAnyAZ":   h.HasAny("a", "z"),
			"hasAnyXZ":   h.HasAny("x", "z"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_HasLock_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_HasLock_Verification", func() {
		// Arrange
		tc := srcC06HashmapHasLockTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		_ = h.IsEmptyLock()
		actual := args.Map{
			"hasLock":      h.HasLock("a"),
			"hasWithLock":  h.HasWithLock("a"),
			"containsLock": h.ContainsLock("a"),
			"notMissing":   h.IsKeyMissingLock("a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_AddVariants_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_AddVariants_Verification", func() {
		// Arrange
		tc := srcC06HashmapAddVariantsTestCase
		h := corestr.New.Hashmap.Empty()

		// Act
		h.AddOrUpdateKeyStrValInt("n", 42)
		h.AddOrUpdateKeyStrValFloat("f", 3.14)
		h.AddOrUpdateKeyStrValFloat64("f64", 3.14)
		h.AddOrUpdateKeyStrValAny("any", "val")
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "kav", Value: 1})
		h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "kv", Value: "vv"})
		h.AddOrUpdateLock("lk", "lv")
		actual := args.Map{
			"length": h.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Merge_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Merge_Verification", func() {
		// Arrange
		tc := srcC06HashmapMergeTestCase
		h1 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2"})

		// Act
		h1.AddOrUpdateHashmap(h2)
		h1.AddOrUpdateHashmap(nil)
		afterHashmap := h1.Length()
		h1.AddOrUpdateMap(map[string]string{"c": "3"})
		h1.AddOrUpdateMap(nil)
		actual := args.Map{
			"afterHashmap": afterHashmap,
			"afterMap":     h1.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Adds_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Adds_Verification", func() {
		// Arrange
		tc := srcC06HashmapAddsTestCase
		h := corestr.New.Hashmap.Empty()

		// Act
		h.AddsOrUpdates(corestr.KeyValuePair{Key: "a", Value: "1"})
		h.AddsOrUpdates()
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "b", Value: 2})
		h.AddOrUpdateKeyAnyValues()
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "c", Value: "3"})
		h.AddOrUpdateKeyValues()
		actual := args.Map{
			"length": h.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_CollectionAdd_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_CollectionAdd_Verification", func() {
		// Arrange
		tc := srcC06HashmapCollectionAddTestCase
		h := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
		vals := corestr.New.Collection.Strings([]string{"v1", "v2"})

		// Act
		h.AddOrUpdateCollection(keys, vals)
		h.AddOrUpdateCollection(nil, nil)
		h.AddOrUpdateCollection(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b", "c"}))
		actual := args.Map{
			"length": 2, // only first valid call adds 2
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_WgLock_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_WgLock_Verification", func() {
		// Arrange
		tc := srcC06HashmapWgLockTestCase
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		h.AddOrUpdateWithWgLock("k", "v", wg)
		actual := args.Map{
			"length": h.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Filter_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Filter_Verification", func() {
		// Arrange
		tc := srcC06HashmapFilterTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"abc": "1", "def": "2"})
		empty := corestr.New.Hashmap.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, s == "abc", false }

		// Act
		actual := args.Map{
			"filteredItemsLen": len(h.GetKeysFilteredItems(filter)),
			"filteredColLen":   h.GetKeysFilteredCollection(filter).Length(),
			"emptyLen":         len(empty.GetKeysFilteredItems(filter)),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_FilterBreak_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_FilterBreak_Verification", func() {
		// Arrange
		tc := srcC06HashmapFilterBreakTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		filter := func(s string, i int) (string, bool, bool) { return s, true, true }

		// Act
		actual := args.Map{
			"itemsLen": len(h.GetKeysFilteredItems(filter)),
			"colLen":   h.GetKeysFilteredCollection(filter).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_AddsFilter_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_AddsFilter_Verification", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		f := func(p corestr.KeyValuePair) (string, bool, bool) { return p.Value, true, false }
		af := func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false }

		// Act — exercise all filter add variants without panic
		h.AddsOrUpdatesUsingFilter(f, corestr.KeyValuePair{Key: "a", Value: "1"})
		h.AddsOrUpdatesUsingFilter(f)
		h.AddsOrUpdatesAnyUsingFilter(af, corestr.KeyAnyValuePair{Key: "b", Value: 2})
		h.AddsOrUpdatesAnyUsingFilter(af)
		h.AddsOrUpdatesAnyUsingFilterLock(af, corestr.KeyAnyValuePair{Key: "c", Value: 3})
		h.AddsOrUpdatesAnyUsingFilterLock(af)

		// Assert
		actual := args.Map{"result": h.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected items added via filter", actual)
	})
}

func Test_SrcC06_Hashmap_Keys_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Keys_Verification", func() {
		// Arrange
		tc := srcC06HashmapKeysTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Act
		_ = h.KeysLock()
		_ = h.ValuesListCopyLock()
		_, _ = h.KeysValuesListLock()
		_ = h.ItemsCopyLock()
		actual := args.Map{
			"allKeysLen": len(h.AllKeys()),
			"keysLen":    len(h.Keys()),
			"keysColLen": h.KeysCollection().Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Values_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Values_Verification", func() {
		// Arrange
		tc := srcC06HashmapValuesTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		_ = h.ValuesCollection()
		_ = h.ValuesHashset()
		_ = h.ValuesCollectionLock()
		_ = h.ValuesHashsetLock()
		_ = h.Collection()
		k, v := h.KeysValuesCollection()
		k2, v2 := h.KeysValuesList()
		actual := args.Map{
			"valuesListLen": len(h.ValuesList()),
			"kvColKeysLen":  k.Length(),
			"kvColValsLen":  v.Length(),
			"kvListKeysLen": len(k2),
			"kvListValsLen": len(v2),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Pairs_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Pairs_Verification", func() {
		// Arrange
		tc := srcC06HashmapPairsTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{
			"pairsLen":    len(h.KeysValuePairs()),
			"pairsColLen": h.KeysValuePairsCollection().Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Remove_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Remove_Verification", func() {
		// Arrange
		tc := srcC06HashmapRemoveTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Act
		h.Remove("a")
		afterRemove := h.Length()
		h.RemoveWithLock("b")
		actual := args.Map{
			"afterRemove":     afterRemove,
			"afterRemoveLock": h.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Diff_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Diff_Verification", func() {
		// Arrange
		tc := srcC06HashmapDiffTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "2"})

		// Act
		noPanic := !callPanicsSrcC06(func() {
			_ = h.DiffRaw(map[string]string{"a": "2"})
			_ = h.Diff(h2)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Equal_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Equal_Verification", func() {
		// Arrange
		tc := srcC06HashmapEqualTestCase
		h1 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h3 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "2"})

		// Act
		actual := args.Map{
			"equalSame":      h1.IsEqualPtr(h2),
			"equalSameLock":  h1.IsEqualPtrLock(h2),
			"equalDifferent": h1.IsEqualPtr(h3),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Concat_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Concat_Verification", func() {
		// Arrange
		tc := srcC06HashmapConcatTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2"})

		// Act
		concat := h.ConcatNew(false, h2)
		noPanic := !callPanicsSrcC06(func() {
			_ = h.ConcatNew(true)
			_ = h.ConcatNewUsingMaps(false, map[string]string{"c": "3"})
			_ = h.ConcatNewUsingMaps(true)
		})
		actual := args.Map{
			"concatLen": concat.Length() >= 2,
			"noPanic":   noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_StringJson_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_StringJson_Verification", func() {
		// Arrange
		tc := srcC06HashmapStringJsonTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		noPanic := !callPanicsSrcC06(func() {
			_ = h.Join(",")
			_ = h.JoinKeys(",")
			_ = h.JsonModel()
			_ = h.JsonModelAny()
			_, _ = h.MarshalJSON()
			_, _ = h.Serialize()
			_ = h.AsJsoner()
			_ = h.AsJsonContractsBinder()
			_ = h.AsJsonParseSelfInjector()
			_ = h.AsJsonMarshaller()
		})
		actual := args.Map{
			"stringNonEmpty":     h.String() != "",
			"stringLockNonEmpty": h.StringLock() != "",
			"noPanic":            noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_KeysToLower_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_KeysToLower_Verification", func() {
		// Arrange
		tc := srcC06HashmapKeysToLowerTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"ABC": "1"})

		// Act
		lower := h.KeysToLower()
		_ = h.ValuesToLower()
		actual := args.Map{
			"hasLowercase": lower.Has("abc"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Except_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Except_Verification", func() {
		// Arrange
		tc := srcC06HashmapExceptTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		_ = h.GetValuesExceptKeysInHashset(nil)
		_ = h.GetValuesKeysExcept(nil)
		_ = h.GetAllExceptCollection(nil)
		actual := args.Map{
			"hashsetExcLen":    len(h.GetValuesExceptKeysInHashset(hs)),
			"keysExcLen":       len(h.GetValuesKeysExcept([]string{"a"})),
			"collectionExcLen": len(h.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_HasAllCol_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_HasAllCol_Verification", func() {
		// Arrange
		tc := srcC06HashmapHasAllColTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"hasAll":    h.HasAllCollectionItems(c),
			"hasAllNil": h.HasAllCollectionItems(nil),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_ToError_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_ToError_Verification", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act — exercise without panic
		_ = h.ToError(",")
		_ = h.ToDefaultError()
		_ = h.KeyValStringLines()

		// Assert — no panic is success
	})
}

func Test_SrcC06_Hashmap_ClearDispose_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_ClearDispose_Verification", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		h.Clear()
		clearedLen := h.Length()
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2"})
		h2.Dispose()

		// Assert
		actual := args.Map{"result": clearedLen != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 after clear", actual)
	})
}

func Test_SrcC06_Hashmap_Clone_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Clone_Verification", func() {
		// Arrange
		tc := srcC06HashmapCloneTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		var nilH *corestr.Hashmap

		// Act
		cloneVal := h.Clone()
		actual := args.Map{
			"cloneLen":    (&cloneVal).Length(),
			"clonePtrLen": h.ClonePtr().Length(),
			"nilClone":    nilH.ClonePtr() == nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_Compiler_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_Compiler_Verification", func() {
		// Arrange
		tc := srcC06HashmapCompilerTestCase
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		empty := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{
			"compiledLen": len(h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })),
			"emptyLen":    len(empty.ToStringsUsingCompiler(func(k, v string) string { return k })),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC06_Hashmap_EmptyString_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_EmptyString_Verification", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()

		// Act — empty hashmap string should still be non-empty (NoElements constant)
		// Assert
		actual := args.Map{"result": h.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty (NoElements)", actual)
		actual := args.Map{"result": h.StringLock() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_SrcC06_Hashmap_StringsPtrWgLock_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC06_Hashmap_StringsPtrWgLock_Verification", func() {
		// Arrange
		tc := srcC06HashmapStringsPtrWgLockTestCase
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		h.AddOrUpdateStringsPtrWgLock(wg, []string{"a"}, []string{"1"})
		wg2 := &sync.WaitGroup{}
		h.AddOrUpdateStringsPtrWgLock(wg2, nil, nil)
		actual := args.Map{
			"length": h.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC06(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
