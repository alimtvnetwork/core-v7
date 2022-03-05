package serializerinf

type MustBytesInToSelfDeserializer interface {
	DeserializeMust(rawBytes []byte)
}
