package coredynamic

import "reflect"

type (
	SimpleInOutConverter                   func(in interface{}, typeMust reflect.Type) *SimpleResult
	SimpleRequestConverter                 func(request SimpleRequest) *SimpleResult
	TypeToTypeConverterFunc                func(in interface{}) (output interface{}, err error)
	TypeToTypeConverterMustFunc            func(in interface{}) (output interface{})
	TypeToTypeConverterWithOptionsFunc     func(in interface{}, options interface{}) (output interface{}, err error)
	TypeToTypeConverterWithOptionsMustFunc func(in interface{}, options interface{}) (output interface{})
)
