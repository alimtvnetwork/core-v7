package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// =============================================
// S22: LeftRightFromSplit
// =============================================

func Test_S22_001_LeftRightFromSplit_basic(t *testing.T) {
	safeTest(t, "Test_S22_001_LeftRightFromSplit_basic", func() {
		// Arrange & Act
		lr := corestr.LeftRightFromSplit("key=value", "=")

		// Assert
		if lr.Left != "key" {
			t.Errorf("LeftRightFromSplit Left returns key, got %s", lr.Left)
		}
		if lr.Right != "value" {
			t.Errorf("LeftRightFromSplit Right returns value, got %s", lr.Right)
		}
	})
}

func Test_S22_002_LeftRightFromSplit_no_sep(t *testing.T) {
	safeTest(t, "Test_S22_002_LeftRightFromSplit_no_sep", func() {
		// Arrange & Act
		lr := corestr.LeftRightFromSplit("nosep", "=")

		// Assert
		if lr.Left != "nosep" {
			t.Errorf("LeftRightFromSplit Left returns full string -- no sep, got %s", lr.Left)
		}
	})
}

func Test_S22_003_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_S22_003_LeftRightFromSplitTrimmed", func() {
		// Arrange & Act
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")

		// Assert
		if lr.Left != "key" {
			t.Errorf("LeftRightFromSplitTrimmed Left returns trimmed key, got %q", lr.Left)
		}
		if lr.Right != "value" {
			t.Errorf("LeftRightFromSplitTrimmed Right returns trimmed value, got %q", lr.Right)
		}
	})
}

func Test_S22_004_LeftRightFromSplitFull_multi_sep(t *testing.T) {
	safeTest(t, "Test_S22_004_LeftRightFromSplitFull_multi_sep", func() {
		// Arrange & Act
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")

		// Assert
		if lr.Left != "a" {
			t.Errorf("LeftRightFromSplitFull Left returns a, got %s", lr.Left)
		}
		if lr.Right != "b:c:d" {
			t.Errorf("LeftRightFromSplitFull Right returns b:c:d, got %s", lr.Right)
		}
	})
}

func Test_S22_005_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_S22_005_LeftRightFromSplitFullTrimmed", func() {
		// Arrange & Act
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")

		// Assert
		if lr.Left != "a" {
			t.Errorf("LeftRightFromSplitFullTrimmed Left returns trimmed a, got %q", lr.Left)
		}
		if lr.Right != "b : c" {
			t.Errorf("LeftRightFromSplitFullTrimmed Right returns trimmed remainder, got %q", lr.Right)
		}
	})
}

func Test_S22_006_LeftRightFromSplit_empty(t *testing.T) {
	safeTest(t, "Test_S22_006_LeftRightFromSplit_empty", func() {
		// Arrange & Act
		lr := corestr.LeftRightFromSplit("", "=")

		// Assert
		if lr.Left != "" {
			t.Errorf("LeftRightFromSplit Left returns empty -- empty input, got %q", lr.Left)
		}
	})
}

// =============================================
// S22: LeftMiddleRightFromSplit
// =============================================

func Test_S22_010_LeftMiddleRightFromSplit_basic(t *testing.T) {
	safeTest(t, "Test_S22_010_LeftMiddleRightFromSplit_basic", func() {
		// Arrange & Act
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")

		// Assert
		if lmr.Left != "a" {
			t.Errorf("LeftMiddleRightFromSplit Left returns a, got %s", lmr.Left)
		}
		if lmr.Middle != "b" {
			t.Errorf("LeftMiddleRightFromSplit Middle returns b, got %s", lmr.Middle)
		}
		if lmr.Right != "c" {
			t.Errorf("LeftMiddleRightFromSplit Right returns c, got %s", lmr.Right)
		}
	})
}

func Test_S22_011_LeftMiddleRightFromSplit_no_sep(t *testing.T) {
	safeTest(t, "Test_S22_011_LeftMiddleRightFromSplit_no_sep", func() {
		// Arrange & Act
		lmr := corestr.LeftMiddleRightFromSplit("nosep", ".")

		// Assert
		if lmr.Left != "nosep" {
			t.Errorf("LeftMiddleRightFromSplit Left returns full string -- no sep, got %s", lmr.Left)
		}
	})
}

func Test_S22_012_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_S22_012_LeftMiddleRightFromSplitTrimmed", func() {
		// Arrange & Act
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")

		// Assert
		if lmr.Left != "a" {
			t.Errorf("LeftMiddleRightFromSplitTrimmed Left returns trimmed a, got %q", lmr.Left)
		}
		if lmr.Middle != "b" {
			t.Errorf("LeftMiddleRightFromSplitTrimmed Middle returns trimmed b, got %q", lmr.Middle)
		}
		if lmr.Right != "c" {
			t.Errorf("LeftMiddleRightFromSplitTrimmed Right returns trimmed c, got %q", lmr.Right)
		}
	})
}

func Test_S22_013_LeftMiddleRightFromSplitN(t *testing.T) {
	safeTest(t, "Test_S22_013_LeftMiddleRightFromSplitN", func() {
		// Arrange & Act
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")

		// Assert
		if lmr.Left != "a" {
			t.Errorf("LeftMiddleRightFromSplitN Left returns a, got %s", lmr.Left)
		}
		if lmr.Middle != "b" {
			t.Errorf("LeftMiddleRightFromSplitN Middle returns b, got %s", lmr.Middle)
		}
		if lmr.Right != "c:d:e" {
			t.Errorf("LeftMiddleRightFromSplitN Right returns c:d:e, got %s", lmr.Right)
		}
	})
}

