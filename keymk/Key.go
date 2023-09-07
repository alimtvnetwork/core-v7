package keymk

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/defaultcapacity"
	"gitlab.com/auk-go/core/errcore"
)

type Key struct {
	option        *Option
	mainName      string
	keyChains     []string
	compiledChain *string
}

func (it *Key) CompiledChain() string {
	if it.IsComplete() {
		return *it.compiledChain
	}

	return constants.EmptyString
}

func (it *Key) MainName() string {
	return it.mainName
}

func (it *Key) IsEmpty() bool {
	return it.Length() == 0 && it.MainName() == ""
}

func (it *Key) Length() int {
	if it == nil {
		return 0
	}

	return len(it.keyChains)
}

func (it *Key) AppendChain(items ...interface{}) *Key {
	if it.IsComplete() {
		// panic
		errcore.CannotModifyCompleteResourceType.HandleUsingPanic(
			cannotModifyErrorMessage,
			items)
	}

	it.keyChains = appendAnyItemsWithBaseStrings(
		it.option.IsSkipEmptyEntry,
		it.keyChains,
		items)

	return it
}

func (it *Key) AppendChainKeys(
	keys ...*Key,
) *Key {
	if len(keys) == 0 {
		return it
	}

	for _, key := range keys {
		if key == nil {
			continue
		}

		it.AppendChainStrings(key.MainName())
		it.AppendChainStrings(key.keyChains...)
	}

	return it
}

func (it *Key) CompileKeys(
	keys ...*Key,
) string {
	if len(keys) == 0 {
		return it.Compile()
	}

	newSlice := make(
		[]interface{},
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

func (it *Key) AppendChainStrings(
	items ...string,
) *Key {
	if it.IsComplete() {
		// panic
		errcore.CannotModifyCompleteResourceType.HandleUsingPanic(
			cannotModifyErrorMessage,
			items)
	}

	isSkipOnEmpty := it.option.IsSkipEmptyEntry

	for _, item := range items {
		if isSkipOnEmpty && item == "" {
			continue
		}

		it.keyChains = append(
			it.keyChains,
			item)
	}

	return it
}

func (it *Key) KeyChains() []string {
	if it == nil {
		return nil
	}

	return it.keyChains
}

// AllRawItems
//
// Returns main + whole chain (raw elements)
func (it *Key) AllRawItems() []string {
	if it == nil {
		return nil
	}

	return stringslice.PrependLineNew(
		it.MainName(),
		it.KeyChains())
}

func (it *Key) HasInChains(
	chainItem string,
) bool {
	if it == nil {
		return false
	}

	for _, chain := range it.keyChains {
		if chain == chainItem {
			return true
		}
	}

	return false
}

func (it *Key) IsComplete() bool {
	return it.compiledChain != nil
}

func (it *Key) Finalized(
	items ...interface{},
) *Key {
	it.AppendChain(items...)
	compiled := it.rootCompile(it.option.Joiner)
	it.compiledChain = &compiled

	return it
}

func (it *Key) rootCompile(
	joiner string,
	items ...interface{},
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
	items []interface{},
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

func (it *Key) compileSingleItem(
	item string,
) string {
	if it.option.IsUseBrackets {
		return it.option.StartBracket + item + it.option.EndBracket
	}

	return item
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
	additionalItems ...interface{},
) string {
	return it.CompileReplaceMapUsingItemsOption(
		true,
		mapToReplace,
		additionalItems...)
}

func (it *Key) CompileReplaceMapUsingItemsOption(
	isConvKeysToCurlyBraceKeys bool, // conv key to {key} before replace
	mapToReplace map[string]string,
	additionalItems ...interface{},
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

func (it *Key) ConcatNewUsingKeys(
	keys ...*Key,
) *Key {
	cloned := it.ClonePtr()

	return cloned.AppendChainKeys(keys...)
}

func (it *Key) ClonePtr(
	newAppendingChains ...interface{},
) *Key {
	if it == nil {
		return nil
	}

	key := NewKey.All(
		it.option.ClonePtr(),
		it.mainName,
	)

	key.AppendChainStrings(
		it.keyChains...)

	return key.AppendChain(
		newAppendingChains...)
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

func (it *Key) CompileDefault() string {
	return it.rootCompile(
		it.option.Joiner,
	)
}

func (it *Key) Compile(
	items ...interface{},
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
	items ...interface{},
) string {
	return it.rootCompile(joiner, items...)
}

func (it *Key) JoinUsingOption(
	tempOption *Option,
	items ...interface{},
) string {
	temp2 := it.option
	it.option = tempOption
	compiled := it.Compile(items...)
	it.option = temp2

	return compiled
}

func (it *Key) String() string {
	return it.Compile()
}

func (it *Key) Strings() []string {
	return it.AllRawItems()
}

func (it *Key) Name() string {
	return it.Compile()
}

func (it *Key) KeyCompiled() string {
	return it.Compile()
}

func (it *Key) compileCompleteAdditional(joiner string, items ...interface{}) string {
	if len(items) == 0 {
		return constants.EmptyString
	}

	finalSlice := make([]string, 0, len(items))
	finalSlice = appendAnyItemsWithBaseStrings(
		it.option.IsSkipEmptyEntry,
		finalSlice,
		items)

	return it.compileFinalStrings(joiner, finalSlice)
}

func (it *Key) compileCompleteAdditionalStrings(joiner string, items ...string) string {
	if len(items) == 0 {
		return constants.EmptyString
	}

	finalSlice := make([]string, 0, len(items))
	finalSlice = stringslice.AppendStringsWithMainSlice(
		it.option.IsSkipEmptyEntry,
		finalSlice,
		items...)

	return it.compileFinalStrings(joiner, finalSlice)
}

func (it *Key) TemplateReplacer() templateReplacer {
	return templateReplacer{
		it,
	}
}

func (it *Key) JsonModel() keyModel {
	return keyModel{
		Option:        it.option,
		MainName:      it.mainName,
		KeyChains:     it.keyChains,
		CompiledChain: it.compiledChain,
	}
}

func (it *Key) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it Key) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Key) UnmarshalJSON(data []byte) error {
	var deserializedModel keyModel
	err := json.Unmarshal(data, &deserializedModel)

	if err == nil {
		it.option = deserializedModel.Option
		it.mainName = deserializedModel.MainName
		it.keyChains = deserializedModel.KeyChains
		it.compiledChain = deserializedModel.CompiledChain
	}

	return err
}

func (it Key) Json() corejson.Result {
	return corejson.New(it)
}

func (it Key) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Key) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *Key) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Key, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *Key) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Key {
	deserialized, err := it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return deserialized
}

func (it *Key) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Key) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Key) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it Key) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return &it
}

func (it Key) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}
