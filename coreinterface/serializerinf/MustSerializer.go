package serializerinf

type MustSerializer interface {
	SerializeMust() []byte
}
