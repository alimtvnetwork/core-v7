package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// IntegerOnce — comparison methods
// ==========================================================================

func Test_IntegerOnce_Comparisons_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return 5 })

	// Act
	actual := args.Map{
		"isAbove3":         io.IsAbove(3),
		"isAbove5":         io.IsAbove(5),
		"isAboveEqual5":    io.IsAboveEqual(5),
		"isLessThan10":     io.IsLessThan(10),
		"isLessThan5":      io.IsLessThan(5),
		"isLessThanEqual5": io.IsLessThanEqual(5),
		"isValidIndex":     io.IsValidIndex(),
		"isInvalidIndex":   io.IsInvalidIndex(),
		"execute":          io.Execute(),
		"string":           io.String(),
	}

	// Assert
	expected := args.Map{
		"isAbove3": true, "isAbove5": false,
		"isAboveEqual5": true, "isLessThan10": true,
		"isLessThan5": false, "isLessThanEqual5": true,
		"isValidIndex": true, "isInvalidIndex": false,
		"execute": 5, "string": "5",
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce comparisons -- value 5", actual)
}

func Test_IntegerOnce_Negative(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return -1 })

	// Act
	actual := args.Map{
		"isLessThanZero":      io.IsLessThanZero(),
		"isLessThanEqualZero": io.IsLessThanEqualZero(),
		"isInvalidIndex":      io.IsInvalidIndex(),
	}

	// Assert
	expected := args.Map{
		"isLessThanZero": true, "isLessThanEqualZero": true,
		"isInvalidIndex": true,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce negative comparisons -- value -1", actual)
}

func Test_IntegerOnce_Serialize(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return 42 })
	bytes, err := io.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce Serialize -- 42", actual)
}

func Test_IntegerOnce_MarshalUnmarshal_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return 42 })
	mb, _ := io.MarshalJSON()
	io2 := coreonce.NewIntegerOnce(func() int { return 0 })
	err := io2.UnmarshalJSON(mb)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"val": io2.Value(),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce Marshal/Unmarshal roundtrip -- 42", actual)
}

// ==========================================================================
// IntegersOnce — Sorted, RangesMap, IsEqual, etc
// ==========================================================================

func Test_IntegersOnce_Sorted(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOnce(func() []int { return []int{3, 1, 2} })
	sorted := io.Sorted()
	// call again for cached path
	sorted2 := io.Sorted()

	// Act
	actual := args.Map{
		"first": sorted[0], "last": sorted[2],
		"cached": sorted2[0],
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"last": 3,
		"cached": 1,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce Sorted returns sorted -- 3,1,2", actual)
}

func Test_IntegersOnce_RangesMap_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOnce(func() []int { return []int{10, 20} })
	rm := io.RangesMap()
	rbm := io.RangesBoolMap()
	um := io.UniqueMap()

	// Act
	actual := args.Map{
		"rmLen": len(rm), "rbmLen": len(rbm), "umLen": len(um),
	}

	// Assert
	expected := args.Map{
		"rmLen": 2,
		"rbmLen": 2,
		"umLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce maps -- 2 items", actual)
}

func Test_IntegersOnce_RangesMap_Empty_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOnce(func() []int { return []int{} })
	rm := io.RangesMap()
	rbm := io.RangesBoolMap()
	um := io.UniqueMap()

	// Act
	actual := args.Map{
		"rmLen": len(rm),
		"rbmLen": len(rbm),
		"umLen": len(um),
	}

	// Assert
	expected := args.Map{
		"rmLen": 0,
		"rbmLen": 0,
		"umLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce maps empty -- empty", actual)
}

func Test_IntegersOnce_IsEqual(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOnce(func() []int { return []int{1, 2} })

	// Act
	actual := args.Map{
		"same":    io.IsEqual(1, 2),
		"diff":    io.IsEqual(1, 3),
		"diffLen": io.IsEqual(1),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
		"diffLen": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce IsEqual -- various", actual)
}

func Test_IntegersOnce_Aliases_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOnce(func() []int { return []int{1} })

	// Act
	actual := args.Map{
		"values":   len(io.Values()),
		"execute":  len(io.Execute()),
		"integers": len(io.Integers()),
		"slice":    len(io.Slice()),
		"list":     len(io.List()),
		"string":   io.String() != "",
		"isZero":   io.IsZero(),
	}

	// Assert
	expected := args.Map{
		"values": 1, "execute": 1, "integers": 1,
		"slice": 1, "list": 1, "string": true, "isZero": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce aliases -- 1 item", actual)
}

func Test_IntegersOnce_Serialize(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOnce(func() []int { return []int{1, 2} })
	bytes, err := io.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce Serialize -- valid", actual)
}

func Test_IntegersOnce_MarshalUnmarshal(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOnce(func() []int { return []int{1, 2} })
	mb, _ := io.MarshalJSON()
	io2 := coreonce.NewIntegersOnce(func() []int { return nil })
	err := io2.UnmarshalJSON(mb)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": len(io2.Value()),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce Marshal/Unmarshal -- roundtrip", actual)
}

