package errcore

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type RawErrCollection struct {
	Items []error
}

func (it RawErrCollection) AddMsg(message string) {
	it.AddString(message)
}

func (it RawErrCollection) AddMsgStackTrace(message string) {
	if len(message) == 0 {
		return
	}

	fullMessage := StackEnhance.MsgSkip(1, message)

	it.AddString(fullMessage)
}

func (it RawErrCollection) AddStackTrace(err error) {
	if err == nil {
		return
	}

	fullMessage := StackEnhance.MsgSkip(1, err.Error())

	it.AddString(fullMessage)
}

func (it RawErrCollection) AddMsgErrStackTrace(msg string, err error) {
	if err == nil {
		return
	}

	fullMessage := StackEnhance.MsgErrorSkip(1, msg, err)

	it.AddString(fullMessage)
}

func (it RawErrCollection) AddMethodName(msg string) {
	if len(msg) == 0 {
		return
	}

	fullMessage := StackEnhance.MsgSkip(1, msg)

	it.AddString(fullMessage)
}

func (it RawErrCollection) AddMessages(
	messages ...string,
) {
	if len(messages) == 0 {
		return
	}

	compiled := strings.Join(
		messages, constants.Space,
	)

	fullMessage := StackEnhance.MsgSkip(1, compiled)

	it.AddString(fullMessage)
}

func (it RawErrCollection) AddErrorWithMessage(
	err error,
	message string,
) {
	if err == nil {
		return
	}

	finalErr := ConcatMessageWithErr(message, err)
	fullMessage := StackEnhance.MsgSkip(
		1,
		finalErr.Error(),
	)

	it.AddMsg(fullMessage)
}

func (it RawErrCollection) AddIf(
	isAdd bool,
	message string,
) {
	if !isAdd {
		return
	}

	it.Add(errors.New(message))
}

func (it RawErrCollection) AddFunc(
	errFunc func() error,
) {
	if errFunc == nil {
		return
	}

	it.Add(errFunc())
}

func (it RawErrCollection) AddFuncIf(
	isAdd bool,
	errFunc func() error,
) {
	if errFunc == nil || !isAdd {
		return
	}

	it.Add(errFunc())
}

func (it RawErrCollection) AddErrorWithMessageRef(
	err error,
	message string,
	reference interface{},
) {
	if err == nil {
		return
	}

	referenceString := constants.NilAngelBracket
	if reference != nil {
		referenceString = fmt.Sprintf(
			constants.ReferenceWrapFormat,
			referenceString,
		)
	}

	it.AddErrorWithMessage(
		ConcatMessageWithErr(message, err), referenceString,
	)
}

func (it RawErrCollection) Fmt(format string, v ...interface{}) {
	if format == "" && len(v) == 0 {
		return
	}

	message := fmt.Sprintf(
		format,
		v...,
	)

	it.AddString(StackEnhance.MsgSkip(1, message))
}

func (it RawErrCollection) FmtIf(
	isAdd bool,
	format string,
	v ...interface{},
) {
	if !isAdd {
		return
	}

	it.Fmt(format, v...)
}

func (it RawErrCollection) References(
	message string,
	v ...interface{},
) {
	referencesCompiled := fmt.Sprintf(
		constants.MessageReferenceWrapFormat,
		message,
		v,
	)

	it.AddString(referencesCompiled)
}

func (it RawErrCollection) MustBeSafe() {
	if it.IsEmpty() {
		return
	}

	panic(it.CompiledError())
}

func (it RawErrCollection) HasAnyIssues() bool {
	return !it.IsEmpty()
}

func (it RawErrCollection) IsDefined() bool {
	return !it.IsEmpty()
}

func (it RawErrCollection) CompiledJsonErrorWithStackTraces() error {
	allBytes, err := it.MarshalJSON()

	if err == nil {
		return errors.New(string(allBytes))
	}

	return ConcatMessageWithErr(string(allBytes), err)
}

func (it RawErrCollection) CompiledJsonStringWithStackTraces() (jsonString string) {
	err := it.CompiledJsonErrorWithStackTraces()

	if err == nil {
		return ""
	}

	return err.Error()
}

func (it RawErrCollection) MustBeEmptyError() {
	it.MustBeSafe()
}

func (it RawErrCollection) IsCollectionType() bool {
	return true
}

func (it RawErrCollection) ReflectSetTo(toPtr interface{}) error {
	switch v := toPtr.(type) {
	case RawErrCollection:
		return FailedToConvertType.Error(
			"cannot convert to value type for RawErrCollection!",
			toPtr,
		)
	case *RawErrCollection:
		if v == nil {
			return FailedToConvertType.
				Error(
					"cannot convert to value type for RawErrCollection to nil ptr!",
					toPtr,
				)
		}

		*v = it

		return nil
	}

	return NotSupportedType.
		Error(
			"RawErrCollection.ReflectSetTo is not supported for other than ptr same time.",
			toPtr,
		)
}

func (it RawErrCollection) HandleError() {
	if it.IsEmpty() {
		return
	}

	panic(it.Items)
}

