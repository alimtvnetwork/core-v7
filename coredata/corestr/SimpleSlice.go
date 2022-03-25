package corestr

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type SimpleSlice struct {
	Items []string `json:"Items,omitempty"`
}

func (it *SimpleSlice) Add(
	item string,
) *SimpleSlice {
	it.Items = append(it.Items, item)

	return it
}

func (it *SimpleSlice) AddIf(
	isAdd bool,
	item string,
) *SimpleSlice {
	if !isAdd {
		return it
	}

	it.Items = append(it.Items, item)

	return it
}

func (it *SimpleSlice) Adds(
	items ...string,
) *SimpleSlice {
	if len(items) == 0 {
		return it
	}

	it.Items = append(it.Items, items...)

	return it
}

func (it *SimpleSlice) Append(
	items ...string,
) *SimpleSlice {
	if len(items) == 0 {
		return it
	}

	it.Items = append(it.Items, items...)

	return it
}

// AppendFmt
//
//  Skips empty format + values
func (it *SimpleSlice) AppendFmt(
	format string,
	v ...interface{},
) *SimpleSlice {
	if format == "" && len(v) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		fmt.Sprintf(format, v...))

	return it
}

// AppendFmtIf
//
//  Skips empty format + values
func (it *SimpleSlice) AppendFmtIf(
	isAppend bool,
	format string,
	v ...interface{},
) *SimpleSlice {
	if !isAppend || format == "" && len(v) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		fmt.Sprintf(format, v...))

	return it
}

// AddAsTitleValue
//
//  Adds Title : value (constants.TitleValueFormat)
func (it *SimpleSlice) AddAsTitleValue(
	title string,
	value interface{},
) *SimpleSlice {
	it.Items = append(
		it.Items,
		fmt.Sprintf(constants.TitleValueFormat, title, value))

	return it
}

// AddAsCurlyTitleWrap
//
//  Adds Title: {value} (constants.CurlyTitleWrapFormat)
func (it *SimpleSlice) AddAsCurlyTitleWrap(
	title string,
	value interface{},
) *SimpleSlice {
	it.Items = append(
		it.Items,
		fmt.Sprintf(constants.CurlyTitleWrapFormat, title, value))

	return it
}

// AddAsCurlyTitleWrapIf
//
//  Adds Title: {value} (constants.CurlyTitleWrapFormat)
func (it *SimpleSlice) AddAsCurlyTitleWrapIf(
	isAppend bool,
	title string,
	value interface{},
) *SimpleSlice {
	if !isAppend {
		return it
	}

	it.Items = append(
		it.Items,
		fmt.Sprintf(constants.CurlyTitleWrapFormat, title, value))

	return it
}

// AddAsTitleValueIf
//
//  Skips if append is false
//  Adds Title : value (constants.TitleValueFormat)
func (it *SimpleSlice) AddAsTitleValueIf(
	isAppend bool,
	title string,
	value interface{},
) *SimpleSlice {
	if !isAppend {
		return it
	}

	it.Items = append(
		it.Items,
		fmt.Sprintf(constants.TitleValueFormat, title, value))

	return it
}

func (it *SimpleSlice) InsertAt(index int, item string) *SimpleSlice {
	it.Items = append(it.Items[:index+1], it.Items[index:]...)
	it.Items[index] = item

	return it
}

func (it *SimpleSlice) AddStruct(
	isIncludeFieldName bool,
	anyStruct interface{},
) *SimpleSlice {
	if anyStruct == nil {
		return it
	}

	val := AnyToString(
		isIncludeFieldName,
		anyStruct)

	return it.Add(val)
}

func (it *SimpleSlice) AddPointer(
	isIncludeFieldName bool,
	anyPtr interface{},
) *SimpleSlice {
	if anyPtr == nil {
		return it
	}

	val := AnyToString(
		isIncludeFieldName,
		anyPtr)

	return it.Add(val)
}

func (it *SimpleSlice) AddsIf(
	isAdd bool,
	items ...string,
) *SimpleSlice {
	if !isAdd {
		return it
	}

	return it.Adds(items...)
}

func (it *SimpleSlice) AddError(err error) *SimpleSlice {
	if err != nil {
		return it.Add(err.Error())
	}

	return it
}

func (it *SimpleSlice) AsDefaultError() error {
	return it.AsError(constants.NewLineUnix)
}

