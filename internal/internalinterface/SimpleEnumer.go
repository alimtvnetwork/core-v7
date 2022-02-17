package internalinterface

type SimpleEnumer interface {
	ValueByte() byte
	ToNamer
	String() string
	IsValidChecker
	IsInvalidChecker
}
