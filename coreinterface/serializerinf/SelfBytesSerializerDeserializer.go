package serializerinf

type SelfBytesSerializerDeserializer interface {
	Serializer
	MustSerializer
	FieldBytesToPointerDeserializer
}