// ==========================================================================
// StringsOnce — extensive
// ==========================================================================

func Test_StringsOnce_UniqueMap(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a", "b", "a"} })
	um := so.UniqueMap()
	// cached path
	um2 := so.UniqueMap()

	// Act
	actual := args.Map{
		"len": len(um),
		"cached": len(um2),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"cached": 2,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce UniqueMap -- 2 unique", actual)
}

func Test_StringsOnce_UniqueMapLock_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a", "b"} })
	um := so.UniqueMapLock()

	// Act
	actual := args.Map{"len": len(um)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsOnce UniqueMapLock -- 2 items", actual)
}

func Test_StringsOnce_Has(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a", "b"} })

	// Act
	actual := args.Map{
		"hasA":      so.Has("a"),
		"hasC":      so.Has("c"),
		"containsA": so.IsContains("a"),
	}

	// Assert
	expected := args.Map{
		"hasA": true,
		"hasC": false,
		"containsA": true,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce Has/IsContains -- a and c", actual)
}

func Test_StringsOnce_HasAll(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a", "b", "c"} })

	// Act
	actual := args.Map{
		"allAB": so.HasAll("a", "b"),
		"allAD": so.HasAll("a", "d"),
	}

	// Assert
	expected := args.Map{
		"allAB": true,
		"allAD": false,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce HasAll -- present and missing", actual)
}

func Test_StringsOnce_Sorted(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"c", "a", "b"} })
	sorted := so.Sorted()
	sorted2 := so.Sorted()

	// Act
	actual := args.Map{
		"first": sorted[0],
		"cached": sorted2[0],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"cached": "a",
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce Sorted -- c,a,b", actual)
}

func Test_StringsOnce_RangesMap_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"x", "y"} })
	rm := so.RangesMap()

	// Act
	actual := args.Map{"len": len(rm)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsOnce RangesMap -- 2 items", actual)
}

func Test_StringsOnce_RangesMap_Empty_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{} })
	rm := so.RangesMap()

	// Act
	actual := args.Map{"len": len(rm)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsOnce RangesMap empty -- empty", actual)
}

func Test_StringsOnce_Csv(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a", "b"} })

	// Act
	actual := args.Map{
		"csv":     so.Csv() != "",
		"options": so.CsvOptions() != "",
		"lines":   len(so.CsvLines()) > 0,
		"string":  so.String() != "",
	}

	// Assert
	expected := args.Map{
		"csv": true,
		"options": true,
		"lines": true,
		"string": true,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce Csv methods -- 2 items", actual)
}

func Test_StringsOnce_SafeStrings_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a"} })

	// Act
	actual := args.Map{"len": len(so.SafeStrings())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringsOnce SafeStrings -- 1 item", actual)
}

func Test_StringsOnce_SafeStrings_Empty_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return nil })

	// Act
	actual := args.Map{"len": len(so.SafeStrings())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsOnce SafeStrings empty -- nil", actual)
}

func Test_StringsOnce_IsEqual(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a", "b"} })

	// Act
	actual := args.Map{
		"same":    so.IsEqual("a", "b"),
		"diff":    so.IsEqual("a", "c"),
		"diffLen": so.IsEqual("a"),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
		"diffLen": false,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce IsEqual -- various", actual)
}

func Test_StringsOnce_Aliases(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a"} })

	// Act
	actual := args.Map{
		"strings": len(so.Strings()),
		"list":    len(so.List()),
		"values":  len(so.Values()),
		"valPtr":  len(so.ValuesPtr()),
		"execute": len(so.Execute()),
	}

	// Assert
	expected := args.Map{
		"strings": 1,
		"list": 1,
		"values": 1,
		"valPtr": 1,
		"execute": 1,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce aliases -- 1 item", actual)
}

func Test_StringsOnce_Serialize(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a"} })
	bytes, err := so.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce Serialize -- valid", actual)
}

func Test_StringsOnce_JsonStringMust(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"a"} })
	jsonStr := so.JsonStringMust()

	// Act
	actual := args.Map{"notEmpty": jsonStr != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsOnce JsonStringMust -- valid", actual)
}

