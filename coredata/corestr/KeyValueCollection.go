package corestr

import (
	"encoding/json"
	"sort"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/defaultcapacity"
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

type KeyValueCollection struct {
	KeyValuePairs []KeyValuePair `json:"KeyValuePairs,omitempty"`
}

func (it *KeyValueCollection) AllKeysSorted() []string {
	keys := it.AllKeys()

	sort.Strings(keys)

	return keys
}

func (it *KeyValueCollection) KeysHashset() map[string]bool {
	panic("implement me")
}

func (it *KeyValueCollection) HasKey(key string) bool {
	for _, pair := range it.KeyValuePairs {
		if pair.IsKey(key) {
			return true
		}
	}

	return false
}

func (it KeyValueCollection) SerializeMust() (jsonBytes []byte) {
	return corejson.NewPtr(it).RawMust()
}

func (it *KeyValueCollection) Compile() string {
	return it.String()
}

func (it *KeyValueCollection) Count() int {
	return it.Length()
}

func (it *KeyValueCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *KeyValueCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *KeyValueCollection) HasIndex(
	index int,
) bool {
	return index != constants.InvalidNotFoundCase && it.LastIndex() >= index
}

func (it *KeyValueCollection) First() *KeyValuePair {
	return &it.KeyValuePairs[0]
}

func (it *KeyValueCollection) FirstOrDefault() *KeyValuePair {
	if it.IsEmpty() {
		return nil
	}

	return &it.KeyValuePairs[0]
}

func (it *KeyValueCollection) Last() *KeyValuePair {
	return &it.KeyValuePairs[it.LastIndex()]
}

func (it *KeyValueCollection) LastOrDefault() *KeyValuePair {
	if it.IsEmpty() {
		return nil
	}

	return it.Last()
}

func (it *KeyValueCollection) Find(
	finder func(index int, currentKeyVal KeyValuePair) (foundItem KeyValuePair, isFound, isBreak bool),
) []KeyValuePair {
	length := it.Length()

	if length == 0 {
		return []KeyValuePair{}
	}

	slice := make(
		[]KeyValuePair,
		0,
		defaultcapacity.OfSearch(length),
	)

	for i, item := range it.KeyValuePairs {
		foundItem, isFound, isBreak := finder(i, item)

		if isFound {
			slice = append(slice, foundItem)
		}

		if isBreak {
			return slice
		}
	}

	return slice
}

func (it *KeyValueCollection) SafeValueAt(index int) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	if it.HasIndex(index) {
		return it.KeyValuePairs[index].Value
	}

	return constants.EmptyString
}

func (it *KeyValueCollection) SafeValuesAtIndexes(indexes ...int) []string {
	requestLength := len(indexes)
	slice := make([]string, requestLength)

	if requestLength == 0 {
		return slice
	}

	for i, index := range indexes {
		slice[i] = it.SafeValueAt(index)
	}

	return slice
}

func (it *KeyValueCollection) Strings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, keyVal := range it.KeyValuePairs {
		slice[i] = keyVal.String()
	}

	return slice
}

func (it *KeyValueCollection) StringsUsingFormat(
	format string,
) []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, keyVal := range it.KeyValuePairs {
		slice[i] = keyVal.FormatString(format)
	}

	return slice
}

func (it *KeyValueCollection) String() string {
	return strutilinternal.AnyToString(it.Strings())
}

func (it *KeyValueCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.KeyValuePairs)
}

func (it *KeyValueCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *KeyValueCollection) Add(key, val string) *KeyValueCollection {
	it.KeyValuePairs = append(
		it.KeyValuePairs, KeyValuePair{
			Key:   key,
			Value: val,
		},
	)

	return it
}

func (it *KeyValueCollection) AddIf(
	isAdd bool,
	key, val string,
) *KeyValueCollection {
	if !isAdd {
		return it
	}

	it.KeyValuePairs = append(
		it.KeyValuePairs, KeyValuePair{
			Key:   key,
			Value: val,
		},
	)

	return it
}

func (it *KeyValueCollection) AddStringBySplit(
	splitter,
	line string,
) *KeyValueCollection {
	key, val := strutilinternal.SplitLeftRight(
		splitter,
		line,
	)

	return it.Add(key, val)
}

func (it *KeyValueCollection) AddStringBySplitTrim(
	splitter,
	line string,
) *KeyValueCollection {
	key, val := strutilinternal.SplitLeftRightTrim(
		splitter,
		line,
	)

	return it.Add(key, val)
}

func (it *KeyValueCollection) Adds(
	keyValues ...KeyValuePair,
) *KeyValueCollection {
	if len(keyValues) == 0 {
		return it
	}

	for _, keyVal := range keyValues {
		it.KeyValuePairs = append(
			it.KeyValuePairs, KeyValuePair{
				Key:   keyVal.Key,
				Value: keyVal.Value,
			},
		)
	}

	return it
}