func (it *SimpleSlice) AsError(joiner string) error {
	if it == nil || it.Length() == 0 {
		return nil
	}

	errStr := strings.Join(
		it.Items,
		joiner)

	return errors.New(errStr)
}

func (it *SimpleSlice) FirstDynamic() interface{} {
	return it.Items[0]
}

func (it *SimpleSlice) First() string {
	return it.Items[0]
}

func (it *SimpleSlice) LastDynamic() interface{} {
	return it.Items[it.LastIndex()]
}

func (it *SimpleSlice) Last() string {
	return it.Items[it.LastIndex()]
}

func (it *SimpleSlice) FirstOrDefaultDynamic() interface{} {
	return it.FirstOrDefault()
}

func (it *SimpleSlice) FirstOrDefault() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.First()
}

func (it *SimpleSlice) LastOrDefaultDynamic() interface{} {
	return it.LastOrDefault()
}

func (it *SimpleSlice) LastOrDefault() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.Last()
}

func (it *SimpleSlice) SkipDynamic(skippingItemsCount int) interface{} {
	return it.Items[skippingItemsCount:]
}

func (it *SimpleSlice) Skip(skippingItemsCount int) []string {
	return it.Items[skippingItemsCount:]
}

func (it *SimpleSlice) TakeDynamic(takeDynamicItems int) interface{} {
	return it.Items[:takeDynamicItems]
}

func (it *SimpleSlice) Take(takeDynamicItems int) []string {
	return it.Items[:takeDynamicItems]
}

func (it *SimpleSlice) LimitDynamic(limit int) interface{} {
	return it.Take(limit)
}

func (it *SimpleSlice) Limit(limit int) []string {
	return it.Take(limit)
}

func (it *SimpleSlice) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *SimpleSlice) Count() int {
	return it.Length()
}

func (it *SimpleSlice) IsEmpty() bool {
	return it == nil || it.Length() == 0
}

func (it *SimpleSlice) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *SimpleSlice) LastIndex() int {
	return it.Length() - 1
}

func (it *SimpleSlice) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *SimpleSlice) Strings() []string {
	return it.Items
}

func (it *SimpleSlice) Hashset() *Hashset {
	return New.Hashset.StringsPtr(&it.Items)
}

func (it *SimpleSlice) Join(joiner string) string {
	return strings.Join(it.Items, joiner)
}

func (it *SimpleSlice) JoinLine() string {
	return strings.Join(it.Items, constants.NewLineUnix)
}

func (it *SimpleSlice) JoinSpace() string {
	return strings.Join(it.Items, constants.Space)
}

func (it *SimpleSlice) JoinComma() string {
	return strings.Join(it.Items, constants.Comma)
}

func (it *SimpleSlice) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it *SimpleSlice) JoinCsvLine() string {
	return strings.Join(it.CsvStrings(), constants.CommaUnixNewLine)
}

func (it *SimpleSlice) EachItemSplitBy(splitBy string) (splitItemsOnly []string) {
	slice := make([]string, 0, it.Length()*constants.Capacity3)

	for _, item := range it.Items {
		splitItems := strings.Split(item, splitBy)

		slice = append(slice, splitItems...)
	}

	return slice
}

func (it *SimpleSlice) PrependJoin(
	joiner string,
	prependItems ...string,
) string {
	prependSlice := &SimpleSlice{
		Items: prependItems,
	}

	return prependSlice.
		ConcatNew(it.Items...).
		Join(joiner)
}

func (it *SimpleSlice) AppendJoin(
	joiner string,
	appendItems ...string,
) string {
	return it.
		ConcatNew(appendItems...).
		Join(joiner)
}

func (it *SimpleSlice) PrependAppend(
	prependItems, appendItems []string,
) *SimpleSlice {
	if len(prependItems) > 0 {
		it.Items = append(prependItems, it.Items...)
	}

	if len(appendItems) > 0 {
		it.Items = append(it.Items, appendItems...)
	}

	return it
}

func (it *SimpleSlice) IsEqual(another *SimpleSlice) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.Length() != another.Length() {
		return false
	}

	if it.Length() == 0 {
		return true
	}

	return it.IsEqualLines(another.Items)
}

// IsEqualUnorderedLines
//
// Will sort both the list,
// output will contain sorted lines
func (it *SimpleSlice) IsEqualUnorderedLines(lines []string) bool {
	if it == nil && lines == nil {
		return true
	}

	if it == nil || lines == nil {
		return false
	}

	if it.Length() != len(lines) {
		return false
	}

	if it.Length() == 0 {
		return true
	}

	sort.Strings(it.Items)
	sort.Strings(lines)

	for i, item := range it.Items {
		anotherItem := lines[i]

		if item != anotherItem {
			return false
		}
	}

	return true
}

