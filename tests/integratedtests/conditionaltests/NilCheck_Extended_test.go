package conditionaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_NilCheck_Extended_Verification(t *testing.T) {
	// Arrange
	// nil case
	result := conditional.NilCheck(nil, "wasNil", "wasNotNil")

	// Act
	actual := args.Map{"result": result != "wasNil"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return onNil", actual)

	// non-nil case
	result2 := conditional.NilCheck("something", "wasNil", "wasNotNil")
	actual = args.Map{"result": result2 != "wasNotNil"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil should return onNonNil", actual)
}

func Test_DefOnNil_Extended_Verification(t *testing.T) {
	// Arrange
	// nil case
	result := conditional.DefOnNil(nil, "default")

	// Act
	actual := args.Map{"result": result != "default"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return default", actual)

	// non-nil case
	result2 := conditional.DefOnNil("original", "default")
	actual = args.Map{"result": result2 != "original"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil should return original", actual)
}

func Test_NilOrEmptyStr_Extended_Verification(t *testing.T) {
	// Arrange
	// nil case
	result := conditional.NilOrEmptyStr(nil, "empty", "full")

	// Act
	actual := args.Map{"result": result != "empty"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return onNilOrEmpty", actual)

	// empty case
	empty := ""
	result2 := conditional.NilOrEmptyStr(&empty, "empty", "full")
	actual = args.Map{"result": result2 != "empty"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty string should return onNilOrEmpty", actual)

	// non-empty case
	val := "hello"
	result3 := conditional.NilOrEmptyStr(&val, "empty", "full")
	actual = args.Map{"result": result3 != "full"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-empty should return onNonNilOrNonEmpty", actual)
}

func Test_NilOrEmptyStrPtr_Extended_Verification(t *testing.T) {
	// Arrange
	// nil case
	result := conditional.NilOrEmptyStrPtr(nil, "empty", "full")

	// Act
	actual := args.Map{"result": result == nil || *result != "empty"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return pointer to onNilOrEmpty", actual)

	// non-empty case
	val := "hello"
	result2 := conditional.NilOrEmptyStrPtr(&val, "empty", "full")
	actual = args.Map{"result": result2 == nil || *result2 != "full"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-empty should return pointer to onNonNilOrNonEmpty", actual)
}

func Test_BoolByOrder_Extended_Verification(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolByOrder()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return false", actual)
	actual = args.Map{"result": conditional.BoolByOrder(false, true)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "one true should return true", actual)
	actual = args.Map{"result": conditional.BoolByOrder(false, false)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all false should return false", actual)
}

func Test_StringsIndexVal_Extended_Verification(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}
	result := conditional.StringsIndexVal(true, slice, 0, 2)

	// Act
	actual := args.Map{"result": result != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "true should return index 0", actual)

	result2 := conditional.StringsIndexVal(false, slice, 0, 2)
	actual = args.Map{"result": result2 != "c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "false should return index 2", actual)
}

func Test_IfSliceAny_Extended_Verification(t *testing.T) {
	// Arrange
	result := conditional.IfSliceAny(true, []any{1, 2}, []any{3})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "true should return first slice", actual)

	result2 := conditional.IfSliceAny(false, []any{1, 2}, []any{3})
	actual = args.Map{"result": len(result2) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "false should return second slice", actual)
}

func Test_IfFuncAny_Extended_Verification(t *testing.T) {
	// Arrange
	result := conditional.IfFuncAny(true, func() any { return "yes" }, func() any { return "no" })

	// Act
	actual := args.Map{"result": result != "yes"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "true should return yes", actual)

	result2 := conditional.IfFuncAny(false, func() any { return "yes" }, func() any { return "no" })
	actual = args.Map{"result": result2 != "no"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "false should return no", actual)
}

func Test_Setter_Extended_Verification(t *testing.T) {
	// This covers conditional.Setter via issetter
	// Just ensure it doesn't panic
	_ = conditional.StringDefault(true, "yes")
	_ = conditional.StringDefault(false, "yes")
}
