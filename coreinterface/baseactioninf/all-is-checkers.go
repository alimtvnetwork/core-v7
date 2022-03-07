package baseactioninf

type IsExecutableChecker interface {
	IsExecutable() bool
}

type IsExecutableUsingMapChecker interface {
	IsExecutableUsingMap(conditionsNamesWithFlag map[string]bool) bool
}
