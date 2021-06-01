package coreinterface

type IsDynamicNullChecker interface {
	// IsDynamicNull may check using reflection that data is nil.
	IsDynamicNull(dynamic interface{}) bool
}
