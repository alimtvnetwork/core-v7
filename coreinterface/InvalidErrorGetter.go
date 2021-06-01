package coreinterface

type InvalidErrorGetter interface {
	// InvalidError get invalid message error
	InvalidError() error
}
