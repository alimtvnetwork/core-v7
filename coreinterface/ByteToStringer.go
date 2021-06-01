package coreinterface

type ByteToStringer interface {
	String(input byte) string
}
