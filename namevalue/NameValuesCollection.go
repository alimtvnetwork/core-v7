package namevalue

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type NameValuesCollection struct {
	Items        []Instance `json:"Items,omitempty"`
	lazyToString *string
}

func NewNameValuesCollection(capacity int) *NameValuesCollection {
	slice := make([]Instance, 0, capacity)

	return &NameValuesCollection{
		Items: slice,
	}
}

func NewCollection() *NameValuesCollection {
	return NewNameValuesCollection(constants.Capacity5)
}

func NewNewNameValuesCollectionUsing(
	isClone bool,
	items ...Instance,
) *NameValuesCollection {
	if items == nil {
		return EmptyNameValuesCollection()
	}

	if !isClone {
		return &NameValuesCollection{
			Items: items,
		}
	}

	slice := NewNameValuesCollection(len(items))

	return slice.Adds(items...)
}

func EmptyNameValuesCollection() *NameValuesCollection {
	return NewNameValuesCollection(0)
}

func (it *NameValuesCollection) Add(
	item Instance,
) *NameValuesCollection {
	it.InvalidateLazyString()
	it.Items = append(it.Items, item)

	return it
}

func (it *NameValuesCollection) Adds(
	items ...Instance,
) *NameValuesCollection {
	if len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		it.Items,
		items...)

	return it
}

func (it *NameValuesCollection) Append(
	items ...Instance,
) *NameValuesCollection {
	if len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		it.Items,
		items...)

	return it
}

func (it *NameValuesCollection) AppendIf(
	isAppend bool,
	items ...Instance,
) *NameValuesCollection {
	if !isAppend || len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		it.Items,
		items...)

	return it
}

func (it *NameValuesCollection) Prepend(
	items ...Instance,
) *NameValuesCollection {
	if len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		items,
		it.Items...)

	return it
}

func (it *NameValuesCollection) PrependIf(
	isPrepend bool,
	items ...Instance,
) *NameValuesCollection {
	if !isPrepend || len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		items,
		it.Items...)

	return it
}

func (it *NameValuesCollection) PrependUsingFuncIf(
	isPrepend bool,
	itemsGetterFunc func() []Instance,
) *NameValuesCollection {
	if !isPrepend || itemsGetterFunc == nil {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		itemsGetterFunc(),
		it.Items...)

	return it
}

func (it *NameValuesCollection) AppendUsingFuncIf(
	isAppend bool,
	itemsGetterFunc func() []Instance,
) *NameValuesCollection {
	if !isAppend || itemsGetterFunc == nil {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		it.Items,
		itemsGetterFunc()...,
	)

	return it
}

func (it *NameValuesCollection) AppendPrependIf(
	isAppendOrPrepend bool,
	prependItems []Instance,
	appendItems []Instance,
) *NameValuesCollection {
	if !isAppendOrPrepend {
		return it
	}

	if len(prependItems) > 0 {
		it.InvalidateLazyString()
		it.Items = append(
			prependItems,
			it.Items...)
	}

	if len(appendItems) > 0 {
		it.InvalidateLazyString()
		it.Items = append(
			it.Items,
			appendItems...)
	}

	return it
}

func (it *NameValuesCollection) AddsPtr(
	items ...*Instance,
) *NameValuesCollection {
	if len(items) == 0 {
		return it
	}

	for _, item := range items {
		if item == nil {
			continue
		}

		it.Items = append(
			it.Items,
			*item)
	}

	return it
}

func (it *NameValuesCollection) HasCompiledString() bool {
	return it != nil && it.lazyToString != nil
}

func (it *NameValuesCollection) InvalidateLazyString() {
	if it == nil {
		return
	}

	it.lazyToString = nil
}

func (it *NameValuesCollection) CompiledLazyString() string {
	if it == nil {
		return constants.EmptyString
	}

	if it.lazyToString != nil {
		return *it.lazyToString
	}

	toString := it.String()
	it.lazyToString = &toString

	return toString
}

func (it *NameValuesCollection) ConcatNew(
	additionalItems ...Instance,
) *NameValuesCollection {
	cloned := it.Clone()

	return cloned.Adds(additionalItems...)
}

func (it *NameValuesCollection) ConcatNewPtr(
	additionalItems ...*Instance,
) *NameValuesCollection {
	cloned := it.Clone()

	return cloned.AddsPtr(
		additionalItems...)
}

func (it *NameValuesCollection) AddsIf(
	isAdd bool,
	items ...Instance,
) *NameValuesCollection {
	if !isAdd {
		return it
	}

	it.InvalidateLazyString()

	return it.Adds(items...)
}

func (it *NameValuesCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *NameValuesCollection) Count() int {
	return it.Length()
}

func (it *NameValuesCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *NameValuesCollection) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *NameValuesCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *NameValuesCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *NameValuesCollection) Strings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.String()
	}

	return list
}

func (it *NameValuesCollection) JsonStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.JsonString()
	}

	return list
}

func (it *NameValuesCollection) JoinJsonStrings(joiner string) string {
	return strings.Join(it.JsonStrings(), joiner)
}

func (it *NameValuesCollection) Join(joiner string) string {
	return strings.Join(it.Strings(), joiner)
}

func (it *NameValuesCollection) JoinLines() string {
	return strings.Join(it.Strings(), constants.DefaultLine)
}

func (it *NameValuesCollection) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it *NameValuesCollection) JoinCsvLine() string {
	return strings.Join(it.CsvStrings(), constants.CommaUnixNewLine)
}

func (it *NameValuesCollection) IsEqual(another *NameValuesCollection) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.Length() != another.Length() {
		return false
	}

	return it.IsEqualItems(another.Items...)
}

func (it *NameValuesCollection) IsEqualItems(lines ...Instance) bool {
	if it == nil && lines == nil {
		return true
	}

	if it == nil || lines == nil {
		return false
	}

	if it.Length() != len(lines) {
		return false
	}

	for i, item := range it.Items {
		anotherItem := lines[i]

		if item != anotherItem {
			return false
		}
	}

	return true
}

func (it NameValuesCollection) JsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	jsonBytes, err := json.Marshal(it)

	if err != nil || jsonBytes == nil {
		return constants.EmptyString
	}

	return string(jsonBytes)
}

func (it *NameValuesCollection) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	if it.HasCompiledString() {
		return *it.lazyToString
	}

	return it.JoinLines()
}

func (it *NameValuesCollection) Error() error {
	if it.IsEmpty() {
		return nil
	}

	return errors.New(it.String())
}

func (it *NameValuesCollection) ErrorUsingMessage(message string) error {
	if it.IsEmpty() {
		return nil
	}

	toCompiled := message + constants.Space + it.String()

	return errors.New(toCompiled)
}

func (it *NameValuesCollection) CsvStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	newSlice := make([]string, it.Length())

	for i, item := range it.Items {
		newSlice[i] = fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			item.String())
	}

	return newSlice
}

func (it *NameValuesCollection) Clear() *NameValuesCollection {
	if it == nil {
		return it
	}

	tempItems := it.Items
	clearFunc := func() {
		for _, item := range tempItems {
			item.Dispose()
		}
	}

	go clearFunc()

	it.Items = []Instance{}
	it.lazyToString = nil

	return it
}

func (it *NameValuesCollection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it NameValuesCollection) Clone() NameValuesCollection {
	list := NewNameValuesCollection(it.Length())

	return *list.Adds(it.Items...)
}

func (it *NameValuesCollection) ClonePtr() *NameValuesCollection {
	if it == nil {
		return nil
	}

	list := NewNameValuesCollection(it.Length())

	return list.Adds(it.Items...)
}
