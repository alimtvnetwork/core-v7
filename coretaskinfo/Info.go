package coretaskinfo

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
)

type Info struct {
	RootName       string            `json:"RootName,omitempty"`
	Description    string            `json:"Description,omitempty"`
	Url            string            `json:"Url,omitempty"`
	HintUrl        string            `json:"HintUrl,omitempty"`
	ErrorUrl       string            `json:"ErrorUrl,omitempty"`
	ExampleUrl     string            `json:"ExampleUrl,omitempty"`
	SingleExample  string            `json:"SingleExample,omitempty"`
	Examples       []string          `json:"Examples,omitempty"` // proves sample examples to call things correctly
	ExcludeOptions *ExcludingOptions `json:"ExcludeOptions,omitempty"`
	lazyMap        map[string]string
}

// SetSecure
//
//  on nil creates and returns new info with secure flag
func (it *Info) SetSecure() *Info {
	if it == nil {
		return &Info{
			ExcludeOptions: &ExcludingOptions{
				IsSecureText: true,
			},
		}
	}

	it.ExcludeOptions = it.
		ExcludeOptions.
		SetSecure()

	return it
}

// SetPlain
//
//  on nil creates and returns
//  new info which is plain not secure
func (it *Info) SetPlain() *Info {
	if it == nil {
		return &Info{} // plain text
	}

	it.ExcludeOptions = it.
		ExcludeOptions.
		SetPlainText()

	return it
}

func (it *Info) IsIncludeRootName() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeRootName() &&
		it.RootName != ""
}

func (it *Info) IsIncludeDescription() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeDescription() &&
		it.Description != ""
}

func (it *Info) IsIncludeUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeUrl() &&
		it.Url != ""
}

func (it *Info) IsIncludeHintUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeHintUrl() &&
		it.HintUrl != ""
}

func (it *Info) IsIncludeErrorUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeErrorUrl() &&
		it.ErrorUrl != ""
}

// IsIncludeAdditionalErrorWrap
//
//  returns true on null or it.ExcludeOptions.IsIncludeAdditionalErrorWrap
func (it *Info) IsIncludeAdditionalErrorWrap() bool {
	return it == nil ||
		it.ExcludeOptions.IsIncludeAdditionalErrorWrap()
}

func (it *Info) IsIncludeExampleUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeExampleUrl() &&
		it.ExampleUrl != ""
}

func (it *Info) IsIncludeSingleExample() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeSingleExample() &&
		it.SingleExample != ""
}

func (it *Info) IsIncludeExamples() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeExamples() &&
		len(it.Examples) > 0
}

func (it *Info) IsSecure() bool {
	return it != nil && it.ExcludeOptions.IsSafeSecureText()
}

func (it *Info) IsPlainText() bool {
	return it == nil || it.ExcludeOptions.IsIncludePayloads()
}

func (it *Info) IsIncludePayloads() bool {
	return it == nil || it.ExcludeOptions.IsIncludePayloads()
}

func (it *Info) IsExcludePayload() bool {
	return it != nil && it.ExcludeOptions.IsSafeSecureText()
}

// IsExcludeRootName
//
//  returns true on defined (not null) and
//  it.ExcludeOptions.IsExcludeRootName
//
//  return false on null
func (it *Info) IsExcludeRootName() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeRootName()
}

// IsExcludeDescription
//
//  returns true on defined (not null) and
//  it.ExcludeOptions.IsExcludeDescription
//
//  return false on null
func (it *Info) IsExcludeDescription() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeDescription()
}

// IsExcludeUrl
//
//  returns true on defined (not null) and
//  it.ExcludeOptions.IsExcludeUrl
//
//  return false on null
func (it *Info) IsExcludeUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeUrl()
}

// IsExcludeHintUrl
//
//  returns true on defined (not null) and
//  it.ExcludeOptions.IsExcludeHintUrl
//
//  return false on null
func (it *Info) IsExcludeHintUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeHintUrl()
}

// IsExcludeErrorUrl
//
//  returns true on defined (not null) and
//  it.ExcludeOptions.IsExcludeErrorUrl
//
//  return false on null
func (it *Info) IsExcludeErrorUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeErrorUrl()
}

// IsExcludeAdditionalErrorWrap
//
//  returns true on defined (not null) and
//  it.ExcludeOptions.IsExcludeAdditionalErrorWrap
//
//  return false on null
func (it *Info) IsExcludeAdditionalErrorWrap() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeAdditionalErrorWrap()
}

// IsExcludeExampleUrl
//
//  return true on null
func (it *Info) IsExcludeExampleUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeExampleUrl()
}

