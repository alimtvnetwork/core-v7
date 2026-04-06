package reflectinternaltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// ── reflectUtils.MaxLimit ──

func Test_Cov4_Utils_MaxLimit(t *testing.T) {
	// Act
	actual := args.Map{
		"noLimit":   reflectinternal.Utils.MaxLimit(10, -1),
		"belowMax":  reflectinternal.Utils.MaxLimit(5, 10),
		"aboveMax":  reflectinternal.Utils.MaxLimit(15, 10),
		"equalMax":  reflectinternal.Utils.MaxLimit(10, 10),
	}

	// Assert
	expected := args.Map{
		"noLimit": 10, "belowMax": 5, "aboveMax": 10, "equalMax": 10,
	}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- MaxLimit", actual)
}

// ── reflectUtils.AppendArgs ──

func Test_Cov4_Utils_AppendArgs_Empty(t *testing.T) {
	// Arrange
	result := reflectinternal.Utils.AppendArgs("first", []any{})

	// Act
	actual := args.Map{"len": len(result), "first": result[0]}

	// Assert
	expected := args.Map{"len": 1, "first": "first"}
	expected.ShouldBeEqual(t, 0, "Utils returns empty -- AppendArgs empty", actual)
}

func Test_Cov4_Utils_AppendArgs_WithItems(t *testing.T) {
	// Arrange
	result := reflectinternal.Utils.AppendArgs("first", []any{"second"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Utils returns non-empty -- AppendArgs with items", actual)
}

// ── reflectUtils.IsReflectTypeMatch ──

func Test_Cov4_Utils_IsReflectTypeMatch(t *testing.T) {
	// Arrange
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")
	ok1, err1 := reflectinternal.Utils.IsReflectTypeMatch(intType, intType)
	ok2, err2 := reflectinternal.Utils.IsReflectTypeMatch(intType, strType)

	// Act
	actual := args.Map{
		"sameOk": ok1, "sameErr": err1 == nil,
		"diffOk": ok2, "diffErr": err2 != nil,
	}

	// Assert
	expected := args.Map{
		"sameOk": true, "sameErr": true,
		"diffOk": false, "diffErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- IsReflectTypeMatch", actual)
}

func Test_Cov4_Utils_IsReflectTypeMatch_InterfaceType(t *testing.T) {
	// Arrange
	var iface interface{}
	ifaceType := reflect.TypeOf(&iface).Elem()
	strType := reflect.TypeOf("")
	ok, err := reflectinternal.Utils.IsReflectTypeMatch(ifaceType, strType)

	// Act
	actual := args.Map{"ok": ok, "noErr": err == nil}

	// Assert
	expected := args.Map{"ok": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- IsReflectTypeMatch interface", actual)
}

// ── reflectUtils.IsReflectTypeMatchAny ──

func Test_Cov4_Utils_IsReflectTypeMatchAny(t *testing.T) {
	// Arrange
	ok, err := reflectinternal.Utils.IsReflectTypeMatchAny(42, 100)
	ok2, err2 := reflectinternal.Utils.IsReflectTypeMatchAny(42, "str")

	// Act
	actual := args.Map{
		"sameOk": ok, "sameErr": err == nil,
		"diffOk": ok2, "diffErr": err2 != nil,
	}

	// Assert
	expected := args.Map{
		"sameOk": true, "sameErr": true,
		"diffOk": false, "diffErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- IsReflectTypeMatchAny", actual)
}

// ── reflectUtils.VerifyReflectTypesAny ──

func Test_Cov4_Utils_VerifyReflectTypesAny_LenMismatch(t *testing.T) {
	// Arrange
	ok, err := reflectinternal.Utils.VerifyReflectTypesAny(
		[]any{1, 2},
		[]any{1},
	)

	// Act
	actual := args.Map{"ok": ok, "hasErr": err != nil}

	// Assert
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- VerifyReflectTypesAny len mismatch", actual)
}

func Test_Cov4_Utils_VerifyReflectTypesAny_Match(t *testing.T) {
	// Arrange
	ok, err := reflectinternal.Utils.VerifyReflectTypesAny(
		[]any{1, "a"},
		[]any{2, "b"},
	)

	// Act
	actual := args.Map{"ok": ok, "noErr": err == nil}

	// Assert
	expected := args.Map{"ok": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- VerifyReflectTypesAny match", actual)
}

func Test_Cov4_Utils_VerifyReflectTypesAny_Mismatch(t *testing.T) {
	// Arrange
	ok, err := reflectinternal.Utils.VerifyReflectTypesAny(
		[]any{1, "a"},
		[]any{2, 3},
	)

	// Act
	actual := args.Map{"ok": ok, "hasErr": err != nil}

	// Assert
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- VerifyReflectTypesAny mismatch", actual)
}

// ── reflectUtils.VerifyReflectTypes ──

func Test_Cov4_Utils_VerifyReflectTypes_LenMismatch(t *testing.T) {
	// Arrange
	ok, err := reflectinternal.Utils.VerifyReflectTypes(
		"TestRoot",
		[]reflect.Type{reflect.TypeOf(0)},
		[]reflect.Type{},
	)

	// Act
	actual := args.Map{"ok": ok, "hasErr": err != nil}

	// Assert
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- VerifyReflectTypes len mismatch", actual)
}

func Test_Cov4_Utils_VerifyReflectTypes_Match(t *testing.T) {
	// Arrange
	ok, err := reflectinternal.Utils.VerifyReflectTypes(
		"TestRoot",
		[]reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")},
		[]reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")},
	)

	// Act
	actual := args.Map{"ok": ok, "noErr": err == nil}

	// Assert
	expected := args.Map{"ok": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- VerifyReflectTypes match", actual)
}

// ── reflectUtils.PkgNameOnly / FullNameToPkgName ──

func Test_Cov4_Utils_PkgNameOnly(t *testing.T) {
	// Arrange
	result := reflectinternal.Utils.PkgNameOnly(Test_Cov4_Utils_PkgNameOnly)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Utils returns correct value -- PkgNameOnly", actual)
}
