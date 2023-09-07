package internalserializer

type Serializer interface {
	Serialize() ([]byte, error)
}