func Test_S22_014_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_S22_014_LeftMiddleRightFromSplitNTrimmed", func() {
		// Arrange & Act
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")

		// Assert
		if lmr.Left != "a" {
			t.Errorf("LeftMiddleRightFromSplitNTrimmed Left returns trimmed a, got %q", lmr.Left)
		}
		if lmr.Middle != "b" {
			t.Errorf("LeftMiddleRightFromSplitNTrimmed Middle returns trimmed b, got %q", lmr.Middle)
		}
	})
}

func Test_S22_015_LeftMiddleRightFromSplit_empty(t *testing.T) {
	safeTest(t, "Test_S22_015_LeftMiddleRightFromSplit_empty", func() {
		// Arrange & Act
		lmr := corestr.LeftMiddleRightFromSplit("", ".")

		// Assert
		if lmr.Left != "" {
			t.Errorf("LeftMiddleRightFromSplit Left returns empty -- empty input, got %q", lmr.Left)
		}
	})
}

// =============================================
// S22: newCollectionCreator
// =============================================

func Test_S22_020_NewCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S22_020_NewCollection_Empty", func() {
		// Act
		c := corestr.New.Collection.Empty()

		// Assert
		if c == nil || !c.IsEmpty() {
			t.Error("New.Collection.Empty returns empty collection")
		}
	})
}

func Test_S22_021_NewCollection_Cap(t *testing.T) {
	safeTest(t, "Test_S22_021_NewCollection_Cap", func() {
		// Act
		c := corestr.New.Collection.Cap(10)

		// Assert
		if c == nil || !c.IsEmpty() {
			t.Error("New.Collection.Cap returns empty collection with capacity")
		}
	})
}

func Test_S22_022_NewCollection_CloneStrings(t *testing.T) {
	safeTest(t, "Test_S22_022_NewCollection_CloneStrings", func() {
		// Arrange
		items := []string{"a", "b"}

		// Act
		c := corestr.New.Collection.CloneStrings(items)

		// Assert
		if c.Length() != 2 {
			t.Errorf("New.Collection.CloneStrings returns 2 items, got %d", c.Length())
		}
	})
}

func Test_S22_023_NewCollection_Create(t *testing.T) {
	safeTest(t, "Test_S22_023_NewCollection_Create", func() {
		// Act
		c := corestr.New.Collection.Create([]string{"x"})

		// Assert
		if c.Length() != 1 {
			t.Errorf("New.Collection.Create returns 1 item, got %d", c.Length())
		}
	})
}

func Test_S22_024_NewCollection_StringsOptions_clone(t *testing.T) {
	safeTest(t, "Test_S22_024_NewCollection_StringsOptions_clone", func() {
		// Act
		c := corestr.New.Collection.StringsOptions(true, []string{"a", "b"})

		// Assert
		if c.Length() != 2 {
			t.Errorf("New.Collection.StringsOptions clone returns 2 items, got %d", c.Length())
		}
	})
}

func Test_S22_025_NewCollection_StringsOptions_no_clone_empty(t *testing.T) {
	safeTest(t, "Test_S22_025_NewCollection_StringsOptions_no_clone_empty", func() {
		// Act
		c := corestr.New.Collection.StringsOptions(false, []string{})

		// Assert
		if !c.IsEmpty() {
			t.Error("New.Collection.StringsOptions no clone empty returns empty")
		}
	})
}

func Test_S22_026_NewCollection_LineUsingSep(t *testing.T) {
	safeTest(t, "Test_S22_026_NewCollection_LineUsingSep", func() {
		// Act
		c := corestr.New.Collection.LineUsingSep(",", "a,b,c")

		// Assert
		if c.Length() != 3 {
			t.Errorf("New.Collection.LineUsingSep returns 3 items, got %d", c.Length())
		}
	})
}

func Test_S22_027_NewCollection_LineDefault(t *testing.T) {
	safeTest(t, "Test_S22_027_NewCollection_LineDefault", func() {
		// Act
		c := corestr.New.Collection.LineDefault("a\nb\nc")

		// Assert
		if c.Length() != 3 {
			t.Errorf("New.Collection.LineDefault returns 3 items, got %d", c.Length())
		}
	})
}

func Test_S22_028_NewCollection_StringsPlusCap_zero_cap(t *testing.T) {
	safeTest(t, "Test_S22_028_NewCollection_StringsPlusCap_zero_cap", func() {
		// Act
		c := corestr.New.Collection.StringsPlusCap(0, []string{"a"})

		// Assert
		if c.Length() != 1 {
			t.Errorf("New.Collection.StringsPlusCap returns 1 item, got %d", c.Length())
		}
	})
}

func Test_S22_029_NewCollection_StringsPlusCap_with_cap(t *testing.T) {
	safeTest(t, "Test_S22_029_NewCollection_StringsPlusCap_with_cap", func() {
		// Act
		c := corestr.New.Collection.StringsPlusCap(5, []string{"a"})

		// Assert
		if c.Length() != 1 {
			t.Errorf("New.Collection.StringsPlusCap returns 1 item, got %d", c.Length())
		}
	})
}

func Test_S22_030_NewCollection_CapStrings_zero_cap(t *testing.T) {
	safeTest(t, "Test_S22_030_NewCollection_CapStrings_zero_cap", func() {
		// Act
		c := corestr.New.Collection.CapStrings(0, []string{"a"})

		// Assert
		if c.Length() != 1 {
			t.Errorf("New.Collection.CapStrings returns 1 item, got %d", c.Length())
		}
	})
}