func Test_StringsOnce_MarshalUnmarshal(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{"x"} })
	mb, _ := so.MarshalJSON()
	so2 := coreonce.NewStringsOnce(func() []string { return nil })
	err := so2.UnmarshalJSON(mb)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": len(so2.Value()),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce Marshal/Unmarshal -- roundtrip", actual)
}

// ==========================================================================
// MapStringStringOnce — AllKeys, AllValues, IsEqual
// ==========================================================================

func Test_MapStringStringOnce_AllKeys(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})
	keys := mso.AllKeys()
	keys2 := mso.AllKeys() // cached
	vals := mso.AllValues()
	vals2 := mso.AllValues() // cached

	// Act
	actual := args.Map{
		"keysLen": len(keys), "cachedKeysLen": len(keys2),
		"valsLen": len(vals), "cachedValsLen": len(vals2),
	}

	// Assert
	expected := args.Map{
		"keysLen": 2, "cachedKeysLen": 2,
		"valsLen": 2, "cachedValsLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce AllKeys/AllValues -- 2 entries", actual)
}

func Test_MapStringStringOnce_AllKeysSorted(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"b": "2", "a": "1"}
	})
	ks := mso.AllKeysSorted()
	ks2 := mso.AllKeysSorted() // cached
	vs := mso.AllValuesSorted()
	vs2 := mso.AllValuesSorted() // cached

	// Act
	actual := args.Map{
		"first": ks[0], "cached": ks2[0],
		"vsLen": len(vs), "vsCached": len(vs2),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"cached": "a",
		"vsLen": 2,
		"vsCached": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce AllKeysSorted -- b,a sorted", actual)
}

func Test_MapStringStringOnce_GetValue(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	val := mso.GetValue("k")
	val2, has := mso.GetValueWithStatus("k")
	_, missing := mso.GetValueWithStatus("nope")

	// Act
	actual := args.Map{
		"val": val,
		"val2": val2,
		"has": has,
		"missing": missing,
	}

	// Assert
	expected := args.Map{
		"val": "v",
		"val2": "v",
		"has": true,
		"missing": false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce GetValue -- k", actual)
}

func Test_MapStringStringOnce_IsEqual(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1"}
	})

	// Act
	actual := args.Map{
		"same":     mso.IsEqual(map[string]string{"a": "1"}),
		"diffVal":  mso.IsEqual(map[string]string{"a": "2"}),
		"diffKey":  mso.IsEqual(map[string]string{"b": "1"}),
		"diffLen":  mso.IsEqual(map[string]string{"a": "1", "b": "2"}),
		"nil":      mso.IsEqual(nil),
	}

	// Assert
	expected := args.Map{
		"same": true, "diffVal": false, "diffKey": false,
		"diffLen": false, "nil": false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce IsEqual -- various", actual)
}

func Test_MapStringStringOnce_Strings(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	strs := mso.Strings()
	strs2 := mso.Strings() // cached

	// Act
	actual := args.Map{
		"len": len(strs),
		"cached": len(strs2),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"cached": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce Strings -- 1 entry", actual)
}

func Test_MapStringStringOnce_Strings_Empty_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string { return map[string]string{} })
	strs := mso.Strings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce Strings empty -- empty", actual)
}

func Test_MapStringStringOnce_Aliases(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"k": "v"}
	})

	// Act
	actual := args.Map{
		"list":    len(mso.List()),
		"items":   len(mso.ItemsMap()),
		"values":  len(mso.Values()),
		"valPtr":  len(*mso.ValuesPtr()),
		"execute": len(mso.Execute()),
		"hasAny":  mso.HasAnyItem(),
		"has":     mso.Has("k"),
		"hasAll":  mso.HasAll("k"),
		"string":  mso.String() != "",
		"jsonStr": mso.JsonStringMust() != "",
	}

	// Assert
	expected := args.Map{
		"list": 1, "items": 1, "values": 1, "valPtr": 1,
		"execute": 1, "hasAny": true, "has": true, "hasAll": true,
		"string": true, "jsonStr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce aliases -- 1 entry", actual)
}

