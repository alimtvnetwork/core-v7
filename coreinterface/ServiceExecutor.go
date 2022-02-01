package coreinterface

type ServiceExecutor interface {
	StandardExecutor
	ServiceNameGetter
	DefaultsInjector
	MustDefaultsInjector
	IsValidChecker
	IsInvalidChecker
	ValidationErrorGetter
}
