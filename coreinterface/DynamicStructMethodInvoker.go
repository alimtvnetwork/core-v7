package coreinterface

import "gitlab.com/evatix-go/core/coredata/coredynamic"

type DynamicStructMethodInvoker interface {
	DynamicMethodInvoke(structInput interface{}, args ...interface{}) *coredynamic.SimpleResult
}