func Test_MapStringStringOnce_HasAll_Missing_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"a": "1"}
	})

	// Act
	actual := args.Map{"hasAll": mso.HasAll("a", "b")}

	// Assert
	expected := args.Map{"hasAll": false}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce HasAll missing -- b missing", actual)
}

func Test_MapStringStringOnce_AllKeys_Empty_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string { return map[string]string{} })

	// Act
	actual := args.Map{
		"keys":       len(mso.AllKeys()),
		"vals":       len(mso.AllValues()),
		"keysSorted": len(mso.AllKeysSorted()),
		"valsSorted": len(mso.AllValuesSorted()),
	}

	// Assert
	expected := args.Map{
		"keys": 0,
		"vals": 0,
		"keysSorted": 0,
		"valsSorted": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce empty maps -- empty", actual)
}

// ==========================================================================
// AnyOnce — Serialize, IsInitialized, IsStringEmpty
// ==========================================================================

func Test_AnyOnce_Serialize_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return "hello" })
	bytes, err := ao.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce Serialize -- valid", actual)
}

func Test_AnyOnce_SerializeMust_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return 42 })
	bytes := ao.SerializeMust()

	// Act
	actual := args.Map{"hasBytes": len(bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyOnce SerializeMust -- 42", actual)
}

func Test_AnyOnce_SerializeSkipExistingError_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return "test" })
	bytes, err := ao.SerializeSkipExistingError()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce SerializeSkipExistingError -- valid", actual)
}

func Test_AnyOnce_IsInitialized_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return nil })
	before := ao.IsInitialized()
	ao.Value()
	after := ao.IsInitialized()

	// Act
	actual := args.Map{
		"before": before,
		"after": after,
	}

	// Assert
	expected := args.Map{
		"before": false,
		"after": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce IsInitialized -- before and after", actual)
}

func Test_AnyOnce_IsStringEmpty_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return nil })

	// Act
	actual := args.Map{
		"isEmpty": ao.IsStringEmpty(),
		"isWs":    ao.IsStringEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"isWs": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce IsStringEmpty -- nil value", actual)
}

func Test_AnyOnce_CastSuccess(t *testing.T) {
	// Arrange
	aoStr := coreonce.NewAnyOnce(func() any { return "hello" })
	aoStrings := coreonce.NewAnyOnce(func() any { return []string{"a"} })
	aoMap := coreonce.NewAnyOnce(func() any { return map[string]string{"k": "v"} })
	aoMapAny := coreonce.NewAnyOnce(func() any { return map[string]any{"k": 1} })
	aoBytes := coreonce.NewAnyOnce(func() any { return []byte("hi") })

	s, sOk := aoStr.CastValueString()
	ss, ssOk := aoStrings.CastValueStrings()
	m, mOk := aoMap.CastValueHashmapMap()
	ma, maOk := aoMapAny.CastValueMapStringAnyMap()
	b, bOk := aoBytes.CastValueBytes()

	// Act
	actual := args.Map{
		"s": s, "sOk": sOk, "ssLen": len(ss), "ssOk": ssOk,
		"mLen": len(m), "mOk": mOk, "maLen": len(ma), "maOk": maOk,
		"bLen": len(b), "bOk": bOk,
	}

	// Assert
	expected := args.Map{
		"s": "hello", "sOk": true, "ssLen": 1, "ssOk": true,
		"mLen": 1, "mOk": true, "maLen": 1, "maOk": true,
		"bLen": 2, "bOk": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce Cast success -- all types", actual)
}

// ==========================================================================
// BoolOnce — Execute, String, Serialize
// ==========================================================================

func Test_BoolOnce_Methods(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return true })

	// Act
	actual := args.Map{
		"execute":  bo.Execute(),
		"string":   bo.String(),
	}

	// Assert
	expected := args.Map{
		"execute": true,
		"string": "true",
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce Execute and String -- true", actual)
}

func Test_BoolOnce_Serialize_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return false })
	bytes, err := bo.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce Serialize -- false", actual)
}

