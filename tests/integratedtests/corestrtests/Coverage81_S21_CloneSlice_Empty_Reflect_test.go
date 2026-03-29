package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// =============================================
// S21: CloneSlice
// =============================================

func Test_S21_001_CloneSlice_with_items(t *testing.T) {
	safeTest(t, "Test_S21_001_CloneSlice_with_items", func() {
		// Arrange
		items := []string{"a", "b", "c"}

		// Act
		clone := corestr.CloneSlice(items)

		// Assert
		if len(clone) != 3 {
			t.Fatalf("CloneSlice returns 3 items, got %d", len(clone))
		}
		if clone[0] != "a" || clone[1] != "b" || clone[2] != "c" {
			t.Error("CloneSlice returns correct values -- a,b,c")
		}
		// Verify independence
		items[0] = "modified"
		if clone[0] == "modified" {
			t.Error("CloneSlice returns independent copy -- mutation should not propagate")
		}
	})
}

func Test_S21_002_CloneSlice_empty(t *testing.T) {
	safeTest(t, "Test_S21_002_CloneSlice_empty", func() {
		// Arrange
		items := []string{}

		// Act
		clone := corestr.CloneSlice(items)

		// Assert
		if clone == nil {
			t.Fatal("CloneSlice returns non-nil -- empty input")
		}
		if len(clone) != 0 {
			t.Errorf("CloneSlice returns empty slice, got %d", len(clone))
		}
	})
}

func Test_S21_003_CloneSlice_nil(t *testing.T) {
	safeTest(t, "Test_S21_003_CloneSlice_nil", func() {
		// Arrange
		var items []string

		// Act
		clone := corestr.CloneSlice(items)

		// Assert
		if clone == nil {
			t.Fatal("CloneSlice returns non-nil -- nil input")
		}
		if len(clone) != 0 {
			t.Errorf("CloneSlice returns empty slice -- nil input, got %d", len(clone))
		}
	})
}

// =============================================
// S21: CloneSliceIf
// =============================================

func Test_S21_010_CloneSliceIf_clone_true(t *testing.T) {
	safeTest(t, "Test_S21_010_CloneSliceIf_clone_true", func() {
		// Arrange
		items := []string{"x", "y"}

		// Act
		clone := corestr.CloneSliceIf(true, items...)

		// Assert
		if len(clone) != 2 {
			t.Fatalf("CloneSliceIf returns 2 items -- clone true, got %d", len(clone))
		}
		items[0] = "modified"
		if clone[0] == "modified" {
			t.Error("CloneSliceIf returns independent copy -- clone true")
		}
	})
}

func Test_S21_011_CloneSliceIf_clone_false(t *testing.T) {
	safeTest(t, "Test_S21_011_CloneSliceIf_clone_false", func() {
		// Arrange
		items := []string{"x", "y"}

		// Act
		result := corestr.CloneSliceIf(false, items...)

		// Assert
		if len(result) != 2 {
			t.Fatalf("CloneSliceIf returns 2 items -- clone false, got %d", len(result))
		}
		// When clone=false, should return same slice (shared backing array)
		items[0] = "modified"
		if result[0] != "modified" {
			t.Error("CloneSliceIf returns same reference -- clone false")
		}
	})
}

func Test_S21_012_CloneSliceIf_empty(t *testing.T) {
	safeTest(t, "Test_S21_012_CloneSliceIf_empty", func() {
		// Arrange & Act
		result := corestr.CloneSliceIf(true)

		// Assert
		if result == nil {
			t.Fatal("CloneSliceIf returns non-nil -- empty variadic")
		}
		if len(result) != 0 {
			t.Errorf("CloneSliceIf returns empty slice, got %d", len(result))
		}
	})
}

// =============================================
// S21: emptyCreator
// =============================================

func Test_S21_020_Empty_Collection(t *testing.T) {
	safeTest(t, "Test_S21_020_Empty_Collection", func() {
		// Act
		c := corestr.Empty.Collection()

		// Assert
		if c == nil || !c.IsEmpty() {
			t.Error("Empty.Collection returns empty collection")
		}
	})
}

