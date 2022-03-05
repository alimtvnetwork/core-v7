package serializerinf

type ByteToJsonMustStringer interface {
	StringJsonMust(input byte) string
}