// IsExcludeSingleExample
//
//  return true on null
func (it *Info) IsExcludeSingleExample() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeSingleExample()
}

// IsExcludeExamples
//
//  return true on null
func (it *Info) IsExcludeExamples() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeExamples()
}

func (it *Info) Name() string {
	if it.IsNull() {
		return ""
	}

	return it.RootName
}

func (it *Info) IsNull() bool {
	return it == nil
}

func (it *Info) IsDefined() bool {
	return it != nil && it.RootName != ""
}

func (it *Info) HasAnyName() bool {
	return it != nil && it.RootName != ""
}

func (it *Info) IsName(name string) bool {
	return it != nil && it.RootName == name
}

func (it *Info) SafeName() string {
	if it.IsNull() {
		return ""
	}

	return it.RootName
}

func (it *Info) SafeDescription() string {
	if it.IsNull() {
		return ""
	}

	return it.Description
}

func (it *Info) SafeUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.Url
}

func (it *Info) SafeHintUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.HintUrl
}

func (it *Info) SafeErrorUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.ErrorUrl
}

func (it *Info) SafeExampleUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.ExampleUrl
}

func (it *Info) SafeChainingExample() string {
	if it.IsNull() {
		return ""
	}

	return it.ExampleUrl
}

func (it *Info) ExamplesAsSlice() *corestr.SimpleSlice {
	if it.IsNull() {
		return corestr.Empty.SimpleSlice()
	}

	return corestr.New.SimpleSlice.Strings(it.Examples)
}

func (it *Info) Options() *ExcludingOptions {
	if it == nil {
		return &ExcludingOptions{}
	}

	return it.ExcludeOptions
}

func (it *Info) IsEmpty() bool {
	return it == nil
}

func (it *Info) HasRootName() bool {
	return it != nil && it.RootName != ""
}

func (it *Info) HasDescription() bool {
	return it != nil && it.Description != ""
}

func (it *Info) HasUrl() bool {
	return it != nil && it.Url != ""
}

func (it *Info) HasHintUrl() bool {
	return it != nil && it.HintUrl != ""
}

func (it *Info) HasErrorUrl() bool {
	return it != nil && it.ErrorUrl != ""
}

func (it *Info) HasExampleUrl() bool {
	return it != nil && it.ExampleUrl != ""
}

func (it *Info) HasChainingExample() bool {
	return it != nil && it.SingleExample != ""
}

func (it *Info) HasExamples() bool {
	return it != nil && len(it.Examples) > 0
}

func (it *Info) HasExcludeOptions() bool {
	return it != nil && !it.ExcludeOptions.IsEmpty()
}

func (it *Info) IsEmptyName() bool {
	return it == nil || it.RootName == ""
}

func (it *Info) IsEmptyDescription() bool {
	return it == nil || it.Description == ""
}

func (it *Info) IsEmptyUrl() bool {
	return it == nil || it.Url == ""
}

func (it *Info) IsEmptyHintUrl() bool {
	return it == nil || it.HintUrl == ""
}

func (it *Info) IsEmptyErrorUrl() bool {
	return it == nil || it.ErrorUrl == ""
}

func (it *Info) IsEmptyExampleUrl() bool {
	return it == nil || it.ExampleUrl == ""
}

func (it *Info) IsEmptySingleExample() bool {
	return it == nil || it.SingleExample == ""
}

func (it *Info) IsEmptyExamples() bool {
	return it == nil || len(it.Examples) == 0
}

func (it *Info) IsEmptyExcludeOptions() bool {
	return it == nil || it.ExcludeOptions.IsZero()
}

func (it *Info) HasAnyItem() bool {
	return it != nil
}

func (it Info) Json() corejson.Result {
	return corejson.New(it)
}

func (it Info) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Info) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	return jsonResult.Deserialize(it)
}

func (it Info) JsonString() string {
	if it.IsNull() {
		return ""
	}

	return it.JsonPtr().JsonString()
}

func (it *Info) PrettyJsonStringWithPayloads(
	payloads []byte,
) string {
	return corejson.
		NewPtr(it.MapWithPayload(payloads)).
		PrettyJsonString()
}

func (it *Info) PrettyJsonString() string {
	if it.IsNull() {
		return ""
	}

	return corejson.
		NewPtr(it).
		PrettyJsonString()
}

func (it *Info) LazyMapPrettyJsonString() string {
	lazyMap := it.LazyMap()

	return corejson.
		NewPtr(lazyMap).
		PrettyJsonString()
}

