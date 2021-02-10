package corestr

import (
	"encoding/json"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type HashsetsCollection struct {
	items *[]*Hashset
}

func (hashsetsCollection *HashsetsCollection) IsEmpty() bool {
	return hashsetsCollection.items == nil ||
		*hashsetsCollection.items == nil ||
		len(*hashsetsCollection.items) == 0
}

func (hashsetsCollection *HashsetsCollection) IndexOf(index int) *Hashset {
	if hashsetsCollection.IsEmpty() ||
		hashsetsCollection.Length()-1 > index {
		return nil
	}

	hashset := (*hashsetsCollection.items)[index]

	return hashset
}

func (hashsetsCollection *HashsetsCollection) ListPtr() *[]*Hashset {
	return hashsetsCollection.items
}

func (hashsetsCollection *HashsetsCollection) List() []*Hashset {
	return *hashsetsCollection.items
}

func (hashsetsCollection *HashsetsCollection) StringsList() *[]string {
	if hashsetsCollection.IsEmpty() {
		return constants.EmptyStringsPtr
	}

	completeLength := 0
	for _, hashset := range *hashsetsCollection.items {
		completeLength += hashset.Length()
	}

	stringsList := make([]string, completeLength)
	index := 0

	for _, hashset := range *hashsetsCollection.items {
		for _, item := range *hashset.ListPtr() {
			stringsList[index] = item
			index++
		}
	}

	return &stringsList
}

// items returns false
// hashsetsCollection empty returns false
func (hashsetsCollection *HashsetsCollection) HasAll(items ...string) bool {
	if hashsetsCollection.IsEmpty() || items == nil {
		return false
	}

	length := hashsetsCollection.Length()
	boolList := make([]bool, length)
	wg := &sync.WaitGroup{}
	wg.Add(length)
	hasFunc := func(i int, wg *sync.WaitGroup) {
		boolList[i] = (*hashsetsCollection.items)[i].HasAllStringsPtr(&items)
		wg.Done()
	}

	for i := 0; i < length; i++ {
		go hasFunc(i, wg)
	}

	wg.Wait()

	for i := 0; i < length; i++ {
		if boolList[i] {
			return true
		}
	}

	return false
}

func (hashsetsCollection *HashsetsCollection) ListDirectPtr() *[]Hashset {
	list := make([]Hashset, hashsetsCollection.Length())

	for i, hashset := range *hashsetsCollection.items {
		//goland:noinspection GoLinterLocal,GoVetCopyLock
		list[i] = *hashset //nolint:govet
	}

	return &list
}

func (hashsetsCollection *HashsetsCollection) Add(hashset *Hashset) *HashsetsCollection {
	*hashsetsCollection.items = append(*hashsetsCollection.items, hashset)

	return hashsetsCollection
}

func (hashsetsCollection *HashsetsCollection) AddNonNil(hashset *Hashset) *HashsetsCollection {
	if hashset == nil {
		return hashsetsCollection
	}

	*hashsetsCollection.items = append(*hashsetsCollection.items, hashset)

	return hashsetsCollection
}

func (hashsetsCollection *HashsetsCollection) AddNonEmpty(
	hashset *Hashset,
) *HashsetsCollection {
	if hashset == nil || hashset.IsEmpty() {
		return hashsetsCollection
	}

	*hashsetsCollection.items = append(*hashsetsCollection.items, hashset)

	return hashsetsCollection
}

// nil will be skipped
func (hashsetsCollection *HashsetsCollection) Adds(
	hashsets ...*Hashset,
) *HashsetsCollection {
	if hashsets == nil {
		return hashsetsCollection
	}

	for _, hashset := range hashsets {
		if hashset == nil || hashset.IsEmpty() {
			continue
		}

		*hashsetsCollection.items = append(
			*hashsetsCollection.items,
			hashset)
	}

	return hashsetsCollection
}

func (hashsetsCollection *HashsetsCollection) Length() int {
	if hashsetsCollection.items == nil ||
		*hashsetsCollection.items == nil {
		return 0
	}

	return len(*hashsetsCollection.items)
}

func (hashsetsCollection *HashsetsCollection) IsEqual(another HashsetsCollection) bool {
	return hashsetsCollection.IsEqualPtr(&another)
}

func (hashsetsCollection *HashsetsCollection) IsEqualPtr(another *HashsetsCollection) bool {
	if hashsetsCollection == nil && another == nil {
		return true
	}

	if hashsetsCollection == nil || another == nil {
		return false
	}

	if hashsetsCollection == another {
		// ptr same
		return true
	}

	if hashsetsCollection.IsEmpty() && another.IsEmpty() {
		return true
	}

	if hashsetsCollection.IsEmpty() || another.IsEmpty() {
		return false
	}

	leftLength := hashsetsCollection.Length()
	rightLength := another.Length()

	if leftLength != rightLength {
		return false
	}

	for i, hashset := range *hashsetsCollection.items {
		anotherHashset := (*another.items)[i]

		if !hashset.IsEqualsPtr(anotherHashset) {
			return false
		}
	}

	return true
}

func (hashsetsCollection *HashsetsCollection) JsonModel() *HashsetsCollectionDataModel {
	return NewHashsetsCollectionDataModelUsing(hashsetsCollection)
}

func (hashsetsCollection *HashsetsCollection) JsonModelAny() interface{} {
	return hashsetsCollection.JsonModel()
}

func (hashsetsCollection *HashsetsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(hashsetsCollection.JsonModel())
}

func (hashsetsCollection *HashsetsCollection) UnmarshalJSON(
	data []byte,
) error {
	var dataModel HashsetsCollectionDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		hashsetsCollection.items = dataModel.Items
	}

	return err
}

