package internalenuminf

type Serializer interface {
	Serialize() ([]byte, error)
}
