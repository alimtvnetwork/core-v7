package coreinterface

type MustBytesInToSelfDeserializer interface {
	DeserializeMust(rawBytes []byte)
}
