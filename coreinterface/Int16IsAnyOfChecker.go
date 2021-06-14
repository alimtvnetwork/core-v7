package coreinterface

type Int16IsAnyOfChecker interface {
	IsAnyOf(value int16, checkingItems ...int16) bool
}
