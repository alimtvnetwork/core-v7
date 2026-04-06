package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": len(clone) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns 3 items", actual)
		actual := args.Map{"result": clone[0] != "a" || clone[1] != "b" || clone[2] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns correct values -- a,b,c", actual)
		// Verify independence
		items[0] = "modified"
		actual := args.Map{"result": clone[0] == "modified"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns independent copy -- mutation should not propagate", actual)
	})
}

func Test_S21_002_CloneSlice_empty(t *testing.T) {
	safeTest(t, "Test_S21_002_CloneSlice_empty", func() {
		// Arrange
		items := []string{}

		// Act
		clone := corestr.CloneSlice(items)

		// Assert
		actual := args.Map{"result": clone == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns non-nil -- empty input", actual)
		actual := args.Map{"result": len(clone) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty slice", actual)
	})
}

func Test_S21_003_CloneSlice_nil(t *testing.T) {
	safeTest(t, "Test_S21_003_CloneSlice_nil", func() {
		// Arrange
		var items []string

		// Act
		clone := corestr.CloneSlice(items)

		// Assert
		actual := args.Map{"result": clone == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns non-nil -- nil input", actual)
		actual := args.Map{"result": len(clone) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty slice -- nil input", actual)
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
		actual := args.Map{"result": len(clone) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns 2 items -- clone true", actual)
		items[0] = "modified"
		actual := args.Map{"result": clone[0] == "modified"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns independent copy -- clone true", actual)
	})
}

func Test_S21_011_CloneSliceIf_clone_false(t *testing.T) {
	safeTest(t, "Test_S21_011_CloneSliceIf_clone_false", func() {
		// Arrange
		items := []string{"x", "y"}

		// Act
		result := corestr.CloneSliceIf(false, items...)

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns 2 items -- clone false", actual)
		// When clone=false, should return same slice (shared backing array)
		items[0] = "modified"
		actual := args.Map{"result": result[0] != "modified"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns same reference -- clone false", actual)
	})
}

func Test_S21_012_CloneSliceIf_empty(t *testing.T) {
	safeTest(t, "Test_S21_012_CloneSliceIf_empty", func() {
		// Arrange & Act
		result := corestr.CloneSliceIf(true)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns non-nil -- empty variadic", actual)
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty slice", actual)
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
		actual := args.Map{"result": c == nil || !c.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.Collection returns empty collection", actual)
	})
}

func Test_S21_021_Empty_LinkedList(t *testing.T) {
	safeTest(t, "Test_S21_021_Empty_LinkedList", func() {
		// Act
		ll := corestr.Empty.LinkedList()

		// Assert
		actual := args.Map{"result": ll == nil || ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.LinkedList returns empty linked list", actual)
	})
}

func Test_S21_022_Empty_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_S21_022_Empty_SimpleSlice", func() {
		// Act
		ss := corestr.Empty.SimpleSlice()

		// Assert
		actual := args.Map{"result": ss == nil || !ss.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.SimpleSlice returns empty simple slice", actual)
	})
}

func Test_S21_023_Empty_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_S21_023_Empty_KeyAnyValuePair", func() {
		// Act
		kv := corestr.Empty.KeyAnyValuePair()

		// Assert
		actual := args.Map{"result": kv == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.KeyAnyValuePair returns non-nil", actual)
	})
}

func Test_S21_024_Empty_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_S21_024_Empty_KeyValuePair", func() {
		// Act
		kv := corestr.Empty.KeyValuePair()

		// Assert
		actual := args.Map{"result": kv == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.KeyValuePair returns non-nil", actual)
	})
}

func Test_S21_025_Empty_KeyValueCollection(t *testing.T) {
	safeTest(t, "Test_S21_025_Empty_KeyValueCollection", func() {
		// Act
		kvc := corestr.Empty.KeyValueCollection()

		// Assert
		actual := args.Map{"result": kvc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.KeyValueCollection returns non-nil", actual)
	})
}

func Test_S21_026_Empty_LinkedCollections(t *testing.T) {
	safeTest(t, "Test_S21_026_Empty_LinkedCollections", func() {
		// Act
		lc := corestr.Empty.LinkedCollections()

		// Assert
		actual := args.Map{"result": lc == nil || lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.LinkedCollections returns empty", actual)
	})
}

func Test_S21_027_Empty_LeftRight(t *testing.T) {
	safeTest(t, "Test_S21_027_Empty_LeftRight", func() {
		// Act
		lr := corestr.Empty.LeftRight()

		// Assert
		actual := args.Map{"result": lr == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.LeftRight returns non-nil", actual)
	})
}

func Test_S21_028_Empty_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_S21_028_Empty_SimpleStringOnce", func() {
		// Act
		sso := corestr.Empty.SimpleStringOnce()

		// Assert
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.SimpleStringOnce returns uninitialized", actual)
	})
}

func Test_S21_029_Empty_SimpleStringOncePtr(t *testing.T) {
	safeTest(t, "Test_S21_029_Empty_SimpleStringOncePtr", func() {
		// Act
		sso := corestr.Empty.SimpleStringOncePtr()

		// Assert
		actual := args.Map{"result": sso == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.SimpleStringOncePtr returns non-nil", actual)
	})
}

func Test_S21_030_Empty_Hashset(t *testing.T) {
	safeTest(t, "Test_S21_030_Empty_Hashset", func() {
		// Act
		hs := corestr.Empty.Hashset()

		// Assert
		actual := args.Map{"result": hs == nil || hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.Hashset returns empty hashset", actual)
	})
}

func Test_S21_031_Empty_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_S21_031_Empty_HashsetsCollection", func() {
		// Act
		hsc := corestr.Empty.HashsetsCollection()

		// Assert
		actual := args.Map{"result": hsc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.HashsetsCollection returns non-nil", actual)
	})
}

func Test_S21_032_Empty_Hashmap(t *testing.T) {
	safeTest(t, "Test_S21_032_Empty_Hashmap", func() {
		// Act
		hm := corestr.Empty.Hashmap()

		// Assert
		actual := args.Map{"result": hm == nil || hm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.Hashmap returns empty hashmap", actual)
	})
}

func Test_S21_033_Empty_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_S21_033_Empty_CharCollectionMap", func() {
		// Act
		ccm := corestr.Empty.CharCollectionMap()

		// Assert
		actual := args.Map{"result": ccm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.CharCollectionMap returns non-nil", actual)
	})
}

func Test_S21_034_Empty_KeyValuesCollection(t *testing.T) {
	safeTest(t, "Test_S21_034_Empty_KeyValuesCollection", func() {
		// Act
		kvc := corestr.Empty.KeyValuesCollection()

		// Assert
		actual := args.Map{"result": kvc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.KeyValuesCollection returns non-nil", actual)
	})
}

func Test_S21_035_Empty_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_S21_035_Empty_CollectionsOfCollection", func() {
		// Act
		coc := corestr.Empty.CollectionsOfCollection()

		// Assert
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.CollectionsOfCollection returns non-nil", actual)
	})
}

