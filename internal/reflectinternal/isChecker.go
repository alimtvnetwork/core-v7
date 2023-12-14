package reflectinternal

import "reflect"

type isChecker struct{}

func (it isChecker) Conclusive(left, right interface{}) (isEqual, isConclusive bool) {
	if left == right {
		return true, true
	}

	if left == nil && right == nil {
		return true, true
	}

	if left == nil || right == nil {
		return false, true
	}

	leftRv := reflect.ValueOf(left)
	rightRv := reflect.ValueOf(right)
	isLeftNull := it.NullRv(leftRv)
	isRightNull := it.NullRv(rightRv)
	isBothEqual := isLeftNull == isRightNull

	if isLeftNull && isBothEqual {
		// both null
		return true, true
	} else if !isBothEqual && isLeftNull || isRightNull {
		// any null but the other is not
		return false, true
	}

	if leftRv.Type() != rightRv.Type() {
		return false, true
	}

	return false, false
}

func (it isChecker) AnyEqual(left, right interface{}) bool {
	isEqual, isConclusive := it.Conclusive(left, right)

	if isConclusive {
		return isEqual
	}

	return reflect.DeepEqual(left, right)
}

func (it isChecker) Func(item interface{}) bool {
	if item == nil {
		return true
	}

	typeOf := reflect.TypeOf(item)

	return it.FuncTypeOf(typeOf)
}

func (it isChecker) SliceOrArray(item interface{}) bool {
	if item == nil {
		return true
	}

	typeOf := reflect.TypeOf(item)

	return it.FuncTypeOf(typeOf)
}

func (it isChecker) NotFunc(item interface{}) bool {
	if item == nil {
		return true
	}

	return !it.Func(item)
}

func (it isChecker) FuncTypeOf(typeOf reflect.Type) bool {
	kind := typeOf.Kind()

	switch kind {
	case reflect.Func:
		return true
	}

	return false
}

func (it isChecker) SliceOrArrayOf(typeOf reflect.Type) bool {
	kind := typeOf.Kind()

	switch kind {
	case reflect.Slice, reflect.Array:
		return true
	}

	return false
}

func (it isChecker) NotNull(item interface{}) bool {
	return !it.Null(item)
}

func (it isChecker) Defined(item interface{}) bool {
	return !it.Null(item)
}

func (it isChecker) Null(item interface{}) bool {
	if item == nil {
		return true
	}

	rv := reflect.ValueOf(item)

	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

func (it isChecker) NullRv(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

// Number
//
// function returns true if the kind passed to it is one of the
// primitive types (reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
//
//	reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
//	reflect.Float32, reflect.Float64)
func (it isChecker) Number(kind reflect.Kind) bool {
	switch kind {
	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// Primitive
//
// function returns true if the kind passed to it is one of the
// primitive types (boolean, int, uint, float, string)
func (it isChecker) Primitive(kind reflect.Kind) bool {
	switch kind {
	case
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.String:
		return true
	default:
		return false
	}
}

// Zero
//
//	returns true if the current value is null
//	or reflect value is zero
//
// Reference:
//   - Stackoverflow Example : https://stackoverflow.com/a/23555352
func (it isChecker) Zero(anyItem interface{}) bool {
	if it.Null(anyItem) {
		return true
	}

	return it.ZeroRv(reflect.ValueOf(anyItem))
}

// ZeroRv
//
//	returns true if the current value is null
//	or reflect value is zero
//
// Reference:
//   - Stackoverflow Example : https://stackoverflow.com/a/23555352
func (it isChecker) ZeroRv(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice, reflect.Ptr:
		return rv.IsNil()
	case reflect.Array:
		isAllZero := true
		for i := 0; i < rv.Len(); i++ {
			isAllZero = isAllZero && it.ZeroRv(rv.Index(i))
		}

		return isAllZero
	case reflect.Struct:
		isAllZero := true
		for i := 0; i < rv.NumField(); i++ {
			isAllZero = isAllZero && it.ZeroRv(rv.Field(i))
		}

		return isAllZero
	}

	// Compare other types directly:
	z := reflect.Zero(rv.Type())

	return rv.Interface() == z.Interface()
}

func (it isChecker) Struct(structObj interface{}) bool {
	structRv := reflect.ValueOf(structObj)
	reducePtr := Looper.ReducePointerRvDefault(structRv)

	if reducePtr.IsValid {
		return reducePtr.Kind == reflect.Struct
	}

	return false
}

func (it isChecker) StructRv(structRv reflect.Value) bool {
	reducePtr := Looper.ReducePointerRvDefault(structRv)

	if reducePtr.IsValid {
		return reducePtr.Kind == reflect.Struct
	}

	return false
}

func (it isChecker) Interface(i interface{}) bool {
	iRv := reflect.ValueOf(i)
	reducePtr := Looper.ReducePointerRvDefault(iRv)

	if reducePtr.IsValid {
		return reducePtr.Kind == reflect.Interface
	}

	return false
}

func (it isChecker) InterfaceRv(iRv reflect.Value) bool {
	reducePtr := Looper.ReducePointerRvDefault(iRv)

	if reducePtr.IsValid {
		return reducePtr.Kind == reflect.Interface
	}

	return false
}