func Test_S21_021_Empty_LinkedList(t *testing.T) {
	safeTest(t, "Test_S21_021_Empty_LinkedList", func() {
		// Act
		ll := corestr.Empty.LinkedList()

		// Assert
		if ll == nil || ll.Length() != 0 {
			t.Error("Empty.LinkedList returns empty linked list")
		}
	})
}

func Test_S21_022_Empty_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_S21_022_Empty_SimpleSlice", func() {
		// Act
		ss := corestr.Empty.SimpleSlice()

		// Assert
		if ss == nil || !ss.IsEmpty() {
			t.Error("Empty.SimpleSlice returns empty simple slice")
		}
	})
}

func Test_S21_023_Empty_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_S21_023_Empty_KeyAnyValuePair", func() {
		// Act
		kv := corestr.Empty.KeyAnyValuePair()

		// Assert
		if kv == nil {
			t.Error("Empty.KeyAnyValuePair returns non-nil")
		}
	})
}

func Test_S21_024_Empty_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_S21_024_Empty_KeyValuePair", func() {
		// Act
		kv := corestr.Empty.KeyValuePair()

		// Assert
		if kv == nil {
			t.Error("Empty.KeyValuePair returns non-nil")
		}
	})
}

func Test_S21_025_Empty_KeyValueCollection(t *testing.T) {
	safeTest(t, "Test_S21_025_Empty_KeyValueCollection", func() {
		// Act
		kvc := corestr.Empty.KeyValueCollection()

		// Assert
		if kvc == nil {
			t.Error("Empty.KeyValueCollection returns non-nil")
		}
	})
}

func Test_S21_026_Empty_LinkedCollections(t *testing.T) {
	safeTest(t, "Test_S21_026_Empty_LinkedCollections", func() {
		// Act
		lc := corestr.Empty.LinkedCollections()

		// Assert
		if lc == nil || lc.Length() != 0 {
			t.Error("Empty.LinkedCollections returns empty")
		}
	})
}

func Test_S21_027_Empty_LeftRight(t *testing.T) {
	safeTest(t, "Test_S21_027_Empty_LeftRight", func() {
		// Act
		lr := corestr.Empty.LeftRight()

		// Assert
		if lr == nil {
			t.Error("Empty.LeftRight returns non-nil")
		}
	})
}

func Test_S21_028_Empty_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_S21_028_Empty_SimpleStringOnce", func() {
		// Act
		sso := corestr.Empty.SimpleStringOnce()

		// Assert
		if sso.IsInitialized() {
			t.Error("Empty.SimpleStringOnce returns uninitialized")
		}
	})
}

func Test_S21_029_Empty_SimpleStringOncePtr(t *testing.T) {
	safeTest(t, "Test_S21_029_Empty_SimpleStringOncePtr", func() {
		// Act
		sso := corestr.Empty.SimpleStringOncePtr()

		// Assert
		if sso == nil {
			t.Error("Empty.SimpleStringOncePtr returns non-nil")
		}
	})
}

func Test_S21_030_Empty_Hashset(t *testing.T) {
	safeTest(t, "Test_S21_030_Empty_Hashset", func() {
		// Act
		hs := corestr.Empty.Hashset()

		// Assert
		if hs == nil || hs.Length() != 0 {
			t.Error("Empty.Hashset returns empty hashset")
		}
	})
}

func Test_S21_031_Empty_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_S21_031_Empty_HashsetsCollection", func() {
		// Act
		hsc := corestr.Empty.HashsetsCollection()

		// Assert
		if hsc == nil {
			t.Error("Empty.HashsetsCollection returns non-nil")
		}
	})
}

func Test_S21_032_Empty_Hashmap(t *testing.T) {
	safeTest(t, "Test_S21_032_Empty_Hashmap", func() {
		// Act
		hm := corestr.Empty.Hashmap()

		// Assert
		if hm == nil || hm.Length() != 0 {
			t.Error("Empty.Hashmap returns empty hashmap")
		}
	})
}