func (it *KeyValueCollection) AddMap(
	inputMap map[string]string,
) *KeyValueCollection {
	if inputMap == nil || len(inputMap) == 0 {
		return it
	}

	for key, val := range inputMap {
		it.KeyValuePairs = append(
			it.KeyValuePairs, KeyValuePair{
				Key:   key,
				Value: val,
			},
		)
	}

	return it
}

func (it *KeyValueCollection) AddHashsetMap(
	inputMap map[string]bool,
) *KeyValueCollection {
	if inputMap == nil || len(inputMap) == 0 {
		return it
	}

	for key := range inputMap {
		it.KeyValuePairs = append(
			it.KeyValuePairs, KeyValuePair{
				Key:   key,
				Value: key,
			},
		)
	}

	return it
}

func (it *KeyValueCollection) AddHashset(
	inputHashset *Hashset,
) *KeyValueCollection {
	if inputHashset == nil || inputHashset.IsEmpty() {
		return it
	}

	for key := range inputHashset.items {
		it.KeyValuePairs = append(
			it.KeyValuePairs, KeyValuePair{
				Key:   key,
				Value: key,
			},
		)
	}

	return it
}

func (it *KeyValueCollection) AddsHashmap(
	hashmap *Hashmap,
) *KeyValueCollection {
	if hashmap == nil || hashmap.IsEmpty() {
		return it
	}

	for key, val := range hashmap.items {
		it.KeyValuePairs = append(
			it.KeyValuePairs, KeyValuePair{
				Key:   key,
				Value: val,
			},
		)
	}

	return it
}

func (it *KeyValueCollection) Hashmap() *Hashmap {
	length := it.Length()
	hashmap := New.Hashmap.Cap(length)

	if length == 0 {
		return hashmap
	}

	for _, keyVal := range it.KeyValuePairs {
		hashmap.AddOrUpdate(keyVal.Key, keyVal.Value)
	}

	return hashmap
}

func (it *KeyValueCollection) IsContains(key string) bool {
	for _, pair := range it.KeyValuePairs {
		if pair.Key == key {
			return true
		}
	}

	return false
}

func (it *KeyValueCollection) Get(key string) (string, bool) {
	for _, pair := range it.KeyValuePairs {
		if pair.Key == key {
			return pair.Value, true
		}
	}

	return "", false
}

func (it *KeyValueCollection) Map() map[string]string {
	hashmap := it.Hashmap()

	return hashmap.items
}

func (it *KeyValueCollection) AddsHashmaps(
	hashmaps ...*Hashmap,
) *KeyValueCollection {
	if hashmaps == nil || len(hashmaps) == 0 {
		return it
	}

	for _, hashmap := range hashmaps {
		it.AddsHashmap(hashmap)
	}

	return it
}

func (it *KeyValueCollection) AllKeys() []string {
	length := len(it.KeyValuePairs)
	keys := make([]string, length)

	if length == 0 {
		return keys
	}

	i := 0
	for _, item := range it.KeyValuePairs {
		keys[i] = item.Key
		i++
	}

	return keys
}

func (it *KeyValueCollection) AllValues() []string {
	length := len(it.KeyValuePairs)
	values := make([]string, length)

	if length == 0 {
		return values
	}

	i := 0
	for _, item := range it.KeyValuePairs {
		values[i] = item.Value
		i++
	}

	return values
}

// Join values
func (it *KeyValueCollection) Join(
	separator string,
) string {
	return strings.Join(it.Strings(), separator)
}

func (it *KeyValueCollection) JoinKeys(
	separator string,
) string {
	return strings.Join(it.AllKeys(), separator)
}

func (it *KeyValueCollection) JoinValues(
	separator string,
) string {
	return strings.Join(it.AllValues(), separator)
}

func (it *KeyValueCollection) JsonModel() []KeyValuePair {
	return it.KeyValuePairs
}

func (it *KeyValueCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *KeyValueCollection) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *KeyValueCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *KeyValueCollection) UnmarshalJSON(data []byte) error {
	var dataModelItems []KeyValuePair
	err := corejson.Deserialize.UsingBytes(
		data,
		&dataModelItems,
	)

	if err == nil && len(dataModelItems) > 0 {
		it.KeyValuePairs = dataModelItems
	} else if err == nil {
		it.KeyValuePairs = []KeyValuePair{}
	}

	return err
}

func (it KeyValueCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it KeyValueCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *KeyValueCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*KeyValueCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

func (it *KeyValueCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *KeyValueCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *KeyValueCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *KeyValueCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *KeyValueCollection) Clear() {
	if it == nil {
		return
	}

	tempItems := it.KeyValuePairs
	clearFunc := func() {
		for i := 0; i < len(tempItems); i++ {
			tempItems[i].Dispose()
		}
	}

	go clearFunc()
	it.KeyValuePairs = []KeyValuePair{}
}

func (it *KeyValueCollection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}

func (it *KeyValueCollection) Deserialize(toPtr interface{}) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}
