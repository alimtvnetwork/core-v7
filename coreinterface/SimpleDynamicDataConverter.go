package coreinterface

type SimpleDynamicDataConverter interface {
	// GetConvertTo convert `type` from, to
	GetConvertTo(from interface{}) (to interface{}, err error)
	// Convert convert `type` from, to
	Convert(from, to interface{}) error
	// ConvertPointers convert `*type` from, to
	ConvertPointers(from, to interface{}) error
}
