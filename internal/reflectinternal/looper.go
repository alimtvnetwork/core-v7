package reflectinternal

import (
	"errors"
	"reflect"
	"strings"
	"unsafe"

	"gitlab.com/auk-go/core/refeflectcore/reflectmodel"
)

type looper struct{}

func (it *looper) FieldsFor(
	anyItem interface{},
	processor func(currentField *reflectmodel.FieldProcessor) (err error),
) error {
	rv := reflect.ValueOf(anyItem)

	return it.FieldsForRv(rv, processor)
}

func (it *looper) FieldsForRv(
	rv reflect.Value,
	processor func(currentField *reflectmodel.FieldProcessor) (err error),
) error {
	reduceRv := it.ReducePointerRvDefault(rv)

	if reduceRv.IsInvalid() || reduceRv.HasError() {
		return reduceRv.Error
	}

	// valid
	structType := reduceRv.FinalReflectVal.Type()
	fieldsLength := structType.NumField()

	for i := 0; i < fieldsLength; i++ {
		field := structType.Field(i)
		input := reflectmodel.FieldProcessor{
			Name:      field.Name,
			Index:     i,
			Field:     field,
			FieldType: field.Type,
		}

		e := processor(&input)

		if e != nil {
			return e
		}
	}

	return nil
}

func (it *looper) FieldNames(
	anyStruct interface{},
) (fieldNames []string, err error) {
	rv := reflect.ValueOf(anyStruct)

	return it.FieldNamesRv(rv)
}

func (it *looper) FieldNamesRv(
	rv reflect.Value,
) (fieldNames []string, err error) {
	reduceRv := it.ReducePointerRvDefault(rv)

	if reduceRv.IsInvalid() || reduceRv.HasError() {
		return []string{}, reduceRv.Error
	}

	// valid
	structType := reduceRv.FinalReflectVal.Type()
	fieldsLength := structType.NumField()

	for i := 0; i < fieldsLength; i++ {
		field := structType.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames, nil
}

func (it *looper) FieldsMap(
	anyItem interface{},
) (resultsMap map[string]*reflect.StructField, err error) {
	rv := reflect.ValueOf(anyItem)

	return it.FieldsMapRv(rv)
}

func (it *looper) FieldsMapRv(
	rv reflect.Value,
) (resultsMap map[string]*reflect.StructField, err error) {
	reduceRv := it.ReducePointerRvDefault(rv)

	if reduceRv.IsInvalid() || reduceRv.HasError() {
		return map[string]*reflect.StructField{}, reduceRv.Error
	}

	// valid
	structType := reduceRv.FinalReflectVal.Type()
	fieldsLength := structType.NumField()
	resultsMap = make(
		map[string]*reflect.StructField,
		fieldsLength,
	)

	for i := 0; i < fieldsLength; i++ {
		field := structType.Field(i)
		resultsMap[field.Name] = &field
	}

	return resultsMap, nil
}

func (it *looper) MethodsMap(
	anyItem interface{},
) (resultsMap map[string]*reflect.Method, err error) {
	rv := reflect.ValueOf(anyItem)

	return it.MethodsMapRv(rv)
}

// ReducePointer
//
//	level -1 means all levels (****...) to Non pointer
func (it *looper) ReducePointer(
	anyItem interface{},
	level int,
) *reflectmodel.ReflectValueKind {
	return it.ReducePointerRv(reflect.ValueOf(anyItem), level)
}

func (it *looper) ReducePointerDefault(
	anyItem interface{},
) *reflectmodel.ReflectValueKind {
	return it.ReducePointerRv(reflect.ValueOf(anyItem), defaultPointerReduction)
}

func (it *looper) ReducePointerRvDefault(
	reflectVal reflect.Value,
) *reflectmodel.ReflectValueKind {
	return it.ReducePointerRv(reflectVal, defaultPointerReduction)
}

// ReducePointerRv
//
//	level -1 means all levels (****...) to Non pointer
func (it *looper) ReducePointerRv(
	reflectVal reflect.Value,
	level int,
) *reflectmodel.ReflectValueKind {
	structValueKind := reflectVal.Kind()
	hasLevel := level > invalid
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

	if !structValue.IsValid() || structValueKind == reflect.Ptr || structValueKind == reflect.Interface {
		return reflectmodel.InvalidReflectValueKindModel(
			"invalid ref value or could not reach in level limit",
		)
	}

	// valid
	return &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: structValue,
		Kind:            structValue.Kind(),
		Error:           nil,
	}
}

