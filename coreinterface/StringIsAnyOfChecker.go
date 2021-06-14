package coreinterface

type StringIsAnyOfChecker interface {
	IsAnyOf(value string, checkingItems ...string) bool
}
