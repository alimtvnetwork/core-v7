package coredynamic

import "reflect"

type CastedResult struct {
	Casted                         interface{}
	SourceReflectType              reflect.Type
	SourceKind                     reflect.Kind
	Error                          error
	IsNull, IsMatchingAcceptedType bool
	IsPointer                      bool // refers to how returned, ptr or non ptr
	IsSourcePointer                bool
}
