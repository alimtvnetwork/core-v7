package internalserializer

type FieldBytesToPointerDeserializer interface {
	Deserialize(toPtr interface{}) error
}
