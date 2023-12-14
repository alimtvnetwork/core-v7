package reflectinternal

import (
	"fmt"
	"reflect"
	"sort"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type mapConverter struct{}

func (it mapConverter) Length(i interface{}) int {
	return SliceConverter.Length(i)
}

// ToStringsRv
//
//	expectation : map[key:string]...value don't care.
func (it mapConverter) ToStringsRv(reflectVal reflect.Value) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToStringsRv(
			reflect.Indirect(reflectVal),
		)
	}

	if reflectVal.Kind() != reflect.Map {
		return []string{},
			it.notAMapErr(reflectVal)
	}

	mapKeys := reflectVal.MapKeys()
	length := len(mapKeys)
	keys := make([]string, length)

	for i, key := range reflectVal.MapKeys() {
		keyAny := key.Interface()
		keyAsString, isString := keyAny.(string)

		if !isString {
			return keys, it.notStringErr(keyAny)
		}

		keys[i] = keyAsString
	}

	return keys, nil
}

func (it mapConverter) notStringErr(keyAny interface{}) error {
	return fmt.Errorf("not string type : %T", keyAny)
}

func (it mapConverter) notAMapErr(reflectVal reflect.Value) error {
	return fmt.Errorf("reflection is not map but %s", reflectVal.String())
}

// ToKeysStrings
//
//	expectation : map[key:string]...value don't care.
func (it mapConverter) ToKeysStrings(i interface{}) ([]string, error) {
	return it.ToStrings(i)
}

// ToValuesAny
//
//	expectation : map[...]...value don't care.
func (it mapConverter) ToValuesAny(i interface{}) ([]interface{}, error) {
	if Is.Null(i) {
		return []interface{}{}, nil
	}

	rv := reflect.ValueOf(i)

	var list []interface{}

	err := Looper.MapForRv(
		rv, func(total int, index int, key, v interface{}) (err error) {
			list = append(list, v)

			return nil
		},
	)

	return list, err
}

// ToKeysAny
//
//	expectation : map[...]...value don't care.
func (it mapConverter) ToKeysAny(i interface{}) ([]interface{}, error) {
	if Is.Null(i) {
		return []interface{}{}, nil
	}

	rv := reflect.ValueOf(i)

	var list []interface{}

	err := Looper.MapForRv(
		rv, func(total int, index int, key, v interface{}) (err error) {
			list = append(list, key)

			return nil
		},
	)

	return list, err
}

// ToKeysValuesAny
//
//	expectation : map[string]...value don't care.
func (it mapConverter) ToKeysValuesAny(i interface{}) (keys []string, values []interface{}, err error) {
	if Is.Null(i) {
		return []string{}, []interface{}{}, nil
	}

	rv := reflect.ValueOf(i)
	toStringFunc := convertinteranl.AnyTo.SmartString

	err = Looper.MapForRv(
		rv, func(total int, index int, key, v interface{}) (err error) {
			keys = append(keys, toStringFunc(key))
			values = append(values, v)

			return nil
		},
	)

	return keys, values, err
}

// ToStrings
//
//	expectation : map[key:string]don't care values
func (it mapConverter) ToStrings(any interface{}) ([]string, error) {
	if Is.Null(any) {
		return []string{}, nil
	}

	reflectVal := reflect.ValueOf(any)

	return it.ToStringsRv(reflectVal)
}

func (it mapConverter) ToStringsMust(any interface{}) []string {
	reflectVal := reflect.ValueOf(any)

	mapKeys, err := it.ToStringsRv(reflectVal)

	if err != nil {
		panic(err)
	}

	return mapKeys
}

func (it mapConverter) ToSortedStrings(any interface{}) ([]string, error) {
	if Is.Null(any) {
		return []string{}, nil
	}

	reflectVal := reflect.ValueOf(any)

	keys, err := it.ToStringsRv(reflectVal)

	if err != nil {
		return keys, err
	}

	sort.Strings(keys)

	return keys, nil
}

func (it mapConverter) ToSortedStringsMust(any interface{}) []string {
	if Is.Null(any) {
		return []string{}
	}

	reflectVal := reflect.ValueOf(any)

	keys := it.ToStringsMust(reflectVal)
	sort.Strings(keys)

	return keys
}

// ToMapStringAnyRv
//
//	expectation : map[key:interface{}]interface{} to map[string]interface{}
func (it mapConverter) ToMapStringAnyRv(
	reflectVal reflect.Value,
) (map[string]interface{}, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToMapStringAnyRv(
			reflect.Indirect(reflectVal),
		)
	}

	if reflectVal.Kind() != reflect.Map {
		return map[string]interface{}{},
			it.notAMapErr(reflectVal)
	}

	mapKeys := reflectVal.MapKeys()
	newMap := make(
		map[string]interface{},
		reflectVal.Len()+1,
	)

	for _, key := range mapKeys {
		value := reflectVal.MapIndex(key)
		keyAny := key.Interface()
		var keyAsString string

		keyAsString, isString := keyAny.(string)
		if isString {
			newMap[keyAsString] = value.Interface()
			continue
		}

		keyAsString = fmt.Sprintf(
			constants.SprintValueFormat,
			keyAny,
		)
		newMap[keyAsString] = value.Interface()
	}

	return newMap, nil
}

// ToMapStringAny
//
//	expectation : map[key:interface{}]interface{} to map[string]interface{}
func (it mapConverter) ToMapStringAny(
	i interface{},
) (map[string]interface{}, error) {
	if Is.Null(i) {
		return map[string]interface{}{}, nil
	}

	return it.ToMapStringAnyRv(
		reflect.ValueOf(i),
	)
}
