package corestr

import (
	"encoding/json"
	"strings"
	"sync"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
)

type HashsetsCollection struct {
	items []*Hashset
}

func (it *HashsetsCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *HashsetsCollection) HasItems() bool {
	return it.Length() > 0
}

func (it *HashsetsCollection) IndexOf(index int) *Hashset {
	if it.IsEmpty() ||
		it.Length()-1 > index {
		return nil
	}

	hashset := it.items[index]

	return hashset
}

func (it *HashsetsCollection) ListPtr() *[]*Hashset {
	return &it.items
}

func (it *HashsetsCollection) List() []*Hashset {
	return it.items
}

func (it *HashsetsCollection) StringsList() []string {
	if it.IsEmpty() {
		return []string{}
	}

	completeLength := 0
	for _, hashset := range it.items {
		completeLength += hashset.Length()
	}

	stringsList := make([]string, completeLength)
	index := 0

	for _, hashset := range it.items {
		for _, item := range *hashset.ListPtr() {
			stringsList[index] = item
			index++
		}
	}

	return stringsList
}

// HasAll items returns false
// hashsetsCollection Empty returns false
func (it *HashsetsCollection) HasAll(items ...string) bool {
	if it.IsEmpty() || items == nil {
		return false
	}

	length := it.Length()
	boolList := make([]bool, length)
	wg := sync.WaitGroup{}
	wg.Add(length)
	hasFunc := func(i int) {
		boolList[i] = it.items[i].
			HasAllStrings(items)
		wg.Done()
	}

	for i := 0; i < length; i++ {
		go hasFunc(i)
	}

	wg.Wait()

	for i := 0; i < length; i++ {
		if boolList[i] {
			return true
		}
	}

	return false
}

func (it *HashsetsCollection) ListDirectPtr() *[]Hashset {
	list := make([]Hashset, it.Length())

	for i, hashset := range it.items {
		//goland:noinspection GoLinterLocal,GoVetCopyLock
		list[i] = *hashset //nolint:govet
	}

	return &list
}

func (it *HashsetsCollection) AddHashsetsCollection(
	next *HashsetsCollection,
) *HashsetsCollection {
	if next == nil || next.IsEmpty() {
		return it
	}

	items := it.items

	for _, nextHashset := range next.items {
		items = append(items, nextHashset)
	}

	it.items = items

	return it
}

func (it *HashsetsCollection) ConcatNew(
	nextCollections ...*HashsetsCollection,
) *HashsetsCollection {
	if nextCollections == nil || len(nextCollections) == 0 {
		return New.HashsetsCollection.UsingHashsetsPointers(it.items...)
	}

	length := it.Length() + constants.Capacity4

	for _, collection := range nextCollections {
		length += collection.Length()
	}

	newHashsetsCollection := New.HashsetsCollection.LenCap(constants.Zero, length)
	newHashsetsCollection.AddHashsetsCollection(it)

	for _, collection := range nextCollections {
		newHashsetsCollection.AddHashsetsCollection(collection)
	}

	return newHashsetsCollection
}

func (it *HashsetsCollection) Add(
	hashset *Hashset,
) *HashsetsCollection {
	it.items = append(
		it.items,
		hashset,
	)

	return it
}

func (it *HashsetsCollection) AddNonNil(
	hashset *Hashset,
) *HashsetsCollection {
	if hashset == nil {
		return it
	}

	it.items = append(it.items, hashset)

	return it
}

func (it *HashsetsCollection) AddNonEmpty(
	hashset *Hashset,
) *HashsetsCollection {
	if hashset.IsEmpty() {
		return it
	}

	it.items = append(it.items, hashset)

	return it
}

// Adds nil will be skipped
func (it *HashsetsCollection) Adds(
	hashsets ...*Hashset,
) *HashsetsCollection {
	if hashsets == nil {
		return it
	}

	for _, hashset := range hashsets {
		if hashset.IsEmpty() {
			continue
		}

		it.items = append(
			it.items,
			hashset,
		)
	}

	return it
}

func (it *HashsetsCollection) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *HashsetsCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *HashsetsCollection) IsEqual(another HashsetsCollection) bool {
	return it.IsEqualPtr(&another)
}

func (it *HashsetsCollection) IsEqualPtr(another *HashsetsCollection) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it == another {
		// ptr same
		return true
	}

	if it.IsEmpty() && another.IsEmpty() {
		return true
	}

	if it.IsEmpty() || another.IsEmpty() {
		return false
	}

	leftLength := it.Length()
	rightLength := another.Length()

	if leftLength != rightLength {
		return false
	}

	for i, hashset := range it.items {
		anotherHashset := another.items[i]

		if !hashset.IsEquals(anotherHashset) {
			return false
		}
	}

	return true
}

func (it *HashsetsCollection) JsonModel() *HashsetsCollectionDataModel {
	return NewHashsetsCollectionDataModelUsing(it)
}

func (it *HashsetsCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *HashsetsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *HashsetsCollection) UnmarshalJSON(
	data []byte,
) error {
	var dataModel HashsetsCollectionDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel.Items
	}

	return err
}

func (it HashsetsCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it HashsetsCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *HashsetsCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*HashsetsCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return Empty.HashsetsCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *HashsetsCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *HashsetsCollection {
	hashSet, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (it *HashsetsCollection) String() string {
	if it.IsEmpty() {
		return commonJoiner + NoElements
	}

	strList := make([]string, it.Length())

	for i, hashset := range it.items {
		strList[i] = hashset.String()
	}

	return strings.Join(
		strList,
		"",
	)
}

func (it *HashsetsCollection) Join(
	separator string,
) string {
	return strings.Join(
		it.StringsList(),
		separator,
	)
}

func (it *HashsetsCollection) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *HashsetsCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *HashsetsCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *HashsetsCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *HashsetsCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *HashsetsCollection) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it HashsetsCollection) Deserialize(toPtr interface{}) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}
