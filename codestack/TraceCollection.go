package codestack

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/defaultcapacity"
	"github.com/alimtvnetwork/core/errcore"
)

type TraceCollection struct {
	Items []Trace `json:"Items,omitempty"`
}

func (it TraceCollection) StackTracesJsonResult() *corejson.Result {
	return it.JsonPtr()
}

func (it TraceCollection) StackTraces() string {
	return it.CodeStacksString()
}

func (it TraceCollection) NewStackTraces(stackSkip int) string {
	return New.
		StackTrace.
		DefaultCount(stackSkip + defaultInternalSkip).
		CodeStacksString()
}

func (it TraceCollection) NewDefaultStackTraces() string {
	return it.NewStackTraces(defaultInternalSkip)
}

func (it TraceCollection) NewStackTracesJsonResult(stackSkip int) *corejson.Result {
	return New.
		StackTrace.
		DefaultCount(defaultInternalSkip + stackSkip).
		JsonPtr()
}

func (it TraceCollection) NewDefaultStackTracesJsonResult() *corejson.Result {
	return New.
		StackTrace.
		DefaultCount(defaultInternalSkip).
		JsonPtr()
}

func (it *TraceCollection) Add(
	trace Trace,
) *TraceCollection {
	it.Items = append(it.Items, trace)

	return it
}

func (it *TraceCollection) AddsUsingSkip(
	isSkipInvalid,
	isBreakOnceInvalid bool,
	startSkipIndex, // will add defaultInternalSkip(2) to skip its own stack trace
	stackCount int,
) *TraceCollection {
	start := startSkipIndex +
		defaultInternalSkip

	for i := start; i < stackCount+start; i++ {
		trace := New.Create(i)
		isSkip := isSkipInvalid && trace.HasIssues()

		if isSkip && isBreakOnceInvalid {
			return it
		} else if isSkip {
			continue
		}

		if trace.IsSkippable {
			continue
		}

		it.Items = append(
			it.Items,
			trace,
		)
	}

	return it
}

func (it *TraceCollection) AddsUsingSkipDefault(
	startSkipIndex int, // will add defaultInternalSkip(2) to skip its own stack trace
) *TraceCollection {
	return it.AddsUsingSkip(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		DefaultStackCount,
	)
}

func (it *TraceCollection) AddsUsingSkipUsingFilter(
	isSkipInvalid,
	isBreakOnceInvalid bool,
	startSkipIndex, // will add defaultInternalSkip(2) to skip its own stack trace
	stackCount int,
	filterFunc FilterFunc,
) *TraceCollection {
	start := startSkipIndex +
		defaultInternalSkip

	for i := start; i < stackCount+start; i++ {
		trace := New.Create(i)
		isSkip := isSkipInvalid && trace.HasIssues()

		if isSkip && isBreakOnceInvalid {
			return it
		} else if isSkip {
			continue
		}

		if trace.IsSkippable {
			continue
		}

		isTake, isBreak := filterFunc(&trace)

		if isTake {
			it.Items = append(
				it.Items,
				trace,
			)
		}

		if isBreak {
			return it
		}
	}

	return it
}

func (it *TraceCollection) Adds(
	traces ...Trace,
) *TraceCollection {
	if len(traces) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		traces...,
	)

	return it
}

func (it *TraceCollection) AddsPtr(
	isSkipOnIssues bool,
	traces ...*Trace,
) *TraceCollection {
	if len(traces) == 0 {
		return it
	}

	for _, trace := range traces {
		if trace.IsNil() {
			continue
		}

		if isSkipOnIssues && trace.HasIssues() {
			continue
		}

		it.Items = append(
			it.Items,
			*trace,
		)
	}

	return it
}

func (it *TraceCollection) ConcatNew(
	additionalTraces ...Trace,
) *TraceCollection {
	cloned := it.Clone()

	return cloned.Adds(additionalTraces...)
}

func (it *TraceCollection) ConcatNewPtr(
	additionalTraces ...*Trace,
) *TraceCollection {
	cloned := it.Clone()

	return cloned.AddsPtr(
		true,
		additionalTraces...,
	)
}

func (it *TraceCollection) ConcatNewUsingSkipPlusCount(
	skipIndex,
	stackCount int,
) *TraceCollection {
	cloned := it.Clone()

	return cloned.AddsUsingSkip(
		true,
		true,
		skipIndex+defaultInternalSkip,
		stackCount,
	)
}