func Test_BoolOnce_MarshalUnmarshal_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	bo := coreonce.NewBoolOnce(func() bool { return true })
	mb, _ := bo.MarshalJSON()
	bo2 := coreonce.NewBoolOnce(func() bool { return false })
	err := bo2.UnmarshalJSON(mb)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"val": bo2.Value(),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"val": true,
	}
	expected.ShouldBeEqual(t, 0, "BoolOnce Marshal/Unmarshal -- roundtrip", actual)
}

// ==========================================================================
// ByteOnce — Int, Execute, String, Serialize
// ==========================================================================

func Test_ByteOnce_Methods_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	bo := coreonce.NewByteOnce(func() byte { return 42 })

	// Act
	actual := args.Map{
		"int":     bo.Int(),
		"execute": int(bo.Execute()),
		"string":  bo.String(),
	}

	// Assert
	expected := args.Map{
		"int": 42,
		"execute": 42,
		"string": "42",
	}
	expected.ShouldBeEqual(t, 0, "ByteOnce methods -- 42", actual)
}

func Test_ByteOnce_Serialize_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	bo := coreonce.NewByteOnce(func() byte { return 1 })
	bytes, err := bo.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "ByteOnce Serialize -- 1", actual)
}

func Test_ByteOnce_MarshalUnmarshal_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	bo := coreonce.NewByteOnce(func() byte { return 5 })
	mb, _ := bo.MarshalJSON()
	bo2 := coreonce.NewByteOnce(func() byte { return 0 })
	err := bo2.UnmarshalJSON(mb)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"val": int(bo2.Value()),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"val": 5,
	}
	expected.ShouldBeEqual(t, 0, "ByteOnce Marshal/Unmarshal -- roundtrip", actual)
}

// ==========================================================================
// BytesOnce — Execute, nil init, String, Length
// ==========================================================================

func Test_BytesOnce_NilInit(t *testing.T) {
	// Arrange
	bo := &coreonce.BytesOnce{}
	val := bo.Value()

	// Act
	actual := args.Map{
		"len": len(val),
		"isEmpty": bo.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce nil init -- no func", actual)
}

func Test_BytesOnce_Execute(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesOnce(func() []byte { return []byte("hi") })

	// Act
	actual := args.Map{
		"execute": string(bo.Execute()),
		"string":  bo.String(),
		"length":  bo.Length(),
	}

	// Assert
	expected := args.Map{
		"execute": "hi",
		"string": "hi",
		"length": 2,
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce Execute/String/Length -- hi", actual)
}

func Test_BytesOnce_Serialize(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesOnce(func() []byte { return []byte("test") })
	bytes, err := bo.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce Serialize -- test", actual)
}

func Test_BytesOnce_MarshalUnmarshal_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesOnce(func() []byte { return []byte("ab") })
	mb, _ := bo.MarshalJSON()
	bo2 := coreonce.NewBytesOnce(func() []byte { return nil })
	err := bo2.UnmarshalJSON(mb)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": len(bo2.Value()),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce Marshal/Unmarshal -- roundtrip", actual)
}

// ==========================================================================
// BytesErrorOnce — Deserialize, IsStringEmpty, etc
// ==========================================================================

func Test_BytesErrorOnce_Deserialize_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte(`"hello"`), nil })
	var result string
	err := beo.Deserialize(&result)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"result": result,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce Deserialize -- valid", actual)
}

func Test_BytesErrorOnce_StringMethods(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("ab"), nil })

	// Act
	actual := args.Map{
		"isStringEmpty": beo.IsStringEmpty(),
		"isStringWs":    beo.IsStringEmptyOrWhitespace(),
		"isBytesEmpty":  beo.IsBytesEmpty(),
		"isEmptyBytes":  beo.IsEmptyBytes(),
		"hasAny":        beo.HasAnyItem(),
		"isDefined":     beo.IsDefined(),
		"isInit":        beo.IsInitialized(),
	}
	// IsInitialized may be false until Value() called - call it
	beo.Value()
	actual["isInitAfter"] = beo.IsInitialized()

	// Assert
	expected := args.Map{
		"isStringEmpty": false, "isStringWs": false,
		"isBytesEmpty": false, "isEmptyBytes": false,
		"hasAny": true, "isDefined": true, "isInit": true,
		"isInitAfter": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce string methods -- ab", actual)
}

