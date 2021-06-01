package coreinterface

type IsWithinRangeInt32Checker interface {
	// IsWithinRange r.Min >= value && value <= r.Max
	//
	// Or, r.Start >= value && value <= r.End
	IsWithinRange(value int32) bool
}
