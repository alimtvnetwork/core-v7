package coretaskinfo

type newInfoPlainTextCreator struct{}

func (it newInfoPlainTextCreator) Default(
	name, desc, url string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
	}
}

func (it newInfoPlainTextCreator) NameDescUrl(
	name, desc, url string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
	}
}

func (it newInfoPlainTextCreator) NameDescUrlExamples(
	name, desc, url string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		Examples:    examples,
	}
}

func (it newInfoPlainTextCreator) NewNameDescUrlErrorUrl(
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

func (it newInfoPlainTextCreator) NameDescUrlErrUrlExamples(
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

func (it newInfoPlainTextCreator) NameDescExamples(
	name, desc string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Examples:    examples,
	}
}

func (it newInfoPlainTextCreator) Examples(
	name, desc string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Examples:    examples,
	}
}

func (it newInfoPlainTextCreator) NameUrlExamples(
	name, url string,
	examples ...string,
) *Info {
	return &Info{
		RootName: name,
		Url:      url,
		Examples: examples,
	}
}

func (it newInfoPlainTextCreator) UrlExamples(
	url string,
	examples ...string,
) *Info {
	return &Info{
		Url:      url,
		Examples: examples,
	}
}

func (it newInfoPlainTextCreator) ExamplesOnly(
	examples ...string,
) *Info {
	return &Info{
		Examples: examples,
	}
}

func (it newInfoPlainTextCreator) UrlOnly(
	url string,
) *Info {
	return &Info{
		Url: url,
	}
}

func (it newInfoPlainTextCreator) ErrorUrlOnly(
	errUrl string,
) *Info {
	return &Info{
		ErrorUrl: errUrl,
	}
}

func (it newInfoPlainTextCreator) HintUrlOnly(
	hintUrl string,
) *Info {
	return &Info{
		HintUrl: hintUrl,
	}
}

func (it newInfoPlainTextCreator) DescHintUrlOnly(
	desc, hintUrl string,
) *Info {
	return &Info{
		Description: desc,
		HintUrl:     hintUrl,
	}
}

func (it newInfoPlainTextCreator) NameHintUrlOnly(
	name, hintUrl string,
) *Info {
	return &Info{
		RootName: name,
		HintUrl:  hintUrl,
	}
}

func (it newInfoPlainTextCreator) SingleExampleOnly(
	singleExample string,
) *Info {
	return &Info{
		SingleExample: singleExample,
	}
}

func (it newInfoPlainTextCreator) AllUrlExamples(
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

func (it newInfoPlainTextCreator) AllUrl(
	name, desc string,
	url, hintUrl, errUrl string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		HintUrl:     hintUrl,
		ErrorUrl:    errUrl,
	}
}

func (it newInfoPlainTextCreator) UrlSingleExample(
	name, desc string,
	url string,
	chainingExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		Url:           url,
		SingleExample: chainingExample,
	}
}

func (it newInfoPlainTextCreator) SingleExample(
	name, desc string,
	singleExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		SingleExample: singleExample,
	}
}

func (it newInfoPlainTextCreator) ExampleUrl(
	name, desc string,
	exampleUrl string,
	singleExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		ExampleUrl:    exampleUrl,
		SingleExample: singleExample,
	}
}

func (it newInfoPlainTextCreator) ExampleUrlSingleExample(
	name, desc string,
	exampleUrl string,
	singleExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		ExampleUrl:    exampleUrl,
		SingleExample: singleExample,
	}
}
