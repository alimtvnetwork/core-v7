package coreinterface

type JsonByter interface {
	JsonBytesPtr() (jsonBytesPtr *[]byte, err error)
	JsonBytes() (jsonBytes []byte, err error)
}
