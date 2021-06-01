package coreinterface

type ListStringsGetter interface {
	ListStringsPtr() *[]string
	ListStrings() []string
}
