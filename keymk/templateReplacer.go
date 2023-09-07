package keymk

import (
	"strconv"
	"strings"
)

type templateReplacer struct {
	key *Key
}

func (it *templateReplacer) RequestIntRange(
	isCurly bool,
	tempReplace TempReplace,
) []string {
	return it.IntRange(
		isCurly,
		tempReplace.KeyName,
		tempReplace.Range.Start,
		tempReplace.Range.End)
}

func (it *templateReplacer) IntRange(
	isCurly bool,
	keyName string,
	startIncluding, endIncluding int,
) []string {
	keyOuts := make(
		[]string,
		0,
		endIncluding-startIncluding+1)

	keyName = curlyWrapIf(isCurly, keyName)
	templateFormat := it.key.KeyCompiled() // format may hold {key-name}

	for i := startIncluding; i <= endIncluding; i++ {
		numString := strconv.Itoa(i)

		keyOuts = append(
			keyOuts,
			strings.ReplaceAll(
				templateFormat,
				keyName,
				numString))
	}

	return keyOuts
}

func (it *templateReplacer) CompileUsingReplacerMap(
	isCurly bool,
	replacerMap map[string]string, // key ==> find, value ==> replace
) string {
	templateFormat := it.key.KeyCompiled() // format may hold {key-name}

	if templateFormat == "" || len(replacerMap) == 0 {
		return templateFormat
	}

	for finder, replacer := range replacerMap {
		finderCurly := curlyWrapIf(isCurly, finder)

		templateFormat = strings.ReplaceAll(
			templateFormat,
			finderCurly,
			replacer)
	}

	return templateFormat
}
