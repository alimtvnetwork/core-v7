package keymk

import "strings"

type fixedLegend struct{}

func (it fixedLegend) FormatKeyMap(
	rootName,
	packageName,
	groupName,
	stateName,
	userName,
	item string,
) (format string, replacerMap map[string]string) {
	return LegendChainSample, map[string]string{
		"{root}":    rootName,
		"{package}": packageName,
		"{group}":   groupName,
		"{state}":   stateName,
		"{user}":    userName,
		"{item}":    item,
	}
}

func (it fixedLegend) Compile(
	isKeepFormatOnEmpty bool,
	rootName,
	packageName,
	groupName,
	stateName,
	userName,
	item string,
) (compiled string) {
	format, compilingMap := it.FormatKeyMap(
		rootName,
		packageName,
		groupName,
		stateName,
		userName,
		item)

	for searcher, replacer := range compilingMap {
		if isKeepFormatOnEmpty && replacer == "" {
			continue
		}

		format = strings.ReplaceAll(
			format,
			searcher,
			replacer)
	}

	return format
}

func (it fixedLegend) CompileKeepFormatOnEmpty(
	rootName,
	packageName,
	groupName,
	stateName,
	userName,
	item string,
) (compiled string) {
	return it.Compile(
		true,
		rootName,
		packageName,
		groupName,
		stateName,
		userName,
		item)
}
