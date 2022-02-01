package coreinterface

type FieldBytesToPointerDeserializer interface {
	Deserialize(toPtr interface{}) error
}
