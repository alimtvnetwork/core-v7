package coretaskinfo

func New(
	isSecure bool,
	name, desc, url,
	hintUrl, errorUrl,
	exampleUrl,
	chainingExample string,
	examples ...string,
) *Info {
	return &Info{
		RootName:        name,
		Description:     desc,
		Url:             url,
		HintUrl:         hintUrl,
		ErrorUrl:        errorUrl,
		ExampleUrl:      exampleUrl,
		ChainingExample: chainingExample,
		Examples:        examples,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: isSecure,
		},
	}
}

func NewSecure(
	name, desc, url,
	hintUrl, errorUrl,
	exampleUrl,
	chainingExample string,
	examples ...string,
) *Info {
	return &Info{
		RootName:        name,
		Description:     desc,
		Url:             url,
		HintUrl:         hintUrl,
		ErrorUrl:        errorUrl,
		ExampleUrl:      exampleUrl,
		ChainingExample: chainingExample,
		Examples:        examples,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func NewPlain(
	name, desc, url,
	hintUrl, errorUrl,
	exampleUrl,
	chainingExample string,
	examples ...string,
) *Info {
	return &Info{
		RootName:        name,
		Description:     desc,
		Url:             url,
		HintUrl:         hintUrl,
		ErrorUrl:        errorUrl,
		ExampleUrl:      exampleUrl,
		ChainingExample: chainingExample,
		Examples:        examples,
	}
}

func NewNameDescUrl(
	name, desc, url string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
	}
}

func NewNameDescUrlErrorUrl(
	name, desc,
	url, errUrl string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		ErrorUrl:    errUrl,
	}
}

func NewNameDescUrlErrUrlExamples(
	name, desc,
	url, errUrl string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		ErrorUrl:    errUrl,
		Examples:    examples,
	}
}

func NewNameDescUrlExamples(
	name, desc,
	url string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		Examples:    examples,
	}
}

func NewNameDescExamples(
	name, desc string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Examples:    examples,
	}
}

func NewExamples(
	name, desc string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Examples:    examples,
	}
}

func NewAllUrl(
	name, desc string,
	url, hintUrl, errUrl string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		HintUrl:     hintUrl,
		ErrorUrl:    errUrl,
		Examples:    examples,
	}
}

func NewChainingExample(
	name, desc string,
	url string,
	chainingExample string,
) *Info {
	return &Info{
		RootName:        name,
		Description:     desc,
		Url:             url,
		ChainingExample: chainingExample,
	}
}

func NewExampleUrl(
	name, desc string,
	exampleUrl string,
	chainingExample string,
) *Info {
	return &Info{
		RootName:        name,
		Description:     desc,
		ExampleUrl:      exampleUrl,
		ChainingExample: chainingExample,
	}
}

func NewExampleUrlSecure(
	name, desc string,
	exampleUrl string,
	chainingExample string,
) *Info {
	return &Info{
		RootName:        name,
		Description:     desc,
		ExampleUrl:      exampleUrl,
		ChainingExample: chainingExample,
		Examples:        nil,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func NewNameDescUrlChainingSecure(
	name, desc string,
	url string,
	chainingExample string,
) *Info {
	return &Info{
		RootName:        name,
		Description:     desc,
		Url:             url,
		ChainingExample: chainingExample,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func NewNameSecure(
	name, desc string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func NewNameExamplesSecure(
	name, desc string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Examples:    examples,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}
