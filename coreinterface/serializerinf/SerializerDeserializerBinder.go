package serializerinf

type SerializerDeserializerBinder interface {
	SerializerDeserializer
	AsSerializerDeserializerBinder() SerializerDeserializer
}
