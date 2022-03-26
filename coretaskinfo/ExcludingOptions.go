package coretaskinfo

type ExcludingOptions struct {
	IsExcludeRootName,
	IsExcludeDescription,
	IsExcludeUrl,
	IsExcludeHintUrl,
	IsExcludeErrorUrl,
	IsExcludeAdditionalErrorWrap,
	IsExcludeExampleUrl,
	IsExcludeChainingExample,
	IsExcludeExamples,
	IsSecureText bool // indicates secure text, invert means log payload, plain text. it will not log payload
}

func (it *ExcludingOptions) IsIncludeRootName() bool {
	return it == nil || !it.IsExcludeRootName
}

func (it *ExcludingOptions) IsIncludeDescription() bool {
	return it == nil || !it.IsExcludeDescription
}

func (it *ExcludingOptions) IsIncludeUrl() bool {
	return it == nil || !it.IsExcludeUrl
}

func (it *ExcludingOptions) IsIncludeHintUrl() bool {
	return it == nil || !it.IsExcludeHintUrl
}

func (it *ExcludingOptions) IsIncludeErrorUrl() bool {
	return it == nil || !it.IsExcludeErrorUrl
}

func (it *ExcludingOptions) IsIncludeExampleUrl() bool {
	return it == nil || !it.IsExcludeExampleUrl
}

func (it *ExcludingOptions) IsIncludeChainingExample() bool {
	return it == nil || !it.IsExcludeChainingExample
}

func (it *ExcludingOptions) IsIncludeExamples() bool {
	return it == nil || !it.IsExcludeExamples
}


func (it *ExcludingOptions) IsIncludeAdditionalErrorWrap() bool {
	return it == nil || !it.IsExcludeAdditionalErrorWrap
}

func (it *ExcludingOptions) IsIncludePayloads() bool {
	return it == nil || !it.IsSecureText
}
