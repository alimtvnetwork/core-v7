package coreinterface

type Deserializer interface {
	Deserialize(toPtr interface{}) error
}
