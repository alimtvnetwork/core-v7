package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// MapStringStringOnce — remaining uncovered methods
// ==========================================================================

func Test_Cov9_MapStringStringOnce_AllKeys(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1", "b": "2"} })
	keys := mo.AllKeys()
	keysCached := mo.AllKeys() // second call uses cache
	actual := args.Map{"len": len(keys), "cachedLen": len(keysCached)}
	expected := args.Map{"len": 2, "cachedLen": 2}
	expected.ShouldBeEqual(t, 0, "AllKeys returns correct value -- returns all keys", actual)
}

func Test_Cov9_MapStringStringOnce_AllKeys_Empty(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })
	keys := mo.AllKeys()
	actual := args.Map{"len": len(keys)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllKeys returns empty -- empty map", actual)
}

func Test_Cov9_MapStringStringOnce_AllValues(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })
	vals := mo.AllValues()
	valsCached := mo.AllValues() // second call uses cache
	actual := args.Map{"len": len(vals), "cachedLen": len(valsCached)}
	expected := args.Map{"len": 1, "cachedLen": 1}
	expected.ShouldBeEqual(t, 0, "AllValues returns non-empty -- returns all values", actual)
}

func Test_Cov9_MapStringStringOnce_AllValues_Empty(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })
	vals := mo.AllValues()
	actual := args.Map{"len": len(vals)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllValues returns empty -- empty map", actual)
}

func Test_Cov9_MapStringStringOnce_AllKeysSorted(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"b": "2", "a": "1"} })
	sorted := mo.AllKeysSorted()
	sortedCached := mo.AllKeysSorted() // second call uses cache
	actual := args.Map{"first": sorted[0], "cachedFirst": sortedCached[0]}
	expected := args.Map{"first": "a", "cachedFirst": "a"}
	expected.ShouldBeEqual(t, 0, "AllKeysSorted returns correct value -- returns sorted keys", actual)
}

func Test_Cov9_MapStringStringOnce_AllKeysSorted_Empty(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })
	sorted := mo.AllKeysSorted()
	actual := args.Map{"len": len(sorted)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllKeysSorted returns empty -- empty map", actual)
}

func Test_Cov9_MapStringStringOnce_AllValuesSorted(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "z", "b": "a"} })
	sorted := mo.AllValuesSorted()
	sortedCached := mo.AllValuesSorted() // second call uses cache
	actual := args.Map{"first": sorted[0], "cachedFirst": sortedCached[0]}
	expected := args.Map{"first": "a", "cachedFirst": "a"}
	expected.ShouldBeEqual(t, 0, "AllValuesSorted returns non-empty -- returns sorted values", actual)
}

func Test_Cov9_MapStringStringOnce_AllValuesSorted_Empty(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })
	sorted := mo.AllValuesSorted()
	actual := args.Map{"len": len(sorted)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllValuesSorted returns empty -- empty map", actual)
}

func Test_Cov9_MapStringStringOnce_GetValue(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"k": "v"} })
	actual := args.Map{"val": mo.GetValue("k"), "missing": mo.GetValue("x")}
	expected := args.Map{"val": "v", "missing": ""}
	expected.ShouldBeEqual(t, 0, "GetValue returns empty -- returns value or empty", actual)
}

func Test_Cov9_MapStringStringOnce_GetValueWithStatus(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"k": "v"} })
	v1, ok1 := mo.GetValueWithStatus("k")
	_, ok2 := mo.GetValueWithStatus("x")
	actual := args.Map{"val": v1, "found": ok1, "notFound": ok2}
	expected := args.Map{"val": "v", "found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "GetValueWithStatus returns non-empty -- returns status", actual)
}

func Test_Cov9_MapStringStringOnce_Has_IsMissing(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"k": "v"} })
	actual := args.Map{"has": mo.Has("k"), "missing": mo.IsMissing("x"), "missingK": mo.IsMissing("k")}
	expected := args.Map{"has": true, "missing": true, "missingK": false}
	expected.ShouldBeEqual(t, 0, "Has returns correct value -- and IsMissing correct", actual)
}

func Test_Cov9_MapStringStringOnce_HasAll(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1", "b": "2"} })
	actual := args.Map{"allFound": mo.HasAll("a", "b"), "oneMissing": mo.HasAll("a", "c")}
	expected := args.Map{"allFound": true, "oneMissing": false}
	expected.ShouldBeEqual(t, 0, "HasAll returns correct value -- correct", actual)
}

