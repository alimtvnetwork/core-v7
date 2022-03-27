package coredynamic

import (
	"reflect"

	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type reflectGetter struct{}

// PublicValuesMapStruct
//
//  returns structs fields map[string]Interface{}
//  map[string:fieldName]Interface{}:PublicValue
//
//  Only public values will be collected into map values
func (it reflectGetter) PublicValuesMapStruct(anyItem interface{}) (
	map[string]interface{}, error,
) {
	if reflectinternal.IsNull(anyItem) {
		return map[string]interface{}{},
			errcore.
				NullResultType.
				ErrorNoRefs(
					"null given to expand map[name]value, failed")
	}

	return ReflectGetterUsingReflectValue.PublicValuesMapStruct(
		reflect.ValueOf(anyItem))
}

// FieldNameWithValuesMap
//
//  returns structs fields map[string]Interface{}
//  map[string:fieldName]reflect.Type:fieldType
//
//  unlike PublicValuesMapStruct to map it collects
//  all fields with values including the private ones.
//
// However, this one will be slower in performance than PublicValuesMapStruct.
func (it reflectGetter) FieldNameWithValuesMap(anyItem interface{}) (
	map[string]interface{}, error,
) {
	if reflectinternal.IsNull(anyItem) {
		return map[string]interface{}{},
			errcore.
				NullResultType.
				ErrorNoRefs(
					"null given to expand map[name]value, failed")
	}

	return ReflectGetterUsingReflectValue.FieldNameWithValuesMap(
		reflect.ValueOf(anyItem))
}

// FieldNamesMap
//
//  returns structs fields map[string]bool names
//  map[string:fieldName]bool:exists
func (it reflectGetter) FieldNamesMap(
	anyItem interface{},
) (
	map[string]bool, error,
) {
	if reflectinternal.IsNull(anyItem) {
		return map[string]bool{},
			errcore.
				NullResultType.
				ErrorNoRefs(
					"null given to expand map[name]bool, failed")
	}

	return ReflectGetterUsingReflectValue.FieldNamesMap(
		reflect.ValueOf(anyItem))
}

// StructFieldsMap
//
//  returns structs all fields (public, private) map[string]reflect.StructField
//  map[string:fieldName]reflect.StructField:StructField
func (it reflectGetter) StructFieldsMap(
	anyItem interface{},
) map[string]reflect.StructField {
	if reflectinternal.IsNull(anyItem) {
		return map[string]reflect.StructField{}
	}

	return ReflectGetterUsingReflectValue.StructFieldsMap(
		reflect.ValueOf(anyItem))
}

// NullFieldsMap
//
//  returns structs all fields (public, private) map[string]bool
//  null fields map only
func (it reflectGetter) NullFieldsMap(
	anyItem interface{},
) map[string]bool {
	if reflectinternal.IsNull(anyItem) {
		return map[string]bool{}
	}

	return ReflectGetterUsingReflectValue.NullFieldsMap(
		defaultMaxLevelOfReflection,
		reflect.ValueOf(anyItem))
}

// NullOrZeroFieldsMap
//
//  returns structs all fields (public, private) map[string]bool
//  null or zero fields map only
func (it reflectGetter) NullOrZeroFieldsMap(
	anyItem interface{},
) map[string]bool {
	if reflectinternal.IsNull(anyItem) {
		return map[string]bool{}
	}

	return ReflectGetterUsingReflectValue.NullOrZeroFieldsMap(
		defaultMaxLevelOfReflection,
		reflect.ValueOf(anyItem))
}
