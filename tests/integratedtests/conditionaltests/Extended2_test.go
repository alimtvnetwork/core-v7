package conditionaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
)

func Test_NilCheck_Extended_Verification(t *testing.T) {
	// nil case
	result := conditional.NilCheck(nil, "wasNil", "wasNotNil")
	if result != "wasNil" {
		t.Error("nil should return onNil")
	}

	// non-nil case
	result2 := conditional.NilCheck("something", "wasNil", "wasNotNil")
	if result2 != "wasNotNil" {
		t.Error("non-nil should return onNonNil")
	}
}

func Test_DefOnNil_Extended_Verification(t *testing.T) {
	// nil case
	result := conditional.DefOnNil(nil, "default")
	if result != "default" {
		t.Error("nil should return default")
	}

	// non-nil case
	result2 := conditional.DefOnNil("original", "default")
	if result2 != "original" {
		t.Error("non-nil should return original")
	}
}

func Test_NilOrEmptyStr_Extended_Verification(t *testing.T) {
	// nil case
	result := conditional.NilOrEmptyStr(nil, "empty", "full")
	if result != "empty" {
		t.Error("nil should return onNilOrEmpty")
	}

	// empty case
	empty := ""
	result2 := conditional.NilOrEmptyStr(&empty, "empty", "full")
	if result2 != "empty" {
		t.Error("empty string should return onNilOrEmpty")
	}

	// non-empty case
	val := "hello"
	result3 := conditional.NilOrEmptyStr(&val, "empty", "full")
	if result3 != "full" {
		t.Error("non-empty should return onNonNilOrNonEmpty")
	}
}

func Test_NilOrEmptyStrPtr_Extended_Verification(t *testing.T) {
	// nil case
	result := conditional.NilOrEmptyStrPtr(nil, "empty", "full")
	if result == nil || *result != "empty" {
		t.Error("nil should return pointer to onNilOrEmpty")
	}

	// non-empty case
	val := "hello"
	result2 := conditional.NilOrEmptyStrPtr(&val, "empty", "full")
	if result2 == nil || *result2 != "full" {
		t.Error("non-empty should return pointer to onNonNilOrNonEmpty")
	}
}

func Test_BoolByOrder_Extended_Verification(t *testing.T) {
	if conditional.BoolByOrder() {
		t.Error("empty should return false")
	}
	if !conditional.BoolByOrder(false, true) {
		t.Error("one true should return true")
	}
	if conditional.BoolByOrder(false, false) {
		t.Error("all false should return false")
	}
}

func Test_StringsIndexVal_Extended_Verification(t *testing.T) {
	slice := []string{"a", "b", "c"}
	result := conditional.StringsIndexVal(true, slice, 0, 2)
	if result != "a" {
		t.Error("true should return index 0")
	}

	result2 := conditional.StringsIndexVal(false, slice, 0, 2)
	if result2 != "c" {
		t.Error("false should return index 2")
	}
}

func Test_IfSliceAny_Extended_Verification(t *testing.T) {
	result := conditional.IfSliceAny(true, []any{1, 2}, []any{3})
	if len(result) != 2 {
		t.Error("true should return first slice")
	}

	result2 := conditional.IfSliceAny(false, []any{1, 2}, []any{3})
	if len(result2) != 1 {
		t.Error("false should return second slice")
	}
}

func Test_IfFuncAny_Extended_Verification(t *testing.T) {
	result := conditional.IfFuncAny(true, func() any { return "yes" }, func() any { return "no" })
	if result != "yes" {
		t.Error("true should return yes")
	}

	result2 := conditional.IfFuncAny(false, func() any { return "yes" }, func() any { return "no" })
	if result2 != "no" {
		t.Error("false should return no")
	}
}

func Test_Setter_Extended_Verification(t *testing.T) {
	// This covers conditional.Setter via issetter
	// Just ensure it doesn't panic
	_ = conditional.StringDefault(true, "yes")
	_ = conditional.StringDefault(false, "yes")
}