func Test_Cov9_MapStringStringOnce_Strings(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })
	s := mo.Strings()
	sCached := mo.Strings() // cached path
	actual := args.Map{"len": len(s), "cachedLen": len(sCached)}
	expected := args.Map{"len": 1, "cachedLen": 1}
	expected.ShouldBeEqual(t, 0, "Strings returns correct value -- returns kv lines", actual)
}

func Test_Cov9_MapStringStringOnce_Strings_Empty(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{} })
	s := mo.Strings()
	actual := args.Map{"len": len(s)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Strings returns empty -- empty map", actual)
}

func Test_Cov9_MapStringStringOnce_JsonStringMust(t *testing.T) {
	mo := coreonce.NewMapStringStringOnce(func() map[string]string { return map[string]string{"a": "1"} })
	s := mo.JsonStringMust()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns json -- valid map", actual)
}

func Test_Cov9_MapStringStringOnce_String(t *testing.T) {
	mo := coreonce.NewMapStringStringOnce(func() map[string]string { return map[string]string{"a": "1"} })
	s := mo.String()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- returns csv-like output", actual)
}

func Test_Cov9_MapStringStringOnce_ValuesPtr(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })
	vp := mo.ValuesPtr()
	actual := args.Map{"notNil": vp != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ValuesPtr returns non-empty -- returns ptr", actual)
}

func Test_Cov9_MapStringStringOnce_Length_Nil(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })
	actual := args.Map{"len": mo.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil map", actual)
}

func Test_Cov9_MapStringStringOnce_IsEqual_BothNil(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })
	actual := args.Map{"eq": mo.IsEqual(nil)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- both nil", actual)
}

func Test_Cov9_MapStringStringOnce_IsEqual_DiffLen(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })
	actual := args.Map{"eq": mo.IsEqual(map[string]string{"a": "1", "b": "2"})}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- diff length", actual)
}

func Test_Cov9_MapStringStringOnce_IsEqual_DiffValue(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })
	actual := args.Map{"eq": mo.IsEqual(map[string]string{"a": "2"})}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- diff value", actual)
}

func Test_Cov9_MapStringStringOnce_IsEqual_MissingKey(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })
	actual := args.Map{"eq": mo.IsEqual(map[string]string{"b": "1"})}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- missing key", actual)
}

func Test_Cov9_MapStringStringOnce_IsEqual_Match(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return map[string]string{"a": "1"} })
	actual := args.Map{"eq": mo.IsEqual(map[string]string{"a": "1"})}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- match", actual)
}

func Test_Cov9_MapStringStringOnce_Serialize(t *testing.T) {
	mo := coreonce.NewMapStringStringOnce(func() map[string]string { return map[string]string{"a": "1"} })
	b, err := mo.Serialize()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- returns bytes", actual)
}

func Test_Cov9_MapStringStringOnce_UnmarshalJSON(t *testing.T) {
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil })
	err := mo.UnmarshalJSON([]byte(`{"x":"y"}`))
	actual := args.Map{"noErr": err == nil, "val": mo.GetValue("x")}
	expected := args.Map{"noErr": true, "val": "y"}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns correct value -- parses json", actual)
}

// ==========================================================================
// StringsOnce — remaining uncovered methods
// ==========================================================================

func Test_Cov9_StringsOnce_UniqueMap(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b", "a"} })
	m := so.UniqueMap()
	mCached := so.UniqueMap() // cached path
	actual := args.Map{"len": len(m), "cachedLen": len(mCached), "hasA": m["a"]}
	expected := args.Map{"len": 2, "cachedLen": 2, "hasA": true}
	expected.ShouldBeEqual(t, 0, "UniqueMap returns correct value -- returns unique map", actual)
}

func Test_Cov9_StringsOnce_UniqueMap_Nil(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return nil })
	m := so.UniqueMap()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "UniqueMap returns empty -- nil values", actual)
}

func Test_Cov9_StringsOnce_UniqueMapLock(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"x"} })
	m := so.UniqueMapLock()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "UniqueMapLock returns correct value -- returns locked unique map", actual)
}

func Test_Cov9_StringsOnce_CsvLines(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	lines := so.CsvLines()
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CsvLines returns correct value -- returns csv lines", actual)
}

