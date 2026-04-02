package args

import (
	"errors"
	"reflect"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// FuncMap is a named map of function names to their FuncWrapAny representations.
// It provides batch function management, lookup, and invocation by name.
type FuncMap map[string]FuncWrapAny

// IsEmpty returns true if the map contains no entries.
func (it FuncMap) IsEmpty() bool {
	return len(it) == 0
}

// Length returns the number of entries in the map.
func (it FuncMap) Length() int {
	return len(it)
}

// Count is an alias for Length.
func (it FuncMap) Count() int {
	return len(it)
}

// HasAnyItem returns true if the map contains at least one entry.
func (it FuncMap) HasAnyItem() bool {
	return !it.IsEmpty()
}

// Has checks if a function with the given name exists in the map.
func (it FuncMap) Has(name string) bool {
	if it.IsEmpty() {
		return false
	}

	_, isFound := it[name]

	return isFound
}

// IsContains is an alias for Has.
func (it FuncMap) IsContains(name string) bool {
	return it.Has(name)
}

// Get returns a pointer to the FuncWrapAny for the given name, or nil if not found.
func (it FuncMap) Get(name string) *FuncWrapAny {
	if it.IsEmpty() {
		return nil
	}

	f, isFound := it[name]

	if isFound {
		return &f
	}

	return nil
}

// Add wraps a function value and adds it to the map by its detected name.
func (it *FuncMap) Add(i any) *FuncMap {
	if *it == nil {
		*it = map[string]FuncWrapAny{}
	}

	v := NewFuncWrap.Single(i)

	if v.IsValid() {
		(*it)[v.Name] = *v
	}

	return it
}

// Adds wraps multiple function values and adds them to the map.
func (it *FuncMap) Adds(iFunctions ...any) *FuncMap {
	if *it == nil {
		*it = map[string]FuncWrapAny{}
	}

	if len(iFunctions) == 0 {
		return it
	}

	for _, function := range iFunctions {
		it.Add(function)
	}

	return it
}

// AddStructFunctions extracts all public methods from structs and adds them to the map.
func (it *FuncMap) AddStructFunctions(iStructs ...any) error {
	if *it == nil {
		*it = map[string]FuncWrapAny{}
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

// GetPascalCaseFuncName returns the PascalCase version of the given function name.
func (it FuncMap) GetPascalCaseFuncName(name string) string {
	if len(it) == 0 {
		return ""
	}

	return reflectinternal.
		GetFunc.
		PascalFuncName(name)
}

// IsValidFuncOf checks if the named function is valid and callable.
func (it FuncMap) IsValidFuncOf(name string) bool {
	f := it.Get(name)

	if f == nil {
		return false
	}

	return f.HasValidFunc()
}

// IsInvalidFunc checks if the named function is invalid or not found.
func (it FuncMap) IsInvalidFunc(name string) bool {
	f := it.Get(name)

	if f == nil {
		return true
	}

	return f.IsInvalid()
}

// PkgPath returns the package path of the named function.
func (it FuncMap) PkgPath(name string) string {
	f := it.Get(name)

	if f == nil {
		return ""
	}

	return f.PkgPath()
}

// PkgNameOnly returns the package name (without path) of the named function.
func (it FuncMap) PkgNameOnly(name string) string {
	f := it.Get(name)

	if f == nil {
		return ""
	}

	return f.PkgNameOnly()
}

// FuncDirectInvokeName returns the direct invocation name of the named function.
func (it FuncMap) FuncDirectInvokeName(name string) string {
	f := it.Get(name)

	if f == nil {
		return ""
	}

	return f.FuncDirectInvokeName()
}

// ArgsCount returns the input argument count of the named function. Returns 0 if not found.
func (it FuncMap) ArgsCount(name string) int {
	f := it.Get(name)

	if f == nil {
		return 0
	}

	return f.ArgsCount()
}

// ArgsLength is an alias for ArgsCount.
func (it FuncMap) ArgsLength(name string) int {
	return it.ArgsCount(name)
}

// ReturnLength returns the return value count of the named function. Returns 0 if not found.
func (it FuncMap) ReturnLength(name string) int {
	f := it.Get(name)

	if f == nil {
		return 0
	}

	return f.ReturnLength()
}

// IsPublicMethod checks if the named function is an exported method.
func (it FuncMap) IsPublicMethod(name string) bool {
	f := it.Get(name)

	if f == nil {
		return false
	}

	return f.IsPublicMethod()
}

// IsPrivateMethod checks if the named function is an unexported method.
func (it FuncMap) IsPrivateMethod(name string) bool {
	f := it.Get(name)

	if f == nil {
		return false
	}

	return f.IsPrivateMethod()
}

// GetType returns the reflect.Type of the named function.
func (it FuncMap) GetType(name string) reflect.Type {
	f := it.Get(name)

	if f == nil {
		return reflect.Type(nil)
	}

	return f.GetType()
}

// GetOutArgsTypes returns the output argument types of the named function.
func (it FuncMap) GetOutArgsTypes(name string) []reflect.Type {
	f := it.Get(name)

	if f == nil {
		return []reflect.Type{}
	}

	return f.GetOutArgsTypes()
}

// GetInArgsTypes returns the input argument types of the named function.
func (it FuncMap) GetInArgsTypes(name string) []reflect.Type {
	f := it.Get(name)

	if f == nil {
		return []reflect.Type{}
	}

	return f.GetOutArgsTypes()
}

// GetInArgsTypesNames returns the string names of input argument types.
func (it FuncMap) GetInArgsTypesNames(name string) []string {
	f := it.Get(name)

	if f == nil {
		return []string{}
	}

	return f.GetInArgsTypesNames()
}

// VerifyInArgs verifies input argument types for the named function.
func (it FuncMap) VerifyInArgs(
	name string,
	args []any,
) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.VerifyInArgs(args)
}

// VerifyOutArgs verifies output argument types for the named function.
func (it FuncMap) VerifyOutArgs(
	name string,
	args []any,
) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.VerifyOutArgs(args)
}

// InArgsVerifyRv verifies input argument types using reflect.Type slices.
func (it FuncMap) InArgsVerifyRv(
	name string,
	args []reflect.Type,
) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.InArgsVerifyRv(args)
}

