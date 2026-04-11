package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── ReflectValueKind — valid model ──

func Test_ReflectValueKind_Valid_FromReflectValueKindVali(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: rv,
		Kind:            rv.Kind(),
		Error:           nil,
	}
	// Act & Assert
	actual := args.Map{
		"isInvalid":  rvk.IsInvalid(),
		"hasError":   rvk.HasError(),
		"isEmptyErr": rvk.IsEmptyError(),
		"typeName":   rvk.TypeName(),
		"pkgPath":    rvk.PkgPath(),
		"actualVal":  rvk.ActualInstance(),
	}
	expected := args.Map{
		"isInvalid": false, "hasError": false,
		"isEmptyErr": true, "typeName": "<int Value>",
		"pkgPath": "", "actualVal": 42,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns non-empty -- valid", actual)
}

func Test_ReflectValueKind_PointerRv_Valid(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: rv,
		Kind:            rv.Kind(),
	}
	// Act
	ptrRv := rvk.PointerRv()
	ptrInf := rvk.PointerInterface()
	// Assert
	actual := args.Map{
		"ptrRvNotNil": ptrRv != nil,
		"ptrInfNotNil": ptrInf != nil,
	}
	expected := args.Map{
		"ptrRvNotNil": true,
		"ptrInfNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns error -- PointerRv valid", actual)
}

func Test_ReflectValueKind_PointerRv_Invalid_FromReflectValueKindVali(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
		Kind:            0,
	}
	// Act
	ptrRv := rvk.PointerRv()
	// Assert
	actual := args.Map{"ptrRvNotNil": ptrRv != nil}
	expected := args.Map{"ptrRvNotNil": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns error -- PointerRv invalid", actual)
}

// ── FieldProcessor ──