func Test_Cov9_StringsOnce_Csv(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	csv := so.Csv()
	actual := args.Map{"notEmpty": csv != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Csv returns correct value -- returns csv string", actual)
}

func Test_Cov9_StringsOnce_Sorted(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"b", "a"} })
	s := so.Sorted()
	sCached := so.Sorted() // cached path
	actual := args.Map{"first": s[0], "cachedFirst": sCached[0]}
	expected := args.Map{"first": "a", "cachedFirst": "a"}
	expected.ShouldBeEqual(t, 0, "Sorted returns non-empty -- returns sorted values", actual)
}

func Test_Cov9_StringsOnce_RangesMap(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"x", "y"} })
	m := so.RangesMap()
	actual := args.Map{"xIdx": m["x"], "yIdx": m["y"]}
	expected := args.Map{"xIdx": 0, "yIdx": 1}
	expected.ShouldBeEqual(t, 0, "RangesMap returns correct value -- returns index map", actual)
}

func Test_Cov9_StringsOnce_RangesMap_Empty(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{} })
	m := so.RangesMap()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangesMap returns empty -- empty slice", actual)
}

func Test_Cov9_StringsOnce_HasAll(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	actual := args.Map{"all": so.HasAll("a", "b"), "missing": so.HasAll("a", "c")}
	expected := args.Map{"all": true, "missing": false}
	expected.ShouldBeEqual(t, 0, "HasAll returns correct value -- correct", actual)
}

func Test_Cov9_StringsOnce_JsonStringMust(t *testing.T) {
	so := coreonce.NewStringsOnce(func() []string { return []string{"a"} })
	s := so.JsonStringMust()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns correct value -- returns json", actual)
}

func Test_Cov9_StringsOnce_String(t *testing.T) {
	so := coreonce.NewStringsOnce(func() []string { return []string{"a", "b"} })
	s := so.String()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- returns csv output", actual)
}

func Test_Cov9_StringsOnce_Length_Nil(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return nil })
	actual := args.Map{"len": so.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil values", actual)
}

func Test_Cov9_StringsOnce_SafeStrings(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return nil })
	ss := so.SafeStrings()
	actual := args.Map{"len": len(ss)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeStrings returns empty slice -- nil", actual)
}

func Test_Cov9_StringsOnce_SafeStrings_NonEmpty(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"x"} })
	ss := so.SafeStrings()
	actual := args.Map{"len": len(ss)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeStrings returns values -- non-nil", actual)
}

// ==========================================================================
// IntegersOnce — remaining uncovered methods
// ==========================================================================

func Test_Cov9_IntegersOnce_Sorted(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{3, 1, 2} })
	s := io.Sorted()
	sCached := io.Sorted() // cached path
	actual := args.Map{"first": s[0], "cachedFirst": sCached[0]}
	expected := args.Map{"first": 1, "cachedFirst": 1}
	expected.ShouldBeEqual(t, 0, "Sorted returns correct value -- returns sorted ints", actual)
}

func Test_Cov9_IntegersOnce_RangesMap(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{10, 20} })
	m := io.RangesMap()
	actual := args.Map{"v10": m[10], "v20": m[20]}
	expected := args.Map{"v10": 0, "v20": 1}
	expected.ShouldBeEqual(t, 0, "RangesMap returns correct value -- returns index map", actual)
}

func Test_Cov9_IntegersOnce_RangesMap_Empty(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := io.RangesMap()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangesMap returns empty -- empty slice", actual)
}

func Test_Cov9_IntegersOnce_RangesBoolMap(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{5, 10} })
	m := io.RangesBoolMap()
	actual := args.Map{"has5": m[5], "has10": m[10]}
	expected := args.Map{"has5": true, "has10": true}
	expected.ShouldBeEqual(t, 0, "RangesBoolMap returns correct value -- returns bool map", actual)
}

func Test_Cov9_IntegersOnce_RangesBoolMap_Empty(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := io.RangesBoolMap()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangesBoolMap returns empty -- empty slice", actual)
}

func Test_Cov9_IntegersOnce_UniqueMap(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2, 1} })
	m := io.UniqueMap()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UniqueMap returns correct value -- returns unique map", actual)
}

