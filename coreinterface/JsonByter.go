package coreinterface

//goland:noinspection SpellCheckingInspection
type JsonByter interface {
	JsonBytesPtr() (jsonBytesPtr *[]byte, err error)
	JsonBytes() (jsonBytes []byte, err error)
}