func (it *looper) MethodsFor(
	anyItem interface{},
	processor func(
		totalMethodsCount int,
		method *reflectmodel.MethodProcessor,
	) (err error),
) error {
	// valid
	// https://stackoverflow.com/q/598defaultPointerReduction1642
	// https://prnt.sc/kmkTmVmO2cPH
	// Pointer connected method and non pointer connect methods will be different
	rv := reflect.ValueOf(anyItem) // can be a pointer or non pointer

	return it.MethodsForRv(rv, processor)
}

func (it *looper) MethodNamesRv(
	rv reflect.Value,
) (methodNames []string, err error) {
	reduceRv := it.ReducePointerRvDefault(rv)

	if reduceRv.IsInvalid() || reduceRv.HasError() {
		return methodNames, reduceRv.Error
	}

	// valid
	structType := rv.Type()
	fieldsLength := structType.NumField()

	for i := 0; i < fieldsLength; i++ {
		field := structType.Field(i)
		methodNames = append(methodNames, field.Name)
	}

	return methodNames, nil
}

func (it *looper) MethodsForRv(
	rv reflect.Value,
	processor func(
		totalMethodsCount int,
		method *reflectmodel.MethodProcessor,
	) (err error),
) error {
	// valid
	// https://stackoverflow.com/q/598defaultPointerReduction1642
	// https://prnt.sc/kmkTmVmO2cPH
	// Pointer connected method and non pointer connect methods will be different
	ptrRv, conErr := it.ToPointerReflectValueRv(rv)

	if conErr != nil {
		return conErr
	}

	err := it.loopBaseMethods(ptrRv, processor)

	if err != nil {
		return err
	}

	// non pointer
	reducer := it.ReducePointerRvDefault(rv)

	return it.loopBaseMethods(reducer.FinalReflectVal, processor)
}

func (it *looper) Slice(
	i interface{},
	processor func(
		total int,
		index int,
		item interface{},
	) (err error),
) error {
	if Is.Null(i) {
		return nil
	}

	toRv := reflect.ValueOf(i)

	return it.SliceForRv(toRv, processor)
}

func (it *looper) SliceForRv(
	rv reflect.Value,
	processor func(
		total int,
		index int,
		item interface{},
	) (err error),
) error {
	valueRvWrap := it.ReducePointerRv(rv, defaultPointerReduction)

	if valueRvWrap.HasError() {
		return valueRvWrap.Error
	}

	valueRv := valueRvWrap.FinalReflectVal

	k := valueRv.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return errors.New("given item is not a slice nor an array")
	}

	length := valueRv.Len()

	if length == 0 {
		return nil
	}

	var errSlice []string

	for i := 0; i < length; i++ {
		err := processor(length, i, valueRv.Index(i))

		if err != nil {
			errSlice = append(errSlice, err.Error())
		}
	}

	if len(errSlice) == 0 {
		return nil
	}

	toMsg := strings.Join(errSlice, "\n")

	return errors.New(toMsg)
}

func (it *looper) MapForRv(
	rv reflect.Value,
	processor func(
		total int,
		index int,
		key,
		value interface{},
	) (err error),
) error {
	valueRvWrap := it.ReducePointerRv(rv, defaultPointerReduction)

	if valueRvWrap.HasError() {
		return valueRvWrap.Error
	}

	valueRv := valueRvWrap.FinalReflectVal

	k := valueRv.Kind()
	isMap := k == reflect.Map

	if !isMap {
		return errors.New("given item is not a map")
	}

	mapKeys := valueRv.MapKeys()
	length := len(mapKeys)

	if length == 0 {
		return nil
	}

	var errSlice []string

	for i, key := range mapKeys {
		value := valueRv.MapIndex(key)
		err := processor(length, i, key, value)

		if err != nil {
			errSlice = append(errSlice, err.Error())
		}
	}

	if len(errSlice) == 0 {
		return nil
	}

	toMsg := strings.Join(errSlice, "\n")

	return errors.New(toMsg)
}

