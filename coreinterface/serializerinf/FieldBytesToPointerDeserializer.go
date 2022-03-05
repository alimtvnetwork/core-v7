package serializerinf

type FieldBytesToPointerDeserializer interface {
	Deserialize(toPtr interface{}) error
}
