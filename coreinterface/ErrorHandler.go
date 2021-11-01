package coreinterface

type ErrorHandler interface {
	// HandleError
	//
	// Only call panic
	// if it has any error
	HandleError()
}
