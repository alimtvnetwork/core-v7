package reflectmodel

import (
	"errors"
	"reflect"
	"unsafe"
)

type ReflectValueKind struct {
	IsValid         bool
	FinalReflectVal reflect.Value
	Kind            reflect.Kind
	Error           error
}

func InvalidReflectValueKindModel(err string) *ReflectValueKind {
	return &ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
		Kind:            0,
		Error:           errors.New(err),
	}
}

func (it *ReflectValueKind) IsInvalid() bool {
	return it == nil || !it.IsValid || it.HasError()
}

func (it *ReflectValueKind) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *ReflectValueKind) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *ReflectValueKind) ActualInstance() interface{} {
	if it == nil {
		return nil
	}

	return it.FinalReflectVal.Interface()
}

func (it *ReflectValueKind) PkgPath() string {
	if it == nil || !it.IsValid {
		return ""
	}

	return it.FinalReflectVal.Type().PkgPath()
}

func (it *ReflectValueKind) PointerRv() *reflect.Value {
	if it == nil {
		return nil
	}

	if !it.IsValid {
		return &it.FinalReflectVal
	}

	rv := it.FinalReflectVal

	toInterface := rv.Interface()
	toPointer := &toInterface
	unsafePtr := unsafe.Pointer(&toPointer)

	newRv := reflect.NewAt(rv.Type(), unsafePtr)

	return &newRv
}

func (it *ReflectValueKind) PointerInterface() interface{} {
	rv := it.PointerRv()

	if rv == nil {
		return nil
	}

	return rv.Interface()
}
