package reflectinternal

import (
	"reflect"
	"runtime"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/refeflectcore/reflectmodel"
)

type getFunc struct{}

func (it getFunc) RunTime(i interface{}) *runtime.Func {
	if Is.Null(i) {
		return nil
	}

	rv := reflect.ValueOf(i)

	if rv.Kind() != reflect.Func {
		return nil
	}

	return runtime.FuncForPC(rv.Pointer())
}

// FullName
//
// Get the function name, passing non function may result panic
func (it getFunc) FullName(i interface{}) string {
	f := it.RunTime(i)

	if f == nil {
		return ""
	}

	return it.fixFinalFuncName(f.Name())
}

func (it getFunc) FullNameWithName(i interface{}) (fullName, name string) {
	fullName = it.FullName(i)

	if len(fullName) == 0 {
		return "", ""
	}

	_, _, funcNameOnly := it.All(fullName)

	return fullName, it.fixFinalFuncName(funcNameOnly)
}

func (it getFunc) Name(i interface{}) string {
	if Is.Null(i) {
		return ""
	}

	funcFullName := it.FullName(i)

	if len(funcFullName) == 0 {
		return ""
	}

	_, _, funcNameOnly := it.All(funcFullName)

	return it.fixFinalFuncName(funcNameOnly)
}

func (it getFunc) fixFinalFuncName(funcNameOnly string) string {
	if strings.HasSuffix(funcNameOnly, "-fm") {
		return funcNameOnly[:len(funcNameOnly)-3]
	}

	return funcNameOnly
}

func (it getFunc) All(fullFuncName string) (fullMethodName, packageName, methodName string) {
	if fullFuncName == "" {
		return "", "", ""
	}

	hasComplexName :=
		strings.HasPrefix(
			fullFuncName,
			gitlabDotCom,
		) ||
			strings.HasPrefix(
				fullFuncName,
				gitHubDotCom,
			) ||
			strings.LastIndexByte(
				fullFuncName,
				constants.ForwardChar,
			) > -1

	if hasComplexName {
		forwardSlashFound := strings.LastIndexByte(
			fullFuncName,
			constants.ForwardChar,
		)

		return it.All(fullFuncName[forwardSlashFound+1:])
	}

	splitsByDot := strings.Split(fullFuncName, constants.Dot)
	packageName, methodName = it.firstLastDefault(splitsByDot)

	return it.fixFinalFuncName(fullFuncName), packageName, it.fixFinalFuncName(methodName)
}

func (it getFunc) FuncDirectInvokeName(i interface{}) string {
	return it.FuncDirectInvokeNameUsingFullName(it.FullName(i))
}

func (it getFunc) FuncDirectInvokeNameUsingFullName(fullName string) string {
	if len(fullName) == 0 {
		return ""
	}

	forwardSlashFoundIndex := strings.LastIndexByte(
		fullName,
		constants.ForwardChar,
	)

	if forwardSlashFoundIndex > -1 {
		invokeName := fullName[forwardSlashFoundIndex+1:]

		splits := strings.Split(invokeName, ".")

		for i, split := range splits {
			// first pkg name
			if i == 0 {
				continue
			}

			splits[i] = it.PascalFuncName(split)
		}

		return strings.Join(splits, ".")
	}

	return fullName
}

func (it getFunc) firstLastDefault(slice []string) (first, last string) {
	length := len(slice)

	if length == 0 {
		return constants.EmptyString, constants.EmptyString
	}

	if length == 1 {
		return slice[0], constants.EmptyString
	}

	// length >= 2
	return slice[0], slice[length-1]
}

func (it getFunc) GetMethod(
	methodName string,
	i interface{},
) *reflect.Method {
	if len(methodName) == 0 || Is.Null(i) {
		return nil
	}

	valStruct := Looper.ReducePointerRv(
		reflect.ValueOf(i),
		defaultPointerReduction,
	)

	if valStruct.IsInvalid() {
		return nil
	}

	return it.GetMethodRv(
		methodName,
		&valStruct.FinalReflectVal,
	)
}

func (it getFunc) GetMethodRv(
	methodName string,
	rv *reflect.Value,
) *reflect.Method {
	if len(methodName) == 0 || Is.Null(rv) {
		return nil
	}

	structType := rv.Type()

	method, isFound := structType.MethodByName(methodName)

	if isFound {
		return &method
	}

	return nil
}

func (it getFunc) GetMethods(
	i interface{},
) []reflect.Method {
	if Is.Null(i) {
		return []reflect.Method{}
	}

	list := make([]reflect.Method, 0, 10)

	_ = Looper.MethodsFor(
		i,
		func(totalMethodsCount int, method *reflectmodel.MethodProcessor) (err error) {
			if method != nil {
				list = append(list, method.ReflectMethod)
			}

			return nil
		},
	)

	return list
}

func (it getFunc) GetMethodsRv(
	rv reflect.Value,
) []reflect.Method {
	list := make([]reflect.Method, 0, 4)

	_ = Looper.MethodsForRv(
		rv,
		func(totalMethodsCount int, method *reflectmodel.MethodProcessor) (err error) {
			if method != nil {
				list = append(list, method.ReflectMethod)
			}

			return nil
		},
	)

	return list
}

func (it getFunc) GetMethodsNames(
	i interface{},
) []string {
	if Is.Null(i) {
		return []string{}
	}

	list, _ := Looper.MethodNamesRv(
		reflect.ValueOf(i),
	)

	return list
}

func (it getFunc) GetMethodsMap(
	i interface{},
) map[string]*reflect.Method {
	if Is.Null(i) {
		return map[string]*reflect.Method{}
	}

	mapList, _ := Looper.MethodsMap(i)

	return mapList
}

func (it getFunc) GetMethodsMapRv(
	rv reflect.Value,
) map[string]*reflect.Method {
	mapList, _ := Looper.MethodsMapRv(rv)

	return mapList
}

func (it getFunc) GetMethodProcessorsMap(
	rv reflect.Value,
) map[string]*reflect.Method {
	mapList, _ := Looper.MethodsMapRv(rv)

	return mapList
}

func (it getFunc) PascalFuncName(
	name string,
) string {
	if len(name) == 0 {
		return ""
	}

	allRunes := []rune(name)
	firstChar := allRunes[0]
	firstCharStr := string(firstChar)
	firstCharUpper := strings.ToUpper(firstCharStr)

	if len(allRunes) == 1 {
		return firstCharUpper
	}

	return firstCharUpper + string(allRunes[1:])
}

func (it getFunc) GetPkgPath(i interface{}) interface{} {
	f := it.FullName(i)

	return it.GetPkgPathFullName(f)
}

func (it getFunc) GetPkgPathFullName(fullName string) string {
	if len(fullName) == 0 {
		return ""
	}

	forwardSlashFoundIndex := strings.LastIndexByte(
		fullName,
		constants.ForwardChar,
	)

	if forwardSlashFoundIndex <= -1 {
		return fullName
	}

	left := fullName[:forwardSlashFoundIndex]
	right := fullName[forwardSlashFoundIndex+1:]
	splits := strings.Split(right, ".")

	if len(splits) == 0 {
		return left
	}

	return left + constants.ForwardSlash + splits[0]
}