func (it *TraceCollection) ConcatNewUsingSkip(
	skipIndex int,
) *TraceCollection {
	cloned := it.Clone()

	return cloned.AddsUsingSkip(
		true,
		true,
		skipIndex+defaultInternalSkip,
		DefaultStackCount,
	)
}

func (it *TraceCollection) AddsIf(
	isAdd bool,
	traces ...Trace,
) *TraceCollection {
	if !isAdd {
		return it
	}

	return it.Adds(traces...)
}

func (it *TraceCollection) InsertAt(index int, item Trace) *TraceCollection {
	it.Items = append(it.Items[:index+1], it.Items[index:]...)
	it.Items[index] = item

	return it
}

func (it *TraceCollection) FirstDynamic() any {
	return it.Items[0]
}

func (it *TraceCollection) First() Trace {
	return it.Items[0]
}

func (it *TraceCollection) LastDynamic() any {
	return it.Items[it.LastIndex()]
}

func (it *TraceCollection) Last() Trace {
	return it.Items[it.LastIndex()]
}

func (it *TraceCollection) FirstOrDefaultDynamic() any {
	return it.FirstOrDefault()
}

func (it *TraceCollection) FirstOrDefault() *Trace {
	if it.IsEmpty() {
		return nil
	}

	first := it.First()

	return &first
}

func (it *TraceCollection) LastOrDefaultDynamic() any {
	return it.LastOrDefault()
}

func (it *TraceCollection) LastOrDefault() *Trace {
	if it.IsEmpty() {
		return nil
	}

	last := it.Last()

	return &last
}

func (it *TraceCollection) SkipDynamic(skippingItemsCount int) any {
	if skippingItemsCount >= len(it.Items) {
		return []Trace{}
	}

	return it.Items[skippingItemsCount:]
}

func (it *TraceCollection) Skip(skippingItemsCount int) []Trace {
	return it.Items[skippingItemsCount:]
}

func (it *TraceCollection) SkipCollection(skippingItemsCount int) *TraceCollection {
	return &TraceCollection{
		Items: it.Items[skippingItemsCount:],
	}
}

func (it *TraceCollection) TakeDynamic(takeDynamicItems int) any {
	return it.Items[:takeDynamicItems]
}

func (it *TraceCollection) Take(takeDynamicItems int) []Trace {
	return it.Items[:takeDynamicItems]
}

func (it *TraceCollection) TakeCollection(takeDynamicItems int) *TraceCollection {
	return &TraceCollection{
		Items: it.Items[:takeDynamicItems],
	}
}

func (it *TraceCollection) LimitCollection(limit int) *TraceCollection {
	return &TraceCollection{
		Items: it.Items[:limit],
	}
}

func (it *TraceCollection) SafeLimitCollection(limit int) *TraceCollection {
	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &TraceCollection{
		Items: it.Items[:limit],
	}
}

func (it *TraceCollection) LimitDynamic(limit int) any {
	return it.Take(limit)
}

func (it *TraceCollection) Limit(limit int) []Trace {
	return it.Take(limit)
}

// GetPagesSize returns the number of pages for the given page size.
// Returns 0 if eachPageSize is zero or negative.
func (it *TraceCollection) GetPagesSize(
	eachPageSize int,
) int {
	if eachPageSize <= 0 {
		return 0
	}

	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *TraceCollection) GetPagedCollection(
	eachPageSize int,
) []*TraceCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*TraceCollection{
			it,
		}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*TraceCollection, pagesPossibleCeiling)

	wg := sync.WaitGroup{}
	addPagedItemsFunc := func(oneBasedPageIndex int) {
		pagedCollection := it.GetSinglePageCollection(
			eachPageSize,
			oneBasedPageIndex,
		)

		collectionOfCollection[oneBasedPageIndex-1] = pagedCollection

		wg.Done()
	}

	wg.Add(pagesPossibleCeiling)
	for i := 1; i <= pagesPossibleCeiling; i++ {
		go addPagedItemsFunc(i)
	}

	wg.Wait()

	return collectionOfCollection
}

