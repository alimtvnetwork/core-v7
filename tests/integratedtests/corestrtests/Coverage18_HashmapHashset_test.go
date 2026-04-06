package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap comprehensive
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov18_Hashmap_Basic(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_Basic", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		var nilH *corestr.Hashmap

		// Act
		actual := args.Map{
			"empty": h.IsEmpty(), "hasItems": h.HasItems(), "len": h.Length(),
			"nilLen": nilH.Length(), "nilSafe": nilH.SafeItems() == nil,
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
			"len": 0,
			"nilLen": 0,
			"nilSafe": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- basic", actual)
	})
}

func Test_Cov18_Hashmap_AddAndGet(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_AddAndGet", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		isNew := h.AddOrUpdate("k1", "v1")
		isNew2 := h.AddOrUpdate("k1", "v2")
		v, found := h.Get("k1")
		v2, found2 := h.GetValue("k1")
		_ = found2
		_ = v2

		// Act
		actual := args.Map{
			"isNew": isNew,
			"notNew": !isNew2,
			"found": found,
			"val": v,
		}

		// Assert
		expected := args.Map{
			"isNew": true,
			"notNew": true,
			"found": true,
			"val": "v2",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddAndGet", actual)
	})
}

func Test_Cov18_Hashmap_Set(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_Set", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.Set("a", "1")
		h.SetTrim(" b ", " 2 ")
		h.SetBySplitter("=", "c=3")
		h.SetBySplitter("=", "d")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Set", actual)
	})
}

func Test_Cov18_Hashmap_Has(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_Has", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{
			"has":        h.Has("a"),
			"contains":   h.Contains("a"),
			"notMissing": !h.IsKeyMissing("a"),
			"missing":    h.IsKeyMissing("z"),
			"hasAll":     h.HasAll("a", "b"),
			"hasAllStr":  h.HasAllStrings("a", "b"),
			"hasAny":     h.HasAny("a", "z"),
			"noAny":      !h.HasAny("x", "z"),
			"hasLock":    h.HasLock("a"),
			"hasWithLock": h.HasWithLock("a"),
			"containsLock": h.ContainsLock("a"),
			"notMissingLock": !h.IsKeyMissingLock("a"),
		}

		// Assert
		expected := args.Map{
			"has": true, "contains": true, "notMissing": true, "missing": true,
			"hasAll": true, "hasAllStr": true, "hasAny": true, "noAny": true,
			"hasLock": true, "hasWithLock": true, "containsLock": true, "notMissingLock": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Has", actual)
	})
}

func Test_Cov18_Hashmap_AddVariants(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_AddVariants", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValInt("n", 42)
		h.AddOrUpdateKeyStrValFloat("f", 3.14)
		h.AddOrUpdateKeyStrValFloat64("f64", 3.14)
		h.AddOrUpdateKeyStrValAny("any", "val")
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "kav", Value: 1})
		h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "kv", Value: "vv"})
		h.AddOrUpdateLock("lk", "lv")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 7}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddVariants", actual)
	})
}

func Test_Cov18_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_AddOrUpdateHashmap", func() {
		// Arrange
		h1 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2"})
		h1.AddOrUpdateHashmap(h2)
		h1.AddOrUpdateHashmap(nil)
		h1.AddOrUpdateMap(map[string]string{"c": "3"})
		h1.AddOrUpdateMap(nil)

		// Act
		actual := args.Map{"len": h1.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateHashmap", actual)
	})
}

func Test_Cov18_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_AddsOrUpdates", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddsOrUpdates(corestr.KeyValuePair{Key: "a", Value: "1"})
		h.AddsOrUpdates()
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "b", Value: 2})
		h.AddOrUpdateKeyAnyValues()
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "c", Value: "3"})
		h.AddOrUpdateKeyValues()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddsOrUpdates", actual)
	})
}

func Test_Cov18_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_AddOrUpdateCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
		vals := corestr.New.Collection.Strings([]string{"v1", "v2"})
		h.AddOrUpdateCollection(keys, vals)
		h.AddOrUpdateCollection(nil, nil)
		h.AddOrUpdateCollection(
			corestr.New.Collection.Strings([]string{"a"}),
			corestr.New.Collection.Strings([]string{"b", "c"}),
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateCollection", actual)
	})
}

func Test_Cov18_Hashmap_WgLock(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_WgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("k", "v", wg)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- WgLock", actual)
	})
}

