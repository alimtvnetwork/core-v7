package coreinterface

type SelfBytesSerializerDeserializer interface {
	Serializer
	MustSerializer
	FieldBytesToPointerDeserializer
}
