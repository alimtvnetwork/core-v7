package coreinterface

type StringRangesGetter interface {
	StringRangesPtr() *[]string
	StringRanges() []string
}
