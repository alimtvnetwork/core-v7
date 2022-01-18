package corepayload

type newCreator struct {
	PayloadWrapper newPayloadWrapperCreator
	Attributes     newAttributesCreator
}
