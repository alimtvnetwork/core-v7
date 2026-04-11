package enumimpl

import "testing"

// In-package tests for unexported interface branches.
// These cover valueByter and exactValueByter type assertions
// in DynamicMap.KeyValueByte and DynamicMap.KeyValueInt.

type stubValueByter struct{ v byte }

func (s stubValueByter) Value() byte { return s.v }

type stubExactValueByter struct{ v byte }

func (s stubExactValueByter) ValueByte() byte { return s.v }

func Test_KeyValueByte_ValueByter(t *testing.T) {
	// Arrange
	dm := DynamicMap(map[string]any{"k": stubValueByter{v: 42}})

	// Act
	val, found, failed := dm.KeyValueByte("k")

	// Assert
	if !found || failed || val != 42 {
		t.Errorf("expected (42, true, false), got (%d, %v, %v)", val, found, failed)
	}
}

func Test_KeyValueByte_ExactValueByter(t *testing.T) {
	// Arrange
	dm := DynamicMap(map[string]any{"k": stubExactValueByter{v: 77}})

	// Act
	val, found, failed := dm.KeyValueByte("k")

	// Assert
	if !found || failed || val != 77 {
		t.Errorf("expected (77, true, false), got (%d, %v, %v)", val, found, failed)
	}
}

func Test_KeyValueInt_ValueByter(t *testing.T) {
	// Arrange
	dm := DynamicMap(map[string]any{"k": stubValueByter{v: 10}})

	// Act
	val, found, failed := dm.KeyValueInt("k")

	// Assert
	if !found || failed || val != 10 {
		t.Errorf("expected (10, true, false), got (%d, %v, %v)", val, found, failed)
	}
}

func Test_KeyValueInt_ExactValueByter(t *testing.T) {
	// Arrange
	dm := DynamicMap(map[string]any{"k": stubExactValueByter{v: 20}})

	// Act
	val, found, failed := dm.KeyValueInt("k")

	// Assert
	if !found || failed || val != 20 {
		t.Errorf("expected (20, true, false), got (%d, %v, %v)", val, found, failed)
	}
}

// DynamicMap.Set nil receiver branch
func Test_DynamicMap_Set_NilInit(t *testing.T) {
	// Arrange
	dm := DynamicMap(nil)
	dmPtr := &dm

	// Act
	isNew := dmPtr.Set("k", "v")

	// Assert
	if !isNew {
		t.Error("expected new key addition")
	}
}

// DynamicMap.AddNewOnly nil receiver branch
func Test_DynamicMap_AddNewOnly_NilInit(t *testing.T) {
	// Arrange
	dm := DynamicMap(nil)
	dmPtr := &dm

	// Act
	isAdded := dmPtr.AddNewOnly("k", "v")

	// Assert
	if !isAdded {
		t.Error("expected key to be added")
	}
}
