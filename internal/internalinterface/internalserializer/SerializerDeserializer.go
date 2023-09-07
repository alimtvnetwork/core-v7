package internalserializer

type SerializerDeserializer interface {
	Serializer
	MustSerializer
	BytesInToSelfDeserializer
}