func (it Info) JsonStringMust() string {
	jsonResult := it.Json()
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

func (it Info) ToPtr() *Info {
	return &it
}

func (it Info) ToNonPtr() Info {
	return it
}

func (it Info) Clone() Info {
	return Info{
		RootName:       it.RootName,
		Description:    it.Description,
		Url:            it.Url,
		HintUrl:        it.HintUrl,
		ErrorUrl:       it.ErrorUrl,
		ExampleUrl:     it.ExampleUrl,
		SingleExample:  it.SingleExample,
		Examples:       it.Examples,
		ExcludeOptions: it.ExcludeOptions.ClonePtr(),
	}
}

func (it *Info) ClonePtr() *Info {
	if it == nil {
		return nil
	}

	return &Info{
		RootName:       it.RootName,
		Description:    it.Description,
		Url:            it.Url,
		HintUrl:        it.HintUrl,
		ErrorUrl:       it.ErrorUrl,
		ExampleUrl:     it.ExampleUrl,
		SingleExample:  it.SingleExample,
		Examples:       it.Examples,
		ExcludeOptions: it.ExcludeOptions.ClonePtr(),
	}
}

func (it *Info) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *Info) Deserialize(toPtr interface{}) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}

func (it *Info) ExamplesAsString() (compiledString string) {
	if it.IsNull() {
		return ""
	}

	return strings.Join(
		it.Examples,
		constants.CommaSpace)
}

func (it *Info) Map() map[string]string {
	if it.IsNull() || it.IsExcludeAdditionalErrorWrap() {
		return map[string]string{}
	}

	newMap := make(
		map[string]string,
		constants.Capacity8)

	if it.IsIncludeRootName() {
		newMap[infoFieldName] = it.RootName
	}

	if it.IsIncludeDescription() {
		newMap[infoFieldDescription] = it.Description
	}

	if it.IsIncludeUrl() {
		newMap[infoFieldUrl] = it.Url
	}

	if it.IsIncludeHintUrl() {
		newMap[infoFieldHintUrl] = it.HintUrl
	}

	if it.IsIncludeErrorUrl() {
		newMap[infoFieldErrorUrl] = it.ErrorUrl
	}

	if it.IsIncludeExampleUrl() {
		newMap[infoFieldExampleUrl] = it.ExampleUrl
	}

	if it.IsIncludeSingleExample() {
		newMap[infoFieldSingleExample] = it.SingleExample
	}

	if it.IsIncludeExamples() {
		newMap[infoFieldExamples] = it.ExamplesAsString()
	}

	return newMap
}

func (it *Info) MapWithPayload(
	payloads []byte,
) map[string]string {
	compiledMap := it.Map()

	if it.IsIncludePayloads() {
		compiledMap[payloadsField] = corejson.BytesToString(payloads)
	}

	return compiledMap
}

func (it *Info) LazyMapWithPayload(
	payloads []byte,
) map[string]string {
	compiledMap := it.LazyMap()

	if it.IsIncludePayloads() {
		compiledMap[payloadsField] = corejson.BytesToString(payloads)
	}

	return compiledMap
}

func (it *Info) MapWithPayloadAsAny(
	payloadsAny interface{},
) map[string]string {
	compiledMap := it.Map()

	if it.IsExcludePayload() {
		return compiledMap
	}

	jsonResult := corejson.
		AnyTo.
		SerializedJsonResult(payloadsAny)

	if jsonResult.HasError() {
		compiledMap[payloadsErrField] = jsonResult.MeaningfulErrorMessage()
	}

	compiledMap[payloadsField] = jsonResult.JsonString()

	return compiledMap
}

func (it *Info) LazyMapWithPayloadAsAny(
	payloadsAny interface{},
) map[string]string {
	compiledMap := it.LazyMap()

	if it.IsExcludePayload() {
		return compiledMap
	}

	jsonResult := corejson.
		AnyTo.
		SerializedJsonResult(payloadsAny)

	if jsonResult.HasError() {
		compiledMap[payloadsErrField] = jsonResult.MeaningfulErrorMessage()
	}

	compiledMap[payloadsField] = jsonResult.JsonString()

	return compiledMap
}

func (it *Info) LazyMap() map[string]string {
	if it.IsNull() || it.IsExcludeAdditionalErrorWrap() {
		return map[string]string{}
	}

	if it.lazyMap != nil {
		return it.lazyMap
	}

	it.lazyMap = it.Map()

	return it.lazyMap
}

func (it *Info) String() string {
	if it == nil {
		return ""
	}

	return it.PrettyJsonString()
}

func (it Info) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
