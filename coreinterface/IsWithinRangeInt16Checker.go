package coreinterface

type IsWithinRangeInt16Checker interface {
	// IsWithinRange r.Min >= value && value <= r.Max
	//
	// Or, r.Start >= value && value <= r.End
	IsWithinRange(value int16) bool
}