func (it *RawErrCollection) IsNull() bool {
	return it == nil || it.Items == nil
}

func (it *RawErrCollection) IsAnyNull() bool {
	return it == nil || it.Items == nil
}

func (it RawErrCollection) ErrorString() string {
	return it.String()
}

func (it RawErrCollection) Compile() string {
	return it.String()
}

func (it RawErrCollection) HandleErrorWithRefs(
	newMessage string,
	refVar, refVal interface{},
) {
	if it.IsEmpty() {
		return
	}

	reference :=
		fmt.Sprintf(
			keyValFormat,
			refVar,
			refVal,
		)

	panic(newMessage + reference + constants.DefaultLine + it.String())
}

func (it RawErrCollection) HandleErrorWithMsg(newMessage string) {
	if it.IsEmpty() {
		return
	}

	panic(newMessage + constants.DefaultLine + it.String())
}

func (it RawErrCollection) FullString() string {
	return it.String()
}

func (it RawErrCollection) FullStringWithTraces() string {
	return it.CompiledStackTracesString()
}

func (it RawErrCollection) FullStringWithTracesIf(isStackTraces bool) string {
	if isStackTraces {
		return it.FullStringWithTraces()
	}

	return it.String()
}

func (it RawErrCollection) ReferencesCompiledString() string {
	return it.String()
}

func (it RawErrCollection) CompiledErrorWithStackTraces() error {
	if it.IsEmpty() {
		return nil
	}

	return errors.New(it.CompiledStackTracesString())
}

func (it RawErrCollection) CompiledStackTracesString() string {
	if it.IsEmpty() {
		return ""
	}

	fullMessage := fmt.Sprintf(
		"%s\n\n%s",
		it.String(),
		reflectinternal.CodeStack.StacksStringDefault(2),
	)

	return fullMessage
}

func (it RawErrCollection) FullStringSplitByNewLine() []string {
	return it.Strings()
}

func (it RawErrCollection) FullStringWithoutReferences() string {
	return it.String()
}

func (it RawErrCollection) SerializeWithoutTraces() ([]byte, error) {
	if it.IsEmpty() {
		return nil, nil
	}

	return json.Marshal(it.Items)
}

func (it RawErrCollection) Serialize() ([]byte, error) {
	if it.IsEmpty() {
		return nil, nil
	}

	return json.Marshal(it.Items)
}

func (it RawErrCollection) SerializeMust() []byte {
	rawBytes, err := it.Serialize()

	MustBeEmpty(err)

	return rawBytes
}

func (it RawErrCollection) MarshalJSON() ([]byte, error) {
	if it.IsEmpty() {
		return nil, nil
	}

	return json.Marshal(it.Items)
}

func (it RawErrCollection) UnmarshalJSON(data []byte) error {
	var errItems []error
	err := json.Unmarshal(data, &errItems)

	if err == nil {
		it.Items = errItems
	}

	return err
}

func (it RawErrCollection) Value() error {
	return it.CompiledError()
}

func (it RawErrCollection) Log() {
	if it.IsEmpty() {
		return
	}

	fmt.Println(it.String())
}

func (it RawErrCollection) LogWithTraces() {
	if it.IsEmpty() {
		return
	}

	log.Println(it.CompiledErrorWithStackTraces())
}

func (it RawErrCollection) LogFatal() {
	if it.IsEmpty() {
		return
	}

	log.Fatalln(it.String())
}

func (it RawErrCollection) LogFatalWithTraces() {
	if it.IsEmpty() {
		return
	}

	log.Fatalln(it.CompiledErrorWithStackTraces())
}

func (it RawErrCollection) LogIf(isLog bool) {
	if isLog {
		it.LogFatal()
	}
}

func (it RawErrCollection) AddErrors(errs ...error) {
	it.Adds(errs...)
}

func (it RawErrCollection) ConditionalAddError(isAdd bool, err error) {
	if !isAdd {
		return
	}

	it.Add(err)
}

func (it RawErrCollection) CountStateChangeTracker() CountStateChangeTracker {
	return NewCountStateChangeTracker(&it)
}

func (it RawErrCollection) IsErrorsCollected(
	errorsItems ...error,
) bool {
	count := it.Length()

	it.Adds(errorsItems...)

	return count != it.Length()
}

func (it RawErrCollection) IsValid() bool {
	return it.IsEmpty()
}

func (it RawErrCollection) IsSuccess() bool {
	return it.IsEmpty()
}

func (it RawErrCollection) IsFailed() bool {
	return !it.IsEmpty()
}

func (it RawErrCollection) IsInvalid() bool {
	return !it.IsEmpty()
}

func (it RawErrCollection) ToRawErrCollection() *RawErrCollection {
	return &it
}

func (it *RawErrCollection) Add(err error) {
	if err == nil {
		return
	}

	it.Items = append(it.Items, err)
}

func (it *RawErrCollection) AddError(err error) {
	if err == nil {
		return
	}

	it.Items = append(it.Items, err)
}

