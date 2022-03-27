package coretaskinfo

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type Info struct {
	RootName          string
	Description, Url  string
	HintUrl, ErrorUrl string
	ExampleUrl        string
	SingleExample     string
	Examples          []string // proves sample examples to call things correctly
	ExcludeOptions    ExcludingOptions
	lazyMap           map[string]string
}

func (it *Info) SetSecure() *Info {
	if it == nil {
		return nil
	}

	it.ExcludeOptions = it.
		ExcludeOptions.
		SetSecure().
		ToNonPtr()

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

func (it *Info) IsIncludeAdditionalErrorWrap() bool {
	return it != nil &&
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

func (it Info) IsSecure() bool {
	return it.ExcludeOptions.IsSecureText
}

func (it Info) IsPlainText() bool {
	return it.ExcludeOptions.IsIncludePayloads()
}

func (it Info) IsIncludePayloads() bool {
	return it.ExcludeOptions.IsIncludePayloads()
}

func (it Info) Name() string {
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

func (it Info) SafeName() string {
	if it.IsNull() {
		return ""
	}

	return it.RootName
}

func (it Info) SafeDescription() string {
	if it.IsNull() {
		return ""
	}

	return it.Description
}

func (it Info) SafeUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.Url
}

func (it Info) SafeHintUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.HintUrl
}

func (it Info) SafeErrorUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.ErrorUrl
}

func (it Info) SafeExampleUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.ExampleUrl
}

func (it Info) SafeChainingExample() string {
	if it.IsNull() {
		return ""
	}

	return it.ExampleUrl
}

func (it Info) Options() ExcludingOptions {
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

func (it Info) PrettyJsonString() string {
	if it.IsNull() {
		return ""
	}

	return corejson.
		NewPtr(it).
		PrettyJsonString()
}

func (it Info) LazyMapPrettyJsonString() string {
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
		ExcludeOptions: it.ExcludeOptions,
	}
}

func (it Info) ClonePtr() *Info {
	return &Info{
		RootName:       it.RootName,
		Description:    it.Description,
		Url:            it.Url,
		HintUrl:        it.HintUrl,
		ErrorUrl:       it.ErrorUrl,
		ExampleUrl:     it.ExampleUrl,
		SingleExample:  it.SingleExample,
		Examples:       it.Examples,
		ExcludeOptions: it.ExcludeOptions,
	}
}

func (it Info) Serialize() ([]byte, error) {
	return it.Json().Raw()
}

func (it Info) Deserialize(toPtr interface{}) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}

func (it Info) ExamplesAsString() (compiledString string) {
	if it.IsNull() {
		return ""
	}

	return strings.Join(
		it.Examples,
		constants.CommaSpace)
}

func (it Info) Map() map[string]string {
	if it.IsNull() {
		return map[string]string{}
	}

	newMap := make(
		map[string]string,
		constants.Capacity8)

	if it.IsIncludeRootName() {
		newMap["Name"] = it.RootName
	}

	if it.IsIncludeDescription() {
		newMap["Description"] = it.Description
	}

	if it.IsIncludeUrl() {
		newMap["Url"] = it.Url
	}

	if it.IsIncludeHintUrl() {
		newMap["HintUrl"] = it.HintUrl
	}

	if it.IsIncludeErrorUrl() {
		newMap["ErrorUrl"] = it.ErrorUrl
	}

	if it.IsIncludeExampleUrl() {
		newMap["ExampleUrl"] = it.ExampleUrl
	}

	if it.IsIncludeSingleExample() {
		newMap["SingleExample"] = it.SingleExample
	}

	if it.IsIncludeExamples() {
		newMap["Examples"] = it.ExamplesAsString()
	}

	return newMap
}

func (it Info) LazyMap() map[string]string {
	if it.IsNull() {
		return map[string]string{}
	}

	if it.lazyMap != nil {
		return it.lazyMap
	}

	it.lazyMap = it.Map()

	return it.lazyMap
}

func (it Info) String() string {
	return it.PrettyJsonString()
}

func (it Info) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
