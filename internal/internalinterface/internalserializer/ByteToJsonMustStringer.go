package internalserializer

type ByteToJsonMustStringer interface {
	StringJsonMust(input byte) string
}
