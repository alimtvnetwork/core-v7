package corepayload

type newCreator struct {
	PayloadWrapper     newPayloadWrapperCreator
	PayloadsCollection newPayloadsCollectionCreator
	Attributes         newAttributesCreator
	User               newUserCreator
}
