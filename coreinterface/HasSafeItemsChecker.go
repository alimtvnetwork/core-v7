package coreinterface

type HasSafeItemsChecker interface {
	// HasSafeItems
	//
	// returns true if has valid item or items and no error
	HasSafeItems() bool
}
