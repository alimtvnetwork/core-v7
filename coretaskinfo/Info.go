package coretaskinfo

import "gitlab.com/evatix-go/core/coredata/corejson"

type Info struct {
	RootName          string
	Description, Url  string
	HintUrl, ErrorUrl string
	ExampleUrl        string
	ChainingExample   string
	Examples          []string // proves sample examples to call things correctly
	ExcludeOptions    ExcludingOptions
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

func (it *Info) IsIncludeChainingExample() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeChainingExample() &&
		it.ChainingExample != ""
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

func (it Info) TaskDescription() string {
	return it.Description
}

func (it Info) MainUrl() string {
	return it.Url
}

func (it Info) UrlError() string {
	return it.ErrorUrl
}

func (it Info) Options() ExcludingOptions {
	return it.ExcludeOptions
}

func (it *Info) IsEmpty() bool {
	return it == nil
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

func (it *Info) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Deserialize(it)
}

func (it Info) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it Info) JsonString() string {
	return it.JsonPtr().PrettyJsonString()
}

func (it Info) JsonStringMust() string {
	jsonResult := it.Json()
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

func (it Info) String() string {
	return it.Json().PrettyJsonString()
}

func (it Info) ToPtr() *Info {
	return &it
}

func (it Info) ToNonPtr() Info {
	return it
}

func (it Info) Clone() Info {
	return Info{
		RootName:        it.RootName,
		Description:     it.Description,
		Url:             it.Url,
		HintUrl:         it.HintUrl,
		ErrorUrl:        it.ErrorUrl,
		ExampleUrl:      it.ExampleUrl,
		ChainingExample: it.ChainingExample,
		Examples:        it.Examples,
		ExcludeOptions:  it.ExcludeOptions,
	}
}

func (it Info) ClonePtr() *Info {
	return &Info{
		RootName:        it.RootName,
		Description:     it.Description,
		Url:             it.Url,
		HintUrl:         it.HintUrl,
		ErrorUrl:        it.ErrorUrl,
		ExampleUrl:      it.ExampleUrl,
		ChainingExample: it.ChainingExample,
		Examples:        it.Examples,
		ExcludeOptions:  it.ExcludeOptions,
	}
}
