package args

import (
	"errors"
	"reflect"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type FuncMap map[string]FuncWrap

func (it FuncMap) IsEmpty() bool {
	return len(it) == 0
}

func (it FuncMap) Length() int {
	return len(it)
}

func (it FuncMap) Count() int {
	return len(it)
}

func (it FuncMap) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it FuncMap) Has(name string) bool {
	if it.IsEmpty() {
		return false
	}

	_, isFound := it[name]

	return isFound
}

func (it FuncMap) IsContains(name string) bool {
	return it.Has(name)
}

func (it FuncMap) Get(name string) *FuncWrap {
	if it.IsEmpty() {
		return nil
	}

	f, isFound := it[name]

	if isFound {
		return &f
	}

	return nil
}

func (it *FuncMap) Add(i interface{}) *FuncMap {
	if it == nil {
		*it = map[string]FuncWrap{}
	}

	v := NewFuncWrap.Single(i)

	if v.IsValid() {
		(*it)[v.Name] = *v
	}

	return it
}

func (it *FuncMap) Adds(iFunctions ...interface{}) *FuncMap {
	if it == nil {
		*it = map[string]FuncWrap{}
	}

	if len(iFunctions) == 0 {
		return it
	}

	for _, function := range iFunctions {
		it.Add(function)
	}

	return it
}

func (it *FuncMap) AddStructFunctions(iStructs ...interface{}) error {
	if it == nil {
		*it = map[string]FuncWrap{}
	}

	if len(iStructs) == 0 {
		return nil
	}

	for _, s := range iStructs {
		funcMap, err := NewFuncWrap.StructToMap(s)

		if err != nil {
			return err
		}

		for _, wrap := range funcMap {
			it.Add(wrap)
		}
	}

	return nil
}

func (it FuncMap) GetPascalCaseFuncName(name string) string {
	if len(it) == 0 {
		return ""
	}

	return reflectinternal.
		GetFunc.
		PascalFuncName(name)
}

func (it FuncMap) IsValidFuncOf(name string) bool {
	f := it.Get(name)

	if f == nil {
		return false
	}

	return f.HasValidFunc()
}

func (it FuncMap) IsInvalidFunc(name string) bool {
	f := it.Get(name)

	if f == nil {
		return true
	}

	return f.IsInvalid()
}

func (it FuncMap) PkgPath(name string) string {
	f := it.Get(name)

	if f == nil {
		return ""
	}

	return f.PkgPath()
}

func (it FuncMap) PkgNameOnly(name string) string {
	f := it.Get(name)

	if f == nil {
		return ""
	}

	return f.PkgNameOnly()
}

func (it FuncMap) FuncDirectInvokeName(name string) string {
	f := it.Get(name)

	if f == nil {
		return ""
	}

	return f.FuncDirectInvokeName()
}

// ArgsCount returns -1 on invalid
func (it FuncMap) ArgsCount(name string) int {
	f := it.Get(name)

	if f == nil {
		return 0
	}

	return f.ArgsCount()
}

// ArgsLength is an Alias for ArgsCount
func (it FuncMap) ArgsLength(name string) int {
	return it.ArgsCount(name)
}

// ReturnLength refers to the return arguments length
func (it FuncMap) ReturnLength(name string) int {
	f := it.Get(name)

	if f == nil {
		return 0
	}

	return f.ReturnLength()
}

func (it FuncMap) IsPublicMethod(name string) bool {
	f := it.Get(name)

	if f == nil {
		return false
	}

	return f.IsPublicMethod()
}

func (it FuncMap) IsPrivateMethod(name string) bool {
	f := it.Get(name)

	if f == nil {
		return false
	}

	return f.IsPrivateMethod()
}

func (it FuncMap) GetType(name string) reflect.Type {
	f := it.Get(name)

	if f == nil {
		return reflect.Type(nil)
	}

	return f.GetType()
}

func (it FuncMap) GetOutArgsTypes(name string) []reflect.Type {
	f := it.Get(name)

	if f == nil {
		return []reflect.Type{}
	}

	return f.GetOutArgsTypes()
}

