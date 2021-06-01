package coreinterface

type CsvLiner interface {
	CsvLines() *[]string
	CsvLinesOptions(isSkipQuoteOnlyOnExistence bool) *[]string
}
