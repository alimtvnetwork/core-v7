package internalinterface

type CommonSliceDefiner interface {
	Length() int
	IsEmpty() bool
	Clear()
	Dispose()
	String() string
}
