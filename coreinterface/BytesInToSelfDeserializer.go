package coreinterface

type BytesInToSelfDeserializer interface {
	Deserialize(rawBytes []byte) error
}
