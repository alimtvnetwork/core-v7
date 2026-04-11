package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── ReflectValueKind ──

func Test_ReflectValueKind_InvalidModel_FromReflectValueKindInva(t *testing.T) {
	// Arrange
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")

	// Act
	actual := args.Map{
		"isValid":   rvk.IsValid,
		"isInvalid": rvk.IsInvalid(),
		"hasError":  rvk.HasError(),
		"isEmptyErr": rvk.IsEmptyError(),
		"errMsg":    rvk.Error.Error(),
	}

	// Assert
	expected := args.Map{
		"isValid": false, "isInvalid": true,
		"hasError": true, "isEmptyErr": false,
		"errMsg": "test error",
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns error -- InvalidModel", actual)
}

func Test_ReflectValueKind_Nil(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{
		"isInvalid":  rvk.IsInvalid(),
		"hasError":   rvk.HasError(),
		"isEmptyErr": rvk.IsEmptyError(),
		"actual":     rvk.ActualInstance() == nil,
		"pkgPath":    rvk.PkgPath(),
		"typeName":   rvk.TypeName(),
		"pointerRv":  rvk.PointerRv() == nil,
		"pointerInf": rvk.PointerInterface() == nil,
	}

	// Assert
	expected := args.Map{
		"isInvalid": true, "hasError": false,
		"isEmptyErr": true, "actual": true,
		"pkgPath": "", "typeName": "",
		"pointerRv": true, "pointerInf": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns nil -- nil receiver", actual)
}

func Test_ReflectValueKind_Valid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf("hello"),
		Kind:            reflect.String,
		Error:           nil,
	}

	// Act
	actual := args.Map{
		"isValid":    rvk.IsValid,
		"isInvalid":  rvk.IsInvalid(),
		"hasError":   rvk.HasError(),
		"actual":     rvk.ActualInstance(),
		"pkgPathEmp": rvk.PkgPath() == "",
		"typeName":   rvk.TypeName() != "",
	}

	// Assert
	expected := args.Map{
		"isValid": true, "isInvalid": false,
		"hasError": false, "actual": "hello",
		"pkgPathEmp": true, "typeName": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns non-empty -- valid", actual)
}

func Test_ReflectValueKind_PointerRv_Invalid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
		Kind:            0,
	}
	rv := rvk.PointerRv()

	// Act
	actual := args.Map{"notNil": rv != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns error -- PointerRv invalid", actual)
}

// ── ReflectValue ──

func Test_ReflectValue_Fields_FromReflectValueKindInva_FromReflectValueKindInva(t *testing.T) {
	// Arrange
	rv := reflectmodel.ReflectValue{
		TypeName:     "TestType",
		FieldsNames:  []string{"Field1", "Field2"},
		MethodsNames: []string{"Method1"},
		RawData:      "raw",
	}

	// Act
	actual := args.Map{
		"typeName":    rv.TypeName,
		"fieldsLen":   len(rv.FieldsNames),
		"methodsLen":  len(rv.MethodsNames),
		"rawData":     rv.RawData,
	}

	// Assert
	expected := args.Map{
		"typeName": "TestType", "fieldsLen": 2,
		"methodsLen": 1, "rawData": "raw",
	}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns correct value -- fields", actual)
}

// ── FieldProcessor ──

func Test_FieldProcessor_Nil(t *testing.T) {
	// Arrange
	var fp *reflectmodel.FieldProcessor

	// Act
	actual := args.Map{
		"isFieldType": fp.IsFieldType(reflect.TypeOf("")),
		"isFieldKind": fp.IsFieldKind(reflect.String),
	}

	// Assert
	expected := args.Map{
		"isFieldType": false, "isFieldKind": false,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns nil -- nil", actual)
}

func Test_FieldProcessor_Valid(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: strType,
	}

	// Act
	actual := args.Map{
		"isFieldType":  fp.IsFieldType(strType),
		"isFieldKind":  fp.IsFieldKind(reflect.String),
		"wrongKind":    fp.IsFieldKind(reflect.Int),
	}

	// Assert
	expected := args.Map{
		"isFieldType": true, "isFieldKind": true, "wrongKind": false,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns non-empty -- valid", actual)
}

// ── MethodProcessor ──

func Test_MethodProcessor_Nil_FromReflectValueKindInva(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{
		"isInvalid":   mp.IsInvalid(),
		"hasValidFunc": mp.HasValidFunc(),
		"func":         mp.Func() == nil,
		"returnLen":    mp.ReturnLength(),
	}

	// Assert
	expected := args.Map{
		"isInvalid": true, "hasValidFunc": false,
		"func": true, "returnLen": -1,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- nil", actual)
}

func Test_MethodProcessor_IsEqual_BothNil_FromReflectValueKindInva(t *testing.T) {
	// Arrange
	var mp1, mp2 *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{
		"equal":    mp1.IsEqual(mp2),
		"notEqual": mp1.IsNotEqual(mp2),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- IsEqual both nil", actual)
}

func Test_MethodProcessor_IsEqual_OneNil_FromReflectValueKindInva(t *testing.T) {
	// Arrange
	mp := &reflectmodel.MethodProcessor{Name: "Test"}

	// Act
	actual := args.Map{"equal": mp.IsEqual(nil)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- IsEqual one nil", actual)
}

func Test_MethodProcessor_GetType_Invalid(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"isNil": mp.GetType() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- GetType nil", actual)
}

func Test_MethodProcessor_GetInOutArgsTypes_Invalid(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{
		"inLen":       len(mp.GetInArgsTypes()),
		"outLen":      len(mp.GetOutArgsTypes()),
		"inNamesLen":  len(mp.GetInArgsTypesNames()),
	}

	// Assert
	expected := args.Map{
		"inLen": 0,
		"outLen": 0,
		"inNamesLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- GetInOutArgsTypes nil", actual)
}
