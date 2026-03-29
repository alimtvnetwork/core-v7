package corestr

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

type SimpleSlice []string

func (it *SimpleSlice) Add(
	item string,
) *SimpleSlice {
	*it = append(*it, item)

	return it
}

func (it *SimpleSlice) AddSplit(
	item string,
	sep string,
) *SimpleSlice {
	lines := strings.Split(item, sep)

	for _, line := range lines {
		*it = append(*it, line)
	}

	return it
}

func (it *SimpleSlice) AddIf(
	isAdd bool,
	item string,
) *SimpleSlice {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	*it = append(*it, item)

	return it
}

func (it *SimpleSlice) Adds(
	items ...string,
) *SimpleSlice {
	if len(items) == 0 {
		return it
	}

	*it = append(*it, items...)

	return it
}

func (it *SimpleSlice) Append(
	items ...string,
) *SimpleSlice {
	if len(items) == 0 {
		return it
	}

	*it = append(*it, items...)

	return it
}

// AppendFmt
//
//	Skips empty format + values
func (it *SimpleSlice) AppendFmt(
	format string,
	v ...any,
) *SimpleSlice {
	if format == "" && len(v) == 0 {
		return it
	}

	*it = append(
		*it,
		fmt.Sprintf(format, v...),
	)

	return it
}

// AppendFmtIf
//
//	Skips empty format + values
func (it *SimpleSlice) AppendFmtIf(
	isAppend bool,
	format string,
	v ...any,
) *SimpleSlice {
	if !isAppend || format == "" && len(v) == 0 {
		return it
	}

	*it = append(
		*it,
		fmt.Sprintf(format, v...),
	)

	return it
}

// AddAsTitleValue
//
//	Adds Title : value (constants.TitleValueFormat)
func (it *SimpleSlice) AddAsTitleValue(
	title string,
	value any,
) *SimpleSlice {
	*it = append(
		*it,
		fmt.Sprintf(constants.TitleValueFormat, title, value),
	)

	return it
}

// AddAsCurlyTitleWrap
//
//	Adds Title: {value} (constants.CurlyTitleWrapFormat)
func (it *SimpleSlice) AddAsCurlyTitleWrap(
	title string,
	value any,
) *SimpleSlice {
	*it = append(
		*it,
		fmt.Sprintf(constants.CurlyTitleWrapFormat, title, value),
	)

	return it
}

// AddAsCurlyTitleWrapIf
//
//	Adds Title: {value} (constants.CurlyTitleWrapFormat)
func (it *SimpleSlice) AddAsCurlyTitleWrapIf(
	isAppend bool,
	title string,
	value any,
) *SimpleSlice {
	isSkip := !isAppend

	if isSkip {
		return it
	}

	*it = append(
		*it,
		fmt.Sprintf(constants.CurlyTitleWrapFormat, title, value),
	)

	return it
}

// AddAsTitleValueIf
//
//	Skips if append is false
//	Adds Title : value (constants.TitleValueFormat)
func (it *SimpleSlice) AddAsTitleValueIf(
	isAppend bool,
	title string,
	value any,
) *SimpleSlice {
	isSkip := !isAppend

	if isSkip {
		return it
	}

	*it = append(
		*it,
		fmt.Sprintf(constants.TitleValueFormat, title, value),
	)

	return it
}

func (it *SimpleSlice) InsertAt(
	index int,
	item string,
) *SimpleSlice {
	if index < 0 || index > it.Length() {
		return it
	}

	s := *it
	s = append(s, "")
	copy(s[index+1:], s[index:])
	s[index] = item
	*it = s

	return it
}

func (it *SimpleSlice) AddStruct(
	isIncludeFieldName bool,
	anyStruct any,
) *SimpleSlice {
	if anyStruct == nil {
		return it
	}

	val := AnyToString(
		isIncludeFieldName,
		anyStruct,
	)

	return it.Add(val)
}