func Test_S22_031_NewCollection_CapStrings_with_cap(t *testing.T) {
	safeTest(t, "Test_S22_031_NewCollection_CapStrings_with_cap", func() {
		// Act
		c := corestr.New.Collection.CapStrings(5, []string{"a"})

		// Assert
		if c.Length() != 1 {
			t.Errorf("New.Collection.CapStrings returns 1 item, got %d", c.Length())
		}
	})
}

func Test_S22_032_NewCollection_LenCap(t *testing.T) {
	safeTest(t, "Test_S22_032_NewCollection_LenCap", func() {
		// Act
		c := corestr.New.Collection.LenCap(3, 10)

		// Assert
		if c.Length() != 3 {
			t.Errorf("New.Collection.LenCap returns 3 items, got %d", c.Length())
		}
	})
}

// =============================================
// S22: newHashsetCreator
// =============================================

func Test_S22_040_NewHashset_Empty(t *testing.T) {
	safeTest(t, "Test_S22_040_NewHashset_Empty", func() {
		// Act
		hs := corestr.New.Hashset.Empty()

		// Assert
		if hs == nil || hs.Length() != 0 {
			t.Error("New.Hashset.Empty returns empty hashset")
		}
	})
}

func Test_S22_041_NewHashset_Cap(t *testing.T) {
	safeTest(t, "Test_S22_041_NewHashset_Cap", func() {
		// Act
		hs := corestr.New.Hashset.Cap(10)

		// Assert
		if hs == nil {
			t.Error("New.Hashset.Cap returns non-nil")
		}
	})
}

func Test_S22_042_NewHashset_Strings(t *testing.T) {
	safeTest(t, "Test_S22_042_NewHashset_Strings", func() {
		// Act
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "a"})

		// Assert
		if hs.Length() != 2 {
			t.Errorf("New.Hashset.Strings returns 2 unique items, got %d", hs.Length())
		}
	})
}

func Test_S22_043_NewHashset_Strings_empty(t *testing.T) {
	safeTest(t, "Test_S22_043_NewHashset_Strings_empty", func() {
		// Act
		hs := corestr.New.Hashset.Strings([]string{})

		// Assert
		if hs.Length() != 0 {
			t.Errorf("New.Hashset.Strings returns 0 -- empty input, got %d", hs.Length())
		}
	})
}

func Test_S22_044_NewHashset_StringsSpreadItems(t *testing.T) {
	safeTest(t, "Test_S22_044_NewHashset_StringsSpreadItems", func() {
		// Act
		hs := corestr.New.Hashset.StringsSpreadItems("x", "y")

		// Assert
		if hs.Length() != 2 {
			t.Errorf("New.Hashset.StringsSpreadItems returns 2, got %d", hs.Length())
		}
	})
}

func Test_S22_045_NewHashset_StringsSpreadItems_empty(t *testing.T) {
	safeTest(t, "Test_S22_045_NewHashset_StringsSpreadItems_empty", func() {
		// Act
		hs := corestr.New.Hashset.StringsSpreadItems()

		// Assert
		if hs.Length() != 0 {
			t.Error("New.Hashset.StringsSpreadItems returns 0 -- empty")
		}
	})
}

func Test_S22_046_NewHashset_StringsOption_nil_zero(t *testing.T) {
	safeTest(t, "Test_S22_046_NewHashset_StringsOption_nil_zero", func() {
		// Act
		hs := corestr.New.Hashset.StringsOption(0, false)

		// Assert
		if hs == nil {
			t.Error("New.Hashset.StringsOption returns non-nil -- nil items zero cap")
		}
	})
}

func Test_S22_047_NewHashset_StringsOption_nil_with_cap(t *testing.T) {
	safeTest(t, "Test_S22_047_NewHashset_StringsOption_nil_with_cap", func() {
		// Act
		hs := corestr.New.Hashset.StringsOption(5, false)

		// Assert
		if hs == nil {
			t.Error("New.Hashset.StringsOption returns non-nil -- nil items with cap")
		}
	})
}

func Test_S22_048_NewHashset_UsingCollection(t *testing.T) {
	safeTest(t, "Test_S22_048_NewHashset_UsingCollection", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		hs := corestr.New.Hashset.UsingCollection(col)

		// Assert
		if hs.Length() != 2 {
			t.Errorf("New.Hashset.UsingCollection returns 2, got %d", hs.Length())
		}
	})
}

func Test_S22_049_NewHashset_UsingCollection_nil(t *testing.T) {
	safeTest(t, "Test_S22_049_NewHashset_UsingCollection_nil", func() {
		// Act
		hs := corestr.New.Hashset.UsingCollection(nil)

		// Assert
		if hs.Length() != 0 {
			t.Error("New.Hashset.UsingCollection returns empty -- nil input")
		}
	})
}

func Test_S22_050_NewHashset_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_S22_050_NewHashset_SimpleSlice", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		hs := corestr.New.Hashset.SimpleSlice(ss)

		// Assert
		if hs.Length() != 2 {
			t.Errorf("New.Hashset.SimpleSlice returns 2, got %d", hs.Length())
		}
	})
}

func Test_S22_051_NewHashset_UsingMap(t *testing.T) {
	safeTest(t, "Test_S22_051_NewHashset_UsingMap", func() {
		// Act
		hs := corestr.New.Hashset.UsingMap(map[string]bool{"a": true, "b": true})

		// Assert
		if hs.Length() != 2 {
			t.Errorf("New.Hashset.UsingMap returns 2, got %d", hs.Length())
		}
	})
}

