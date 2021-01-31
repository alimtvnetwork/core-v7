package corejson

type JsonMarshaller interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