func (it *SimpleSlice) AddPointer(
	isIncludeFieldName bool,
	anyPtr any,
) *SimpleSlice {
	if anyPtr == nil {
		return it
	}

	val := AnyToString(
		isIncludeFieldName,
		anyPtr,
	)

	return it.Add(val)
}

func (it *SimpleSlice) AddsIf(
	isAdd bool,
	items ...string,
) *SimpleSlice {
	isSkip := !isAdd

	if isSkip {
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
		*it,
		joiner,
	)

	return errors.New(errStr)
}

func (it *SimpleSlice) FirstDynamic() any {
	return (*it)[0]
}

func (it *SimpleSlice) First() string {
	return (*it)[0]
}

func (it *SimpleSlice) LastDynamic() any {
	return (*it)[it.LastIndex()]
}

func (it *SimpleSlice) Last() string {
	return (*it)[it.LastIndex()]
}

func (it *SimpleSlice) FirstOrDefaultDynamic() any {
	return it.FirstOrDefault()
}

func (it *SimpleSlice) FirstOrDefault() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.First()
}

func (it *SimpleSlice) LastOrDefaultDynamic() any {
	return it.LastOrDefault()
}

func (it *SimpleSlice) LastOrDefault() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.Last()
}

func (it *SimpleSlice) SkipDynamic(skippingItemsCount int) any {
	if it == nil || skippingItemsCount < 0 || skippingItemsCount >= it.Length() {
		return []string{}
	}

	return []string((*it)[skippingItemsCount:])
}

func (it *SimpleSlice) Skip(skippingItemsCount int) []string {
	if skippingItemsCount < 0 || skippingItemsCount >= it.Length() {
		return []string{}
	}

	return (*it)[skippingItemsCount:]
}

func (it *SimpleSlice) TakeDynamic(takeDynamicItems int) any {
	if it == nil || takeDynamicItems < 0 {
		return []string{}
	}

	if takeDynamicItems >= it.Length() {
		return []string(*it)
	}

	return []string((*it)[:takeDynamicItems])
}

func (it *SimpleSlice) Take(takeDynamicItems int) []string {
	if takeDynamicItems < 0 || takeDynamicItems >= it.Length() {
		return *it
	}

	return (*it)[:takeDynamicItems]
}

func (it *SimpleSlice) LimitDynamic(limit int) any {
	return it.Take(limit)
}

func (it *SimpleSlice) Limit(limit int) []string {
	return it.Take(limit)
}

func (it *SimpleSlice) Length() int {
	if it == nil {
		return 0
	}

	return len(*it)
}

func (it *SimpleSlice) Count() int {
	return it.Length()
}

func (it *SimpleSlice) CountFunc(
	counterFunc func(index int, item string) (isCount bool),
) int {
	if it.IsEmpty() {
		return 0
	}

	counter := 0
	for i, item := range *it {
		isCount := counterFunc(i, item)

		if isCount {
			counter++
		}
	}

	return counter
}

func (it *SimpleSlice) IsEmpty() bool {
	return it == nil || it.Length() == 0
}

func (it *SimpleSlice) IsContains(item string) bool {
	if it.IsEmpty() {
		return false
	}

	slice := *it

	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}

func (it *SimpleSlice) IsContainsFunc(
	searchingItem string,
	isSearchFunc func(item, searching string) bool,
) bool {
	if it.IsEmpty() {
		return false
	}

	slice := *it

	for _, s := range slice {
		if isSearchFunc(s, searchingItem) {
			return true
		}
	}

	return false
}

func (it *SimpleSlice) IndexOfFunc(
	searchingItem string,
	isSearchFunc func(item, searching string) bool,
) int {
	if it.IsEmpty() {
		return constants.InvalidIndex
	}

	slice := *it

	for i, s := range slice {
		if isSearchFunc(s, searchingItem) {
			return i
		}
	}

	return constants.InvalidIndex
}

func (it *SimpleSlice) IndexOf(item string) int {
	if it.IsEmpty() {
		return constants.InvalidIndex
	}

	slice := *it

	for i, s := range slice {
		if s == item {
			return i
		}
	}

	return constants.InvalidIndex
}