func Test_S22_052_NewHashset_UsingMap_empty(t *testing.T) {
	safeTest(t, "Test_S22_052_NewHashset_UsingMap_empty", func() {
		// Act
		hs := corestr.New.Hashset.UsingMap(map[string]bool{})

		// Assert
		if hs.Length() != 0 {
			t.Error("New.Hashset.UsingMap returns empty -- empty map")
		}
	})
}

func Test_S22_053_NewHashset_UsingMapOption_clone(t *testing.T) {
	safeTest(t, "Test_S22_053_NewHashset_UsingMapOption_clone", func() {
		// Act
		hs := corestr.New.Hashset.UsingMapOption(0, true, map[string]bool{"a": true})

		// Assert
		if hs.Length() != 1 {
			t.Errorf("New.Hashset.UsingMapOption clone returns 1, got %d", hs.Length())
		}
	})
}

func Test_S22_054_NewHashset_UsingMapOption_no_clone(t *testing.T) {
	safeTest(t, "Test_S22_054_NewHashset_UsingMapOption_no_clone", func() {
		// Act
		hs := corestr.New.Hashset.UsingMapOption(0, false, map[string]bool{"a": true})

		// Assert
		if hs.Length() != 1 {
			t.Errorf("New.Hashset.UsingMapOption no clone returns 1, got %d", hs.Length())
		}
	})
}

func Test_S22_055_NewHashset_UsingMapOption_empty(t *testing.T) {
	safeTest(t, "Test_S22_055_NewHashset_UsingMapOption_empty", func() {
		// Act
		hs := corestr.New.Hashset.UsingMapOption(5, false, map[string]bool{})

		// Assert
		if hs == nil {
			t.Error("New.Hashset.UsingMapOption returns non-nil -- empty map with cap")
		}
	})
}

// =============================================
// S22: newHashmapCreator
// =============================================

func Test_S22_060_NewHashmap_Empty(t *testing.T) {
	safeTest(t, "Test_S22_060_NewHashmap_Empty", func() {
		// Act
		hm := corestr.New.Hashmap.Empty()

		// Assert
		if hm == nil || hm.Length() != 0 {
			t.Error("New.Hashmap.Empty returns empty hashmap")
		}
	})
}

func Test_S22_061_NewHashmap_Cap(t *testing.T) {
	safeTest(t, "Test_S22_061_NewHashmap_Cap", func() {
		// Act
		hm := corestr.New.Hashmap.Cap(10)

		// Assert
		if hm == nil {
			t.Error("New.Hashmap.Cap returns non-nil")
		}
	})
}

func Test_S22_062_NewHashmap_UsingMap(t *testing.T) {
	safeTest(t, "Test_S22_062_NewHashmap_UsingMap", func() {
		// Act
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Assert
		if hm.Length() != 1 {
			t.Errorf("New.Hashmap.UsingMap returns 1, got %d", hm.Length())
		}
	})
}

func Test_S22_063_NewHashmap_UsingMapOptions_clone(t *testing.T) {
	safeTest(t, "Test_S22_063_NewHashmap_UsingMapOptions_clone", func() {
		// Act
		hm := corestr.New.Hashmap.UsingMapOptions(true, 0, map[string]string{"a": "1"})

		// Assert
		if hm.Length() != 1 {
			t.Errorf("New.Hashmap.UsingMapOptions clone returns 1, got %d", hm.Length())
		}
	})
}

func Test_S22_064_NewHashmap_UsingMapOptions_no_clone(t *testing.T) {
	safeTest(t, "Test_S22_064_NewHashmap_UsingMapOptions_no_clone", func() {
		// Act
		hm := corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{"a": "1"})

		// Assert
		if hm.Length() != 1 {
			t.Errorf("New.Hashmap.UsingMapOptions no clone returns 1, got %d", hm.Length())
		}
	})
}

func Test_S22_065_NewHashmap_UsingMapOptions_empty(t *testing.T) {
	safeTest(t, "Test_S22_065_NewHashmap_UsingMapOptions_empty", func() {
		// Act
		hm := corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{})

		// Assert
		if hm == nil {
			t.Error("New.Hashmap.UsingMapOptions returns non-nil -- empty with cap")
		}
	})
}

func Test_S22_066_NewHashmap_MapWithCap_zero_cap(t *testing.T) {
	safeTest(t, "Test_S22_066_NewHashmap_MapWithCap_zero_cap", func() {
		// Act
		hm := corestr.New.Hashmap.MapWithCap(0, map[string]string{"a": "1"})

		// Assert
		if hm.Length() != 1 {
			t.Errorf("New.Hashmap.MapWithCap returns 1, got %d", hm.Length())
		}
	})
}

func Test_S22_067_NewHashmap_MapWithCap_with_cap(t *testing.T) {
	safeTest(t, "Test_S22_067_NewHashmap_MapWithCap_with_cap", func() {
		// Act
		hm := corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})

		// Assert
		if hm.Length() != 1 {
			t.Errorf("New.Hashmap.MapWithCap returns 1, got %d", hm.Length())
		}
	})
}

func Test_S22_068_NewHashmap_MapWithCap_empty(t *testing.T) {
	safeTest(t, "Test_S22_068_NewHashmap_MapWithCap_empty", func() {
		// Act
		hm := corestr.New.Hashmap.MapWithCap(5, map[string]string{})

		// Assert
		if hm == nil {
			t.Error("New.Hashmap.MapWithCap returns non-nil -- empty with cap")
		}
	})
}

