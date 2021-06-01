package coreinterface

type StringHasCombineChecker interface {
	StringHasChecker
	StringHasAllChecker
	StringHasAnyChecker
	StringHasAnyItemChecker
}