func (it *SimpleSlice) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *SimpleSlice) LastIndex() int {
	return it.Length() - 1
}

func (it *SimpleSlice) HasIndex(index int) bool {
	return index >= 0 && it.LastIndex() >= index
}

func (it *SimpleSlice) Strings() []string {
	return *it
}

func (it *SimpleSlice) List() []string {
	return *it
}

func (it *SimpleSlice) WrapDoubleQuote() *SimpleSlice {
	return it.Transpile(StringUtils.WrapDouble)
}

func (it *SimpleSlice) WrapSingleQuote() *SimpleSlice {
	return it.Transpile(StringUtils.WrapSingle)
}

func (it *SimpleSlice) WrapTildaQuote() *SimpleSlice {
	return it.Transpile(StringUtils.WrapTilda)
}

func (it *SimpleSlice) WrapDoubleQuoteIfMissing() *SimpleSlice {
	return it.Transpile(StringUtils.WrapDoubleIfMissing)
}

func (it *SimpleSlice) WrapSingleQuoteIfMissing() *SimpleSlice {
	return it.Transpile(StringUtils.WrapSingleIfMissing)
}

func (it *SimpleSlice) Transpile(
	fmtFunc func(s string) string,
) *SimpleSlice {
	if it.IsEmpty() {
		return it.emptySimpleSlice()
	}

	newSlice := make([]string, it.Length())

	for i, s := range *it {
		newSlice[i] = fmtFunc(s)
	}

	return it.convertToSimpleSlice(newSlice)
}

func (it *SimpleSlice) emptySimpleSlice() *SimpleSlice {
	list := SimpleSlice([]string{})

	return &list
}

func (it *SimpleSlice) convertToSimpleSlice(list []string) *SimpleSlice {
	output := SimpleSlice(list)

	return &output
}

func (it *SimpleSlice) TranspileJoin(
	fmtFunc func(s string) string,
	sep string,
) string {
	toTranspile := it.Transpile(fmtFunc)

	return toTranspile.Join(sep)
}

func (it *SimpleSlice) Hashset() *Hashset {
	return New.Hashset.Strings(*it)
}

func (it *SimpleSlice) Join(joiner string) string {
	if it.IsEmpty() {
		return ""
	}

	return strings.Join(*it, joiner)
}

func (it *SimpleSlice) JoinLine() string {
	if it.IsEmpty() {
		return ""
	}

	return strings.Join(
		*it,
		constants.DefaultLine,
	)
}

// JoinLineEofLine
//
//	contains new line at the end of the file
func (it *SimpleSlice) JoinLineEofLine() string {
	if it.IsEmpty() {
		return ""
	}

	joined := strings.Join(
		*it,
		constants.DefaultLine,
	)

	if strings.HasSuffix(joined, constants.DefaultLine) {
		// already contains

		return joined
	}

	return joined + constants.DefaultLine
}

func (it *SimpleSlice) JoinSpace() string {
	return strings.Join(*it, constants.Space)
}

func (it *SimpleSlice) JoinComma() string {
	return strings.Join(*it, constants.Comma)
}

func (it *SimpleSlice) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it *SimpleSlice) JoinCsvLine() string {
	return strings.Join(it.CsvStrings(), constants.CommaUnixNewLine)
}

func (it *SimpleSlice) EachItemSplitBy(splitBy string) (splitItemsOnly *SimpleSlice) {
	slice := make([]string, 0, it.Length()*constants.Capacity3)

	for _, item := range *it {
		splitItems := strings.Split(item, splitBy)

		slice = append(slice, splitItems...)
	}

	return it.convertToSimpleSlice(slice)
}