// IsEqualUnorderedLinesClone
//
// Will sort both the list,
// output will contain sorted lines
func (it *SimpleSlice) IsEqualUnorderedLinesClone(lines []string) bool {
	if it == nil && lines == nil {
		return true
	}

	if it == nil || lines == nil {
		return false
	}

	if it.Length() != len(lines) {
		return false
	}

	if it.Length() == 0 {
		return true
	}

	clonedLeftSlice := it.Clone(true)
	clonedRightSlice := New.SimpleSlice.Direct(true, lines)

	sort.Strings(clonedLeftSlice.Items)
	sort.Strings(clonedRightSlice.Items)

	for i, item := range clonedLeftSlice.Items {
		anotherItem := clonedRightSlice.Items[i]

		if item != anotherItem {
			return false
		}
	}

	go func() {
		clonedLeftSlice.Dispose()
		clonedRightSlice.Dispose()
	}()

	return true
}

func (it *SimpleSlice) IsEqualLines(lines []string) bool {
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

func (it *SimpleSlice) Collection(isClone bool) *Collection {
	return New.Collection.StringsOptions(
		isClone,
		it.Items,
	)
}

func (it SimpleSlice) NonPtr() SimpleSlice {
	return it
}

func (it *SimpleSlice) Ptr() *SimpleSlice {
	return it
}

func (it *SimpleSlice) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return strings.Join(
		it.Items,
		constants.NewLineUnix)
}

func (it *SimpleSlice) ConcatNewSimpleSlices(items ...*SimpleSlice) *SimpleSlice {
	items2 := append(
		items,
		it)
	length := AllIndividualsLengthOfSimpleSlices(items2...)
	slice := make(
		[]string,
		0,
		length)

	slice = append(slice, it.Items...)

	for _, simpleSlice := range items {
		slice = append(slice, simpleSlice.Items...)
	}

	return &SimpleSlice{Items: slice}
}

func (it *SimpleSlice) ConcatNewStrings(items ...string) []string {
	if it == nil {
		return CloneSlice(items)
	}

	slice := make(
		[]string,
		0,
		it.Length()+len(items))

	slice = append(slice, it.Items...)
	slice = append(slice, items...)

	return slice
}

func (it *SimpleSlice) ConcatNew(items ...string) *SimpleSlice {
	concatNew := it.ConcatNewStrings(items...)

	return &SimpleSlice{
		concatNew,
	}
}

func (it *SimpleSlice) ToCollection(isClone bool) *Collection {
	return New.Collection.StringsOptions(isClone, it.Items)
}

func (it *SimpleSlice) CsvStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	newSlice := make([]string, it.Length())

	for i, item := range it.Items {
		newSlice[i] = fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			item)
	}

	return newSlice
}

func (it *SimpleSlice) JoinCsvString(joiner string) string {
	if it.IsEmpty() {
		return ""
	}

	newSlice := it.CsvStrings()

	return strings.Join(newSlice, joiner)
}

func (it *SimpleSlice) JoinWith(
	joiner string,
) string {
	if it.IsEmpty() {
		return ""
	}

	return joiner + strings.Join(it.Items, joiner)
}

func (it *SimpleSlice) JsonModel() []string {
	return it.Items
}

func (it *SimpleSlice) Sort() *SimpleSlice {
	sort.Strings(it.Items)

	return it
}

func (it *SimpleSlice) Reverse() *SimpleSlice {
	length := it.Length()

	if length <= 1 {
		return it
	}

	if length == 2 {
		it.Items[0], it.Items[1] = it.Items[1], it.Items[0]

		return it
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		first := it.Items[i]
		it.Items[i] = it.Items[lastIndex-i]
		it.Items[lastIndex-i] = first
	}

	return it
}

func (it *SimpleSlice) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *SimpleSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *SimpleSlice) UnmarshalJSON(
	data []byte,
) error {
	var dataModel []string
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.Items = dataModel
	}

	return err
}

func (it SimpleSlice) Json() corejson.Result {
	return corejson.New(it)
}

func (it SimpleSlice) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *SimpleSlice) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*SimpleSlice, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return Empty.SimpleSlice(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *SimpleSlice) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *SimpleSlice {
	parsedResult, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return parsedResult
}

