package coreinterface

type ValidationErrorGetter interface {
	ValidationError() error
}
