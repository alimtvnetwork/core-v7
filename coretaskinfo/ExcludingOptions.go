package coretaskinfo

type ExcludingOptions struct {
	IsExcludeRootName,
	IsExcludeDescription,
	IsExcludeUrl,
	IsExcludeHintUrl,
	IsExcludeErrorUrl,
	IsExcludeAdditionalErrorWrap,
	IsExcludeExampleUrl,
	IsExcludeSingleExample,
	IsExcludeExamples,
	IsSecureText bool // indicates secure text, invert means log payload, plain text. it will not log payload
}

func (it *ExcludingOptions) SetSecure() *ExcludingOptions {
	if it == nil {
		return &ExcludingOptions{
			IsSecureText: true,
		}
	}

	it.IsSecureText = true

	return it
}

func (it *ExcludingOptions) SetPlainText() *ExcludingOptions {
	if it == nil {
		return &ExcludingOptions{
			IsSecureText: false,
		}
	}

	it.IsSecureText = false

	return it
}

func (it *ExcludingOptions) IsEmpty() bool {
	return it == nil ||
		!it.IsExcludeRootName &&
			!it.IsExcludeDescription &&
			!it.IsExcludeUrl &&
			!it.IsExcludeHintUrl &&
			!it.IsExcludeErrorUrl &&
			!it.IsExcludeAdditionalErrorWrap &&
			!it.IsExcludeExampleUrl &&
			!it.IsExcludeSingleExample &&
			!it.IsExcludeExamples &&
			!it.IsSecureText
}

func (it *ExcludingOptions) IsZero() bool {
	return it == nil ||
		!it.IsExcludeRootName &&
			!it.IsExcludeDescription &&
			!it.IsExcludeUrl &&
			!it.IsExcludeHintUrl &&
			!it.IsExcludeErrorUrl &&
			!it.IsExcludeAdditionalErrorWrap &&
			!it.IsExcludeExampleUrl &&
			!it.IsExcludeSingleExample &&
			!it.IsExcludeExamples &&
			!it.IsSecureText
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

func (it *ExcludingOptions) IsIncludeSingleExample() bool {
	return it == nil || !it.IsExcludeSingleExample
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

func (it ExcludingOptions) ToPtr() *ExcludingOptions {
	return &it
}

func (it ExcludingOptions) ToNonPtr() ExcludingOptions {
	return it
}

func (it ExcludingOptions) Clone() ExcludingOptions {
	return ExcludingOptions{
		IsExcludeRootName:            false,
		IsExcludeDescription:         false,
		IsExcludeUrl:                 false,
		IsExcludeHintUrl:             false,
		IsExcludeErrorUrl:            false,
		IsExcludeAdditionalErrorWrap: false,
		IsExcludeExampleUrl:          false,
		IsExcludeSingleExample:       false,
		IsExcludeExamples:            false,
		IsSecureText:                 false,
	}
}
