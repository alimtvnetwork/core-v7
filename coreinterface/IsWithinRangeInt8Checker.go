package coreinterface

type IsWithinRangeInt8Checker interface {
	// IsWithinRange r.Min >= value && value <= r.Max
	//
	// Or, r.Start >= value && value <= r.End
	IsWithinRange(value int8) bool
}
