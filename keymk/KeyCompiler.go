package keymk

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/defaultcapacity"
)

func (it *Key) CompileKeys(
	keys ...*Key,
) string {
	if len(keys) == 0 {
		return it.Compile()
	}

	newSlice := make(
		[]any,
		0,
		it.Length()+
			defaultcapacity.PredictiveDefaultSmall(len(keys)))

	for _, key := range keys {
		if key == nil {
			continue
		}

		newSlice = append(
			newSlice,
			key.MainName())

		newSlice = appendStringsWithBaseAnyItems(
			it.option.IsSkipEmptyEntry,
			newSlice,
			key.keyChains)
	}

	return it.Compile(newSlice...)
}

func (it *Key) Finalized(
	items ...any,
) *Key {
	it.AppendChain(items...)
	compiled := it.rootCompile(it.option.Joiner)
	it.compiledChain = &compiled

	return it
}

func (it *Key) rootCompile(
	joiner string,
	items ...any,
) string {
	if it.IsComplete() {
		return it.onCompleteCompileInternal(joiner, items)
	}

	finalSlice := make([]string, 0, it.Length()+len(items)+constants.Capacity2)
	finalSlice = append(finalSlice, it.MainName())
	finalSlice = append(finalSlice, it.keyChains...)
	finalSlice = appendAnyItemsWithBaseStrings(
		it.option.IsSkipEmptyEntry,
		finalSlice,
		items)

	return it.compileFinalStrings(joiner, finalSlice)
}

func (it *Key) rootCompileUsingStrings(
	joiner string,
	items ...string,
) string {
	if it.IsComplete() {
		return it.onCompleteCompileInternalStrings(joiner, items)
	}

	finalSlice := make([]string, 0, it.Length()+len(items)+constants.Capacity2)
	finalSlice = append(finalSlice, it.MainName())
	finalSlice = append(finalSlice, it.keyChains...)
	finalSlice = stringslice.AppendStringsWithMainSlice(
		it.option.IsSkipEmptyEntry,
		finalSlice,
		items...)

	return it.compileFinalStrings(joiner, finalSlice)
}

func (it *Key) onCompleteCompileInternal(
	joiner string,
	items []any,
) string {
	if len(items) == 0 {
		return *it.compiledChain
	}

	additionalCompiled := it.compileCompleteAdditional(
		joiner,
		items...)

	if additionalCompiled == constants.EmptyString {
		return *it.compiledChain
	}

	compiledTerms := []string{
		*it.compiledChain,
		additionalCompiled,
	}

	return strings.Join(compiledTerms, joiner)
}

func (it *Key) onCompleteCompileInternalStrings(
	joiner string,
	items []string,
) string {
	if len(items) == 0 {
		return *it.compiledChain
	}

	additionalCompiled := it.compileCompleteAdditionalStrings(
		joiner,
		items...)

	if additionalCompiled == constants.EmptyString {
		return *it.compiledChain
	}

	compiledTerms := []string{
		*it.compiledChain,
		additionalCompiled,
	}

	return strings.Join(compiledTerms, joiner)
}

// CompileReplaceCurlyKeyMap
//
// Keys will be converted to {Key} then replaced
func (it *Key) CompileReplaceCurlyKeyMap(
	mapToReplace map[string]string,
) string {
	return it.CompileReplaceMapUsingItemsOption(
		true,
		mapToReplace,
	)
}

// CompileReplaceCurlyKeyMapUsingItems
//
// Keys will be converted to {Key} then replaced
func (it *Key) CompileReplaceCurlyKeyMapUsingItems(
	mapToReplace map[string]string,
	additionalItems ...any,
) string {
	return it.CompileReplaceMapUsingItemsOption(
		true,
		mapToReplace,
		additionalItems...)
}

func (it *Key) CompileReplaceMapUsingItemsOption(
	isConvKeysToCurlyBraceKeys bool, // conv key to {key} before replace
	mapToReplace map[string]string,
	additionalItems ...any,
) string {
	format := it.Compile(additionalItems...)

	if len(mapToReplace) == 0 {
		return format
	}

	if isConvKeysToCurlyBraceKeys {
		for key, valueToReplace := range mapToReplace {
			keyCurly := fmt.Sprintf(
				constants.CurlyWrapFormat,
				key)

			format = strings.ReplaceAll(
				format,
				keyCurly,
				valueToReplace)
		}

		return format
	}

	for key, valueToReplace := range mapToReplace {
		format = strings.ReplaceAll(
			format,
			key,
			valueToReplace)
	}

	return format
}

func (it *Key) compileFinalStrings(
	joiner string, items []string,
) string {
	if it.option.IsUseBrackets {
		items = it.addBracketsStrings(items)
	}

	return strings.Join(items, joiner)
}

func (it *Key) addBracketsStrings(
	items []string,
) []string {
	for i, item := range items {
		items[i] = it.option.StartBracket + item + it.option.EndBracket
	}

	return items
}

func (it *Key) CompileDefault() string {
	return it.rootCompile(
		it.option.Joiner,
	)
}

func (it *Key) Compile(
	items ...any,
) string {
	return it.rootCompile(
		it.option.Joiner,
		items...)
}

func (it *Key) CompileStrings(
	items ...string,
) string {
	return it.rootCompileUsingStrings(
		it.option.Joiner,
		items...)
}

func (it *Key) JoinUsingJoiner(
	joiner string,
	items ...any,
) string {
	return it.rootCompile(joiner, items...)
}

func (it *Key) JoinUsingOption(
	tempOption *Option,
	items ...any,
) string {
	cloned := it.ClonePtr()
	cloned.option = tempOption

	return cloned.Compile(items...)
}

func (it *Key) compileCompleteAdditional(joiner string, items ...any) string {
	finalSlice := make([]string, 0, len(items))
	finalSlice = appendAnyItemsWithBaseStrings(
		it.option.IsSkipEmptyEntry,
		finalSlice,
		items)

	return it.compileFinalStrings(joiner, finalSlice)
}

func (it *Key) compileCompleteAdditionalStrings(joiner string, items ...string) string {
	finalSlice := make([]string, 0, len(items))
	finalSlice = stringslice.AppendStringsWithMainSlice(
		it.option.IsSkipEmptyEntry,
		finalSlice,
		items...)

	return it.compileFinalStrings(joiner, finalSlice)
}

func (it *Key) IntRange(
	startIncluding, endIncluding int,
) []string {
	keyOuts := make(
		[]string,
		0,
		endIncluding-startIncluding+1)

	for i := startIncluding; i <= endIncluding; i++ {
		keyOuts = append(
			keyOuts,
			it.CompileStrings(strconv.Itoa(i)))
	}

	return keyOuts
}

func (it *Key) IntRangeEnding(
	endIncluding int,
) []string {
	return it.IntRange(
		constants.Zero,
		endIncluding)
}
