package serializerinf

type SerializerDeserializer interface {
	Serializer
	MustSerializer
	BytesInToSelfDeserializer
}
