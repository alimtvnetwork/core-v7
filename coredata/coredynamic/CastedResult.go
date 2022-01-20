package coredynamic

import "reflect"

type CastedResult struct {
	Casted                         interface{}
	SourceReflectType              reflect.Type
	SourceKind                     reflect.Kind
	Error                          error
	IsNull, IsMatchingAcceptedType bool
	IsValid                        bool
	IsPointer                      bool // refers to how returned, ptr or non ptr
	IsSourcePointer                bool
}

func (it *CastedResult) IsInvalid() bool {
	return it == nil || !it.IsValid
}

func (it *CastedResult) IsNotNull() bool {
	return it != nil && !it.IsNull
}

func (it *CastedResult) IsNotPointer() bool {
	return it != nil && !it.IsPointer
}

func (it *CastedResult) IsNotMatchingAcceptedType() bool {
	return it != nil && !it.IsMatchingAcceptedType
}

func (it *CastedResult) IsSourceKind(kind reflect.Kind) bool {
	return it != nil && it.SourceKind == kind
}

func (it *CastedResult) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *CastedResult) HasAnyIssues() bool {
	return it.IsInvalid() || it.IsNull || !it.IsMatchingAcceptedType
}