func (it *RawErrCollection) AddWithTraceRef(
	err error,
	traces []string,
	referenceItem interface{},
) {
	if err == nil {
		return
	}

	it.Items = append(
		it.Items,
		ErrorWithTracesRefToError(err, traces, referenceItem),
	)
}

func (it *RawErrCollection) AddWithCompiledTraceRef(
	err error,
	compiledTrace string,
	referenceItem interface{},
) {
	if err == nil {
		return
	}

	compiledErr := ErrorWithCompiledTraceRefToError(
		err,
		compiledTrace,
		referenceItem,
	)

	it.Items = append(
		it.Items,
		compiledErr,
	)
}

func (it *RawErrCollection) AddWithRef(
	err error,
	referenceItem interface{},
) {
	if err == nil {
		return
	}

	compiledErr := ErrorWithRefToError(
		err,
		referenceItem,
	)

	it.Items = append(
		it.Items,
		compiledErr,
	)
}

func (it *RawErrCollection) Adds(
	errorItems ...error,
) {
	if len(errorItems) == 0 {
		return
	}

	for _, err := range errorItems {
		if err == nil {
			continue
		}

		it.Items = append(
			it.Items,
			err,
		)
	}
}

// AddString
//
//	Empty string will be ignored
func (it *RawErrCollection) AddString(
	message string,
) {
	if message == "" {
		return
	}

	it.Items = append(
		it.Items,
		errors.New(message),
	)
}

func (it *RawErrCollection) AddStringSliceAsErr(
	errSliceStrings ...string,
) {
	if len(errSliceStrings) == 0 {
		return
	}

	for _, errString := range errSliceStrings {
		if errString == "" {
			continue
		}

		it.Items = append(it.Items, errors.New(errString))
	}
}

func (it *RawErrCollection) AddErrorGetters(
	errorGetter ...errorGetter,
) {
	if len(errorGetter) == 0 {
		return
	}

	for _, errGetter := range errorGetter {
		if errGetter == nil {
			continue
		}

		err := errGetter.Error()

		if err == nil {
			continue
		}

		it.Items = append(it.Items, err)
	}
}

func (it *RawErrCollection) AddCompiledErrorGetters(
	errorGetter ...compiledErrorGetter,
) {
	if len(errorGetter) == 0 {
		return
	}

	for _, errGetter := range errorGetter {
		if errGetter == nil {
			continue
		}

		err := errGetter.CompiledError()

		if err == nil {
			continue
		}

		it.Items = append(it.Items, err)
	}
}

func (it *RawErrCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *RawErrCollection) IsEmpty() bool {
	return it == nil || len(it.Items) == 0
}

func (it *RawErrCollection) HasError() bool {
	return it != nil && len(it.Items) > 0
}

func (it *RawErrCollection) HasAnyError() bool {
	return it != nil && len(it.Items) > 0
}

func (it *RawErrCollection) Clear() {
	if it.IsEmpty() {
		return
	}

	tempItems := it.Items
	clearFunc := func() {
		for i := 0; i < len(tempItems); i++ {
			tempItems[i] = nil
		}
	}

	go clearFunc()
	it.Items = []error{}
}

func (it *RawErrCollection) Dispose() {
	if it.IsEmpty() {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it RawErrCollection) Strings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, err := range it.Items {
		slice[i] = err.Error()
	}

	return slice
}

func (it RawErrCollection) StringUsingJoiner(joiner string) string {
	if it.IsEmpty() {
		return ""
	}

	return strings.Join(
		it.Strings(),
		joiner,
	)
}

func (it RawErrCollection) StringUsingJoinerAdditional(joiner, additionalMessage string) string {
	if it.IsEmpty() {
		return ""
	}

	return strings.Join(
		it.Strings(),
		joiner,
	) + additionalMessage
}

func (it RawErrCollection) String() string {
	if it.IsEmpty() {
		return ""
	}

	return it.StringUsingJoiner(constants.NewLineUnix)
}

func (it RawErrCollection) CompiledError() error {
	if it.IsEmpty() {
		return nil
	}

	toString := it.String()

	return errors.New(toString)
}

func (it RawErrCollection) CompiledErrorUsingJoiner(joiner string) error {
	if it.IsEmpty() {
		return nil
	}

	toString := it.StringUsingJoiner(joiner)

	return errors.New(toString)
}

func (it RawErrCollection) CompiledErrorUsingJoinerAdditionalMessage(joiner, additionalMessage string) error {
	if it.IsEmpty() {
		return nil
	}

	toString := it.StringUsingJoinerAdditional(
		joiner,
		additionalMessage,
	)

	return errors.New(toString)
}

func (it RawErrCollection) CompiledErrorUsingStackTraces(joiner string, stackTraces []string) error {
	if it.IsEmpty() {
		return nil
	}

	return ErrorWithTracesRefToError(
		it.CompiledErrorUsingJoiner(joiner),
		stackTraces,
		nil,
	)
}

func (it RawErrCollection) StringWithAdditionalMessage(additionalMessage string) string {
	if it.IsEmpty() {
		return ""
	}

	return it.StringUsingJoiner(constants.NewLineUnix) + additionalMessage
}
