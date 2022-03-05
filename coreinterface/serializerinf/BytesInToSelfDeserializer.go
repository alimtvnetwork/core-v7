package serializerinf

type BytesInToSelfDeserializer interface {
	Deserialize(rawBytes []byte) error
}
