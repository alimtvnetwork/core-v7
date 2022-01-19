package coreinterface

type ActionReturnsErrorFuncBinder interface {
	Exec() (err error)
}