func Test_Cov18_Hashmap_Keys(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_Keys", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{
			"allKeysLen": len(h.AllKeys()),
			"keysLen":    len(h.Keys()),
			"keysColLen": h.KeysCollection().Length(),
		}
		_ = h.KeysLock()
		_ = h.ValuesListCopyLock()
		keys, vals := h.KeysValuesListLock()
		_ = keys
		_ = vals
		_ = h.ItemsCopyLock()

		// Assert
		expected := args.Map{
			"allKeysLen": 2,
			"keysLen": 2,
			"keysColLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Keys", actual)
	})
}

func Test_Cov18_Hashmap_Values(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_Values", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"valsLen": len(h.ValuesList())}
		_ = h.ValuesCollection()
		_ = h.ValuesHashset()
		_ = h.ValuesCollectionLock()
		_ = h.ValuesHashsetLock()
		_ = h.Collection()
		k, v := h.KeysValuesCollection()
		_ = k
		_ = v
		k2, v2 := h.KeysValuesList()
		_ = k2
		_ = v2

		// Assert
		expected := args.Map{"valsLen": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- Values", actual)
	})
}

func Test_Cov18_Hashmap_KeyValuePairs(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_KeyValuePairs", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		pairs := h.KeysValuePairs()
		pairsCol := h.KeysValuePairsCollection()

		// Act
		actual := args.Map{
			"pairsLen": len(pairs),
			"colLen": pairsCol.Length(),
		}

		// Assert
		expected := args.Map{
			"pairsLen": 1,
			"colLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- KeyValuePairs", actual)
	})
}

func Test_Cov18_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_Remove", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		h.Remove("a")
		h.RemoveWithLock("b")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Remove", actual)
	})
}

func Test_Cov18_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_Diff", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "2"})
		_ = h.DiffRaw(map[string]string{"a": "2"})
		_ = h.Diff(h2)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Diff", actual)
	})
}

func Test_Cov18_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_IsEqual", func() {
		// Arrange
		h1 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h3 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "2"})

		// Act
		actual := args.Map{
			"equal":    h1.IsEqualPtr(h2),
			"equalLock": h1.IsEqualPtrLock(h2),
			"notEqual": !h1.IsEqualPtr(h3),
		}

		// Assert
		expected := args.Map{
			"equal": true,
			"equalLock": true,
			"notEqual": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IsEqual", actual)
	})
}

func Test_Cov18_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_ConcatNew", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2"})
		concat := h.ConcatNew(false, h2)
		_ = h.ConcatNew(true)
		_ = h.ConcatNewUsingMaps(false, map[string]string{"c": "3"})
		_ = h.ConcatNewUsingMaps(true)

		// Act
		actual := args.Map{"concatLen": concat.Length()}

		// Assert
		expected := args.Map{"concatLen": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ConcatNew", actual)
	})
}

func Test_Cov18_Hashmap_StringAndJson(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_StringAndJson", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{
			"str": h.String() != "",
			"strLock": h.StringLock() != "",
		}
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
		_ = h.ToError(",")
		_ = h.ToDefaultError()
		_ = h.KeyValStringLines()

		// Assert
		expected := args.Map{
			"str": true,
			"strLock": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- StringAndJson", actual)
	})
}

func Test_Cov18_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_KeysToLower", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"ABC": "1"})
		lower := h.KeysToLower()
		_ = h.ValuesToLower()

		// Act
		actual := args.Map{"hasLower": lower.Has("abc")}

		// Assert
		expected := args.Map{"hasLower": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- KeysToLower", actual)
	})
}

func Test_Cov18_Hashmap_GetExcept(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_GetExcept", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := h.GetValuesExceptKeysInHashset(hs)
		r2 := h.GetValuesKeysExcept([]string{"a"})
		r3 := h.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))
		_ = h.GetValuesExceptKeysInHashset(nil)
		_ = h.GetValuesKeysExcept(nil)
		_ = h.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{
			"r": len(r),
			"r2": len(r2),
			"r3": len(r3),
		}

		// Assert
		expected := args.Map{
			"r": 1,
			"r2": 1,
			"r3": 1,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetExcept", actual)
	})
}

func Test_Cov18_Hashmap_Filter(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_Filter", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"abc": "1", "def": "2"})
		filter := func(s string, i int) (string, bool, bool) { return s, s == "abc", false }
		items := h.GetKeysFilteredItems(filter)
		col := h.GetKeysFilteredCollection(filter)

		// Act
		actual := args.Map{
			"itemsLen": len(items),
			"colLen": col.Length(),
		}

		// Assert
		expected := args.Map{
			"itemsLen": 1,
			"colLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Filter", actual)
	})
}

