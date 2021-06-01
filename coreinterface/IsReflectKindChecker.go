package coreinterface

import "reflect"

type IsReflectKindChecker interface {
	IsReflectKind(checkingKind reflect.Kind) bool
}