func (it *SimpleSlice) PrependJoin(
	joiner string,
	prependItems ...string,
) string {
	prependSlice := SimpleSlice(prependItems)

	return prependSlice.
		ConcatNew(*it...).
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
		*it = append(prependItems, *it...)
	}

	if len(appendItems) > 0 {
		*it = append(*it, appendItems...)
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

	return it.IsEqualLines(*another)
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

	sort.Strings(*it)
	sort.Strings(lines)

	for i, item := range *it {
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

	sort.Strings(clonedLeftSlice)
	sort.Strings(*clonedRightSlice)

	for i, item := range clonedLeftSlice {
		anotherItem := (*clonedRightSlice)[i]

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

	for i, item := range *it {
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
		*it,
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
		*it,
		constants.NewLineUnix,
	)
}

func (it *SimpleSlice) ConcatNewSimpleSlices(items ...*SimpleSlice) *SimpleSlice {
	nextItems := append(
		items,
		it,
	)
	length := AllIndividualsLengthOfSimpleSlices(nextItems...)
	slice := make(
		[]string,
		0,
		length,
	)

	slice = append(slice, *it...)

	for _, simpleSlice := range items {
		slice = append(slice, *simpleSlice...)
	}

	newSlice := SimpleSlice(slice)

	return &newSlice
}

func (it *SimpleSlice) ConcatNewStrings(items ...string) []string {
	if it == nil {
		return CloneSlice(items)
	}

	slice := make(
		[]string,
		0,
		it.Length()+len(items),
	)

	slice = append(slice, *it...)
	slice = append(slice, items...)

	return slice
}

func (it *SimpleSlice) ConcatNew(items ...string) *SimpleSlice {
	concatNew := it.ConcatNewStrings(items...)
	slice := SimpleSlice(concatNew)

	return &slice
}

func (it *SimpleSlice) ToCollection(isClone bool) *Collection {
	return New.Collection.StringsOptions(isClone, *it)
}

func (it *SimpleSlice) CsvStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	newSlice := make([]string, it.Length())
	slice := *it

	for i, item := range slice {
		newSlice[i] = fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			item,
		)
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

	return joiner + strings.Join(*it, joiner)
}

func (it SimpleSlice) JsonModel() []string {
	return it
}

func (it *SimpleSlice) Sort() *SimpleSlice {
	sort.Strings(*it)

	return it
}

func (it *SimpleSlice) Reverse() *SimpleSlice {
	length := it.Length()

	if length <= 1 {
		return it
	}

	if length == 2 {
		(*it)[0], (*it)[1] = (*it)[1], (*it)[0]

		return it
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		first := (*it)[i]
		(*it)[i] = (*it)[lastIndex-i]
		(*it)[lastIndex-i] = first
	}

	return it
}

func (it SimpleSlice) JsonModelAny() any {
	return it.JsonModel()
}

func (it SimpleSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *SimpleSlice) UnmarshalJSON(
	rawBytes []byte,
) error {
	var dataModel []string
	err := json.Unmarshal(rawBytes, &dataModel)

	if err == nil {
		*it = dataModel
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

	*it = (*it)[:0]

	return it
}

func (it *SimpleSlice) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	*it = nil
}

func (it SimpleSlice) Clone(isDeepClone bool) SimpleSlice {
	return CloneSliceIf(
		isDeepClone,
		it...,
	)
}

func (it *SimpleSlice) ClonePtr(isDeepClone bool) *SimpleSlice {
	if it == nil {
		return nil
	}

	cloned := it.Clone(
		isDeepClone,
	)

	return &cloned
}

func (it SimpleSlice) DeepClone() *SimpleSlice {
	return it.ClonePtr(true)
}

func (it SimpleSlice) ShadowClone() *SimpleSlice {
	return it.ClonePtr(false)
}

func (it SimpleSlice) IsDistinctEqualRaw(rightLines ...string) bool {
	return it.Hashset().IsEquals(New.Hashset.Strings(rightLines))
}

func (it SimpleSlice) IsDistinctEqual(rightSlice *SimpleSlice) bool {
	return it.Hashset().IsEquals(rightSlice.Hashset())
}

// IsUnorderedEqualRaw
//
//	sort both and then compare, if length are not equal then mismatch
//
//	isClone:
//	    Don't mutate, create copy and then sort and then returns the final result.
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
		rightLines,
	)
	rightSort.Sort()

	return it.IsEqual(rightSort)
}

