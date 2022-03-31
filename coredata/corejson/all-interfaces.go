package corejson

type bytesSerializer interface {
	Serialize() ([]byte, error)
}

type bytesDeserializer interface {
	Deserialize(toPtr interface{}) error
}
