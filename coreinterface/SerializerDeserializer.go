package coreinterface

type SerializerDeserializer interface {
	Serializer
	MustSerializer
	BytesInToSelfDeserializer
}