func Test_S22_069_NewHashmap_KeyValues(t *testing.T) {
	safeTest(t, "Test_S22_069_NewHashmap_KeyValues", func() {
		// Act
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Assert
		if hm.Length() != 1 {
			t.Errorf("New.Hashmap.KeyValues returns 1, got %d", hm.Length())
		}
	})
}

func Test_S22_070_NewHashmap_KeyValues_empty(t *testing.T) {
	safeTest(t, "Test_S22_070_NewHashmap_KeyValues_empty", func() {
		// Act
		hm := corestr.New.Hashmap.KeyValues()

		// Assert
		if hm == nil {
			t.Error("New.Hashmap.KeyValues returns non-nil -- empty")
		}
	})
}

func Test_S22_071_NewHashmap_KeyAnyValues(t *testing.T) {
	safeTest(t, "Test_S22_071_NewHashmap_KeyAnyValues", func() {
		// Act
		hm := corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "a", Value: "1"})

		// Assert
		if hm.Length() != 1 {
			t.Errorf("New.Hashmap.KeyAnyValues returns 1, got %d", hm.Length())
		}
	})
}

func Test_S22_072_NewHashmap_KeyAnyValues_empty(t *testing.T) {
	safeTest(t, "Test_S22_072_NewHashmap_KeyAnyValues_empty", func() {
		// Act
		hm := corestr.New.Hashmap.KeyAnyValues()

		// Assert
		if hm == nil {
			t.Error("New.Hashmap.KeyAnyValues returns non-nil -- empty")
		}
	})
}

func Test_S22_073_NewHashmap_KeyValuesCollection(t *testing.T) {
	safeTest(t, "Test_S22_073_NewHashmap_KeyValuesCollection", func() {
		// Arrange
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})

		// Act
		hm := corestr.New.Hashmap.KeyValuesCollection(keys, vals)

		// Assert
		if hm.Length() != 2 {
			t.Errorf("New.Hashmap.KeyValuesCollection returns 2, got %d", hm.Length())
		}
	})
}

func Test_S22_074_NewHashmap_KeyValuesCollection_nil_keys(t *testing.T) {
	safeTest(t, "Test_S22_074_NewHashmap_KeyValuesCollection_nil_keys", func() {
		// Act
		hm := corestr.New.Hashmap.KeyValuesCollection(nil, nil)

		// Assert
		if hm.Length() != 0 {
			t.Error("New.Hashmap.KeyValuesCollection returns empty -- nil keys")
		}
	})
}

func Test_S22_075_NewHashmap_KeyValuesStrings(t *testing.T) {
	safeTest(t, "Test_S22_075_NewHashmap_KeyValuesStrings", func() {
		// Act
		hm := corestr.New.Hashmap.KeyValuesStrings([]string{"a"}, []string{"1"})

		// Assert
		if hm.Length() != 1 {
			t.Errorf("New.Hashmap.KeyValuesStrings returns 1, got %d", hm.Length())
		}
	})
}

func Test_S22_076_NewHashmap_KeyValuesStrings_empty(t *testing.T) {
	safeTest(t, "Test_S22_076_NewHashmap_KeyValuesStrings_empty", func() {
		// Act
		hm := corestr.New.Hashmap.KeyValuesStrings([]string{}, []string{})

		// Assert
		if hm.Length() != 0 {
			t.Error("New.Hashmap.KeyValuesStrings returns empty -- empty keys")
		}
	})
}

// =============================================
// S22: newLinkedListCreator
// =============================================

func Test_S22_080_NewLinkedList_Create(t *testing.T) {
	safeTest(t, "Test_S22_080_NewLinkedList_Create", func() {
		// Act
		ll := corestr.New.LinkedList.Create()

		// Assert
		if ll == nil || ll.Length() != 0 {
			t.Error("New.LinkedList.Create returns empty list")
		}
	})
}

func Test_S22_081_NewLinkedList_Strings(t *testing.T) {
	safeTest(t, "Test_S22_081_NewLinkedList_Strings", func() {
		// Act
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Assert
		if ll.Length() != 2 {
			t.Errorf("New.LinkedList.Strings returns 2, got %d", ll.Length())
		}
	})
}

func Test_S22_082_NewLinkedList_Strings_empty(t *testing.T) {
	safeTest(t, "Test_S22_082_NewLinkedList_Strings_empty", func() {
		// Act
		ll := corestr.New.LinkedList.Strings([]string{})

		// Assert
		if ll.Length() != 0 {
			t.Error("New.LinkedList.Strings returns empty -- empty input")
		}
	})
}

func Test_S22_083_NewLinkedList_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_S22_083_NewLinkedList_SpreadStrings", func() {
		// Act
		ll := corestr.New.LinkedList.SpreadStrings("x", "y", "z")

		// Assert
		if ll.Length() != 3 {
			t.Errorf("New.LinkedList.SpreadStrings returns 3, got %d", ll.Length())
		}
	})
}

func Test_S22_084_NewLinkedList_SpreadStrings_empty(t *testing.T) {
	safeTest(t, "Test_S22_084_NewLinkedList_SpreadStrings_empty", func() {
		// Act
		ll := corestr.New.LinkedList.SpreadStrings()

		// Assert
		if ll.Length() != 0 {
			t.Error("New.LinkedList.SpreadStrings returns empty -- no args")
		}
	})
}

func Test_S22_085_NewLinkedList_UsingMap(t *testing.T) {
	safeTest(t, "Test_S22_085_NewLinkedList_UsingMap", func() {
		// Act
		ll := corestr.New.LinkedList.UsingMap(map[string]bool{"a": true, "b": true})

		// Assert
		if ll.Length() != 2 {
			t.Errorf("New.LinkedList.UsingMap returns 2, got %d", ll.Length())
		}
	})
}

