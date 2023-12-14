package reflectmodel

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type MethodProcessor struct {
	Name             string
	Index            int
	ReflectMethod    reflect.Method
	inArgsTypesNames []string       `json:"-"`
	inArgsTypes      []reflect.Type `json:"-"`
	outArgsTypes     []reflect.Type `json:"-"`
}

func (it *MethodProcessor) HasValidFunc() bool {
	return it != nil
}

func (it *MethodProcessor) GetFuncName() string {
	return it.Name
}

func (it *MethodProcessor) IsInvalid() bool {
	return it == nil
}

func (it *MethodProcessor) Func() *reflect.Value {
	if it.IsInvalid() {
		return nil
	}

	return &it.ReflectMethod.Func
}

// ArgsCount is same as ArgsLength
//
// Reference:
//
//	https://stackoverflow.com/a/47626214
func (it *MethodProcessor) ArgsCount() int {
	return it.ReflectMethod.Type.NumIn()
}

// ReturnLength refers to the return arguments length
func (it *MethodProcessor) ReturnLength() int {
	if it.IsInvalid() {
		return -1
	}

	// https://stackoverflow.com/a/47626214

	return it.GetType().NumOut()
}

func (it *MethodProcessor) IsPublicMethod() bool {
	return it != nil && it.ReflectMethod.PkgPath == ""
}

func (it *MethodProcessor) IsPrivateMethod() bool {
	return it != nil && it.ReflectMethod.PkgPath != ""
}

// ArgsLength
//
// https://stackoverflow.com/a/47626214
// It is same as ArgsCount
func (it *MethodProcessor) ArgsLength() int {
	return it.ReflectMethod.Type.NumIn()
}

func (it *MethodProcessor) Invoke(args ...interface{}) (
	responses []interface{},
	err error,
) {
	firstErr := it.validationError()

	if firstErr != nil {
		return nil, firstErr
	}

	argsValidationErr := it.ValidateMethodArgs(args)

	if argsValidationErr != nil {
		return nil, argsValidationErr
	}

	rvs := util.ArgsToReflectValues(args)
	resultsRawValues := it.Func().Call(rvs)

	return util.ReflectValuesToInterfaces(resultsRawValues), nil
}

func (it *MethodProcessor) GetFirstResponseOfInvoke(
	args ...interface{},
) (firstResponse interface{}, err error) {
	result, err := it.InvokeResultOfIndex(0, args...)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (it *MethodProcessor) InvokeResultOfIndex(
	index int,
	args ...interface{},
) (firstResponse interface{}, err error) {
	results, err := it.Invoke(args...)

	if err != nil {
		return nil, err
	}

	return results[index], err
}

func (it *MethodProcessor) InvokeError(
	args ...interface{},
) (funcErr, processingErr error) {
	result, err := it.GetFirstResponseOfInvoke(args...)

	if err != nil {
		return nil, err
	}

	return result.(error), err
}

// InvokeFirstAndError
//
//	useful for method which looks like ReflectMethod() (soemthing, error)
func (it *MethodProcessor) InvokeFirstAndError(
	args ...interface{},
) (firstResponse interface{}, funcErr, processingErr error) {
	results, processingErr := it.Invoke(args...)

	if processingErr != nil {
		return nil, nil, processingErr
	}

	if len(results) <= 1 {
		return results,
			nil,
			errors.New(it.GetFuncName() + " doesn't return at least 2 return args")
	}

	first := results[0]
	second := results[1].(error)

	return first, second, processingErr
}

// IsNotEqual
//
// Based on predication.
//
// Warning: it can be wrong as well
func (it *MethodProcessor) IsNotEqual(
	another *MethodProcessor,
) bool {
	return !it.IsEqual(another)
}

// IsEqual
//
// Based on predication.
//
// Warning: it can be wrong as well
func (it *MethodProcessor) IsEqual(
	another *MethodProcessor,
) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it == another {
		return true
	}

	if it.IsInvalid() != another.IsInvalid() {
		return false
	}

	if it.Name != it.Name {
		return false
	}

	// can be skipped,
	// because name also refers to public or private
	if it.IsPublicMethod() != it.IsPublicMethod() {
		return false
	}

	if it.ArgsCount() != it.ArgsCount() {
		return false
	}

	if it.ReturnLength() != it.ReturnLength() {
		return false
	}

	isInArgsOkay, _ := it.InArgsVerifyRv(another.GetInArgsTypes())

	if !isInArgsOkay {
		return false
	}

	isOutArgsOkay, _ := it.OutArgsVerifyRv(another.GetOutArgsTypes())

	if !isOutArgsOkay {
		return false
	}

	// most probably true,
	// but can be false as well

	return true
}

