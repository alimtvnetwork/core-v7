package coredynamic

import (
	"reflect"
	"unsafe"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type reflectGetUsingReflectValue struct{}

// PublicValuesMapStruct
//
//  returns structs fields map[string]Interface{}
//  map[string:fieldName]Interface{}:PublicValue
//
//  Only public values will be collected into map values
func (it reflectGetUsingReflectValue) PublicValuesMapStruct(structValue reflect.Value) (
	map[string]interface{}, error,
) {
	if structValue.Kind() != reflect.Struct {
		return nil, errcore.Expected.ReflectButFound(
			reflect.Struct,
			structValue.Kind())
	}

	structType := structValue.Type()
	structNumFields := structType.NumField()
	fieldToValueMap := make(map[string]interface{}, structNumFields)

	for i := 0; i < structNumFields; i++ {
		fieldStruct := structType.Field(i)

		// ignore unexported fields
		if fieldStruct.PkgPath != "" {
			continue
		}

		field := structValue.Field(i)
		fieldToValueMap[fieldStruct.Name] = field.Interface()
	}

	return fieldToValueMap, nil
}

// FieldNameWithTypeMap
//
//  returns structs fields map[string]Interface{}
//  map[string:fieldName]reflect.Type:fieldType
//
//  Only public values will be collected into map values
func (it reflectGetUsingReflectValue) FieldNameWithTypeMap(
	rv reflect.Value,
) map[string]reflect.Type {
	structValue := rv
	structValueKind := structValue.Kind()

	for structValueKind == reflect.Ptr || structValueKind == reflect.Interface {
		// mutating dangerous code
		structValue = structValue.Elem()
		structValueKind = structValue.Kind()
	}

	if !structValue.IsValid() || structValueKind != reflect.Struct {
		return nil
	}

	structType := structValue.Type()
	fieldsLength := structType.NumField()
	fieldsHashset :=
		make(
			map[string]reflect.Type,
			fieldsLength)

	var name string

	for i := 0; i < fieldsLength; i++ {
		field := structType.Field(i)
		name = field.Name
		fieldsHashset[name] = field.Type
	}

	return fieldsHashset
}

// FieldNameWithValuesMap
//
//  returns structs all fields (public, private) map[string]Interface{}
//  map[string:fieldName]interface{}:fieldValuePublicOrPrivate
//
//  unlike PublicValuesMapStruct to map it collects
//  all fields with values including the private ones.
//
// However, this one will be slower in performance than PublicValuesMapStruct.
func (it reflectGetUsingReflectValue) FieldNameWithValuesMap(
	structValue reflect.Value,
) (
	map[string]interface{}, error,
) {
	structType := structValue.Type()
	structNumFields := structType.NumField()
	fieldToValueMap := make(map[string]interface{}, structNumFields)

	// structValue is not addressable, create a temporary copy
	if !structValue.CanAddr() {
		newType := reflect.New(structType).Elem()
		newType.Set(structValue)
		// structValue is now addressable
		structValue = newType
	}

	for i := 0; i < structNumFields; i++ {
		fieldType := structType.Field(i)
		fieldValue := structValue.Field(i)

		if fieldType.PkgPath != "" {
			unexportedField := reflect.NewAt(
				fieldType.Type,
				unsafe.Pointer(fieldValue.UnsafeAddr())).Elem()
			fieldToValueMap[fieldType.Name] = unexportedField.Interface()
		} else {
			fieldToValueMap[fieldType.Name] = fieldValue.Interface()
		}
	}

	return fieldToValueMap, nil
}

// FieldNamesMap
//
//  returns structs fields map[string]bool
//  map[string:fieldName]bool
func (it reflectGetUsingReflectValue) FieldNamesMap(
	rv reflect.Value,
) (map[string]bool, error) {
	structValue := rv
	structValueKind := structValue.Kind()

	for structValueKind == reflect.Ptr || structValueKind == reflect.Interface {
		// mutating dangerous code
		structValue = structValue.Elem()
		structValueKind = structValue.Kind()
	}

	if !structValue.IsValid() || structValueKind != reflect.Struct {
		return map[string]bool{},
			errcore.Expected.ReflectButFound(
				reflect.Struct, structValueKind)
	}

	structType := structValue.Type()
	fieldsLength := structType.NumField()
	fieldsMap := make(
		map[string]bool,
		fieldsLength+1)

	for i := 0; i < fieldsLength; i++ {
		name := structType.Field(i).Name
		fieldsMap[name] = true
	}

	return fieldsMap, nil
}

// StructFieldsMap
//
//  returns structs all fields (public, private) map[string]reflect.StructField
//  map[string:fieldName]reflect.StructField:StructField
func (it reflectGetUsingReflectValue) StructFieldsMap(
	rv reflect.Value,
) map[string]reflect.StructField {
	structValue := rv
	structValueKind := structValue.Kind()

	for structValueKind == reflect.Ptr || structValueKind == reflect.Interface {
		// mutating dangerous code
		structValue = structValue.Elem()
		structValueKind = structValue.Kind()
	}

	if !structValue.IsValid() || structValueKind != reflect.Struct {
		return nil
	}

	structType := structValue.Type()
	fieldsLength := structType.NumField()
	fieldsHashset :=
		make(
			map[string]reflect.StructField,
			fieldsLength)

	var name string

	for i := 0; i < fieldsLength; i++ {
		field := structType.Field(i)
		name = field.Name
		fieldsHashset[name] = field
	}

	return fieldsHashset
}

// NullFieldsMap
//
//  returns structs all fields (public, private) map[string]bool
//  null fields map only
func (it reflectGetUsingReflectValue) NullFieldsMap(
	level int,
	reflectVal reflect.Value,
) map[string]bool {
	structType := reflectVal.Type()
	structValueKind := reflectVal.Kind()
	hasLevel := level > constants.InvalidIndex
	structValue := reflectVal

	// reducing ****ToValue to ToValue
	for structValueKind == reflect.Ptr || structValueKind == reflect.Interface {
		// mutating dangerous code
		structValue = structValue.Elem()
		structValueKind = structValue.Kind()

		level--
		if hasLevel && level <= 0 {
			break
		}
	}

	if !structValue.IsValid() || structValueKind != reflect.Struct {
		return map[string]bool{}
	}

	structNumFields := structType.NumField()
	hashset := make(
		map[string]bool,
		structNumFields+1)
	var fieldValue reflect.Value
	var fieldType reflect.StructField

	for i := 0; i < structNumFields; i++ {
		fieldValue = structValue.Field(i)

		if reflectinternal.IsNullUsingReflectValue(fieldValue) {
			fieldType = structType.Field(i)
			hashset[fieldType.Name] = true
		}
	}

	return hashset
}

// NullOrZeroFieldsMap
//
//  returns structs all fields (public, private) map[string]bool
//  null or zero fields map only
func (it reflectGetUsingReflectValue) NullOrZeroFieldsMap(
	level int,
	reflectVal reflect.Value,
) map[string]bool {
	structType := reflectVal.Type()
	structValueKind := reflectVal.Kind()
	hasLevel := level > constants.InvalidIndex
	structValue := reflectVal

	// reducing ****ToValue to ToValue
	for structValueKind == reflect.Ptr || structValueKind == reflect.Interface {
		// mutating dangerous code
		structValue = structValue.Elem()
		structValueKind = structValue.Kind()

		level--
		if hasLevel && level <= 0 {
			break
		}
	}

	if !structValue.IsValid() || structValueKind != reflect.Struct {
		return map[string]bool{}
	}

	structNumFields := structType.NumField()
	hashset := make(
		map[string]bool,
		structNumFields+1)
	var fieldValue reflect.Value
	var fieldType reflect.StructField

	for i := 0; i < structNumFields; i++ {
		fieldValue = structValue.Field(i)

		if reflectinternal.IsZeroReflectValue(fieldValue) {
			fieldType = structType.Field(i)
			hashset[fieldType.Name] = true
		}
	}

	return hashset
}
