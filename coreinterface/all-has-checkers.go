package coreinterface

type HasAnyIssueChecker interface {
	HasAnyIssues() bool
}

type HasAnyItemChecker interface {
	HasAnyItem() bool
}

type HasErrorChecker interface {
	HasError() bool
}

type HasIndexChecker interface {
	HasIndex(index int) bool
}

type HasIssueChecker interface {
	HasAnyItemChecker
	HasIssues() bool
	IsEmptyOrIssues() bool
	// HasValidItems Has items and there is no issues
	HasValidItems() bool
}

type HasIssuesOrEmptyChecker interface {
	HasIssuesOrEmpty() bool
}

type HasItemAtChecker interface {
	HasItemAt(index int) bool
}

type HasSafeItemsChecker interface {
	// HasSafeItems
	//
	// returns true if has valid item or items and no error
	HasSafeItems() bool
}

type HasFlagByNameChecker interface {
	HasFlagByName(flagName string) bool
}

type HasKeyChecker interface {
	HasKey(key string) bool
}

type HasAllKeysChecker interface {
	HasAllKeys(keys ...string) bool
}

type HasAnyKeysChecker interface {
	HasAnyKeys(keys ...string) bool
}