// OutArgsVerifyRv verifies output argument types using reflect.Type slices.
func (it FuncMap) OutArgsVerifyRv(
	name string,
	args []reflect.Type,
) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.OutArgsVerifyRv(args)
}

// VoidCallNoReturn invokes the named function ignoring return values.
func (it FuncMap) VoidCallNoReturn(
	name string,
	args ...any,
) (processingErr error) {
	f := it.Get(name)

	if f == nil {
		return it.notFoundErr(name)
	}

	return f.VoidCallNoReturn(args...)
}

// MustBeValid panics if the named function is not found or invalid.
func (it FuncMap) MustBeValid(name string) {
	f := it.Get(name)

	if f == nil {
		panic(it.notFoundErr(name))
	}

	f.MustBeValid()
}

// ValidationError returns an error if the named function is not found or invalid.
func (it FuncMap) ValidationError(name string) error {
	f := it.Get(name)

	if f == nil {
		return it.notFoundErr(name)
	}

	return f.ValidationError()
}

// InvokeMust invokes the named function, panicking on error.
func (it FuncMap) InvokeMust(
	name string,
	args ...any,
) []any {
	results, err := it.Invoke(name, args...)

	if err != nil {
		panic(err)
	}

	return results
}

// Invoke dynamically calls the named function with the given arguments.
func (it FuncMap) Invoke(
	name string,
	args ...any,
) (results []any, processingErr error) {
	return it.InvokeSkip(codestack.Skip1, name, args...)
}

// InvokeSkip invokes the named function with a custom stack skip.
func (it FuncMap) InvokeSkip(
	skipStack int,
	name string,
	args ...any,
) (results []any, processingErr error) {
	f := it.Get(name)

	if f == nil {
		return []any{}, it.notFoundErr(name)
	}

	return f.InvokeSkip(skipStack+1, args...)
}

// VoidCall invokes the named function with no arguments.
func (it FuncMap) VoidCall(name string) ([]any, error) {
	return it.Invoke(name)
}

// ValidateMethodArgs validates argument types for the named function.
func (it FuncMap) ValidateMethodArgs(
	name string,
	args []any,
) error {
	f := it.Get(name)

	if f == nil {
		return it.notFoundErr(name)
	}

	return f.ValidateMethodArgs(args)
}

// GetFirstResponseOfInvoke invokes the named function and returns the first result.
func (it FuncMap) GetFirstResponseOfInvoke(
	name string,
	args ...any,
) (firstResponse any, err error) {
	result, err := it.InvokeResultOfIndex(name, 0, args...)

	if err != nil {
		return nil, err
	}

	return result, err
}

// InvokeResultOfIndex invokes the named function and returns the result at the given index.
func (it FuncMap) InvokeResultOfIndex(
	name string,
	index int,
	args ...any,
) (firstResponse any, err error) {
	f := it.Get(name)

	if f == nil {
		return nil, it.notFoundErr(name)
	}

	return f.InvokeResultOfIndex(index, args...)
}

// InvokeError invokes the named function and returns the first result as an error.
func (it FuncMap) InvokeError(
	name string,
	args ...any,
) (funcErr, processingErr error) {
	result, err := it.GetFirstResponseOfInvoke(name, args...)

	if err != nil {
		return nil, err
	}

	return result.(error), err
}

// InvokeFirstAndError invokes the named function and separates the first result from the error.
func (it FuncMap) InvokeFirstAndError(
	name string,
	args ...any,
) (firstResponse any, funcErr, processingErr error) {
	f := it.Get(name)

	if f == nil {
		return nil, nil, it.notFoundErr(name)
	}

	return f.InvokeFirstAndError(args...)
}

// InvalidError returns an error if the map is empty.
func (it FuncMap) InvalidError() error {
	if it.IsEmpty() {
		return errors.New("func-wrap map is empty")
	}

	return nil
}

// InvalidErrorByName returns an error if the map is empty or the named function is invalid.
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