func Test_S22_086_NewLinkedList_UsingMap_nil(t *testing.T) {
	safeTest(t, "Test_S22_086_NewLinkedList_UsingMap_nil", func() {
		// Act
		ll := corestr.New.LinkedList.UsingMap(nil)

		// Assert
		if ll.Length() != 0 {
			t.Error("New.LinkedList.UsingMap returns empty -- nil input")
		}
	})
}

// =============================================
// S22: newLinkedListCollectionsCreator
// =============================================

func Test_S22_090_NewLinkedCollection_Create(t *testing.T) {
	safeTest(t, "Test_S22_090_NewLinkedCollection_Create", func() {
		// Act
		lc := corestr.New.LinkedCollection.Create()

		// Assert
		if lc == nil || lc.Length() != 0 {
			t.Error("New.LinkedCollection.Create returns empty")
		}
	})
}

func Test_S22_091_NewLinkedCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S22_091_NewLinkedCollection_Empty", func() {
		// Act
		lc := corestr.New.LinkedCollection.Empty()

		// Assert
		if lc == nil || lc.Length() != 0 {
			t.Error("New.LinkedCollection.Empty returns empty")
		}
	})
}

func Test_S22_092_NewLinkedCollection_UsingCollections(t *testing.T) {
	safeTest(t, "Test_S22_092_NewLinkedCollection_UsingCollections", func() {
		// Arrange
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"b"})

		// Act
		lc := corestr.New.LinkedCollection.UsingCollections(col1, col2)

		// Assert
		if lc.Length() != 2 {
			t.Errorf("New.LinkedCollection.UsingCollections returns 2, got %d", lc.Length())
		}
	})
}

func Test_S22_093_NewLinkedCollection_UsingCollections_nil(t *testing.T) {
	safeTest(t, "Test_S22_093_NewLinkedCollection_UsingCollections_nil", func() {
		// Act
		lc := corestr.New.LinkedCollection.UsingCollections()

		// Assert
		if lc.Length() != 0 {
			t.Error("New.LinkedCollection.UsingCollections returns empty -- nil")
		}
	})
}

func Test_S22_094_NewLinkedCollection_Strings(t *testing.T) {
	safeTest(t, "Test_S22_094_NewLinkedCollection_Strings", func() {
		// Act
		lc := corestr.New.LinkedCollection.Strings("a", "b")

		// Assert — Strings bundles all items into a single collection node
		if lc.Length() != 1 {
			t.Errorf("New.LinkedCollection.Strings returns 1 node, got %d", lc.Length())
		}
	})
}

func Test_S22_095_NewLinkedCollection_Strings_empty(t *testing.T) {
	safeTest(t, "Test_S22_095_NewLinkedCollection_Strings_empty", func() {
		// Act
		lc := corestr.New.LinkedCollection.Strings()

		// Assert
		if lc.Length() != 0 {
			t.Error("New.LinkedCollection.Strings returns empty -- no args")
		}
	})
}

// =============================================
// S22: newKeyValuesCreator
// =============================================

func Test_S22_100_NewKeyValues_Empty(t *testing.T) {
	safeTest(t, "Test_S22_100_NewKeyValues_Empty", func() {
		// Act
		kvc := corestr.New.KeyValues.Empty()

		// Assert
		if kvc == nil {
			t.Error("New.KeyValues.Empty returns non-nil")
		}
	})
}

func Test_S22_101_NewKeyValues_Cap(t *testing.T) {
	safeTest(t, "Test_S22_101_NewKeyValues_Cap", func() {
		// Act
		kvc := corestr.New.KeyValues.Cap(10)

		// Assert
		if kvc == nil {
			t.Error("New.KeyValues.Cap returns non-nil")
		}
	})
}

func Test_S22_102_NewKeyValues_UsingMap(t *testing.T) {
	safeTest(t, "Test_S22_102_NewKeyValues_UsingMap", func() {
		// Act
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{"a": "1"})

		// Assert
		if kvc.Length() != 1 {
			t.Errorf("New.KeyValues.UsingMap returns 1, got %d", kvc.Length())
		}
	})
}

func Test_S22_103_NewKeyValues_UsingMap_empty(t *testing.T) {
	safeTest(t, "Test_S22_103_NewKeyValues_UsingMap_empty", func() {
		// Act
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{})

		// Assert
		if kvc.Length() != 0 {
			t.Error("New.KeyValues.UsingMap returns empty -- empty map")
		}
	})
}