func (it *looper) MethodsMapRv(
	rv reflect.Value,
) (map[string]*reflect.Method, error) {
	// valid
	// https://stackoverflow.com/q/598defaultPointerReduction1642
	// https://prnt.sc/kmkTmVmO2cPH
	// Pointer connected method and non pointer connect methods will be different
	ptrRv, conErr := it.ToPointerReflectValueRv(rv)

	if conErr != nil {
		return map[string]*reflect.Method{}, conErr
	}

	resultsMap := it.baseMethodsMap(ptrRv)

	// non pointer
	reducer := it.ReducePointerRvDefault(rv)

	resultsMapNext := it.baseMethodsMap(
		reducer.FinalReflectVal,
	)

	for s, method := range resultsMapNext {
		resultsMap[s] = method
	}

	return resultsMap, nil
}

// ToPointerReflectValue
//
// anyItem must be a struct or pointer to struct
func (it *looper) ToPointerReflectValue(
	anyItem interface{},
) (reflect.Value, error) {
	// valid
	// https://stackoverflow.com/q/598defaultPointerReduction1642
	// https://prnt.sc/kmkTmVmO2cPH
	// Pointer connected method and non pointer connect methods will be different
	rv := reflect.ValueOf(anyItem) // can be a pointer or non pointer

	return it.ToPointerReflectValueRv(rv)
}

// ToPointerReflectValueRv
//
// Rv must be a struct or pointer to struct
func (it *looper) ToPointerReflectValueRv(
	rv reflect.Value,
) (reflect.Value, error) {
	// valid
	// https://stackoverflow.com/q/598defaultPointerReduction1642
	// https://prnt.sc/kmkTmVmO2cPH
	// Pointer connected method and non pointer connect methods will be different
	k := rv.Kind()
	switch k {
	case reflect.Ptr:
		return rv, nil
	case reflect.Struct:
		toInterface := rv.Interface()
		toPointer := &toInterface
		unsafePtr := unsafe.Pointer(&toPointer)

		return reflect.NewAt(rv.Type(), unsafePtr), nil
	}

	return reflect.Value{}, errors.New("pointer and Struct is only allowed - given type - " + k.String())
}

// loopBaseMethods
//
// Pointer and non pointer methods are attached differently.
// Call this twice
func (it *looper) loopBaseMethods(
	rv reflect.Value, // can be a pointer or non pointer
	processor func(
		totalMethodsCount int,
		method *reflectmodel.MethodProcessor,
	) (err error),
) error {
	// valid
	// https://stackoverflow.com/q/598defaultPointerReduction1642
	// https://prnt.sc/kmkTmVmO2cPH
	// Pointer connected method and non pointer connect methods will be different
	structType := rv.Type()
	methodsCount := rv.NumMethod()

	for i := 0; i < methodsCount; i++ {
		method := structType.Method(i)
		input := reflectmodel.MethodProcessor{
			Name:          method.Name,
			Index:         i,
			ReflectMethod: method,
		}

		e := processor(methodsCount, &input)

		if e != nil {
			return e
		}
	}

	return nil
}

// loopBaseMethods
//
// Pointer and non pointer methods are attached differently.
// Call this twice
func (it *looper) baseMethodsMap(
	rv reflect.Value, // can be a pointer or non pointer
) map[string]*reflect.Method {
	// valid
	// https://stackoverflow.com/q/598defaultPointerReduction1642
	// https://prnt.sc/kmkTmVmO2cPH
	// Pointer connected method and non pointer connect methods will be different
	structType := rv.Type()
	methodsCount := rv.NumMethod()
	methodsMap := make(
		map[string]*reflect.Method,
		methodsCount,
	)

	for i := 0; i < methodsCount; i++ {
		method := structType.Method(i)
		methodsMap[method.Name] = &method
	}

	return methodsMap
}
