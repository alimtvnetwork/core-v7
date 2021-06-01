package coredynamic

import "reflect"

type (
	SimpleInOutConverter   func(in interface{}, typeMust reflect.Type) *SimpleResult
	SimpleRequestConverter func(request SimpleRequest) *SimpleResult
)