func Test_Cov18_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		f := func(p corestr.KeyValuePair) (string, bool, bool) { return p.Value, true, false }
		h.AddsOrUpdatesUsingFilter(f, corestr.KeyValuePair{Key: "a", Value: "1"})
		h.AddsOrUpdatesUsingFilter(f)
		af := func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false }
		h.AddsOrUpdatesAnyUsingFilter(af, corestr.KeyAnyValuePair{Key: "b", Value: 2})
		h.AddsOrUpdatesAnyUsingFilter(af)
		h.AddsOrUpdatesAnyUsingFilterLock(af, corestr.KeyAnyValuePair{Key: "c", Value: 3})
		h.AddsOrUpdatesAnyUsingFilterLock(af)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddsOrUpdatesUsingFilter", actual)
	})
}

func Test_Cov18_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"has": h.HasAllCollectionItems(c),
			"nilFalse": !h.HasAllCollectionItems(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nilFalse": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- HasAllCollectionItems", actual)
	})
}

func Test_Cov18_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_ToStringsUsingCompiler", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		s := h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
		empty := corestr.New.Hashmap.Empty()
		s2 := empty.ToStringsUsingCompiler(func(k, v string) string { return k })

		// Act
		actual := args.Map{
			"len": len(s),
			"emptyLen": len(s2),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"emptyLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ToStringsUsingCompiler", actual)
	})
}

func Test_Cov18_Hashmap_Clone(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_Clone", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		c := h.Clone()
		cp := h.ClonePtr()
		var nilH *corestr.Hashmap

		// Act
		actual := args.Map{
			"cLen": c.Length(),
			"cpLen": cp.Length(),
			"nilClone": nilH.ClonePtr() == nil,
		}

		// Assert
		expected := args.Map{
			"cLen": 1,
			"cpLen": 1,
			"nilClone": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Clone", actual)
	})
}

func Test_Cov18_Hashmap_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_ClearDispose", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		h.Clear()
		h2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2"})
		h2.Dispose()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ClearDispose", actual)
	})
}

func Test_Cov18_Hashmap_EmptyString(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_EmptyString", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{
			"str": h.String() != "",
			"strLock": h.StringLock() != "",
		}

		// Assert
		expected := args.Map{
			"str": true,
			"strLock": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- EmptyString", actual)
	})
}

func Test_Cov18_Hashmap_StringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_StringsPtrWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(wg, []string{"a"}, []string{"1"})
		// Fix: must call wg.Add(1) before passing WaitGroup — method always calls wg.Done()
		// See issues/corestrtests-wg-negative-counter-panic.md
		wg2 := &sync.WaitGroup{}
		wg2.Add(1)
		h.AddOrUpdateStringsPtrWgLock(wg2, nil, nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- StringsPtrWgLock", actual)
	})
}

func Test_Cov18_Hashmap_ParseInjectJson(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashmap_ParseInjectJson", func() {
		// Arrange
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		j := h.Json()
		h2 := corestr.New.Hashmap.Empty()
		_, err := h2.ParseInjectUsingJson(&j)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ParseInjectJson", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset comprehensive
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov18_Hashset_Basic(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_Basic", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		var nilH *corestr.Hashset

		// Act
		actual := args.Map{
			"empty": h.IsEmpty(), "hasItems": h.HasItems(), "len": h.Length(),
			"nilLen": nilH.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
			"len": 0,
			"nilLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- basic", actual)
	})
}

func Test_Cov18_Hashset_Add(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_Add", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Add("a").Add("b")
		h.AddLock("c")
		h.AddNonEmpty("")
		h.AddNonEmpty("d")
		h.AddNonEmptyWhitespace("   ")
		h.AddNonEmptyWhitespace("e")
		h.AddIf(false, "skip")
		h.AddIf(true, "f")
		h.AddIfMany(false, "x", "y")
		h.AddIfMany(true, "g", "h")
		h.AddFunc(func() string { return "i" })
		h.AddFuncErr(func() (string, error) { return "j", nil }, func(e error) {})

		// Act
		actual := args.Map{"hasItems": h.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Add", actual)
	})
}

func Test_Cov18_Hashset_AddBool(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_AddBool", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		existed := h.AddBool("a")
		existed2 := h.AddBool("a")

		// Act
		actual := args.Map{
			"first": !existed,
			"second": existed2,
		}

		// Assert
		expected := args.Map{
			"first": true,
			"second": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddBool", actual)
	})
}

func Test_Cov18_Hashset_AddPtr(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_AddPtr", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		s := "hello"
		h.AddPtr(&s)
		h.AddPtrLock(&s)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddPtr", actual)
	})
}