func Test_S21_033_Empty_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_S21_033_Empty_CharCollectionMap", func() {
		// Act
		ccm := corestr.Empty.CharCollectionMap()

		// Assert
		if ccm == nil {
			t.Error("Empty.CharCollectionMap returns non-nil")
		}
	})
}

func Test_S21_034_Empty_KeyValuesCollection(t *testing.T) {
	safeTest(t, "Test_S21_034_Empty_KeyValuesCollection", func() {
		// Act
		kvc := corestr.Empty.KeyValuesCollection()

		// Assert
		if kvc == nil {
			t.Error("Empty.KeyValuesCollection returns non-nil")
		}
	})
}

func Test_S21_035_Empty_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_S21_035_Empty_CollectionsOfCollection", func() {
		// Act
		coc := corestr.Empty.CollectionsOfCollection()

		// Assert
		if coc == nil {
			t.Error("Empty.CollectionsOfCollection returns non-nil")
		}
	})
}

func Test_S21_036_Empty_CharHashsetMap(t *testing.T) {
	safeTest(t, "Test_S21_036_Empty_CharHashsetMap", func() {
		// Act
		chm := corestr.Empty.CharHashsetMap()

		// Assert
		if chm == nil {
			t.Error("Empty.CharHashsetMap returns non-nil")
		}
	})
}

// =============================================
// S21: AnyToString
// =============================================

func Test_S21_040_AnyToString_empty_string(t *testing.T) {
	safeTest(t, "Test_S21_040_AnyToString_empty_string", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, "")

		// Assert
		if result != "" {
			t.Errorf("AnyToString returns empty -- empty input, got %q", result)
		}
	})
}

func Test_S21_041_AnyToString_with_value_no_field_name(t *testing.T) {
	safeTest(t, "Test_S21_041_AnyToString_with_value_no_field_name", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, "hello")

		// Assert
		if result == "" {
			t.Error("AnyToString returns non-empty -- has value")
		}
	})
}

func Test_S21_042_AnyToString_with_value_include_field_name(t *testing.T) {
	safeTest(t, "Test_S21_042_AnyToString_with_value_include_field_name", func() {
		// Arrange & Act
		result := corestr.AnyToString(true, "hello")

		// Assert
		if result == "" {
			t.Error("AnyToString returns non-empty -- include field name")
		}
	})
}

func Test_S21_043_AnyToString_with_int(t *testing.T) {
	safeTest(t, "Test_S21_043_AnyToString_with_int", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, 42)

		// Assert
		if result == "" {
			t.Error("AnyToString returns non-empty -- int input")
		}
	})
}

func Test_S21_044_AnyToString_with_pointer(t *testing.T) {
	safeTest(t, "Test_S21_044_AnyToString_with_pointer", func() {
		// Arrange
		val := "test"

		// Act
		result := corestr.AnyToString(false, &val)

		// Assert
		if result == "" {
			t.Error("AnyToString returns non-empty -- pointer input")
		}
	})
}

func Test_S21_045_AnyToString_with_nil(t *testing.T) {
	safeTest(t, "Test_S21_045_AnyToString_with_nil", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, nil)

		// Assert
		// nil is not "" so it goes through reflectInterfaceVal
		// Result depends on implementation
		_ = result
	})
}

func Test_S21_046_AnyToString_with_struct(t *testing.T) {
	safeTest(t, "Test_S21_046_AnyToString_with_struct", func() {
		// Arrange
		type sample struct {
			Name string
		}
		s := sample{Name: "test"}

		// Act
		result := corestr.AnyToString(true, s)

		// Assert
		if result == "" {
			t.Error("AnyToString returns non-empty -- struct input with field names")
		}
	})
}

func Test_S21_047_AnyToString_with_struct_pointer(t *testing.T) {
	safeTest(t, "Test_S21_047_AnyToString_with_struct_pointer", func() {
		// Arrange
		type sample struct {
			Name string
		}
		s := &sample{Name: "test"}

		// Act
		result := corestr.AnyToString(false, s)

		// Assert
		if result == "" {
			t.Error("AnyToString returns non-empty -- struct pointer input")
		}
	})
}
