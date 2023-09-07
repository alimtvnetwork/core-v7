package internalserializer

type SerializerDeserializerBinder interface {
	SerializerDeserializer
	AsSerializerDeserializerBinder() SerializerDeserializerBinder
}
