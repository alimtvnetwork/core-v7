package internalserializer

type MustSerializer interface {
	SerializeMust() []byte
}
