package codestack

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/defaultcapacity"
)

type TraceCollection struct {
	Items []Trace `json:"Items,omitempty"`
}

func NewTraceCollection(capacity int) *TraceCollection {
	slice := make([]Trace, 0, capacity)

	return &TraceCollection{
		slice,
	}
}

func NewStacksCollection() *TraceCollection {
	return NewTraceCollection(DefaultStackCount * 2)
}

func NewNewTraceCollectionUsing(
	isClone bool,
	traces ...Trace,
) *TraceCollection {
	if traces == nil {
		return EmptyTraceCollection()
	}

	if !isClone {
		return &TraceCollection{
			traces,
		}
	}

	slice := NewTraceCollection(len(traces))

	return slice.Adds(traces...)
}

func EmptyTraceCollection() *TraceCollection {
	return NewTraceCollection(0)
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
		trace := New(i)
		isSkip := isSkipInvalid && trace.HasIssues()

		if isSkip && isBreakOnceInvalid {
			return it
		} else if isSkip {
			continue
		}

		it.Items = append(
			it.Items,
			trace)
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
		DefaultStackCount)
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
		trace := New(i)
		isSkip := isSkipInvalid && trace.HasIssues()

		if isSkip && isBreakOnceInvalid {
			return it
		} else if isSkip {
			continue
		}

		isTake, isBreak := filterFunc(&trace)

		if isTake {
			it.Items = append(
				it.Items,
				trace)
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
		traces...)

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
			*trace)
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
		additionalTraces...)
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
		stackCount)
}

func (it *TraceCollection) ConcatNewUsingSkip(
	skipIndex int,
) *TraceCollection {
	cloned := it.Clone()

	return cloned.AddsUsingSkip(
		true,
		true,
		skipIndex+defaultInternalSkip,
		DefaultStackCount)
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

func (it *TraceCollection) FirstDynamic() interface{} {
	return it.Items[0]
}

func (it *TraceCollection) First() Trace {
	return it.Items[0]
}

func (it *TraceCollection) LastDynamic() interface{} {
	return it.Items[it.LastIndex()]
}

func (it *TraceCollection) Last() Trace {
	return it.Items[it.LastIndex()]
}

func (it *TraceCollection) FirstOrDefaultDynamic() interface{} {
	return it.FirstOrDefault()
}

func (it *TraceCollection) FirstOrDefault() *Trace {
	if it.IsEmpty() {
		return nil
	}

	first := it.First()

	return &first
}

func (it *TraceCollection) LastOrDefaultDynamic() interface{} {
	return it.LastOrDefault()
}

func (it *TraceCollection) LastOrDefault() *Trace {
	if it.IsEmpty() {
		return nil
	}

	last := it.Last()

	return &last
}

func (it *TraceCollection) SkipDynamic(skippingItemsCount int) interface{} {
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

func (it *TraceCollection) TakeDynamic(takeDynamicItems int) interface{} {
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

func (it *TraceCollection) LimitDynamic(limit int) interface{} {
	return it.Take(limit)
}

func (it *TraceCollection) Limit(limit int) []Trace {
	return it.Take(limit)
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
	return it == nil || it.Length() == 0
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
		limit)
	list := make(
		[]Trace,
		0,
		length)

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

	traceCollection := NewNewTraceCollectionUsing(
		false, list...)

	return traceCollection
}

func (it *TraceCollection) FilterPackageNameTraceCollection(
	packageName string,
) *TraceCollection {
	return it.FilterTraceCollection(func(trace *Trace) (isTake, isBreak bool) {
		return trace.PackageName == packageName, false
	})
}

func (it *TraceCollection) SkipFilterPackageNameTraceCollection(
	packageName string,
) *TraceCollection {
	return it.FilterTraceCollection(func(trace *Trace) (isTake, isBreak bool) {
		return trace.PackageName != packageName, false
	})
}

func (it *TraceCollection) FilterMethodNameTraceCollection(
	methodName string,
) *TraceCollection {
	return it.FilterTraceCollection(func(trace *Trace) (isTake, isBreak bool) {
		return trace.MethodName == methodName, false
	})
}

func (it *TraceCollection) SkipFilterMethodNameTraceCollection(
	methodName string,
) *TraceCollection {
	return it.FilterTraceCollection(func(trace *Trace) (isTake, isBreak bool) {
		return trace.MethodName != methodName, false
	})
}

// FilterFullMethodNameTraceCollection
//
// fullMethodName := packageName.struct.methodName
func (it *TraceCollection) FilterFullMethodNameTraceCollection(
	fullMethodName string,
) *TraceCollection {
	return it.FilterTraceCollection(func(trace *Trace) (isTake, isBreak bool) {
		return trace.PackageMethodName == fullMethodName, false
	})
}

// SkipFilterFullMethodNameTraceCollection
//
// fullMethodName := packageName.struct.methodName
func (it *TraceCollection) SkipFilterFullMethodNameTraceCollection(
	fullMethodName string,
) *TraceCollection {
	return it.FilterTraceCollection(func(trace *Trace) (isTake, isBreak bool) {
		return trace.PackageMethodName != fullMethodName, false
	})
}

func (it *TraceCollection) SkipFilterFilenameTraceCollection(
	skipFilename string,
) *TraceCollection {
	return it.FilterTraceCollection(func(trace *Trace) (isTake, isBreak bool) {
		return trace.FileName != skipFilename, false
	})
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

func (it *TraceCollection) JsonStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.JsonString()
	}

	return list
}

func (it *TraceCollection) JoinFileWithLinesStrings(joiner string) string {
	return strings.Join(it.FileWithLinesStrings(), joiner)
}

func (it *TraceCollection) JoinJsonStrings(joiner string) string {
	return strings.Join(it.JsonStrings(), joiner)
}

func (it *TraceCollection) Join(joiner string) string {
	return strings.Join(it.Strings(), joiner)
}

func (it *TraceCollection) JoinLine() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}

func (it *TraceCollection) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it *TraceCollection) JoinCsvLine() string {
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

func (it *TraceCollection) JsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.Json().JsonString()
}

func (it *TraceCollection) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JoinLine()
}

func (it *TraceCollection) CsvStrings() []string {
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

func (it *TraceCollection) JsonModel() []Trace {
	return it.Items
}

func (it *TraceCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it TraceCollection) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it TraceCollection) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

func (it *TraceCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*TraceCollection, error) {
	err := jsonResult.Unmarshal(&it)

	if err != nil {
		return EmptyTraceCollection(), err
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

func (it TraceCollection) Clone() TraceCollection {
	list := NewTraceCollection(it.Length())

	return *list.Adds(it.Items...)
}

func (it *TraceCollection) ClonePtr() *TraceCollection {
	if it == nil {
		return nil
	}

	list := NewTraceCollection(it.Length())

	return list.Adds(it.Items...)
}
