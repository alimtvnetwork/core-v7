package coreinterface

import "reflect"

type ReflectKindGetter interface {
	ReflectKind() reflect.Kind
}