func Test_Cov9_IntegersOnce_UniqueMap_Empty(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{} })
	m := io.UniqueMap()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "UniqueMap returns empty -- empty slice", actual)
}

func Test_Cov9_IntegersOnce_IsEqual_DiffCount(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2, 1} })
	actual := args.Map{"eq": io.IsEqual(1, 2)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- diff count", actual)
}

func Test_Cov9_IntegersOnce_IsEqual_DiffValues(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })
	actual := args.Map{"eq": io.IsEqual(1, 3)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns non-empty -- diff values", actual)
}

func Test_Cov9_IntegersOnce_Aliases(t *testing.T) {
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })
	actual := args.Map{
		"valuesLen":  len(io.Values()),
		"executeLen": len(io.Execute()),
		"integersLen": len(io.Integers()),
		"sliceLen":   len(io.Slice()),
		"listLen":    len(io.List()),
		"strNotEmpty": io.String() != "",
	}
	expected := args.Map{
		"valuesLen":  2,
		"executeLen": 2,
		"integersLen": 2,
		"sliceLen":   2,
		"listLen":    2,
		"strNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce returns correct value -- aliases work", actual)
}

// ==========================================================================
// StringsOnce — IsEqual with excess value
// ==========================================================================

func Test_Cov9_StringsOnce_IsEqual_DiffCount(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	actual := args.Map{"eq": so.IsEqual("a")}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- diff count", actual)
}

func Test_Cov9_StringsOnce_IsEqual_DiffValues(t *testing.T) {
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "b"} })
	actual := args.Map{"eq": so.IsEqual("a", "c")}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns non-empty -- diff values", actual)
}

// ==========================================================================
// ByteOnce — remaining uncovered methods
// ==========================================================================

func Test_Cov9_ByteOnce_UnmarshalJSON(t *testing.T) {
	bo := coreonce.NewByteOncePtr(func() byte { return 0 })
	err := bo.UnmarshalJSON([]byte(`65`))
	actual := args.Map{"noErr": err == nil, "val": int(bo.Value())}
	expected := args.Map{"noErr": true, "val": 65}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns correct value -- parses byte", actual)
}

func Test_Cov9_ByteOnce_Int(t *testing.T) {
	bo := coreonce.NewByteOncePtr(func() byte { return 42 })
	actual := args.Map{"int": bo.Int()}
	expected := args.Map{"int": 42}
	expected.ShouldBeEqual(t, 0, "Int returns correct value -- returns int value", actual)
}

