package enuminf

type PatternsSplitter interface {
	// PatternSplit
	//
	//  based on pattern string it will split
	//  and try to convert each string split to enum
	PatternSplit(splitter, format string) []BasicEnumer
}