func Test_S22_104_NewKeyValues_UsingKeyValuePairs(t *testing.T) {
	safeTest(t, "Test_S22_104_NewKeyValues_UsingKeyValuePairs", func() {
		// Act
		kvc := corestr.New.KeyValues.UsingKeyValuePairs(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Assert
		if kvc.Length() != 1 {
			t.Errorf("New.KeyValues.UsingKeyValuePairs returns 1, got %d", kvc.Length())
		}
	})
}

func Test_S22_105_NewKeyValues_UsingKeyValuePairs_empty(t *testing.T) {
	safeTest(t, "Test_S22_105_NewKeyValues_UsingKeyValuePairs_empty", func() {
		// Act
		kvc := corestr.New.KeyValues.UsingKeyValuePairs()

		// Assert
		if kvc.Length() != 0 {
			t.Error("New.KeyValues.UsingKeyValuePairs returns empty -- no args")
		}
	})
}

func Test_S22_106_NewKeyValues_UsingKeyValueStrings(t *testing.T) {
	safeTest(t, "Test_S22_106_NewKeyValues_UsingKeyValueStrings", func() {
		// Act
		kvc := corestr.New.KeyValues.UsingKeyValueStrings(
			[]string{"a", "b"},
			[]string{"1", "2"},
		)

		// Assert
		if kvc.Length() != 2 {
			t.Errorf("New.KeyValues.UsingKeyValueStrings returns 2, got %d", kvc.Length())
		}
	})
}

func Test_S22_107_NewKeyValues_UsingKeyValueStrings_empty(t *testing.T) {
	safeTest(t, "Test_S22_107_NewKeyValues_UsingKeyValueStrings_empty", func() {
		// Act
		kvc := corestr.New.KeyValues.UsingKeyValueStrings([]string{}, []string{})

		// Assert
		if kvc.Length() != 0 {
			t.Error("New.KeyValues.UsingKeyValueStrings returns empty -- empty input")
		}
	})
}

// =============================================
// S22: newHashsetsCollectionCreator
// =============================================

func Test_S22_110_NewHashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S22_110_NewHashsetsCollection_Empty", func() {
		// Act
		hsc := corestr.New.HashsetsCollection.Empty()

		// Assert
		if hsc == nil {
			t.Error("New.HashsetsCollection.Empty returns non-nil")
		}
	})
}

func Test_S22_111_NewHashsetsCollection_Cap(t *testing.T) {
	safeTest(t, "Test_S22_111_NewHashsetsCollection_Cap", func() {
		// Act
		hsc := corestr.New.HashsetsCollection.Cap(5)

		// Assert
		if hsc == nil {
			t.Error("New.HashsetsCollection.Cap returns non-nil")
		}
	})
}

func Test_S22_112_NewHashsetsCollection_LenCap(t *testing.T) {
	safeTest(t, "Test_S22_112_NewHashsetsCollection_LenCap", func() {
		// Act
		hsc := corestr.New.HashsetsCollection.LenCap(0, 5)

		// Assert
		if hsc == nil {
			t.Error("New.HashsetsCollection.LenCap returns non-nil")
		}
	})
}

func Test_S22_113_NewHashsetsCollection_UsingHashsets(t *testing.T) {
	safeTest(t, "Test_S22_113_NewHashsetsCollection_UsingHashsets", func() {
		// Arrange
		hs := *corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		hsc := corestr.New.HashsetsCollection.UsingHashsets(hs)

		// Assert
		if hsc == nil {
			t.Error("New.HashsetsCollection.UsingHashsets returns non-nil")
		}
	})
}

func Test_S22_114_NewHashsetsCollection_UsingHashsets_empty(t *testing.T) {
	safeTest(t, "Test_S22_114_NewHashsetsCollection_UsingHashsets_empty", func() {
		// Act
		hsc := corestr.New.HashsetsCollection.UsingHashsets()

		// Assert
		if hsc == nil {
			t.Error("New.HashsetsCollection.UsingHashsets returns non-nil -- empty")
		}
	})
}

func Test_S22_115_NewHashsetsCollection_UsingHashsetsPointers(t *testing.T) {
	safeTest(t, "Test_S22_115_NewHashsetsCollection_UsingHashsetsPointers", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		hsc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Assert
		if hsc == nil {
			t.Error("New.HashsetsCollection.UsingHashsetsPointers returns non-nil")
		}
	})
}

func Test_S22_116_NewHashsetsCollection_UsingHashsetsPointers_empty(t *testing.T) {
	safeTest(t, "Test_S22_116_NewHashsetsCollection_UsingHashsetsPointers_empty", func() {
		// Act
		hsc := corestr.New.HashsetsCollection.UsingHashsetsPointers()

		// Assert
		if hsc == nil {
			t.Error("New.HashsetsCollection.UsingHashsetsPointers returns non-nil -- empty")
		}
	})
}

// =============================================
// S22: newCollectionsOfCollectionCreator
// =============================================

func Test_S22_120_NewCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S22_120_NewCollectionsOfCollection_Empty", func() {
		// Act
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Assert
		if coc == nil {
			t.Error("New.CollectionsOfCollection.Empty returns non-nil")
		}
	})
}

func Test_S22_121_NewCollectionsOfCollection_Cap(t *testing.T) {
	safeTest(t, "Test_S22_121_NewCollectionsOfCollection_Cap", func() {
		// Act
		coc := corestr.New.CollectionsOfCollection.Cap(5)

		// Assert
		if coc == nil {
			t.Error("New.CollectionsOfCollection.Cap returns non-nil")
		}
	})
}

func Test_S22_122_NewCollectionsOfCollection_LenCap(t *testing.T) {
	safeTest(t, "Test_S22_122_NewCollectionsOfCollection_LenCap", func() {
		// Act
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 5)

		// Assert
		if coc == nil {
			t.Error("New.CollectionsOfCollection.LenCap returns non-nil")
		}
	})
}

func Test_S22_123_NewCollectionsOfCollection_Strings(t *testing.T) {
	safeTest(t, "Test_S22_123_NewCollectionsOfCollection_Strings", func() {
		// Act
		coc := corestr.New.CollectionsOfCollection.Strings([]string{"a", "b"})

		// Assert
		if coc == nil {
			t.Error("New.CollectionsOfCollection.Strings returns non-nil")
		}
	})
}

func Test_S22_124_NewCollectionsOfCollection_CloneStrings(t *testing.T) {
	safeTest(t, "Test_S22_124_NewCollectionsOfCollection_CloneStrings", func() {
		// Act
		coc := corestr.New.CollectionsOfCollection.CloneStrings([]string{"a"})

		// Assert
		if coc == nil {
			t.Error("New.CollectionsOfCollection.CloneStrings returns non-nil")
		}
	})
}