// GetSinglePageCollection PageIndex is one based index. Should be above or equal 1
func (it *TraceCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *TraceCollection {
	length := it.Length()

	if length < eachPageSize {
		return it
	}

	/**
	 * eachPageItems = 10
	 * pageIndex = 4
	 * skipItems = 10 * (4 - 1) = 30
	 */
	skipItems := eachPageSize * (pageIndex - 1)
	if skipItems < 0 {
		errcore.
			CannotBeNegativeIndexType.
			HandleUsingPanic(
				"pageIndex cannot be negative or zero.",
				pageIndex,
			)
	}

	if skipItems >= length {
		return New.traces.Using(false)
	}

	endingIndex := skipItems + eachPageSize

	if endingIndex > length {
		endingIndex = length
	}

	list := it.Items[skipItems:endingIndex]

	return New.traces.Using(false, list...)
}

func (it *TraceCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *TraceCollection) Count() int {
	return it.Length()
}

func (it *TraceCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *TraceCollection) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *TraceCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *TraceCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *TraceCollection) Strings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.String()
	}

	return list
}

func (it *TraceCollection) Filter(
	filterFunc FilterFunc,
) []Trace {
	list := make([]Trace, 0, it.Length())

	for _, item := range it.Items {
		isTake, isBreak := filterFunc(&item)

		if isTake {
			list = append(list, item)
		}

		if isBreak {
			return list
		}
	}

	return list
}

func (it *TraceCollection) FilterWithLimit(
	limit int,
	filterFunc FilterFunc,
) []Trace {
	length := defaultcapacity.MaxLimit(
		it.Length(),
		limit,
	)
	list := make(
		[]Trace,
		0,
		length,
	)

	collectedItems := 0
	for _, item := range it.Items {
		isTake, isBreak := filterFunc(&item)

		if isTake {
			list = append(list, item)
			collectedItems++
		}

		if isBreak {
			return list
		}

		if collectedItems >= length {
			return list
		}
	}

	return list
}

func (it *TraceCollection) FilterTraceCollection(
	filterFunc FilterFunc,
) *TraceCollection {
	list := it.Filter(filterFunc)

	traceCollection := New.traces.Using(
		false, list...,
	)

	return traceCollection
}

func (it *TraceCollection) FilterPackageNameTraceCollection(
	packageName string,
) *TraceCollection {
	return it.FilterTraceCollection(
		func(trace *Trace) (isTake, isBreak bool) {
			return trace.PackageName == packageName, false
		},
	)
}

func (it *TraceCollection) SkipFilterPackageNameTraceCollection(
	packageName string,
) *TraceCollection {
	return it.FilterTraceCollection(
		func(trace *Trace) (isTake, isBreak bool) {
			return trace.PackageName != packageName, false
		},
	)
}

func (it *TraceCollection) FilterMethodNameTraceCollection(
	methodName string,
) *TraceCollection {
	return it.FilterTraceCollection(
		func(trace *Trace) (isTake, isBreak bool) {
			return trace.MethodName == methodName, false
		},
	)
}

func (it *TraceCollection) SkipFilterMethodNameTraceCollection(
	methodName string,
) *TraceCollection {
	return it.FilterTraceCollection(
		func(trace *Trace) (isTake, isBreak bool) {
			return trace.MethodName != methodName, false
		},
	)
}

// FilterFullMethodNameTraceCollection
//
// fullMethodName := packageName.struct.methodName
func (it *TraceCollection) FilterFullMethodNameTraceCollection(
	fullMethodName string,
) *TraceCollection {
	return it.FilterTraceCollection(
		func(trace *Trace) (isTake, isBreak bool) {
			return trace.PackageMethodName == fullMethodName, false
		},
	)
}

// SkipFilterFullMethodNameTraceCollection
//
// fullMethodName := packageName.struct.methodName
func (it *TraceCollection) SkipFilterFullMethodNameTraceCollection(
	fullMethodName string,
) *TraceCollection {
	return it.FilterTraceCollection(
		func(trace *Trace) (isTake, isBreak bool) {
			return trace.PackageMethodName != fullMethodName, false
		},
	)
}

func (it *TraceCollection) SkipFilterFilenameTraceCollection(
	skipFilename string,
) *TraceCollection {
	return it.FilterTraceCollection(
		func(trace *Trace) (isTake, isBreak bool) {
			return trace.FilePath != skipFilename, false
		},
	)
}

func (it *TraceCollection) FileWithLines() []FileWithLine {
	list := make([]FileWithLine, it.Length())

	for i, item := range it.Items {
		list[i] = item.FileWithLine()
	}

	return list
}

func (it *TraceCollection) FileWithLinesStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.FileWithLineString()
	}

	return list
}

func (it *TraceCollection) StringsUsingFmt(formatter Formatter) []string {
	list := make([]string, it.Length())

	for i := range it.Items {
		list[i] = formatter(&it.Items[i])
	}

	return list
}

