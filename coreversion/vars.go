package coreversion

var (
	New           = newCreator{}
	skipValuesMap = map[string]bool{
		"*": true,
		"":  true,
		" ": true,
	}
)