func (hashsetsCollection *HashsetsCollection) Json() *corejson.Result {
	if hashsetsCollection.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(hashsetsCollection)

	return corejson.NewPtr(jsonBytes, err)
}

func (hashsetsCollection *HashsetsCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*HashsetsCollection, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyHashsetsCollection(), nil
	}

	err := json.Unmarshal(*jsonResult.Bytes, &hashsetsCollection)

	if err != nil {
		return EmptyHashsetsCollection(), err
	}

	return hashsetsCollection, nil
}

// Panic if error
func (hashsetsCollection *HashsetsCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *HashsetsCollection {
	hashSet, err := hashsetsCollection.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (hashsetsCollection *HashsetsCollection) String() string {
	if hashsetsCollection.IsEmpty() {
		return commonJoiner + NoElements
	}

	strList := make([]string, hashsetsCollection.Length())

	for i, hashset := range *hashsetsCollection.items {
		strList[i] = hashset.String()
	}

	return strings.Join(
		strList,
		"")
}

func (hashsetsCollection *HashsetsCollection) Join(
	separator string,
) string {
	return strings.Join(
		*hashsetsCollection.StringsList(),
		separator)
}

func (hashsetsCollection *HashsetsCollection) AsJsoner() *corejson.Jsoner {
	var jsoner corejson.Jsoner = hashsetsCollection

	return &jsoner
}

func (hashsetsCollection *HashsetsCollection) JsonParseSelfInject(jsonResult *corejson.Result) {
	hashsetsCollection.ParseInjectUsingJsonMust(jsonResult)
}

func (hashsetsCollection *HashsetsCollection) AsJsonParseSelfInjector() *corejson.ParseSelfInjector {
	var jsonMarshaller corejson.ParseSelfInjector = hashsetsCollection

	return &jsonMarshaller
}

func (hashsetsCollection *HashsetsCollection) AsJsonMarshaller() *corejson.JsonMarshaller {
	var jsonMarshaller corejson.JsonMarshaller = hashsetsCollection

	return &jsonMarshaller
}