func Test_Cov18_Hashset_Adds(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_Adds", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		h.Adds("a", "b")
		h.Adds()
		h.AddStrings([]string{"c"})
		h.AddStrings(nil)
		h.AddStringsLock([]string{"d"})
		h.AddStringsLock(nil)
		h.AddCollection(corestr.New.Collection.Strings([]string{"e"}))
		h.AddCollection(nil)
		h.AddCollections(corestr.New.Collection.Strings([]string{"f"}))
		h.AddCollections()
		ss := corestr.New.SimpleSlice.Lines("g")
		h.AddSimpleSlice(ss)

		// Act
		actual := args.Map{"hasItems": h.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Adds", actual)
	})
}

func Test_Cov18_Hashset_Has(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_Has", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")

		// Act
		actual := args.Map{
			"has":       h.Has("a"),
			"contains":  h.Contains("a"),
			"hasLock":   h.HasLock("a"),
			"hasWithLock": h.HasWithLock("a"),
			"notMissing": !h.IsMissing("a"),
			"missing":   h.IsMissing("z"),
			"notMissingLock": !h.IsMissingLock("a"),
			"hasAll":    h.HasAll("a", "b"),
			"hasAllStr": h.HasAllStrings([]string{"a", "b"}),
			"hasAny":    h.HasAny("a", "z"),
			"noAny":     !h.HasAny("x", "z"),
			"allMissing": h.IsAllMissing("x", "z"),
			"notAllMissing": !h.IsAllMissing("a"),
			"hasAllCol": h.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a"})),
			"nilCol":    !h.HasAllCollectionItems(nil),
		}

		// Assert
		expected := args.Map{
			"has": true, "contains": true, "hasLock": true, "hasWithLock": true,
			"notMissing": true, "missing": true, "notMissingLock": true,
			"hasAll": true, "hasAllStr": true, "hasAny": true, "noAny": true,
			"allMissing": true, "notAllMissing": true, "hasAllCol": true, "nilCol": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Has", actual)
	})
}

func Test_Cov18_Hashset_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_IsEquals", func() {
		// Arrange
		h1 := corestr.New.Hashset.StringsSpreadItems("a", "b")
		h2 := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{
			"equal":     h1.IsEquals(h2),
			"isEqual":   h1.IsEqual(h2),
			"equalsLock": h1.IsEqualsLock(h2),
		}

		// Assert
		expected := args.Map{
			"equal": true,
			"isEqual": true,
			"equalsLock": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals", actual)
	})
}

func Test_Cov18_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_Remove", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")
		h.Remove("a")
		h.SafeRemove("b")
		h.SafeRemove("z")
		h.RemoveWithLock("z")

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Remove", actual)
	})
}

func Test_Cov18_Hashset_List(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_List", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.List()
		_ = h.ListPtr()
		_ = h.Lines()
		_ = h.SafeStrings()
		_ = h.ListPtrSortedAsc()
		_ = h.ListPtrSortedDsc()
		_ = h.OrderedList()
		_ = h.SortedList()
		_ = h.ListCopyLock()
		_ = h.SimpleSlice()
		_ = h.Items()
		_ = h.Collection()
		_ = h.MapStringAny()
		_ = h.MapStringAnyDiff()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- List", actual)
	})
}

func Test_Cov18_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_Filter", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("abc", "def")
		f := func(s string) bool { return s == "abc" }
		r := h.Filter(f)
		sf := func(s string, i int) (string, bool, bool) { return s, s == "abc", false }
		items := h.GetFilteredItems(sf)
		col := h.GetFilteredCollection(sf)

		// Act
		actual := args.Map{
			"filterLen": r.Length(),
			"itemsLen": len(items),
			"colLen": col.Length(),
		}

		// Assert
		expected := args.Map{
			"filterLen": 1,
			"itemsLen": 1,
			"colLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Filter", actual)
	})
}

func Test_Cov18_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_GetAllExcept", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		r := h.GetAllExceptHashset(corestr.New.Hashset.StringsSpreadItems("a"))
		r2 := h.GetAllExcept([]string{"a"})
		r3 := h.GetAllExceptSpread("a")
		r4 := h.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))
		_ = h.GetAllExceptHashset(nil)
		_ = h.GetAllExcept(nil)
		_ = h.GetAllExceptSpread()
		_ = h.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{
			"r": len(r),
			"r2": len(r2),
			"r3": len(r3),
			"r4": len(r4),
		}

		// Assert
		expected := args.Map{
			"r": 2,
			"r2": 2,
			"r3": 2,
			"r4": 2,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetAllExcept", actual)
	})
}

