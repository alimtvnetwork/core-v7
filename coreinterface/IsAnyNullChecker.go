package coreinterface

// IsAnyNullChecker
//
// Returns true if self is null or values is null
// Values have to be null to have true return.
// False: Any empty slice will return false.
type IsAnyNullChecker interface {
	// IsAnyNull
	//
	// Returns true if self is null or values is null
	// Values have to be null to have true return.
	// False: Any empty slice will return false.
	IsAnyNull() bool
}