func (it SimpleSlice) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it SimpleSlice) AsJsoner() corejson.Jsoner {
	return &it
}

func (it SimpleSlice) ToPtr() *SimpleSlice {
	return &it
}

func (it SimpleSlice) ToNonPtr() SimpleSlice {
	return it
}

func (it *SimpleSlice) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it SimpleSlice) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return &it
}

func (it SimpleSlice) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}

func (it *SimpleSlice) Clear() *SimpleSlice {
	if it == nil {
		return nil
	}

	it.Items = it.Items[:0]

	return it
}

func (it *SimpleSlice) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it SimpleSlice) Clone(isDeepClone bool) SimpleSlice {
	return SimpleSlice{
		Items: CloneSliceIf(
			isDeepClone,
			it.Items...),
	}
}

func (it *SimpleSlice) ClonePtr(isDeepClone bool) *SimpleSlice {
	if it == nil {
		return nil
	}

	cloned := it.Clone(
		isDeepClone)

	return &cloned
}

func (it SimpleSlice) DeepClone() *SimpleSlice {
	return it.ClonePtr(true)
}

func (it SimpleSlice) ShadowClone() *SimpleSlice {
	return it.ClonePtr(false)
}

func (it SimpleSlice) IsDistinctEqualRaw(rightLines ...string) bool {
	return it.Hashset().IsEqualsPtr(New.Hashset.Strings(rightLines))
}

func (it SimpleSlice) IsDistinctEqual(rightSlice *SimpleSlice) bool {
	return it.Hashset().IsEqualsPtr(rightSlice.Hashset())
}

// IsUnorderedEqualRaw
//
//  sort both and then compare, if length are not equal then mismatch
//
//  isClone:
//      Don't mutate, create copy and then sort and then returns the final result.
func (it SimpleSlice) IsUnorderedEqualRaw(
	isClone bool,
	rightLines ...string,
) bool {
	if it.Length() != len(rightLines) {
		return false
	}

	if it.Length() == 0 {
		// no checking required
		return true
	}

	if isClone {
		leftSored := it.DeepClone()
		leftSored.Sort()

		rightSort := New.SimpleSlice.Direct(true, rightLines)
		rightSort.Sort()

		return leftSored.IsEqual(rightSort)
	}

	it.Sort()
	rightSort := New.SimpleSlice.Direct(
		false,
		rightLines)
	rightSort.Sort()

	return it.IsEqual(rightSort)
}

// IsUnorderedEqual
//
//  sort both and then compare, if length are not equal then mismatch
//
//  isClone:
//      Don't mutate, create copy and then sort and then returns the final result.
func (it SimpleSlice) IsUnorderedEqual(
	isClone bool,
	rightSlice *SimpleSlice,
) bool {
	if it.IsEmpty() && rightSlice.IsEmpty() {
		return true
	}

	if rightSlice == nil {
		// is nil left is not empty
		return false
	}

	return it.IsUnorderedEqualRaw(
		isClone,
		rightSlice.Items...)
}

// IsEqualByFunc
//
//  checks comparison by the given function.
func (it SimpleSlice) IsEqualByFunc(
	isMatchCheckerFunc func(index int, left, right string) (isMatch bool),
	rightLines ...string,
) bool {
	if it.Length() != len(rightLines) {
		return false
	}

	if it.Length() == 0 {
		return true
	}

	for i, rightLine := range rightLines {
		leftLine := it.Items[i]

		if !isMatchCheckerFunc(i, leftLine, rightLine) {
			return false
		}
	}

	return true
}

// IsEqualByFuncLinesSplit
//
//  first splits the line and then takes
//  lines and process same as EqualByFunc
func (it SimpleSlice) IsEqualByFuncLinesSplit(
	isTrim bool,
	splitter string,
	rightLine string,
	isMatchCheckerFunc func(index int, left, right string) (isMatch bool),
) bool {
	rightLines := strings.Split(rightLine, splitter)

	if len(rightLines) != it.Length() {
		return false
	}

	if it.IsEmpty() {
		return true
	}

	for i, curLeftLine := range it.Items {
		curRightLine := rightLines[i]

		if isTrim {
			curLeftLine = strings.TrimSpace(curLeftLine)
			curRightLine = strings.TrimSpace(curRightLine)
		}

		if !isMatchCheckerFunc(i, curLeftLine, curRightLine) {
			return false
		}
	}

	return true
}