// IsUnorderedEqual
//
//	sort both and then compare, if length are not equal then mismatch
//
//	isClone:
//	    Don't mutate, create copy and then sort and then returns the final result.
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
		*rightSlice...,
	)
}

// IsEqualByFunc
//
//	checks comparison by the given function.
func (it *SimpleSlice) IsEqualByFunc(
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
		leftLine := (*it)[i]

		isMismatch := !isMatchCheckerFunc(i, leftLine, rightLine)

		if isMismatch {
			return false
		}
	}

	return true
}

// IsEqualByFuncLinesSplit
//
//	first splits the line and then takes
//	lines and process same as EqualByFunc
func (it *SimpleSlice) IsEqualByFuncLinesSplit(
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

	for i, curLeftLine := range *it {
		curRightLine := rightLines[i]

		if isTrim {
			curLeftLine = strings.TrimSpace(curLeftLine)
			curRightLine = strings.TrimSpace(curRightLine)
		}

		isMismatch := !isMatchCheckerFunc(i, curLeftLine, curRightLine)

		if isMismatch {
			return false
		}
	}

	return true
}

func (it *SimpleSlice) DistinctDiffRaw(
	rightLines ...string,
) []string {
	if it == nil && rightLines == nil {
		return []string{}
	}

	if it == nil && rightLines != nil {
		return rightLines
	}

	if it != nil && rightLines == nil {
		return *it
	}

	return New.
		Hashset.
		Strings(*it).
		DistinctDiffLinesRaw(rightLines...)
}

func (it *SimpleSlice) AddedRemovedLinesDiff(
	rightLines ...string,
) (addedLines, removedLines []string) {
	if it == nil && rightLines == nil {
		return addedLines, removedLines
	}

	oldDomainsHashSet := New.Hashset.Strings(it.SafeStrings())
	newDomainsHashSet := New.Hashset.Strings(rightLines)

	addedDomainsPtr := newDomainsHashSet.
		GetAllExceptHashset(oldDomainsHashSet)
	deletedDomainsPtr := oldDomainsHashSet.
		GetAllExceptHashset(newDomainsHashSet)

	if addedDomainsPtr != nil {
		addedLines = addedDomainsPtr
	}

	if deletedDomainsPtr != nil {
		removedLines = deletedDomainsPtr
	}

	return addedLines, removedLines

}

func (it *SimpleSlice) RemoveIndexes(
	indexes ...int,
) (newSlice *SimpleSlice, err error) {
	if it.IsEmpty() {
		return Empty.SimpleSlice(), fmt.Errorf(
			"cannot remove indexes(%+v) from an empty slice",
			indexes,
		)
	}

	length := it.Length() - len(indexes)
	if length <= 0 {
		length = 0
	}

	newSlice = New.SimpleSlice.Cap(length)

	indexMap := make(map[int]bool, len(indexes))
	for _, index := range indexes {
		indexMap[index] = true
	}

	for i, s := range *it {
		if indexMap[i] {
			delete(indexMap, i)

			continue
		}

		newSlice.Add(s)
	}

	if len(indexMap) > 0 {
		err = fmt.Errorf(
			"cannot remove indexes(%+v) from the slice (%+v)",
			indexMap,
			*it,
		)
	}

	return newSlice, err
}

func (it *SimpleSlice) DistinctDiff(
	rightSlice *SimpleSlice,
) []string {
	if it == nil && rightSlice == nil {
		return []string{}
	}

	if it == nil && rightSlice != nil {
		return *rightSlice
	}

	if it != nil && rightSlice == nil {
		return *it
	}

	return New.
		Hashset.
		Strings(*it).
		DistinctDiffLinesRaw(*rightSlice...)
}

func (it *SimpleSlice) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *SimpleSlice) Deserialize(toPtr any) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}

func (it *SimpleSlice) SafeStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return *it
}