func Test_Cov9_ByteOnce_IsEmpty_IsZero(t *testing.T) {
	bo := coreonce.NewByteOncePtr(func() byte { return 0 })
	actual := args.Map{"empty": bo.IsEmpty(), "zero": bo.IsZero()}
	expected := args.Map{"empty": true, "zero": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- and IsZero on zero byte", actual)
}

func Test_Cov9_ByteOnce_IsPositive_IsNegative(t *testing.T) {
	bo := coreonce.NewByteOncePtr(func() byte { return 5 })
	actual := args.Map{"pos": bo.IsPositive(), "neg": bo.IsNegative()}
	expected := args.Map{"pos": true, "neg": false}
	expected.ShouldBeEqual(t, 0, "IsPositive returns correct value -- and IsNegative on positive byte", actual)
}

// ==========================================================================
// IntegerOnce — remaining uncovered methods
// ==========================================================================

func Test_Cov9_IntegerOnce_UnmarshalJSON(t *testing.T) {
	io := coreonce.NewIntegerOncePtr(func() int { return 0 })
	err := io.UnmarshalJSON([]byte(`42`))
	actual := args.Map{"noErr": err == nil, "val": io.Value()}
	expected := args.Map{"noErr": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns correct value -- parses int", actual)
}

func Test_Cov9_IntegerOnce_ComparisonMethods(t *testing.T) {
	io := coreonce.NewIntegerOncePtr(func() int { return 5 })
	actual := args.Map{
		"aboveZero":     io.IsAboveZero(),
		"aboveEqZero":   io.IsAboveEqualZero(),
		"lessThanZero":  io.IsLessThanZero(),
		"lessThanEqZero": io.IsLessThanEqualZero(),
		"above3":        io.IsAbove(3),
		"aboveEq5":      io.IsAboveEqual(5),
		"lessThan10":    io.IsLessThan(10),
		"lessThanEq5":   io.IsLessThanEqual(5),
		"validIndex":    io.IsValidIndex(),
		"invalidIndex":  io.IsInvalidIndex(),
		"positive":      io.IsPositive(),
		"negative":      io.IsNegative(),
	}
	expected := args.Map{
		"aboveZero":     true,
		"aboveEqZero":   true,
		"lessThanZero":  false,
		"lessThanEqZero": false,
		"above3":        true,
		"aboveEq5":      true,
		"lessThan10":    true,
		"lessThanEq5":   true,
		"validIndex":    true,
		"invalidIndex":  false,
		"positive":      true,
		"negative":      false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce returns correct value -- comparison methods", actual)
}

// ==========================================================================
// BoolOnce — remaining uncovered methods
// ==========================================================================

func Test_Cov9_BoolOnce_UnmarshalJSON(t *testing.T) {
	bo := coreonce.NewBoolOncePtr(func() bool { return false })
	err := bo.UnmarshalJSON([]byte(`true`))
	actual := args.Map{"noErr": err == nil, "val": bo.Value()}
	expected := args.Map{"noErr": true, "val": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns correct value -- parses bool", actual)
}

// ==========================================================================
// BytesOnce — remaining uncovered methods
// ==========================================================================

func Test_Cov9_BytesOnce_UnmarshalJSON(t *testing.T) {
	bo := coreonce.NewBytesOncePtr(func() []byte { return nil })
	err := bo.UnmarshalJSON([]byte(`"aGVsbG8="`))
	actual := args.Map{"noErr": err == nil, "hasData": len(bo.Value()) > 0}
	expected := args.Map{"noErr": true, "hasData": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns correct value -- parses bytes", actual)
}

func Test_Cov9_BytesOnce_NilInitializerFunc(t *testing.T) {
	bo := &coreonce.BytesOnce{}
	val := bo.Value()
	actual := args.Map{"isNil": val == nil, "len": bo.Length(), "empty": bo.IsEmpty()}
	expected := args.Map{"isNil": true, "len": 0, "empty": true}
	expected.ShouldBeEqual(t, 0, "BytesOnce returns nil -- nil initializer", actual)
}

// ==========================================================================
// StringOnce — remaining uncovered methods
// ==========================================================================

func Test_Cov9_StringOnce_SplitLeftRight_Single(t *testing.T) {
	so := coreonce.NewStringOncePtr(func() string { return "only" })
	left, right := so.SplitLeftRight("=")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "only", "right": ""}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns correct value -- single element", actual)
}

func Test_Cov9_StringOnce_SplitLeftRight_Two(t *testing.T) {
	so := coreonce.NewStringOncePtr(func() string { return "k=v" })
	left, right := so.SplitLeftRight("=")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "k", "right": "v"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns correct value -- two parts", actual)
}

func Test_Cov9_StringOnce_SplitLeftRightTrim(t *testing.T) {
	so := coreonce.NewStringOncePtr(func() string { return " k = v " })
	left, right := so.SplitLeftRightTrim("=")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "k", "right": "v"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrim returns correct value -- trims spaces", actual)
}

func Test_Cov9_StringOnce_MoreMethods(t *testing.T) {
	so := coreonce.NewStringOncePtr(func() string { return "hello world" })
	actual := args.Map{
		"hasPrefix":  so.HasPrefix("hello"),
		"hasSuffix":  so.HasSuffix("world"),
		"startsWith": so.IsStartsWith("hello"),
		"endsWith":   so.IsEndsWith("world"),
		"contains":   so.IsContains("lo wo"),
		"isEmpty":    so.IsEmpty(),
		"isEmptyWs":  so.IsEmptyOrWhitespace(),
		"bytesLen":   len(so.Bytes()),
		"errNotNil":  so.Error() != nil,
		"vp":         *so.ValuePtr() == "hello world",
	}
	expected := args.Map{
		"hasPrefix":  true,
		"hasSuffix":  true,
		"startsWith": true,
		"endsWith":   true,
		"contains":   true,
		"isEmpty":    false,
		"isEmptyWs":  false,
		"bytesLen":   11,
		"errNotNil":  true,
		"vp":         true,
	}
	expected.ShouldBeEqual(t, 0, "StringOnce returns correct value -- methods", actual)
}