func Test_S22_125_NewCollectionsOfCollection_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_S22_125_NewCollectionsOfCollection_SpreadStrings", func() {
		// Act
		coc := corestr.New.CollectionsOfCollection.SpreadStrings(true, "a", "b")

		// Assert
		if coc == nil {
			t.Error("New.CollectionsOfCollection.SpreadStrings returns non-nil")
		}
	})
}

func Test_S22_126_NewCollectionsOfCollection_StringsOption(t *testing.T) {
	safeTest(t, "Test_S22_126_NewCollectionsOfCollection_StringsOption", func() {
		// Act
		coc := corestr.New.CollectionsOfCollection.StringsOption(true, 5, []string{"a"})

		// Assert
		if coc == nil {
			t.Error("New.CollectionsOfCollection.StringsOption returns non-nil")
		}
	})
}

func Test_S22_127_NewCollectionsOfCollection_StringsOptions(t *testing.T) {
	safeTest(t, "Test_S22_127_NewCollectionsOfCollection_StringsOptions", func() {
		// Act
		coc := corestr.New.CollectionsOfCollection.StringsOptions(false, 5, []string{"a"})

		// Assert
		if coc == nil {
			t.Error("New.CollectionsOfCollection.StringsOptions returns non-nil")
		}
	})
}

func Test_S22_128_NewCollectionsOfCollection_StringsOfStrings(t *testing.T) {
	safeTest(t, "Test_S22_128_NewCollectionsOfCollection_StringsOfStrings", func() {
		// Act
		coc := corestr.New.CollectionsOfCollection.StringsOfStrings(
			true,
			[]string{"a", "b"},
			[]string{"c", "d"},
		)

		// Assert
		if coc == nil {
			t.Error("New.CollectionsOfCollection.StringsOfStrings returns non-nil")
		}
	})
}

// =============================================
// S22: newSimpleStringOnceCreator
// =============================================

func Test_S22_130_NewSSO_Init(t *testing.T) {
	safeTest(t, "Test_S22_130_NewSSO_Init", func() {
		// Act
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Assert
		if sso.Value() != "hello" {
			t.Errorf("New.SimpleStringOnce.Init returns hello, got %s", sso.Value())
		}
		if !sso.IsInitialized() {
			t.Error("New.SimpleStringOnce.Init returns initialized")
		}
	})
}

func Test_S22_131_NewSSO_InitPtr(t *testing.T) {
	safeTest(t, "Test_S22_131_NewSSO_InitPtr", func() {
		// Act
		sso := corestr.New.SimpleStringOnce.InitPtr("world")

		// Assert
		if sso == nil {
			t.Fatal("New.SimpleStringOnce.InitPtr returns non-nil")
		}
		if sso.Value() != "world" {
			t.Errorf("New.SimpleStringOnce.InitPtr returns world, got %s", sso.Value())
		}
	})
}

func Test_S22_132_NewSSO_Uninitialized(t *testing.T) {
	safeTest(t, "Test_S22_132_NewSSO_Uninitialized", func() {
		// Act
		sso := corestr.New.SimpleStringOnce.Uninitialized("test")

		// Assert
		if sso.IsInitialized() {
			t.Error("New.SimpleStringOnce.Uninitialized returns not initialized")
		}
	})
}

func Test_S22_133_NewSSO_Create(t *testing.T) {
	safeTest(t, "Test_S22_133_NewSSO_Create", func() {
		// Act
		sso := corestr.New.SimpleStringOnce.Create("val", true)

		// Assert
		if sso.Value() != "val" {
			t.Errorf("New.SimpleStringOnce.Create returns val, got %s", sso.Value())
		}
	})
}

func Test_S22_134_NewSSO_CreatePtr(t *testing.T) {
	safeTest(t, "Test_S22_134_NewSSO_CreatePtr", func() {
		// Act
		sso := corestr.New.SimpleStringOnce.CreatePtr("val", false)

		// Assert
		if sso == nil {
			t.Fatal("New.SimpleStringOnce.CreatePtr returns non-nil")
		}
	})
}

func Test_S22_135_NewSSO_Empty(t *testing.T) {
	safeTest(t, "Test_S22_135_NewSSO_Empty", func() {
		// Act
		sso := corestr.New.SimpleStringOnce.Empty()

		// Assert
		if sso.Value() != "" {
			t.Error("New.SimpleStringOnce.Empty returns empty value")
		}
	})
}

func Test_S22_136_NewSSO_Any_no_field_names(t *testing.T) {
	safeTest(t, "Test_S22_136_NewSSO_Any_no_field_names", func() {
		// Act
		sso := corestr.New.SimpleStringOnce.Any(false, 42, true)

		// Assert
		if sso.Value() == "" {
			t.Error("New.SimpleStringOnce.Any returns non-empty -- int input")
		}
		if !sso.IsInitialized() {
			t.Error("New.SimpleStringOnce.Any returns initialized -- isInit true")
		}
	})
}

func Test_S22_137_NewSSO_Any_with_field_names(t *testing.T) {
	safeTest(t, "Test_S22_137_NewSSO_Any_with_field_names", func() {
		// Act
		sso := corestr.New.SimpleStringOnce.Any(true, "test", false)

		// Assert
		if sso.IsInitialized() {
			t.Error("New.SimpleStringOnce.Any returns not initialized -- isInit false")
		}
	})
}
