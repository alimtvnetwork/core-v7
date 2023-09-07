package coretaskinfo

import (
	"gitlab.com/auk-go/core/coredata/corejson"
)

type newInfoCreator struct {
	Plain  newInfoPlainTextCreator
	Secure newInfoSecureTextCreator
}

func (it newInfoCreator) Deserialized(
	rawBytes []byte,
) (parsedInfo *Info, parsingErr error) {
	emptyInfo := Info{}
	parsingErr = corejson.Deserialize.UsingBytes(
		rawBytes, &emptyInfo)

	return emptyInfo.ToPtr(), parsingErr
}

func (it newInfoCreator) DeserializedUsingJsonResult(
	jsonResult *corejson.Result,
) (parsedInfo *Info, parsingErr error) {
	emptyInfo := Info{}
	parsingErr = jsonResult.Deserialize(
		&emptyInfo)

	return emptyInfo.ToPtr(), parsingErr
}

func (it newInfoCreator) Default(
	name, desc, url string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
	}
}

func (it newInfoCreator) Examples(
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

func (it newInfoCreator) Create(
	isSecure bool,
	name, desc, url,
	hintUrl, errorUrl,
	exampleUrl,
	chainingExample string,
	examples ...string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		Url:           url,
		HintUrl:       hintUrl,
		ErrorUrl:      errorUrl,
		ExampleUrl:    exampleUrl,
		SingleExample: chainingExample,
		Examples:      examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: isSecure,
		},
	}
}

func (it newInfoCreator) SecureCreate(
	name, desc, url,
	hintUrl, errorUrl,
	exampleUrl,
	chainingExample string,
	examples ...string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		Url:           url,
		HintUrl:       hintUrl,
		ErrorUrl:      errorUrl,
		ExampleUrl:    exampleUrl,
		SingleExample: chainingExample,
		Examples:      examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoCreator) PlainCreate(
	name, desc, url,
	hintUrl, errorUrl,
	exampleUrl,
	chainingExample string,
	examples ...string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		Url:           url,
		HintUrl:       hintUrl,
		ErrorUrl:      errorUrl,
		ExampleUrl:    exampleUrl,
		SingleExample: chainingExample,
		Examples:      examples,
	}
}
