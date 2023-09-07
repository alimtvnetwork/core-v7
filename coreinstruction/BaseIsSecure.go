package coreinstruction

type BaseIsSecure struct {
	IsSecure bool // indicates secure text, invert means log payload, plain text. it will not log payload
}

func NewSecure() BaseIsSecure {
	return BaseIsSecure{
		IsSecure: true,
	}
}

func NewPlain() BaseIsSecure {
	return BaseIsSecure{
		IsSecure: false,
	}
}

func (it *BaseIsSecure) IsPlainText() bool {
	return it == nil || !it.IsSecure
}

func (it *BaseIsSecure) IsIncludePayload() bool {
	return it == nil || !it.IsSecure
}
