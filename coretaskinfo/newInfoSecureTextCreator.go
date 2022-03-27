package coretaskinfo

type newInfoSecureTextCreator struct{}

func (it newInfoSecureTextCreator) Default(
	name, desc, url string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NameDescUrl(
	name, desc, url string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NameDescUrlExamples(
	name, desc, url string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		Examples:    examples,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NewNameDescUrlErrorUrl(
	name, desc,
	url, errUrl string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		ErrorUrl:    errUrl,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NameDescUrlErrUrlExamples(
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
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NameDescExamples(
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

func (it newInfoSecureTextCreator) Examples(
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

func (it newInfoSecureTextCreator) ExamplesOnly(
	examples ...string,
) *Info {
	return &Info{
		Examples: examples,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) UrlOnly(
	url string,
) *Info {
	return &Info{
		Url: url,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) ErrorUrlOnly(
	errUrl string,
) *Info {
	return &Info{
		ErrorUrl: errUrl,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) HintUrlOnly(
	hintUrl string,
) *Info {
	return &Info{
		HintUrl: hintUrl,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) DescHintUrlOnly(
	desc, hintUrl string,
) *Info {
	return &Info{
		Description: desc,
		HintUrl:     hintUrl,
	}
}

func (it newInfoSecureTextCreator) NameHintUrlOnly(
	name, hintUrl string,
) *Info {
	return &Info{
		RootName: name,
		HintUrl:  hintUrl,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) SingleExampleOnly(
	singleExample string,
) *Info {
	return &Info{
		SingleExample: singleExample,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) AllUrlExamples(
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
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) AllUrl(
	name, desc string,
	url, hintUrl, errUrl string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		HintUrl:     hintUrl,
		ErrorUrl:    errUrl,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) UrlSingleExample(
	name, desc string,
	url string,
	chainingExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		Url:           url,
		SingleExample: chainingExample,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) SingleExample(
	name, desc string,
	singleExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		SingleExample: singleExample,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) ExampleUrl(
	name, desc string,
	exampleUrl string,
	singleExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		ExampleUrl:    exampleUrl,
		SingleExample: singleExample,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) ExampleUrlSingleExample(
	name, desc string,
	exampleUrl string,
	singleExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		ExampleUrl:    exampleUrl,
		SingleExample: singleExample,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NewExampleUrlSecure(
	name, desc string,
	exampleUrl string,
	chainingExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		ExampleUrl:    exampleUrl,
		SingleExample: chainingExample,
		Examples:      nil,
		ExcludeOptions: ExcludingOptions{
			IsSecureText: true,
		},
	}
}
