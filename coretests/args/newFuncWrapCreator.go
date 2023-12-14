package args

import (
	"reflect"

	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/iserror"
)

type newFuncWrapCreator struct{}

func (it newFuncWrapCreator) Default(anyFunc interface{}) *FuncWrap {
	if reflectinternal.Is.Null(anyFunc) {
		return &FuncWrap{
			Func:      anyFunc,
			isInvalid: true,
		}
	}

	switch v := anyFunc.(type) {
	case *FuncWrap:
		return v
	case FuncWrapGetter:
		return v.FuncWrap()
	}

	typeOf := reflect.TypeOf(anyFunc)
	kind := typeOf.Kind()

	if kind != reflect.Func {
		// invalid

		return &FuncWrap{
			Func:      anyFunc,
			isInvalid: true,
			rvType:    typeOf,
		}
	}

	// valid
	fullName, nameOnly := reflectinternal.
		GetFunc.
		FullNameWithName(anyFunc)

	return &FuncWrap{
		Name:      nameOnly,
		FullName:  fullName,
		Func:      anyFunc,
		isInvalid: false,
		rvType:    typeOf,
		rv:        reflect.ValueOf(anyFunc),
	}
}

func (it newFuncWrapCreator) Single(
	anyFunc interface{},
) *FuncWrap {
	return it.Default(anyFunc)
}

func (it newFuncWrapCreator) Invalid() *FuncWrap {
	return &FuncWrap{
		isInvalid: true,
	}
}

func (it newFuncWrapCreator) Map(
	anyFunctions ...interface{},
) FuncMap {
	if len(anyFunctions) == 0 {
		return map[string]FuncWrap{}
	}

	newMap := make(
		map[string]FuncWrap,
		len(anyFunctions),
	)

	for _, function := range anyFunctions {
		v := it.Default(function)

		if v.IsValid() {
			newMap[v.GetFuncName()] = *v
		}
	}

	return newMap
}

func (it newFuncWrapCreator) Many(
	anyFunctions ...interface{},
) []*FuncWrap {
	if len(anyFunctions) == 0 {
		return []*FuncWrap{}
	}

	slice := make(
		[]*FuncWrap,
		len(anyFunctions),
	)

	for i, function := range anyFunctions {
		v := it.Default(function)

		slice[i] = v
	}

	return slice
}

func (it newFuncWrapCreator) MethodToFunc(
	m *reflect.Method,
) (*FuncWrap, error) {
	if m == nil {
		return it.Invalid(), errcore.CannotBeNilType.ErrorNoRefs("m * method cannot be nil")
	}

	name := m.Name
	fullName := m.PkgPath + name

	return &FuncWrap{
		Name:      name,
		FullName:  fullName,
		Func:      m.Func.Interface(),
		isInvalid: false,
		rvType:    m.Func.Type(),
		rv:        m.Func,
	}, nil
}

func (it newFuncWrapCreator) StructToMap(
	i interface{},
) (FuncMap, error) {
	methods, err := reflectinternal.Looper.MethodsMap(i)

	if iserror.Defined(err) {
		return Empty.FuncMap(), err
	}

	newMap := make(
		map[string]FuncWrap,
		len(methods),
	)

	var rawErr errcore.RawErrCollection

	for index, method := range methods {
		v, nErr := it.MethodToFunc(method)

		rawErr.Add(nErr)

		if v.IsValid() {
			newMap[index] = *v
		}
	}

	return newMap, rawErr.CompiledErrorWithStackTraces()
}
