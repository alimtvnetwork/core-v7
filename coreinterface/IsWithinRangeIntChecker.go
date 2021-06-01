package coreinterface

type IsWithinRangeIntChecker interface {
	// IsWithinRange r.Min >= value && value <= r.Max
	//
	// Or, r.Start >= value && value <= r.End
	IsWithinRange(value int) bool
}