func (it *MethodProcessor) GetType() reflect.Type {
	if it.IsInvalid() {
		return nil
	}

	return it.ReflectMethod.Type
}

func (it *MethodProcessor) GetOutArgsTypes() []reflect.Type {
	if it.IsInvalid() {
		return []reflect.Type{}
	}

	argsOutCount := it.ReturnLength()

	if argsOutCount == 0 {
		return []reflect.Type{}
	}

	if len(it.outArgsTypes) == argsOutCount {
		return it.outArgsTypes
	}

	// https://go.dev/play/p/dpIspUFfbu0
	mainType := it.GetType()
	slice := make([]reflect.Type, 0, argsOutCount)

	for i := 0; i < argsOutCount; i++ {
		slice = append(slice, mainType.Out(i))
	}

	it.outArgsTypes = slice

	return slice
}

func (it *MethodProcessor) GetInArgsTypes() []reflect.Type {
	if it.IsInvalid() {
		return []reflect.Type{}
	}

	argsCount := it.ArgsCount()

	if argsCount == 0 {
		return []reflect.Type{}
	}

	if len(it.inArgsTypes) == argsCount {
		return it.inArgsTypes
	}

	// https://go.dev/play/p/dpIspUFfbu0
	mainType := it.GetType()
	slice := make([]reflect.Type, 0, argsCount)

	for i := 0; i < argsCount; i++ {
		slice = append(slice, mainType.In(i))
	}

	it.inArgsTypes = slice

	return slice
}

func (it *MethodProcessor) GetInArgsTypesNames() []string {
	if it.IsInvalid() {
		return []string{}
	}

	argsCount := it.ArgsCount()

	if argsCount == 0 {
		return []string{}
	}

	if len(it.inArgsTypesNames) == argsCount {
		return it.inArgsTypesNames
	}

	// https://go.dev/play/p/dpIspUFfbu0
	mainType := it.GetType()
	slice := make([]string, 0, argsCount)

	for i := 0; i < argsCount; i++ {
		slice = append(slice, mainType.In(i).Name())
	}

	it.inArgsTypesNames = slice

	return slice
}

func (it *MethodProcessor) validationError() error {
	if it == nil {
		return errors.New("cannot execute on nil func-wrap")
	}

	if it.IsInvalid() {
		return fmt.Errorf(
			"func-wrap is invalid:\n"+
				"    given type: %T\n"+
				"    name: %s",
			it.Func(),
			it.Name,
		)
	}

	return nil
}

func (it *MethodProcessor) ValidateMethodArgs(args []interface{}) error {
	expectedCount := it.ArgsCount()
	given := len(args)

	if given != expectedCount {
		return errors.New(it.argsCountMismatchErrorMessage(expectedCount, given, args))
	}

	_, err := it.VerifyInArgs(args)

	return err
}

func (it *MethodProcessor) argsCountMismatchErrorMessage(
	expectedCount int,
	given int,
	args []interface{},
) string {
	expectedTypes := it.GetInArgsTypesNames()
	expectedToNames := strings.Join(expectedTypes, newLineSpaceIndent)
	actualTypes := util.InterfacesToTypesNamesWithValues(args)
	actualTypesName := strings.Join(actualTypes, newLineSpaceIndent)

	return fmt.Sprintf(
		"%s [Func] =>\n"+
			"  arguments count doesn't match for - count:\n"+
			"    expected : %d\n"+
			"    given    : %d\n"+
			"  expected types listed :\n"+
			"    - %s\n"+
			"  actual given types list :\n"+
			"    - %s",
		it.Name,
		expectedCount,
		given,
		expectedToNames,
		actualTypesName,
	)
}

func (it *MethodProcessor) VerifyInArgs(args []interface{}) (isOkay bool, err error) {
	toTypes := util.InterfacesToTypes(args)

	return it.InArgsVerifyRv(toTypes)
}

func (it *MethodProcessor) VerifyOutArgs(args []interface{}) (isOkay bool, err error) {
	toTypes := util.InterfacesToTypes(args)

	return it.OutArgsVerifyRv(toTypes)
}

func (it *MethodProcessor) InArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return util.
		VerifyReflectTypes(
			it.Name,
			it.GetInArgsTypes(),
			args,
		)
}

func (it *MethodProcessor) OutArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return util.
		VerifyReflectTypes(
			it.Name,
			it.GetOutArgsTypes(),
			args,
		)
}
