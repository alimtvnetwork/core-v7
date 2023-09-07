package internalserializer

type BytesInToSelfDeserializer interface {
	Deserialize(rawBytes []byte) error
}
