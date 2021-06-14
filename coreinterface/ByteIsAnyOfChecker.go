package coreinterface

type ByteIsAnyOfChecker interface {
	IsAnyOf(value byte, checkingItems ...byte) bool
}
