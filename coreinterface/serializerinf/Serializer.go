package serializerinf

type Serializer interface {
	Serialize() ([]byte, error)
}
