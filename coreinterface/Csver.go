package coreinterface

type Csver interface {
	Csv() string
	CsvOptions(isSkipQuoteOnlyOnExistence bool) string
}
