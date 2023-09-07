package internalserializer

type MustBytesInToSelfDeserializer interface {
	DeserializeMust(rawBytes []byte)
}
