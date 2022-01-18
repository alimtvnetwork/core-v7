package coreinterface

type Serializer interface {
	Serialize() ([]byte, error)
	SerializeMust() []byte
}
