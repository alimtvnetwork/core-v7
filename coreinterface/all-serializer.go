package coreinterface

type Serializer interface {
	Serialize() ([]byte, error)
}

type MustSerializer interface {
	SerializeMust() (jsonBytes []byte)
}

type AllSerializer interface {
	Serializer
	MustSerializer
}