func (it *TraceCollection) JoinUsingFmt(formatter Formatter, joiner string) string {
	lines := it.StringsUsingFmt(formatter)

	return strings.Join(lines, joiner)
}

// ShortStrings
//
// Returns slice of Trace.ShortString() which gives
// "Method (LineNumber) -> FileFullPath:LineNumber"
// using format shortStringFormat "%s (%d) -> %s:%d"
//
// Format :
//   - https://prnt.sc/25ypcyc : "%s (%d) -> %s:%d"
//
// Example :
//   - Slice of []String{"Method (LineNumber) -> FileFullPath:LineNumber"}
func (it *TraceCollection) ShortStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.ShortString()
	}

	return list
}

// JoinShortStrings
//
// Returns a join slice of Trace.ShortString() from ShortStrings()
// which gives "Method (LineNumber) -> FileFullPath:LineNumber"
// using format shortStringFormat "%s (%d) -> %s:%d"
//
// Format :
//   - https://prnt.sc/25ypcyc : "%s (%d) -> %s:%d"
//
// Example :
//   - Compiled joined string of slice []String{"Method (LineNumber) -> FileFullPath:LineNumber"}
func (it TraceCollection) JoinShortStrings(joiner string) string {
	return strings.Join(it.ShortStrings(), joiner)
}

func (it *TraceCollection) Reverse() *TraceCollection {
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
		it.Items[i], it.Items[lastIndex-i] =
			it.Items[lastIndex-i], it.Items[i]
	}

	return it
}

func (it TraceCollection) JsonStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.JsonString()
	}

	return list
}

func (it TraceCollection) JoinFileWithLinesStrings(joiner string) string {
	return strings.Join(it.FileWithLinesStrings(), joiner)
}

func (it TraceCollection) JoinJsonStrings(joiner string) string {
	return strings.Join(it.JsonStrings(), joiner)
}

// CodeStacksString
//
// Returns a join "Code Stack :\n- JoinLinesWith(\n- )"
//
// Format :
//   - https://prnt.sc/25ypwem : "Code Stack :\n- JoinLinesWith(\n- )"
//
// Sample :
//   - "Code Stack :\n- JoinLinesWith(\n- )"
func (it TraceCollection) CodeStacksString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	toString := errcore.StackTracesCompiled(
		it.ShortStrings(),
	)

	return toString
}

func (it *TraceCollection) FileWithLinesString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	toString := errcore.StackTracesCompiled(
		it.FileWithLinesStrings(),
	)

	return toString
}

func (it TraceCollection) CodeStacksStringLimit(limit int) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	collection := it.SafeLimitCollection(limit)

	return collection.CodeStacksString()
}

func (it TraceCollection) Join(joiner string) string {
	return strings.Join(it.Strings(), joiner)
}

func (it TraceCollection) JoinLines() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}

func (it TraceCollection) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it TraceCollection) JoinCsvLine() string {
	return strings.Join(it.CsvStrings(), constants.CommaUnixNewLine)
}

func (it *TraceCollection) IsEqual(another *TraceCollection) bool {
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

func (it *TraceCollection) IsEqualItems(lines ...Trace) bool {
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

func (it TraceCollection) JsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JsonPtr().JsonString()
}

func (it TraceCollection) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JoinLines()
}

func (it TraceCollection) CsvStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	newSlice := make([]string, it.Length())

	for i, item := range it.Items {
		newSlice[i] = fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			item.String(),
		)
	}

	return newSlice
}

func (it TraceCollection) JsonModel() []Trace {
	return it.Items
}

func (it TraceCollection) JsonModelAny() any {
	return it.JsonModel()
}

func (it TraceCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it TraceCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it TraceCollection) Serializer() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it TraceCollection) StackTracesBytes() []byte {
	if it.IsEmpty() {
		return []byte{}
	}

	return []byte(it.CodeStacksString())
}

func (it *TraceCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*TraceCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return New.traces.Empty(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *TraceCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *TraceCollection {
	hashSet, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (it TraceCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it *TraceCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *TraceCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *TraceCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *TraceCollection) Clear() *TraceCollection {
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

	it.Items = []Trace{}

	return it
}

func (it *TraceCollection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it TraceCollection) Clone() TraceCollection {
	list := New.traces.Cap(it.Length())

	return *list.Adds(it.Items...)
}

func (it *TraceCollection) ClonePtr() *TraceCollection {
	if it == nil {
		return nil
	}

	list := New.traces.Cap(it.Length())

	return list.Adds(it.Items...)
}
