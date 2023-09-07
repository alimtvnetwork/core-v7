package internalserializer

type SelfBytesSerializerDeserializer interface {
	Serializer
	MustSerializer
	FieldBytesToPointerDeserializer
}
