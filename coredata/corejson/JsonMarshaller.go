package corejson

type JsonMarshaller interface {
	// MarshalJSON
	//
	//  alias for Serialize (from any to json)
	MarshalJSON() (jsonBytes []byte, parsedErr error)
	// UnmarshalJSON
	//
	//  alias for Deserialize (from json to any)
	UnmarshalJSON(rawJsonBytes []byte) error
}