func Test_BytesErrorOnce_HasIssuesOrEmpty(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("ok"), nil })

	// Act
	actual := args.Map{
		"hasIssues": beo.HasIssuesOrEmpty(),
		"hasSafe":   beo.HasSafeItems(),
	}

	// Assert
	expected := args.Map{
		"hasIssues": false,
		"hasSafe": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce HasIssuesOrEmpty -- valid", actual)
}

func Test_BytesErrorOnce_ValueWithError_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("ok"), nil })
	val, err := beo.ValueWithError()

	// Act
	actual := args.Map{
		"hasVal": len(val) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasVal": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce ValueWithError -- valid", actual)
}

func Test_BytesErrorOnce_Serialize_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("ok"), nil })
	bytes, err := beo.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce Serialize -- valid", actual)
}

func Test_BytesErrorOnce_MarshalJSON_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("ok"), nil })
	bytes, err := beo.MarshalJSON()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce MarshalJSON -- valid", actual)
}

// ==========================================================================
// AnyErrorOnce — ValueStringOnly, SafeString, IsStringEmpty
// ==========================================================================

func Test_AnyErrorOnce_ValueStringOnly(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })

	// Act
	actual := args.Map{
		"valueStringOnly": aeo.ValueStringOnly() != "",
		"safeString":      aeo.SafeString() != "",
		"valueStringMust": aeo.ValueStringMust() != "",
	}

	// Assert
	expected := args.Map{
		"valueStringOnly": true, "safeString": true, "valueStringMust": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce string aliases -- hello", actual)
}

func Test_AnyErrorOnce_StringMethods(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "hello", nil })

	// Act
	actual := args.Map{
		"isStrEmpty": aeo.IsStringEmpty(),
		"isStrWs":    aeo.IsStringEmptyOrWhitespace(),
		"string":     aeo.String() != "",
	}

	// Assert
	expected := args.Map{
		"isStrEmpty": false,
		"isStrWs": false,
		"string": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce string methods -- hello", actual)
}

func Test_AnyErrorOnce_ValueOnly(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return 42, nil })
	val := aeo.ValueOnly()

	// Act
	actual := args.Map{
		"val": val,
		"isInit": aeo.IsInitialized(),
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"isInit": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueOnly -- 42", actual)
}

func Test_AnyErrorOnce_Serialize_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "test", nil })
	bytes, err := aeo.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce Serialize -- valid", actual)
}

func Test_AnyErrorOnce_SerializeSkipExistingError_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return 42, nil })
	bytes, err := aeo.SerializeSkipExistingError()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce SerializeSkipExistingError -- 42", actual)
}

func Test_AnyErrorOnce_SerializeMust(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "ok", nil })
	bytes := aeo.SerializeMust()

	// Act
	actual := args.Map{"hasBytes": len(bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce SerializeMust -- valid", actual)
}

func Test_AnyErrorOnce_ValueMust_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "ok", nil })
	val := aeo.ValueMust()

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "ok"}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueMust -- ok", actual)
}

func Test_AnyErrorOnce_ExecuteMust_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return 42, nil })
	val := aeo.ExecuteMust()

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ExecuteMust -- 42", actual)
}

func Test_AnyErrorOnce_ValueString_Cached_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return 42, nil })
	first, _ := aeo.ValueString()
	second, _ := aeo.ValueString() // cached

	// Act
	actual := args.Map{"same": first == second}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueString cached -- 42", actual)
}

// ==========================================================================
// ErrorOnce — MarshalJSON, UnmarshalJSON, Serialize
// ==========================================================================

func Test_ErrorOnce_MarshalJSON_NoError(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })
	mb, err := eo.MarshalJSON()

	// Act
	actual := args.Map{
		"hasBytes": len(mb) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce MarshalJSON no error -- nil", actual)
}

func Test_ErrorOnce_MarshalJSON_WithError_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })
	mb, _ := eo.MarshalJSON()
	eo2 := coreonce.NewErrorOnce(func() error { return nil })
	err := eo2.UnmarshalJSON(mb)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ErrorOnce UnmarshalJSON -- roundtrip", actual)
}

func Test_ErrorOnce_Serialize_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })
	bytes, err := eo.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce Serialize -- nil error", actual)
}

func Test_ErrorOnce_ConcatNewString_NoError_FromIntegerOnceCompariso(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })
	result := eo.ConcatNewString("extra")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorOnce ConcatNewString no error -- extra", actual)
}
