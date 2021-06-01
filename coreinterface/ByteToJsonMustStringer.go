package coreinterface

type ByteToJsonMustStringer interface {
	StringJsonMust(input byte) string
}
