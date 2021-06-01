package coreinterface

type StringJsoner interface {
	StringJson() (jsonString string, err error)
}