func Test_FieldProcessor_Methods(t *testing.T) {
	// Arrange
	type sample struct{ Name string }
	sampleType := reflect.TypeOf(sample{})
	field := sampleType.Field(0)
	fp := &reflectmodel.FieldProcessor{
		Name:      field.Name,
		Index:     0,
		Field:     field,
		FieldType: field.Type,
	}
	// Act & Assert
	actual := args.Map{
		"isFieldType":  fp.IsFieldType(reflect.TypeOf("")),
		"isFieldKind":  fp.IsFieldKind(reflect.String),
		"wrongKind":    fp.IsFieldKind(reflect.Int),
	}
	expected := args.Map{
		"isFieldType": true, "isFieldKind": true, "wrongKind": false,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns correct value -- methods", actual)
}

func Test_FieldProcessor_Nil_FromReflectValueKindVali(t *testing.T) {
	// Arrange
	var fp *reflectmodel.FieldProcessor
	// Act & Assert
	actual := args.Map{
		"isFieldType": fp.IsFieldType(reflect.TypeOf("")),
		"isFieldKind": fp.IsFieldKind(reflect.String),
	}
	expected := args.Map{
		"isFieldType": false,
		"isFieldKind": false,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns nil -- nil", actual)
}

// ── MethodProcessor ──

type cov4SampleStruct struct{}

func (s cov4SampleStruct) Hello(name string) string { return "Hello " + name }
func (s cov4SampleStruct) Greet() string             { return "Hi" }

func Test_MethodProcessor_Basic(t *testing.T) {
	// Arrange
	sType := reflect.TypeOf(cov4SampleStruct{})
	method, _ := sType.MethodByName("Hello")
	mp := &reflectmodel.MethodProcessor{
		Name:          method.Name,
		Index:         method.Index,
		ReflectMethod: method,
	}
	// Act & Assert
	actual := args.Map{
		"hasValidFunc":  mp.HasValidFunc(),
		"funcName":      mp.GetFuncName(),
		"isInvalid":     mp.IsInvalid(),
		"argsCount":     mp.ArgsCount(),
		"argsLen":       mp.ArgsLength(),
		"returnLen":     mp.ReturnLength(),
		"isPublic":      mp.IsPublicMethod(),
		"isPrivate":     mp.IsPrivateMethod(),
		"funcNotNil":    mp.Func() != nil,
		"typeNotNil":    mp.GetType() != nil,
	}
	expected := args.Map{
		"hasValidFunc": true, "funcName": "Hello",
		"isInvalid": false, "argsCount": 2, "argsLen": 2,
		"returnLen": 1, "isPublic": true, "isPrivate": false,
		"funcNotNil": true, "typeNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- basic", actual)
}

func Test_MethodProcessor_Nil_FromReflectValueKindVali(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	// Act & Assert
	actual := args.Map{
		"hasValidFunc": mp.HasValidFunc(),
		"isInvalid":    mp.IsInvalid(),
		"returnLen":    mp.ReturnLength(),
		"funcNil":      mp.Func() == nil,
		"typeNil":      mp.GetType() == nil,
	}
	expected := args.Map{
		"hasValidFunc": false, "isInvalid": true,
		"returnLen": -1, "funcNil": true, "typeNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- nil", actual)
}

func Test_MethodProcessor_GetInOutArgsTypes(t *testing.T) {
	// Arrange
	sType := reflect.TypeOf(cov4SampleStruct{})
	method, _ := sType.MethodByName("Hello")
	mp := &reflectmodel.MethodProcessor{
		Name:          method.Name,
		Index:         method.Index,
		ReflectMethod: method,
	}
	// Act
	inTypes := mp.GetInArgsTypes()
	outTypes := mp.GetOutArgsTypes()
	inNames := mp.GetInArgsTypesNames()
	// Call again to hit cache
	inTypes2 := mp.GetInArgsTypes()
	outTypes2 := mp.GetOutArgsTypes()
	inNames2 := mp.GetInArgsTypesNames()
	// Assert
	actual := args.Map{
		"inLen":       len(inTypes),
		"outLen":      len(outTypes),
		"inNamesLen":  len(inNames),
		"inLen2":      len(inTypes2),
		"outLen2":     len(outTypes2),
		"inNames2Len": len(inNames2),
	}
	expected := args.Map{
		"inLen": 2, "outLen": 1, "inNamesLen": 2,
		"inLen2": 2, "outLen2": 1, "inNames2Len": 2,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- GetInOutArgsTypes", actual)
}

func Test_MethodProcessor_GetInOutArgsTypes_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	// Act
	inTypes := mp.GetInArgsTypes()
	outTypes := mp.GetOutArgsTypes()
	inNames := mp.GetInArgsTypesNames()
	// Assert
	actual := args.Map{
		"inLen": len(inTypes),
		"outLen": len(outTypes),
		"inNamesLen": len(inNames),
	}
	expected := args.Map{
		"inLen": 0,
		"outLen": 0,
		"inNamesLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- nil GetInOutArgsTypes", actual)
}

func Test_MethodProcessor_Invoke(t *testing.T) {
	// Arrange
	s := cov4SampleStruct{}
	sType := reflect.TypeOf(s)
	method, _ := sType.MethodByName("Hello")
	mp := &reflectmodel.MethodProcessor{
		Name:          method.Name,
		Index:         method.Index,
		ReflectMethod: method,
	}
	// Act
	results, err := mp.Invoke(s, "World")
	// Assert
	actual := args.Map{
		"err": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"err": true,
		"result": "Hello World",
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- Invoke", actual)
}

func Test_MethodProcessor_Invoke_WrongArgCount(t *testing.T) {
	// Arrange
	sType := reflect.TypeOf(cov4SampleStruct{})
	method, _ := sType.MethodByName("Hello")
	mp := &reflectmodel.MethodProcessor{
		Name:          method.Name,
		Index:         method.Index,
		ReflectMethod: method,
	}
	// Act
	_, err := mp.Invoke("wrong")
	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- Invoke wrong arg count", actual)
}

func Test_MethodProcessor_GetFirstResponseOfInvoke_FromReflectValueKindVali(t *testing.T) {
	// Arrange
	s := cov4SampleStruct{}
	sType := reflect.TypeOf(s)
	method, _ := sType.MethodByName("Greet")
	mp := &reflectmodel.MethodProcessor{
		Name:          method.Name,
		Index:         method.Index,
		ReflectMethod: method,
	}
	// Act
	result, err := mp.GetFirstResponseOfInvoke(s)
	// Assert
	actual := args.Map{
		"err": err == nil,
		"result": result,
	}
	expected := args.Map{
		"err": true,
		"result": "Hi",
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- GetFirstResponseOfInvoke", actual)
}

func Test_MethodProcessor_IsEqual_FromReflectValueKindVali(t *testing.T) {
	// Arrange
	sType := reflect.TypeOf(cov4SampleStruct{})
	method1, _ := sType.MethodByName("Hello")
	method2, _ := sType.MethodByName("Hello")
	mp1 := &reflectmodel.MethodProcessor{Name: method1.Name, ReflectMethod: method1}
	mp2 := &reflectmodel.MethodProcessor{Name: method2.Name, ReflectMethod: method2}
	// Act & Assert
	actual := args.Map{
		"selfEqual":   mp1.IsEqual(mp1),
		"sameMethod":  mp1.IsEqual(mp2),
		"notEqual":    mp1.IsNotEqual(mp2),
		"bothNil":     (*reflectmodel.MethodProcessor)(nil).IsEqual(nil),
		"oneNil":      mp1.IsEqual(nil),
	}
	expected := args.Map{
		"selfEqual": true, "sameMethod": true,
		"notEqual": false, "bothNil": true, "oneNil": false,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- IsEqual", actual)
}

func Test_MethodProcessor_VerifyInOutArgs(t *testing.T) {
	// Arrange
	sType := reflect.TypeOf(cov4SampleStruct{})
	method, _ := sType.MethodByName("Hello")
	mp := &reflectmodel.MethodProcessor{Name: method.Name, ReflectMethod: method}
	// Act
	inOk, inErr := mp.VerifyInArgs([]any{cov4SampleStruct{}, "test"})
	outOk, outErr := mp.VerifyOutArgs([]any{"result"})
	// Assert
	actual := args.Map{
		"inOk": inOk,
		"inErr": inErr == nil,
		"outOk": outOk,
		"outErr": outErr == nil,
	}
	expected := args.Map{
		"inOk": true,
		"inErr": true,
		"outOk": true,
		"outErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- VerifyInOutArgs", actual)
}

func Test_MethodProcessor_ValidateMethodArgs_TypeMismatch(t *testing.T) {
	// Arrange
	sType := reflect.TypeOf(cov4SampleStruct{})
	method, _ := sType.MethodByName("Hello")
	mp := &reflectmodel.MethodProcessor{Name: method.Name, ReflectMethod: method}
	// Act
	err := mp.ValidateMethodArgs([]any{cov4SampleStruct{}, 42})
	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns non-empty -- ValidateMethodArgs type mismatch", actual)
}

// ── rvUtils ──

func Test_RvUtils_IndexToPosition(t *testing.T) {
	// Arrange
	u := reflect.TypeOf(cov4SampleStruct{})
	_ = u // just need the utils struct
	// The utils are not exported, but tested via MethodProcessor's error messages
	// Test VerifyReflectTypes with mismatched lengths
	sType := reflect.TypeOf(cov4SampleStruct{})
	method, _ := sType.MethodByName("Hello")
	mp := &reflectmodel.MethodProcessor{Name: method.Name, ReflectMethod: method}
	// Act
	inOk, inErr := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})
	// Assert
	actual := args.Map{
		"ok": inOk,
		"hasErr": inErr != nil,
	}
	expected := args.Map{
		"ok": false,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "rvUtils returns correct value -- VerifyReflectTypes length mismatch", actual)
}

// ── ReflectValue struct ──

func Test_ReflectValue_Struct(t *testing.T) {
	// Arrange
	rv := reflectmodel.ReflectValue{
		TypeName:     "MyType",
		FieldsNames:  []string{"Field1", "Field2"},
		MethodsNames: []string{"Method1"},
		RawData:      42,
	}
	// Assert
	actual := args.Map{
		"typeName":    rv.TypeName,
		"fieldsLen":   len(rv.FieldsNames),
		"methodsLen":  len(rv.MethodsNames),
		"rawData":     rv.RawData,
	}
	expected := args.Map{
		"typeName": "MyType", "fieldsLen": 2, "methodsLen": 1, "rawData": 42,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns correct value -- struct", actual)
}
