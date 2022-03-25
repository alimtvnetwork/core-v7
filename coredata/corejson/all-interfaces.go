package corejson

type bytesSerializer interface {
	Serialize() ([]byte, error)
}
