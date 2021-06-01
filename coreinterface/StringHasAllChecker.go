package coreinterface

type StringHasAllChecker interface {
	HasAll(searchTerms ...string) bool
}