func (it FuncMap) GetInArgsTypes(name string) []reflect.Type {
	f := it.Get(name)

	if f == nil {
		return []reflect.Type{}
	}

	return f.GetOutArgsTypes()
}

func (it FuncMap) GetInArgsTypesNames(name string) []string {
	f := it.Get(name)

	if f == nil {
		return []string{}
	}

	return f.GetInArgsTypesNames()
}

func (it FuncMap) VerifyInArgs(name string, args []interface{}) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.VerifyInArgs(args)
}

func (it FuncMap) VerifyOutArgs(name string, args []interface{}) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.VerifyOutArgs(args)
}

func (it FuncMap) InArgsVerifyRv(name string, args []reflect.Type) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.InArgsVerifyRv(args)
}

func (it FuncMap) OutArgsVerifyRv(name string, args []reflect.Type) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.OutArgsVerifyRv(args)
}

func (it FuncMap) VoidCallNoReturn(
	name string,
	args ...interface{},
) (processingErr error) {
	f := it.Get(name)

	if f == nil {
		return it.notFoundErr(name)
	}

	return f.VoidCallNoReturn(args...)
}

func (it FuncMap) MustBeValid(name string) {
	f := it.Get(name)

	if f == nil {
		panic(it.notFoundErr(name))
	}

	f.MustBeValid()
}

func (it FuncMap) ValidationError(name string) error {
	f := it.Get(name)

	if f == nil {
		return it.notFoundErr(name)
	}

	return f.ValidationError()
}

func (it FuncMap) InvokeMust(
	name string,
	args ...interface{},
) []interface{} {
	results, err := it.Invoke(name, args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it FuncMap) Invoke(
	name string,
	args ...interface{},
) (results []interface{}, processingErr error) {
	return it.InvokeSkip(codestack.Skip1, name, args...)
}

func (it FuncMap) InvokeSkip(
	skipStack int,
	name string,
	args ...interface{},
) (results []interface{}, processingErr error) {
	f := it.Get(name)

	if f == nil {
		return []interface{}{}, it.notFoundErr(name)
	}

	return f.InvokeSkip(skipStack+1, args)
}

func (it FuncMap) VoidCall(name string) ([]interface{}, error) {
	return it.Invoke(name)
}

func (it FuncMap) ValidateMethodArgs(name string, args []interface{}) error {
	f := it.Get(name)

	if f == nil {
		return it.notFoundErr(name)
	}

	return f.ValidateMethodArgs(args)
}

func (it FuncMap) GetFirstResponseOfInvoke(
	name string,
	args ...interface{},
) (firstResponse interface{}, err error) {
	result, err := it.InvokeResultOfIndex(name, 0, args...)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (it FuncMap) InvokeResultOfIndex(
	name string,
	index int,
	args ...interface{},
) (firstResponse interface{}, err error) {
	f := it.Get(name)

	if f == nil {
		return nil, it.notFoundErr(name)
	}

	return f.InvokeResultOfIndex(index, args...)
}

func (it FuncMap) InvokeError(
	name string,
	args ...interface{},
) (funcErr, processingErr error) {
	result, err := it.GetFirstResponseOfInvoke(name, args...)

	if err != nil {
		return nil, err
	}

	return result.(error), err
}

// InvokeFirstAndError
//
//	useful for method which looks like ReflectMethod() (soemthing, error)
func (it FuncMap) InvokeFirstAndError(
	name string,
	args ...interface{},
) (firstResponse interface{}, funcErr, processingErr error) {
	f := it.Get(name)

	if f == nil {
		return nil, nil, it.notFoundErr(name)
	}

	return f.InvokeFirstAndError(args...)
}

func (it FuncMap) InvalidError() error {
	if it.IsEmpty() {
		return errors.New("func-wrap map is empty")
	}

	return nil
}

func (it FuncMap) InvalidErrorByName(name string) error {
	if it.IsEmpty() {
		return errors.New("func-wrap map is empty")
	}

	f := it.Get(name)

	if f == nil {
		return it.notFoundErr(name)
	}

	return f.InvalidError()
}

func (it FuncMap) notFoundErr(name string) error {
	return errcore.NotFound.Error("func-wrap not found by the name", name)
}