func Test_S21_036_Empty_CharHashsetMap(t *testing.T) {
	safeTest(t, "Test_S21_036_Empty_CharHashsetMap", func() {
		// Act
		chm := corestr.Empty.CharHashsetMap()

		// Assert
		actual := args.Map{"result": chm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty.CharHashsetMap returns non-nil", actual)
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
		actual := args.Map{"result": result != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- empty input", actual)
	})
}

func Test_S21_041_AnyToString_with_value_no_field_name(t *testing.T) {
	safeTest(t, "Test_S21_041_AnyToString_with_value_no_field_name", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, "hello")

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- has value", actual)
	})
}

func Test_S21_042_AnyToString_with_value_include_field_name(t *testing.T) {
	safeTest(t, "Test_S21_042_AnyToString_with_value_include_field_name", func() {
		// Arrange & Act
		result := corestr.AnyToString(true, "hello")

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- include field name", actual)
	})
}

func Test_S21_043_AnyToString_with_int(t *testing.T) {
	safeTest(t, "Test_S21_043_AnyToString_with_int", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, 42)

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- int input", actual)
	})
}

func Test_S21_044_AnyToString_with_pointer(t *testing.T) {
	safeTest(t, "Test_S21_044_AnyToString_with_pointer", func() {
		// Arrange
		val := "test"

		// Act
		result := corestr.AnyToString(false, &val)

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- pointer input", actual)
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
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- struct input with field names", actual)
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
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- struct pointer input", actual)
	})
}