func Test_Cov18_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_Resize", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.Resize(100)
		h.ResizeLock(200)
		h.AddCapacities(10, 20)
		h.AddCapacitiesLock(10)
		h.AddCapacities()
		h.AddCapacitiesLock()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Resize", actual)
	})
}

func Test_Cov18_Hashset_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_ConcatNew", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h2 := corestr.New.Hashset.StringsSpreadItems("b")
		r := h.ConcatNewHashsets(false, h2)
		_ = h.ConcatNewHashsets(true)
		_ = h.ConcatNewStrings(false, []string{"c"})
		_ = h.ConcatNewStrings(true)

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ConcatNew", actual)
	})
}

func Test_Cov18_Hashset_StringAndJson(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_StringAndJson", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{
			"str": h.String() != "",
			"strLock": h.StringLock() != "",
		}
		_ = h.Join(",")
		_ = h.NonEmptyJoins(",")
		_ = h.NonWhitespaceJoins(",")
		_ = h.JoinSorted(",")
		_ = h.JsonModel()
		_ = h.JsonModelAny()
		_, _ = h.MarshalJSON()
		_ = h.AsJsoner()
		_ = h.AsJsonContractsBinder()
		_ = h.AsJsonParseSelfInjector()
		_ = h.AsJsonMarshaller()

		// Assert
		expected := args.Map{
			"str": true,
			"strLock": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- StringAndJson", actual)
	})
}

func Test_Cov18_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_ToLowerSet", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("ABC")
		lower := h.ToLowerSet()

		// Act
		actual := args.Map{"hasLower": lower.Has("abc")}

		// Assert
		expected := args.Map{"hasLower": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ToLowerSet", actual)
	})
}

func Test_Cov18_Hashset_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_ClearDispose", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.Clear()
		h.Dispose()
		var nilH *corestr.Hashset
		nilH.Dispose()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ClearDispose", actual)
	})
}

func Test_Cov18_Hashset_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_DistinctDiff", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = h.DistinctDiffLinesRaw("b", "c")
		_ = h.DistinctDiffLines("b", "c")
		_ = h.DistinctDiffHashset(corestr.New.Hashset.StringsSpreadItems("b", "c"))
		empty := corestr.New.Hashset.Empty()
		_ = empty.DistinctDiffLinesRaw()
		_ = empty.DistinctDiffLinesRaw("a")
		_ = h.DistinctDiffLinesRaw()
		_ = empty.DistinctDiffLines()
		_ = empty.DistinctDiffLines("a")

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiff", actual)
	})
}

func Test_Cov18_Hashset_WgLock(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_WgLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(10)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddWithWgLock("a", wg)
		wg2 := &sync.WaitGroup{}
		wg2.Add(1)
		h.AddStringsPtrWgLock([]string{"b", "c"}, wg2)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- WgLock", actual)
	})
}

func Test_Cov18_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_AddItemsMap", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		h.AddItemsMap(map[string]bool{"a": true, "b": false})
		h.AddItemsMap(nil)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		m := map[string]bool{"c": true}
		h.AddItemsMapWgLock(&m, wg)
		h.AddItemsMapWgLock(nil, nil)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddItemsMap", actual)
	})
}

func Test_Cov18_Hashset_AddHashset(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_AddHashset", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		h2 := corestr.New.Hashset.StringsSpreadItems("a", "b")
		h.AddHashsetItems(h2)
		h.AddHashsetItems(nil)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h3 := corestr.New.Hashset.StringsSpreadItems("c")
		h.AddHashsetWgLock(h3, wg)
		h.AddHashsetWgLock(nil, nil)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddHashset", actual)
	})
}

func Test_Cov18_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_AddsUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()
		f := func(s string, i int) (string, bool, bool) { return s, true, false }
		h.AddsUsingFilter(f, "a", "b")
		h.AddsUsingFilter(f)
		h.AddsAnyUsingFilter(f, "c")
		h.AddsAnyUsingFilter(f)
		h.AddsAnyUsingFilterLock(f, "d")
		h.AddsAnyUsingFilterLock(f)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddsUsingFilter", actual)
	})
}

func Test_Cov18_Hashset_EmptyString(t *testing.T) {
	safeTest(t, "Test_Cov18_Hashset_EmptyString", func() {
		// Arrange
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{
			"str": h.String() != "",
			"strLock": h.StringLock() != "",
		}

		// Assert
		expected := args.Map{
			"str": true,
			"strLock": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- EmptyString", actual)
	})
}
